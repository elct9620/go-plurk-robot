package main

import (
	"bytes"
	"github.com/elct9620/go-plurk-robot/plurk"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
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
		id := 999999
		content := r.PostFormValue("content")
		lang := r.PostFormValue("lang")
		qualifier := r.PostFormValue("qualifier")

		resData := &plurk.Plurk{
			PlurkID:    id,
			RawContent: content,
			Language:   lang,
			Qualifier:  qualifier,
		}

		ts, _ := template.ParseFiles(responseFile)
		ts.Execute(w, resData)
	}))
}

func buildResponseAddServer(responseFile string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		idString := r.PostFormValue("plurk_id")
		content := r.PostFormValue("content")
		qualifier := r.PostFormValue("qualifier")

		id, _ := strconv.Atoi(idString)

		resData := &plurk.Plurk{
			PlurkID:    id,
			RawContent: content,
			Language:   "en",
			Qualifier:  qualifier,
		}

		ts, _ := template.ParseFiles(responseFile)
		ts.Execute(w, resData)
	}))
}

func Test_AddPlurkCommand(t *testing.T) {

	testPath := "test/plurk_add_response.json"
	server := buildPlurkAddServer(testPath)
	defer server.Close()

	buffer := bytes.NewBuffer(make([]byte, 0))
	cmd := &cobra.Command{}
	cmd.SetOutput(buffer)

	Client = &plurk.PlurkClient{ApiBase: server.URL}
	addPlurk(cmd, []string{"Hello World"})

	assert.Contains(t, buffer.String(), "Plurk ID: 999999")
}

func Test_AddPlurkCommandWithoutContent(t *testing.T) {

	testPath := "test/plurk_add_response.json"
	server := buildPlurkAddServer(testPath)
	defer server.Close()

	buffer := bytes.NewBuffer(make([]byte, 0))
	cmd := &cobra.Command{}
	cmd.SetOutput(buffer)

	Client = &plurk.PlurkClient{ApiBase: server.URL}
	addPlurk(cmd, make([]string, 0))

	assert.Contains(t, buffer.String(), "No plurk content specified")
}

func Test_AddResponseCommand(t *testing.T) {

	testPath := "test/response_add.json"
	server := buildResponseAddServer(testPath)
	defer server.Close()

	buffer := bytes.NewBuffer(make([]byte, 0))
	cmd := &cobra.Command{}
	cmd.SetOutput(buffer)

	Client = &plurk.PlurkClient{ApiBase: server.URL}
	addResponse(cmd, []string{"123456", "Message"})

	assert.Contains(t, buffer.String(), "Respons success add to 123456")
}

func Test_AddResponseCommandInvalidPlurkID(t *testing.T) {

	testPath := "test/response_add.json"
	server := buildResponseAddServer(testPath)
	defer server.Close()

	buffer := bytes.NewBuffer(make([]byte, 0))
	cmd := &cobra.Command{}
	cmd.SetOutput(buffer)

	Client = &plurk.PlurkClient{ApiBase: server.URL}
	addResponse(cmd, []string{"null", "message"})

	assert.Contains(t, buffer.String(), "Convert plurk id error")
}

func Test_AddResponseCommandWithoutContent(t *testing.T) {

	testPath := "test/response_add.json"
	server := buildResponseAddServer(testPath)
	defer server.Close()

	buffer := bytes.NewBuffer(make([]byte, 0))
	cmd := &cobra.Command{}
	cmd.SetOutput(buffer)

	Client = &plurk.PlurkClient{ApiBase: server.URL}
	addResponse(cmd, []string{"123456"})

	assert.Contains(t, buffer.String(), "No plurk id or response content specified")
}
