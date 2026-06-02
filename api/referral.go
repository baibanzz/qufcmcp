package api

// UpdateReferral 修改推荐关系 POST /api/referral/update
func (h *HTTP) UpdateReferral(data map[string]any) (string, error) {
	body, err := h.POST("/api/referral/update", data)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
