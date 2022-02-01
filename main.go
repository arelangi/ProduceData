package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	dateTicker := time.NewTicker(20 * time.Second)
	var maxUserID int64
	for {
		select {
		case <-dateTicker.C:
			maxUserID = generateUser()
		}
	}
	fmt.Println("MaxUserID ", maxUserID)
}

func generateUser() int64 {

	faker := gofakeit.New(time.Now().Local().UnixMilli())
	data := UserPayload{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://127.0.0.1:1984/user", body)
	if err != nil {
		// handle errr)
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		log.Fatal(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var userResponse UserResponse
	err = json.Unmarshal(respBody, &userResponse)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(userResponse)

	return userResponse.UserID
}

type UserPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserResponse struct {
	UserID int64 `json:"user_id"`
}
