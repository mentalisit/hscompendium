package Compendium

import (
	generate2 "compendium/Compendium/generate"
	"compendium/models"
	"fmt"
)

func (c *Compendium) connect() {

	err := c.sendDM(fmt.Sprintf("Вот код подключения для подключения приложения к серверу %s.", c.in.GuildName))
	if err != nil && err.Error() == "Forbidden: bot can't initiate conversation with a user" {
		c.sendChat(c.in.MentionName +
			" пожалуйста отправьте мне команду старт в личных сообщениях, " +
			"я как бот не могу первый отправить вам личное сообщение. " +
			"И после повторите команду  ")
		return
	}
	c.sendChat(c.in.MentionName + ", Инструкцию отправили вам в Директ.")
	code := generate2.GenerateFormattedString(c.generate())
	err = c.sendDM(code)
	if err != nil {
		return
	}
	err = c.sendDM("Пожалуйста, вставьте код в приложение\nhttps://mentalisit.github.io/HadesSpace/")
	if err != nil {
		return
	}
}

func (c *Compendium) generate() models.Identity {
	//проверить если есть NameId то предложить соединить для двух корпораций
	identity := models.Identity{
		User: models.User{
			ID:            c.in.NameId,
			Username:      c.in.Name,
			Discriminator: "",
			Avatar:        c.in.AvatarF,
			AvatarURL:     c.in.Avatar,
		},
		Guild: models.Guild{
			URL:  c.in.GuildAvatar,
			ID:   c.in.GuildId,
			Name: c.in.GuildName,
			Icon: c.in.GuildAvatarF,
		},
		Token: generate2.GenerateToken(),
	}
	return identity
}
