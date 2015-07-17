package plurk

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func Test_PollingGetPlurks(t *testing.T) {
	testPath := "test/get_plurks.json"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := ioutil.ReadFile(testPath)
		fmt.Fprintln(w, string(jsonData))
	}))

	defer server.Close()

	plurkClient := &PlurkClient{credential: credential, ApiBase: server.URL}
	polling := plurkClient.GetPolling()
	res, err := polling.GetPlurks(Now(), 1)

	if err != nil {
		assert.Error(t, err)
	}

	expectedContent := "Test Data"
	expectedUserName := "Tester"

	assert.NotNil(t, res)

	assert.Equal(t, expectedContent, res.Plurks[0].Content)
	assert.Equal(t, expectedUserName, res.Users[strconv.Itoa(res.Plurks[0].OwnerID)].DisplayName)
}
