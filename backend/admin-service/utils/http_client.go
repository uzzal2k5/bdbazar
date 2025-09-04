package utils

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
)

func HttpRequest(method, url string, payload interface{}, headers map[string]string) ([]byte, error) {
    var body []byte
    if payload != nil {
        jsonBody, err := json.Marshal(payload)
        if err != nil {
            return nil, err
        }
        body = jsonBody
    }

    req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
    if err != nil {
        return nil, err
    }

    for k, v := range headers {
        req.Header.Set(k, v)
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return ioutil.ReadAll(resp.Body)
}
