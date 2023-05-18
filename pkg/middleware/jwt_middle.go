package middleware

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"newbooking/pkg/utils"
	"os"
	"strings"
)

type whiteList struct {
	Whitelist []string
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//过滤是否验证token
		list, err := getWhiteList("whitelist.json")
		if err != nil {
			panic("no whitelist")
			return
		}
		enableUrl := list.Whitelist
		fmt.Println(enableUrl)
		for _, url := range enableUrl {
			fmt.Println(url, strings.Split(c.Request.RequestURI, "?")[0])
			if url == strings.Split(c.Request.RequestURI, "?")[0] {
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

func getWhiteList(path string) (*whiteList, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	var wl whiteList

	if err = decoder.Decode(&wl); err != nil {
		return nil, err
	}
	return &wl, nil
}
