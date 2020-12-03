package git

import (
	"os"
	"os/exec"
)

func Clone(path, url, branch string) error {
	args := []string{"clone", "-b", branch, url}
	return gitRun(args, path)
}

func Checkout(path, branch string) error {
	args := []string{"checkout", branch}
	return gitRun(args, path)
}

func Pull(path string) error {
	args := []string{"pull"}
	return gitRun(args, path)
}

func gitRun(args []string, path string) error {
	git, err := exec.LookPath("git")
	if err != nil {
		return err
	}
	arguments := []string{git}
	for _, a := range args {
		arguments = append(arguments, a)
	}
	gitCmd := &exec.Cmd{
		Path:   git,
		Args:   arguments,
		Stderr: os.Stderr,
		Stdout: os.Stdout,
		Dir:    path,
	}
	return gitCmd.Run()
}
