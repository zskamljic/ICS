package ics

import (
	"errors"
	"strings"
	"time"
)

type Event struct {
	Uid         string    `json:"uid"`
	Location    string    `json:"location"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
}

type Events []Event

func NewEvent(tokens *[]string) (event *Event, err error) {
	err = consumeToken(tokens, TOKEN_EVENT_START)
	if err != nil {
		return nil, err
	}

	event = &Event{}

	for len(*tokens) > 0 {
		switch {
		case strings.HasPrefix((*tokens)[0], TOKEN_UID):
			event.Uid = consumePrefix(tokens, TOKEN_UID)
		case strings.HasPrefix((*tokens)[0], TOKEN_LOCATION):
			event.Location = consumePrefix(tokens, TOKEN_LOCATION)
		case strings.HasPrefix((*tokens)[0], TOKEN_DTSTART):
			stamp := consumePrefix(tokens, TOKEN_DTSTART)

			event.Start, err = time.Parse("20060102T150405", stamp)
			if err != nil {
				return nil, err
			}
		case strings.HasPrefix((*tokens)[0], TOKEN_DTSTAMP):
			consumePrefix(tokens, TOKEN_DTSTAMP)
		case strings.HasPrefix((*tokens)[0], TOKEN_DTEND):
			stamp := consumePrefix(tokens, TOKEN_DTEND)

			event.End, err = time.Parse("20060102T150405", stamp)
			if err != nil {
				return nil, err
			}
		case strings.HasPrefix((*tokens)[0], TOKEN_SUMMARY):
			event.Summary = consumePrefix(tokens, TOKEN_SUMMARY)
		case strings.HasPrefix((*tokens)[0], TOKEN_DESCRIPTION):
			event.Description = consumePrefix(tokens, TOKEN_DESCRIPTION)
		case strings.HasPrefix((*tokens)[0], TOKEN_EVENT_END):
			consumeToken(tokens, TOKEN_EVENT_END)
			return event, nil
		default:
			return nil, errors.New("Unrecognized token: " + (*tokens)[0])
		}
	}

	return nil, errors.New("Ran out of tokens")
}
