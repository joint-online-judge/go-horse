// Package schemas provides primitives to interact with the openapi HTTP API.
package schemas

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/pkg/configs"
)

type JWTCommon struct {
	Type      string `json:"type"`
	OauthName string `json:"oauthName"`
	jwt.RegisteredClaims
}

type JWTClaims struct {
	JWTCommon
	Fresh    bool   `json:"fresh"`
	Csrf     string `json:"csrf"`
	Category string `json:"category"`
	Username string `json:"username"`
	Gravatar string `json:"gravatar"`
	Role     string `json:"role"`
	IsActive bool   `json:"isActive"`
}

func JWT(c *fiber.Ctx) *JWTClaims {
	return c.Locals("jwt").(*JWTClaims)
}

func JWTUser(c *fiber.Ctx) *User {
	claims := c.Locals("jwt").(*JWTClaims)
	user := User{
		Gravatar: &claims.Gravatar,
		Id:       uuid.MustParse(claims.ID),
		IsActive: &claims.IsActive,
		Role:     &claims.Role,
		Username: claims.Username,
	}
	return &user
}

func NewAccessToken(
	user User,
	oauth_name, category string,
	fresh bool,
) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		JWTCommon: JWTCommon{
			Type:      "access",
			OauthName: oauth_name,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(
					time.Now().
						Add(time.Duration(configs.Conf.JwtExpireSeconds) * time.Second),
				),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				Issuer:    uuid.New().String(),
				ID:        user.Id.String(),
				Subject:   user.Id.String(),
				// Audience:  []string{"somebody_else"},
			},
		},
		Fresh:    fresh,
		Csrf:     "test", // FIXME: do we need it as we can use csrf middleware?
		Category: category,
		Username: user.Username,
		Gravatar: *user.Gravatar,
		Role:     *user.Role,
		IsActive: *user.IsActive,
	})
	return token.SignedString([]byte(configs.Conf.JwtSecret))
}

func NewRefreshToken(user User, oauth_name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTCommon{
		Type:      "refresh",
		OauthName: oauth_name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().
					Add(time.Duration(configs.Conf.JwtExpireSeconds) * time.Second),
			),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    uuid.New().String(),
			ID:        user.Id.String(),
			Subject:   user.Id.String(),
			// Audience:  []string{"somebody_else"},
		},
	})
	return token.SignedString([]byte(configs.Conf.JwtSecret))
}

// AuthTokens defines model for AuthTokens.
type AuthTokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType"`
}

// AuthTokensWithLakefs defines model for AuthTokensWithLakefs.
type AuthTokensWithLakefs struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessToken     string `json:"accessToken"`
	RefreshToken    string `json:"refreshToken"`
	SecretAccessKey string `json:"secretAccessKey"`
	TokenType       string `json:"tokenType"`
}

// Defines values for JWTAccessTokenCategory.
const (
	JWTAccessTokenCategoryOauth JWTAccessTokenCategory = "oauth"
	JWTAccessTokenCategoryUser  JWTAccessTokenCategory = "user"
)

// JWTAccessToken defines model for JWTAccessToken.
type JWTAccessToken struct {
	Category  JWTAccessTokenCategory `json:"category"`
	Csrf      *string                `json:"csrf,omitempty"`
	Exp       int                    `json:"exp"`
	Fresh     *bool                  `json:"fresh,omitempty"`
	Gravatar  *string                `json:"gravatar,omitempty"`
	Iat       int                    `json:"iat"`
	IsActive  bool                   `json:"isActive"`
	Jti       string                 `json:"jti"`
	Nbf       int                    `json:"nbf"`
	OauthName *string                `json:"oauthName,omitempty"`
	Role      *string                `json:"role,omitempty"`
	Sub       string                 `json:"sub"`
	Type      string                 `json:"type"`
	Username  string                 `json:"username"`
}

// JWTAccessTokenCategory defines model for JWTAccessToken.Category.
type JWTAccessTokenCategory string

