package requests

import (
	"reflect"
	"strings"
	"time"

	config "github.com/Chris-Kellett/workflow-manager/Config"
	datasets "github.com/Chris-Kellett/workflow-manager/Datasets"
	logger "github.com/Chris-Kellett/workflow-manager/Logger"
	"github.com/bwmarrin/discordgo"
)

func Add(i *discordgo.InteractionCreate, commandName string, correlationId string) {
	// Create the Interaction object
	logger.Debug(i.GuildID, "Request ID: %v, Command: %v, Storing Inputs", correlationId, commandName)
	requestsMu.RLock()
	_, exists := requests[correlationId]
	requestsMu.RUnlock()
	if exists {
		update(i, correlationId)
		return
	}

	requestsMu.Lock()
	requests[correlationId] = datasets.RequestCache{
		StartInteraction: *i,
		Interactions: map[string]discordgo.InteractionCreate{
			i.ID: *i,
		},
		Values: datasets.RequestCacheValues{
			String:  make(map[string]string),
			Integer: make(map[string]int64),
			Bool:    make(map[string]bool),
			User:    make(map[string]*discordgo.User),
			Channel: make(map[string]*discordgo.Channel),
			Role:    make(map[string]*discordgo.Role),
			Number:  make(map[string]float64),
		},
		Started: time.Now(),
	}

	// Add any Options in the Interaction to the Map Cache
	if len(i.ApplicationCommandData().Options) > 0 {
		for _, option := range i.ApplicationCommandData().Options {
			switch option.Type {
			case discordgo.ApplicationCommandOptionString:
				requests[correlationId].Values.String[option.Name] = option.StringValue()
				logger.Info(i.GuildID, "Request ID: %s, Obtained String Value for [%v]: %v", correlationId, option.Name, requests[correlationId].Values.String[option.Name])
			case discordgo.ApplicationCommandOptionInteger:
				requests[correlationId].Values.Integer[option.Name] = option.IntValue()
				logger.Info(i.GuildID, "Request ID: %s, Obtained Integer Value for [%v]: %v", correlationId, option.Name, requests[correlationId].Values.Integer[option.Name])
			case discordgo.ApplicationCommandOptionBoolean:
				requests[correlationId].Values.Bool[option.Name] = option.BoolValue()
				logger.Info(i.GuildID, "Request ID: %s, Obtained Bool Value for [%v]: %v", correlationId, option.Name, requests[correlationId].Values.Bool[option.Name])
			case discordgo.ApplicationCommandOptionUser:
				requests[correlationId].Values.User[option.Name] = option.UserValue(config.DISCORD_SESSION)
				logger.Info(i.GuildID, "Request ID: %s, Obtained User Value for [%v]: %v", correlationId, option.Name, requests[correlationId].Values.User[option.Name])
			case discordgo.ApplicationCommandOptionChannel:
				requests[correlationId].Values.Channel[option.Name] = option.ChannelValue(config.DISCORD_SESSION)
				logger.Info(i.GuildID, "Request ID: %s, Obtained Channel Value for [%v]: %v", correlationId, option.Name, requests[correlationId].Values.Channel[option.Name])
			case discordgo.ApplicationCommandOptionRole:
				requests[correlationId].Values.Role[option.Name] = option.RoleValue(config.DISCORD_SESSION, i.GuildID)
				logger.Info(i.GuildID, "Request ID: %s, Obtained Role Value for [%v]: %v", correlationId, option.Name, requests[correlationId].Values.Role[option.Name])
			case discordgo.ApplicationCommandOptionNumber:
				requests[correlationId].Values.Number[option.Name] = option.FloatValue()
				logger.Info(i.GuildID, "Request ID: %s, Obtained Number Value for [%v]: %v", correlationId, option.Name, requests[correlationId].Values.Number[option.Name])
			default:
				logger.ErrorText(i.GuildID, "AddInteraction encountered an unknown data type [%v]", option.Type.String())
			}
		}
	}
	requestsMu.Unlock()
}

