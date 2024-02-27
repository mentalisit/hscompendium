package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

func (t *Telegram) getFileLink(fileId string) string {
	fileconfig := tgbotapi.FileConfig{FileID: fileId}
	file, _ := t.t.GetFile(fileconfig)
	if file.FileID != "" {
		return "https://api.telegram.org/file/bot" + t.t.Token + "/" + file.FilePath
	}
	return ""
}

// отправка сообщения в телегу
func (t *Telegram) Send(chatid string, text string) error {
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
	_, errs := t.t.Send(m)
	if errs != nil {
		return errs
	}
	return nil
}
func (t *Telegram) GetAvatar(userid int64) string {
	AvatarTG := ""
	userProfilePhotos, err := t.t.GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{UserID: userid})
	if err != nil || len(userProfilePhotos.Photos) == 0 {
		return AvatarTG
	}

	fileconfig := tgbotapi.FileConfig{FileID: userProfilePhotos.Photos[0][0].FileID}
	file, err := t.t.GetFile(fileconfig)
	if err != nil {
		t.log.ErrorErr(err)
		return AvatarTG
	}
	return "https://api.telegram.org/file/bot" + t.t.Token + "/" + file.FilePath
}
func (t *Telegram) nameOrNick(UserName, FirstName string) (name string) {
	if UserName != "" {
		name = UserName

	} else {
		name = FirstName
	}
	return name
}
