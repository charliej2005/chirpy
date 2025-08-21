package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestHashAndCheckPassword(t *testing.T) {
	password := "supersecret"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword failed: %v", err)
	}
	if hash == "" {
		t.Fatal("HashPassword returned empty hash")
	}

	// Correct password
	if err := CheckPasswordHash(password, hash); err != nil {
		t.Errorf("CheckPasswordHash failed for correct password: %v", err)
	}

	// Incorrect password
	if err := CheckPasswordHash("wrongpassword", hash); err == nil {
		t.Error("CheckPasswordHash did not fail for incorrect password")
	}
}

func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()
	secret := "testsecret"
	expiresIn := time.Minute

	tokenString, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("MakeJWT failed: %v", err)
	}
	if tokenString == "" {
		t.Fatal("MakeJWT returned empty token")
	}

	parsedID, err := ValidateJWT(tokenString, secret)
	if err != nil {
		t.Fatalf("ValidateJWT failed: %v", err)
	}
	if parsedID != userID {
		t.Errorf("ValidateJWT returned wrong userID: got %v, want %v", parsedID, userID)
	}

	// Incorrect secret
	_, err = ValidateJWT(tokenString, "wrongsecret")
	if err == nil {
		t.Error("ValidateJWT did not fail with invalid secret")
	}

	// Expired token
	expiredToken, err := MakeJWT(userID, secret, -time.Minute)
	if err != nil {
		t.Fatalf("MakeJWT (expired) failed: %v", err)
	}
	_, err = ValidateJWT(expiredToken, secret)
	if err == nil {
		t.Error("ValidateJWT did not fail for expired token")
	}
}
