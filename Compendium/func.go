package Compendium

func (c *Compendium) sendChat(text string) {
	if c.in.Type == "ds" {
		c.ds.Send(c.in.ChannelId, text)
	} else if c.in.Type == "tg" {
		c.tg.Send(c.in.ChannelId, text)
	}
}

func (c *Compendium) sendDM(text string) {
	if c.in.Type == "ds" {
		c.ds.Send(c.in.DmChat, text)
	} else if c.in.Type == "tg" {
		c.tg.Send(c.in.DmChat, text)
	}
}
