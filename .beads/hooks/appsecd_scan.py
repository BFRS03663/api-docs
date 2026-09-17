#!/usr/bin/env python3
# >>> appsecd >>>  (managed by the AppsecD extension - do not edit)
# Scans staged/pushed changes via the AppsecD backend. stdlib only:
# no jq, no curl, no third-party packages. Fail-OPEN on any error.
import os, sys, json, base64, subprocess, time, ssl
import urllib.request, urllib.error

def _sslctx():
    # Honor a corporate / custom CA bundle so TLS-intercepting proxies
    # (common in enterprises) don't break the scan. python.org builds on
    # macOS/Windows do NOT use the OS trust store, so without this a
    # private root CA => CERTIFICATE_VERIFY_FAILED => fail-open (no scan).
    # APPSECD_CA_BUNDLE / SSL_CERT_FILE / REQUESTS_CA_BUNDLE add the CA;
    # APPSECD_SKIP_SSL_VERIFY=1 is a last-resort escape hatch (logged).
    try:
        ctx = ssl.create_default_context()
    except Exception:
        return None
    cab = os.environ.get('APPSECD_CA_BUNDLE') or os.environ.get('SSL_CERT_FILE') or os.environ.get('REQUESTS_CA_BUNDLE')
    if cab and os.path.isfile(cab):
        try:
            ctx.load_verify_locations(cab)
        except Exception:
            pass
    if os.environ.get('APPSECD_SKIP_SSL_VERIFY', '').lower() in ('1', 'true', 'yes'):
        warn('APPSECD_SKIP_SSL_VERIFY set - TLS verification DISABLED for this scan')
        ctx.check_hostname = False
        ctx.verify_mode = ssl.CERT_NONE
    return ctx

MAX_FILES = 200
MAX_BYTES = 5242880       # 5 MB default per file; overridden by config.json 'max_bytes'
POLL_TRIES = 36           # x5s = 180s (commit-time scan budget)
EXT_VERSION = '0.3.9-hook'

def warn(m):
    sys.stderr.write('[appsecd] ' + m + '\n')

def git(*args):
    try:
        r = subprocess.run(['git', *args], capture_output=True, text=True)
        return r.stdout.strip() if r.returncode == 0 else ''
    except Exception:
        return ''

def _find_cfg():
    # The extension writes config to os.homedir()/.appsecd/config.json:
    # USERPROFILE on Windows, HOME on macOS/Linux. In Git Bash $HOME is an
    # MSYS path (/c/Users/..) that native python can't open, so prefer the
    # native USERPROFILE and return the candidate where the file exists.
    bases = []
    for e in ('USERPROFILE', 'HOME'):
        v = os.environ.get(e)
        if v and v not in bases:
            bases.append(v)
    eu = os.path.expanduser('~')
    if eu and eu not in bases:
        bases.append(eu)
    for b in bases:
        p = os.path.join(b, '.appsecd', 'config.json')
        if os.path.isfile(p):
            return p
    return os.path.join(bases[0] if bases else '.', '.appsecd', 'config.json')

