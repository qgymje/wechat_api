package token

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	appid     = "wx94347eee565b4df7"
	appsecret = "ee5333d5a0675b9ac9c129a3d3509a89"
)

var url = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

type AccessTokenResponse struct {
	ErrorCode int    `json:"errorcode"`
	ErrMsg    string `json:"errmsg"`

	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func AccessToken() (*AccessTokenResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(url, appid, appsecret), nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	accessTokenResponse := &AccessTokenResponse{}
	err = json.Unmarshal(body, accessTokenResponse)
	if err != nil {
		return nil, err
	}
	return accessTokenResponse, nil
}
