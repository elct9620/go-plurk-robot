package server

import (
	"github.com/elct9620/go-plurk-robot/db"
	"github.com/labstack/echo"
)

// TODO(elct9620): Should improve this using global template variable
type JobPage struct {
	Jobs    []db.Job
	HideNav bool
}

func jobs(c *echo.Context) error {

	mdb, _ := getDatabase()
	jobs, err := db.GetJobs(mdb, nil, 0)

	defer mdb.Session.Close()

	if err != nil {
		return err
	}

	return c.Render(200, "jobs", JobPage{jobs, false})
}
