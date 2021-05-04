package git

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type PublishRequest struct {
	Method string  `json:"method"`
	Params *Params `json:"params"`
}

type Params struct {
	Channel string `json:"channel"`
	Data    *Data  `json:"data"`
}

type Data struct {
	Log string `json:"log"`
}

func Publish(request *PublishRequest) error {
	publishBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	url := "http://localhost:9000/api"
	method := "POST"
	payload := strings.NewReader(string(publishBody))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "apikey abc")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	log.Println("Successfully pushed to socket")

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return err
	}
	return nil
}
