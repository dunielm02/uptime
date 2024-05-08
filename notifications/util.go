package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type responseHandler func([]byte) error

func aliveMessage(name string) string {
	return fmt.Sprintf("\u2705 The Service: \"%s\" is Up\u2705", name)
}

func deadMessage(name string) string {
	return fmt.Sprintf("\u203c\ufe0f The Service \"%s\" is Down \u203c\ufe0f", name)
}

func sendToWebHook(url string, body any, respHandler responseHandler) error {

	jsonFormatted, _ := json.Marshal(body)

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonFormatted))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return respHandler(read)
}
