package db

import (
	"gopkg.in/mgo.v2"
)

// CronJob Model
type Job struct {
	Name     string
	Schedule string
	Script   string
}

func GetJobs(db *mgo.Database, query interface{}, limit int) (result []Job, err error) {
	jobs := db.C("jobs")
	statement := jobs.Find(query)
	if limit > 0 {
		statement = statement.Limit(limit)
	}

	err = statement.All(&result)

	return
}
