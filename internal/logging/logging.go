package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger zerolog.Logger

func init() {
	logConfig := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.Kitchen,
	}

	logConfig.FormatLevel = func(i any) string {
		switch i {
		case "info":
			return "\x1b[38;5;120mINF\x1b[0m \x1b[38;5;239m>\x1b[0m"
		case "debug":
			return "\x1b[38;5;221mDBG\x1b[0m \x1b[38;5;239m>\x1b[0m"
		case "warn":
			return "\x1b[38;5;203mWRN\x1b[0m \x1b[38;5;239m>\x1b[0m"
		case "error":
			return "\033[1m\x1b[38;5;196mERR\x1b[0m\033[0m \x1b[38;5;239m>\x1b[0m"
		case "fatal":
			return "\033[1m\x1b[38;5;209mFTL\x1b[0m\033[0m \x1b[38;5;239m>\x1b[0m"
		default:
			return fmt.Sprintf("%v", i)
		}
	}

	logConfig.FormatCaller = func(i any) string {
		return fmt.Sprintf("\u001B[0m%v \u001B[38;5;239m>\u001B[0m", i)
	}

	zerolog.CallerMarshalFunc = func(_ uintptr, file string, line int) string {
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				file = file[i+1:]
				break
			}
		}
		return fmt.Sprintf("%v:%v", file, line)
	}

	logConfig.FormatFieldName = func(i any) string {
		return fmt.Sprintf("\u001B[38;5;239m%v=\u001B[0m", i)
	}

	logConfig.FormatErrFieldName = func(i any) string {
		return fmt.Sprintf("\u001B[38;5;239m%v=\u001B[0m", i)
	}

	logConfig.FormatErrFieldValue = func(i any) string {
		return fmt.Sprintf("%v", i)
	}

	log.Logger = log.Output(logConfig).With().Caller().Logger()
	Logger = log.Logger
}
