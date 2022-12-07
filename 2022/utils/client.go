package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	SessionToken string
}

func (c *Client) GetInput(day int) ([]byte, error) {

	url := fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: c.SessionToken,
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Non-OK HTTP status:", resp.StatusCode)

		return nil, errors.New(string(body))
	}

	return body, nil
}

func (c *Client) MustGetInput(day int) []byte {
	input, err := c.GetInput(day)
	if err != nil {
		log.Fatalf("failed to get input for day %d: %s", day, err)
	}
	return input
}
