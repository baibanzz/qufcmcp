package api

// FindPageAdminUser 管理员分页查询 GET /api/adminUser/findPage
func (h *HTTP) FindPageAdminUser(params map[string]string) (string, error) {
	queryParams := make(map[string]any)
	for k, v := range params {
		queryParams[k] = v
	}
	body, err := h.GET("/api/adminUser/findPage", queryParams)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
