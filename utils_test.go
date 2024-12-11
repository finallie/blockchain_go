package main

import (
	"bytes"
	"testing"
)

func TestNumberToBytesWithInt(t *testing.T) {
	value := int32(123)
	expected := []byte{0, 0, 0, 123}
	result := NumberToBytes(value)
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestNumberToBytesWithNegativeInt(t *testing.T) {
	value := int32(-123)
	expected := []byte{255, 255, 255, 133}
	result := NumberToBytes(value)
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestNumberToBytesWithZero(t *testing.T) {
	value := int32(0)
	expected := []byte{0, 0, 0, 0}
	result := NumberToBytes(value)
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestNumberToBytesWithLargeInt(t *testing.T) {
	value := int64(9223372036854775807)
	expected := []byte{127, 255, 255, 255, 255, 255, 255, 255}
	result := NumberToBytes(value)
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
