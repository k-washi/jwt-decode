package main

import (
	"fmt"

	"github.com/k-washi/jwt-decode/jwtdecode"
)

func main() {
	jwt := "eyJhbGciOiJSUzI1NiIsImtpZCI6ImEwYjQwY2NjYmQ0OWQxNmVkMjg2MGRiNzIyNmQ3NDZiNmZhZmRmYzAiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoidGVzdHVzZXIiLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vZXgtZmlyZWJhc2UtYXV0aCIsImF1ZCI6ImV4LWZpcmViYXNlLWF1dGgiLCJhdXRoX3RpbWUiOjE1NzIxMTQzMjIsInVzZXJfaWQiOiJxWmhzRjJIZnVXWkVCZ2hGYTRubDJLaWR5cDIyIiwic3ViIjoicVpoc0YySGZ1V1pFQmdoRmE0bmwyS2lkeXAyMiIsImlhdCI6MTU3MjExNDMyMiwiZXhwIjoxNTcyMTE3OTIyLCJlbWFpbCI6InRlc3RAdGVzdC5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsidGVzdEB0ZXN0LmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6InBhc3N3b3JkIn19.gFfq8ut-n-qxlkHiMCT6zusVvkNK2iCg0_xvEi1vBXK-KCWPG1dlyT5UUnuVlj3vHgFj3uERLsQo6fh8ReofN3KjNRomD8uLJlfykFJ6AbczzpxiXLqK8Y57qXsziF-bz78NKQG01RYxxeVgLnq_euPLkPkmOwEB5QayGjD3uTaHb0BphznsXdo4_7xjjW4GTfKxGt4EFq0WE88eEptBV83kWMUrO5jIsMt45vMrreaHOSnSAGDdge7vuEClL-GzcJdr_uaFCAJVEGYbmgV9VLkGu2XEgDJI0plrKLAqJRQBEKvzPtbaBuQVEDTqk-GOYPKrsUCPJryo43dHU4gUSQ"

	hCS, err := jwtdecode.JwtDecode.DecomposeFB(jwt)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	payload, err := jwtdecode.JwtDecode.DecodeClaimFB(hCS[1])
	if err != nil {
		fmt.Println("Error :", err)
	}

	user := payload.Subject
	email := payload.Email
	fmt.Printf("User ID: " + user + " ,Email: " + email + "\n")

}
