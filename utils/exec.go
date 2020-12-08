package utils

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os/exec"
)

func Run(name string, args []string, dir string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	log.Debug().Msg(fmt.Sprintf("%s in %s", cmd.Args, cmd.Dir))
	return cmd.Output()
}
