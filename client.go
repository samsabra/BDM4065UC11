package bdm

import (
	"fmt"
	"io"

	"github.com/tarm/serial"
)

type Client struct {
	serial io.ReadWriteCloser
}

var fixedHeader = []byte{
	0xA6, // Header
	0x01, // monitor ID
	0x00, // Category
	0x00, // Page
	0x00, // Function
}

const control byte = 0x01

func New(port string, baud int) (*Client, error) {
	conf := &serial.Config{
		Name: port,
		Baud: baud,
	}
	s, err := serial.OpenPort(conf)
	if err != nil {
		return nil, err
	}

	c := &Client{
		serial: s,
	}
	return c, nil
}

func (c *Client) Send(data []byte) ([]byte, error) {
	_, err := c.write(data)
	if err != nil {
		return nil, err
	}
	resData, err := c.read()
	if err != nil {
		return nil, err
	}

	return resData, nil
}

func (c *Client) Close() error {
	return c.serial.Close()
}

func (c *Client) write(data []byte) (int, error) {
	msg := c.build(data)
	return c.serial.Write(msg)
}

// [Header] [MonitorID] [Category] [Page] [Length] [Control] [Command] [Data 0] ... [Data N] [Checksum]
// Or
// [Header] [MonitorID] [Category] [Page] [Length] [Control] [Data 0] [Status] [Checksum]
func (c *Client) read() (Result, error) {
	header := make([]byte, 6)
	_, err := io.ReadFull(c.serial, header)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, header[4]-1)
	_, err = io.ReadFull(c.serial, buf)
	if err != nil {
		return nil, err
	}

	res := Result(append(header, buf...))

	data := res.Data()
	var checksum byte
	if data0 := data[0]; data0 == 0x00 {
		// Response the status
		checksum = c.checksum(res[:len(res)-1])
	} else {
		// Response the data
		// ignore command
		checksum = c.checksum(header) ^ c.checksum(data[1:])
	}

	if checksum != res.checksum() {
		return nil, fmt.Errorf("Checksum Error. Expected: %v but received %v", checksum, res.checksum())
	}

	return res, nil
}

// [Header] [MonitorID] [Category] [Page] [Function] [Length] [Control] [Data 0] ... [Data N] [Checksum]
func (c *Client) build(data []byte) []byte {
	res := make([]byte, 0, len(data)+8)
	res = append(res, fixedHeader...)
	length := len(data) + 2
	res = append(res, byte(length))
	res = append(res, control)
	res = append(res, data...)

	checksum := c.checksum(res)
	res = append(res, checksum)

	return res
}

func (_ *Client) checksum(b []byte) byte {
	res := byte(0)
	for _, v := range b {
		res ^= v
	}
	return res
}
