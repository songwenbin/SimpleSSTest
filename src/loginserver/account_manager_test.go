package loginserver

import (
	"testing"
)

func TestAccountManager(t *testing.T) {
	am := NewAccountManager()
	am.Add("abcde")
	info := am.Get("abcde")

	if info.ip != "localhost" {
		t.Error("ip is error")
	}
	if info.key != "foobar" {
		t.Error("key is error")
	}
	if info.port != "8888" {
		t.Error("port is error")
	}
}
