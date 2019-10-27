package jwtdecode

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"firebase.google.com/go/auth"
)

//FireBaseCustomToken auth.token add email
type FireBaseCustomToken struct {
	auth.Token
	Email string `json:"email"`
}

//DecomposeFB  jwt in firebase authorization decompose header and claim, signature
func (s *jwtDecode) DecomposeFB(jwt string) ([]string, error) {
	hCS := strings.Split(jwt, ".")
	if len(hCS) == 3 {
		return hCS, nil
	}
	return nil, errors.New("Error jwt str decompose: jwt structure don't match")

}

//DecodeClaim
func (s *jwtDecode) DecodeClaimFB(payload string) (*FireBaseCustomToken, error) {
	payloadByte, err := base64.StdEncoding.DecodeString(payload)

	if err != nil {
		return nil, errors.New("Error jwt token decode: " + err.Error())
	}

	var tokenJSON FireBaseCustomToken
	err = json.Unmarshal(payloadByte, &tokenJSON)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, errors.New("Error jwt token unmarshal: " + err.Error())
	}
	return &tokenJSON, nil
}