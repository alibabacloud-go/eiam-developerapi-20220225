// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateOrganizationalUnitHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateOrganizationalUnitHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationalUnitHeaders) GoString() string {
	return s.String()
}

func (s *CreateOrganizationalUnitHeaders) SetCommonHeaders(v map[string]*string) *CreateOrganizationalUnitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateOrganizationalUnitHeaders) SetAuthorization(v string) *CreateOrganizationalUnitHeaders {
	s.Authorization = &v
	return s
}

type CreateOrganizationalUnitRequest struct {
	Description                  *string `json:"description,omitempty" xml:"description,omitempty"`
	OrganizationalUnitExternalId *string `json:"organizationalUnitExternalId,omitempty" xml:"organizationalUnitExternalId,omitempty"`
	OrganizationalUnitName       *string `json:"organizationalUnitName,omitempty" xml:"organizationalUnitName,omitempty"`
	ParentId                     *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s CreateOrganizationalUnitRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *CreateOrganizationalUnitRequest) SetDescription(v string) *CreateOrganizationalUnitRequest {
	s.Description = &v
	return s
}

func (s *CreateOrganizationalUnitRequest) SetOrganizationalUnitExternalId(v string) *CreateOrganizationalUnitRequest {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *CreateOrganizationalUnitRequest) SetOrganizationalUnitName(v string) *CreateOrganizationalUnitRequest {
	s.OrganizationalUnitName = &v
	return s
}

func (s *CreateOrganizationalUnitRequest) SetParentId(v string) *CreateOrganizationalUnitRequest {
	s.ParentId = &v
	return s
}

type CreateOrganizationalUnitResponseBody struct {
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
}

func (s CreateOrganizationalUnitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrganizationalUnitResponseBody) SetOrganizationalUnitId(v string) *CreateOrganizationalUnitResponseBody {
	s.OrganizationalUnitId = &v
	return s
}

type CreateOrganizationalUnitResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrganizationalUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *CreateOrganizationalUnitResponse) SetHeaders(v map[string]*string) *CreateOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *CreateOrganizationalUnitResponse) SetStatusCode(v int32) *CreateOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrganizationalUnitResponse) SetBody(v *CreateOrganizationalUnitResponseBody) *CreateOrganizationalUnitResponse {
	s.Body = v
	return s
}

type CreateUserHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateUserHeaders) GoString() string {
	return s.String()
}

func (s *CreateUserHeaders) SetCommonHeaders(v map[string]*string) *CreateUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateUserHeaders) SetAuthorization(v string) *CreateUserHeaders {
	s.Authorization = &v
	return s
}

type CreateUserRequest struct {
	Description                 *string `json:"description,omitempty" xml:"description,omitempty"`
	DisplayName                 *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Email                       *string `json:"email,omitempty" xml:"email,omitempty"`
	EmailVerified               *bool   `json:"emailVerified,omitempty" xml:"emailVerified,omitempty"`
	Password                    *string `json:"password,omitempty" xml:"password,omitempty"`
	PhoneNumber                 *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	PhoneNumberVerified         *bool   `json:"phoneNumberVerified,omitempty" xml:"phoneNumberVerified,omitempty"`
	PhoneRegion                 *string `json:"phoneRegion,omitempty" xml:"phoneRegion,omitempty"`
	PrimaryOrganizationalUnitId *string `json:"primaryOrganizationalUnitId,omitempty" xml:"primaryOrganizationalUnitId,omitempty"`
	UserExternalId              *string `json:"userExternalId,omitempty" xml:"userExternalId,omitempty"`
	Username                    *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetDescription(v string) *CreateUserRequest {
	s.Description = &v
	return s
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetEmailVerified(v bool) *CreateUserRequest {
	s.EmailVerified = &v
	return s
}

func (s *CreateUserRequest) SetPassword(v string) *CreateUserRequest {
	s.Password = &v
	return s
}

func (s *CreateUserRequest) SetPhoneNumber(v string) *CreateUserRequest {
	s.PhoneNumber = &v
	return s
}

func (s *CreateUserRequest) SetPhoneNumberVerified(v bool) *CreateUserRequest {
	s.PhoneNumberVerified = &v
	return s
}

func (s *CreateUserRequest) SetPhoneRegion(v string) *CreateUserRequest {
	s.PhoneRegion = &v
	return s
}

func (s *CreateUserRequest) SetPrimaryOrganizationalUnitId(v string) *CreateUserRequest {
	s.PrimaryOrganizationalUnitId = &v
	return s
}

func (s *CreateUserRequest) SetUserExternalId(v string) *CreateUserRequest {
	s.UserExternalId = &v
	return s
}

func (s *CreateUserRequest) SetUsername(v string) *CreateUserRequest {
	s.Username = &v
	return s
}

type CreateUserResponseBody struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) SetUserId(v string) *CreateUserResponseBody {
	s.UserId = &v
	return s
}

type CreateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) SetHeaders(v map[string]*string) *CreateUserResponse {
	s.Headers = v
	return s
}

func (s *CreateUserResponse) SetStatusCode(v int32) *CreateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type DeleteOrganizationalUnitHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteOrganizationalUnitHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrganizationalUnitHeaders) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationalUnitHeaders) SetCommonHeaders(v map[string]*string) *DeleteOrganizationalUnitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteOrganizationalUnitHeaders) SetAuthorization(v string) *DeleteOrganizationalUnitHeaders {
	s.Authorization = &v
	return s
}

type DeleteOrganizationalUnitResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteOrganizationalUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationalUnitResponse) SetHeaders(v map[string]*string) *DeleteOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *DeleteOrganizationalUnitResponse) SetStatusCode(v int32) *DeleteOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

type DeleteUserHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserHeaders) GoString() string {
	return s.String()
}

func (s *DeleteUserHeaders) SetCommonHeaders(v map[string]*string) *DeleteUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteUserHeaders) SetAuthorization(v string) *DeleteUserHeaders {
	s.Authorization = &v
	return s
}

type DeleteUserResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetStatusCode(v int32) *DeleteUserResponse {
	s.StatusCode = &v
	return s
}

type GenerateDeviceCodeRequest struct {
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s GenerateDeviceCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDeviceCodeRequest) GoString() string {
	return s.String()
}

func (s *GenerateDeviceCodeRequest) SetScope(v string) *GenerateDeviceCodeRequest {
	s.Scope = &v
	return s
}

