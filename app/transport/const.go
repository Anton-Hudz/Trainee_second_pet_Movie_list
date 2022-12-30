package transport

const (
	MsgInternalSeverErr      = "Internal server error"
	MsgBadRequest            = "Bad request"
	MsgNotFound              = "Not found"
	MsgEmptyAuthHeader       = "Empty authorization header"
	MsgInvalidAuthHeader     = "Invalid authorization header"
	MsgProblemWithParseToken = "Problem while parsing token"
	readTimeoutSeconds       = 10
	writeTimeoutSeconds      = 10
)
