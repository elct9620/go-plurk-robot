package plurk

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_UnmarshalTime(t *testing.T) {

	expectedDate := "Fri, 09 Dec 1904 00:01:01 GMT"
	rawJson := `{"date_of_birth": "Fri, 09 Dec 1904 00:01:01 GMT"}`
	var user User
	json.Unmarshal([]byte(rawJson), &user)

	assert.Equal(t, expectedDate, user.Birthday.Format(timeLayout))
}

func Test_Now(t *testing.T) {
	now := Now()
	expectedTime := now.Time

	assert.Equal(t, expectedTime.String(), now.String())
}

func Test_PollingOffset(t *testing.T) {
	originTime, _ := time.Parse(timeLayout, "Fri, 17 Jul 2015 03:54:02 GMT")
	expectedDate := &Time{originTime}
	expectedPollingOffset := "2015-7-17T03:54:02"

	assert.Equal(t, expectedPollingOffset, expectedDate.PollingOffset())
}
