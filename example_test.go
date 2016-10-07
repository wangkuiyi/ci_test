package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentDir(t *testing.T) {
	assert.True(t, strings.HasSuffix(CurrentDir(), "src/github.com/wangkuiyi/ci_test"))
}
