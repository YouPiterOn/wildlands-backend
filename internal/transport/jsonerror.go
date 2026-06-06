package transport

import "youpiteron.dev/wildlands-backend/internal/utils"

type JsonError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func ToJsonError(err error) JsonError {
	switch err {
	case utils.ErrEventLoad:
		return JsonError{
			Code:    "EVENT_LOAD_ERROR",
			Message: err.Error(),
		}
	case utils.ErrEventAppend:
		return JsonError{
			Code:    "EVENT_APPEND_ERROR",
			Message: err.Error(),
		}
	case utils.ErrEventApply:
		return JsonError{
			Code:    "EVENT_APPLY_ERROR",
			Message: err.Error(),
		}
	case utils.ErrCommandHandle:
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
