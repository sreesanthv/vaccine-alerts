package notification

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type TelegramNotifier struct {
	apiUrl string
}

func NewTelegramNotifier(token, chatId string) *TelegramNotifier {
	return &TelegramNotifier{
		apiUrl: fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=", token, chatId),
	}
}

func (n *TelegramNotifier) Notify(content []string) error {
	text := strings.Join(content, "\n")
	text = strings.ReplaceAll(text, "*", "")
	res, err := http.Get(n.apiUrl + url.QueryEscape(text))
	if err != nil {
		log.Println("Error sending alert to telegram channel:", err)
	} else if res.StatusCode != 200 {
		buf := new(strings.Builder)
		io.Copy(buf, res.Body)
		log.Println("Error sending alert to telegram channel:", buf.String())
	}

	return nil
}
