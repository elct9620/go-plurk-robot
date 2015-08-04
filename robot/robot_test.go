package robot

import (
	"github.com/stretchr/testify/assert"
	"syscall"
	"testing"
)

func TestNew(t *testing.T) {
	robot := New()

	assert.NotNil(t, robot)
}

func TestSetupSignal(t *testing.T) {
	r := New()
	r.SetupSignal()

	go func() {
		select {
		case sign := <-r.Signal:
			assert.Equal(t, syscall.SIGTERM, sign)
		}
	}()

	r.Signal <- syscall.SIGTERM
}

func TestHandleSignal(t *testing.T) {
	robot := New()
	robot.cron.Start()
	exit := robot.HandleSignal(syscall.SIGTERM)
	assert.True(t, exit)

	robot.cron.Start()
	exit = robot.HandleSignal(syscall.SIGINT)
	assert.True(t, exit)

	robot.cron.Start()
	exit = robot.HandleSignal(syscall.SIGKILL)
	assert.True(t, exit)
}
