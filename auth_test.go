package echo

import (
	"net/http"
	"net/url"
	"testing"
)

func TestAuth(t *testing.T) {
	vals := url.Values{
		"client_id":    {"8ba2991113e68b4805c1"},
		"redirect_uri": {"http://shaalx-echouj.daoapp.io/callback"},
		"scope":        {"repo"},
		"state":        {"state"},
	}
	v := vals.Encode()
	url_, _ := url.Parse("https://github.com/login/oauth/authorize")
	url_.RawQuery = v

	t.Log(url_)
	req, err := http.NewRequest("GET", url_.String(), nil)
	t.Log(req, err)
}
