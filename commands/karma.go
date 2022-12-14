package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/corentings/stouuf-bot/services"
)

var KarmaCommand = discordgo.ApplicationCommand{
	Name:        "karma",
	Description: "Karma main command",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "add",
			Description: "add karma",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user-option",
					Description: "User option",
					Required:    true,
				},
			},
		},
		{
			Name:        "show",
			Description: "show karma",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user-option",
					Description: "User option",
					Required:    false,
				},
			},
		},
		{
			Name:        "help",
			Description: "help",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
	},
}

func KarmaCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options

		karmaCommand := services.GetServiceContainer().InjectKarmaCommandHandler()

		var embed *discordgo.MessageEmbed
		var err error

		switch options[0].Name {
		case "add":
			embed, err = karmaCommand.AddKarmaCommandHandler(s, i)
		case "show":
			embed, err = karmaCommand.ShowKarmaCommandHandler(s, i)
		default:
			embed, err = karmaCommand.HelpKarmaCommandHandler(s, i)
		}

		if err != nil {
			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{embed},
					Flags:  discordgo.MessageFlagsEphemeral,
				},
			})

			return
		}
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		})
	}
}
