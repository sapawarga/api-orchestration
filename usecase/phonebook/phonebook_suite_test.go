package phonebook_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPhonebook(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Phonebook Suite")
}
