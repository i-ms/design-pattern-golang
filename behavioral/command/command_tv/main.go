package command_tv

func App() {
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
