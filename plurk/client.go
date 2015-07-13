// Package for Plurk API 2.0 Client
package plurk

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Endpoint
const apiBase = "http://www.plurk.com/APP"

// API Instance
type Plurk struct {
	credential Credential
}

// Robot Authorize Information
type Credential struct {
	AppKey      string
	AppSecret   string
	Token       string
	TokenSecret string
}

// Signature from params and url to generate OAuth 1.0 signature
func (c *Credential) Signature(uri *url.URL, method string, params url.Values) string {
	// HMAC-SHA1
	var signatureURL = fmt.Sprintf(
		"%s&%s&%s", // Method&URI&Params
		method,
		url.QueryEscape(uri.String()),
		strings.Replace(url.QueryEscape(params.Encode()), "%2B", "%2520", -1), // Resolve space " " change to "+" after encode
	)

	key := []byte(fmt.Sprint(c.AppSecret, "&", c.TokenSecret))
	h := hmac.New(sha1.New, key)
	h.Write([]byte(signatureURL))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// Create OAuth 1.0 nonce
// Simple use timestamp to generate unique nonce
func nonce() string {
	return strconv.FormatInt(time.Now().UnixNano(), 16)
}

// Add OAuth 1.0 request params into request
func signParams(token *Credential, method string, uri *url.URL, params url.Values) url.Values {

	// OAuth 1.0 Basic
	params.Set("oauth_consumer_key", token.AppKey)
	params.Set("oauth_signature_method", "HMAC-SHA1")
	params.Set("oauth_version", "1.0")

	// OAuth 1.0 Timestamp
	params.Set("oauth_timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	params.Set("oauth_nonce", nonce())

	params.Set("oauth_token", token.Token)

	params.Set("oauth_signature", token.Signature(uri, method, params))

	return params
}

// Send GET Request to Plurk API
func get(endpoint string, token *Credential, params url.Values) {

	requestUri := fmt.Sprintf("%s/%s", apiBase, endpoint)
	uri, _ := url.Parse(requestUri)
	params = signParams(token, "GET", uri, params)
	requestUri = fmt.Sprint(requestUri, "?", params.Encode())
	res, _ := http.Get(requestUri)

	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)

	// NOTE(elct9620): Should return a JSON string for parse
	fmt.Println(string(data))
}

// Send POST Request to Plurk API
func post() {
	// NOTE:(elct9620): Should implement POST Method same as GET
}

// Helper to generate Pluck Instance
func New(AppKey string, AppSecret string, Token string, TokenSecret string) *Plurk {
	credential := Credential{
		AppKey:      AppKey,
		AppSecret:   AppSecret,
		Token:       Token,
		TokenSecret: TokenSecret,
	}
	return &Plurk{credential: credential}
}

// Echo, Plruk API which can return same data
func (plurk *Plurk) Echo(data string) {
	params := make(url.Values)

	// If has data, add data as parameter
	if len(data) > 0 {
		params.Set("data", data)
	}

	get("echo", &plurk.credential, params)
}
