package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

// Test the loginHandler function.
func TestLoginHandler(t *testing.T) {
	var jsonStr = []byte(`{"username":"user","password":"pass"}`)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(loginHandler)

	handler.ServeHTTP(rr, req)

	assert.Truef(t, rr.Code == http.StatusOK, "handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	assert.Truef(t, rr.Body.String() != "", "handler returned unexpected body: got empty want not empty")
}

// TestInvalidToken checks what happens when an invalid token is provided
func TestInvalidToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/files", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer InvalidToken")

	rr := httptest.NewRecorder()
	handler := authMiddleware(http.HandlerFunc(fileHandler))

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code, "Expected unauthorized status code")
}

// TestExpiredToken checks what happens when an expired token is provided
func TestExpiredToken(t *testing.T) {
	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(-5 * time.Minute).Unix(),
	})
	expiredTokenString, _ := expiredToken.SignedString(jwtKey)

	req, err := http.NewRequest("GET", "/files", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+expiredTokenString)

	rr := httptest.NewRecorder()
	handler := authMiddleware(http.HandlerFunc(fileHandler))

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code, "Expected unauthorized status code")
}

// TestIncorrectLogin checks what happens when the login credentials are incorrect
func TestIncorrectLogin(t *testing.T) {
	var jsonStr = []byte(`{"username":"wronguser","password":"wrongpass"}`)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(loginHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code, "Expected unauthorized status code")
}

// Test the auth middleawrae on the fileHandler function.
func TestFileHandlerAuth(t *testing.T) {
	req, err := http.NewRequest("GET", "/files", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := authMiddleware(http.HandlerFunc(fileHandler))

	handler.ServeHTTP(rr, req)

	assert.Equalf(t, http.StatusUnauthorized, rr.Code, "handler returned wrong status code: got %v want %v", rr.Code, http.StatusUnauthorized)
}
