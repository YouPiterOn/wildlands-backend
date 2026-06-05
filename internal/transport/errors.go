package transport

import "youpiteron.dev/wildlands-backend/internal/application"

type JsonError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func ToJsonError(err error) JsonError {
	switch err {
	case application.ErrEventsLoad:
		return JsonError{
			Code:    "EVENTS_LOAD_ERROR",
			Message: err.Error(),
		}
	case application.ErrEventsAppend:
		return JsonError{
			Code:    "EVENTS_APPEND_ERROR",
			Message: err.Error(),
		}
	case application.ErrEventApply:
		return JsonError{
			Code:    "EVENT_APPLY_ERROR",
			Message: err.Error(),
		}
	case application.ErrCommandHandle:
		return JsonError{
			Code:    "COMMAND_HANDLE_ERROR",
			Message: err.Error(),
		}
	default:
		return JsonError{
			Code:    "UNKNOWN_ERROR",
			Message: err.Error(),
		}
	}
}
