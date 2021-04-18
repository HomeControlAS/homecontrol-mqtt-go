package endpoints

import (
	c "homecontrol-mqtt-go/internal/pkg/commands"
)

type OnOffEndpoint struct {
	*endpoint
}

func NewOnOffEndpoint(
	epId string,
	epName string,
	onStateChange func(ep Endpoint, cmd string, state string),
) *OnOffEndpoint {
	return &OnOffEndpoint{
		endpoint: newEndpoint(
			"pwr",
			"60",
			epId,
			epName,
			onStateChange,
			map[string]c.Command{
				c.CP: c.NewCommand(c.CP),
				c.SP: c.NewCommand(c.SP),
			}),
	}
}

func (obj *OnOffEndpoint) SendStatus() {
	obj.SendFeedbackMessage(c.SP, obj.commands[c.SP].GetState())
}
