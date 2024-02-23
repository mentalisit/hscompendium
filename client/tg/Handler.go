package tg

import (
	"compendium/models"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
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

		userProfilePhotos, _ := t.t.GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{UserID: m.From.ID})
		AvatarUser := t.getFileLink(userProfilePhotos.Photos[0][0].FileID)

		ThreadID := m.MessageThreadID
		if !m.IsTopicMessage && ThreadID != 0 {
			ThreadID = 0
		}
		ChatId := strconv.FormatInt(m.Chat.ID, 10) + fmt.Sprintf("/%d", ThreadID)

		i := models.IncomingMessage{
			Text:        m.Text,
			DmChat:      strconv.FormatInt(m.From.ID, 10),
			Name:        m.From.UserName,
			MentionName: "@" + m.From.UserName,
			NameId:      strconv.FormatInt(m.From.ID, 10),
			Avatar:      AvatarUser,
			ChannelId:   ChatId,
			GuildId:     strconv.FormatInt(m.Chat.ID, 10),
			GuildName:   m.Chat.Title,

			Type: "tg",
		}
		chat, err := t.t.GetChat(tgbotapi.ChatInfoConfig{ChatConfig: m.Chat.ChatConfig()})
		if err != nil {
			log.Panic(err)
		}
		if chat.Photo != nil {
			i.GuildAvatar = t.getFileLink(chat.Photo.BigFileID)
		}

		t.ChanMessage <- i

	}
}
func (t *Telegram) getFileLink(fileId string) string {
	fileconfig := tgbotapi.FileConfig{FileID: fileId}
	file, _ := t.t.GetFile(fileconfig)
	if file.FileID != "" {
		return "https://api.telegram.org/file/bot" + t.t.Token + "/" + file.FilePath
	}
	return ""
}

// отправка сообщения в телегу
func (t *Telegram) Send(chatid string, text string) int {
	a := strings.SplitN(chatid, "/", 2)
	chatId, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		t.log.ErrorErr(err)
	}
	ThreadID := 0
	if len(a) > 1 {
		ThreadID, err = strconv.Atoi(a[1])
		if err != nil {
			t.log.ErrorErr(err)
		}
	}
	m := tgbotapi.NewMessage(chatId, text)
	m.MessageThreadID = ThreadID
	tMessage, _ := t.t.Send(m)
	return tMessage.MessageID
}
