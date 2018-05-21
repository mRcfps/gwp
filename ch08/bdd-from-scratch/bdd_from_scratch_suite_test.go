package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBddFromScratch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BddFromScratch Suite")
}
