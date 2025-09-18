package datasets

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type RequestCache struct {
	StartInteraction discordgo.InteractionCreate
	Interactions     map[string]discordgo.InteractionCreate
	Values           RequestCacheValues
	Started          time.Time
}

type RequestCacheValues struct {
	String  map[string]string
	Integer map[string]int64
	Bool    map[string]bool
	User    map[string]*discordgo.User
	Channel map[string]*discordgo.Channel
	Role    map[string]*discordgo.Role
	Number  map[string]float64
}
