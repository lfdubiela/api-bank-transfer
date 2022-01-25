package env

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetWithDefaultValues(t *testing.T) {
	assert.NotEmpty(t, Get().AwsRegion, "should not be null")
	assert.Empty(t, Get().AwsSecret, "should not be empty")
	assert.Empty(t, Get().AwsKey, "should not be empty")
	assert.Empty(t, Get().AwsEndpoint, "should not be null")

	assert.NotEmpty(t, Get().DbHost, "should not be null")
	assert.NotEmpty(t, Get().DbUser, "should not be null")
	assert.NotEmpty(t, Get().DbPass, "should not be null")
	assert.NotEmpty(t, Get().DbName, "should not be null")
	assert.NotEmpty(t, Get().DbConnStr, "should not be null")
	assert.NotEmpty(t, Get().DbPool, "should not be null")

	assert.NotEmpty(t, Get().HttpTimeout, "should not be null")
	assert.NotEmpty(t, Get().Port, "should not be null")

	assert.Equal(t, "local", Get().Env)
	assert.NotEmpty(t, "info", Get().LogLevel)
}

func TestSetEnvsForTestingPurpose(t *testing.T) {
	SetEnvs(&Environment{
		Env: "prod",
	})

	assert.Equal(t, "prod", Get().Env)
}
