package types

type ProviderParameters struct {
	API_URL        string
	API_KEY        string
	API_HEADERS    []string
	API_URL_PARAMS []string
}

type ProviderOutput struct {
	JobTitle               string
	EmployerName           string
	EmployerWebsite        string
	JobPublisher           string
	JobEmploymentType      string
	JobApplyLink           *string
	JobApplyIsDirect       *bool
	JobDescription         string
	JobPostedAtDatetimeUTC *string
	JobLocation            string
	JobBenefits            *[]string
}
