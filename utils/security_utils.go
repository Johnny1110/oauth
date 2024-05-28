package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"oauth/model"
	"oauth/sys"
	"time"
)

var (
	RSAPrivateKey *rsa.PrivateKey
	RSAPublicKey  *rsa.PublicKey
)

func init() {
	privateKey, err := loadPrivateKey()
	if err != nil {
		sys.Logger().Fatalf("load private key err: %v", err)
	}
	publicKey, err := loadPublicKey()
	if err != nil {
		sys.Logger().Fatalf("load public key err: %v", err)
	}
	RSAPrivateKey = privateKey
	RSAPublicKey = publicKey
}

// 加載 RSA 私鑰
func loadPrivateKey() (*rsa.PrivateKey, error) {
	// TODO 未來搬到 db 去
	privateKeyPemString := `
-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCcalnn5Qiffkkuc0z4LXhbBnP41tDZ9xXZCWD2bjtK49zl8Dq1jAnW1eeQJWy1XXs8ZicPPek6JF9p3XxuhYX9mSSuZCd2NXfaBI079TlLXK8A2KInhG9mfxkXtGrBHC5LaZSGJf7S1ROoWNTrC949YUIuQ9CqEXlJ6rSKFYIgQD9IXnn4DVArLiZWkRCgG9bhQSdLSh4lxf2g916CMLcG8umVP68fg6cD3k7NZikzfS9p93tYgPnd9m70poC8Ip08hZxKyEXXltqiePUdTAntJNhfv6VOQQuSYF66bYfquY6HzHQ3VyDqK5xMbSI12qMgazAaTbzigwcUvOFbA7rLAgMBAAECggEAZnIRqJyN5UilTdMe2n2NkrDLTHQGHnyCxsApzb2tAAP2LnQFP9sUpq07GWIprTApAwo1EvqNwxMHmyMB2LGDbPHc4IugfP/QG+9XQan/eKifxoIc0p0fCZa9LJVyRkDa4XGpYaOJHzWHxn9IFRqU2MbWvc6U4I7JTex3iulDbTMhF4R0M76Z3jtlhGlWsq/2kintBbje+ZXAFvPTGdo3DV/vih2uli9/YqUsDMphdW/m8tCVYmW9hpKiQYcp0/Dv8kLzB2/jeZ+BRoTeS09JLpO2sFEiJATyXmpq2Qhk8g0y6zy5drzjxkun5x5wBgSOulNT4EZddaTkmcIWrVOJoQKBgQDLL+elScwNJBmA3CZdXGcbbfE4qK29kXcB2kG//0FFBwWVhFUXg6pgOvWINVfar3gguUFnIKklA3Wb859Q9kW3gMyA/w2104AO4G71o2fyNydabeYsxh1WKHWpobQWxyXqJne6pr3ukTb3tp8UPKfb1OFdtjMEduGjQfBO3EfhhQKBgQDFEkCbhS6MjBhGJnlUUSUcY5aNFvIixv+arXheR+IzJCuosgJrmFRXQbCBx/3/4xQJZuEdPehjWCDq7kTJdZgrKMzBoVrVnn3uuA3t6r6u7HjkYflsYdT+OfYaZJdjMUHho94n4PSlYvRqQ6eCGjAaJ26Kie1Dnu5QAc2w3gy0DwKBgQCVrNtSWhNzVrVmxEWKnqfhf9KjLzaVH5PwDGxE1+6nv61wX8QjBz25l5UJWmo2UO4IBQ/VvSx8dJjtYcBpbpEaxUlgeQILgBqkWtXCIzZOKizWI4DcWLCBqFpMtC6qXNdkLiQinfPpypUYUzHKQYhRmvbNBot8bWp3zfoMzZ4x1QKBgQC1L2jTE4mOkqcmp+zZBpnWFgGuyi/opYkPTvnhxLlFR7YULUVoYu74Il8DkzoF72LWmg3Scr6bx8TL+jCoAEPdOm+2foEi8cralcHIwhB6htNHNoS5juDis6t+7Ij7G6h0qdJwW7TR8b7BjF4PkcAz65kIKnNHvnFggaf5Os33JQKBgQDGPTQxfg0wIpbTzjCIXlQ7jH5u8M3J5vzFS48egZl3gSuV37QXQ2XsM0SOi1bviZmUp5XSfuH3wjk9tLuVRjYkYIzIotc8wi46UwyJ7A01psqvIMgLF1tIvEfOiqZCvkqehDnp98Wsv/d6VPOmvkt4HJzGxQ9vPbMWtQmKDj3jdQ==
-----END PRIVATE KEY-----`
	block, _ := pem.Decode([]byte(privateKeyPemString))
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA private key")
	}

	return rsaPrivateKey, nil
}

