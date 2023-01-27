package response

func Content(content interface{}) response {
	return DefaultResponseBuilder.Content(content)
}

func Status(status int) response {
	return DefaultResponseBuilder.Status(status)
}

/*
"Status shortcut methods"
*/

func BadGateway() response {
	return DefaultResponseBuilder.BadGateway()
}

func BadRequest() response {
	return DefaultResponseBuilder.BadRequest()
}

func Conflict() response {
	return DefaultResponseBuilder.Conflict()
}

func Created() response {
	return DefaultResponseBuilder.Created()
}

func Forbidden() response {
	return DefaultResponseBuilder.Forbidden()
}

func Found() response {
	return DefaultResponseBuilder.Found()
}

func GatewayTimeout() response {
	return DefaultResponseBuilder.GatewayTimeout()
}

func InternalServerError() response {
	return DefaultResponseBuilder.InternalServerError()
}

func OK() response {
	return DefaultResponseBuilder.OK()
}

func MovedPermanently() response {
	return DefaultResponseBuilder.MovedPermanently()
}

func NotFound() response {
	return DefaultResponseBuilder.NotFound()
}

func NotImplemented() response {
	return DefaultResponseBuilder.NotImplemented()
}

func Unauthorized() response {
	return DefaultResponseBuilder.Unauthorized()
}
