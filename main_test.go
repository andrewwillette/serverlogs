package main_test

import (
	"fmt"
	"testing"

	"github.com/andrewwillette/serverlogs"
)

func TestRemoteRun(t *testing.T) {
	logs, err := main.RemoteRun(main.User, main.WebsiteIP, main.PrivateKeyPath, "cat "+main.LogLocation)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("logs: %s\n", logs)
}
