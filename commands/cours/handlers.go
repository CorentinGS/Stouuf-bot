package cours

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/corentings/stouuf-bot/commands/embeds"
	"github.com/corentings/stouuf-bot/interfaces"
	"github.com/corentings/stouuf-bot/utils"
)

type CoursCommand struct {
	interfaces.ICoursService
}

func (c *CoursCommand) AddCoursCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.MessageEmbed, error) {
	options := i.ApplicationCommandData().Options

	if !utils.HasPermissionsAdmin(i.Member) {
		return embeds.CreateForbiddenEmbed(s, i), fmt.Errorf("you don't have permissions to add tags")
	}

	commandOption := options[0].Options
	user := commandOption[0].UserValue(s)
	number := commandOption[1].IntValue()

	_, err := c.AddCours(user.ID, uint(number))
	if err != nil {
		return embeds.CreateErrorEmbed(s, i, err), err
	}
	return embeds.CreateSuccessEmbed(s, i, fmt.Sprintf("Added %d cours to %s", number, user.Mention())), nil
}

func (c *CoursCommand) ShowCoursCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.MessageEmbed, error) {
	options := i.ApplicationCommandData().Options

	commandOptions := options[0].Options
	var user *discordgo.User
	if len(commandOptions) > 0 {
		if commandOptions[0].Name == "user-option" {
			user = commandOptions[0].UserValue(s)
		} else {
			return embeds.CreateErrorEmbed(s, i, fmt.Errorf("an error occurred while processing the command")), fmt.Errorf("an error occurred while processing the command")
		}
	} else {
		user = i.Member.User
	}
	cours, err := c.GetCours(user.ID)
	if err != nil {
		return embeds.CreateErrorEmbed(s, i, err), err
	}
	return embeds.CreateResponseEmbed(s, i, user.Username, fmt.Sprintf("%d", cours.Value)), nil
}

func (c *CoursCommand) HelpCoursCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.MessageEmbed, error) {
	embed := embeds.CreateHelpEmbed(s, i, "cours", "Manage cousrs")
	commands := []struct {
		Name string
		Desc string
	}{
		{
			Name: "add",
			Desc: "Add cours to an user",
		},
		{
			Name: "show",
			Desc: "Show cours of a user",
		},
		{
			Name: "help",
			Desc: "Show help",
		},
		{
			Name: "remove",
			Desc: "Remove cours of an user",
		},
	}
	for _, command := range commands {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   command.Name,
			Value:  command.Desc,
			Inline: false,
		})
	}
	return embed, nil
}

func (c *CoursCommand) RemoveCoursCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.MessageEmbed, error) {
	options := i.ApplicationCommandData().Options

	if !utils.HasPermissionsAdmin(i.Member) {
		return embeds.CreateForbiddenEmbed(s, i), fmt.Errorf("you don't have permissions to add tags")
	}

	commandOption := options[0].Options
	user := commandOption[0].UserValue(s)
	number := commandOption[1].IntValue()

	_, err := c.RemoveCours(user.ID, uint(number))
	if err != nil {
		return embeds.CreateErrorEmbed(s, i, err), err
	}
	return embeds.CreateSuccessEmbed(s, i, fmt.Sprintf("Removed %d cours to %s", number, user.Mention())), nil
}

func (c *CoursCommand) RemoveCoursChannelCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.MessageEmbed, error) {
	options := i.ApplicationCommandData().Options

	if !utils.HasPermissionsAdmin(i.Member) {
		return embeds.CreateForbiddenEmbed(s, i), fmt.Errorf("you don't have permissions to add tags")
	}

	commandOption := options[0].Options
	channel := commandOption[0].ChannelValue(s)
	if channel.Type != discordgo.ChannelTypeGuildVoice {
		return embeds.CreateErrorEmbed(s, i, fmt.Errorf("channel is not a voice channel")), fmt.Errorf("channel is not a voice channel")
	}
	guild, err := s.State.Guild(i.GuildID)
	if err != nil {
		return embeds.CreateErrorEmbed(s, i, err), err
	}
	var users []string
	for _, vs := range guild.VoiceStates {
		if vs.ChannelID == channel.ID {
			users = append(users, vs.UserID)
		}
	}
	for _, user := range users {
		_, err := c.RemoveCours(user, 1)
		if err != nil {
			return embeds.CreateErrorEmbed(s, i, err), err
		}
	}
	return embeds.CreateSuccessEmbed(s, i, fmt.Sprintf("Removed %s from cours channel", channel.Mention())), nil
}
