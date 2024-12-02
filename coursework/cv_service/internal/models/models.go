package models

type Resume struct {
	Id            int
	CandidateName string
	Email         string
	Phone         string
	Experience    string
	Education     string
}

type Vacancy struct {
	Id          int
	Title       string
	Company     string
	Location    string
	Description string
}
