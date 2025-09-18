package helpers

import (
	"fmt"
	"time"
)

func GetNiceCommandDuration(duration time.Duration) string {
	minutes := int(duration.Minutes())
	seconds := int(duration.Seconds()) % 60
	milliseconds := int(duration.Milliseconds()) % 1000
	return fmt.Sprintf("%02dm %02d.%03ds", minutes, seconds, milliseconds)
}
