package v1

import (
	"encoding/json"
	"os"
)

type Resp struct {
	Data []*Info `json:"data"`
}

type Info struct {
	IType        int    `json:"itype"`
	Provider     string `json:"provider"`
	URLTemplate  string `json:"urlTemplate"`
	ClientID     string `json:"client_id"`
	RedirectUri  string `json:"redirect_uri"`
	State        string `json:"state"`
	Scope        string `json:"scope"`
	ResponseType string `json:"response_type"`
	ClientSecret string `json:"client_secret"`
}

type Config struct {
	Path string
}

func InfoFromFile(path string) []*Info {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = f.Close()
	}()

	info := &Resp{}
	err = json.NewDecoder(f).Decode(info)

	// To maintain the order all Infos according to their types
	res := make([]*Info, 10)
	for _, in := range info.Data {
		res[in.IType] = in
	}

	return res
}
