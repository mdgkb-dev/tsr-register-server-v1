package answervariants

import (
	"github.com/pro-assistance/pro-assister/helper"
)

type Handler struct {
	helper *helper.Helper
}
type Service struct {
	helper *helper.Helper
}

type Repository struct {
	helper *helper.Helper
}

var H *Handler
var S *Service
var R *Repository

func Init(h *helper.Helper) {
	H = &Handler{helper: h}
	S = &Service{helper: h}
	R = &Repository{helper: h}
}
