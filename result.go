package bdm

import "fmt"

// Result is result to read from display.
//                           Header                         |              Data             | Checksum
// [Header] [MonitorID] [Category] [Page] [Length] [Control] [Command] [Data 0] ... [Data N] [Checksum]
// [Header] [MonitorID] [Category] [Page] [Length] [Control] [Data 0] [Status]               [Checksum]
type Result []byte

// Header returns header.
func (r Result) Header() []byte {
	return r[:6]
}

// Data returns data.
func (r Result) Data() []byte {
	return r[6 : len(r)-1]
}

func (r Result) checksum() byte {
	return r[len(r)-1]
}

// CHeckChecksum checks checksum. Returns an error if data is incosistent.
func (r Result) CheckChecksum() error {
	sum := checksum(r[:len(r)-1])

	if sum != r.checksum() {
		return fmt.Errorf("Checksum Error. Expected: %v but received %v", sum, r.checksum())
	}
	return nil
}

// For implement fmt.Stringer
func (r Result) String() string {
	return string(r.Data()[1:])
}
