package tests

import (
    "crypto/rsa"
    "io/ioutil"
    "strings"
    "fmt"

    "github.com/dgrijalva/jwt-go"
    "github.com/revel/revel/testing"
)

type RevelJWTMysqlTest struct {
    testing.TestSuite
}

var rsaPSSTestData = []struct {
    name        string
    tokenString string
    alg         string
    claims      map[string]interface{}
    valid       bool
}{
    {
        "Basic RS256",
        "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDExODgxOSwiZXhwIjoxNDU0MTIyNDE5LCJpYXQiOjE0NTQxMTg4MTksImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.iAenRMV9u2CelASLw8Ol4w-3HRbpApIRfN-PCo0SdSUd9kvTO6EeWhSn6EReb4u4-zEajohLrgaS5f308j5NZW4K8aabd-GaBrpWHlf8YY5ueqOYbrmOQcevjnyZlfc05nCUPBWaG0cPeADh_zppXRkiamHfUSwgJQA-Dhs6J04",
        "RS256",
        map[string]interface{}{
        "iss": "https://jwt-idp.example.com",
        "sub": "mailto:mike@example.com",
        "nbf": 1454118819,
        "exp": 1454122419,
        "iat": 1454118819,
        "jti": "id123456",
        "typ": "https://example.com/register",
        },
		true,
    },
}

// verify key check
func (t *RevelJWTMysqlTest) TestRSAPSSVerify() {
    var err error

    key, _ := ioutil.ReadFile("demo.rsa.pub")
    var rsaPSSKey *rsa.PublicKey
    if rsaPSSKey, err = jwt.ParseRSAPublicKeyFromPEM(key); err != nil {
        fmt.Println("Unable to parse RSA public key: %v", err)
    }

    for _, data := range rsaPSSTestData {
        parts := strings.Split(data.tokenString, ".")

        method := jwt.GetSigningMethod(data.alg)
        err := method.Verify(strings.Join(parts[0:2], "."), parts[2], rsaPSSKey)
        if data.valid && err != nil {
            fmt.Println("[%v] Error while verifying key: %v", data.name, err)
        }
        if !data.valid && err == nil {
            fmt.Println("[%v] Invalid key passed validation", data.name)
        }
    }
}

// Token Check
func (t *RevelJWTMysqlTest) TestRSAPSSSign() {
    var err error

    key, _ := ioutil.ReadFile("demo.rsa")
    var rsaPSSKey *rsa.PrivateKey
    if rsaPSSKey, err = jwt.ParseRSAPrivateKeyFromPEM(key); err != nil {
        fmt.Println("Unable to parse RSA private key: %v", err)
    }

    for _, data := range rsaPSSTestData {
        if data.valid {
            parts := strings.Split(data.tokenString, ".")
            method := jwt.GetSigningMethod(data.alg)
            sig, err := method.Sign(strings.Join(parts[0:2], "."), rsaPSSKey)
            if err != nil {
                fmt.Println("[%v] Error signing token: %v", data.name, err)
            }
            if sig == parts[2] {
                fmt.Println("[%v] Signatures shouldn't match\nnew:\n%v\noriginal:\n%v", data.name, sig, parts[2])
            }
        }
    }
}