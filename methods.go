package bdm

func (c *Client) SICPVersion() (string, error) {
	res, err := c.Send([]byte{GetVersion, SICPVersion})
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func (c *Client) PlatformVersion() (string, error) {
	res, err := c.Send([]byte{GetVersion, PlatformVersion})
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func (c *Client) IsPowerOn() (bool, error) {
	res, err := c.Send([]byte{GetPowerState})
	if err != nil {
		return false, err
	}

	return res.Data()[1] == PowerStateOn, nil
}

func (c *Client) PowerOn() error {
	_, err := c.Send([]byte{SetPowerState, PowerStateOn})
	return err
}

func (c *Client) PowerOff() error {
	_, err := c.Send([]byte{SetPowerState, PowerStateOff})
	return err
}
