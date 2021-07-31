package client

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"br.com.github/raphasalomao/go-marvel-heroes-api/model"
)

var (
	PublicKey  string
	PrivateKey string
)

func GetCharacters(name string) (model.MarvelHeroResponse, error) {
	ts := time.Now().Unix()
	response, err := http.Get(
		fmt.Sprintf("https://gateway.marvel.com:443/v1/public/characters?name=%s&limit=10&apikey=%s&hash=%s&ts=%d", url.PathEscape(name), PublicKey, hash(ts), ts),
	)
	if err != nil {
		return model.MarvelHeroResponse{}, model.ErrorResponse{
			Timestamp: time.Now(),
			Message:   "failed to get marvel character",
		}
	}
	var heroes model.MarvelHeroResponse
	err = json.NewDecoder(response.Body).Decode(&heroes)
	if err != nil || len(heroes.Data.Results) == 0 {
		return model.MarvelHeroResponse{}, model.ErrorResponse{
			Timestamp: time.Now(),
			Message: func() string {
				if len(heroes.Data.Results) == 0 {
					return "the informed character name bring no result"
				}
				return "failed to parse marvel's response"
			}(),
		}
	}
	return heroes, nil
}

func hash(timestamp int64) string {
	key := fmt.Sprintf("%v%s%s", timestamp, PrivateKey, PublicKey)
	hash := md5.Sum([]byte(key))
	return fmt.Sprintf("%x", hash)
}
