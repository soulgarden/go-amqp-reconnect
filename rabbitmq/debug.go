package rabbitmq

import "github.com/rs/zerolog/log"

// Debug debug log flag
var Debug bool

func debug(args ...interface{}) {
	if !Debug {
		return
	}
	log.Print(args...)
}

func debugf(format string, args ...interface{}) {
	if !Debug {
		return
	}
	log.Printf(format, args...)
}

func logErr(msg string, err error) {
	log.Err(err).Msg(msg)
}