// 加載 RSA 公鑰
func loadPublicKey() (*rsa.PublicKey, error) {
	// TODO 未來搬到 db 去
	publicKeyPemString := `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnGpZ5+UIn35JLnNM+C14WwZz+NbQ2fcV2Qlg9m47SuPc5fA6tYwJ1tXnkCVstV17PGYnDz3pOiRfad18boWF/ZkkrmQndjV32gSNO/U5S1yvANiiJ4RvZn8ZF7RqwRwuS2mUhiX+0tUTqFjU6wvePWFCLkPQqhF5Seq0ihWCIEA/SF55+A1QKy4mVpEQoBvW4UEnS0oeJcX9oPdegjC3BvLplT+vH4OnA95OzWYpM30vafd7WID53fZu9KaAvCKdPIWcSshF15baonj1HUwJ7STYX7+lTkELkmBeum2H6rmOh8x0N1cg6iucTG0iNdqjIGswGk284oMHFLzhWwO6ywIDAQAB
-----END PUBLIC KEY-----`
	block, _ := pem.Decode([]byte(publicKeyPemString))

	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return rsaPub, nil
}

func ValidateToken(tokenString string) (authCode string, roles []string, scopes []string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return RSAPublicKey, nil
	})
	if err != nil {
		sys.Logger().Warningf("validate token fail: %s", err.Error())
		return "", nil, nil, errors.New("token is invalid")
	}

	// Extract claims
	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claims.AuthCode, claims.Roles, claims.Scopes, nil
	}
	return "", nil, nil, errors.New("invalid token.")
}

// Function to check if the user has the required permissions
func HasRequiredPermissions(userRoles, requiredRoles, userScopes, requiredScopes []string) bool {
	roleMatch := false
	scopeMatch := true

	sys.Logger().Debugf("Checking required permissions for userRoles %s requiredRoles: %v", userRoles, requiredRoles)

	for _, requiredRole := range requiredRoles {
		for _, userRole := range userRoles {
			if requiredRole == userRole {
				roleMatch = true
				break
			}
		}
	}

	sys.Logger().Debugf("Checking required permissions for userScopes %s requiredScopes: %v", userScopes, requiredScopes)

	for _, requiredScope := range requiredScopes {
		tempScopeCheck := false
		for _, userScope := range userScopes {
			if requiredScope == userScope {
				tempScopeCheck = true
				break
			}
		}
		if !tempScopeCheck {
			scopeMatch = false
			break
		}
	}

	return roleMatch && scopeMatch
}

// secret key used for signing JWT tokens
var jwtSecretKey = []byte("your-very-secret-key")

func GenerateJWT(authCode string, email string, username string, userRoles []string, userScopes []string, expiresIn int) (string, error) {
	// Generate JWT token with roles and scopes

	jwtID := GenRandomString(32, true)
	// expire time
	expirationTime := time.Now().Add(time.Duration(expiresIn) * time.Minute)
	claims := &model.Claims{
		AuthCode: authCode,
		Email:    email,
		Roles:    userRoles,
		Scopes:   userScopes,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "frizo-oauth",
			Subject:   username,
			ID:        jwtID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(RSAPrivateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
