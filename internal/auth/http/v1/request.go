package authv1

import (
	"net/http"
	"strings"
)

type BasicLoginPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func getBearer(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	auths := strings.Split(authHeader, "Bearer")
	if len(auths) < 2 {
		return ""
	}

	return strings.TrimSpace(auths[1])
}
