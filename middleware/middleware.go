package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"
)

type Middleware struct {
	helper *helper.Helper
}

func CreateMiddleware(helper *helper.Helper) *Middleware {
	return &Middleware{helper: helper}
}

func (m *Middleware) InjectFTSP() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.Request.URL.Path, "ftsp") {
			return
		}
		ftspQuery := &sqlHelper.FTSPQuery{}
		err := ftspQuery.FromForm(c)
		if m.helper.HTTP.HandleError(c, err) {
			return
		}

		ftsp, found := ftspStore.GetOrCreateFTSP(ftspQuery)

		if !found {
			c.JSON(http.StatusOK, nil)
			c.Abort()
			return
		}

		m.helper.SQL.InjectFTSP2(c.Request, &ftsp)

		if err != nil {
			return
		}
		c.Next()
	}
}

func (m *Middleware) InjectRequestInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := Claims{ClaimUserID, ClaimDomainIDS}.Inject(c.Request, m.helper.Token)
		fmt.Println(err)
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
