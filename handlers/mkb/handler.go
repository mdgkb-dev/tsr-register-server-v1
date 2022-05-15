package mkb

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllClasses(c *gin.Context) {
	items, err := h.service.GetAllClasses()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Update(c *gin.Context) {
	name := c.Query("update")
	fmt.Println(name)
	id := c.Param("id")
	mkbType := c.Query("mkbType")
	if name == "name" {
		type Name struct {
			Name string `json:"name"`
		}
		nameName := Name{}
		err := c.Bind(&nameName)
		if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
			return
		}
		fmt.Println(nameName.Name)
		err = h.service.UpdateName(&id, &nameName.Name, &mkbType)
		if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
			return
		}
	} else {
		err := h.service.UpdateRelevant(&id, &mkbType)
		if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
			return
		}
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) GetGroupByClassId(c *gin.Context) {
	classId := c.Param("classId")
	items, err := h.service.GetGroupByClassId(&classId)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetGroupChildrens(c *gin.Context) {
	groupId := c.Param("groupId")
	items, err := h.service.GetGroupChildrens(&groupId)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetSubGroupChildrens(c *gin.Context) {
	subGroupId := c.Param("subGroupId")
	items, err := h.service.GetGroupByClassId(&subGroupId)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetGroupsBySearch(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusOK, nil)
	}
	items, err := h.service.GetGroupsBySearch(&query)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetDiagnosisBySearch(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusOK, nil)
	}
	items, err := h.service.GetDiagnosisBySearch(&query)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetDiagnosisByGroupId(c *gin.Context) {
	groupId := c.Param("groupId")
	items, err := h.service.GetDiagnosisByGroupId(&groupId)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetSubDiagnosisByDiagnosisId(c *gin.Context) {
	diagnosisId := c.Param("diagnosisId")
	items, err := h.service.GetSubDiagnosisByDiagnosisId(&diagnosisId)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}
