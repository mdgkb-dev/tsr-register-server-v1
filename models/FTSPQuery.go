package models

import (
	"github.com/pro-assistance/pro-assister/helpers/sql"
)

type FTSPQuery struct {
	QID string `json:"qid"`
}
type FTSPAnswer struct {
	Data interface{} `json:"data"`
	FTSP sql.FTSP    `json:"ftsp"`
}
