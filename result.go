package bdm

import "fmt"

type Result []byte

func (r Result) Header() []byte {
	return r[:6]
}

func (r Result) Data() []byte {
	return r[6 : len(r)-1]
}

func (r Result) checksum() byte {
	return r[len(r)-1]
}

func (r Result) CheckChecksum() error {
	sum := checksum(r[:len(r)-1])

	if sum != r.checksum() {
		return fmt.Errorf("Checksum Error. Expected: %v but received %v", sum, r.checksum())
	}
	return nil
}
