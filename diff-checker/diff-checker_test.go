package main

import (
	"testing"
)

func TestIsDiff(t *testing.T) {
	/* textに差分があればtrue,なければfalse */
	text1 := "Hello World"
	text2 := "Hello World"
	text3 := "HelloWorld"

	if isDiff(text1, text2) {
		t.Fatal("faild test")
	}
	if !isDiff(text1, text3) {
		t.Fatal("faild test.")
	}
}