type GenerateDeviceCodeResponseBody struct {
	DeviceCode              *string `json:"device_code,omitempty" xml:"device_code,omitempty"`
	ExpiresAt               *int64  `json:"expires_at,omitempty" xml:"expires_at,omitempty"`
	ExpiresIn               *int64  `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	Interval                *int64  `json:"interval,omitempty" xml:"interval,omitempty"`
	UserCode                *string `json:"user_code,omitempty" xml:"user_code,omitempty"`
	VerificationUri         *string `json:"verification_uri,omitempty" xml:"verification_uri,omitempty"`
	VerificationUriComplete *string `json:"verification_uri_complete,omitempty" xml:"verification_uri_complete,omitempty"`
}

func (s GenerateDeviceCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateDeviceCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDeviceCodeResponseBody) SetDeviceCode(v string) *GenerateDeviceCodeResponseBody {
	s.DeviceCode = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetExpiresAt(v int64) *GenerateDeviceCodeResponseBody {
	s.ExpiresAt = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetExpiresIn(v int64) *GenerateDeviceCodeResponseBody {
	s.ExpiresIn = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetInterval(v int64) *GenerateDeviceCodeResponseBody {
	s.Interval = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetUserCode(v string) *GenerateDeviceCodeResponseBody {
	s.UserCode = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetVerificationUri(v string) *GenerateDeviceCodeResponseBody {
	s.VerificationUri = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetVerificationUriComplete(v string) *GenerateDeviceCodeResponseBody {
	s.VerificationUriComplete = &v
	return s
}

type GenerateDeviceCodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateDeviceCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateDeviceCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateDeviceCodeResponse) GoString() string {
	return s.String()
}

func (s *GenerateDeviceCodeResponse) SetHeaders(v map[string]*string) *GenerateDeviceCodeResponse {
	s.Headers = v
	return s
}

func (s *GenerateDeviceCodeResponse) SetStatusCode(v int32) *GenerateDeviceCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDeviceCodeResponse) SetBody(v *GenerateDeviceCodeResponseBody) *GenerateDeviceCodeResponse {
	s.Body = v
	return s
}

type GenerateTokenRequest struct {
	ClientId     *string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty" xml:"client_secret,omitempty"`
	Code         *string `json:"code,omitempty" xml:"code,omitempty"`
	CodeVerifier *string `json:"code_verifier,omitempty" xml:"code_verifier,omitempty"`
	DeviceCode   *string `json:"device_code,omitempty" xml:"device_code,omitempty"`
	ExclusiveTag *string `json:"exclusive_tag,omitempty" xml:"exclusive_tag,omitempty"`
	GrantType    *string `json:"grant_type,omitempty" xml:"grant_type,omitempty"`
	Password     *string `json:"password,omitempty" xml:"password,omitempty"`
	RedirectUri  *string `json:"redirect_uri,omitempty" xml:"redirect_uri,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	Scope        *string `json:"scope,omitempty" xml:"scope,omitempty"`
	Username     *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GenerateTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateTokenRequest) SetClientId(v string) *GenerateTokenRequest {
	s.ClientId = &v
	return s
}

func (s *GenerateTokenRequest) SetClientSecret(v string) *GenerateTokenRequest {
	s.ClientSecret = &v
	return s
}

func (s *GenerateTokenRequest) SetCode(v string) *GenerateTokenRequest {
	s.Code = &v
	return s
}

func (s *GenerateTokenRequest) SetCodeVerifier(v string) *GenerateTokenRequest {
	s.CodeVerifier = &v
	return s
}

func (s *GenerateTokenRequest) SetDeviceCode(v string) *GenerateTokenRequest {
	s.DeviceCode = &v
	return s
}

func (s *GenerateTokenRequest) SetExclusiveTag(v string) *GenerateTokenRequest {
	s.ExclusiveTag = &v
	return s
}

func (s *GenerateTokenRequest) SetGrantType(v string) *GenerateTokenRequest {
	s.GrantType = &v
	return s
}

func (s *GenerateTokenRequest) SetPassword(v string) *GenerateTokenRequest {
	s.Password = &v
	return s
}

func (s *GenerateTokenRequest) SetRedirectUri(v string) *GenerateTokenRequest {
	s.RedirectUri = &v
	return s
}

func (s *GenerateTokenRequest) SetRefreshToken(v string) *GenerateTokenRequest {
	s.RefreshToken = &v
	return s
}

func (s *GenerateTokenRequest) SetScope(v string) *GenerateTokenRequest {
	s.Scope = &v
	return s
}

func (s *GenerateTokenRequest) SetUsername(v string) *GenerateTokenRequest {
	s.Username = &v
	return s
}

type GenerateTokenResponseBody struct {
	AccessToken  *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ExpiresAt    *int64  `json:"expires_at,omitempty" xml:"expires_at,omitempty"`
	ExpiresIn    *int64  `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	IdToken      *string `json:"id_token,omitempty" xml:"id_token,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	TokenType    *string `json:"token_type,omitempty" xml:"token_type,omitempty"`
}

func (s GenerateTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTokenResponseBody) SetAccessToken(v string) *GenerateTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GenerateTokenResponseBody) SetExpiresAt(v int64) *GenerateTokenResponseBody {
	s.ExpiresAt = &v
	return s
}

func (s *GenerateTokenResponseBody) SetExpiresIn(v int64) *GenerateTokenResponseBody {
	s.ExpiresIn = &v
	return s
}

func (s *GenerateTokenResponseBody) SetIdToken(v string) *GenerateTokenResponseBody {
	s.IdToken = &v
	return s
}

func (s *GenerateTokenResponseBody) SetRefreshToken(v string) *GenerateTokenResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *GenerateTokenResponseBody) SetTokenType(v string) *GenerateTokenResponseBody {
	s.TokenType = &v
	return s
}

type GenerateTokenResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateTokenResponse) SetHeaders(v map[string]*string) *GenerateTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateTokenResponse) SetStatusCode(v int32) *GenerateTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateTokenResponse) SetBody(v *GenerateTokenResponseBody) *GenerateTokenResponse {
	s.Body = v
	return s
}

type GetApplicationProvisioningScopeHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetApplicationProvisioningScopeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningScopeHeaders) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeHeaders) SetCommonHeaders(v map[string]*string) *GetApplicationProvisioningScopeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetApplicationProvisioningScopeHeaders) SetAuthorization(v string) *GetApplicationProvisioningScopeHeaders {
	s.Authorization = &v
	return s
}

type GetApplicationProvisioningScopeResponseBody struct {
	OrganizationalUnitIds []*string `json:"organizationalUnitIds,omitempty" xml:"organizationalUnitIds,omitempty" type:"Repeated"`
}

func (s GetApplicationProvisioningScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningScopeResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeResponseBody) SetOrganizationalUnitIds(v []*string) *GetApplicationProvisioningScopeResponseBody {
	s.OrganizationalUnitIds = v
	return s
}

type GetApplicationProvisioningScopeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetApplicationProvisioningScopeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetApplicationProvisioningScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningScopeResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeResponse) SetHeaders(v map[string]*string) *GetApplicationProvisioningScopeResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationProvisioningScopeResponse) SetStatusCode(v int32) *GetApplicationProvisioningScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationProvisioningScopeResponse) SetBody(v *GetApplicationProvisioningScopeResponseBody) *GetApplicationProvisioningScopeResponse {
	s.Body = v
	return s
}

type GetOrganizationalUnitHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetOrganizationalUnitHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitHeaders) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitHeaders) SetCommonHeaders(v map[string]*string) *GetOrganizationalUnitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrganizationalUnitHeaders) SetAuthorization(v string) *GetOrganizationalUnitHeaders {
	s.Authorization = &v
	return s
}

type GetOrganizationalUnitResponseBody struct {
	CreateTime                   *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description                  *string `json:"description,omitempty" xml:"description,omitempty"`
	InstanceId                   *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	OrganizationalUnitExternalId *string `json:"organizationalUnitExternalId,omitempty" xml:"organizationalUnitExternalId,omitempty"`
	OrganizationalUnitId         *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
	OrganizationalUnitName       *string `json:"organizationalUnitName,omitempty" xml:"organizationalUnitName,omitempty"`
	OrganizationalUnitSourceId   *string `json:"organizationalUnitSourceId,omitempty" xml:"organizationalUnitSourceId,omitempty"`
	OrganizationalUnitSourceType *string `json:"organizationalUnitSourceType,omitempty" xml:"organizationalUnitSourceType,omitempty"`
	ParentId                     *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	UpdateTime                   *int64  `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetOrganizationalUnitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitResponseBody) SetCreateTime(v int64) *GetOrganizationalUnitResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetDescription(v string) *GetOrganizationalUnitResponseBody {
	s.Description = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetInstanceId(v string) *GetOrganizationalUnitResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnitExternalId(v string) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnitId(v string) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnitId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnitName(v string) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnitName = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnitSourceId(v string) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnitSourceId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnitSourceType(v string) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnitSourceType = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetParentId(v string) *GetOrganizationalUnitResponseBody {
	s.ParentId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetUpdateTime(v int64) *GetOrganizationalUnitResponseBody {
	s.UpdateTime = &v
	return s
}

type GetOrganizationalUnitResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrganizationalUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitResponse) SetHeaders(v map[string]*string) *GetOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationalUnitResponse) SetStatusCode(v int32) *GetOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrganizationalUnitResponse) SetBody(v *GetOrganizationalUnitResponseBody) *GetOrganizationalUnitResponse {
	s.Body = v
	return s
}

type GetOrganizationalUnitIdByExternalIdHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetOrganizationalUnitIdByExternalIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitIdByExternalIdHeaders) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitIdByExternalIdHeaders) SetCommonHeaders(v map[string]*string) *GetOrganizationalUnitIdByExternalIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdHeaders) SetAuthorization(v string) *GetOrganizationalUnitIdByExternalIdHeaders {
	s.Authorization = &v
	return s
}

type GetOrganizationalUnitIdByExternalIdRequest struct {
	OrganizationalUnitExternalId *string `json:"organizationalUnitExternalId,omitempty" xml:"organizationalUnitExternalId,omitempty"`
	OrganizationalUnitSourceId   *string `json:"organizationalUnitSourceId,omitempty" xml:"organizationalUnitSourceId,omitempty"`
	OrganizationalUnitSourceType *string `json:"organizationalUnitSourceType,omitempty" xml:"organizationalUnitSourceType,omitempty"`
}

func (s GetOrganizationalUnitIdByExternalIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitIdByExternalIdRequest) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitIdByExternalIdRequest) SetOrganizationalUnitExternalId(v string) *GetOrganizationalUnitIdByExternalIdRequest {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdRequest) SetOrganizationalUnitSourceId(v string) *GetOrganizationalUnitIdByExternalIdRequest {
	s.OrganizationalUnitSourceId = &v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdRequest) SetOrganizationalUnitSourceType(v string) *GetOrganizationalUnitIdByExternalIdRequest {
	s.OrganizationalUnitSourceType = &v
	return s
}

type GetOrganizationalUnitIdByExternalIdResponseBody struct {
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
}

func (s GetOrganizationalUnitIdByExternalIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitIdByExternalIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitIdByExternalIdResponseBody) SetOrganizationalUnitId(v string) *GetOrganizationalUnitIdByExternalIdResponseBody {
	s.OrganizationalUnitId = &v
	return s
}

type GetOrganizationalUnitIdByExternalIdResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOrganizationalUnitIdByExternalIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrganizationalUnitIdByExternalIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitIdByExternalIdResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitIdByExternalIdResponse) SetHeaders(v map[string]*string) *GetOrganizationalUnitIdByExternalIdResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdResponse) SetStatusCode(v int32) *GetOrganizationalUnitIdByExternalIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdResponse) SetBody(v *GetOrganizationalUnitIdByExternalIdResponseBody) *GetOrganizationalUnitIdByExternalIdResponse {
	s.Body = v
	return s
}

type GetUserHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserHeaders) GoString() string {
	return s.String()
}

func (s *GetUserHeaders) SetCommonHeaders(v map[string]*string) *GetUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserHeaders) SetAuthorization(v string) *GetUserHeaders {
	s.Authorization = &v
	return s
}

type GetUserResponseBody struct {
	AccountExpireTime           *int64                                    `json:"accountExpireTime,omitempty" xml:"accountExpireTime,omitempty"`
	CreateTime                  *int64                                    `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description                 *string                                   `json:"description,omitempty" xml:"description,omitempty"`
	DisplayName                 *string                                   `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Email                       *string                                   `json:"email,omitempty" xml:"email,omitempty"`
	EmailVerified               *bool                                     `json:"emailVerified,omitempty" xml:"emailVerified,omitempty"`
	InstanceId                  *string                                   `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	LockExpireTime              *int64                                    `json:"lockExpireTime,omitempty" xml:"lockExpireTime,omitempty"`
	OrganizationalUnits         []*GetUserResponseBodyOrganizationalUnits `json:"organizationalUnits,omitempty" xml:"organizationalUnits,omitempty" type:"Repeated"`
	PhoneNumber                 *string                                   `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	PhoneNumberVerified         *bool                                     `json:"phoneNumberVerified,omitempty" xml:"phoneNumberVerified,omitempty"`
	PhoneRegion                 *string                                   `json:"phoneRegion,omitempty" xml:"phoneRegion,omitempty"`
	PrimaryOrganizationalUnitId *string                                   `json:"primaryOrganizationalUnitId,omitempty" xml:"primaryOrganizationalUnitId,omitempty"`
	RegisterTime                *int64                                    `json:"registerTime,omitempty" xml:"registerTime,omitempty"`
	Status                      *string                                   `json:"status,omitempty" xml:"status,omitempty"`
	UpdateTime                  *int64                                    `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserExternalId              *string                                   `json:"userExternalId,omitempty" xml:"userExternalId,omitempty"`
	UserId                      *string                                   `json:"userId,omitempty" xml:"userId,omitempty"`
	UserSourceId                *string                                   `json:"userSourceId,omitempty" xml:"userSourceId,omitempty"`
	UserSourceType              *string                                   `json:"userSourceType,omitempty" xml:"userSourceType,omitempty"`
	Username                    *string                                   `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetAccountExpireTime(v int64) *GetUserResponseBody {
	s.AccountExpireTime = &v
	return s
}

func (s *GetUserResponseBody) SetCreateTime(v int64) *GetUserResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetUserResponseBody) SetDescription(v string) *GetUserResponseBody {
	s.Description = &v
	return s
}

func (s *GetUserResponseBody) SetDisplayName(v string) *GetUserResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetUserResponseBody) SetEmail(v string) *GetUserResponseBody {
	s.Email = &v
	return s
}

func (s *GetUserResponseBody) SetEmailVerified(v bool) *GetUserResponseBody {
	s.EmailVerified = &v
	return s
}

func (s *GetUserResponseBody) SetInstanceId(v string) *GetUserResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetUserResponseBody) SetLockExpireTime(v int64) *GetUserResponseBody {
	s.LockExpireTime = &v
	return s
}

