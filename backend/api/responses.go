package api

type APIResponse struct {
	success bool
	message string
}

type APIJSONResponse struct {
	success bool
	message string
}

func NewSuccessResponse(message string) APIResponse {
	return APIResponse{
		success: true,
		message: message,
	}
}

func NewFailureResponse(message string) APIResponse {
	return APIResponse{
		success: false,
		message: message,
	}
}