// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func SearchExecutable(name string) (string, error) {
	if len(name) == 0 {
		return "", fmt.Errorf("error, empty name")
	}

	if path, err := exec.LookPath(name); err != nil {
		return "", err
	} else {
		return path, nil
	}

}

func RunFrondendCmd(ctx context.Context, cmdName string, env []string, stdin_msg string) (stdoutMsg, stderrMsg string, exitedCode int, e error) {

	var outMsg bytes.Buffer
	var outErr bytes.Buffer
	var cmd *exec.Cmd

	if len(cmdName) == 0 {
		e = fmt.Errorf("error, empty cmd")
		return
	}

	rootCmd := "bash"
	if path, _ := SearchExecutable(rootCmd); len(path) != 0 {
		cmd = exec.CommandContext(ctx, rootCmd, "-c", cmdName)
		goto EXE
	}

	rootCmd = "sh"
	if path, _ := SearchExecutable(rootCmd); len(path) != 0 {
		cmd = exec.CommandContext(ctx, rootCmd, "-c", cmdName)
		goto EXE
	}

	e = fmt.Errorf("error, no sh or bash installed")
	return

EXE:

	cmd.Env = append(os.Environ(), env...)

	if len(stdin_msg) != 0 {
		cmd.Stdin = strings.NewReader(stdin_msg)
	}

	cmd.Stdout = &outMsg
	cmd.Stderr = &outErr

	e = cmd.Run()
	if a := strings.TrimSpace(outMsg.String()); len(a) > 0 {
		stdoutMsg = a
	}
	if b := strings.TrimSpace(outErr.String()); len(b) > 0 {
		stderrMsg = b
	}
	exitedCode = cmd.ProcessState.ExitCode()

	return
}

func MergeMap[M map[Key]Value, Key comparable, Value any](m1 M, m2 M) M {
	m := make(M)
	for k, v := range m1 {
		m[k] = v
	}
	for k, v := range m2 {
		m[k] = v
	}
	return m

}
