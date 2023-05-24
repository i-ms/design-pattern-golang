package command_tv

import "testing"

func TestCommand(t *testing.T) {
	tv := &TV{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}

	onButton.press()

	offButton := &Button{
		command: offCommand,
	}

	offButton.press()
}
