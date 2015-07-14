package plurk

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// Using Plurk Test Console generate Fake Signature
var credential Credential = Credential{
	AppKey:      "TestAppKey",
	AppSecret:   "TestAppSecret",
	Token:       "rLXmPnnQpviV",
	TokenSecret: "56Fgl58yOfqXOhHXX0ybvOmSnPQFvR2miYmm30A",
}

func mockParams() url.Values {
	var params url.Values = make(url.Values)

	// OAuth 1.0 Basic
	params.Set("oauth_consumer_key", credential.AppKey)
	params.Set("oauth_signature_method", "HMAC-SHA1")
	params.Set("oauth_version", "1.0")

	// OAuth 1.0 Timestamp
	params.Set("oauth_timestamp", "1436769226")
	params.Set("oauth_nonce", "97990917")

	params.Set("oauth_token", credential.Token)
	return params
}

func Test_Signature(t *testing.T) {

	uri, _ := url.Parse("http://www.plurk.com/APP/echo")

	signature := credential.Signature(uri, "GET", mockParams())

	expectedSignature := "FEgaoJyXWYy3FBWYCog8NI63xRo="
	assert.Equal(t, signature, expectedSignature)
}

func Test_Nonce(t *testing.T) {
	n := nonce()

	assert.False(t, len(n) < 8, "nonce %s should longer then 8", n)
}

func Test_SignParams(t *testing.T) {
	uri, _ := url.Parse("http://www.plurk.com/APP/echo")

	signedParams := signParams(&credential, "GET", uri, make(url.Values))

	signature := signedParams.Get("oauth_signature")
	signedParams.Del("oauth_signature")

	expectedSignature := credential.Signature(uri, "GET", signedParams)

	assert.Equal(t, signature, expectedSignature)
}

func Test_New(t *testing.T) {
	plurk := New(credential.AppKey, credential.AppSecret, credential.Token, credential.TokenSecret)

	assert.Equal(t, plurk.credential, credential)
	assert.Equal(t, plurk.ApiBase, apiBase)
}

func Test_PlurkGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"message": "message"}`)
	}))

	defer server.Close()

	plurkClient := &Plurk{credential: credential, ApiBase: server.URL}
	data, _ := plurkClient.Get("/", make(url.Values))

	var result map[string]string
	json.Unmarshal(data, &result)

	assert.Equal(t, result["message"], "message")
}

func Test_PlurkGetError(t *testing.T) {
	errorText := "request error!"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"error_text\": \"%s\"}", errorText)
	}))

	defer server.Close()

	plurkClient := &Plurk{credential: credential, ApiBase: server.URL}
	_, err := plurkClient.Get("/", make(url.Values))

	assert.Equal(t, err.Error(), errorText)
}

func Test_PlurkEcho(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		data := r.URL.Query().Get("data")
		fmt.Fprintf(w, "{\"length\": %d, \"data\": \"%s\"}", len(data), data)
	}))

	defer server.Close()

	requestData := "Hello World"

	plurkClient := &Plurk{credential: credential, ApiBase: server.URL}
	result, _ := plurkClient.Echo(requestData)

	assert.Equal(t, result.Length, len(requestData))
	assert.Equal(t, result.Data, requestData)
}
