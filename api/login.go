package api

import "encoding/json"

type resLogin struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Token         string `json:"Token"`
		Nickname      string `json:"Nickname"`
		MerchantCodes string `json:"MerchantCodes"`
		MerchantNames string `json:"MerchantNames"`
		Id            string `json:"Id"`
		RoleIds       string `json:"RoleIds"`
		PermsList     string `json:"PermsList"`
		AccountType   int    `json:"account_type"`
		Department    string `json:"department"`
		LoginIp       string `json:"login_ip"`
		Mobile        string `json:"mobile"`
		IpInfo        struct {
			Country string `json:"country"`
			City    string `json:"city"`
		} `json:"ip_info"`
	} `json:"data"`
}

func (h *HTTP) BaseLogin(user, pass, MFAcode string) (string, error) {
	// 先发一个 GET 请求获取 aws_session Cookie（登录接口需要先有 Cookie）
	_, err := h.GET("/api/fcUserMaterial/findPage", map[string]any{})
	if err != nil {
		return "", err
	}

	post, err := h.POST("/api/base/nologin/login", map[string]any{
		"userName":   user,
		"password":   pass,
		"verifyCode": MFAcode,
	})
	if err != nil {
		return "", err
	}
	var ret resLogin
	if err := json.Unmarshal(post, &ret); err != nil {
		return "", err
	}
	h.token = ret.Data.Token
	return h.token, nil
}
