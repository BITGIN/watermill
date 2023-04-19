package message_test

import (
	"testing"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/stretchr/testify/assert"
)

func TestStringLength(t *testing.T) {
	assert.Equal(t, 5, message.Count("Hello"))
	assert.Equal(t, 0, message.Count(""))
	assert.Equal(t, 7, message.Count("rubbish"))
	assert.Equal(t, 6, message.Count("$%^&*("))
}

func TestNumberOfCharactersInANumber(t *testing.T) {
	assert.Equal(t, 3, message.CountCharactersInt(123))
	assert.Equal(t, 1, message.CountCharactersInt(0))
	assert.Equal(t, 7, message.CountCharactersInt(1234567))
	assert.Equal(t, 6, message.CountCharactersInt(123456))
}

func TestCountSubstrings(t *testing.T) {
	assert.Equal(t, 3, message.CountSubstrings("Hello World", "l"))
	assert.Equal(t, 0, message.CountSubstrings("Hello World", "z"))
	assert.Equal(t, 2, message.CountSubstrings("Hello World", "o"))
	assert.Equal(t, 1, message.CountSubstrings("Hello World", "H"))

	assert.Equal(t, 8, message.CountSubstrings("********", "*"))
	assert.Equal(t, 0, message.CountSubstrings("********", "**_"))
	assert.Equal(t, 6, message.CountSubstrings("********", "***"))
	assert.Equal(t, 0, message.CountSubstrings("********", "*********"))
	assert.Equal(t, 0, message.CountSubstrings("*-*-*-*-*", "*********"))

	assert.Equal(t, 3, message.CountSubstrings("Hey man where-where-where's your cup holder?", "where"))
	assert.Equal(t, 2, message.CountSubstrings("Hey man where-where-where's your cup holder?", "where-where"))
	assert.Equal(t, 1, message.CountSubstrings("And some Skittles", "Skittles"))
	assert.Equal(t, 0, message.CountSubstrings("And some Skittles", "chocolate"))
}
