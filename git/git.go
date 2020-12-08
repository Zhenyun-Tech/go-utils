package git

import (
	"github.com/Zhenyun-Tech/go-utils/utils"
	"github.com/rs/zerolog/log"
	"strings"
)

var command = "git"

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

func Clean(path string) error {
	//清理未跟踪的文件
	args := []string{"clean", "-f"}
	err := gitRun(args, path)
	utils.PanicWhenErr(err)
	//清理已跟踪的文件
	args = []string{"reset", "--hard"}
	return gitRun(args, path)
}

func gitRun(args []string, path string) error {
	output, err := utils.Run(command, args, path)
	o := strings.TrimSpace(string(output))
	isOutNil := len(o) == 0
	if !isOutNil {
		log.Info().Msg(o)
	}
	return err
}
