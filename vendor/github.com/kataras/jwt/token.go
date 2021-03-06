package jwt

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
)

var (
	// ErrMissing indicates that a given token to `Verify` is empty.
	ErrMissing = errors.New("token is empty")
	// ErrTokenForm indicates that the extracted token has not the expected form .
	ErrTokenForm = errors.New("invalid token form")
	// ErrTokenAlg indicates that the given algorithm does not match the extracted one.
	ErrTokenAlg = errors.New("unexpected token algorithm")
)

type (
	// PrivateKey is a generic type, this key is responsible for signing the token.
	PrivateKey interface{}
	// PublicKey is a generic type, this key is responsible to verify the token.
	PublicKey interface{}
)

func encodeToken(alg Alg, key PrivateKey, payload []byte) ([]byte, error) {
	header := createHeader(alg.Name())
	payload = Base64Encode(payload)

	headerPayload := joinParts(header, payload)

	signature, err := createSignature(alg, key, headerPayload)
	if err != nil {
		return nil, fmt.Errorf("encodeToken: signature: %w", err)
	}

	// header.payload.signature
	token := joinParts(headerPayload, signature)

	return token, nil
}

// We could omit the "alg" because the token contains it
// BUT, for security reason the algorithm MUST explicitly match
// (even if we perform hash comparison later on).
//
// Decodes and verifies the given compact "token".
// It returns the header, payoad and signature parts (decoded).
func decodeToken(alg Alg, key PublicKey, token []byte) ([]byte, []byte, []byte, error) {
	parts := bytes.Split(token, sep)
	if len(parts) != 3 {
		return nil, nil, nil, ErrTokenForm
	}

	header := parts[0]
	payload := parts[1]
	signature := parts[2]

	headerDecoded, err := Base64Decode(header)
	if err != nil {
		return nil, nil, nil, err
	}
	// validate header equality.
	if !CompareHeader(alg.Name(), headerDecoded) {
		return nil, nil, nil, ErrTokenAlg
	}

	signatureDecoded, err := Base64Decode(signature)
	if err != nil {
		return nil, nil, nil, err
	}
	// validate signature.
	headerPayload := joinParts(header, payload)
	if err := alg.Verify(key, headerPayload, signatureDecoded); err != nil {
		return nil, nil, nil, err
	}

	payload, err = Base64Decode(payload)
	if err != nil {
		return nil, nil, nil, err
	}
	return headerDecoded, payload, signatureDecoded, nil
}

var (
	sep    = []byte(".")
	pad    = []byte("=")
	padStr = string(pad)
)

func joinParts(parts ...[]byte) []byte {
	return bytes.Join(parts, sep)
}

// A builtin list of fixed headers for builtin algorithms (to boost the performance a bit).
// key = alg, value = the base64encoded full header
// (when kid or any other extra headers are not required to be inside).
type fixedHeader struct {
	// the json raw byte value.
	raw []byte
	// the base64 encoded value of raw.
	encoded []byte
	// same as raw but reversed order, e.g. first type then alg.
	// Useful to validate external jwt tokens that are not using the standard form order.
	reversed []byte
}

var fixedHeaders = map[string]*fixedHeader{
	NONE.Name():  nil,
	HS256.Name(): nil,
	HS384.Name(): nil,
	HS512.Name(): nil,
	RS256.Name(): nil,
	RS384.Name(): nil,
	RS512.Name(): nil,
	ES256.Name(): nil,
	ES384.Name(): nil,
	ES512.Name(): nil,
	EdDSA.Name(): nil,
}

func init() {
	for k := range fixedHeaders {
		fixedHeaders[k] = &fixedHeader{
			raw:      createHeaderRaw(k),
			encoded:  createHeader(k),
			reversed: createHeaderReversed(k),
		}
	}
}

func createHeader(alg string) []byte {
	if header := fixedHeaders[alg]; header != nil {
		return header.encoded
	}

	return Base64Encode([]byte(`{"alg":"` + alg + `","typ":"JWT"}`))
}

func createHeaderRaw(alg string) []byte {
	if header := fixedHeaders[alg]; header != nil {
		return header.raw
	}

	return []byte(`{"alg":"` + alg + `","typ":"JWT"}`)
}

func createHeaderReversed(alg string) []byte {
	if header := fixedHeaders[alg]; header != nil {
		return header.reversed
	}

	return []byte(`{"typ":"JWT","alg":"` + alg + `"}`)
}

// Note that this check is fully hard coded for known
// algorithms and it is fully hard coded in terms of
// its serialized format.
func compareHeader(alg string, headerDecoded []byte) bool {
	if len(headerDecoded) < 25 /* 28 but allow custom short algs*/ {
		return false
	}

	// Fast check if the order is reversed.
	// The specification says otherwise but
	// some other programming languages' libraries
	// don't actually follow the correct order.
	if headerDecoded[2] == 't' {
		expectedHeader := createHeaderReversed(alg)
		return bytes.Equal(expectedHeader, headerDecoded)
	}

	expectedHeader := createHeaderRaw(alg)
	return bytes.Equal(expectedHeader, headerDecoded)
}

func createSignature(alg Alg, key PrivateKey, headerAndPayload []byte) ([]byte, error) {
	signature, err := alg.Sign(key, headerAndPayload)
	if err != nil {
		return nil, err
	}
	return Base64Encode(signature), nil
}

// Base64Encode encodes "src" to jwt base64 url format.
// We could use the base64.RawURLEncoding but the below is a bit faster.
func Base64Encode(src []byte) []byte {
	buf := make([]byte, base64.URLEncoding.EncodedLen(len(src)))
	base64.URLEncoding.Encode(buf, src)

	return bytes.TrimRight(buf, padStr) // JWT: no trailing '='.
}

// Base64Decode decodes "src" to jwt base64 url format.
// We could use the base64.RawURLEncoding but the below is a bit faster.
func Base64Decode(src []byte) ([]byte, error) {
	if n := len(src) % 4; n > 0 {
		// JWT: Because of no trailing '=' let's suffix it
		// with the correct number of those '=' before decoding.
		src = append(src, bytes.Repeat(pad, 4-n)...)
	}

	buf := make([]byte, base64.URLEncoding.DecodedLen(len(src)))
	n, err := base64.URLEncoding.Decode(buf, src)
	return buf[:n], err
}

// Decode decodes the token of compact form WITHOUT verification and validation.
//
// This function is only useful to read a token's claims
// when the source is trusted and no algorithm verification or direct signature and
// content validation is required.
//
// Use `Verify/VerifyEncrypted` functions instead.
func Decode(token []byte) (*UnverifiedToken, error) {
	parts := bytes.Split(token, sep)
	if len(parts) != 3 {
		return nil, ErrTokenForm
	}

	header := parts[0]
	payload := parts[1]
	signature := parts[2]

	headerDecoded, err := Base64Decode(header)
	if err != nil {
		return nil, err
	}

	signatureDecoded, err := Base64Decode(signature)
	if err != nil {
		return nil, err
	}

	payload, err = Base64Decode(payload)
	if err != nil {
		return nil, err
	}

	tok := &UnverifiedToken{
		Header:    headerDecoded,
		Payload:   payload,
		Signature: signatureDecoded,
	}
	return tok, nil
}

// UnverifiedToken contains the compact form token parts.
// Look its `Claims` method to decode to a custom structure.
type UnverifiedToken struct {
	Header    []byte
	Payload   []byte
	Signature []byte
}

// Claims decodes the `Payload` field to the "dest".
func (t *UnverifiedToken) Claims(dest interface{}) error {
	return Unmarshal(t.Payload, dest)
}
