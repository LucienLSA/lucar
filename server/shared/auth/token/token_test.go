package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAl1T0t+B9f5ei+IiUodgP
9D9MJxJidCsl6xtcPitvGSODz60Gku1WtcosPkPkD09hSs0IbucVQKXmP+aWO+2a
8OHeaMu5lI8zrGE/Bg5hImOzKJW1rCMn1dDccVqH+VgUsg4Dokdo2YRGNy0kadi7
VV9C65bAjPeuosMP4iyLJuuqPYTNqdtIWOBMfII2uA7Iy1HcPXYnyvNRLrgTJwFK
L02tUj9bvtFCwbfcmauHM2iCw5SYeW7Gxas3i8mLtIbvyKXBKwLQaHGA8SyZzuGv
9kJg8vymDggwRyrvOy68j/rpH7Yk8w+eDrWAekNZZlNR8UZyJcZOrr5ziP2Og1SJ
D5JTx4HzFpZFgYWop/oKMH2S37pgNxQo+kY3Abw+OnjK9Z4zh2o9q3uGP56f/TGe
UwgPCkC0zQJC+P09oLDvhLzcHJ2QSDa4IyXq2s75c5TGdzLRyhc7eLmCbwGpWTXu
iAYDTQlh3xx2bkl+z/EWCdOUFpMJ38Scy9/QitIOkPDVdtc/cxxDbj65h4uYkHXP
Gt7hOsrFtK+nWZmuqYU4jg+1cpUW4C7vsk0ll1OlOhfp5tc5DKgBBfFKhs2YZV7P
/GhzfjiwAyK6nqDAyJh02oOau6hrRPhDiGsQfK9ZLnDmr/mmmghqJy2qRSsTotfO
snRqiN1aC3PEpKZrHsF36YcCAwEAAQ==
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("can not parse public key:%v", err)
	}
	v := JWTTokenVerifier{
		PublicKey: pubKey,
	}
	cases := []struct {
		name    string
		tkn     string
		now     time.Time
		want    string
		wantErr bool
	}{
		{
			name:    "valid_token",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoibHVjYXIvYXV0aCIsInN1YiI6IjVmN2MyNDVhYjAzNjFlMDBmZmI5ZmQ2ZiJ9.EJwI-tyqjNVbF1w2jdFDffR6cnPB0gHQYhuxhUvAKMD_RrNe9emARxsEuNGIedBD2rlSmvK36sxvfnU9V13Z46qstet1vc5IZltyF_hJd2dmBT7uBQPIcp_9JQ6j-lXXNBAcXOW_kIXglX8OAKflNIBoGZVoQwdNOIGt5b2UP0TAn23pxBkYVIVqGezRb8AvZvy7RqgY37Rh34mqqYTn29nXpvRADRjkkmaMFBHVSET1BiUhYVmf_IpvA-SexcbFkV2HhVQlCf9Kgo9toOZhGKFishSz934V68Ie2YfdViB3f400ZX2ammdHvRIRYlfoXtjA-AOpSFZgmBDlTYIIoSMYJVt5-ozPLfFd5zYKtuN8sP1PJI80PSof5UTsNcPbOwNh9NYqG7JuLcR0zzSjpsInEmOucJo1fBE_Zy4a7KWzA7c4LSH73ZB6H8MdCVBeZdkLZxGVyqeSehj9B2b-NA-BOhf6vTc5831Qssntezv6IDlxq7M4v0PEMI1KHL44rHlgVVKds6V4zEZs0hrZy30YBbBXkEWcnRg3i5bVCLYTAnsSprHUbP2zjm-H8WL_fBdKHeFcbg_ZmtF3LgKyYTetYqhSkDT2hEwuLOY-JLi0-ZhanHBbNBaKP2vvJ5s7T2HuVILT8YEZxj27lD4xU8WPbjUxy11wtTSOtz8EZag",
			now:     time.Unix(1516239022, 0),
			want:    "5f7c245ab0361e00ffb9fd6f",
			wantErr: false,
		},
		{
			name:    "token_expired",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoibHVjYXIvYXV0aCIsInN1YiI6IjVmN2MyNDVhYjAzNjFlMDBmZmI5ZmQ2ZiJ9.EJwI-tyqjNVbF1w2jdFDffR6cnPB0gHQYhuxhUvAKMD_RrNe9emARxsEuNGIedBD2rlSmvK36sxvfnU9V13Z46qstet1vc5IZltyF_hJd2dmBT7uBQPIcp_9JQ6j-lXXNBAcXOW_kIXglX8OAKflNIBoGZVoQwdNOIGt5b2UP0TAn23pxBkYVIVqGezRb8AvZvy7RqgY37Rh34mqqYTn29nXpvRADRjkkmaMFBHVSET1BiUhYVmf_IpvA-SexcbFkV2HhVQlCf9Kgo9toOZhGKFishSz934V68Ie2YfdViB3f400ZX2ammdHvRIRYlfoXtjA-AOpSFZgmBDlTYIIoSMYJVt5-ozPLfFd5zYKtuN8sP1PJI80PSof5UTsNcPbOwNh9NYqG7JuLcR0zzSjpsInEmOucJo1fBE_Zy4a7KWzA7c4LSH73ZB6H8MdCVBeZdkLZxGVyqeSehj9B2b-NA-BOhf6vTc5831Qssntezv6IDlxq7M4v0PEMI1KHL44rHlgVVKds6V4zEZs0hrZy30YBbBXkEWcnRg3i5bVCLYTAnsSprHUbP2zjm-H8WL_fBdKHeFcbg_ZmtF3LgKyYTetYqhSkDT2hEwuLOY-JLi0-ZhanHBbNBaKP2vvJ5s7T2HuVILT8YEZxj27lD4xU8WPbjUxy11wtTSOtz8EZag",
			now:     time.Unix(1517239022, 0),
			wantErr: true,
		},
		{
			name:    "bad_token",
			tkn:     "dsad",
			now:     time.Unix(1516239022, 0),
			wantErr: true,
		},
		{
			name:    "fake_token",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoibHVjYXIvYXV0aCIsInN1YiI6IjVmN2MyNDVhYjAzNjFlMDBmZmI5ZmQ2ZiJ9.EJwI-tyqjNVbF1w2jdFDffR6cnPB0gHQYhuxhUvAKMD_RrNe9emARxsEuNGIedBD2rlSmvK36sxvfnU9V13Z46qstet1vc5IZltyF_hJd2dmBT7uBQPIcp_9JQ6j-lXXNBAcXOW_kIXglX8OAKflNIBoGZVoQwdNOIGt5b2UP0TAn23pxBkYVIVqGezRb8AvZvy7RqgY37Rh34mqqYTn29nXpvRADRjkkmaMFBHVSET1BiUhYVmf_IpvA-SexcbFkV2HhVQlCf9Kgo9toOZhGKFishSz934V68Ie2YfdViB3f400ZX2ammdHvRIRYlfoXtjA-AOpSFZgmBDlTYIIoSMYJVt5-ozPLfFd5zYKtuN8sP1PJI80PSof5UTsNcPbOwNh9NYqG7JuLcR0zzSjpsInEmOucJo1fBE_Zy4a7KWzA7c4LSH73ZB6H8MdCVBeZdkLZxGVyqeSehj9B2b-NA-BOhf6vTc5831Qssntezv6IDlxq7M4v0PEMI1KHL44rHlgVVKds6V4zEZs0hrZy30YBbBXkEWcnRg3i5bVCLYTAnsSprHUbP2zjm-H8WL_fBdKHeFcbg_ZmtF3LgKyYTetYqhSkDT2hEwuLOY-JLi0-ZhanHBbNBaKP2vvJ5s7T2HuVILT8YEZxj27lD4xU8WPbjUxy11wtTSOtz8E",
			now:     time.Unix(1516239022, 0),
			wantErr: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			jwt.TimeFunc = func() time.Time {
				return c.now
			}
			accountID, err := v.Verifier(c.tkn)
			if !c.wantErr && err != nil {
				t.Errorf("verification failed:%v", err)
			}
			if c.wantErr && err == nil {
				t.Errorf("want error; got not error")
			}
			if accountID != c.want {
				t.Errorf("wrong account id, want: %q, got:%q", c.want, accountID)
			}
		})
	}
	// tkn := `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoibHVjYXIvYXV0aCIsInN1YiI6IjVmN2MyNDVhYjAzNjFlMDBmZmI5ZmQ2ZiJ9.EJwI-tyqjNVbF1w2jdFDffR6cnPB0gHQYhuxhUvAKMD_RrNe9emARxsEuNGIedBD2rlSmvK36sxvfnU9V13Z46qstet1vc5IZltyF_hJd2dmBT7uBQPIcp_9JQ6j-lXXNBAcXOW_kIXglX8OAKflNIBoGZVoQwdNOIGt5b2UP0TAn23pxBkYVIVqGezRb8AvZvy7RqgY37Rh34mqqYTn29nXpvRADRjkkmaMFBHVSET1BiUhYVmf_IpvA-SexcbFkV2HhVQlCf9Kgo9toOZhGKFishSz934V68Ie2YfdViB3f400ZX2ammdHvRIRYlfoXtjA-AOpSFZgmBDlTYIIoSMYJVt5-ozPLfFd5zYKtuN8sP1PJI80PSof5UTsNcPbOwNh9NYqG7JuLcR0zzSjpsInEmOucJo1fBE_Zy4a7KWzA7c4LSH73ZB6H8MdCVBeZdkLZxGVyqeSehj9B2b-NA-BOhf6vTc5831Qssntezv6IDlxq7M4v0PEMI1KHL44rHlgVVKds6V4zEZs0hrZy30YBbBXkEWcnRg3i5bVCLYTAnsSprHUbP2zjm-H8WL_fBdKHeFcbg_ZmtF3LgKyYTetYqhSkDT2hEwuLOY-JLi0-ZhanHBbNBaKP2vvJ5s7T2HuVILT8YEZxj27lD4xU8WPbjUxy11wtTSOtz8EZag`
	// jwt.TimeFunc = func() time.Time {
	// 	return time.Unix(1517239122, 0)
	// }
	// accountID, err := v.Verifier(tkn)
	// if err != nil {
	// 	t.Errorf("verificate failed:%v", err)
	// }
	// want := "5f7c245ab0361e00ffb9fd6f"
	// if accountID != want {
	// 	t.Errorf("wrong account id, want: %q, got:%q", want, accountID)
	// }
}
