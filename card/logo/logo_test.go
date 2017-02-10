package logo

import (
	"path/filepath"
	"testing"

	"github.com/qgymje/wechat_api/token"
)

func TestUploadLogo(t *testing.T) {
	accessToken, err := token.AccessToken()
	if err != nil {
		t.Errorf("request access token failed, error: %s", err)
	}

	path, _ := filepath.Abs("./testdata/")
	filename := "/logo.jpg"
	resp, err := UploadLogo(accessToken.AccessToken, path+filename)
	if err != nil {
		t.Errorf("upload logo failed, error: %s", err)
	}
	if resp.URL == "" {
		t.Errorf("upload logo request failed, errcode: %d, errormsg: %s", resp.ErrorCode, resp.ErrorMsg)
	}
}
