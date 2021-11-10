package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Responce struct {
	PostLink  string `json:"postLink"`
	Subreddit string `json:"subreddit"`
	Title     string `json:"title"`
	URL       string `json:"url"`
}

type MemeApi interface {
	GetMeme() (r Responce)
}

type Api struct {
	BaseUrl string
}

func (a *Api) getJson(url string, target interface{}) error {
	client := http.Client{}
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func (a Api) GetMeme() (r Responce) {
	url := a.BaseUrl + "/gimme/"
	fmt.Println(url)
	a.getJson(url, &r)
	return
}
