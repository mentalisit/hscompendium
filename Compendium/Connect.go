package Compendium

import (
	generate2 "compendium/Compendium/generate"
	"compendium/models"
	"fmt"
)

func (c *Compendium) connect() {

	c.sendDM(fmt.Sprintf("Here is a connect code to connect your app to the %s server.", c.in.GuildName))
	code := generate2.GenerateFormattedString(c.generate())
	c.sendDM(code)
	c.sendDM("Please insert code in app\nhttps://userxinos.github.io/HadesSpace/ or https://userxinos.github.io/HadesSpaceNew")
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
		Guild: []models.Guild{models.Guild{
			URL:  c.in.GuildAvatar,
			ID:   c.in.GuildId,
			Name: c.in.GuildName,
			Icon: c.in.GuildAvatarF,
		}},
		Token: generate2.GenerateToken(),
	}
	return identity
}
