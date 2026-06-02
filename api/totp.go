package api

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

// GenerateTOTP 根据 base32 密钥生成当前时间窗口的 6 位 TOTP 验证码
// 使用标准 TOTP 参数：30 秒周期、SHA1、6 位数字
func GenerateTOTP(secret string) (string, error) {
	// 去除空格并转为大写
	secret = strings.ToUpper(strings.TrimSpace(secret))

	// Base32 解码密钥
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secret)
	if err != nil {
		// 尝试带填充的解码
		key, err = base32.StdEncoding.DecodeString(secret)
		if err != nil {
			return "", fmt.Errorf("无效的 base32 密钥: %v", err)
		}
	}

	// 计算当前时间窗口（30 秒周期）
	counter := time.Now().Unix() / 30

	// 将 counter 转为 8 字节大端字节序
	msg := make([]byte, 8)
	binary.BigEndian.PutUint64(msg, uint64(counter))

	// HMAC-SHA1
	mac := hmac.New(sha1.New, key)
	mac.Write(msg)
	hash := mac.Sum(nil)

	// 动态截断
	offset := hash[len(hash)-1] & 0x0f
	code := (int32(hash[offset]&0x7f) << 24) |
		(int32(hash[offset+1]&0xff) << 16) |
		(int32(hash[offset+2]&0xff) << 8) |
		int32(hash[offset+3]&0xff)

	// 取模 10^6 得到 6 位数字
	otp := code % 1000000

	return fmt.Sprintf("%06d", otp), nil
}
