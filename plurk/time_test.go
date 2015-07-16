package plurk

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_UnmarshalTime(t *testing.T) {

	expectedDate := "Fri, 09 Dec 1904 00:01:01 GMT"
	rawJson := `{"date_of_birth": "Fri, 09 Dec 1904 00:01:01 GMT"}`
	var user User
	json.Unmarshal([]byte(rawJson), &user)

	assert.Equal(t, expectedDate, user.Birthday.Format(timeLayout))
}
