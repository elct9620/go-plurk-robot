package plurk

import (
	"encoding/json"
	"fmt"
	"github.com/elct9620/go-plurk-robot/logger"
	"net/url"
)

type Polling struct {
	*PlurkClient
}

// Get Polling Endpoint
func (p *PlurkClient) GetPolling() *Polling {
	return &Polling{p}
}

// Get Plurks using "Polling" way
func (p *Polling) GetPlurks(offset *Time, limit int) (result *GetPlurkResponse, err error) {
	params := make(url.Values)
	params.Add("offset", offset.PollingOffset())
	params.Add("limit", fmt.Sprintf("%d", limit))

	logger.Info("Polling from %s", offset.PollingOffset())

	data, err := p.Get("Polling/getPlurks", params)
	if err != nil {
		return nil, err
	}

	result = &GetPlurkResponse{}
	err = json.Unmarshal(data, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
