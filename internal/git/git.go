package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func ExecuteGitCommand(args ...string) (string, error) {
	cmd := exec.Command("git", args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err:= cmd.Run()
	if err != nil {
		return "", fmt.Errorf("git error: %s", stderr.String())
	}

	return strings.TrimSpace(stdout.String()), nil
}

func GetCurrentBranch() (string, error) {
	return ExecuteGitCommand("rev-parse", "--abbrev-ref", "HEAD")
}

func CreateBranch(name string) error {
	_, err := ExecuteGitCommand("checkout", "-b", name)
	return err
}

func GetStatus() (string, error) {
	return ExecuteGitCommand("status", "--short")
}

func CheckIfGitRepo() error {
	_, err := ExecuteGitCommand("rev-parse", "--git-dir")
	return err
}

func Commit(message string) error {
	_, err := ExecuteGitCommand("commit", "-m", message)
	return err
}

func GetStagedFiles() ([]string, error) {
	output, err := ExecuteGitCommand("diff", "--cached", "--name-only")
	if err != nil {
		return nil, err
	}
	
	if output == "" {
		return []string{}, nil
	}
	
	files := strings.Split(output, "\n")
	return files, nil
}