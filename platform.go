package destiny

type Platform struct {
	*Client
	ID   int
	Base string
}

func (c *Client) XBL() *Platform {
	return &Platform{
		Client: c,
		ID:     1,
		Base:   "https://www.bungie.net/platform/destiny",
	}
}

func (c *Client) PSN() *Platform {
	return &Platform{
		Client: c,
		ID:     2,
		Base:   "https://www.bungie.net/platform/destiny",
	}
}
