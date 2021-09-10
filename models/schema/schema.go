package schema

type Schema struct {
	HumanSchema            map[string]string `json:"humanSchema"`
	PatientDiagnosisSchema map[string]string `json:"patientDiagnosisSchema"`
}

func CreateSchema() Schema {
	return Schema{
		HumanSchema:            createHumanSchema(),
		PatientDiagnosisSchema: createPatientDiagnosisSchema(),
	}
}

func createHumanSchema() map[string]string {
	return map[string]string{
		"tableName": "human",
		"dateBirth": "date_birth",
		"fullName":  "full_name",
		"isMale":    "is_male",
	}
}

func createPatientDiagnosisSchema() map[string]string {
	return map[string]string{
		"tableName":   "patient_diagnosis",
		"joinTable":   "patient",
		"joinTableFk": "patient_id",
		"joinTablePk": "id",

		"mkbDiagnosisId":    "mkb_diagnosis_id",
		"mkbSubDiagnosisId": "mkb_sub_diagnosis_id",
	}
}
