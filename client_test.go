package bdm

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestWrite(t *testing.T) {
	b := newRWCMock([]byte{})
	c := &Client{serial: b}
	defer c.Close()

	data := []byte{0x18, 0x01}
	_, err := c.write(data)
	if err != nil {
		t.Error(err)
	}

	expected := []byte{0xA6, 0x01, 0x00, 0x00, 0x00, 0x04, 0x01, 0x18, 0x01, 0xBB}
	got := b.Buffer.Bytes()
	if !reflect.DeepEqual(got, expected) {
		t.Error("Expected: %v, but got %v", expected, got)
	}
}

func TestRead(t *testing.T) {
	assert := func(b []byte) {
		rwc := newRWCMock(b)
		c := &Client{serial: rwc}
		defer c.Close()
		res, err := c.read()
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(res, Result(b)) {
			t.Errorf("Expected %v, but got %v", b, res)
		}
	}

	// Response the data
	assert([]byte{0x21, 0x01, 0x00, 0x00, 0x0C, 0x01, 0xA2, 0x42, 0x44, 0x4D, 0x34, 0x30, 0x36, 0x35, 0x55, 0x43, 0xD5})
	// Response the status
	assert([]byte{0x21, 0x01, 0x00, 0x00, 0x04, 0x01, 0x00, 0x03, 0x26})
}

func TestChecksum(t *testing.T) {
	sum := checksum([]byte{0xA6, 0x01, 0x00, 0x00, 0x00, 0x04, 0x01, 0x18, 0x01})
	if sum != 0xBB {
		t.Errorf("Expected: %d, but got %d", 0xBB, sum)
	}
}

var _ io.Closer = &Client{}

// Mock
type RWCMock struct {
	*bytes.Buffer
	Closed bool
}

func newRWCMock(b []byte) *RWCMock {
	return &RWCMock{Buffer: bytes.NewBuffer(b)}
}

func (m RWCMock) Close() error {
	m.Closed = true
	return nil
}

var _ io.ReadWriteCloser = &RWCMock{}
