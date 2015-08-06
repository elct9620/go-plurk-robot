package robot

import (
	"github.com/ddliu/motto"
	"github.com/elct9620/go-plurk-robot/plurk"
	"github.com/robertkrimen/otto"
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

func Test_plurkModuleLoader(t *testing.T) {
	vm := motto.New()
	module, _ := plurkModuleLoader(vm)
	// Convert to Object
	moduleObject := module.Object()
	keys := moduleObject.Keys()

	assert.Contains(t, keys, "addPlurk")
	assert.Contains(t, keys, "addResponse")
}

func TestAddPlurk(t *testing.T) {
	testPath := "test/plurk_add.json"
	server := buildPlurkAddServer(testPath)
	defer server.Close()

	client = &plurk.PlurkClient{ApiBase: server.URL}

	_, content, _ := otto.Run(`("Hello World")`)
	funcCall := otto.FunctionCall{}
	funcCall.ArgumentList = append(funcCall.ArgumentList, content)

	result := plurk_AddPlurk(funcCall)
	success, _ := result.ToBoolean()

	assert.True(t, success)
}

func TestAddPlurkNoContent(t *testing.T) {
	testPath := "test/plurk_add.json"
	server := buildPlurkAddServer(testPath)
	defer server.Close()

	client = &plurk.PlurkClient{ApiBase: server.URL}

	_, content, _ := otto.Run(`("")`)
	funcCall := otto.FunctionCall{}
	funcCall.ArgumentList = append(funcCall.ArgumentList, content)

	result := plurk_AddPlurk(funcCall)
	success, _ := result.ToBoolean()

	assert.False(t, success)
}

func TestAddResponse(t *testing.T) {
	testPath := "test/response_add.json"
	server := buildResponseAddServer(testPath)
	defer server.Close()

	client = &plurk.PlurkClient{ApiBase: server.URL}

	_, plurkID, _ := otto.Run(`(123456)`)
	_, content, _ := otto.Run(`("Hello World")`)
	funcCall := otto.FunctionCall{}
	funcCall.ArgumentList = append(funcCall.ArgumentList, plurkID)
	funcCall.ArgumentList = append(funcCall.ArgumentList, content)

	result := plurk_AddResponse(funcCall)
	success, _ := result.ToBoolean()

	assert.True(t, success)
}

func TestAddResponseInvalidID(t *testing.T) {
	testPath := "test/response_add.json"
	server := buildResponseAddServer(testPath)
	defer server.Close()

	client = &plurk.PlurkClient{ApiBase: server.URL}

	_, plurkID, _ := otto.Run(`("Invalid")`)
	_, content, _ := otto.Run(`("Hello World")`)
	funcCall := otto.FunctionCall{}
	funcCall.ArgumentList = append(funcCall.ArgumentList, plurkID)
	funcCall.ArgumentList = append(funcCall.ArgumentList, content)

	result := plurk_AddResponse(funcCall)
	success, _ := result.ToBoolean()

	assert.False(t, success)
}

func TestAddResponseNoContent(t *testing.T) {
	testPath := "test/response_add.json"
	server := buildResponseAddServer(testPath)
	defer server.Close()

	client = &plurk.PlurkClient{ApiBase: server.URL}

	_, plurkID, _ := otto.Run(`(123456)`)
	funcCall := otto.FunctionCall{}
	funcCall.ArgumentList = append(funcCall.ArgumentList, plurkID)

	result := plurk_AddResponse(funcCall)
	success, _ := result.ToBoolean()

	assert.False(t, success)
}
