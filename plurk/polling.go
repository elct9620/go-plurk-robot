package plurk

type Polling struct {
	*PlurkClient
}

// Get Polling Endpoint
func (p *PlurkClient) GetPolling() *Polling {
	return &Polling{p}
}
