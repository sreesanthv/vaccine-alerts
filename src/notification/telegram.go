package notification

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type TelegramNotifier struct {
	apiUrl  string
	content []string
	lock    sync.Mutex
}

func NewTelegramNotifier(token, chatId string) *TelegramNotifier {
	n := &TelegramNotifier{
		apiUrl: fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=", token, chatId),
	}
	go n.Send()

	return n
}

func (n *TelegramNotifier) Send() {
	for {
		time.Sleep(40 * time.Millisecond)
		if len(n.content) == 0 {
			continue
		}

		res, err := http.Get(n.apiUrl + url.QueryEscape(strings.Join(n.content, "\n\n")))
		if err != nil {
			log.Println("Error sending alert to telegram channel:", err)
		} else if res.StatusCode != 200 {
			buf := new(strings.Builder)
			io.Copy(buf, res.Body)
			log.Println("Error sending alert to telegram channel:", buf.String())
		}
		n.lock.Lock()
		n.content = []string{}
		n.lock.Unlock()
	}
}

func (n *TelegramNotifier) Notify(content []string) error {
	text := strings.Join(content, "\n")
	text = strings.ReplaceAll(text, "*", "")
	n.lock.Lock()
	n.content = append(n.content, text)
	n.lock.Unlock()
	return nil
}
