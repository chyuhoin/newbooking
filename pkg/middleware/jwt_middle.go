package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"newbooking/pkg/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//过滤是否验证token， 示例代码没有完整设置白名单路径
		enableUrl := make([]string, 0)
		enableUrl = append(enableUrl, "/login")
		enableUrl = append(enableUrl, "/register")
		for _, url := range enableUrl {
			if url == c.Request.RequestURI {
				return
			}
		}

		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		j := utils.NewJWT()
		// parse token, get the user and role info
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				c.JSON(http.StatusForbidden, gin.H{
					"status": -1,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set(utils.GinContextKey, claims)
	}
}
