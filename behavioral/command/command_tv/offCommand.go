package command_tv

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}
