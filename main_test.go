package main

import (
  "testing"
)

func TestGetMessage(t *testing.T)  {
  actual := GetMessage()
  expected := "Hello"
  if actual != expected {
    t.Errorf("Expected %s, got %s\n", expected, actual)
  }
}


