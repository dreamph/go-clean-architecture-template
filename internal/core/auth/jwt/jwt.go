package jwt

import (
	"backend/internal/core/utils"
	"errors"
	"strings"
	"time"

	jwttoken "github.com/golang-jwt/jwt/v5"
	errs "github.com/pkg/errors"
)

var (
	errTokenRequired         = errs.New("Token Required")
	errTokenInvalidOrExpired = errs.New("Token Invalid or Expired")
)

type ConfigInfo struct {
	Issuer            string
	SecretKey         []byte
	Expiration        time.Duration
	RefreshExpiration time.Duration
	HandleEncryptInfo func(data string) string
	HandleDecryptInfo func(data string) string
}

type TokenInfo struct {
	AccessToken           string        `json:"accessToken"`
	RefreshToken          string        `json:"refreshToken"`
	Expiration            time.Duration `json:"expiration"  swaggertype:"number"`
	RefreshExpiration     time.Duration `json:"refreshExpiration"  swaggertype:"number"`
	AccessTokenExpiresAt  *time.Time    `json:"-"`
	RefreshTokenExpiresAt *time.Time    `json:"-"`
}

type TokenData struct {
	jwttoken.RegisteredClaims
	ID        string `json:"id,omitempty"`
	Scope     string `json:"scope,omitempty"`
	Org       string `json:"org,omitempty"`
	OrgIssuer string `json:"orgIssuer,omitempty"`
	Info      string `json:"info,omitempty"`
	UserType  string `json:"userType,omitempty"`
}

type JwtToken interface {
	Create(tokenData *TokenData) (*TokenInfo, error)
	GetTokenData(token string) (*TokenData, error)
	GetInfoData(info string, infoObjPtr interface{}) error
	GetTokenDataWithoutValidate(token string) (*TokenData, error)
	Validate(token string) error
	CreateJwtToken(tokenData *TokenData, issuer string, key []byte, expiresAt time.Time) (string, error)
	CreateToken(tokenData *TokenData, timeoutDuration time.Duration) (string, *time.Time, error)
}

type jwtToken struct {
	config *ConfigInfo
}

func NewJwtToken(config *ConfigInfo) JwtToken {
	return &jwtToken{config: config}
}

func (j *jwtToken) Create(tokenData *TokenData) (*TokenInfo, error) {
	accessToken, accessTokenExpiresAt, err := j.CreateToken(tokenData, j.config.Expiration)
	if err != nil {
		return nil, err
	}
	refreshToken, refreshTokenExpiresAt, err := j.CreateToken(tokenData, j.config.RefreshExpiration)
	if err != nil {
		return nil, err
	}
	return &TokenInfo{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		Expiration:            j.config.Expiration,
		RefreshExpiration:     j.config.RefreshExpiration,
		AccessTokenExpiresAt:  accessTokenExpiresAt,
		RefreshTokenExpiresAt: refreshTokenExpiresAt,
	}, nil
}

func (j *jwtToken) GetTokenData(token string) (*TokenData, error) {
	token = ExtractToken(token)
	tkn, err := j.validate(token)
	if err != nil {
		return nil, err
	}
	if claimsData, ok := tkn.Claims.(*TokenData); ok && tkn.Valid {
		return claimsData, nil
	} else {
		return nil, err
	}
}

func (j *jwtToken) GetInfoData(info string, infoObjPtr interface{}) error {
	if info == "" {
		return nil
	}
	if j.config.HandleEncryptInfo == nil {
		return errors.New("handleEncryptInfo is null")
	}
	decryptedInfo := j.config.HandleDecryptInfo(info)
	err := utils.JsonToObj([]byte(decryptedInfo), infoObjPtr)
	if err != nil {
		return err
	}
	return nil
}

func (j *jwtToken) CreateToken(tokenData *TokenData, timeoutDuration time.Duration) (string, *time.Time, error) {
	if timeoutDuration <= 0 {
		timeoutDuration = j.config.Expiration
	}

	expiredAt := time.Now().Add(timeoutDuration)
	accessToken, err := j.CreateJwtToken(tokenData, j.config.Issuer, j.config.SecretKey, expiredAt)
	if err != nil {
		return "", nil, err
	}
	return accessToken, &expiredAt, nil
}

