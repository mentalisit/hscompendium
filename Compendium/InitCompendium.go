package Compendium

import "C"
import (
	"compendium/client/ds"
	"compendium/client/tg"
	"compendium/models"
	"github.com/mentalisit/logger"
)

type Compendium struct {
	ds  *ds.Discord
	tg  *tg.Telegram
	log *logger.Logger
	in  models.IncomingMessage
}

func NewCompendium(cds *ds.Discord, ctg *tg.Telegram, log *logger.Logger) *Compendium {
	c := &Compendium{
		ds:  cds,
		log: log,
		tg:  ctg,
	}
	go c.inbox()
	return c
}

func (c *Compendium) inbox() {
	for {
		select {
		case in := <-c.ds.ChanMessage:
			go c.logic(in)
		case in := <-c.tg.ChanMessage:
			go c.logic(in)
		}
	}
}
