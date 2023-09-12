package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/helper"
)

type Middleware struct {
	helper *helper.Helper
}

func CreateMiddleware(helper *helper.Helper) *Middleware {
	return &Middleware{helper: helper}
}

func (m *Middleware) InjectRequestInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := Claims{ClaimUserID, ClaimDomainIDS}.Inject(c.Request, m.helper.Token)
		if m.helper.HTTP.HandleError(c, err) {
			return
		}
		if err != nil {
			return
		}
		err = m.helper.SQL.InjectQueryFilter(c)
		if m.helper.HTTP.HandleError(c, err) {
			return
		}

		if err != nil {
			return
		}
		c.Next()
	}
}
