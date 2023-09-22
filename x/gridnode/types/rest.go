package types

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	io "io"
	"net/http"
)

type HedgehogData struct {
	Active bool `json:"active"`
}

func IsGridnode(voterAddr string) bool {
	fmt.Println("Voter address ", voterAddr)

	// hedgehogUrl := viper.GetString("hedgehog.hedgehog_url") + "/gridspork/mint-storage/"
	hedgehogUrl := "http://127.0.0.1:5000/is-active?address="
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(hedgehogUrl + voterAddr)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if len(body) == 0 {
		return false
	}

	var data *HedgehogData

	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Is active ", data.Active)

	return data.Active
}