func (s *GetUserResponseBody) SetOrganizationalUnits(v []*GetUserResponseBodyOrganizationalUnits) *GetUserResponseBody {
	s.OrganizationalUnits = v
	return s
}

func (s *GetUserResponseBody) SetPhoneNumber(v string) *GetUserResponseBody {
	s.PhoneNumber = &v
	return s
}

func (s *GetUserResponseBody) SetPhoneNumberVerified(v bool) *GetUserResponseBody {
	s.PhoneNumberVerified = &v
	return s
}

func (s *GetUserResponseBody) SetPhoneRegion(v string) *GetUserResponseBody {
	s.PhoneRegion = &v
	return s
}

func (s *GetUserResponseBody) SetPrimaryOrganizationalUnitId(v string) *GetUserResponseBody {
	s.PrimaryOrganizationalUnitId = &v
	return s
}

func (s *GetUserResponseBody) SetRegisterTime(v int64) *GetUserResponseBody {
	s.RegisterTime = &v
	return s
}

func (s *GetUserResponseBody) SetStatus(v string) *GetUserResponseBody {
	s.Status = &v
	return s
}

func (s *GetUserResponseBody) SetUpdateTime(v int64) *GetUserResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetUserResponseBody) SetUserExternalId(v string) *GetUserResponseBody {
	s.UserExternalId = &v
	return s
}

func (s *GetUserResponseBody) SetUserId(v string) *GetUserResponseBody {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBody) SetUserSourceId(v string) *GetUserResponseBody {
	s.UserSourceId = &v
	return s
}

func (s *GetUserResponseBody) SetUserSourceType(v string) *GetUserResponseBody {
	s.UserSourceType = &v
	return s
}

func (s *GetUserResponseBody) SetUsername(v string) *GetUserResponseBody {
	s.Username = &v
	return s
}

type GetUserResponseBodyOrganizationalUnits struct {
	OrganizationalUnitId   *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
	OrganizationalUnitName *string `json:"organizationalUnitName,omitempty" xml:"organizationalUnitName,omitempty"`
	Primary                *bool   `json:"primary,omitempty" xml:"primary,omitempty"`
}

func (s GetUserResponseBodyOrganizationalUnits) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyOrganizationalUnits) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyOrganizationalUnits) SetOrganizationalUnitId(v string) *GetUserResponseBodyOrganizationalUnits {
	s.OrganizationalUnitId = &v
	return s
}

func (s *GetUserResponseBodyOrganizationalUnits) SetOrganizationalUnitName(v string) *GetUserResponseBodyOrganizationalUnits {
	s.OrganizationalUnitName = &v
	return s
}

func (s *GetUserResponseBodyOrganizationalUnits) SetPrimary(v bool) *GetUserResponseBodyOrganizationalUnits {
	s.Primary = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type GetUserIdByEmailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUserIdByEmailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByEmailHeaders) GoString() string {
	return s.String()
}

func (s *GetUserIdByEmailHeaders) SetCommonHeaders(v map[string]*string) *GetUserIdByEmailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserIdByEmailHeaders) SetAuthorization(v string) *GetUserIdByEmailHeaders {
	s.Authorization = &v
	return s
}

type GetUserIdByEmailRequest struct {
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
}

func (s GetUserIdByEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByEmailRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdByEmailRequest) SetEmail(v string) *GetUserIdByEmailRequest {
	s.Email = &v
	return s
}

type GetUserIdByEmailResponseBody struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserIdByEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByEmailResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserIdByEmailResponseBody) SetUserId(v string) *GetUserIdByEmailResponseBody {
	s.UserId = &v
	return s
}

type GetUserIdByEmailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserIdByEmailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserIdByEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByEmailResponse) GoString() string {
	return s.String()
}

func (s *GetUserIdByEmailResponse) SetHeaders(v map[string]*string) *GetUserIdByEmailResponse {
	s.Headers = v
	return s
}

func (s *GetUserIdByEmailResponse) SetStatusCode(v int32) *GetUserIdByEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserIdByEmailResponse) SetBody(v *GetUserIdByEmailResponseBody) *GetUserIdByEmailResponse {
	s.Body = v
	return s
}

type GetUserIdByPhoneNumberHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUserIdByPhoneNumberHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByPhoneNumberHeaders) GoString() string {
	return s.String()
}

func (s *GetUserIdByPhoneNumberHeaders) SetCommonHeaders(v map[string]*string) *GetUserIdByPhoneNumberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserIdByPhoneNumberHeaders) SetAuthorization(v string) *GetUserIdByPhoneNumberHeaders {
	s.Authorization = &v
	return s
}

type GetUserIdByPhoneNumberRequest struct {
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
}

func (s GetUserIdByPhoneNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdByPhoneNumberRequest) SetPhoneNumber(v string) *GetUserIdByPhoneNumberRequest {
	s.PhoneNumber = &v
	return s
}

type GetUserIdByPhoneNumberResponseBody struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserIdByPhoneNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserIdByPhoneNumberResponseBody) SetUserId(v string) *GetUserIdByPhoneNumberResponseBody {
	s.UserId = &v
	return s
}

type GetUserIdByPhoneNumberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserIdByPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserIdByPhoneNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *GetUserIdByPhoneNumberResponse) SetHeaders(v map[string]*string) *GetUserIdByPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *GetUserIdByPhoneNumberResponse) SetStatusCode(v int32) *GetUserIdByPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserIdByPhoneNumberResponse) SetBody(v *GetUserIdByPhoneNumberResponseBody) *GetUserIdByPhoneNumberResponse {
	s.Body = v
	return s
}

type GetUserIdByUserExternalIdHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUserIdByUserExternalIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByUserExternalIdHeaders) GoString() string {
	return s.String()
}

func (s *GetUserIdByUserExternalIdHeaders) SetCommonHeaders(v map[string]*string) *GetUserIdByUserExternalIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserIdByUserExternalIdHeaders) SetAuthorization(v string) *GetUserIdByUserExternalIdHeaders {
	s.Authorization = &v
	return s
}

type GetUserIdByUserExternalIdRequest struct {
	UserExternalId *string `json:"userExternalId,omitempty" xml:"userExternalId,omitempty"`
	UserSourceId   *string `json:"userSourceId,omitempty" xml:"userSourceId,omitempty"`
	UserSourceType *string `json:"userSourceType,omitempty" xml:"userSourceType,omitempty"`
}

func (s GetUserIdByUserExternalIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByUserExternalIdRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdByUserExternalIdRequest) SetUserExternalId(v string) *GetUserIdByUserExternalIdRequest {
	s.UserExternalId = &v
	return s
}

func (s *GetUserIdByUserExternalIdRequest) SetUserSourceId(v string) *GetUserIdByUserExternalIdRequest {
	s.UserSourceId = &v
	return s
}

func (s *GetUserIdByUserExternalIdRequest) SetUserSourceType(v string) *GetUserIdByUserExternalIdRequest {
	s.UserSourceType = &v
	return s
}

type GetUserIdByUserExternalIdResponseBody struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserIdByUserExternalIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByUserExternalIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserIdByUserExternalIdResponseBody) SetUserId(v string) *GetUserIdByUserExternalIdResponseBody {
	s.UserId = &v
	return s
}

type GetUserIdByUserExternalIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserIdByUserExternalIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserIdByUserExternalIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByUserExternalIdResponse) GoString() string {
	return s.String()
}

func (s *GetUserIdByUserExternalIdResponse) SetHeaders(v map[string]*string) *GetUserIdByUserExternalIdResponse {
	s.Headers = v
	return s
}

func (s *GetUserIdByUserExternalIdResponse) SetStatusCode(v int32) *GetUserIdByUserExternalIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserIdByUserExternalIdResponse) SetBody(v *GetUserIdByUserExternalIdResponseBody) *GetUserIdByUserExternalIdResponse {
	s.Body = v
	return s
}

type GetUserInfoHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUserInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetUserInfoHeaders) SetCommonHeaders(v map[string]*string) *GetUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserInfoHeaders) SetAuthorization(v string) *GetUserInfoHeaders {
	s.Authorization = &v
	return s
}

type GetUserInfoResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponse) SetHeaders(v map[string]*string) *GetUserInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUserInfoResponse) SetStatusCode(v int32) *GetUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserInfoResponse) SetBody(v map[string]interface{}) *GetUserInfoResponse {
	s.Body = v
	return s
}

type ListOrganizationalUnitParentIdsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListOrganizationalUnitParentIdsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitParentIdsHeaders) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentIdsHeaders) SetCommonHeaders(v map[string]*string) *ListOrganizationalUnitParentIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListOrganizationalUnitParentIdsHeaders) SetAuthorization(v string) *ListOrganizationalUnitParentIdsHeaders {
	s.Authorization = &v
	return s
}

type ListOrganizationalUnitParentIdsResponseBody struct {
	ParentIds []*string `json:"parentIds,omitempty" xml:"parentIds,omitempty" type:"Repeated"`
}

func (s ListOrganizationalUnitParentIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitParentIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentIdsResponseBody) SetParentIds(v []*string) *ListOrganizationalUnitParentIdsResponseBody {
	s.ParentIds = v
	return s
}

type ListOrganizationalUnitParentIdsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOrganizationalUnitParentIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrganizationalUnitParentIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitParentIdsResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentIdsResponse) SetHeaders(v map[string]*string) *ListOrganizationalUnitParentIdsResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationalUnitParentIdsResponse) SetStatusCode(v int32) *ListOrganizationalUnitParentIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationalUnitParentIdsResponse) SetBody(v *ListOrganizationalUnitParentIdsResponseBody) *ListOrganizationalUnitParentIdsResponse {
	s.Body = v
	return s
}

type ListOrganizationalUnitsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListOrganizationalUnitsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsHeaders) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsHeaders) SetCommonHeaders(v map[string]*string) *ListOrganizationalUnitsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListOrganizationalUnitsHeaders) SetAuthorization(v string) *ListOrganizationalUnitsHeaders {
	s.Authorization = &v
	return s
}

type ListOrganizationalUnitsRequest struct {
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ParentId   *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s ListOrganizationalUnitsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsRequest) SetPageNumber(v int32) *ListOrganizationalUnitsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetPageSize(v int32) *ListOrganizationalUnitsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetParentId(v string) *ListOrganizationalUnitsRequest {
	s.ParentId = &v
	return s
}

type ListOrganizationalUnitsResponseBody struct {
	Data       []*ListOrganizationalUnitsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	TotalCount *int64                                     `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListOrganizationalUnitsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsResponseBody) SetData(v []*ListOrganizationalUnitsResponseBodyData) *ListOrganizationalUnitsResponseBody {
	s.Data = v
	return s
}

func (s *ListOrganizationalUnitsResponseBody) SetTotalCount(v int64) *ListOrganizationalUnitsResponseBody {
	s.TotalCount = &v
	return s
}

type ListOrganizationalUnitsResponseBodyData struct {
	CreateTime                   *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description                  *string `json:"description,omitempty" xml:"description,omitempty"`
	InstanceId                   *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	OrganizationalUnitExternalId *string `json:"organizationalUnitExternalId,omitempty" xml:"organizationalUnitExternalId,omitempty"`
	OrganizationalUnitId         *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
	OrganizationalUnitName       *string `json:"organizationalUnitName,omitempty" xml:"organizationalUnitName,omitempty"`
	OrganizationalUnitSourceId   *string `json:"organizationalUnitSourceId,omitempty" xml:"organizationalUnitSourceId,omitempty"`
	OrganizationalUnitSourceType *string `json:"organizationalUnitSourceType,omitempty" xml:"organizationalUnitSourceType,omitempty"`
	ParentId                     *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	UpdateTime                   *int64  `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListOrganizationalUnitsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsResponseBodyData) SetCreateTime(v int64) *ListOrganizationalUnitsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetDescription(v string) *ListOrganizationalUnitsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetInstanceId(v string) *ListOrganizationalUnitsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetOrganizationalUnitExternalId(v string) *ListOrganizationalUnitsResponseBodyData {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetOrganizationalUnitId(v string) *ListOrganizationalUnitsResponseBodyData {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetOrganizationalUnitName(v string) *ListOrganizationalUnitsResponseBodyData {
	s.OrganizationalUnitName = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetOrganizationalUnitSourceId(v string) *ListOrganizationalUnitsResponseBodyData {
	s.OrganizationalUnitSourceId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetOrganizationalUnitSourceType(v string) *ListOrganizationalUnitsResponseBodyData {
	s.OrganizationalUnitSourceType = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetParentId(v string) *ListOrganizationalUnitsResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetUpdateTime(v int64) *ListOrganizationalUnitsResponseBodyData {
	s.UpdateTime = &v
	return s
}

type ListOrganizationalUnitsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOrganizationalUnitsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrganizationalUnitsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsResponse) SetHeaders(v map[string]*string) *ListOrganizationalUnitsResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationalUnitsResponse) SetStatusCode(v int32) *ListOrganizationalUnitsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationalUnitsResponse) SetBody(v *ListOrganizationalUnitsResponseBody) *ListOrganizationalUnitsResponse {
	s.Body = v
	return s
}

type ListUsersHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListUsersHeaders) GoString() string {
	return s.String()
}

func (s *ListUsersHeaders) SetCommonHeaders(v map[string]*string) *ListUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListUsersHeaders) SetAuthorization(v string) *ListUsersHeaders {
	s.Authorization = &v
	return s
}

type ListUsersRequest struct {
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
	PageNumber           *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize             *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetOrganizationalUnitId(v string) *ListUsersRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v int32) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

type ListUsersResponseBody struct {
	Data       []*ListUsersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	TotalCount *int64                       `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetData(v []*ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int64) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

type ListUsersResponseBodyData struct {
	AccountExpireTime   *int64  `json:"accountExpireTime,omitempty" xml:"accountExpireTime,omitempty"`
	CreateTime          *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description         *string `json:"description,omitempty" xml:"description,omitempty"`
	DisplayName         *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Email               *string `json:"email,omitempty" xml:"email,omitempty"`
	EmailVerified       *bool   `json:"emailVerified,omitempty" xml:"emailVerified,omitempty"`
	InstanceId          *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	LockExpireTime      *int64  `json:"lockExpireTime,omitempty" xml:"lockExpireTime,omitempty"`
	PhoneNumber         *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	PhoneNumberVerified *bool   `json:"phoneNumberVerified,omitempty" xml:"phoneNumberVerified,omitempty"`
	PhoneRegion         *string `json:"phoneRegion,omitempty" xml:"phoneRegion,omitempty"`
	RegisterTime        *int64  `json:"registerTime,omitempty" xml:"registerTime,omitempty"`
	Status              *string `json:"status,omitempty" xml:"status,omitempty"`
	UpdateTime          *int64  `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserExternalId      *string `json:"userExternalId,omitempty" xml:"userExternalId,omitempty"`
	UserId              *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UserSourceId        *string `json:"userSourceId,omitempty" xml:"userSourceId,omitempty"`
	UserSourceType      *string `json:"userSourceType,omitempty" xml:"userSourceType,omitempty"`
	Username            *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) SetAccountExpireTime(v int64) *ListUsersResponseBodyData {
	s.AccountExpireTime = &v
	return s
}

