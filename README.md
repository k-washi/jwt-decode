# jwt-decode

---

Parse user information from JWT created by firebase

## Install

```bash
go get -u "github.com/k-washi/jwt-decode/jwtdecode"
``` 

## example

```bash
> go run main.go
User ID: qZhsF2HfuWZEBghFa4nl2Kidyp22 ,Email: test@test.com
```

## payload

jwtのpayloadは以下のように表現されている。

```json
{"name":"testuser","iss":"https://securetoken.google.com/ex-firebase-auth","aud":"ex-firebase-auth","auth_time":1572322174,"user_id":"qZhsF2HfuWZEBghFa4nl2Kidyp22","sub":"qZhsF2HfuWZEBghFa4nl2Kidyp22","iat":1572322174,"exp":1572325774,"email":"test@test.com","email_verified":false,"firebase":{"identities":{"email":["test@test.com"]},"sign_in_provider":"password"}}

```