package notification

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type SlackNotifier struct {
	webhookUrl  string
	isHeaderSet bool
}

func NewSlackNotifier(webhookUrl string) *SlackNotifier {
	return &SlackNotifier{
		webhookUrl: webhookUrl,
	}
}

type slackPayload struct {
	Text string `json:"text"`
}

func (n *SlackNotifier) Notify(content []string) error {

	payload := &slackPayload{
		Text: strings.Join(content, "\n"),
	}
	pBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error creating JSON payload:", err)
		return err
	}

	res, err := http.Post(n.webhookUrl, "application/json", bytes.NewBuffer(pBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		buf := new(strings.Builder)
		io.Copy(buf, res.Body)
		log.Println("Error connecting slack Webhook URL:", buf.String())
	}

	return nil
}
