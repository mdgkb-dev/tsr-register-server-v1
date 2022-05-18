package schema

type Schema struct {
	Patient          map[string]string `json:"patient"`
	Representative   map[string]string `json:"representative"`
	Human            map[string]string `json:"human"`
	PatientDiagnosis map[string]string `json:"patientDiagnosis"`
	Disability       map[string]string `json:"disability"`
}

func CreateSchema() Schema {
	return Schema{
		Patient:          createPatientSchema(),
		Human:            createHumanSchema(),
		PatientDiagnosis: createPatientDiagnosisSchema(),
		Representative:   createRepresentativeSchema(),
		Disability:       createDisabilitiesSchema(),
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

func createPatientSchema() map[string]string {
	return map[string]string{
		"tableName": "patients_view",
		"key":       "patient",
		"fullName":  "full_name",
		"createdAt": "created_at",
		"updatedAt": "updated_at",
	}
}

func createRepresentativeSchema() map[string]string {
	return map[string]string{
		"tableName": "representatives_view",
		"key":       "representative",
		"fullName":  "full_name",
		"createdAt": "created_at",
		"updatedAt": "updated_at",
	}
}

func createPatientDiagnosisSchema() map[string]string {
	return map[string]string{
		"tableName":         "patient_diagnosis",
		"joinTable":         "patients",
		"joinTableFk":       "patient_id",
		"joinTablePk":       "id",
		"mkbDiagnosisId":    "mkb_diagnosis_id",
		"mkbSubDiagnosisId": "mkb_sub_diagnosis_id",
	}
}

func createDisabilitiesSchema() map[string]string {
	return map[string]string{
		"tableName":   "disability",
		"id":          "id",
		"patientId":   "patient_id",
		"joinTable":   "patients",
		"joinTableFk": "patient_id",
		"joinTablePk": "id",
	}
}
