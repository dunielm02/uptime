package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestTelegram(t *testing.T) {
	body := map[string]string{
		"chat_id": "827317315",
		"text":    "\u203c\ufe0f The Service \"My http Service\" is down \u203c\ufe0f\n \u2705 The Service is Up\u2705",
	}

	json, _ := json.Marshal(body)

	res, err := http.Post(
		"https://api.telegram.org/bot5601115331:AAF3l9w3_1SwjMfsxXUhtmlGCwyMdBRB6T/sendMessage",
		"application/json",
		bytes.NewBuffer(json),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	b, _ := io.ReadAll(res.Body)

	fmt.Println(string(b))
}
