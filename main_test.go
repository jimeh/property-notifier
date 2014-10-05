package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

type MainSuite struct{}

var _ = Suite(&MainSuite{})

// Test Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

/*
   Tests
*/
