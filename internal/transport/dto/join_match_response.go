package dto

import "youpiteron.dev/wildlands-backend/internal/transport"

type JoinMatchResponse struct {
	Match transport.JsonMatch `json:"match"`
}
