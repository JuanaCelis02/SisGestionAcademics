package models

type StudentSubjectRelationship struct {
	StudentID uint     `gorm:"primaryKey" json:"student_id"`
	SubjectID uint     `gorm:"primaryKey" json:"subject_id"`
	Student   *Student `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	Subject   *Subject `gorm:"foreignKey:SubjectID" json:"subject,omitempty"`
}

func (StudentSubjectRelationship) TableName() string {
	return "student_subjects"
}
