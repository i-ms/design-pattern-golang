package adapter_computer

import "testing"

func TestAdapter(t *testing.T) {
	client := &Client{}
	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowsMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