func (s *ListUsersResponseBodyData) SetCreateTime(v int64) *ListUsersResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListUsersResponseBodyData) SetDescription(v string) *ListUsersResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListUsersResponseBodyData) SetDisplayName(v string) *ListUsersResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyData) SetEmail(v string) *ListUsersResponseBodyData {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyData) SetEmailVerified(v bool) *ListUsersResponseBodyData {
	s.EmailVerified = &v
	return s
}

func (s *ListUsersResponseBodyData) SetInstanceId(v string) *ListUsersResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListUsersResponseBodyData) SetLockExpireTime(v int64) *ListUsersResponseBodyData {
	s.LockExpireTime = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPhoneNumber(v string) *ListUsersResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPhoneNumberVerified(v bool) *ListUsersResponseBodyData {
	s.PhoneNumberVerified = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPhoneRegion(v string) *ListUsersResponseBodyData {
	s.PhoneRegion = &v
	return s
}

func (s *ListUsersResponseBodyData) SetRegisterTime(v int64) *ListUsersResponseBodyData {
	s.RegisterTime = &v
	return s
}

func (s *ListUsersResponseBodyData) SetStatus(v string) *ListUsersResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUpdateTime(v int64) *ListUsersResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUserExternalId(v string) *ListUsersResponseBodyData {
	s.UserExternalId = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUserId(v string) *ListUsersResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUserSourceId(v string) *ListUsersResponseBodyData {
	s.UserSourceId = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUserSourceType(v string) *ListUsersResponseBodyData {
	s.UserSourceType = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUsername(v string) *ListUsersResponseBodyData {
	s.Username = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type PatchOrganizationalUnitHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PatchOrganizationalUnitHeaders) String() string {
	return tea.Prettify(s)
}

func (s PatchOrganizationalUnitHeaders) GoString() string {
	return s.String()
}

func (s *PatchOrganizationalUnitHeaders) SetCommonHeaders(v map[string]*string) *PatchOrganizationalUnitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PatchOrganizationalUnitHeaders) SetAuthorization(v string) *PatchOrganizationalUnitHeaders {
	s.Authorization = &v
	return s
}

type PatchOrganizationalUnitRequest struct {
	Description            *string `json:"description,omitempty" xml:"description,omitempty"`
	OrganizationalUnitName *string `json:"organizationalUnitName,omitempty" xml:"organizationalUnitName,omitempty"`
}

func (s PatchOrganizationalUnitRequest) String() string {
	return tea.Prettify(s)
}

func (s PatchOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *PatchOrganizationalUnitRequest) SetDescription(v string) *PatchOrganizationalUnitRequest {
	s.Description = &v
	return s
}

func (s *PatchOrganizationalUnitRequest) SetOrganizationalUnitName(v string) *PatchOrganizationalUnitRequest {
	s.OrganizationalUnitName = &v
	return s
}

type PatchOrganizationalUnitResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PatchOrganizationalUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s PatchOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *PatchOrganizationalUnitResponse) SetHeaders(v map[string]*string) *PatchOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *PatchOrganizationalUnitResponse) SetStatusCode(v int32) *PatchOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

type PatchUserHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Authorization *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PatchUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s PatchUserHeaders) GoString() string {
	return s.String()
}

func (s *PatchUserHeaders) SetCommonHeaders(v map[string]*string) *PatchUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PatchUserHeaders) SetAuthorization(v string) *PatchUserHeaders {
	s.Authorization = &v
	return s
}

type PatchUserRequest struct {
	DisplayName         *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Email               *string `json:"email,omitempty" xml:"email,omitempty"`
	EmailVerified       *bool   `json:"emailVerified,omitempty" xml:"emailVerified,omitempty"`
	PhoneNumber         *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	PhoneNumberVerified *bool   `json:"phoneNumberVerified,omitempty" xml:"phoneNumberVerified,omitempty"`
	PhoneRegion         *string `json:"phoneRegion,omitempty" xml:"phoneRegion,omitempty"`
	Username            *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s PatchUserRequest) String() string {
	return tea.Prettify(s)
}

func (s PatchUserRequest) GoString() string {
	return s.String()
}

func (s *PatchUserRequest) SetDisplayName(v string) *PatchUserRequest {
	s.DisplayName = &v
	return s
}

func (s *PatchUserRequest) SetEmail(v string) *PatchUserRequest {
	s.Email = &v
	return s
}

func (s *PatchUserRequest) SetEmailVerified(v bool) *PatchUserRequest {
	s.EmailVerified = &v
	return s
}

func (s *PatchUserRequest) SetPhoneNumber(v string) *PatchUserRequest {
	s.PhoneNumber = &v
	return s
}

func (s *PatchUserRequest) SetPhoneNumberVerified(v bool) *PatchUserRequest {
	s.PhoneNumberVerified = &v
	return s
}

func (s *PatchUserRequest) SetPhoneRegion(v string) *PatchUserRequest {
	s.PhoneRegion = &v
	return s
}

func (s *PatchUserRequest) SetUsername(v string) *PatchUserRequest {
	s.Username = &v
	return s
}

type PatchUserResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PatchUserResponse) String() string {
	return tea.Prettify(s)
}

func (s PatchUserResponse) GoString() string {
	return s.String()
}

func (s *PatchUserResponse) SetHeaders(v map[string]*string) *PatchUserResponse {
	s.Headers = v
	return s
}

func (s *PatchUserResponse) SetStatusCode(v int32) *PatchUserResponse {
	s.StatusCode = &v
	return s
}

type RevokeTokenRequest struct {
	ClientId      *string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	ClientSecret  *string `json:"client_secret,omitempty" xml:"client_secret,omitempty"`
	Token         *string `json:"token,omitempty" xml:"token,omitempty"`
	TokenTypeHint *string `json:"token_type_hint,omitempty" xml:"token_type_hint,omitempty"`
}

func (s RevokeTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeTokenRequest) GoString() string {
	return s.String()
}

func (s *RevokeTokenRequest) SetClientId(v string) *RevokeTokenRequest {
	s.ClientId = &v
	return s
}

func (s *RevokeTokenRequest) SetClientSecret(v string) *RevokeTokenRequest {
	s.ClientSecret = &v
	return s
}

func (s *RevokeTokenRequest) SetToken(v string) *RevokeTokenRequest {
	s.Token = &v
	return s
}

