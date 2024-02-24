package ds

import (
	"compendium/models"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func (d *Discord) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Message.WebhookID != "" {
		return
	}
	if strings.HasPrefix(m.Content, "%") {
		g, err := d.s.Guild(m.GuildID)
		if err != nil {
			d.log.ErrorErr(err)
		}

		i := models.IncomingMessage{
			Text:         m.Content,
			DmChat:       d.dmChannel(m.Author.ID),
			Name:         m.Author.Username,
			MentionName:  m.Author.Mention(),
			NameId:       m.Author.ID,
			Avatar:       m.Author.AvatarURL("128"),
			AvatarF:      m.Author.Avatar,
			ChannelId:    m.ChannelID,
			GuildId:      m.GuildID,
			GuildName:    g.Name,
			GuildAvatar:  g.IconURL("128"),
			GuildAvatarF: g.Icon,
			Type:         "ds",
		}

		d.ChanMessage <- i
	}

}

func (d *Discord) dmChannel(AuthorID string) (chatidDM string) {
	create, err := d.s.UserChannelCreate(AuthorID)
	if err != nil {
		return ""
	}
	chatidDM = create.ID
	return chatidDM
}
func (d *Discord) Send(chatid, text string) (mesId string) { //отправка текста
	message, err := d.s.ChannelMessageSend(chatid, text)
	if err != nil {
		d.log.ErrorErr(err)
	}
	return message.ID
}
