package main

import (
	"2022/five"
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
	_ = c

	// log.Println("day1pt1", day1Pt1(c.MustGetInput(1)))
	// log.Println("day1pt2", day1Pt2(c.MustGetInput(1)))
	// log.Println("day2pt2", two.Run(c.MustGetInput(2)))
	// log.Println(three.Run(c.MustGetInput(3)))
	// log.Println(four.Run(c.MustGetInput(4)))
	log.Println(five.Run())

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
