package models

type SubjectGroupStudent struct {
	SubjectID uint `gorm:"primaryKey;index" json:"subject_id"`
	GroupNum  int  `gorm:"primaryKey" json:"group_num"`
	StudentID uint `gorm:"primaryKey;index" json:"student_id"`
}

func (SubjectGroupStudent) TableName() string {
	return "subject_group_students"
}
