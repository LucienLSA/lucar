package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`

func TestGenerateToken(t *testing.T) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("can not parse private key: %v", err)
	}
	g := NewJwtTokenGen("lucar/auth", key)
	g.nowFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}
	tkn, err := g.GenerateToken("5f7c245ab0361e00ffb9fd6f", 2*time.Hour)
	if err != nil {
		t.Errorf("can nt generate token %v", err)
	}
	want := `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoibHVjYXIvYXV0aCIsInN1YiI6IjVmN2MyNDVhYjAzNjFlMDBmZmI5ZmQ2ZiJ9.EJwI-tyqjNVbF1w2jdFDffR6cnPB0gHQYhuxhUvAKMD_RrNe9emARxsEuNGIedBD2rlSmvK36sxvfnU9V13Z46qstet1vc5IZltyF_hJd2dmBT7uBQPIcp_9JQ6j-lXXNBAcXOW_kIXglX8OAKflNIBoGZVoQwdNOIGt5b2UP0TAn23pxBkYVIVqGezRb8AvZvy7RqgY37Rh34mqqYTn29nXpvRADRjkkmaMFBHVSET1BiUhYVmf_IpvA-SexcbFkV2HhVQlCf9Kgo9toOZhGKFishSz934V68Ie2YfdViB3f400ZX2ammdHvRIRYlfoXtjA-AOpSFZgmBDlTYIIoSMYJVt5-ozPLfFd5zYKtuN8sP1PJI80PSof5UTsNcPbOwNh9NYqG7JuLcR0zzSjpsInEmOucJo1fBE_Zy4a7KWzA7c4LSH73ZB6H8MdCVBeZdkLZxGVyqeSehj9B2b-NA-BOhf6vTc5831Qssntezv6IDlxq7M4v0PEMI1KHL44rHlgVVKds6V4zEZs0hrZy30YBbBXkEWcnRg3i5bVCLYTAnsSprHUbP2zjm-H8WL_fBdKHeFcbg_ZmtF3LgKyYTetYqhSkDT2hEwuLOY-JLi0-ZhanHBbNBaKP2vvJ5s7T2HuVILT8YEZxj27lD4xU8WPbjUxy11wtTSOtz8EZag`
	if tkn != want {
		t.Errorf("wrong token generated. want:%q; got:%q", want, tkn)
	}
}
