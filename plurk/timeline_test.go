package plurk

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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

	plurkClient := &Plurk{credential: credential, ApiBase: server.URL}
	timeline := plurkClient.GetTimeline()
	res, _ := timeline.PlurkAdd(content, qualifier, make([]int, 0), false, lang)

	assert.Equal(t, content, res.Content)
	assert.Equal(t, lang, res.Lang)
	assert.Equal(t, qualifier, res.Qualifier)
}
