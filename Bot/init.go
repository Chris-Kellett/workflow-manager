package bot

import (
	"fmt"
	"time"

	commands "github.com/Chris-Kellett/workflow-manager/Commands"
	config "github.com/Chris-Kellett/workflow-manager/Config"
	datasets "github.com/Chris-Kellett/workflow-manager/Datasets"
	logger "github.com/Chris-Kellett/workflow-manager/Logger"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

func Init() {
	if !sessionInit() {
		logger.ErrorText("BOTINIT", "Failed to initialise discord session")
		panic("Failed to initialise discord session")
	}
	logger.Info("BOTINIT", "Discord session initialised")

	if !addHandlers() {
		logger.ErrorText("BOTINIT", "Failed to add discord event handlers")
		panic("Failed to add discord event handlers")
	}
	logger.Info("BOTINIT", "Handlers added")

	if !sessionOpen() {
		logger.ErrorText("BOTINIT", "Failed to open discord session")
		panic("Failed to open discord session")
	}

	if !commands.Init() {
		logger.ErrorText("BOTINIT", "Failed to initialise commands")
		panic("Failed to initialise commands")
	}

	logger.Info("BOTINIT", "Discord session opened")
	go worker()
}

func sessionInit() bool {
	session, err := discordgo.New("Bot " + config.DISCORD_BOTTOKEN)
	if err != nil {
		logger.Error("BOTINIT", err)
		return false
	} else if session == nil {
		logger.Error("BOTINIT", fmt.Errorf("discord session was nil"))
		return false
	}

	config.DISCORD_SESSION = session
	return true
}

func sessionOpen() bool {
	err := config.DISCORD_SESSION.Open()
	if err != nil {
		logger.Error("BOTINIT", err)
		return false
	}

	if config.DISCORD_SESSION == nil {
		logger.Error("BOTINIT", err)
		return false
	}

	return true
}

func addHandlers() bool {
	config.DISCORD_SESSION.AddHandler(newGuild)
	config.DISCORD_SESSION.AddHandler(newInteractionCreate)
	return true
}

func newInteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	queue <- &datasets.Request{
		Interaction:   interaction,
		CorrelationID: uuid.NewString(),
		TimeStart:     time.Now(),
	}
}

func newGuild(session *discordgo.Session, newGuild *discordgo.GuildCreate) {
	newGuildConnected(newGuild)
}
