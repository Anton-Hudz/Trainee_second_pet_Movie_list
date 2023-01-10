package transport

const (
	MsgInternalSeverErr      = "Internal server error"
	MsgBadRequest            = "Bad request"
	MsgNotFound              = "Not found"
	MsgEmptyAuthHeader       = "Empty authorization header"
	MsgInvalidAuthHeader     = "Invalid authorization header"
	MsgProblemWithParseToken = "Problem while parsing token"
	MsgHaveNotPermission     = "You don’t have permission to access"
	readTimeoutSeconds       = 10
	writeTimeoutSeconds      = 10
)
