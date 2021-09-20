package main

import (
	"iceblue/pkg/storage"
	"strconv"
	"testing"
)

func TestHelloIceBlud(t *testing.T) {
	storage.InitializeStore()
	testKey := "internal"
	testValue := "Hello, Would!"
	actual := storage.Store(testKey, testValue)
	if actual < 0 {
		t.Errorf("Expect - %d, but got - %d", 0, actual)
	}
}

func TestSimpleInsert(t *testing.T) {
	testCount := 1024
	testKey := "insertTestKey"
	testValue := "simpleValue"

	var result int
	for i := 0; i < testCount; i++ {
		result = storage.Store(testKey+strconv.Itoa(i), testValue)
		if result < 0 {
			t.Errorf("Expect - %d, but got - %d", 0, result)
		}
	}
}

func TestSimpleGetSet(t *testing.T) {
	testKey := "simpleGetSet"
	testValue := "testKey"
	_, result := storage.Get(testKey)
	if result > 0 {
		t.Errorf("Expect - %d, but got - %d", -1, result)
	}

	result = storage.Store(testKey, testValue)
	if result < 0 {
		t.Errorf("Expect - %d, but got - %d", 0, result)
	}

	value, result := storage.Get(testKey)
	if value != testValue {
		t.Errorf("Expect - %s, but got - %s", testValue, value)
	}
}

func TestSimpleUpdate(t *testing.T) {
	testKey := "simpleUpdate"
	originValue := "testValue1"
	changeValue := "testValue2"
	_, result := storage.Get(testKey)
	if result > 0 {
		t.Errorf("Expect - %d, but got - %d", -1, result)
	}

	result = storage.Store(testKey, originValue)
	if result < 0 {
		t.Errorf("Expect - %d, but got - %d", 0, result)
	}

	result = storage.Update(testKey, changeValue)
	if result < 0 {
		t.Errorf("Expect - %d, but got - %d", 0, result)
	}

	value, result := storage.Get(testKey)
	if value != changeValue {
		t.Errorf("Expect - %s, but got - %s", changeValue, value)
	}
}
