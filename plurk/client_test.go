package plurk

import (
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
	if signature != expectedSignature {
		t.Fatalf("Expected signature is %s but generated %s", expectedSignature, signature)
	}
}

func Test_Nonce(t *testing.T) {
	n := nonce()

	if len(n) < 8 {
		t.Fatalf("nonce is %s, exected something longer", n)
	}
}

func Test_SignParams(t *testing.T) {
	uri, _ := url.Parse("http://www.plurk.com/APP/echo")

	signedParams := signParams(&credential, "GET", uri, make(url.Values))

	signature := signedParams.Get("oauth_signature")
	signedParams.Del("oauth_signature")

	expectedSignature := credential.Signature(uri, "GET", signedParams)

	if signature != expectedSignature {
		t.Fatalf("Expected signature is %s but generated %s", expectedSignature, signature)
	}
}
