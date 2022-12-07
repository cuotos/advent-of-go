package main

import (
	"2022/two"
	"2022/utils"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	c := &utils.Client{SessionToken: mustGetEnv("SESSION_TOKEN")}

	// log.Println("day1pt1", day1Pt1(c.MustGetInput(1)))
	// log.Println("day1pt2", day1Pt2(c.MustGetInput(1)))
	// log.Println("day2pt1", day2Pt1(c.MustGetInput(2)))
	log.Println("day2pt2", two.Day2(c.MustGetInput(2)))

	return nil
}

func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic("missing env var " + key)
	} else {
		return val
	}
}
