package response

// Content creates a new Builder with the supplied content using DefaultConfig.
func Content(content any) Builder {
	return NewBuilder().WithContent(content)
}

// Status creates a new Builder with the supplied status using DefaultConfig.
func Status(status int) Builder {
	return NewBuilder().WithStatus(status)
}

// BadGateway returns a 502 Bad Gateway Builder using DefaultConfig.
func BadGateway() Builder {
	return NewBuilder().BadGateway()
}

// BadRequest returns a 400 Bad Request Builder using DefaultConfig.
func BadRequest() Builder {
	return NewBuilder().BadRequest()
}

// Conflict returns a 409 Conflict Builder using DefaultConfig.
func Conflict() Builder {
	return NewBuilder().Conflict()
}

// Created returns a 201 Created Builder using DefaultConfig.
func Created() Builder {
	return NewBuilder().Created()
}

// Forbidden returns a 403 Forbidden Builder using DefaultConfig.
func Forbidden() Builder {
	return NewBuilder().Forbidden()
}

// Found returns a 302 Found Builder using DefaultConfig.
func Found(location string) Builder {
	return NewBuilder().Found(location)
}

// GatewayTimeout returns a 504 Gateway Timeout Builder using DefaultConfig.
func GatewayTimeout() Builder {
	return NewBuilder().GatewayTimeout()
}

// InternalServerError returns a 500 Internal Server Error Builder using DefaultConfig.
func InternalServerError() Builder {
	return NewBuilder().InternalServerError()
}

// OK returns a 200 OK Builder using DefaultConfig.
func OK() Builder {
	return NewBuilder().OK()
}

// MovedPermanently returns a 301 Moved Permanently Builder using DefaultConfig.
func MovedPermanently(location string) Builder {
	return NewBuilder().MovedPermanently(location)
}

// NotFound returns a 404 Not Found Builder using DefaultConfig.
func NotFound() Builder {
	return NewBuilder().NotFound()
}

// NotImplemented returns a 501 Not Implemented Builder using DefaultConfig.
func NotImplemented() Builder {
	return NewBuilder().NotImplemented()
}

// PermanentRedirect returns a 308 Permanent Redirect Builder using DefaultConfig.
func PermanentRedirect(location string) Builder {
	return NewBuilder().PermanentRedirect(location)
}

// TemporaryRedirect returns a 307 Temporary Redirect Builder using DefaultConfig.
func TemporaryRedirect(location string) Builder {
	return NewBuilder().TemporaryRedirect(location)
}

// Unauthorized returns a 401 Unauthorized Builder using DefaultConfig.
func Unauthorized() Builder {
	return NewBuilder().Unauthorized()
}