func (s *RevokeTokenRequest) SetTokenTypeHint(v string) *RevokeTokenRequest {
	s.TokenTypeHint = &v
	return s
}

type RevokeTokenResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokeTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeTokenResponse) GoString() string {
	return s.String()
}

func (s *RevokeTokenResponse) SetHeaders(v map[string]*string) *RevokeTokenResponse {
	s.Headers = v
	return s
}

func (s *RevokeTokenResponse) SetStatusCode(v int32) *RevokeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeTokenResponse) SetBody(v map[string]interface{}) *RevokeTokenResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("eiam-developerapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrganizationalUnit(instanceId *string, applicationId *string, request *CreateOrganizationalUnitRequest) (_result *CreateOrganizationalUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateOrganizationalUnitHeaders{}
	_result = &CreateOrganizationalUnitResponse{}
	_body, _err := client.CreateOrganizationalUnitWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrganizationalUnitWithOptions(instanceId *string, applicationId *string, request *CreateOrganizationalUnitRequest, headers *CreateOrganizationalUnitHeaders, runtime *util.RuntimeOptions) (_result *CreateOrganizationalUnitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitExternalId)) {
		body["organizationalUnitExternalId"] = request.OrganizationalUnitExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitName)) {
		body["organizationalUnitName"] = request.OrganizationalUnitName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		body["parentId"] = request.ParentId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrganizationalUnit"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/organizationalUnits"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUser(instanceId *string, applicationId *string, request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateUserHeaders{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserWithOptions(instanceId *string, applicationId *string, request *CreateUserRequest, headers *CreateUserHeaders, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.EmailVerified)) {
		body["emailVerified"] = request.EmailVerified
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["phoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumberVerified)) {
		body["phoneNumberVerified"] = request.PhoneNumberVerified
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneRegion)) {
		body["phoneRegion"] = request.PhoneRegion
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryOrganizationalUnitId)) {
		body["primaryOrganizationalUnitId"] = request.PrimaryOrganizationalUnitId
	}

	if !tea.BoolValue(util.IsUnset(request.UserExternalId)) {
		body["userExternalId"] = request.UserExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["username"] = request.Username
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUser"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/users"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteOrganizationalUnit(instanceId *string, applicationId *string, organizationalUnitId *string) (_result *DeleteOrganizationalUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteOrganizationalUnitHeaders{}
	_result = &DeleteOrganizationalUnitResponse{}
	_body, _err := client.DeleteOrganizationalUnitWithOptions(instanceId, applicationId, organizationalUnitId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteOrganizationalUnitWithOptions(instanceId *string, applicationId *string, organizationalUnitId *string, headers *DeleteOrganizationalUnitHeaders, runtime *util.RuntimeOptions) (_result *DeleteOrganizationalUnitResponse, _err error) {
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	organizationalUnitId = openapiutil.GetEncodeParam(organizationalUnitId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteOrganizationalUnit"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/organizationalUnits/" + tea.StringValue(organizationalUnitId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUser(instanceId *string, applicationId *string, userId *string) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteUserHeaders{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(instanceId, applicationId, userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserWithOptions(instanceId *string, applicationId *string, userId *string, headers *DeleteUserHeaders, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	userId = openapiutil.GetEncodeParam(userId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/users/" + tea.StringValue(userId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateDeviceCode(instanceId *string, applicationId *string, request *GenerateDeviceCodeRequest) (_result *GenerateDeviceCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateDeviceCodeResponse{}
	_body, _err := client.GenerateDeviceCodeWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateDeviceCodeWithOptions(instanceId *string, applicationId *string, request *GenerateDeviceCodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateDeviceCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["scope"] = request.Scope
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateDeviceCode"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/oauth2/device/code"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateDeviceCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateToken(instanceId *string, applicationId *string, request *GenerateTokenRequest) (_result *GenerateTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateTokenResponse{}
	_body, _err := client.GenerateTokenWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateTokenWithOptions(instanceId *string, applicationId *string, request *GenerateTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["client_id"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSecret)) {
		query["client_secret"] = request.ClientSecret
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.CodeVerifier)) {
		query["code_verifier"] = request.CodeVerifier
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceCode)) {
		query["device_code"] = request.DeviceCode
	}

	if !tea.BoolValue(util.IsUnset(request.ExclusiveTag)) {
		query["exclusive_tag"] = request.ExclusiveTag
	}

	if !tea.BoolValue(util.IsUnset(request.GrantType)) {
		query["grant_type"] = request.GrantType
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RedirectUri)) {
		query["redirect_uri"] = request.RedirectUri
	}

	if !tea.BoolValue(util.IsUnset(request.RefreshToken)) {
		query["refresh_token"] = request.RefreshToken
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateToken"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/oauth2/token"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApplicationProvisioningScope(instanceId *string, applicationId *string) (_result *GetApplicationProvisioningScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetApplicationProvisioningScopeHeaders{}
	_result = &GetApplicationProvisioningScopeResponse{}
	_body, _err := client.GetApplicationProvisioningScopeWithOptions(instanceId, applicationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetApplicationProvisioningScopeWithOptions(instanceId *string, applicationId *string, headers *GetApplicationProvisioningScopeHeaders, runtime *util.RuntimeOptions) (_result *GetApplicationProvisioningScopeResponse, _err error) {
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplicationProvisioningScope"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/provisioningScope"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplicationProvisioningScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrganizationalUnit(instanceId *string, applicationId *string, organizationalUnitId *string) (_result *GetOrganizationalUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOrganizationalUnitHeaders{}
	_result = &GetOrganizationalUnitResponse{}
	_body, _err := client.GetOrganizationalUnitWithOptions(instanceId, applicationId, organizationalUnitId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrganizationalUnitWithOptions(instanceId *string, applicationId *string, organizationalUnitId *string, headers *GetOrganizationalUnitHeaders, runtime *util.RuntimeOptions) (_result *GetOrganizationalUnitResponse, _err error) {
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	organizationalUnitId = openapiutil.GetEncodeParam(organizationalUnitId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetOrganizationalUnit"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/organizationalUnits/" + tea.StringValue(organizationalUnitId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrganizationalUnitIdByExternalId(instanceId *string, applicationId *string, request *GetOrganizationalUnitIdByExternalIdRequest) (_result *GetOrganizationalUnitIdByExternalIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOrganizationalUnitIdByExternalIdHeaders{}
	_result = &GetOrganizationalUnitIdByExternalIdResponse{}
	_body, _err := client.GetOrganizationalUnitIdByExternalIdWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrganizationalUnitIdByExternalIdWithOptions(instanceId *string, applicationId *string, request *GetOrganizationalUnitIdByExternalIdRequest, headers *GetOrganizationalUnitIdByExternalIdHeaders, runtime *util.RuntimeOptions) (_result *GetOrganizationalUnitIdByExternalIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitExternalId)) {
		body["organizationalUnitExternalId"] = request.OrganizationalUnitExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitSourceId)) {
		body["organizationalUnitSourceId"] = request.OrganizationalUnitSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitSourceType)) {
		body["organizationalUnitSourceType"] = request.OrganizationalUnitSourceType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOrganizationalUnitIdByExternalId"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/organizationalUnits/_/actions/getOrganizationalUnitIdByExternalId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrganizationalUnitIdByExternalIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(instanceId *string, applicationId *string, userId *string) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserHeaders{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(instanceId, applicationId, userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(instanceId *string, applicationId *string, userId *string, headers *GetUserHeaders, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	userId = openapiutil.GetEncodeParam(userId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/users/" + tea.StringValue(userId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserIdByEmail(instanceId *string, applicationId *string, request *GetUserIdByEmailRequest) (_result *GetUserIdByEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserIdByEmailHeaders{}
	_result = &GetUserIdByEmailResponse{}
	_body, _err := client.GetUserIdByEmailWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserIdByEmailWithOptions(instanceId *string, applicationId *string, request *GetUserIdByEmailRequest, headers *GetUserIdByEmailHeaders, runtime *util.RuntimeOptions) (_result *GetUserIdByEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["email"] = request.Email
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserIdByEmail"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/users/_/actions/getUserIdByEmail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserIdByEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserIdByPhoneNumber(instanceId *string, applicationId *string, request *GetUserIdByPhoneNumberRequest) (_result *GetUserIdByPhoneNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserIdByPhoneNumberHeaders{}
	_result = &GetUserIdByPhoneNumberResponse{}
	_body, _err := client.GetUserIdByPhoneNumberWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserIdByPhoneNumberWithOptions(instanceId *string, applicationId *string, request *GetUserIdByPhoneNumberRequest, headers *GetUserIdByPhoneNumberHeaders, runtime *util.RuntimeOptions) (_result *GetUserIdByPhoneNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["phoneNumber"] = request.PhoneNumber
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserIdByPhoneNumber"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/users/_/actions/getUserIdByPhoneNumber"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserIdByPhoneNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserIdByUserExternalId(instanceId *string, applicationId *string, request *GetUserIdByUserExternalIdRequest) (_result *GetUserIdByUserExternalIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserIdByUserExternalIdHeaders{}
	_result = &GetUserIdByUserExternalIdResponse{}
	_body, _err := client.GetUserIdByUserExternalIdWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserIdByUserExternalIdWithOptions(instanceId *string, applicationId *string, request *GetUserIdByUserExternalIdRequest, headers *GetUserIdByUserExternalIdHeaders, runtime *util.RuntimeOptions) (_result *GetUserIdByUserExternalIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserExternalId)) {
		body["userExternalId"] = request.UserExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.UserSourceId)) {
		body["userSourceId"] = request.UserSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserSourceType)) {
		body["userSourceType"] = request.UserSourceType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserIdByUserExternalId"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/users/_/actions/getUserIdByExternalId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserIdByUserExternalIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserInfo(instanceId *string, applicationId *string) (_result *GetUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserInfoHeaders{}
	_result = &GetUserInfoResponse{}
	_body, _err := client.GetUserInfoWithOptions(instanceId, applicationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserInfoWithOptions(instanceId *string, applicationId *string, headers *GetUserInfoHeaders, runtime *util.RuntimeOptions) (_result *GetUserInfoResponse, _err error) {
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserInfo"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/oauth2/userinfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrganizationalUnitParentIds(instanceId *string, applicationId *string, organizationalUnitId *string) (_result *ListOrganizationalUnitParentIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListOrganizationalUnitParentIdsHeaders{}
	_result = &ListOrganizationalUnitParentIdsResponse{}
	_body, _err := client.ListOrganizationalUnitParentIdsWithOptions(instanceId, applicationId, organizationalUnitId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrganizationalUnitParentIdsWithOptions(instanceId *string, applicationId *string, organizationalUnitId *string, headers *ListOrganizationalUnitParentIdsHeaders, runtime *util.RuntimeOptions) (_result *ListOrganizationalUnitParentIdsResponse, _err error) {
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	organizationalUnitId = openapiutil.GetEncodeParam(organizationalUnitId)
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("ListOrganizationalUnitParentIds"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/organizationalUnits/" + tea.StringValue(organizationalUnitId) + "/parentIds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOrganizationalUnitParentIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrganizationalUnits(instanceId *string, applicationId *string, request *ListOrganizationalUnitsRequest) (_result *ListOrganizationalUnitsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListOrganizationalUnitsHeaders{}
	_result = &ListOrganizationalUnitsResponse{}
	_body, _err := client.ListOrganizationalUnitsWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrganizationalUnitsWithOptions(instanceId *string, applicationId *string, request *ListOrganizationalUnitsRequest, headers *ListOrganizationalUnitsHeaders, runtime *util.RuntimeOptions) (_result *ListOrganizationalUnitsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["parentId"] = request.ParentId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOrganizationalUnits"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/organizationalUnits"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOrganizationalUnitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsers(instanceId *string, applicationId *string, request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListUsersHeaders{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersWithOptions(instanceId *string, applicationId *string, request *ListUsersRequest, headers *ListUsersHeaders, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitId)) {
		query["organizationalUnitId"] = request.OrganizationalUnitId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PatchOrganizationalUnit(instanceId *string, applicationId *string, organizationalUnitId *string, request *PatchOrganizationalUnitRequest) (_result *PatchOrganizationalUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PatchOrganizationalUnitHeaders{}
	_result = &PatchOrganizationalUnitResponse{}
	_body, _err := client.PatchOrganizationalUnitWithOptions(instanceId, applicationId, organizationalUnitId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PatchOrganizationalUnitWithOptions(instanceId *string, applicationId *string, organizationalUnitId *string, request *PatchOrganizationalUnitRequest, headers *PatchOrganizationalUnitHeaders, runtime *util.RuntimeOptions) (_result *PatchOrganizationalUnitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	organizationalUnitId = openapiutil.GetEncodeParam(organizationalUnitId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitName)) {
		body["organizationalUnitName"] = request.OrganizationalUnitName
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PatchOrganizationalUnit"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/organizationalUnits/" + tea.StringValue(organizationalUnitId)),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &PatchOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PatchUser(instanceId *string, applicationId *string, userId *string, request *PatchUserRequest) (_result *PatchUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PatchUserHeaders{}
	_result = &PatchUserResponse{}
	_body, _err := client.PatchUserWithOptions(instanceId, applicationId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PatchUserWithOptions(instanceId *string, applicationId *string, userId *string, request *PatchUserRequest, headers *PatchUserHeaders, runtime *util.RuntimeOptions) (_result *PatchUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.EmailVerified)) {
		body["emailVerified"] = request.EmailVerified
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["phoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumberVerified)) {
		body["phoneNumberVerified"] = request.PhoneNumberVerified
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneRegion)) {
		body["phoneRegion"] = request.PhoneRegion
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["username"] = request.Username
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PatchUser"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/users/" + tea.StringValue(userId)),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &PatchUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeToken(instanceId *string, applicationId *string, request *RevokeTokenRequest) (_result *RevokeTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RevokeTokenResponse{}
	_body, _err := client.RevokeTokenWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeTokenWithOptions(instanceId *string, applicationId *string, request *RevokeTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RevokeTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	applicationId = openapiutil.GetEncodeParam(applicationId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["client_id"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSecret)) {
		query["client_secret"] = request.ClientSecret
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.TokenTypeHint)) {
		query["token_type_hint"] = request.TokenTypeHint
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeToken"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(instanceId) + "/" + tea.StringValue(applicationId) + "/oauth2/revoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
