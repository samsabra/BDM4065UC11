package main

import (
	"fmt"

	"github.com/pocke/BDM4065UC11"
)

func main() {
	c, err := bdm.New("/dev/ttyUSB0", 9600)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	data, err := c.Send([]byte{0xA2, 0x01})
	if err != nil {
		panic(err)
	}
	for _, v := range data {
		fmt.Printf("%X ", v)
	}
	fmt.Println()
	fmt.Println(data.String())
}
