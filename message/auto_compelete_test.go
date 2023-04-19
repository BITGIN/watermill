package message_test

import (
	"testing"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/stretchr/testify/assert"
)

func TestEmptySearchArray(t *testing.T) {
	assert.Equal(t, 0, len(message.AutoComplete([]string{}, "searchterm")))
	assert.Equal(t, 0, len(message.AutoComplete([]string{""}, "searchterm")))
	assert.Equal(t, 0, len(message.AutoComplete([]string{"  "}, "searchterm")))
}

func TestEmptySearchTerm(t *testing.T) {
	assert.Equal(t, 0, len(message.AutoComplete([]string{"searchterm"}, "")))
	assert.Equal(t, 0, len(message.AutoComplete([]string{"searchterm"}, "  ")))
}

func TestCorrectAnswer(t *testing.T) {
	assert.Equal(t, []string{"term1"}, message.AutoComplete([]string{"term1"}, "t"))
	assert.Equal(t, []string{"term1"}, message.AutoComplete([]string{"term1"}, "te"))
	assert.Equal(t, []string{"term1"}, message.AutoComplete([]string{"term1"}, "ter"))
	assert.Equal(t, []string{}, message.AutoComplete([]string{"term1"}, "tud"))
	assert.Equal(t, []string{"term1"}, message.AutoComplete([]string{"term1"}, " ter"))
	assert.Equal(t, []string{"term1"}, message.AutoComplete([]string{"term1"}, "ter "))
}

var searchTerms = []string{
	"Kentucky",
	"North Dakota",
	"North Carolina",
	"New York",
	"New Mexico",
	"New Hampshire",
	"New Jersey",
	"Nebraska",
	"Tennessee",
}

func TestExactMatchMultiple(t *testing.T) {
	assert.Equal(t, []string{}, message.AutoComplete(searchTerms, "tud"))
	assert.Equal(t, []string{
		"Nebraska",
		"New Hampshire",
		"New Jersey",
		"New Mexico",
		"New York",
		"North Carolina",
		"North Dakota",
	}, message.AutoComplete(searchTerms, "n"))
	assert.Equal(t, []string{
		"Nebraska",
		"New Hampshire",
		"New Jersey",
		"New Mexico",
		"New York",
	}, message.AutoComplete(searchTerms, "ne"))
	assert.Equal(t, []string{"Nebraska"}, message.AutoComplete(searchTerms, "nebr"))
}

func TestFuzzyMatchMultiple(t *testing.T) {
	// city names
	assert.Equal(t, []string{}, message.AutoComplete(searchTerms, "r"))
	assert.Equal(t, []string{}, message.AutoComplete(searchTerms, "re"))
	assert.Equal(t, []string{"Kentucky", "Tennessee"}, message.AutoComplete(searchTerms, "jen"))
	assert.Equal(t, []string{
		"New Hampshire",
		"New Jersey",
		"New Mexico",
		"New York",
		"North Carolina",
		"North Dakota",
	}, message.AutoComplete(searchTerms, "now"))
}
