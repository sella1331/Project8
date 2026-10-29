package main

import "errors"

var (
	ErrEmptyList     = errors.New("list is empty")
	ErrNilNode       = errors.New("node is nil")
	ErrNodeNotFound  = errors.New("node not found in this list")
	ErrValueNotFound = errors.New("value not found")
)
