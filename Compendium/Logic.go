package Compendium

import (
	"compendium/models"
	"fmt"
	"strings"
)

func (c *Compendium) logic(m models.IncomingMessage) {
	fmt.Printf("%+v\n", m)
	cutPrefix, found := strings.CutPrefix(m.Text, "%")

	if found {
		c.in = m
		switch cutPrefix {
		case "connect": //проверить айди пользователя на наличие токена
			{
				c.connect()
			}

		}
	}
}
