package jwthelp

import (
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	coreError "iceline-hosting.com/core/error"
)

type kvStore interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type JWTHelper struct {
	ecdsaPrivateKey *ecdsa.PrivateKey
	ecdsaPublicKey  *ecdsa.PublicKey

	kvs kvStore
}

type options func(*JWTHelper) error

const (
	invalidTokenDuration = 1 * time.Hour
	validTokenDuration   = 1 * time.Hour
)

func WithPrivateKey() func(*JWTHelper) error {
	return func(j *JWTHelper) error {

		data := os.Getenv("JWT_PRIVATE_KEY")
		if data == "" {
			return errors.New("JWT_PRIVATE_KEY not set")
		}
		return j.initJWT([]byte(strings.ReplaceAll(data, "\\n", "\n")))
	}
}

func WithPublicKey() func(*JWTHelper) error {
	return func(j *JWTHelper) error {

		// // remove this godotenv code while running in docker containers
		// if err := godotenv.Load("../scheduler-manager/.env"); err != nil {
		// 	log.Println(err)
		// }
		data := os.Getenv("JWT_PUBLIC_KEY")
		if data == "" {
			return errors.New("JWT_PUBLIC_KEY not set")
		}
		return j.initJWT([]byte(strings.ReplaceAll(data, "\\n", "\n")))
	}
}
func NewJWTHelper(kvs kvStore, opts ...options) (*JWTHelper, error) {
	jwth := &JWTHelper{
		kvs: kvs,
	}

	for _, opt := range opts {
		err := opt(jwth)
		if err != nil {
			return nil, err
		}
	}

	return jwth, nil
}

func (j *JWTHelper) initJWT(key []byte) error {
	var err error
	// Print the raw key data for debugging
	block, _ := pem.Decode(key)
	if block == nil {
		fmt.Println("PEM block is nil") // Print a message if the PEM block is nil
		return fmt.Errorf("cannot decode pem. block returned is nil")
	}

	switch block.Type {
	case "EC PRIVATE KEY":
		err = j.initWithPrivateKey(block)
	case "PUBLIC KEY":
		err = j.initWithPublicKey(block)
	default:
		err = fmt.Errorf("bad key")
	}

	return err
}

func (j *JWTHelper) initWithPrivateKey(p *pem.Block) error {
	x509Encoded := p.Bytes
	ecdsaPrivateKey, err := x509.ParseECPrivateKey(x509Encoded)
	if err != nil {
		log.Printf("Error parsing private key: %s", err)

		return err
	}

	j.ecdsaPrivateKey = ecdsaPrivateKey
	return nil
}

func (j *JWTHelper) initWithPublicKey(p *pem.Block) error {
	x509bytes := p.Bytes
	tmpPublicKey, err := x509.ParsePKIXPublicKey(x509bytes)
	if err != nil {
		log.Printf("Error parsing public key: %s", err)

		return err
	}

	ecdsaPublicKey, ok := tmpPublicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("could not convert parsed key into public key")
	}

	j.ecdsaPublicKey = ecdsaPublicKey
	return nil
}

// GenerateJWT returns you a jwt string with expTime and params inside
func (j *JWTHelper) GenerateJWT(ctx context.Context, userID string) (string, error) {
	if j.ecdsaPrivateKey == nil {
		return "", fmt.Errorf("ecdsaPrivateKey not found")
	}

	jwtClaims := map[string]interface{}{
		"user": userID,
		"exp":  time.Now().Add(1 * time.Hour).Unix(),
	}

	var claims jwt.MapClaims = jwtClaims

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	str, err := token.SignedString(j.ecdsaPrivateKey)
	if err != nil {
		return "", err
	}

	if err := j.kvs.Set(ctx, userID, str, validTokenDuration); err != nil {
		return "", err
	}

	return str, nil
}

func (j *JWTHelper) RegenerateJWT(ctx context.Context, userID string) (string, error) {
	token, err := j.GenerateJWT(ctx, userID)
	if err != nil {
		return "", err
	}

	if err := j.kvs.Set(ctx, userID, token, validTokenDuration); err != nil {
		return "", err
	}

	return token, nil
}

// VerifyJWT verifies the jwt token and returns params inside the jwt
// verification validates the signing method, expiration date and signed token
// ... (your existing code)

// VerifyJWT verifies the jwt token and returns params inside the jwt
// verification validates the signing method, expiration date, and signed token
func (j *JWTHelper) VerifyJWT(ctx context.Context, tokenStr string) (map[string]interface{}, error) {
	fmt.Println("Verifying Token:", tokenStr) // Print the token being verified

	if j.ecdsaPublicKey == nil {
		return nil, fmt.Errorf("ecdsaPublicKey not found")
	}

	fmt.Println("Public Key for Verification:", j.ecdsaPublicKey) // Print the public key for verification

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		alg := token.Method.Alg()
		fmt.Println("Signing Algorithm:", alg) // Log the signing algorithm

		_, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			return nil, fmt.Errorf("signing method mismatch %s", alg)
		}

		return j.ecdsaPublicKey, nil
	})
	if err != nil {
		log.Printf("could not parse token %s: %s", tokenStr, err)
		return nil, fmt.Errorf("could not parse token: %w", err)
	}

	fmt.Println("Decoded Token Payload:", token.Claims) // Print the decoded payload for inspection

	if !token.Valid {
		return nil, coreError.BadJWT{}
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("could not convert token claims into map claims")
	}

	fmt.Println("Decoded Token Payload:", mapClaims) // Print the decoded payload for inspection

	userID, ok := mapClaims["user"].(string)
	if !ok {
		return nil, fmt.Errorf("could not convert user id into string")
	}

	originalToken, err := j.kvs.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting original token: %w", err)
	}

	fmt.Println("Original Token from Storage:", originalToken) // Print the original token for debugging

	if originalToken != tokenStr {
		log.Println("originalToken---------->", originalToken)
		return nil, coreError.BadJWT{Msg: "invalid token provided"}
	}

	return mapClaims, nil
}

// ... (your existing code)

func (j *JWTHelper) DeleteJWT(ctx context.Context, userID string) error {
	return j.kvs.Delete(ctx, userID)
}
