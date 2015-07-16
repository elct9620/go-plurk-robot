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

func Test_PlurkAdd(t *testing.T) {

	content := "Hello World"
	lang := "en"
	qualifier := "says"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		id := 1
		content := r.PostFormValue("content")
		lang := r.PostFormValue("lang")
		qualifier := r.PostFormValue("qualifier")
		fmt.Fprintf(w, "{\"plurk_id\": %d, \"content\": \"%s\", \"lang\": \"%s\", \"qualifier\": \"%s\"}", id, content, lang, qualifier)
	}))

	defer server.Close()

	plurkClient := &PlurkClient{credential: credential, ApiBase: server.URL}
	timeline := plurkClient.GetTimeline()
	res, _ := timeline.PlurkAdd(content, qualifier, make([]int, 0), false, lang)

	assert.Equal(t, content, res.Content)
	assert.Equal(t, lang, res.Lang)
	assert.Equal(t, qualifier, res.Qualifier)
}

func Test_GetPlurks(t *testing.T) {
	testPath := "test/get_plurks.json"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := ioutil.ReadFile(testPath)
		fmt.Fprintln(w, string(jsonData))
	}))

	defer server.Close()

	plurkClient := &PlurkClient{credential: credential, ApiBase: server.URL}
	timeline := plurkClient.GetTimeline()
	res, err := timeline.GetPlurks(0, 1, "")

	if err != nil {
		assert.Error(t, err)
	}

	expectedContent := "Test Data"
	expectedUserName := "Tester"

	assert.NotNil(t, res)

	assert.Equal(t, expectedContent, res.Plurks[0].Content)
	assert.Equal(t, expectedUserName, res.Users[strconv.Itoa(res.Plurks[0].OwnerID)].DisplayName)
}
