package message_test

import (
	"testing"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/stretchr/testify/suite"
)

type autoCompleteTestSuite struct {
	suite.Suite

	searchTerms []string
}

func (s *autoCompleteTestSuite) SetupSuite() {
	s.searchTerms = []string{
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
}

func (s *autoCompleteTestSuite) SetupTest() {
}

func (s *autoCompleteTestSuite) TearDownTest() {
}

func (s *autoCompleteTestSuite) TearDownSuite() {
}

func (s *autoCompleteTestSuite) TestEmptySearchArray() {
	s.T().Run("return 0 when search empty array", func(t *testing.T) {
		s.Require().Equal(0, len(message.AutoComplete([]string{}, "searchterm")))
		s.Require().Equal(0, len(message.AutoComplete([]string{""}, "searchterm")))
		s.Require().Equal(0, len(message.AutoComplete([]string{"  "}, "searchterm")))
	})
}

func (s *autoCompleteTestSuite) TestEmptySearchTerm() {
	s.T().Run("return 0 when search with empty search term", func(t *testing.T) {
		s.Require().Equal(0, len(message.AutoComplete([]string{"searchterm"}, "")))
		s.Require().Equal(0, len(message.AutoComplete([]string{"searchterm"}, "  ")))
	})
}

func (s *autoCompleteTestSuite) TestCorrectAnswer() {
	s.T().Run("simulate type one by one", func(t *testing.T) {
		s.Require().Equal([]string{"term1"}, message.AutoComplete([]string{"term1"}, "t"))
		s.Require().Equal([]string{"term1"}, message.AutoComplete([]string{"term1"}, "te"))
		s.Require().Equal([]string{"term1"}, message.AutoComplete([]string{"term1"}, "ter"))
		s.Require().Equal([]string{}, message.AutoComplete([]string{"term1"}, "tud"))
		s.Require().Equal([]string{"term1"}, message.AutoComplete([]string{"term1"}, " ter"))
		s.Require().Equal([]string{"term1"}, message.AutoComplete([]string{"term1"}, "ter "))
	})

	s.T().Run("exact match multiple", func(t *testing.T) {
		s.Require().Equal([]string{}, message.AutoComplete(s.searchTerms, "tud"))
		s.Require().Equal([]string{
			"Nebraska",
			"New Hampshire",
			"New Jersey",
			"New Mexico",
			"New York",
			"North Carolina",
			"North Dakota",
		}, message.AutoComplete(s.searchTerms, "n"))
		s.Require().Equal([]string{
			"Nebraska",
			"New Hampshire",
			"New Jersey",
			"New Mexico",
			"New York",
		}, message.AutoComplete(s.searchTerms, "ne"))
		s.Require().Equal([]string{"Nebraska"}, message.AutoComplete(s.searchTerms, "nebr"))
	})

	s.T().Run("fuzzy match multiple", func(t *testing.T) {
		// city names
		s.Require().Equal([]string{}, message.AutoComplete(s.searchTerms, "r"))
		s.Require().Equal([]string{}, message.AutoComplete(s.searchTerms, "re"))
		s.Require().Equal([]string{"Kentucky", "Tennessee"}, message.AutoComplete(s.searchTerms, "jen"))
		s.Require().Equal([]string{
			"New Hampshire",
			"New Jersey",
			"New Mexico",
			"New York",
			"North Carolina",
			"North Dakota",
		}, message.AutoComplete(s.searchTerms, "now"))
	})
}

func TestAutoComplete(t *testing.T) {
	suite.Run(t, &autoCompleteTestSuite{})
}
