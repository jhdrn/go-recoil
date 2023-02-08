package response

// Content creates a new response with the supplied content using the default
// response builder.
func Content(content interface{}) response {
	return DefaultResponseBuilder.Content(content)
}

// Status creates a new response with the supplied status using the default
// response builder.
func Status(status int) response {
	return DefaultResponseBuilder.Status(status)
}

// BadGateway returns a 502 Bad Gateway response using the default response builder.
func BadGateway() response {
	return DefaultResponseBuilder.BadGateway()
}

// BadRequest returns a 400 Bad Request response using the default response builder.
func BadRequest() response {
	return DefaultResponseBuilder.BadRequest()
}

// Conflict returns a 409 Conflict response using the default response builder.
func Conflict() response {
	return DefaultResponseBuilder.Conflict()
}

// Created returns a 201 Created response using the default response builder.
func Created() response {
	return DefaultResponseBuilder.Created()
}

// Forbidden returns a 403 Forbidden response using the default response builder.
func Forbidden() response {
	return DefaultResponseBuilder.Forbidden()
}

// Found returns a 302 Found response using the default response builder.
func Found() response {
	return DefaultResponseBuilder.Found()
}

// GatewayTimeout returns a 504 Gateway Timeout response using the default
// response builder.
func GatewayTimeout() response {
	return DefaultResponseBuilder.GatewayTimeout()
}

// InternalServerError returns a 500 Internal Server Error response using the
// default response builder.
func InternalServerError() response {
	return DefaultResponseBuilder.InternalServerError()
}

// MethodNotAllowed returns a 405 Method Not Allowed response using the default
// response builder.
func OK() response {
	return DefaultResponseBuilder.OK()
}

// MovedPermanently returns a 301 Moved Permanently response using the default
// response builder.
func MovedPermanently() response {
	return DefaultResponseBuilder.MovedPermanently()
}

// NoContent returns a 204 No Content response using the default response builder.
func NotFound() response {
	return DefaultResponseBuilder.NotFound()
}

// NotImplemented returns a 501 Not Implemented response using the default
// response builder.
func NotImplemented() response {
	return DefaultResponseBuilder.NotImplemented()
}

// Unauthorized returns a 401 Unauthorized response using the default response
// builder.
func Unauthorized() response {
	return DefaultResponseBuilder.Unauthorized()
}
