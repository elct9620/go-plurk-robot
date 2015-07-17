package plurk

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"text/template"
)

func buildResponseAddServer(responseFile string, isError bool, errorText string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Should return error or not
		if isError {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(200)
		}

		idString := r.PostFormValue("plurk_id")
		content := r.PostFormValue("content")
		qualifier := r.PostFormValue("qualifier")

		id, _ := strconv.Atoi(idString)

		var resData interface{}
		resData = &Response{
			PlurkID:    id,
			RawContent: content,
			Content:    content,
			Qualifier:  qualifier,
		}

		if isError {
			resData = &struct {
				ErrorText string
			}{
				ErrorText: errorText,
			}
		}

		ts, _ := template.ParseFiles(responseFile)
		ts.Execute(os.Stdout, resData)
		ts.Execute(w, resData)
	}))
}

func Test_ResponseAdd(t *testing.T) {
	testPath := "test/response_add.json"

	expectedPlurkID := 999999
	expectedContent := "Hello World"
	expectedQualifier := "says"

	server := buildResponseAddServer(testPath, false, "")
	defer server.Close()

	plurkClient := &PlurkClient{credential: credential, ApiBase: server.URL}
	responses := plurkClient.GetResponses()
	res, _ := responses.ResponseAdd(expectedPlurkID, expectedContent, expectedQualifier)

	assert.Equal(t, expectedContent, res.RawContent)
	assert.Equal(t, expectedPlurkID, res.PlurkID)
	assert.Equal(t, expectedQualifier, res.Qualifier)
}

func Test_ResponseAddError(t *testing.T) {
	testPath := "test/error.json"

	expectedError := "Some error"

	server := buildResponseAddServer(testPath, true, expectedError)
	defer server.Close()

	plurkClient := &PlurkClient{credential: credential, ApiBase: server.URL}
	responses := plurkClient.GetResponses()
	res, err := responses.ResponseAdd(999999, "Hello World", "says")

	assert.Nil(t, res)
	assert.EqualError(t, err, expectedError)
}
