package adapter_computer

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(comm Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	comm.InsertIntoLightningPort()
}
