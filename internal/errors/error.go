package errors

type Error struct {
	Type     string `json:"type,omitempty"`
	Message  string `json:"message,omitempty"`
	MoreInfo string `json:"moreInfo,omitempty"`
}

var Unknown = Error{
	Type:    "UnknownError",
	Message: "Unknown error",
}
