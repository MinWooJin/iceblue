package main

import "testing"

func TestHelloIceBlud(t *testing.T) {
	initializeStore()
	testKey := "internal"
	testValue := "Hello, Would!"
	actual := store(testKey, testValue)
	if actual < 0 {
		t.Errorf("Expect - %d, but got - %d", 0, actual)
	}
}

func TestSimpleGetSet(t *testing.T) {
	testKey := "simpleGetSet"
	testValue := "testKey"
	_, result := get(testKey)
	if result > 0 {
		t.Errorf("Expect - %d, but got - %d", -1, result)
	}

	result = store(testKey, testValue)
	if result < 0 {
		t.Errorf("Expect - %d, but got - %d", 0, result)
	}

	value, result := get(testKey)
	if value != testValue {
		t.Errorf("Expect - %s, but got - %s", testValue, value)
	}
}

func TestSimpleUpdate(t *testing.T) {
	testKey := "simpleUpdate"
	originValue := "testValue1"
	changeValue := "testValue2"
	_, result := get(testKey)
	if result > 0 {
		t.Errorf("Expect - %d, but got - %d", -1, result)
	}

	result = store(testKey, originValue)
	if result < 0 {
		t.Errorf("Expect - %d, but got - %d", 0, result)
	}

	result = update(testKey, changeValue)
	if result < 0 {
		t.Errorf("Expect - %d, but got - %d", 0, result)
	}

	value, result := get(testKey)
	if value != changeValue {
		t.Errorf("Expect - %s, but got - %s", changeValue, value)
	}
}
