package logo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	maxsize = 1024 * 1024
)

var (
	url              = "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%s"
	allowedImageType = []string{"jpg", "png"}
)

type LogoResponse struct {
	ErrorCode int    `json:"errcode"`
	ErrorMsg  string `json:"errmsg"`
	URL       string `json:"url"`
}

func UploadLogo(token, path string) (*LogoResponse, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)
	w, err := mw.CreateFormFile("buffer", path)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(w, file); err != nil {
		return nil, err
	}
	defer mw.Close()

	_ = mw.WriteField("hack", "")
	req, err := http.NewRequest("POST", fmt.Sprintf(url, token), body)
	req.Header.Set("Content-Type", "multipart/form-data")
	//req.Header.Set("Content-Type", mw.FormDataContentType())

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	logoResponse := &LogoResponse{}
	err = json.Unmarshal(respbody, logoResponse)
	if err != nil {
		return nil, err
	}
	return logoResponse, nil
}
