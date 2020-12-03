package utils

import (
	"github.com/rs/zerolog/log"
)

func PanicWhenErr(err error)  {
	if err!=nil {
		panic(err)
	}
}

func LogWhenErr(err error) {
	if err != nil {
		log.Error().Msg(err.Error())
	}
}
