package mkbitems

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//
//func (h *Handler) GetAllClasses(c *gin.Context) {
//	items, err := h.service.GetAllClasses()
//	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
//		return
//	}
//	c.JSON(http.StatusOK, items)
//}
//
//func (h *Handler) Update(c *gin.Context) {
//	name := c.Query("update")
//	id := c.Param("id")
//	mkbType := c.Query("mkbType")
//	if name == "name" {
//		type Name struct {
//			Name string `json:"name"`
//		}
//		nameName := Name{}
//		err := c.Bind(&nameName)
//		if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
//			return
//		}
//		fmt.Println(nameName)
//		err = h.service.UpdateName(id, nameName.Name, mkbType)
//		if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
//			return
//		}
//	} else {
//		err := h.service.UpdateRelevant(id, mkbType)
//		if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
//			return
//		}
//	}
//	c.JSON(http.StatusOK, nil)
//}
//
//func (h *Handler) GetGroupByClassID(c *gin.Context) {
//	classID := c.Param("classId")
//	items, err := h.service.GetGroupByClassID(classID)
//	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
//		return
//	}
//	c.JSON(http.StatusOK, items)
//}
//
//func (h *Handler) GetGroupChildrens(c *gin.Context) {
//	groupID := c.Param("groupId")
//	items, err := h.service.GetGroupChildrens(groupID)
//	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
//		return
//	}
//	c.JSON(http.StatusOK, items)
//}
//
//func (h *Handler) GetSubGroupChildrens(c *gin.Context) {
//	subGroupID := c.Param("subGroupId")
//	items, err := h.service.GetGroupByClassID(subGroupID)
//	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
//		return
//	}
//	c.JSON(http.StatusOK, items)
//}
//
//func (h *Handler) GetGroupsBySearch(c *gin.Context) {
//	query := c.Query("query")
//	if query == "" {
//		c.JSON(http.StatusOK, nil)
//	}
//	items, err := h.service.GetGroupsBySearch(query)
//	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
//		return
//	}
//	c.JSON(http.StatusOK, items)
//}
//
//func (h *Handler) GetDiagnosisBySearch(c *gin.Context) {
//	query := c.Query("query")
//	if query == "" {
//		c.JSON(http.StatusOK, nil)
//	}
//	items, err := h.service.GetDiagnosisBySearch(query)
//	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
//		return
//	}
//	c.JSON(http.StatusOK, items)
//}
//
//func (h *Handler) GetSubDiagnosesBySearch(c *gin.Context) {
//	query := c.Query("query")
//	if query == "" {
//		c.JSON(http.StatusOK, nil)
//	}
//	items, err := h.service.GetSubDiagnosesBySearch(query)
//	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
//		return
//	}
//	c.JSON(http.StatusOK, items)
//}
//
//func (h *Handler) GetDiagnosisByGroupID(c *gin.Context) {
//	groupID := c.Param("groupId")
//	items, err := h.service.GetDiagnosisByGroupID(groupID)
//	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
//		return
//	}
//	c.JSON(http.StatusOK, items)
//}
//
//func (h *Handler) GetSubDiagnosisByDiagnosisID(c *gin.Context) {
//	diagnosisID := c.Param("diagnosisId")
//	items, err := h.service.GetSubDiagnosisByDiagnosisID(diagnosisID)
//	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
//		return
//	}
//	c.JSON(http.StatusOK, items)
//}
//
//func (h *Handler) GetConcreteDiagnosisBySubDiagnosisID(c *gin.Context) {
//	diagnosisID := c.Param("subDiagnosisId")
//	items, err := h.service.GetConcreteDiagnosisBySubDiagnosisID(diagnosisID)
//	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
//		return
//	}
//	c.JSON(http.StatusOK, items)
//}
//
//func (h *Handler) GetConcreteDiagnosisBySearch(c *gin.Context) {
//	query := c.Query("query")
//	if query == "" {
//		c.JSON(http.StatusOK, nil)
//	}
//	items, err := h.service.GetConcreteDiagnosisBySearch(query)
//	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
//		return
//	}
//	c.JSON(http.StatusOK, items)
//}
//
//func (h *Handler) SelectMkbElement(c *gin.Context) {
//	id := c.Param("id")
//	class, mkbElement, err := h.service.SelectMkbElement(id)
//	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
//		return
//	}
//	c.JSON(http.StatusOK, map[string]interface{}{"mkbClass": class, "mkbElement": mkbElement})
//}

func (h *Handler) GetTree(c *gin.Context) {
	item, err := h.service.GetTree()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}
func (h *Handler) Get(c *gin.Context) {
	item, err := h.service.Get(c.Param("id"))
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}
