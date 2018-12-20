package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLogLine(t *testing.T) {
	logLine := `[1518-02-05 23:52] Guard #3109 begins shift`

	el := eventLog{}

	parsedEvent, _ := el.parseEvent(logLine)

	assert.Equal(t, 3109, parsedEvent.gid)
}
