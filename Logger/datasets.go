package logger

const (
	LevelError = iota
	LevelWarn
	LevelEvent
	LevelInfo
	LevelDebug
)

type LoggingOptions struct {
	Name   string
	Colour Colour
}

var LoggingLevels map[int]LoggingOptions = map[int]LoggingOptions{
	LevelError: {
		Name:   "Error",
		Colour: Colours["red"],
	},
	LevelWarn: {
		Name:   "Warn",
		Colour: Colours["yellow"],
	},
	LevelEvent: {
		Name:   "Event",
		Colour: Colours["green"],
	},
	LevelInfo: {
		Name:   "Info",
		Colour: Colours["white"],
	},
	LevelDebug: {
		Name:   "Debug",
		Colour: Colours["blue"],
	},
}

type Colour struct {
	Terminal         string `json:"Terminal,omitempty"`
	Html             string `json:"Html,omitempty"`
	GraphOpaque      string `json:"GraphOpaque,omitempty"`
	GraphTransparent string `json:"GraphTransparent,omitempty"`
}

var Colours map[string]Colour = map[string]Colour{
	"default": {
		Terminal:         "\033[0m",
		Html:             "#000000",
		GraphOpaque:      "rgba(0, 0, 0, 1)",
		GraphTransparent: "rgba(0, 0, 0, 0.2)",
	},
	"white": {
		Terminal:         "\033[97m",
		Html:             "#FFFFFF",
		GraphOpaque:      "rgba(, , , 1)",
		GraphTransparent: "rgba(, , , 0.2)",
	},
	"magenta": {
		Terminal:         "\033[35m",
		Html:             "#C30CC9",
		GraphOpaque:      "rgba(195, 12, 201, 1)",
		GraphTransparent: "rgba(195, 12, 201, 0.2)",
	},
	"yellow": {
		Terminal:         "\033[33m",
		Html:             "#FAF200",
		GraphOpaque:      "rgba(250, 242, 0, 1)",
		GraphTransparent: "rgba(250, 242, 0, 0.2)",
	},
	"green": {
		Terminal:         "\033[32m",
		Html:             "#28F200",
		GraphOpaque:      "rgba(40, 242, 0, 1)",
		GraphTransparent: "rgba(40, 242, 0, 0.2)",
	},
	"red": {
		Terminal:         "\033[31m",
		Html:             "#FF9EA0",
		GraphOpaque:      "rgba(242, 0, 8, 1)",
		GraphTransparent: "rgba(242, 0, 8, 0.2)",
	},
	"blue": {
		Terminal:         "\033[34m",
		Html:             "#25B7FF",
		GraphOpaque:      "rgba(0, 0, 255, 1)",
		GraphTransparent: "rgba(0, 0, 255, 0.2)",
	},
}
