package meta

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/projecthelper"
)

func (h *Handler) GetCount(c *gin.Context) {
	table := c.Param("table")
	items, err := h.service.GetCount(&table)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetSchema(c *gin.Context) {
	c.JSON(http.StatusOK, projecthelper.SchemasLib)
}

type Address struct {
	ID        string      `json:"id"`
	FullName  string      `json:"fullName"`
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	TypeShort string      `json:"typeShort"`
	Zip       interface{} `json:"zip"`
	Parents   Addresses   `json:"parents"`
}

type Addresses []*Address

type Result struct {
	Result Addresses `json:"result"`
}

func (h *Handler) GetAddress(c *gin.Context) {
	var item models.KladrAPI
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	ctx := context.Background()
	url := item.GetURLOneString()
	fmt.Println(url)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
	}
	cli := &http.Client{}
	resp, err := cli.Do(request)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	resp.Body.Close()
	res := Result{}
	err = json.Unmarshal(body, &res)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	for _, v := range res.Result {
		fmt.Println(v)
	}
	// c.JSON(http.StatusOK, res.Result[1:])
	c.JSON(http.StatusOK, res.Result)
}
