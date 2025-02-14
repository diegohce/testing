package nettest

import "net"

type Mock struct {
	m map[string][]byte
}

func NewMock() *Mock {
	return &Mock{m: map[string][]byte{}}
}

func (m *Mock) ExpectBytes(in []byte, out []byte) *Mock {

	m.m[string(in)] = out
	return m
}

func (m *Mock) Client(c net.Conn) {
	var out []byte
	var exists bool

	b := []byte{0}
	inBuffer := make([]byte, 0, 4096)

	for _, err := c.Read(b); err == nil; _, err = c.Read(b) {
		inBuffer = append(inBuffer, b...)

		out, exists = m.m[string(inBuffer)]
		if exists {
			c.Write(out)
			inBuffer = make([]byte, 0, 4096)
		}

	}

}