func update(i *discordgo.InteractionCreate, correlationId string) {
	// Check we have the associated Interaction in the Cache
	requestsMu.Lock()
	if _, exists := requests[correlationId]; !exists {
		logger.ErrorText(i.GuildID, "Interaction Update could not find the associated CorrelationId [%v]", correlationId)
		requestsMu.Unlock()
		return
	}

	// Add the Interaction
	requests[correlationId].Interactions[i.ID] = *i

	switch i.Type {
	// Slash Commands (directly) or Autocomplete
	case discordgo.InteractionApplicationCommand, discordgo.InteractionApplicationCommandAutocomplete:
		if len(i.ApplicationCommandData().Options) > 0 {
			for _, option := range i.ApplicationCommandData().Options {
				switch option.Type {
				case discordgo.ApplicationCommandOptionString:
					requests[correlationId].Values.String[option.Name] = option.StringValue()
					logger.Info(i.GuildID, "Request ID: %s, Obtained String Value for [%v]: %v", correlationId, option.Name, requests[correlationId].Values.String[option.Name])
				case discordgo.ApplicationCommandOptionInteger:
					requests[correlationId].Values.Integer[option.Name] = option.IntValue()
					logger.Info(i.GuildID, "Request ID: %s, Obtained Integer Value for [%v]: %v", correlationId, option.Name, requests[correlationId].Values.Integer[option.Name])
				case discordgo.ApplicationCommandOptionBoolean:
					requests[correlationId].Values.Bool[option.Name] = option.BoolValue()
					logger.Info(i.GuildID, "Request ID: %s, Obtained Bool Value for [%v]: %v", correlationId, option.Name, requests[correlationId].Values.Bool[option.Name])
				case discordgo.ApplicationCommandOptionUser:
					requests[correlationId].Values.User[option.Name] = option.UserValue(config.DISCORD_SESSION)
					logger.Info(i.GuildID, "Request ID: %s, Obtained User Value for [%v]: %v", correlationId, option.Name, requests[correlationId].Values.User[option.Name])
				case discordgo.ApplicationCommandOptionChannel:
					requests[correlationId].Values.Channel[option.Name] = option.ChannelValue(config.DISCORD_SESSION)
					logger.Info(i.GuildID, "Request ID: %s, Obtained Channel Value for [%v]: %v", correlationId, option.Name, requests[correlationId].Values.Channel[option.Name])
				case discordgo.ApplicationCommandOptionRole:
					requests[correlationId].Values.Role[option.Name] = option.RoleValue(config.DISCORD_SESSION, i.GuildID)
					logger.Info(i.GuildID, "Request ID: %s, Obtained Role Value for [%v]: %v", correlationId, option.Name, requests[correlationId].Values.Role[option.Name])
				case discordgo.ApplicationCommandOptionNumber:
					requests[correlationId].Values.Number[option.Name] = option.FloatValue()
					logger.Info(i.GuildID, "Request ID: %s, Obtained Number Value for [%v]: %v", correlationId, option.Name, requests[correlationId].Values.Number[option.Name])
				default:
					logger.ErrorText(i.GuildID, "UpdateInteraction encountered an unknown CommandData data type: [%v]", option.Type.String())
				}
			}
		}

	case discordgo.InteractionModalSubmit:
		// Modal Submissions
		data := i.ModalSubmitData()

		// Loop through the components (ActionRow contains InputText components)
		for _, actionRow := range data.Components {
			if row, ok := actionRow.(*discordgo.ActionsRow); ok {
				for _, comp := range row.Components {
					if input, ok := comp.(*discordgo.TextInput); ok {
						objectId := input.CustomID
						if strings.Contains(objectId, "|") {
							objectId = strings.Split(objectId, "|")[0]
						}
						requests[correlationId].Values.String[objectId] = input.Value
						logger.Info(i.GuildID, "Interaction ID: [%s] Obtained String Value for [%s]: %s", correlationId, objectId, requests[correlationId].Values.String[objectId])
					}
				}
			}
		}
	case discordgo.InteractionMessageComponent:
		// Message Component (selects/buttons etc.)
		switch data := i.Data.(type) {
		case *discordgo.MessageComponentInteractionData:

			var objectID string
			if strings.Contains(data.CustomID, "|") {
				objectID = strings.Split(data.CustomID, "|")[0]
			} else {
				objectID = data.CustomID
			}

			if len(data.Values) > 0 {
				// Handle select menu values
				requests[correlationId].Values.String[objectID] = data.Values[0]
				logger.Info(i.GuildID, "Request ID: %s, Obtained String Value for [%v]: %v", correlationId, objectID, requests[correlationId].Values.String[objectID])
			} else {
				// Handle button interactions
				requests[correlationId].Values.Bool[objectID] = true
				logger.Info(i.GuildID, "Request ID: %s, Obtained Button Selected (Bool=True) Value for [%v]: %v", correlationId, objectID, requests[correlationId].Values.Bool[objectID])
			}
		default:
			// Handle situations where a type could not be asserted yet we can access the returned value.
			// This will always be added as a string.
			val := reflect.ValueOf(i.Data)
			if val.Kind() == reflect.Ptr && !val.IsNil() {
				val = val.Elem()
			}

			if val.Kind() == reflect.Struct {
				// Extract CustomID
				var providedObjectID string
				if field := val.FieldByName("CustomID"); field.IsValid() && field.Kind() == reflect.String {
					providedObjectID = field.String()
				}

				var objectID string
				if strings.Contains(providedObjectID, "|") {
					objectID = strings.Split(providedObjectID, "|")[0]
				} else {
					objectID = providedObjectID
				}

				// Extract Values
				if field := val.FieldByName("Values"); field.IsValid() && field.Kind() == reflect.Slice {
					for j := 0; j < field.Len(); j++ {
						value := field.Index(j).Interface()
						switch v := value.(type) {
						case string:
							requests[correlationId].Values.String[objectID] = v
							logger.Info(i.GuildID, "Request ID: %s, Obtained String Value for [%v]: %v", correlationId, objectID, requests[correlationId].Values.String[objectID])
						case int:
							requests[correlationId].Values.Integer[objectID] = int64(v)
							logger.Info(i.GuildID, "Request ID: %s, Obtained Integer Value for [%v]: %v", correlationId, objectID, requests[correlationId].Values.Integer[objectID])
						case float64:
							requests[correlationId].Values.Number[objectID] = v
							logger.Info(i.GuildID, "Request ID: %s, Obtained Number Value for [%v]: %v", correlationId, objectID, requests[correlationId].Values.Number[objectID])
						case bool:
							requests[correlationId].Values.Bool[objectID] = v
							logger.Info(i.GuildID, "Request ID: %s, Obtained Bool Value for [%v]: %v", correlationId, objectID, requests[correlationId].Values.Bool[objectID])
						case *discordgo.User:
							requests[correlationId].Values.User[objectID] = v
							logger.Info(i.GuildID, "Request ID: %s, Obtained User Value for [%v]: %v", correlationId, objectID, requests[correlationId].Values.User[objectID])
						case *discordgo.Channel:
							requests[correlationId].Values.Channel[objectID] = v
							logger.Info(i.GuildID, "Request ID: %s, Obtained Channel Value for [%v]: %v", correlationId, objectID, requests[correlationId].Values.Channel[objectID])
						case *discordgo.Role:
							requests[correlationId].Values.Role[objectID] = v
							logger.Info(i.GuildID, "Request ID: %s, Obtained Role Value for [%v]: %v", correlationId, objectID, requests[correlationId].Values.Role[objectID])
						default:
							logger.ErrorText(i.GuildID, "Request ID: %s, Unable to obtain Type Assertion for value", i.GuildID)
						}
					}
				}
			}
		}
	default:
		logger.ErrorText(i.GuildID, "UpdateInteraction encountered an unknown data type: [%v]", i.Type.String())
	}
	requestsMu.Unlock()
}

func Complete(correlationId string) {
	requestsMu.Lock()
	delete(requests, correlationId)
	requestsMu.Unlock()
}
