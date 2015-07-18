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

func TestChecksum(t *testing.T) {
	c := &Client{}
	sum := c.checksum([]byte{0xA6, 0x01, 0x00, 0x00, 0x00, 0x04, 0x01, 0x18, 0x01})
	if sum != 0xBB {
		t.Errorf("Expected: %d, but got %d", 0xBB, sum)
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

		if !reflect.DeepEqual(res, b) {
			t.Errorf("Expected %v, but got %v", b, res)
		}
	}

	// Response the data
	assert([]byte{0x21, 0x01, 0x00, 0x00, 0x04, 0x01, 0x1D, 0x03, 0x26})
	// Response the status
	assert([]byte{0x21, 0x01, 0x00, 0x00, 0x04, 0x01, 0x00, 0x00, 0x25})
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
