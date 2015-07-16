package plurk

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"text/template"
)

func buildPlurkAddServer(responseFile string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		id := 1
		content := r.PostFormValue("content")
		lang := r.PostFormValue("lang")
		qualifier := r.PostFormValue("qualifier")

		resData := &Plurk{
			PlurkID:    id,
			RawContent: content,
			Language:   lang,
			Qualifier:  qualifier,
		}

		ts, _ := template.ParseFiles(responseFile)
		ts.Execute(w, resData)
	}))
}

func Test_PlurkAdd(t *testing.T) {

	testPath := "test/plurk_add_response.json"

	expectedContent := "Hello World"
	expectedLanguage := "en"
	expectedQualifier := "says"

	server := buildPlurkAddServer(testPath)
	defer server.Close()

	plurkClient := &PlurkClient{credential: credential, ApiBase: server.URL}
	timeline := plurkClient.GetTimeline()
	res, _ := timeline.PlurkAdd(expectedContent, expectedQualifier, make([]int, 0), false, expectedLanguage, false)

	assert.Equal(t, expectedContent, res.RawContent)
	assert.Equal(t, expectedLanguage, res.Language)
	assert.Equal(t, expectedQualifier, res.Qualifier)
}

func Test_PlurkAddIgnoreSocial(t *testing.T) {

	testPath := "test/plurk_add_response.json"

	expectedContent := "Hello World"
	expectedLanguage := "en"
	expectedQualifier := "says"

	server := buildPlurkAddServer(testPath)
	defer server.Close()

	plurkClient := &PlurkClient{credential: credential, ApiBase: server.URL}
	timeline := plurkClient.GetTimeline()
	res, _ := timeline.PlurkAdd(expectedContent, expectedQualifier, make([]int, 0), false, expectedLanguage, true)

	assert.Equal(t, expectedContent+" !FB !TW", res.RawContent)
	assert.Equal(t, expectedLanguage, res.Language)
	assert.Equal(t, expectedQualifier, res.Qualifier)
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
