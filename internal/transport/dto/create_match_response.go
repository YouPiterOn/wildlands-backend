package dto

import "youpiteron.dev/wildlands-backend/internal/transport"

type CreateMatchResponse struct {
	Match transport.JsonMatch `json:"match"`
}
