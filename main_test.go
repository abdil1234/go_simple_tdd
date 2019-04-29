package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNumber(t *testing.T) {
	testResult := AddNumber(1, 2)
	assert.Equal(t, testResult, 3)
}

func TestCheckFalse(t *testing.T) {
	testResult := CheckFalse(12)
	assert.False(t, testResult, "hasil harus false")
}

func AddNumber(A int, B int) int {
	total := A + B
	return total
}

func CheckFalse(B int) bool {
	if B == 1 {
		return true
	}
	return false
}
