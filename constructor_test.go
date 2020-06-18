package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArgument_IsService(t *testing.T) {
	assert.True(t, Argument("@{ServiceName}").IsService())
	assert.False(t, Argument("${ServiceName}").IsService())
	assert.False(t, Argument("#{ServiceName}").IsService())
	assert.False(t, Argument("var").IsService())

	assert.False(t, Argument("@{ServiceName}").IsEnv())
	assert.True(t, Argument("${ServiceName}").IsEnv())
	assert.False(t, Argument("#{ServiceName}").IsEnv())
	assert.False(t, Argument("var").IsEnv())

	assert.False(t, Argument("@{ServiceName}").IsContainerValue())
	assert.False(t, Argument("${ServiceName}").IsContainerValue())
	assert.True(t, Argument("#{ServiceName}").IsContainerValue())
	assert.False(t, Argument("var").IsContainerValue())

	assert.False(t, Argument("@{ServiceName}").IsValue())
	assert.False(t, Argument("${ServiceName}").IsValue())
	assert.False(t, Argument("#{ServiceName}").IsValue())
	assert.True(t, Argument("var").IsValue())
}
