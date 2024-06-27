package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"my-gin/app/common/response"
	"my-gin/global"
	"sort"

	"github.com/gin-gonic/gin"
)

func CheckSignMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求体
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			response.Fail(c, 1001, "Unable to read request body")
			c.Abort() // 终止后续处理
			return
		}

		// 重新填充请求体以便后续处理
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		// 解析请求体参数
		var params map[string]interface{}
		if err := json.Unmarshal(body, &params); err != nil {
			response.Fail(c, 1002, "Invalid request body")
			c.Abort() // 终止后续处理
			return
		}

		sign, ok := params["sign"].(string)
		if !ok || sign == "" {
			response.Fail(c, 1003, "Missing signature")
			c.Abort() // 终止后续处理
			return
		}

		// 删除签名字段以进行签名验证
		delete(params, "sign")

		// 按照字段名对参数进行排序
		keys := make([]string, 0, len(params))
		for k := range params {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		// 按照排序后的顺序生成 JSON 数据
		sortedParams := make(map[string]interface{})
		for _, k := range keys {
			sortedParams[k] = params[k]
		}

		bodyWithoutSign, err := json.Marshal(sortedParams)
		if err != nil {
			response.Fail(c, 1004, "Unable to marshal request body without sign")
			c.Abort() // 终止后续处理
			return
		}

		// 计算签名
		key := []byte(global.App.Config.App.SignKey)
		mac := hmac.New(sha256.New, key)
		mac.Write(bodyWithoutSign)
		expectedSign := hex.EncodeToString(mac.Sum(nil))

		// 验证签名
		if !hmac.Equal([]byte(sign), []byte(expectedSign)) {
			response.Fail(c, 1005, "Invalid signature")
			c.Abort() // 终止后续处理
			return
		}

		c.Next()
	}
}
