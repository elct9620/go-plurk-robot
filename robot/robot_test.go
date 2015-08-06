package robot

import (
	"github.com/elct9620/go-plurk-robot/plurk"
	"github.com/stretchr/testify/assert"
	"syscall"
	"testing"
)

func TestSetupPlurk(t *testing.T) {
	clientExpect := plurk.New("A", "B", "C", "D")
	SetupPlurk("A", "B", "C", "D")

	assert.Equal(t, clientExpect, client)
}

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

func TestGenerateJobScript(t *testing.T) {
	jobName := GenerateJobScript("Example", "")
	assert.Equal(t, "Job_Example", jobName)
}

func TestGenerateTaskScript(t *testing.T) {
	taskName := GenerateTaskScript("Example", "")
	assert.Equal(t, "Task_Example", taskName)
}
