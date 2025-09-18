package logger

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	config "github.com/Chris-Kellett/workflow-manager/Config"
)

type LogRequest struct {
	GuildID string
	Type    int
	LogText string
	Stack   string
	Error   error
	LogTime time.Time
	Params  any
}

var (
	queue chan (LogRequest)  = make(chan LogRequest)
	Stop  chan (interface{}) = make(chan interface{})
)

func Init() {
	go worker()
}

func worker() {
	for {
		select {
		case log := <-queue:
			sendLog(log)
		case <-Stop:
			fmt.Println("Logger Queue stopping...")
			return
		}
	}
}

func sendLog(log LogRequest) {
	// Do we need to ignore this log?
	if config.APP_LOGGINGLEVEL > log.Type {
		return
	}

	// Format any ...any params into the LogText if present
	logText := log.LogText
	if len(log.Params.([]any)) > 0 {
		logText = fmt.Sprintf(logText, log.Params.([]any)...)
	}

	// Get the Logging Level type, to use for Colours etc.
	logType := LoggingLevels[log.Type]

	// Define the Logging info line
	infoLine := fmt.Sprintf("%v | %v", log.LogTime.Format("02/01/06 15:04:05.000"), log.Stack)
	if log.GuildID != "" {
		infoLine += " | " + log.GuildID
	}
	if config.APP_ISDEV {
		infoLine += " | " + config.APP_HOSTNAME
	}
	fmt.Printf("%v[%v] %v :: %v %v \n",
		logType.Colour.Terminal,
		strings.ToUpper(logType.Name),
		infoLine,
		logText,
		Colours["default"].Terminal)
}

func getStack() string {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return parseStackTrace(string(buf[:n]))
		}
		buf = make([]byte, len(buf)*2)
	}
}

func parseStackTrace(stack string) string {
	retVal := ""
	lines := strings.Split(stack, "\n")

	isFirst := true

	for i := 5; i < len(lines)-1; i++ {

		line := lines[i]
		if strings.Contains(line, "logger") || strings.Contains(line, "created by") || strings.Contains(line, "main.go") {
			continue
		}

		botIndex := strings.Index(line, config.APP_FOLDERNAME)
		botIndexLength := len(config.APP_FOLDERNAME)

		lastIndex := 0

		isFileLine := strings.Contains(line, " +")
		if isFileLine {
			lastIndex = strings.LastIndex(line, " +") - 1
		} else {
			lastIndex = strings.LastIndex(line, ")")
		}

		if !isFileLine {
			continue
		}

		if botIndex != -1 && lastIndex != -1 {

			appended := false
			if isFirst {
				isFirst = false
			} else {
				break
			}

			retVal += line[botIndex+botIndexLength : lastIndex+1]

			if appended {
				break
			}

		}
	}

	return "." + removeTextInParentheses(retVal)
}

func removeTextInParentheses(input string) string {
	var result strings.Builder
	inParens := false

	for _, char := range input {
		if char == '(' {
			inParens = true
			result.WriteRune(char)
		} else if char == ')' {
			inParens = false
			result.WriteRune(char)
		} else if !inParens {
			result.WriteRune(char)
		}
	}

	return result.String()
}
