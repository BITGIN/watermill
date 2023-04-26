package message_test

import (
	"testing"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/stretchr/testify/suite"
)

type stringsTestSuite struct {
	suite.Suite
}

func (s *stringsTestSuite) SetupSuite() {
}

func (s *stringsTestSuite) SetupTest() {
}

func (s *stringsTestSuite) TearDownTest() {
}

func (s *stringsTestSuite) TearDownSuite() {
}

func (s *stringsTestSuite) TestString() {
	s.T().Run("string length", func(t *testing.T) {
		s.Require().Equal(5, message.Count("Hello"))
		s.Require().Equal(0, message.Count(""))
		s.Require().Equal(7, message.Count("rubbish"))
		s.Require().Equal(6, message.Count("$%^&*("))
	})
}

func (s *stringsTestSuite) TestNumberOfCharactersInANumber() {
	s.T().Run("match number of characters in a integer", func(t *testing.T) {
		s.Require().Equal(3, message.CountCharactersInt(123))
		s.Require().Equal(1, message.CountCharactersInt(0))
		s.Require().Equal(7, message.CountCharactersInt(1234567))
		s.Require().Equal(6, message.CountCharactersInt(123456))
	})
}

func (s *stringsTestSuite) TestCountSubstrings() {
	s.T().Run("match substring count with single character", func(t *testing.T) {
		s.Require().Equal(3, message.CountSubstrings("Hello World", "l"))
		s.Require().Equal(0, message.CountSubstrings("Hello World", "z"))
		s.Require().Equal(2, message.CountSubstrings("Hello World", "o"))
		s.Require().Equal(1, message.CountSubstrings("Hello World", "H"))
	})

	//s.T().Run("match substring count with single vary symbols", func(t *testing.T) {
	//	s.Require().Equal(8, message.CountSubstrings("********", "*"))
	//	s.Require().Equal(0, message.CountSubstrings("********", "**_"))
	//	s.Require().Equal(6, message.CountSubstrings("********", "***"))
	//	s.Require().Equal(0, message.CountSubstrings("********", "*********"))
	//	s.Require().Equal(0, message.CountSubstrings("*-*-*-*-*", "*********"))
	//})

	s.T().Run("match substring count with single vary strings", func(t *testing.T) {
		s.Require().Equal(3, message.CountSubstrings("Hey man where-where-where's your cup holder?", "where"))
		s.Require().Equal(2, message.CountSubstrings("Hey man where-where-where's your cup holder?", "where-where"))
		s.Require().Equal(1, message.CountSubstrings("And some Skittles", "Skittles"))
		s.Require().Equal(0, message.CountSubstrings("And some Skittles", "chocolate"))
	})
}

func TestStrings(t *testing.T) {
	suite.Run(t, &stringsTestSuite{})
}
