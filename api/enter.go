package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type HTTP struct {
	baseURL string //平台链接
	client  *http.Client
	token   string //token
}

func (h *HTTP) POST(URL string, data map[string]any) ([]byte, error) {
	marshal, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", h.baseURL+URL, bytes.NewBuffer(marshal))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", h.token)
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (h *HTTP) POSTWithCookie(URL string, data map[string]any, cookies map[string]string) ([]byte, error) {
	marshal, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", h.baseURL+URL, bytes.NewBuffer(marshal))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", h.token)
	for key, value := range cookies {
		req.AddCookie(&http.Cookie{Name: key, Value: value})
	}
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (h *HTTP) GET(URL string, data map[string]any) ([]byte, error) {
	var queryParams []string
	for key, value := range data {
		queryParams = append(queryParams, key+"="+url.QueryEscape(fmt.Sprintf("%v", value)))
	}
	if len(queryParams) > 0 {
		URL = URL + "?" + strings.Join(queryParams, "&")
	}
	req, err := http.NewRequest("GET", h.baseURL+URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Token", h.token)
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (h *HTTP) SetToken(token string) {
	h.token = token
}

func New(URL string, Token string) *HTTP {
	return &HTTP{
		baseURL: URL,
		token:   Token,
		client:  &http.Client{},
	}
}
