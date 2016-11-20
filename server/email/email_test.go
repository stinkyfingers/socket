package email

import (
	"testing"
)

func TestSend(t *testing.T) {
	to := []string{"john_shenk@hotmail.com"}
	msg := []byte("this is a test")
	err := Send(to, msg)
	if err != nil {
		t.Error(err)
	}
}
