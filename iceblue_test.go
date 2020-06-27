package iceblue

import "testing"

func TestHelloIceBlud(t *testing.T) {
	expected := "Hello, would!"
	actual := helloWould()
	if actual != expected {
		t.Errorf("Expect - %v, but got - %v", expected, actual)
	}
}
