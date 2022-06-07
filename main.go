package spworlds

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// spapi ...
type spapi struct {
	id, token, header string
}

type Target struct {
	Url string `json:"url"`
}

// New ...
func New(id, token string) spapi {
	base64string := string(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", id, token))))
	return (spapi{id: id, token: token, header: base64string})
}

// NewPayment ...
func (s *spapi) NewPayment(amount int, redirectUrl, webhookUrl, data string) (string, error) {
	message := map[string]string{
		"amount":      fmt.Sprintf("%v", amount),
		"redirectUrl": redirectUrl,
		"webhookUrl":  webhookUrl,
		"data":        data,
	}
	bytesRepresantation, _ := json.Marshal(message)
	client := http.Client{}
	req, _ := http.NewRequest("POST", "https://spworlds.ru/api/public/payment", bytes.NewBuffer(bytesRepresantation))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.header))
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	result := Target{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	return result.Url, nil
}

// NewTransaction ...
func (s *spapi) NewTransaction(receiver, comment string, amount int) (string, error) {
	message := map[string]string{
		"receiver": receiver,
		"amount":   fmt.Sprintf("%v", amount),
		"comment":  comment,
	}
	bytesRepresentation, _ := json.Marshal(message)
	client := http.Client{}
	req, _ := http.NewRequest("POST", "https://spworlds.ru/api/public/transactions", bytes.NewBuffer(bytesRepresentation))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.header))
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}

// getName ...
func (s *spapi) GetName(discrodId string) (string, error) {
	resp, err := http.Get("https://spworlds.ru/api/public/users/" + discrodId)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	return string(body), nil
}