func (j *jwtToken) validate(token string) (*jwttoken.Token, error) {
	if token == "" {
		return nil, errTokenRequired
	}
	claims := &TokenData{}
	tkn, err := jwttoken.ParseWithClaims(token, claims, func(token *jwttoken.Token) (interface{}, error) {
		return j.config.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, errTokenInvalidOrExpired
	}
	return tkn, nil
}

func (j *jwtToken) Validate(token string) error {
	if token == "" {
		return errTokenRequired
	}
	_, err := j.validate(token)
	if err != nil {
		return err
	}
	return nil
}

func (j *jwtToken) CreateJwtToken(tokenData *TokenData, issuer string, key []byte, expiresAt time.Time) (string, error) {
	now := time.Now()

	id := ""
	if utils.IsNotEmpty(tokenData.ID) {
		id = tokenData.ID
	} else {
		id = utils.NewID()
	}
	tokenData.RegisteredClaims = jwttoken.RegisteredClaims{
		ID:        id,
		Issuer:    issuer,
		ExpiresAt: jwttoken.NewNumericDate(expiresAt),
		IssuedAt:  jwttoken.NewNumericDate(now),
		NotBefore: jwttoken.NewNumericDate(now),
	}

	/*claims := &TokenData{
		ID:        tokenData.ID,
		Scope:     tokenData.Scope,
		Info:      tokenData.Info,
		Org:       tokenData.Org,
		OrgIssuer: tokenData.OrgIssuer,
		UserType:  tokenData.UserType,
		RegisteredClaims: jwttoken.RegisteredClaims{
			ID:        utils.NewID(),
			Issuer:    issuer,
			ExpiresAt: jwttoken.NewNumericDate(expiresAt),
			IssuedAt:  jwttoken.NewNumericDate(now),
			NotBefore: jwttoken.NewNumericDate(now),
		},
	}*/

	return SignObjectClaims(tokenData, key)
}

func SignString(data string, keyName string, key []byte) (string, error) {
	claims := jwttoken.MapClaims{}
	claims[keyName] = data
	return SignObjectClaims(claims, key)
}

func SignObjectClaims(claims jwttoken.Claims, key []byte) (string, error) {
	token := jwttoken.NewWithClaims(jwttoken.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/*
func ValidateAndGet(token string, key []byte, claims jwttoken.Claims) error {
	if token == "" {
		return errTokenRequired
	}
	tkn, err := jwttoken.ParseWithClaims(token, claims, func(token *jwttoken.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return err
	}
	if !tkn.Valid {
		return errTokenInvalidOrExpired
	}
	return nil
}
*/

func ExtractToken(bearToken string) string {
	if strings.Contains(bearToken, "Bearer") {
		strArr := strings.Split(bearToken, " ")
		if len(strArr) == 2 {
			return strings.TrimSpace(strArr[1])
		}
	}
	return strings.TrimSpace(bearToken)
}

func (j *jwtToken) GetTokenDataWithoutValidate(token string) (*TokenData, error) {
	if token == "" {
		return nil, errTokenRequired
	}
	claims := &TokenData{}
	_, err := jwttoken.ParseWithClaims(token, claims, func(token *jwttoken.Token) (interface{}, error) {
		return j.config.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func ValidateToken(tokenString string) (bool, *jwttoken.RegisteredClaims, error) {
	if tokenString == "" {
		return false, nil, errTokenRequired
	}
	claims := &jwttoken.RegisteredClaims{}
	_, _, err := jwttoken.NewParser().ParseUnverified(tokenString, claims)
	if err != nil {
		return false, nil, err
	}

	token := jwttoken.NewValidator()
	err = token.Validate(claims)
	if err != nil {
		return false, nil, errTokenInvalidOrExpired
	}

	return true, claims, nil
}
