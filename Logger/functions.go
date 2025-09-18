package logger

import (
	"time"
)

// Logging functions
func Error(guildId string, err error, a ...any) {
	queue <- LogRequest{
		GuildID: guildId,
		Type:    LevelError,
		LogText: err.Error(),
		Stack:   getStack(),
		LogTime: time.Now(),
		Params:  a,
	}
}

func ErrorWithCID(guildId string, correlationId string, err error, a ...any) {
	queue <- LogRequest{
		GuildID: guildId,
		Type:    LevelError,
		LogText: correlationId + " | " + err.Error(),
		Stack:   getStack(),
		LogTime: time.Now(),
		Params:  a,
	}
}

func ErrorText(guildId string, logText string, a ...any) {
	queue <- LogRequest{
		GuildID: guildId,
		Type:    LevelError,
		LogText: logText,
		Stack:   getStack(),
		LogTime: time.Now(),
		Params:  a,
	}
}

func ErrorTextWithCID(guildId string, correlationId string, logText string, a ...any) {
	queue <- LogRequest{
		GuildID: guildId,
		Type:    LevelError,
		LogText: correlationId + " | " + logText,
		Stack:   getStack(),
		LogTime: time.Now(),
		Params:  a,
	}
}

func Event(guildId string, logText string, a ...any) {
	queue <- LogRequest{
		GuildID: guildId,
		Type:    LevelEvent,
		LogText: logText,
		Stack:   getStack(),
		LogTime: time.Now(),
		Params:  a,
	}
}

func EventWithCID(guildId string, correlationId string, logText string, a ...any) {
	queue <- LogRequest{
		GuildID: guildId,
		Type:    LevelEvent,
		LogText: correlationId + " | " + logText,
		Stack:   getStack(),
		LogTime: time.Now(),
		Params:  a,
	}
}

func Info(guildId string, logText string, a ...any) {
	queue <- LogRequest{
		GuildID: guildId,
		Type:    LevelInfo,
		LogText: logText,
		Stack:   getStack(),
		LogTime: time.Now(),
		Params:  a,
	}
}

func InfoWithCID(guildId string, correlationId string, logText string, a ...any) {
	queue <- LogRequest{
		GuildID: guildId,
		Type:    LevelInfo,
		LogText: correlationId + " | " + logText,
		Stack:   getStack(),
		LogTime: time.Now(),
		Params:  a,
	}
}

func Debug(guildId string, logText string, a ...any) {
	queue <- LogRequest{
		GuildID: guildId,
		Type:    LevelDebug,
		LogText: logText,
		Stack:   getStack(),
		LogTime: time.Now(),
		Params:  a,
	}
}

func DebugWithCID(guildId string, correlationId string, logText string, a ...any) {
	queue <- LogRequest{
		GuildID: guildId,
		Type:    LevelDebug,
		LogText: correlationId + " | " + logText,
		Stack:   getStack(),
		LogTime: time.Now(),
		Params:  a,
	}
}
