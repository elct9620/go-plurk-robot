package plurk

import "time"

// Plurk Time format cannot be parsed using builtin format
type Time struct {
	time.Time
}

const timeLayout = "Mon, 02 Jan 2006 15:04:05 MST"

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	// Trim " char
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	t.Time, err = time.Parse(timeLayout, string(b))
	return
}
