package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"oauth/model"
	"oauth/sys"
	"time"
)

func ValidateToken(tokenString string) (roles []string, scopes []string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecretKey, nil
	})
	if err != nil {
		return nil, nil, err
	}

	// Extract claims
	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claims.Roles, claims.Scopes, nil
	}
	return nil, nil, errors.New("invalid token.")
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
