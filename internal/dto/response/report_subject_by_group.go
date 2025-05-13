package response

type ReportSubjectCancellationsByGroup struct {
	SubjectID         uint     `json:"subject_id"`
	SubjectCode       string   `json:"subject_code"`
	SubjectName       string   `json:"subject_name"`
	GroupCancellations []struct {
			Group            string `json:"group"`
			CancellationCount int    `json:"cancellation_count"`
	} `json:"group_cancellations"`
}