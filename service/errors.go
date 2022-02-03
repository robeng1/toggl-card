package service

import "errors"

var (
	ErrInvalidNumberOfCardsToDraw = errors.New("number of cards to draw is larger than remaining cards in the deck")
	ErrDeckNotFound               = errors.New("deck not found")
)
