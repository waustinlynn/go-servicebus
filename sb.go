package servicebus

import (
	sha "crypto/sha256"
	"crypto/hmac"
	"encoding/base64"
	"fmt"
	"text/template"
	"net/http"
	"bytes"
	"strconv"
	"time"
)

type SendMessage interface {
	Send(msg string)
}

type SbConfig struct {
	Key string
	KeyType string
	Endpoint string
}

type SbMessage struct {
	Body string
	Endpoint string
	Props map[string]string
}

func (config *SbConfig) Send(msg *SbMessage) (bool, error) {
	url := config.Endpoint
	req, err := http.NewRequest("POST", url + "/" + msg.Endpoint + "/messages?timeout=60", bytes.NewReader([]byte(msg.Body)))
	if(err != nil){
		return false, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", getToken(config, msg.Endpoint))
	for propName, propValue := range msg.Props {
		req.Header.Add(propName, propValue)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if(err != nil){
		return false, err
	}
	defer resp.Body.Close()
	return true, nil
}

func getToken(config *SbConfig, sbEndpoint string) string {
	now := timeNowUnixPlusHour()
	return ShaIt(now, config.Endpoint + "/" + sbEndpoint, config.KeyType, config.Key)
}

func timeNowUnixPlusHour() string {
	return strconv.Itoa(int(time.Now().Add(time.Hour * 1).Unix()))
}

func ShaIt(now string, uri string, sasType string, secret string) string {
	encoded := template.URLQueryEscaper(uri)
	data := fmt.Sprintf("%v\n%v", encoded, now)
	h := hmac.New(sha.New, []byte(secret))
	h.Write([]byte(data))
	sha := template.URLQueryEscaper(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	token := fmt.Sprintf("SharedAccessSignature sr=%v&sig=%v&se=%v&skn=%v", encoded, sha, now, sasType)
	return token
}

