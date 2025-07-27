package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanInput(t *testing.T) {
	result := cleanInput("hello, world")

	assert.Equal(t, "hello", result[0])
}
