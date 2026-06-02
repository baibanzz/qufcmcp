package api

// PageUserRepeat 用户查重查询 GET /api/fcUserMaterial/repeat
func (h *HTTP) PageUserRepeat(params map[string]string) (string, error) {
	queryParams := make(map[string]any)
	for k, v := range params {
		queryParams[k] = v
	}
	body, err := h.GET("/api/fcUserMaterial/repeat", queryParams)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// SaveRepeatUser 保存/更新用户信息 POST /api/fcUserMaterial/saverepeat
func (h *HTTP) SaveRepeatUser(data map[string]any) (string, error) {
	body, err := h.POST("/api/fcUserMaterial/saverepeat", data)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
