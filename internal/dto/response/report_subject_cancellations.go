package response

type ReportSubjectCancellations struct {
	SubjectID         uint   `json:"subject_id"`
	SubjectCode       string `json:"subject_code"`
	SubjectName       string `json:"subject_name"`
	Semester          int    `json:"semester"`
	CancellationCount int    `json:"cancellation_count"`
}
