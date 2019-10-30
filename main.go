package main

import (
	"fmt"
	"log"

	"github.com/k-washi/jwt-decode/jwtdecode"
)

/*
https://github.com/firebase/firebase-admin-go/blob/master/auth/token_verifier.go
*/

func main() {
	//jwt := "eyJhbGciOiJSUzI1NiIsImtpZCI6ImEwYjQwY2NjYmQ0OWQxNmVkMjg2MGRiNzIyNmQ3NDZiNmZhZmRmYzAiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoidGVzdHVzZXIyIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL2V4LWZpcmViYXNlLWF1dGgiLCJhdWQiOiJleC1maXJlYmFzZS1hdXRoIiwiYXV0aF90aW1lIjoxNTcyNDQ3MTU0LCJ1c2VyX2lkIjoicVpoc0YySGZ1V1pFQmdoRmE0bmwyS2lkeXAyMiIsInN1YiI6InFaaHNGMkhmdVdaRUJnaEZhNG5sMktpZHlwMjIiLCJpYXQiOjE1NzI0NDcxNTQsImV4cCI6MTU3MjQ1MDc1NCwiZW1haWwiOiJ0ZXN0QHRlc3QuY29tIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7ImVtYWlsIjpbInRlc3RAdGVzdC5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9fQ.Y7DyZ70gkDBrPSfjIbmMWfObWH0GKLM7THyFMJAmMCgWu7dUfM5s1GKWUrzpnOw6qONJL6pHELWYqm0zn4UI-Gz0xlChA8CbllpiI5cbyTWpRGz4YyWfv0-8b983TmaWZr6RbZczTtfMefSo_fVnz-EexldmqYAmya6eVoqSideLqpUJR4MQA6Lx-VlsmQDhjSAaX0ULK6zCz3zGoFojgWfWDssBinSAcUivXDmqjg516cLkxIHExqIZb4WgR1x6-rD5KtnprPZjTa5IlaP5pK2FVZTC_JeU78CYm6Dut9Y6VuF9j7s-0mxqJdgPuxU-aZpq-rDBkdUMY6Br8CbiXw"
	//jwt := "eyJhbGciOiJSUzI1NiIsImtpZCI6ImEwYjQwY2NjYmQ0OWQxNmVkMjg2MGRiNzIyNmQ3NDZiNmZhZmRmYzAiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoidGVzdHVzZXIiLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vZXgtZmlyZWJhc2UtYXV0aCIsImF1ZCI6ImV4LWZpcmViYXNlLWF1dGgiLCJhdXRoX3RpbWUiOjE1NzIzMjIxNzQsInVzZXJfaWQiOiJxWmhzRjJIZnVXWkVCZ2hGYTRubDJLaWR5cDIyIiwic3ViIjoicVpoc0YySGZ1V1pFQmdoRmE0bmwyS2lkeXAyMiIsImlhdCI6MTU3MjMyMjE3NCwiZXhwIjoxNTcyMzI1Nzc0LCJlbWFpbCI6InRlc3RAdGVzdC5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsidGVzdEB0ZXN0LmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6InBhc3N3b3JkIn19.TMG8WXLrmcA6oiZpYB-Vt5QkZlrIU4_EoaZ-MEAXtAv9qYibCkztOKW2wXGYzq457TvrddIfQ8JdFYPfVNoZSPae7L_IWJqBHH_S_IKijhYWiyvRk-YkyyhThOvry5QexDP4Bvmfs9VQAOK-njjhiVYqAmTS2IPI_c4ep9Dl_R8L-ElK3NiRKY32RK3Y2paxKoCVG0xyc2Rd2HPl8BUk10pzDnmBbYDeFK3buP0irjT5B5w7ub6hSwyj4MlCKxdyvHM4GMNipW-hoMDyXbyAZVyvPARKfLjAQvfLU5Ni4Pb6UTYLtPXPsjPLILzV3mVDqBXPp1ZaAZJ15UyadSjDVA"
	jwt := "eyJhbGciOiJSUzI1NiIsImtpZCI6ImEwYjQwY2NjYmQ0OWQxNmVkMjg2MGRiNzIyNmQ3NDZiNmZhZmRmYzAiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoidGVzdHVzZXIyIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL2V4LWZpcmViYXNlLWF1dGgiLCJhdWQiOiJleC1maXJlYmFzZS1hdXRoIiwiYXV0aF90aW1lIjoxNTcyNDUwNTIwLCJ1c2VyX2lkIjoicVpoc0YySGZ1V1pFQmdoRmE0bmwyS2lkeXAyMiIsInN1YiI6InFaaHNGMkhmdVdaRUJnaEZhNG5sMktpZHlwMjIiLCJpYXQiOjE1NzI0NTA1MjAsImV4cCI6MTU3MjQ1NDEyMCwiZW1haWwiOiJ0ZXN0QHRlc3QuY29tIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7ImVtYWlsIjpbInRlc3RAdGVzdC5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9fQ.PD78j1K17EYus1IgmvKhLQE9bq1XUKBvON2OC1I80fG22I84koRjblWEJGPHYt7nX3cTUAJ1u8Opq3cMOogR0MnvTHVDCpIHCDwlIwuLlNgyr4cDaGX2iDWGU2OSOJddpKXa1X9NPWZGywRaRitnNvkUu5ztnxifEKaWJfL1ubl_giacLMPwjri2tJjhvO6xZlkHOpyYUn4mP-BcFFJNWE7K9vEVz05k-_dz4xY0hlkdstyq1u9tfWaUy-HZR9hn69zR4AuygVcHKqGtshV1yyPU7C2oyCkDhcUXa8MwF9xj9B2qzkNgF3ftoj7DXVWqEi_QdeQAf3X84YqadscbbA"
	hCS, err := jwtdecode.JwtDecode.DecomposeFB(jwt)
	if err != nil {
		log.Fatalln("Error : ", err)
	}

	payload, err := jwtdecode.JwtDecode.DecodeClaimFB(hCS[1])
	if err != nil {
		log.Fatalln("Error :", err)
	}

	user := payload.Subject
	email := payload.Email
	fmt.Printf("User ID: " + user + " ,Email: " + email + "\n")

}
