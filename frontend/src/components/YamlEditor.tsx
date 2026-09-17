import CodeMirror, { EditorView } from "@uiw/react-codemirror";
import { yaml } from "@codemirror/lang-yaml";

const extensions = [yaml(), EditorView.lineWrapping];
const setup = { lineNumbers: true, foldGutter: true, highlightActiveLine: true, tabSize: 2, bracketMatching: true };

/** CodeMirror configured for OpenAPI YAML. Fills its parent's height. */
export default function YamlEditor({ value, onChange }: { value: string; onChange: (next: string) => void }) {
  return (
    <CodeMirror
      value={value}
      onChange={onChange}
      extensions={extensions}
      basicSetup={setup}
      height="100%"
      className="h-full text-[13px] [&_.cm-editor]:h-full [&_.cm-scroller]:font-mono"
      aria-label="YAML"
      placeholder="Paste or write an OpenAPI document here…"
    />
  );
}
