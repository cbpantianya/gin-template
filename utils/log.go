package utils

import (
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

func LogInit() zerolog.Logger {
	var writers []io.Writer
	writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	writers = append(writers, &lumberjack.Logger{
        Filename:   "./log/server.log",
        MaxSize:    10,
        MaxBackups: 500,
        MaxAge:     365,
        Compress:   true,
    })

	mw := io.MultiWriter(writers...)

	logger := zerolog.New(mw).With().Timestamp().Logger()

	logger.Info().Msg("Successfully initialized Logger")
	
	return logger
}