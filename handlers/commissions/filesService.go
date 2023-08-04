package commissions

import (
	"mdgkb/tsr-tegister-server-v1/models"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *FilesService) Upload(c *gin.Context, item *models.Commission, files map[string][]*multipart.FileHeader) (err error) {
	//for i, file := range files {
	//	err := s.helper.Uploader.Upload(c, file, item.SetFilePath(&i))
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}

func (s *FilesService) FillApplicationTemplate(item *models.Commission) ([]byte, error) {
	const point = `✓`
	m := map[string]interface{}{
		"item.PatientDiagnosis.MkbItem.Name": item.PatientDiagnosis.MkbItem.Name,
		"item.Number":                        item.Number,
		//"item.Drug.Name":                     item.Drug.Name,
		"year":                              time.Now().Year(),
		"item.Date":                         item.Date.Format("02.01.2006"),
		"item.FormValue.User.Human.Surname": item.Patient.Human,
	}

	p := make([]string, 0)
	for i, commissionDoctor := range item.CommissionsDoctors {
		//p = append(p, strconv.Itoa(i+1)+". "+point.PointsAchievement.Name)
		name := strconv.Itoa(i+1) + ". " + commissionDoctor.Doctor.Position + " - " + commissionDoctor.Doctor.Name
		p = append(p, name+"\nподпись: ________\n\n")
	}
	m["item.CommissionsDoctors"] = strings.Join(p, "\n")

	return s.helper.Templater.ReplaceDoc(m, "commissionProtocol.docx")
}
