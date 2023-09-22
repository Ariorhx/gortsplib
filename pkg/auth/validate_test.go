package auth

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Ariorhx/gortsplib/pkg/base"
)

func TestValidateErrors(t *testing.T) {
	for _, ca := range []struct {
		name string
		hv   base.HeaderValue
		err  string
	}{
		{
			"invalid auth",
			base.HeaderValue{`Invalid`},
			"invalid authorization header",
		},
		{
			"digest missing realm",
			base.HeaderValue{`Digest `},
			"realm is missing",
		},
		{
			"digest missing nonce",
			base.HeaderValue{`Digest realm=123`},
			"nonce is missing",
		},
		{
			"digest missing username",
			base.HeaderValue{`Digest realm=123,nonce=123`},
			"username is missing",
		},
		{
			"digest missing uri",
			base.HeaderValue{`Digest realm=123,nonce=123,username=123`},
			"uri is missing",
		},
		{
			"digest missing response",
			base.HeaderValue{`Digest realm=123,nonce=123,username=123,uri=123`},
			"response is missing",
		},
		{
			"digest wrong nonce",
			base.HeaderValue{`Digest realm=123,nonce=123,username=123,uri=123,response=123`},
			"wrong nonce",
		},
		{
			"digest wrong realm",
			base.HeaderValue{`Digest realm=123,nonce=abcde,username=123,uri=123,response=123`},
			"wrong realm",
		},
	} {
		t.Run(ca.name, func(t *testing.T) {
			err := Validate(
				&base.Request{
					Method: base.Describe,
					URL:    nil,
					Header: base.Header{
						"Authorization": ca.hv,
					},
				},
				"myuser",
				"mypass",
				nil,
				nil,
				"IPCAM",
				"abcde",
			)
			require.EqualError(t, err, ca.err)
		})
	}
}
