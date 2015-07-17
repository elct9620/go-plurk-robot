package main

import (
	"bytes"
	"github.com/elct9620/go-plurk-robot/logger"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SetupLogger(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	LogFile = buffer
	RobotName = "Test Robot"

	setupLogger()
	logger.Info("Test Message")

	// Validate logger wirte into buffer and conatains spacify robot name
	assert.Contains(t, buffer.String(), RobotName)
}

func Test_SetupClient(t *testing.T) {
	Client = nil
	// Default client should nil
	assert.Nil(t, Client)

	// Client should be setup after setup
	setupClient(&cobra.Command{}, make([]string, 0))
	assert.NotNil(t, Client)

	// Client should not changed after first time setup
	oldClient := Client
	AppKey = "Test"
	setupClient(&cobra.Command{}, make([]string, 0))
	assert.Equal(t, oldClient, Client)
}
