package config

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

const (
	APP_FOLDERNAME = "workflow-manager"
	APP_NAME       = "Workflow Manager"
)

var (
	// APP_x - relate to the current running instance of the bot
	APP_ROOTDIR      string
	APP_HOSTNAME     string
	APP_SESSIONID    string
	APP_ISDEV        bool
	APP_LOGGINGLEVEL int

	// DISCORD_x - relate to interaction with discord
	DISCORD_BOTTOKEN string
	DISCORD_SESSION  *discordgo.Session

	// DB_x - relate to the database
	DB_IPADDRESS string
	DB_NAME      string
	DB_USER      string
	DB_PASSWORD  string
	DB_PORT      string

	// From here: https://gist.github.com/thomasbnt/b6f455e2c7d743b796917fa3c205f812
	EmbedColourGold  = 15844367
	EmbedColourRed   = 15548997
	EmbedColourGreen = 5763719
)

func Init() {
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		fmt.Printf("wdErr ERROR: %s\n", wdErr.Error())
		panic(wdErr)
	}

	trimIndex := strings.LastIndex(wd, APP_FOLDERNAME+`\`)
	if trimIndex != -1 {
		APP_ROOTDIR = wd[:trimIndex+len(APP_FOLDERNAME+`\`)]
	} else {
		APP_ROOTDIR = wd
	}

	loadErr := godotenv.Load(path.Join(APP_ROOTDIR, ".env"))
	if loadErr != nil {
		fmt.Printf("loadErr ERROR: %s\n", loadErr.Error())
		panic(loadErr)
	}

	envErr := parseEnvVariables()
	if envErr != nil {
		fmt.Printf("parseEnvVariables ERROR: %s\n", envErr.Error())
		panic(envErr)
	}

	currentHostName, err := os.Hostname()
	if err != nil {
		APP_HOSTNAME = "Unknown"
	} else {
		APP_HOSTNAME = currentHostName
	}

	APP_SESSIONID = uuid.New().String()
}

func parseEnvVariables() error {
	var (
		err               error
		value             string
		notFoundText      string = "could not find .env value: "
		skippingText      string = "optional .env value not provided: "
		optionalErrorText string = "error processing optional .env value: "
	)

	value = os.Getenv("APP_ISDEV")
	if value == "" {
		return fmt.Errorf("%s%s", notFoundText, "APP_ISDEV")
	} else {
		APP_ISDEV, err = strconv.ParseBool(value)
		if err != nil {
			return err
		}
	}

	DISCORD_BOTTOKEN = os.Getenv("DISCORD_BOTTOKEN")
	if DISCORD_BOTTOKEN == "" {
		return fmt.Errorf("%s%s", notFoundText, "DISCORD_BOTTOKEN")
	}

	value = os.Getenv("APP_LOGGINGLEVEL")
	if value == "" {
		fmt.Printf("%s%s\n", skippingText, "APP_LOGGINGLEVEL")
	} else {
		APP_LOGGINGLEVEL, err = strconv.Atoi(value)
		if err != nil {
			fmt.Printf("%s%s :: %s\n", optionalErrorText, "APP_LOGGINGLEVEL", err)
		}
	}

	return nil
}
