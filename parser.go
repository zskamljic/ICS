package ics

import (
	"errors"
	"strings"
)

func consumeToken(tokens *[]string, token string) error {
	(*tokens)[0] = strings.TrimSpace((*tokens)[0])
	token = strings.TrimSpace(token)

	if (*tokens)[0] != token {
		return errors.New("Invalid token: " + (*tokens)[0] + ", expected: " + token)
	}

	*tokens = (*tokens)[1:]
	return nil
}

func consumePrefix(tokens *[]string, token string) (ret string) {
	(*tokens)[0] = strings.TrimSpace((*tokens)[0])
	token = strings.TrimSpace(token)

	ret = strings.TrimPrefix((*tokens)[0], token)
	*tokens = (*tokens)[1:]
	return
}
