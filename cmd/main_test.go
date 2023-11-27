package main

import (
	"github.com/nameof/sample-controller/pkg/client/clientset/versioned"
	"testing"
)

func Test_createOne(t *testing.T) {
	client := createClient()

	beforeCount := count(client)

	createOne(client)

	afterCount := count(client)

	if beforeCount != afterCount-1 {
		t.Errorf("before: %d, after: %d", beforeCount, afterCount)
	}
}

func Test_printall(t *testing.T) {
	type args struct {
		clientset *versioned.Clientset
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printall(tt.args.clientset)
		})
	}
}
