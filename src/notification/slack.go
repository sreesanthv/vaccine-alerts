package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type SlackNotifier struct {
	webhookUrl  string
	isHeaderSet bool
	headers     []string
}

func NewSlackNotifier(webhookUrl string, headers []string) *SlackNotifier {
	return &SlackNotifier{
		webhookUrl: webhookUrl,
		headers:    headers,
	}
}

type slackPayload struct {
	Text string `json:"text"`
}

func (n *SlackNotifier) Notify(content []string) error {

	contentFormatted := strings.Builder{}
	if n.isHeaderSet == false {
		contentFormatted.WriteString(fmt.Sprintf("*%s*\n", time.Now().Format("02-01-2006 15:04:05 MST")))
		contentFormatted.WriteString("*")
		contentFormatted.WriteString(strings.Join(n.headers, "\t\t"))
		contentFormatted.WriteString("*")
		contentFormatted.WriteString("\n")
		n.isHeaderSet = true
	}
	contentFormatted.WriteString(strings.Join(content, "\t\t"))

	payload := &slackPayload{
		Text: contentFormatted.String(),
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
