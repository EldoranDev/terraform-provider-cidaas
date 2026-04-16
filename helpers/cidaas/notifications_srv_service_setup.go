package cidaas

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

// NotificationsSrvServiceSetup calls GET /servicesetups.
type NotificationsSrvServiceSetup struct {
	ClientConfig
}

// NewNotificationsSrvServiceSetup builds a client for notification-srv service setup listing.
func NewNotificationsSrvServiceSetup(cfg ClientConfig) *NotificationsSrvServiceSetup {
	return &NotificationsSrvServiceSetup{ClientConfig: cfg}
}

func (s *NotificationsSrvServiceSetup) segmentURL(parts ...string) string {
	return SegmentNotificationsURL(s.ClientConfig, parts...)
}

// NotificationsSrvServiceSetupModel is a subset of fields used for Terraform datasource (full API may return more).
type NotificationsSrvServiceSetupModel struct {
	ID                   string `json:"_id"`
	Name                 string `json:"name,omitempty"`
	Status               string `json:"status,omitempty"`
	Description          string `json:"description,omitempty"`
	ParentServiceSetupID string `json:"parentServiceSetupId,omitempty"`
	HasRemoteTemplates   bool   `json:"hasRemoteTemplates,omitempty"`
	ServiceDescInfo      struct {
		ServiceID       string `json:"serviceId"`
		ServiceCategory string `json:"serviceCategory,omitempty"`
		Name            string `json:"name,omitempty"`
	} `json:"serviceDescInfo"`
}

// List GET /servicesetups/
func (s *NotificationsSrvServiceSetup) List(ctx context.Context) ([]NotificationsSrvServiceSetupModel, error) {
	urlStr := s.segmentURL("servicesetups")
	client, err := util.NewHTTPClient(urlStr, http.MethodGet, s.AccessToken)
	if err != nil {
		return nil, err
	}
	res, err := client.MakeRequest(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read servicesetups body: %w", err)
	}
	out, err := ParseNotificationSrvData[[]NotificationsSrvServiceSetupModel](bodyBytes, res.StatusCode)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, nil
	}
	return *out, nil
}
