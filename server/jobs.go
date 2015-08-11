package server

import (
	"github.com/elct9620/go-plurk-robot/db"
	"github.com/labstack/echo"
)

// TODO(elct9620): Should improve this using global template variable
type JobsPage struct {
	Jobs    []db.Job
	HideNav bool
	WithAdd bool
}

type JobPage struct {
	*db.Job
	HideNav bool
	WithAdd bool
}

func jobs(c *echo.Context) error {

	mdb, _ := getDatabase()
	jobs, err := db.GetJobs(mdb, nil, 0)

	defer mdb.Session.Close()

	if err != nil {
		return err
	}

	return c.Render(200, "jobs", JobsPage{jobs, false, true})
}

func getJob(c *echo.Context) error {
	mdb, _ := getDatabase()
	id := c.Param("id")
	job, err := db.GetJobById(mdb, id)

	defer mdb.Session.Close()

	if err != nil {
		return c.Render(404, "job", JobPage{nil, false, false})
	}

	return c.Render(200, "job", JobPage{&job, false, false})
}

func newJob(c *echo.Context) error {
	return c.Render(200, "newJob", JobPage{nil, false, false})
}

func createJob(c *echo.Context) error {
	mdb, _ := getDatabase()

	name := c.Form("name")
	schedule := c.Form("schedule")
	script := c.Form("script")

	job := db.Job{
		Name:     name,
		Schedule: schedule,
		Script:   script,
	}

	defer mdb.Session.Close()

	err := db.CreateJob(mdb, job)

	if err != nil {
		return c.JSON(500, struct{ Error string }{err.Error()})
	}

	return c.JSON(200, job)
}

func updateJob(c *echo.Context) error {

	mdb, _ := getDatabase()

	id := c.Param("id")
	name := c.Form("name")
	schedule := c.Form("schedule")
	script := c.Form("script")

	defer mdb.Session.Close()

	err := db.UpdateJob(mdb, id, db.Job{Name: name, Schedule: schedule, Script: script})

	if err != nil {
		return c.JSON(500, struct{ Error string }{err.Error()})
	}

	return c.JSON(200, struct {
		Name     string
		Schedule string
		Script   string
	}{name, schedule, script})
}
