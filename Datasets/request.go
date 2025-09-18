package datasets

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type Request struct {
	Interaction   *discordgo.InteractionCreate
	CorrelationID string
	TimeStart     time.Time
}
