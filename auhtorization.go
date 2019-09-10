package Siege

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GenToken(email string, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(email + ":" + password))
}

func GenerateTicket(token string) (*UbiTicket, error) {
	client := &http.Client{}

	var uri string = "https://connect.ubi.com/ubiservices/v2/profiles/sessions"

	req, _ := http.NewRequest("POST", uri, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))
	req.Header.Add("Ubi-AppId", "39baebad-39e5-4552-8c25-2c9b919064e2")
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	} else {
		body, _ := ioutil.ReadAll(res.Body)

		ticket := UbiTicket{}
		if err := json.Unmarshal(body, &ticket); err != nil {
			fmt.Println(err)
			return &ticket, err
		}

		return &ticket, nil
	}
}

func (ticket *UbiTicket) Validate(token string) (*UbiTicket, error) {

	if ticket.Expiration == "" {
		newTicket, err := GenerateTicket(token)

		if err != nil {
			return nil, err
		}

		return newTicket, nil
	}

	expiration, err := time.Parse("2006-01-02T15:04:05Z", ticket.Expiration)
	current := time.Now()

	if err != nil {
		return nil, err
	}

	if expiration.After(current) {
		return ticket, nil
	} else {
		newTicket, err := GenerateTicket(token)

		if err != nil {
			return nil, err
		}

		return newTicket, err
	}
}
