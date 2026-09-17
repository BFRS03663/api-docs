// Command hashpw prints a bcrypt hash for use as ADMIN_PASSWORD_HASH.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Fprint(os.Stderr, "Password: ")
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && line == "" {
		fmt.Fprintln(os.Stderr, "no password provided")
		os.Exit(1)
	}
	pw := strings.TrimRight(line, "\r\n")
	if len(pw) < 8 {
		fmt.Fprintln(os.Stderr, "password must be at least 8 characters")
		os.Exit(1)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		fmt.Fprintln(os.Stderr, "hash failed:", err)
		os.Exit(1)
	}
	// Single quotes keep the $ signs literal in .env files read by docker compose or a shell.
	fmt.Printf("ADMIN_PASSWORD_HASH='%s'\n", hash)
}
