package schema

type Schema struct {
	Patient              map[string]string `json:"patient"`
	Representative       map[string]string `json:"representative"`
	Human                map[string]string `json:"human"`
	PatientDiagnosis     map[string]string `json:"patientDiagnosis"`
	Disability           map[string]string `json:"disability"`
	MkbGroup             map[string]string `json:"mkbGroup"`
	MkbDiagnosis         map[string]string `json:"mkbDiagnosis"`
	MkbSubDiagnosis      map[string]string `json:"mkbSubDiagnosis"`
	MkbConcreteDiagnosis map[string]string `json:"mkbConcreteDiagnosis"`
	Drug                 map[string]string `json:"drug"`
	RegisterToPatient    map[string]string `json:"registerToPatient"`
}

func CreateSchema() Schema {
	return Schema{
		Patient:              createPatientSchema(),
		Human:                createHumanSchema(),
		PatientDiagnosis:     createPatientDiagnosisSchema(),
		Representative:       createRepresentativeSchema(),
		Disability:           createDisabilitiesSchema(),
		MkbGroup:             createMkbGroupsSchema(),
		MkbDiagnosis:         createMkbDiagnosisSchema(),
		MkbSubDiagnosis:      createMkbSubDiagnosisSchema(),
		MkbConcreteDiagnosis: createMkbConcreteDiagnosisSchema(),
		Drug:                 createDrugSchema(),
		RegisterToPatient:    createRegisterToPatientSchema(),
	}
}

func createHumanSchema() map[string]string {
	return map[string]string{
		"tableName": "humans_view",
		"id":        "id",
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
		"humanId":   "human_id",
		"isMale":    "is_male",
		"dateBirth": "date_birth",
		"createdAt": "created_at",
		"updatedAt": "updated_at",
	}
}

func createRepresentativeSchema() map[string]string {
	return map[string]string{
		"tableName": "representatives_view",
		"key":       "representative",
		"isMale":    "is_male",
		"dateBirth": "date_birth",
		"fullName":  "full_name",
		"createdAt": "created_at",
		"updatedAt": "updated_at",
	}
}

func createPatientDiagnosisSchema() map[string]string {
	return map[string]string{
		"tableName":         "patient_diagnosis",
		"isMale":            "is_male",
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

func createMkbGroupsSchema() map[string]string {
	return map[string]string{
		"tableName": "mkb_groups_view",
		"key":       "mkbGroup",
		"fullName":  "full_name",
		"name":      "name",
	}
}

func createMkbDiagnosisSchema() map[string]string {
	return map[string]string{
		"tableName": "mkb_diagnosis_view",
		"key":       "mkbDiagnosis",
		"fullName":  "full_name",
		"name":      "name",
	}
}

func createMkbSubDiagnosisSchema() map[string]string {
	return map[string]string{
		"tableName": "mkb_sub_diagnosis_view",
		"key":       "mkbSubDiagnosis",
		"fullName":  "full_name",
		"name":      "name",
	}
}

func createMkbConcreteDiagnosisSchema() map[string]string {
	return map[string]string{
		"tableName": "mkb_concrete_diagnosis",
		"key":       "mkbConcreteDiagnosis",
		"name":      "name",
	}
}

func createDrugSchema() map[string]string {
	return map[string]string{
		"tableName": "drugs",
		"key":       "drug",
		"name":      "name",
	}
}

func createRegisterToPatientSchema() map[string]string {
	return map[string]string{
		"tableName":  "register_to_patient",
		"key":        "registerToPatient",
		"registerId": "register_id",
		"patientId":  "patient_id",
	}
}
