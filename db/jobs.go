package db

// CronJob Model
type Job struct {
	Name     string
	Schedule string
	Script   string
}
