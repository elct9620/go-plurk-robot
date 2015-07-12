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
	"time"
)

// Endpoint
const apiBase = "http://www.plurk.com/APP"

type Plurk struct {
	credential Credential
}

// Robot Authorize Information
type Credential struct {
	AppKey      string
	AppSecret   string
	OauthToken  string
	TokenSecret string
}

func nonce() string { // Simple use timestamp to generate unique nonce
	return strconv.FormatInt(time.Now().UnixNano(), 16)
}

func signParams(token *Credential, method string, uri *url.URL, params url.Values) url.Values {

	// OAuth 1.0 Basic
	params.Set("oauth_consumer_key", token.AppKey)
	params.Set("oauth_signature_method", "HMAC-SHA1")
	params.Set("oauth_version", "1.0")

	// OAuth 1.0 Timestamp
	params.Set("oauth_timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	params.Set("oauth_nonce", nonce())

	params.Set("oauth_token", token.OauthToken)

	// HMAC-SHA1
	var signatureURL = fmt.Sprintf(
		"%s&%s&%s", // Method&URI&Params
		method,
		url.QueryEscape(uri.String()),
		url.QueryEscape(params.Encode()),
	)

	key := []byte(fmt.Sprint(token.AppSecret, "&", token.TokenSecret))
	h := hmac.New(sha1.New, key)
	h.Write([]byte(signatureURL))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	params.Set("oauth_signature", signature)

	return params
}

func get(endpoint string, token *Credential, params url.Values) {

	requestUri := fmt.Sprintf("%s/%s", apiBase, endpoint)
	uri, _ := url.Parse(requestUri)
	params = signParams(token, "GET", uri, params)
	requestUri = fmt.Sprint(requestUri, "?", params.Encode())
	res, _ := http.Get(requestUri)

	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)

	// Simple debug result
	fmt.Println(string(data))
}

func post() {

}

func New(AppKey string, AppSecret string, OauthToken string, TokenSecret string) *Plurk {
	credential := Credential{
		AppKey:      AppKey,
		AppSecret:   AppSecret,
		OauthToken:  OauthToken,
		TokenSecret: TokenSecret,
	}
	return &Plurk{credential: credential}
}

func (plurk *Plurk) Echo() {
	get("echo", &plurk.credential, make(url.Values))
}
