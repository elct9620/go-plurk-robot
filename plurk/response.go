package plurk

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Response Endpoint
type Responses struct {
	*PlurkClient
}

// Response data struct
type Response struct {
	Id         int
	Qualifier  string
	Content    string
	RawContent string `json:"content_raw"`
	PlurkID    int    `json:"plurk_id"`
	UserID     int    `json:"user_id"`
	Posted     Time
	Language   string `json:"lang"`
}

// Get Response API Endpoint from Plurk Client
func (p *PlurkClient) GetResponses() *Responses {
	return &Responses{p}
}

// Add response to specify plurk id
func (r *Responses) ResponseAdd(plurkID int, content string, qualifier string) (result *Response, err error) {

	params := make(url.Values)
	params.Add("plurk_id", fmt.Sprintf("%d", plurkID))
	params.Add("content", content)
	params.Add("qualifier", qualifier)

	data, err := r.Post("Responses/responseAdd", params)
	if err != nil {
		return nil, err
	}

	result = &Response{}
	err = json.Unmarshal(data, result)

	if err != nil {
		return nil, err
	}

	return result, err
}
