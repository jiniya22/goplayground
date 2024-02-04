package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonnacci1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0, fibonacci1(-1), "fibonacci1(-1) 결과는 0이어야 합니다")
	assert.Equal(0, fibonacci1(0), "fibonacci1(0) 결과는 0이어야 합니다")
	assert.Equal(1, fibonacci1(1), "fibonacci1(1) 결과는 1이어야 합니다")
	assert.Equal(21, fibonacci1(8), "fibonacci1(8) 결과는 21이어야 합니다")
	assert.Equal(55, fibonacci1(10), "fibonacci1(10) 결과는 55이어야 합니다")
}

func TestFibonnacci2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0, fibonacci2(-1), "fibonacci2(-1) 결과는 0이어야 합니다")
	assert.Equal(0, fibonacci2(0), "fibonacci2(0) 결과는 0이어야 합니다")
	assert.Equal(1, fibonacci2(1), "fibonacci2(1) 결과는 1이어야 합니다")
	assert.Equal(21, fibonacci2(8), "fibonacci2(8) 결과는 21이어야 합니다")
	assert.Equal(55, fibonacci2(10), "fibonacci2(10) 결과는 55이어야 합니다")
}

func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}