// OAuth2Client defines model for OAuth2Client.
type OAuth2Client struct {
	DisplayName string `json:"displayName"`
	Icon        string `json:"icon"`
	OauthName   string `json:"oauthName"`
}

// OAuth2ClientList defines model for OAuth2ClientList.
type OAuth2ClientList struct {
	Count   *int            `json:"count,omitempty"`
	Results *[]OAuth2Client `json:"results,omitempty"`
}

// OAuth2PasswordRequestForm defines model for OAuth2PasswordRequestForm.
type OAuth2PasswordRequestForm struct {
	ClientId     *string `json:"clientId,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty"`
	GrantType    *string `json:"grantType,omitempty"`
	Password     *string `json:"password,omitempty"`
	Scope        *string `json:"scope,omitempty"`
	Username     *string `json:"username,omitempty"`
}

// LoginParams defines parameters for Login.
type LoginParams struct {
	// Cookie Add Set/Delete-Cookie on response header
	Cookie       *bool                   `form:"cookie,omitempty" json:"cookie,omitempty"`
	ResponseType LoginParamsResponseType `form:"responseType"     json:"responseType"`

	// RedirectUrl The redirect url after the operation
	RedirectUrl *string `form:"redirectUrl,omitempty" json:"redirectUrl,omitempty"`
}

// LogoutParams defines parameters for Logout.
type LogoutParams struct {
	// Cookie Add Set/Delete-Cookie on response header
	Cookie       *bool                    `form:"cookie,omitempty" json:"cookie,omitempty"`
	ResponseType LogoutParamsResponseType `form:"responseType"     json:"responseType"`

	// RedirectUrl The redirect url after the operation
	RedirectUrl *string `form:"redirectUrl,omitempty" json:"redirectUrl,omitempty"`
}

// OauthAuthorizeParams defines parameters for OauthAuthorize.
type OauthAuthorizeParams struct {
	Scopes *[]string `form:"scopes,omitempty" json:"scopes,omitempty"`

	// Cookie Add Set/Delete-Cookie on response header
	Cookie       *bool                            `form:"cookie,omitempty" json:"cookie,omitempty"`
	ResponseType OauthAuthorizeParamsResponseType `form:"responseType"     json:"responseType"`

	// RedirectUrl The redirect url after the operation
	RedirectUrl *string `form:"redirectUrl,omitempty" json:"redirectUrl,omitempty"`
}

// RefreshParams defines parameters for Refresh.
type RefreshParams struct {
	// Cookie Add Set/Delete-Cookie on response header
	Cookie       *bool                     `form:"cookie,omitempty" json:"cookie,omitempty"`
	ResponseType RefreshParamsResponseType `form:"responseType"     json:"responseType"`

	// RedirectUrl The redirect url after the operation
	RedirectUrl *string `form:"redirectUrl,omitempty" json:"redirectUrl,omitempty"`
}

// RegisterParams defines parameters for Register.
type RegisterParams struct {
	// Cookie Add Set/Delete-Cookie on response header
	Cookie       *bool                      `form:"cookie,omitempty" json:"cookie,omitempty"`
	ResponseType RegisterParamsResponseType `form:"responseType"     json:"responseType"`

	// RedirectUrl The redirect url after the operation
	RedirectUrl *string `form:"redirectUrl,omitempty" json:"redirectUrl,omitempty"`
}

// GetTokenParams defines parameters for GetToken.
type GetTokenParams struct {
	// Cookie Add Set/Delete-Cookie on response header
	Cookie       *bool                      `form:"cookie,omitempty" json:"cookie,omitempty"`
	ResponseType GetTokenParamsResponseType `form:"responseType"     json:"responseType"`

	// RedirectUrl The redirect url after the operation
	RedirectUrl *string `form:"redirectUrl,omitempty" json:"redirectUrl,omitempty"`
}

// Redirect defines model for Redirect.
type Redirect struct {
	RedirectUrl string `json:"redirectUrl"`
}