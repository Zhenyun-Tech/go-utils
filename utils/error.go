package utils

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
)

func PanicWhenErr(err error) {
	if err != nil {
		panic(err)
	}
}

func LogWhenErr(err error) {
	if err != nil {
		log.Error().Msg(fmt.Sprintf("exit with error: %s", err.Error()))
		os.Exit(1)
	}
}
