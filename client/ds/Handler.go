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
			Avatar:       m.Author.AvatarURL(""),
			AvatarF:      m.Author.Avatar,
			ChannelId:    m.ChannelID,
			GuildId:      m.GuildID,
			GuildName:    g.Name,
			GuildAvatar:  g.IconURL(""),
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
func (d *Discord) GetRoles(guildId string) []models.CorpRole {
	roles, err := d.s.GuildRoles(guildId)
	if err != nil {
		d.log.ErrorErr(err)
		return nil
	}
	var guildRole []models.CorpRole
	for _, role := range roles {
		r := models.CorpRole{
			Name: role.Name,
			Id:   role.ID,
		}
		if r.Name == "@everyone" {
			r.Id = ""
		}

		guildRole = append(guildRole, r)
	}
	return guildRole
}

func (d *Discord) CheckRole(guildId, memderId, roleid string) bool {
	if roleid == "" {
		return true
	}
	member, err := d.s.GuildMember(guildId, memderId)
	if err != nil {
		d.log.ErrorErr(err)
		return false
	}
	for _, role := range member.Roles {
		if roleid == role {
			return true
		}
	}
	return false
}
