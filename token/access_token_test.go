package token

import (
	"testing"
)

func TestAccessToken(t *testing.T) {
	resp, err := AccessToken()
	if err != nil {
		t.Errorf("request access token failed, error: %s", err)
	}
	if resp.AccessToken == "" {
		t.Errorf("request access token failed, errcode: %d, errmsg: %s", resp.ErrorCode, resp.ErrMsg)
	}
}
