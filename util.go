package wechatwork

import (
	"fmt"
	"net/url"
)

// FormatURL 格式化url
func FormatURL(apiPath string, values url.Values) string {
	return fmt.Sprintf("%s?%s", apiPath, values.Encode())
}
