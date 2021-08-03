package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, World!!!!"
	if want != hello() {
		t.Error("失敗しました。")
	}
}
