package main

import "errors"

var (
	ErrInternalDatabase = errors.New("internal database error")
	ErrNoContext        = errors.New("no context provided")
	ErrUnauthorized     = errors.New("authorization error")
)
