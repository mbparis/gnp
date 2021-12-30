package ch03

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listener.Close() }()

	t.Logf("bound to %q", listener.Addr())
}

func TestListenerShouldFail(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:58185")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listener.Close() }()

	t.Logf("listener bound to %q", listener.Addr())

	_, err = net.Listen("tcp", "127.0.0.1:58185")
	if err == nil {
		t.Errorf("Expect listener's duplicate port to fail")
	}

}
