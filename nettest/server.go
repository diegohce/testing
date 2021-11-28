package nettest

import "net"

type TestServer struct {
	l  net.Listener
	fn func(net.Conn)
}

func NewTestServer(network string, fn func(net.Conn)) (*TestServer, error) {

	l, err := net.Listen(network, "127.0.0.1:0")
	if err != nil {
		return nil, err
	}

	s := TestServer{
		l:  l,
		fn: fn,
	}

	go func(l net.Listener, fn func(net.Conn)) {

		for {
			if c, err := l.Accept(); err != nil {
				return
			} else {
				fn(c)
			}
		}

	}(l, fn)

	return &s, nil
}

func (s *TestServer) Address() string {
	return s.l.Addr().String()
}

func (s *TestServer) Stop() {
	s.l.Close()
}
