package nettest

import (
	"net"
	"testing"
)

func TestMockServer(t *testing.T) {

	mock := NewMock()
	mock.ExpectBytes([]byte("hello"), []byte("goodbye"))
	mock.ExpectBytes([]byte{0, 1, 2, 3}, []byte{3, 2, 1, 0})

	srv, _ := NewTestServer("tcp", mock.Client)
	defer srv.Close()

	c, err := net.Dial("tcp", srv.Address())
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	_, err = c.Write([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}

	res := make([]byte, 7)

	_, err = c.Read(res)
	if err != nil {
		t.Fatal(err)
	}

	if string(res) != "goodbye" {
		t.Errorf("got %s want goodbye", string(res))
	}

	_, err = c.Write([]byte{0, 1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}

	res = make([]byte, 4)
	_, err = c.Read(res)
	if err != nil {
		t.Fatal(err)
	}
	if string(res) != string([]byte{3, 2, 1, 0}) {
		t.Error("bad binary mock")
	}

}
