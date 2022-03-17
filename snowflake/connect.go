package snowflake

import (
	"context"
	"crypto/rsa"
	"database/sql"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	gosnowflake "github.com/snowflakedb/gosnowflake"
	"github.com/youmark/pkcs8"
	"golang.org/x/crypto/ssh"

	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*sql.DB, error) {

	// have we already created and cached the session?
	cacheKey := "snowflake"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*sql.DB), nil
	}

	config := GetConfig(d.Connection)
	var account, user, password, oauthAccessToken, privateKeyPath, privateKey, privateKeyPassphrase, region, role, oauthEndpoint, oauthClientSecret, oauthClientID, oauthRefreshToken, oauthRedirectURL string
	var browserAuth bool
	if config.Account != nil {
		account = *config.Account
	}
	if config.User != nil {
		user = *config.User
	}
	if config.Password != nil {
		password = *config.Password
	}
	if config.BrowserAuth != nil {
		browserAuth = *config.BrowserAuth
	}
	if config.PrivateKeyPath != nil {
		privateKeyPath = *config.PrivateKeyPath
	}
	if config.PrivateKey != nil {
		privateKey = *config.PrivateKey
	}
	if config.PrivateKeyPassphrase != nil {
		privateKeyPassphrase = *config.PrivateKeyPassphrase
	}
	if config.PrivateKeyPassphrase != nil {
		privateKeyPassphrase = *config.PrivateKeyPassphrase
	}
	if config.Region != nil {
		region = *config.Region
	}
	if config.Role != nil {
		role = *config.Role
	}
	if config.OAuthEndpoint != nil {
		oauthEndpoint = *config.OAuthEndpoint
	}
	if config.OAuthClientSecret != nil {
		oauthClientSecret = *config.OAuthClientSecret
	}
	if config.OAuthClientID != nil {
		oauthClientID = *config.OAuthClientID
	}
	if config.OAuthRefreshToken != nil {
		oauthRefreshToken = *config.OAuthRefreshToken
	}
	if config.OAuthRedirectURL != nil {
		oauthRedirectURL = *config.OAuthRedirectURL
	}
	if config.OAuthAccessToken != nil {
		oauthAccessToken = *config.OAuthAccessToken
	}

	if config.OAuthRefreshToken != nil {
		accessToken, err := GetOauthAccessToken(oauthEndpoint, oauthClientID, oauthClientSecret, GetOauthData(oauthRefreshToken, oauthRedirectURL))
		if err != nil {
			return nil, fmt.Errorf("could not retreive access token from refresh token: %w", err)
		}
		oauthAccessToken = accessToken
	}

	dsn, err := DSN(ctx,
		account,
		user,
		password,
		browserAuth,
		privateKeyPath,
		privateKey,
		privateKeyPassphrase,
		oauthAccessToken,
		region,
		role,
	)
	if err != nil {
		plugin.Logger(ctx).Error("DSN", "could not build dsn for snowflake connection", err)
		return nil, fmt.Errorf("could not build dsn for snowflake connection: %w", err)
	}

	db, err := sql.Open("snowflake", dsn)
	if err != nil {
		return nil, fmt.Errorf("Could not open snowflake database: %w", err)
	}
	d.ConnectionManager.Cache.Set(cacheKey, db)
	return db, nil
}

