package errors

type Error struct {
	Code     int    `json:"-"` // don't include in response
	Type     string `json:"type,omitempty"`
	Message  string `json:"message,omitempty"`
	MoreInfo string `json:"moreInfo,omitempty"`
}

func (err Error) Error() string {
	return err.Message
}

var Unknown = Error{
	Type:    "UnknownError",
	Message: "Unknown error",
}
