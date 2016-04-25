package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestConcourseMeetupGoserv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ConcourseMeetupGoserv Suite")
}