def main():
    hook = os.environ.get('APPSECD_HOOK_KIND') or (sys.argv[1] if len(sys.argv) > 1 else 'pre-commit')
    cfg_path = _find_cfg()
    try:
        with open(cfg_path) as f:
            cfg = json.load(f)
    except Exception:
        warn('config ' + cfg_path + ' unreadable - run \'AppsecD: Install Git hooks\' once. (commit allowed)')
        return 0
    token = cfg.get('token')
    url = (cfg.get('backend_url') or 'https://appsec.shiprocket.in').rstrip('/')
    thresh = cfg.get('block_severity') or 'high'
    try:
        max_bytes = int(cfg.get('max_bytes') or MAX_BYTES)
    except Exception:
        max_bytes = MAX_BYTES
    if not token:
        warn('no token in config - run \'AppsecD: Install Git hooks\' after sign-in. (commit allowed)')
        return 0

    if hook == 'pre-commit':
        head = git('rev-parse', 'HEAD') or 'STAGED'
        base = git('merge-base', 'HEAD', 'origin/main') or git('rev-parse', 'HEAD~1') or 'STAGED'
        names = _nstat(git('diff', '--cached', '--name-status', '-M', '--no-color'))
    else:
        base = head = ''
        for line in sys.stdin:
            p = line.split()
            if len(p) >= 4:
                head, base = p[1], p[3]
                if base == '0' * 40:
                    base = git('merge-base', 'HEAD', 'origin/main') or git('rev-parse', 'HEAD~1')
                break
        names = _nstat(git('diff', base + '...' + head, '--name-status', '-M', '--no-color')) if base and head else []

    branch = git('rev-parse', '--abbrev-ref', 'HEAD') or 'HEAD'
    remote = git('config', '--get', 'remote.origin.url')
    if not remote:
        warn('no origin remote - skipping scan (local-only repo). commit allowed')
        return 0
    if not names:
        warn('no changed files - skipping scan. commit allowed')
        return 0

    files = []
    act = {'A': 'added', 'M': 'modified', 'D': 'deleted', 'R': 'renamed', 'C': 'renamed'}
    for action, p in names[:MAX_FILES]:
        sym = action[0]
        canon = act.get(sym, 'modified')
        if sym == 'D':
            files.append({'path': p, 'content_b64': None, 'action': 'deleted', 'size_bytes': 0})
            continue
        try:
            with open(p, 'rb') as fh:
                data = fh.read()
        except Exception:
            continue
        if len(data) > max_bytes:
            files.append({'path': p, 'content_b64': None, 'action': canon, 'size_bytes': len(data)})
            continue
        files.append({'path': p, 'content_b64': base64.b64encode(data).decode(), 'action': canon, 'size_bytes': len(data)})

    body = json.dumps({
        'repo_remote_url': remote, 'repo_full_name': None,
        'base_sha': base or 'STAGED', 'head_sha': head or 'STAGED',
        'branch': branch, 'hook_type': hook, 'files': files,
        'ext_version': EXT_VERSION, 'ide_name': 'git-hook', 'ide_version': '-',
        'platform': sys.platform, 'arch': '-',
    }).encode()

    try:
        resp = _call(url, '/api/extension/scan/diff', token, 'POST', body)
        scan_id = json.loads(resp).get('scan_id')
    except urllib.error.HTTPError as e:
        warn('scan submission failed (HTTP %s): %s (commit allowed)' % (e.code, _read(e)[:300]))
        return 0
    except Exception as e:
        warn('scan submission error: %s (commit allowed)' % e)
        return 0
    if not scan_id:
        warn('scan submission returned no scan_id (commit allowed)')
        return 0
    warn('scan submitted: %s (hook=%s)' % (scan_id, hook))

    outcome, res = 'not_scanned_yet', {}
    for _ in range(POLL_TRIES):
        time.sleep(5)
        try:
            res = json.loads(_call(url, '/api/extension/scan/%s/result' % scan_id, token, 'GET'))
            outcome = res.get('outcome') or 'not_scanned_yet'
            if outcome not in ('not_scanned_yet', None):
                break
        except urllib.error.HTTPError as e:
            if e.code in (401, 403):
                warn('session expired/revoked during scan (HTTP %s) - run \'AppsecD: Sign in\' to re-authenticate. (commit allowed)' % e.code)
                return 0
            # transient 5xx / rate-limit - keep polling
        except Exception:
            pass
    if outcome in ('not_scanned_yet', None, ''):
        warn('scan did not complete in %ds - allowing commit (fail-open)' % (POLL_TRIES * 5))
        return 0

    # New-Code Gate: when the org opted into new-code-only gating, the
    # backend returns effective_* = NET-NEW counts (issues this change
    # introduced). Gate on those instead of total repo debt. Default OFF
    # -> effective_* == findings_* so older behavior is unchanged.
    gate_new = bool(res.get('gate_new_code_only'))
    if gate_new:
        c = int(res.get('effective_critical') or 0)
        h = int(res.get('effective_high') or 0)
        m = int(res.get('effective_medium') or 0)
        lo = int(res.get('effective_low') or 0)
    else:
        c = int(res.get('findings_critical') or 0)
        h = int(res.get('findings_high') or 0)
        m = int(res.get('findings_medium') or 0)
        lo = int(res.get('findings_low') or 0)
    _mode = ' [new-code-gate]' if gate_new else ''
    warn('scan %s: critical=%d high=%d medium=%d low=%d (block_severity=%s)%s' % (outcome, c, h, m, lo, thresh, _mode))
    # Print the ACTUAL findings (file -> line -> severity -> message + snippet)
    # to the terminal, grouped by file, so the dev sees exactly what & where
    # without leaving the commit. Best-effort; never affects the gate.
    try:
        _flist = json.loads(_call(url, '/api/extension/scan/%s/findings' % scan_id, token, 'GET'))
    except Exception:
        _flist = []
    if _flist:
        _order = {'critical': 0, 'high': 1, 'medium': 2, 'low': 3, 'info': 4}
        _byfile = {}
        for _f in _flist:
            _byfile.setdefault(_f.get('file') or '(unknown)', []).append(_f)
        warn('')
        warn('AppsecD findings (%d) by file:' % len(_flist))
        for _fp in sorted(_byfile):
            warn('  %s' % _fp)
            for _f in sorted(_byfile[_fp], key=lambda x: (_order.get((x.get('severity') or 'info').lower(), 9), x.get('line') or 0))[:50]:
                _sev = (_f.get('severity') or 'info').upper()
                _ln = _f.get('line') or '?'
                _rid = _f.get('rule_id') or ''
                _msg = (_f.get('message') or _rid or '').replace(chr(10), ' ').replace(chr(9), ' ')[:140]
                _nb = 'NEW ' if (_f.get('baseline_state') == 'new') else ''
                warn('    L%-5s [%-8s] %s%s' % (_ln, _sev, _nb, _msg))
                _snip = (_f.get('snippet') or '').strip()
                if _snip:
                    warn('             > %s' % _snip.split(chr(10))[0][:140])
                # Enriched context (send all the info to the terminal):
                # CWE | OWASP | confidence | rule, then a fix hint and a
                # per-finding deep link into the dashboard.
                _meta = []
                if _f.get('cwe_id'): _meta.append(str(_f['cwe_id']))
                if _f.get('owasp_category'): _meta.append(str(_f['owasp_category'])[:48])
                if _f.get('confidence'): _meta.append('conf:' + str(_f['confidence']))
                if _rid: _meta.append(_rid)
                if _meta: warn('             %s' % ' | '.join(_meta))
                _fix = (_f.get('fix_suggestion') or '').strip()
                if _fix: warn('             fix> %s' % _fix.replace(chr(10), ' ').replace(chr(9), ' ')[:160])
                _fid = _f.get('id')
                if _fid: warn('             %s/sast/scan/%s?finding=%s' % (url, scan_id, _fid))
        warn('')
    block = ((thresh == 'critical' and c > 0) or
             (thresh == 'high' and (c or h)) or
             (thresh == 'medium' and (c or h or m)) or
             (thresh == 'low' and (c or h or m or lo)))
    def _disp(d):
        # Report the gate verdict so 'blocked commits' stats are real.
        try:
            _call(url, '/api/extension/scan/%s/disposition' % scan_id, token,
                  'POST', json.dumps({'disposition': d}).encode())
        except Exception:
            pass
    if not block:
        warn('no blocking findings - commit allowed')
        _disp('allowed')
        return 0

    warn('')
    warn('=' * 52)
    warn('COMMIT BLOCKED - ' + ('NET-NEW ' if gate_new else '') + 'findings at or above ' + thresh)
    warn('Review: %s/sast/scan/%s' % (url, scan_id))
    warn('Override: set APPSECD_OVERRIDE_REASON="<at least 20 chars>" and re-commit')
    warn('=' * 52)
    reason = os.environ.get('APPSECD_OVERRIDE_REASON', '')
    if len(reason) >= 20:
        warn('override accepted - logging audit, commit allowed')
        try:
            _call(url, '/api/extension/override-log', token, 'POST', json.dumps({
                'repo_full_name': remote, 'commit_sha': head or 'STAGED', 'branch': branch,
                'hook_type': hook, 'finding_ids': [], 'severities': {}, 'reason': reason,
                'scan_id': scan_id,
            }).encode())
        except Exception:
            _disp('overridden')
        return 0
    _disp('blocked')
    return 1

def _nstat(out):
    rows = []
    for ln in out.splitlines():
        cols = ln.split('\t')
        if len(cols) < 2:
            continue
        action = cols[0]
        p = cols[2] if (action[0] in ('R', 'C') and len(cols) >= 3) else cols[1]
        if p:
            rows.append((action, p))
    return rows

def _call(url, pathpart, token, method, data=None):
    req = urllib.request.Request(url + pathpart, data=data, method=method, headers={
        'Authorization': 'AppsecD-Extension ' + token,
        'Content-Type': 'application/json',
        'User-Agent': 'appsecd-git-hook/' + EXT_VERSION,
    })
    with urllib.request.urlopen(req, timeout=20, context=_sslctx()) as r:
        return r.read().decode()

def _read(e):
    try:
        return e.read().decode()
    except Exception:
        return str(e)

if __name__ == '__main__':
    try:
        sys.exit(main())
    except Exception as e:
        sys.stderr.write('[appsecd] hook error: %s (commit allowed)\n' % e)
        sys.exit(0)
# <<< appsecd <<<
