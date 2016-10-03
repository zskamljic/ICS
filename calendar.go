package ics

import (
	"errors"
	"strconv"
	"strings"
)

type Calendar struct {
	Method string `json:"method"`
	Prodid string `json:"prodid"`
	Events Events `json:"events"`
}

func NewCalendar(data string) (*Calendar, error) {
	cal := &Calendar{}

	err := cal.parse(data)
	if err != nil {
		return nil, err
	}

	return cal, nil
}

func (this *Calendar) parse(data string) error {
	tokens := strings.Split(data, "\n")

	err := consumeToken(&tokens, TOKEN_CAL_START)
	if err != nil {
		return err
	}

	err = consumeToken(&tokens, TOKEN_VERSION)
	if err != nil {
		return err
	}

	ids := make(map[string]string)

	for len(tokens) > 0 {
		switch {
		case strings.HasPrefix(tokens[0], TOKEN_METHOD):
			this.Method = consumePrefix(&tokens, TOKEN_METHOD)
		case strings.HasPrefix(tokens[0], TOKEN_PRODID):
			this.Prodid = consumePrefix(&tokens, TOKEN_PRODID)
		case strings.HasPrefix(tokens[0], TOKEN_EVENT_START):
			event, err := NewEvent(&tokens)
			if err != nil {
				return err
			}

			if id, ok := ids[event.Uid]; ok {
				event.Uid = id
			} else {
				ids[event.Uid] = "ll-" + strconv.Itoa(len(ids))
				event.Uid = ids[event.Uid]
			}

			this.Events = append(this.Events, *event)
		case strings.HasPrefix(tokens[0], TOKEN_CAL_END):
			consumeToken(&tokens, TOKEN_CAL_END)
			if len(tokens) > 0 {
				return errors.New("Tokens remaining: " + strings.Join(tokens, ","))
			}
			return nil

		default:
			return errors.New("Unknown token: " + tokens[0])
		}
	}

	return nil
}
