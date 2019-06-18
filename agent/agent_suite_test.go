package agent_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/joho/godotenv/autoload"
)

// var app *wechatwork.App

func TestAgent(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Agent Suite")

	// corpID := os.Getenv("CORP_ID")
	// corpSecret := os.Getenv("CORP_SECRET")
	// agentID, _ := strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)

	// client := wechatwork.New(corpID)
	// app = client.WithApp(corpSecret, agentID)
}
