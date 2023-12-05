package main

import (
	"github.com/nameof/sample-controller/cmd/util"
	"testing"
)

func Test_createOne(t *testing.T) {
	client := util.CreateClientset()

	beforeCount := count(client)

	createOne(client)

	afterCount := count(client)

	if beforeCount != afterCount-1 {
		t.Errorf("before: %d, after: %d", beforeCount, afterCount)
	}
}
