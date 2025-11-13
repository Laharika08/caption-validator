
package sender

import (
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "solution/internal/parser"
)

type LangResponse struct {
    Lang string `json:"lang"`
}

func SendCaptions(captions []parser.Caption, endpoint string) (bool, error) {
    if endpoint == "" {
        return false, errors.New("endpoint not provided")
    }

    var textBuilder bytes.Buffer
    for _, c := range captions {
        textBuilder.WriteString(c.Text + "
")
    }

    resp, err := http.Post(endpoint, "text/plain", bytes.NewReader(textBuilder.Bytes()))
    if err != nil {
        return false, err
    }
    defer resp.Body.Close()

    var langResp LangResponse
    if err := json.NewDecoder(resp.Body).Decode(&langResp); err != nil {
        return false, err
    }

    if langResp.Lang != "en-US" {
        return false, errors.New("language is " + langResp.Lang)
    }
    return true, nil
}
