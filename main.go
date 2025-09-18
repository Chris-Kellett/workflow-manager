package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	bot "github.com/Chris-Kellett/workflow-manager/Bot"
	config "github.com/Chris-Kellett/workflow-manager/Config"
	logger "github.com/Chris-Kellett/workflow-manager/Logger"
	ping "github.com/Chris-Kellett/workflow-manager/Ping"
)

func main() {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
	config.Init()
	logger.Init()
	logger.Event("MAIN", "%s bot initialisation successful, Session ID: %s", config.APP_NAME, config.APP_SESSIONID)
	//TODO: Start bot services here
	go ping.Init()
	go bot.Init()
	<-osSignal
	stop()
}

func stop() {
	logger.Info("MAIN", "Quit signal received, %s bot stopping...", config.APP_NAME)
	//TODO: Stop workers here for graceful shutdown
	bot.Stop <- true
	logger.Stop <- true
	fmt.Println(config.APP_NAME + " bot has stopped...")
	os.Exit(0)
}
