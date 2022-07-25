package fileinfo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Download(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(&id)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	fullPath := h.filesService.GetFullPath(&item.FileSystemPath)
	c.Header("Content-Description", "File Transfer")
	c.Header("Download-File-Name", item.OriginalName)
	c.File(*fullPath)
}
