package slack

import (
	"context"
	"errors"
	"net/url"
)

type OAuthResponseIncomingWebhook struct {
	URL              string `json:"url"`
	Channel          string `json:"channel"`
	ChannelID        string `json:"channel_id,omitempty"`
	ConfigurationURL string `json:"configuration_url"`
}

type OAuthResponseBot struct {
	BotUserID      string `json:"bot_user_id"`
	BotAccessToken string `json:"bot_access_token"`
}

type OAuthResponse struct {
	AccessToken     string                       `json:"access_token"`
	Scope           string                       `json:"scope"`
	TeamName        string                       `json:"team_name"`
	TeamID          string                       `json:"team_id"`
	IncomingWebhook OAuthResponseIncomingWebhook `json:"incoming_webhook"`
	Bot             OAuthResponseBot             `json:"bot"`
	UserID          string                       `json:"user_id,omitempty"`
	SlackResponse
}

// OAuthV2Response ...
type OAuthV2Response struct {
	AccessToken     string                       `json:"access_token"`
	TokenType       string                       `json:"token_type"`
	Scope           string                       `json:"scope"`
	BotUserID       string                       `json:"bot_user_id"`
	AppID           string                       `json:"app_id"`
	Team            OAuthV2ResponseTeam          `json:"team"`
	IncomingWebhook OAuthResponseIncomingWebhook `json:"incoming_webhook"`
	Enterprise      OAuthV2ResponseEnterprise    `json:"enterprise"`
	AuthedUser      OAuthV2ResponseAuthedUser    `json:"authed_user"`
	SlackResponse
}

// OAuthV2ResponseTeam ...
type OAuthV2ResponseTeam struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OAuthV2ResponseEnterprise ...
type OAuthV2ResponseEnterprise struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OAuthV2ResponseAuthedUser ...
type OAuthV2ResponseAuthedUser struct {
	ID          string `json:"id"`
	Scope       string `json:"scope"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// GetOAuthToken retrieves an AccessToken
func GetOAuthToken(clientID, clientSecret, code, redirectURI string, debug bool) (accessToken string, scope string, err error) {
	return GetOAuthTokenContext(context.Background(), clientID, clientSecret, code, redirectURI, debug)
}

// GetOAuthTokenContext retrieves an AccessToken with a custom context
func GetOAuthTokenContext(ctx context.Context, clientID, clientSecret, code, redirectURI string, debug bool) (accessToken string, scope string, err error) {
	response, err := GetOAuthResponseContext(ctx, clientID, clientSecret, code, redirectURI, debug)
	if err != nil {
		return "", "", err
	}
	return response.AccessToken, response.Scope, nil
}

func GetOAuthResponse(clientID, clientSecret, code, redirectURI string, debug bool) (resp *OAuthResponse, err error) {
	return GetOAuthResponseContext(context.Background(), clientID, clientSecret, code, redirectURI, debug)
}

func GetOAuthResponseContext(ctx context.Context, clientID, clientSecret, code, redirectURI string, debug bool) (resp *OAuthResponse, err error) {
	values := url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"code":          {code},
		"redirect_uri":  {redirectURI},
	}
	response := &OAuthResponse{}
	err = post(ctx, customHTTPClient, "oauth.access", values, response, debug)
	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, errors.New(response.Error)
	}
	return response, nil
}

// GetOAuthV2Response gets a V2 OAuth access token response - https://api.slack.com/methods/oauth.v2.access
func GetOAuthV2Response(clientID, clientSecret, code, redirectURI string, debug bool) (resp *OAuthV2Response, err error) {
	return GetOAuthV2ResponseContext(context.Background(), clientID, clientSecret, code, redirectURI, debug)
}

// GetOAuthV2ResponseContext with a context, gets a V2 OAuth access token response
func GetOAuthV2ResponseContext(ctx context.Context, clientID, clientSecret, code, redirectURI string, debug bool) (resp *OAuthV2Response, err error) {
	values := url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"code":          {code},
		"redirect_uri":  {redirectURI},
	}
	response := &OAuthV2Response{}
	if err = post(ctx, customHTTPClient, "oauth.v2.access", values, response, debug); err != nil {
		return nil, err
	}
	return response, response.Err()
}
