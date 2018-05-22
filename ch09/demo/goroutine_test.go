package main

import (
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	print()
}

func TestGoPrint(t *testing.T) {
	goPrint()
	time.Sleep(1 * time.Millisecond)
}
