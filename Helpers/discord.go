package helpers

import (
	config "github.com/Chris-Kellett/workflow-manager/Config"
	logger "github.com/Chris-Kellett/workflow-manager/Logger"
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

func SendEmbed(i *discordgo.InteractionCreate, embed *embed.Embed) {
	err := config.DISCORD_SESSION.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed.MessageEmbed},
			Flags:  discordgo.MessageFlagsEphemeral,
		},
	})

	if err != nil {
		logger.Error(i.GuildID, err)
	}
}

func SendError(i *discordgo.InteractionCreate, text string) {
	errEmbed := embed.NewEmbed()
	errEmbed.SetTitle("Error")
	errEmbed.SetDescription(text)
	errEmbed.SetColor(config.EmbedColourRed)
	err := config.DISCORD_SESSION.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{errEmbed.MessageEmbed},
			Flags:  discordgo.MessageFlagsEphemeral,
		},
	})

	if err != nil {
		logger.Error(i.GuildID, err)
	}
}
