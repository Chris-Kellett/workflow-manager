package bot

import (
	"time"

	requests "github.com/Chris-Kellett/workflow-manager/Cache/Requests"
	commands "github.com/Chris-Kellett/workflow-manager/Commands"
	datasets "github.com/Chris-Kellett/workflow-manager/Datasets"
	helpers "github.com/Chris-Kellett/workflow-manager/Helpers"
	logger "github.com/Chris-Kellett/workflow-manager/Logger"
)

var (
	queue chan (*datasets.Request) = make(chan *datasets.Request)
	Stop  chan (interface{})       = make(chan interface{})
)

func worker() {
	logger.Info("BOT", "Worker starting...")
	for {
		select {
		case request := <-queue:
			handleItem(request)
		case <-Stop:
			logger.Info("BOT", "Worker stopping...")
			return
		}
	}
}

func handleItem(request *datasets.Request) {
	logger.Info(request.Interaction.GuildID, "Handling request: %s", request.CorrelationID)
	cmdName := request.Interaction.ApplicationCommandData().Name
	if cmd, exists := commands.Commands[cmdName]; exists {
		go func(c commands.Command, r *datasets.Request) {
			requests.Add(r.Interaction, cmdName, r.CorrelationID)
			c.Handler(r.Interaction, r.CorrelationID)
			logger.Event(request.Interaction.GuildID, "Request: %s completed after %s",
				r.CorrelationID, helpers.GetNiceCommandDuration(time.Since(r.TimeStart)))
		}(cmd, request)
	} else {
		logger.ErrorText(request.Interaction.GuildID, "Unknown Command: %s", cmdName)
	}
}
