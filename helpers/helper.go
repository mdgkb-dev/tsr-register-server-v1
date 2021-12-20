package helpers

import (
	"mdgkb/tsr-tegister-server-v1/config"
	httpHelper "mdgkb/tsr-tegister-server-v1/helpers/httpHelperV2"
	"mdgkb/tsr-tegister-server-v1/helpers/sqlHelper"
	"mdgkb/tsr-tegister-server-v1/helpers/tokenHelper"
	"mdgkb/tsr-tegister-server-v1/helpers/uploadHelper"

	"mdgkb/tsr-tegister-server-v1/helpers/pdfHelper"
)

type Helper struct {
	HTTP     *httpHelper.HTTPHelper
	PDF      *pdfHelper.PDFHelper
	Uploader uploadHelper.Uploader
	SQL      *sqlHelper.SQLHelper
	Token    *tokenHelper.TokenHelper
}

func NewHelper(config config.Config) *Helper {
	http := httpHelper.NewHTTPHelper()
	pdf := pdfHelper.NewPDFHelper(config)
	sql := sqlHelper.NewSQLHelper()
	uploader := uploadHelper.NewLocalUploader(&config.UploadPath)
	token := tokenHelper.NewTokenHelper()
	return &Helper{HTTP: http, Uploader: uploader, PDF: pdf, SQL: sql, Token: token}
}
