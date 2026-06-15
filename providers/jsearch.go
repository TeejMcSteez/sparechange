package providers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"sparechange/types"

	"gopkg.in/yaml.v3"
)

type JSearchApplyOption struct {
	ApplyLink string `json:"apply_link"`
	IsDirect  bool   `json:"is_direct"`
	Publisher string `json:"publisher"`
}

type JSearchJob struct {
	JobID                  string               `json:"job_id"`
	JobTitle               string               `json:"job_title"`
	EmployerName           string               `json:"employer_name"`
	EmployerLogo           *string              `json:"employer_logo"`
	EmployerWebsite        *string              `json:"employer_website"`
	JobPublisher           string               `json:"job_publisher"`
	JobEmploymentType      string               `json:"job_employment_type"`
	JobEmploymentTypes     []string             `json:"job_employment_types"`
	JobApplyLink           string               `json:"job_apply_link"`
	JobApplyIsDirect       bool                 `json:"job_apply_is_direct"`
	ApplyOptions           []JSearchApplyOption `json:"apply_options"`
	JobDescription         string               `json:"job_description"`
	JobIsRemote            bool                 `json:"job_is_remote"`
	JobPostedAt            string               `json:"job_posted_at"`
	JobPostedAtTimestamp   int64                `json:"job_posted_at_timestamp"`
	JobPostedAtDatetimeUTC string               `json:"job_posted_at_datetime_utc"`
	JobLocation            string               `json:"job_location"`
	JobCity                string               `json:"job_city"`
	JobState               string               `json:"job_state"`
	JobCountry             string               `json:"job_country"`
	JobLatitude            float64              `json:"job_latitude"`
	JobLongitude           float64              `json:"job_longitude"`
	JobBenefits            []string             `json:"job_benefits"`
	JobBenefitsStrings     []string             `json:"job_benefits_strings"`
	JobGoogleLink          string               `json:"job_google_link"`
	JobSalary              *float64             `json:"job_salary"`
	JobSalaryString        *string              `json:"job_salary_string"`
	JobMinSalary           *float64             `json:"job_min_salary"`
	JobMaxSalary           *float64             `json:"job_max_salary"`
	JobSalaryPeriod        *string              `json:"job_salary_period"`
	JobHighlights          map[string]any       `json:"job_highlights"`
	JobOnetSoc             *string              `json:"job_onet_soc"`
	JobOnetJobZone         *string              `json:"job_onet_job_zone"`
	EmployerReviews        any                  `json:"employer_reviews"`
}

type JSearchData struct {
	Jobs   []JSearchJob `json:"jobs"`
	Cursor *string      `json:"cursor"`
}

type JSearchParameters struct {
	Query      string `json:"query"`
	NumPages   int    `json:"num_pages"`
	DatePosted string `json:"date_posted"`
	Country    string `json:"country"`
	Language   string `json:"language"`
}

type JSearchResponse struct {
	Status     string            `json:"status"`
	RequestID  string            `json:"request_id"`
	Parameters JSearchParameters `json:"parameters"`
	Data       JSearchData       `json:"data"`
}

type JSearchConfig struct {
	RAPID_API_URL     string            `yaml:"rapid_api_url"`
	RAPID_API_HEADERS map[string]string `yaml:"rapid_api_headers"`
	RAPID_API_PARAMS  *[]string         `yaml:"rapid_api_params"`
}

func JSearch(file string) ([]types.ProviderOutput, error) {
	fb, err := os.ReadFile(file)
	if err != nil {
		return []types.ProviderOutput{}, err
	}
	var conf JSearchConfig
	err = yaml.Unmarshal(fb, &conf)
	if err != nil {
		return []types.ProviderOutput{}, err
	}

	rawJson, err := fetch(conf)
	if err != nil {
		return []types.ProviderOutput{}, err
	}

	res, err := parse(rawJson)
	if err != nil {
		return []types.ProviderOutput{}, err
	}

	data := normalize(res)

	return data, nil
}

func fetch(conf JSearchConfig) ([]byte, error) {
	var sb strings.Builder
	sb.WriteString(conf.RAPID_API_URL)
	if conf.RAPID_API_PARAMS != nil {
		for _, str := range *conf.RAPID_API_PARAMS {
			sb.WriteString(str)
		}
	}
	req, err := http.NewRequest("GET", sb.String(), nil)
	if err != nil {
		return nil, err
	}
	for key, val := range conf.RAPID_API_HEADERS {
		req.Header.Add(key, val)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func parse(body []byte) (JSearchResponse, error) {
	var resp JSearchResponse
	err := json.Unmarshal(body, &resp)
	return resp, err
}

func normalize(j JSearchResponse) (pOut []types.ProviderOutput) {
	for _, job := range j.Data.Jobs {
		var newP types.ProviderOutput
		newP.JobTitle = job.JobTitle
		newP.EmployerName = job.EmployerName
		if job.EmployerWebsite != nil {
			newP.EmployerWebsite = *job.EmployerWebsite
		}
		newP.JobPublisher = job.JobPublisher
		newP.JobEmploymentType = job.JobEmploymentType
		newP.JobApplyLink = &job.JobApplyLink
		newP.JobApplyIsDirect = &job.JobApplyIsDirect
		newP.JobDescription = job.JobDescription
		newP.JobPostedAtDatetimeUTC = &job.JobPostedAtDatetimeUTC
		newP.JobLocation = job.JobLocation
		newP.JobBenefits = &job.JobBenefits
		pOut = append(pOut, newP)
	}
	return pOut
}
