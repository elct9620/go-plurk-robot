package plurk

import (
	"encoding/json"
	"fmt"
	"github.com/elct9620/go-plurk-robot/logger"
	"net/url"
	"strings"
)

// Timeline API
type Timeline struct {
	*PlurkClient
}

// PlurkAdd Response
type PlurkAddResponse struct {
	PlurkID             int `json:"plurk_id"`
	Content             string
	QualifierTranslated string `json:"qualifier_translated"`
	Qualifier           string
	Lang                string
}

// GetPlurkResponse
type GetPlurkResponse struct {
	Users  map[string]User `json:"plurk_users"`
	Plurks []Plurk
}

func (t *Timeline) PlurkAdd(content string, qualifier string, limitTo []int, disableComment bool, language string) (*PlurkAddResponse, error) {

	params := make(url.Values)
	params.Add("content", content)
	params.Add("qualifier", qualifier)

	// Should limit to user
	if len(limitTo) > 0 {
		// []int hack, to JSON array
		params.Add("limit_to", strings.Replace(fmt.Sprintf("%v", limitTo), " ", ", ", -1))
	}

	params.Add("no_comments", fmt.Sprintf("%d", BoolToInt(disableComment)))
	params.Add("lang", language)

	data, err := t.Post("Timeline/plurkAdd", params)
	if err != nil {
		return nil, err
	}

	var result PlurkAddResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (t *Timeline) GetPlurks(offset int, limit int, filter string) (result *GetPlurkResponse, err error) {
	params := make(url.Values)
	params.Add("offset", fmt.Sprintf("%d", offset))
	params.Add("limit", fmt.Sprintf("%d", limit))
	if len(filter) > 0 {
		params.Add("filter", fmt.Sprintf("%d", filter))
	}

	data, err := t.Get("Timeline/getPlurks", params)
	if err != nil {
		return nil, err
	}

	logger.Info("Plruks JSON: %s", string(data))

	result = &GetPlurkResponse{}
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}

	return result, err
}
