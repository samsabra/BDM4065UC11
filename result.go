package bdm

type Result []byte

func (r Result) Header() []byte {
	return r[:5]
}

func (r Result) Data() []byte {
	return r[6 : len(r)-1]
}

func (r Result) checksum() byte {
	return r[len(r)-1]
}
