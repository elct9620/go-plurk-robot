package robot

import (
	"github.com/elct9620/go-plurk-robot/logger"
	"github.com/robfig/cron"
	"gopkg.in/mgo.v2"
	"os"
	"os/signal"
	"syscall"
)

type Robot struct {
	cron   *cron.Cron
	db     *mgo.Session
	Signal chan os.Signal
}

// Create a new robot instance
func New() *Robot {
	return &Robot{cron: cron.New(), Signal: make(chan os.Signal, 1)}
}

// Load cron jobs from database
func (r *Robot) LoadCronJobs() {
}

// Load plurk timeline task from database
func (r *Robot) LoadTasks() {

}

// Handle program stop to graceful down
func (r *Robot) SetupSignal() {
	signal.Notify(r.Signal, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
}

// Handle signal
func (r *Robot) HandleSignal(sign os.Signal) (exit bool) {
	switch sign {
	case syscall.SIGTERM:
		r.Stop()
		exit = true
	case syscall.SIGINT:
		r.Stop()
		exit = true
	case syscall.SIGKILL:
		r.Stop()
		exit = true
	}

	return exit
}

func (r *Robot) Start() {
	r.cron.Start()

	// Setup
	r.LoadCronJobs()
	r.LoadTasks()
	r.SetupSignal()

	logger.Info("Robot started!")

	for {
		sign := <-r.Signal
		if r.HandleSignal(sign) {
			logger.Info("Graceful down robot...")
			break
		}
	}

}

func (r *Robot) Stop() {
	r.cron.Stop()
	logger.Info("Robot stopped!")
}
