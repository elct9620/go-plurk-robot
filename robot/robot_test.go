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
