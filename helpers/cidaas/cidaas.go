package cidaas

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type Client struct {
<<<<<<< HEAD
	Role           RoleService
	CustomProvider CustomProvideService
	SocialProvider SocialProviderService
	Scope          ScopeService
	ScopeGroup     ScopeGroupService
	ConsentGroup   ConsentGroupService
	GroupType      GroupTypeService
	UserGroup      UserGroupService
	HostedPage     HostedPageService
	Webhook        WebhookService
	App            AppService
	RegField       RegFieldService
	TemplateGroup            TemplateGroupService
	Template                 TemplateService
	TemplateType             TemplateTypeService
	NotificationTemplateType TemplateTypeService
	PasswordPolicy           PasswordPolicyService
	Consent        ConsentService
	ConsentVersion ConsentVersionService
=======
	Roles          *Role
	CustomProvider *CustomProvider
	SocialProvider *SocialProvider
	Scopes         *Scope
	ScopeGroup     *ScopeGroup
	ConsentGroup   *ConsentGroup
	GroupType      *GroupType
	UserGroup      *UserGroup
	HostedPages    *HostedPage
	Webhook        *Webhook
	Apps           *App
	RegFields      *RegField
	TemplateGroup  *TemplateGroup
	Templates      *Template
	PasswordPolicy *PasswordPolicy
	Consent        *Consent
	ConsentVersion *ConsentVersion
>>>>>>> master
}

type ClientConfig struct {
	ClientID     string
	ClientSecret string
	BaseURL      string
	AccessToken  string
}

func (c *ClientConfig) makeRequest(ctx context.Context, method, endpoint string, body interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)
	client, err := util.NewHTTPClient(url, method, c.AccessToken)
	if err != nil {
		return nil, err
	}

	res, err := client.MakeRequest(ctx, body)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	return res, nil
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func NewClient(ctx context.Context, config ClientConfig) (*Client, error) {
	re := regexp.MustCompile(`/*$`)
	config.BaseURL = re.ReplaceAllString(config.BaseURL, "")
	tokenURL, err := url.JoinPath(config.BaseURL, "token-srv/token")
	if err != nil {
		return nil, fmt.Errorf("failed to create token url %s", err.Error())
	}
	httpClient, err := util.NewHTTPClient(tokenURL, http.MethodPost)
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"client_id":     config.ClientID,
		"client_secret": config.ClientSecret,
		"grant_type":    "client_credentials",
	}
	res, err := httpClient.MakeRequest(ctx, payload)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, fmt.Errorf("failed to generate access token %s", err.Error())
	}
	defer res.Body.Close()
	var response TokenResponse
	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, fmt.Errorf("failed to generate access token %s", err.Error())
	}
	config.AccessToken = response.AccessToken
	client := &Client{
		Roles:          NewRole(config),
		CustomProvider: NewCustomProvider(config),
		Scopes:         NewScope(config),
		ScopeGroup:     NewScopeGroup(config),
		GroupType:      NewGroupType(config),
		UserGroup:      NewUserGroup(config),
		HostedPages:    NewHostedPage(config),
		Webhook:        NewWebhook(config),
<<<<<<< HEAD
		App:            NewApp(config),
		RegField:       NewRegField(config),
		TemplateGroup:            NewTemplateGroup(config),
		Template:                 NewTemplate(config),
		TemplateType:             NewTemplateType(config),
		NotificationTemplateType: NewTemplateType(config),
		SocialProvider:           NewSocialProvider(config),
		PasswordPolicy:           NewPasswordPolicy(config),
=======
		Apps:           NewApp(config),
		RegFields:      NewRegField(config),
		TemplateGroup:  NewTemplateGroup(config),
		Templates:      NewTemplate(config),
		SocialProvider: NewSocialProvider(config),
		PasswordPolicy: NewPasswordPolicy(config),
>>>>>>> master
		ConsentGroup:   NewConsentGroup(config),
		Consent:        NewConsent(config),
		ConsentVersion: NewConsentVersion(config),
	}
	return client, nil
}
