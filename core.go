package Siege

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func makeHTTPReq(uri string, ticket UbiTicket) ([]byte, error) {
	return ticket.makeHTTPReq(uri)
}
func (ticket UbiTicket) makeHTTPReq(uri string) ([]byte, error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Ubi_v1 t=%s", ticket.Ticket))
	req.Header.Add("Ubi-AppId", "39baebad-39e5-4552-8c25-2c9b919064e2")
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	} else {
		body, _ := ioutil.ReadAll(res.Body)
		return body, nil
	}
}
