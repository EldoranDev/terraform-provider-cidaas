package cidaas

import (
	"encoding/json"
	"errors"
	"fmt"
)

// notificationSrvEnvelope matches basetype.Response[T] JSON from notification-srv handlers.
type notificationSrvEnvelope struct {
	Success  bool            `json:"success"`
	Status   int             `json:"status"`
	Code     string          `json:"code,omitempty"`
	ErrorMsg string          `json:"errorMsg,omitempty"`
	ErrorAlt string          `json:"error,omitempty"`
	Data     json.RawMessage `json:"data"`
}

// ParseNotificationSrvData unmarshals the `data` field from a notification-srv JSON envelope into T.
func ParseNotificationSrvData[T any](body []byte, httpStatus int) (*T, error) {
	var env notificationSrvEnvelope
	if err := json.Unmarshal(body, &env); err != nil {
		return nil, fmt.Errorf("failed to parse notification-srv response: %w", errors.New(truncateBody(body)))
	}
	errMsg := env.ErrorMsg
	if errMsg == "" {
		errMsg = env.ErrorAlt
	}
	if !env.Success || string(env.Data) == "null" || len(env.Data) == 0 {
		if errMsg != "" {
			return nil, fmt.Errorf("notification-srv error (status %d, code %s): %s", env.Status, env.Code, errMsg)
		}
		if httpStatus >= 400 {
			return nil, fmt.Errorf("notification-srv request failed with HTTP %d: %s", httpStatus, string(body))
		}
		return nil, fmt.Errorf("notification-srv: empty or unsuccessful response: %s", string(body))
	}
	var out T
	if err := json.Unmarshal(env.Data, &out); err != nil {
		return nil, fmt.Errorf("failed to parse notification-srv data: %w", err)
	}
	return &out, nil
}

// ParseNotificationSrvDataOrNil unmarshals envelope data into T when present; if `data` is JSON null or empty, returns (nil, nil) without error (e.g. GET by id with no match).
func ParseNotificationSrvDataOrNil[T any](body []byte, httpStatus int) (*T, error) {
	var env notificationSrvEnvelope
	if err := json.Unmarshal(body, &env); err != nil {
		return nil, fmt.Errorf("failed to parse notification-srv response: %w", errors.New(truncateBody(body)))
	}
	errMsg := env.ErrorMsg
	if errMsg == "" {
		errMsg = env.ErrorAlt
	}
	if !env.Success {
		if errMsg != "" {
			return nil, fmt.Errorf("notification-srv error (status %d, code %s): %s", env.Status, env.Code, errMsg)
		}
		if httpStatus >= 400 {
			return nil, fmt.Errorf("notification-srv request failed with HTTP %d: %s", httpStatus, string(body))
		}
		return nil, fmt.Errorf("notification-srv: unsuccessful response: %s", string(body))
	}
	if len(env.Data) == 0 || string(env.Data) == "null" {
		return nil, nil
	}
	var out T
	if err := json.Unmarshal(env.Data, &out); err != nil {
		return nil, fmt.Errorf("failed to parse notification-srv data: %w", err)
	}
	return &out, nil
}
