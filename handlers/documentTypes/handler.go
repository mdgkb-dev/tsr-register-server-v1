package documentTypes

import (
	"github.com/gin-gonic/gin"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.DocumentType
	err := c.Bind(&item)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	err = h.service.Create(&item)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, item)
}

func (h *Handler) GetAll(c *gin.Context) {
	items, err := h.service.GetAll()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, items)
}

//
func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(&id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, item)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(&id)
	if err != nil {
		c.JSON(500, err.Error())
	}
	c.JSON(200, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.DocumentType
	err := c.Bind(&item)
	if err != nil {
		c.JSON(500, err)
	}
	err = h.service.Update(&item)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, item)
}
