package models

type Member struct {
	Id               int    `json:"id"`
	No               string `json:"no"`
	ProfileImg       string `json:"profile_img"`
	FullName         string `json:"full_name"`
	KanaName         string `json:"kana_name"`
	Motto            string `json:"motto"`
	Biography        string `json:"biography"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	EmploymentStatus *int   `json:"employment_status"`
	Status           int    `json:"status"`
}

type Members struct {
	Members []Member `json:"members"`
}
