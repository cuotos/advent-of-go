package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type eventType int

const (
	begin eventType = iota + 1
	sleeps
	wakes
)

const (
	dateFormat = `2006-01-02 15:04`
)

var (
	beginShiftRegex = regexp.MustCompile(`.+ Guard #(\d+) begins shift`)
	sleepsRegex     = regexp.MustCompile(`.+ falls asleep`)
	wakesRegex      = regexp.MustCompile(`.+ wakes up`)
	datetimeRe      = regexp.MustCompile(`\[(.+)\]`)
)

type event struct {
	timestamp time.Time
	gid       int
	eventType eventType
}

type eventLog struct {
	events []event
}

func parseDate(eventString string) time.Time {

	var (
		t   time.Time
		err error
	)

	if match := datetimeRe.FindStringSubmatch(eventString); len(match) > 1 && match[1] != "" {
		t, err = time.Parse(dateFormat, match[1])
		if err != nil {
			panic(err)
		}

	}

	return t
}

func parseEventFromInput(eventString string) *event {
	e := &event{}

	if match := beginShiftRegex.FindStringSubmatch(eventString); len(match) > 1 && match[1] != "" {
		e.gid, _ = strconv.Atoi(match[1])
		e.eventType = begin
	}

	switch {
	case sleepsRegex.MatchString(eventString):
		e.eventType = sleeps
		e.timestamp = parseDate(eventString)
	case wakesRegex.MatchString(eventString):
		e.eventType = wakes
		e.timestamp = parseDate(eventString)
	}

	e.timestamp = parseDate(eventString)

	return e
}

func main() {

	var rawEvents []string

	data, err := ioutil.ReadFile("2018/4-2/input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(bytes.NewReader(data))

	for s.Scan() {
		rawEvents = append(rawEvents, s.Text())
	}

	// sort the raw events into chrono order
	sort.Slice(rawEvents, func(i, j int) bool {
		var datetimeI, datetimeJ time.Time
		var err error

		if rawI := datetimeRe.FindStringSubmatch(rawEvents[i]); len(rawI) > 1 && rawI[1] != "" {
			datetimeI, err = time.Parse(dateFormat, rawI[1])
			if err != nil {
				log.Fatal(err)
			}
		}

		if rawJ := datetimeRe.FindStringSubmatch(rawEvents[j]); len(rawJ) > 1 && rawJ[1] != "" {
			datetimeJ, err = time.Parse(dateFormat, rawJ[1])
			if err != nil {
				log.Fatal(err)
			}
		}

		return datetimeI.Before(datetimeJ)
	})

	var (
		events    []*event
		lastGuard int
	)

	for _, re := range rawEvents {
		event := parseEventFromInput(re)
		if event.gid == 0 {
			event.gid = lastGuard
		} else {
			lastGuard = event.gid
		}

		events = append(events, event)
	}

	gid, min := timeMostLikelyAsleep(events)

	fmt.Println(gid * min)

}

type guardmin struct {
	guard, min int
}

func timeMostLikelyAsleep(events []*event) (int, int) {
	var guardmins = make(map[guardmin]int)

	for i, e := range events {
		if e.eventType != wakes {
			continue
		} else {
			for i := events[i-1].timestamp.Minute(); i < e.timestamp.Minute(); i++ {
				guardmins[guardmin{e.gid, i}]++
			}
		}
	}

	maxcount := 0
	gid := 0
	min := 0

	for g, count := range guardmins {
		if count > maxcount {
			maxcount = count
			gid = g.guard
			min = g.min
		}
	}

	return gid, min
}

func sleptMost(events []*event) int {
	var sleep = make(map[int]time.Duration)

	for i, e := range events {
		if e.eventType == wakes {
			amountOfSleep := e.timestamp.Sub(events[i-1].timestamp)

			sleep[e.gid] += amountOfSleep
		}
	}

	var gid int
	var maxSleep time.Duration

	for g, amount := range sleep {
		if amount > maxSleep {
			maxSleep = amount
			gid = g
		}
	}

	return gid
}
