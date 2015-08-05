package robot

import (
	"github.com/ddliu/motto"
	"github.com/elct9620/go-plurk-robot/db"
	"github.com/elct9620/go-plurk-robot/logger"
	"github.com/robertkrimen/otto"
	"github.com/robfig/cron"
	"gopkg.in/mgo.v2"
	"os"
	"os/signal"
	"syscall"
)

type Robot struct {
	cron   *cron.Cron
	db     *mgo.Database
	Signal chan os.Signal
}

// Create a new robot instance
func New() *Robot {
	session, err := db.OpenSession("")
	if err != nil {
		logger.Error("Initialize robot failed, because %s", err.Error())
		return nil
	}

	return &Robot{cron: cron.New(), Signal: make(chan os.Signal, 1), db: session.DB("")}
}

// Auto generate "Motto" module let can run in Javascript VM
func GenerateScriptModule(prefix string, name string, script string) string {
	name = prefix + "_" + name
	motto.AddModule(name, func(vm *motto.Motto) (otto.Value, error) {
		return motto.CreateLoaderFromSource(script, "")(vm)
	})
	return name
}

// Alias to fast generate "Job" type module
func GenerateJobScript(name string, script string) string {
	return GenerateScriptModule("Job", name, script)
}

// Alias to fast generate "Task" type module
func GenerateTaskScript(name string, script string) string {
	return GenerateScriptModule("Task", name, script)
}

// Load cron jobs from database
func (r *Robot) LoadCronJobs() {
	jobsIt := r.db.C("jobs").Find(nil).Iter()

	job := db.Job{}
	// Read all jobs and throw into motto vm to run script
	for jobsIt.Next(&job) {
		jobName := GenerateJobScript(job.Name, job.Script)
		r.cron.AddFunc(job.Schedule, func() {
			motto.Run(jobName)
		})
	}
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

	if r.db != nil {
		r.db.Session.Close()
	}

	logger.Info("Robot stopped!")
}
