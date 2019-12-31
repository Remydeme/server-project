package models

type MedicalInfo struct {
	Privacy      bool     `json:"privacy"`
	Pathology    string   `json:"medicalCondition"`
	MedicalNotes string   `json:"medicalNotes"`
	Allergies    []string `json:"allergies"`
	Medications  []string `json:"medications"`
	BloodType    string   `json:"bloodType"`
	Donor        bool     `json:"donor"`
	Weight       int64    `json:"weight"`
	Height       int64    `json:"weight"`
}
