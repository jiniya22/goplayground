package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPower(t *testing.T) {
	result := power(5)
	if result != 25 {
		t.Errorf("power(5)의 결과는 25입니다. 하지만 결과가 %d가 나왔습니다.", result)
	}
}

func TestPower2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(25, power(5), "power(5)의 결과는 25입니다")
}
