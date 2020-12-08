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

func gitRun(args []string, path string) error {
	output, err := utils.Run(command, args, path)
	log.Info().Msg(strings.TrimSpace(string(output)))
	return err
}
