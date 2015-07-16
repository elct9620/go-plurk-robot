package plurk

type Plurk struct {
	PlurkID             int `json:"plurk_id"`
	Content             string
	RawContent          string `json:"content_raw"`
	OwnerID             int    `json:"owner_id"`
	Replurked           bool
	ReplurkerID         int `json:"replurker_id"`
	Posted              Time
	Anonymous           bool
	PlurkType           int    `json:"plurk_type"`
	QualifierTranslated string `json:"qualifier_translated"`
	Qualifier           string
	ResponseCount       int    `json:"response_count"`
	NoComments          int    `json:"no_comments"`
	Language            string `json:"lang"`
	Favorite            bool
	FacebookDisabled    bool   `json:"facebook_disabled,omitempty"`
	TwitterDisabled     bool   `json:"twitter_disabled,omitempty"`
	LimitedTo           string `json:"limited_to"` // TODO(elct9620): Convert `LimitedTo` as a array
}
