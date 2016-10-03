package ics

import (
	"errors"
	"strings"
)

func consumeToken(tokens *[]string, token string) error {
	if (*tokens)[0] != token {
		return errors.New("Invalid token: " + (*tokens)[0] + ", expected: " + token)
	}

	*tokens = (*tokens)[1:]
	return nil
}

func consumePrefix(tokens *[]string, token string) (ret string) {
	ret = strings.TrimPrefix((*tokens)[0], token)
	*tokens = (*tokens)[1:]
	return
}
