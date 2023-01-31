package transport

const (
	MsgInternalServerErr       = "Internal server error"
	MsgBadRequest              = "Bad request"
	MsgNotFound                = "Not found"
	MsgEmptyAuthHeader         = "Empty authorization header"
	MsgInvalidAuthHeader       = "Invalid authorization header"
	MsgProblemWithParseToken   = "Problem while parsing token"
	MsgHaveNotPermission       = "You don’t have permission to access"
	MsgInvalidIDParam          = "Invalid id parameter"
	MsgProblemFormatOutputData = "Problem with format of output data"
	MsgProblemFieldList        = "Problem with field of list"
	readTimeoutSeconds         = 10
	writeTimeoutSeconds        = 10
)
