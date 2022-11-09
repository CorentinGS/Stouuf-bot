package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/corentings/stouuf-bot/services"
)

var CoursCommand = discordgo.ApplicationCommand{
	Name:        "cours",
	Description: "Cours command",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "add",
			Description: "add cours to an user",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user-option",
					Description: "User option",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "number-option",
					Description: "Number option",
					Required:    true,
				},
			},
		},
		{
			Name:        "show",
			Description: "show cours",
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
		{
			Name:        "remove",
			Description: "remove cours",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user-option",
					Description: "User option",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "number-option",
					Description: "Number option",
					Required:    true,
				},
			},
		},
		{
			Name:        "remove-channel",
			Description: "remove cours channel",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "channel-option",
					Description: "Channel option",
					Required:    true,
				},
			},
		},
	},
}

func CoursCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options

		coursCommand := services.GetServiceContainer().InjectCoursCommandHandler()

		var embed *discordgo.MessageEmbed
		var err error

		switch options[0].Name {
		case "add":
			embed, err = coursCommand.AddCoursCommandHandler(s, i)
		case "show":
			embed, err = coursCommand.ShowCoursCommandHandler(s, i)
		case "remove":
			embed, err = coursCommand.RemoveCoursCommandHandler(s, i)
		case "remove-channel":
			embed, err = coursCommand.RemoveCoursChannelCommandHandler(s, i)
		default:
			embed, err = coursCommand.HelpCoursCommandHandler(s, i)
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
				Flags:  discordgo.MessageFlagsEphemeral,
			},
		})
	}
}
