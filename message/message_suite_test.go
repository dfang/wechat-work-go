package message_test

import (
	"testing"

	_ "github.com/joho/godotenv/autoload"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMessage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Message Suite")
}
