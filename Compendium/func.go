package Compendium

func (c *Compendium) sendChat(text string) {
	if c.in.Type == "ds" {
		c.ds.Send(c.in.ChannelId, text)
	} else if c.in.Type == "tg" {
		err := c.tg.Send(c.in.ChannelId, text)
		if err != nil {
			return
		}
	}
}

func (c *Compendium) sendDM(text string) error {
	if c.in.Type == "ds" {
		c.ds.Send(c.in.DmChat, text)
	} else if c.in.Type == "tg" {
		err := c.tg.Send(c.in.DmChat, text)
		if err != nil {
			return err
		}
	}
	return nil
}
