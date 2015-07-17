package main

import (
	"bytes"
	"github.com/elct9620/go-plurk-robot/logger"
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
