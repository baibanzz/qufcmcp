package api

import "encoding/json"

type resLogin struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Token         string `json:"token"`
		Nickname      string `json:"nickname"`
		MerchantCodes string `json:"merchant_codes"`
		MerchantNames string `json:"merchant_names"`
		Id            string `json:"id"`
		RoleIds       string `json:"role_ids"`
		PermsList     string `json:"perms_list"`
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
	post, err := h.POST("/api/base/nologin/login", map[string]any{
		"user_name":   user,
		"pass_word":   pass,
		"verify_code": MFAcode,
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