func DSN(ctx context.Context, account, user,
	password string,
	browserAuth bool,
	privateKeyPath,
	privateKey,
	privateKeyPassphrase,
	oauthAccessToken,
	region,
	role string) (string, error) {

	// us-west-2 is their default region for snowflake instance, if it is mentioned in the connection config
	// don't add it into connection string
	if region == "us-west-2" {
		region = ""
	}

	config := gosnowflake.Config{
		Account: account,
		User:    user,
		Region:  region,
		Role:    role,
	}

	if privateKeyPath != "" {
		privateKeyBytes, err := ReadPrivateKeyFile(privateKeyPath)
		if err != nil {
			return "", fmt.Errorf("Private Key file could not be read: %w", err)
		}
		rsaPrivateKey, err := ParsePrivateKey(privateKeyBytes, []byte(privateKeyPassphrase))
		if err != nil {
			return "", fmt.Errorf("Private Key could not be parsed: %w", err)
		}
		config.PrivateKey = rsaPrivateKey
		config.Authenticator = gosnowflake.AuthTypeJwt
	} else if privateKey != "" {
		rsaPrivateKey, err := ParsePrivateKey([]byte(privateKey), []byte(privateKeyPassphrase))
		if err != nil {
			return "", fmt.Errorf("Private Key could not be parsed: %w", err)
		}
		config.PrivateKey = rsaPrivateKey
		config.Authenticator = gosnowflake.AuthTypeJwt
	} else if browserAuth {
		config.Authenticator = gosnowflake.AuthTypeExternalBrowser
	} else if oauthAccessToken != "" {
		config.Authenticator = gosnowflake.AuthTypeOAuth
		config.Token = oauthAccessToken
	} else if password != "" {
		config.Password = password
	} else {
		return "", errors.New("no authentication method provided")
	}

	return gosnowflake.DSN(&config)
}

func ReadPrivateKeyFile(privateKeyPath string) ([]byte, error) {
	expandedPrivateKeyPath, err := homedir.Expand(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("Invalid Path to private key: %w", err)
	}

	privateKeyBytes, err := ioutil.ReadFile(expandedPrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("Could not read private key: %w", err)
	}

	if len(privateKeyBytes) == 0 {
		return nil, errors.New("Private key is empty")
	}

	return privateKeyBytes, nil
}

func ParsePrivateKey(privateKeyBytes []byte, passhrase []byte) (*rsa.PrivateKey, error) {
	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyBlock == nil {
		return nil, fmt.Errorf("Could not parse private key, key is not in PEM format")
	}

	if privateKeyBlock.Type == "ENCRYPTED PRIVATE KEY" {
		if len(passhrase) == 0 {
			return nil, fmt.Errorf("Private key requires a passphrase, but private_key_passphrase was not supplied")
		}
		privateKey, err := pkcs8.ParsePKCS8PrivateKeyRSA(privateKeyBlock.Bytes, passhrase)
		if err != nil {
			return nil, fmt.Errorf(
				"Could not parse encrypted private key with passphrase, only ciphers aes-128-cbc, aes-128-gcm, aes-192-cbc, aes-192-gcm, aes-256-cbc, aes-256-gcm, and des-ede3-cbc are supported: %w", err)
		}
		return privateKey, nil
	}

	privateKey, err := ssh.ParseRawPrivateKey(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("Could not parse private key: %w", err)
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("privateKey not of type RSA")
	}
	return rsaPrivateKey, nil
}

type Result struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetOauthData(refreshToken, redirectUrl string) url.Values {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	data.Set("redirect_uri", redirectUrl)
	return data
}

func GetOauthRequest(dataContent io.Reader, endPoint, clientId, clientSecret string) (*http.Request, error) {
	request, err := http.NewRequest("POST", endPoint, dataContent)
	if err != nil {
		return nil, fmt.Errorf("Request to the endpoint could not be completed: %w", err)
	}
	request.SetBasicAuth(clientId, clientSecret)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	return request, nil

}

func GetOauthAccessToken(
	endPoint,
	client_id,
	client_secret string,
	data url.Values) (string, error) {

	client := &http.Client{}
	request, err := GetOauthRequest(strings.NewReader(data.Encode()), endPoint, client_id, client_secret)
	if err != nil {
		return "", fmt.Errorf("Oauth request returned error: %w", err)
	}

	var result Result

	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("Response status returned error: %w", err)
	}
	if response.StatusCode != 200 {
		return "", fmt.Errorf("Response status code: %s: %s", strconv.Itoa(response.StatusCode), http.StatusText(response.StatusCode))
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Response body was not able to be parsed: %w", err)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("Error parsing JSON from Snowflake: %w", err)
	}
	return result.AccessToken, nil
}
