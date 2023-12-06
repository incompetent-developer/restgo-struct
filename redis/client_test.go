package redis

import "testing"

func TestConnection(t *testing.T) {
	rd := Redis{
		Host:     "127.0.0.1",
		Port:     "6379",
		Database: 0,
	}

	if err := rd.Connection(); err != nil {
		t.Error()
	}
}
