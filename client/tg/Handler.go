package tg

import (
	"compendium/models"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

func (t *Telegram) update() {
	ut := tgbotapi.NewUpdate(0)
	ut.Timeout = 60
	//получаем обновления от телеграм
	updates := t.t.GetUpdatesChan(ut)
	for update := range updates {
		if update.Message != nil {
			if update.Message.Chat.IsPrivate() { //если пишут боту в личку
				//t.messagePrivatHandler(update.Message)
			} else if update.Message.IsCommand() { //если сообщение является командой
			} else { //остальные сообщения
				t.messageHandler(update.Message)
			}
		} else if update.EditedMessage != nil {
			//t.logicMix(update.EditedMessage, true)
		} else {
		}
	}
}
func (t *Telegram) messageHandler(m *tgbotapi.Message) {

	if strings.HasPrefix(m.Text, "%") {
		ThreadID := m.MessageThreadID
		if !m.IsTopicMessage && ThreadID != 0 {
			ThreadID = 0
		}
		ChatId := strconv.FormatInt(m.Chat.ID, 10) + fmt.Sprintf("/%d", ThreadID)

		i := models.IncomingMessage{
			Text:         m.Text,
			DmChat:       strconv.FormatInt(m.From.ID, 10),
			Name:         m.From.String(),
			MentionName:  "@" + m.From.String(),
			NameId:       strconv.FormatInt(m.From.ID, 10),
			Avatar:       t.GetAvatar(m.From.ID),
			AvatarF:      "tg",
			ChannelId:    ChatId,
			GuildId:      strconv.FormatInt(m.Chat.ID, 10),
			GuildName:    m.Chat.Title,
			GuildAvatarF: "tg",

			Type: "tg",
		}
		chat, err := t.t.GetChat(tgbotapi.ChatInfoConfig{ChatConfig: m.Chat.ChatConfig()})
		if err != nil {
			t.log.Error(err.Error())
		}
		if chat.Photo != nil {
			i.GuildAvatar = t.getFileLink(chat.Photo.BigFileID)
		}

		t.ChanMessage <- i

	}
}
func (t *Telegram) messagePrivatHandler(m *tgbotapi.Message) {

	after, _ := strings.CutPrefix(m.Text, "%")
	ChatId := strconv.FormatInt(m.Chat.ID, 10)

	i := models.IncomingMessage{
		Text:         after,
		DmChat:       strconv.FormatInt(m.From.ID, 10),
		Name:         m.From.String(),
		MentionName:  "@" + m.From.String(),
		NameId:       strconv.FormatInt(m.From.ID, 10),
		Avatar:       t.GetAvatar(m.From.ID),
		AvatarF:      "tg",
		ChannelId:    ChatId,
		GuildId:      strconv.FormatInt(m.Chat.ID, 10),
		GuildName:    m.Chat.Title,
		GuildAvatarF: "tg",

		Type: "tg",
	}
	chat, err := t.t.GetChat(tgbotapi.ChatInfoConfig{ChatConfig: m.Chat.ChatConfig()})
	if err != nil {
		t.log.Error(err.Error())
	}
	if chat.Photo != nil {
		i.GuildAvatar = t.getFileLink(chat.Photo.BigFileID)
	}

	t.ChanMessage <- i

}
