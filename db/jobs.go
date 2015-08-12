package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CronJob Model
type Job struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
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

func GetJobById(db *mgo.Database, id string) (result Job, err error) {
	jobs := db.C("jobs")
	// NOTE(elct9620): Confuse API, only `bson.ObjectIdHex` working here
	err = jobs.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func CreateJob(db *mgo.Database, job Job) error {
	jobs := db.C("jobs")

	return jobs.Insert(job)
}

func UpdateJob(db *mgo.Database, id string, job Job) (err error) {
	jobs := db.C("jobs")

	return jobs.UpdateId(bson.ObjectIdHex(id), job)
}

func DeleteJob(db *mgo.Database, id string) (err error) {
	jobs := db.C("jobs")

	return jobs.RemoveId(bson.ObjectIdHex(id))
}
