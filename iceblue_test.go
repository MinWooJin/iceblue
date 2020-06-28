package main

import "testing"

func TestHelloIceBlud(t *testing.T) {
	testKey := "internal"
	testValue := "Hello, Would!"
	actual := store(testKey, testValue)
	if actual < 0 {
		t.Errorf("Expect - %d, but got - %d", 0, actual)
	}
}
