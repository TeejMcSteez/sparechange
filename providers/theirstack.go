package providers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"sparechange/types"

	"gopkg.in/yaml.v3"
)

type TheirstackCompany struct {
	AlexaRanking                   int      `json:"alexa_ranking"`
	AnnualRevenueUSD               float64  `json:"annual_revenue_usd"`
	AnnualRevenueUSDReadable       string   `json:"annual_revenue_usd_readable"`
	ApolloID                       string   `json:"apollo_id"`
	City                           string   `json:"city"`
	CompanyKeywords                []string `json:"company_keywords"`
	CompanyTags                    []any    `json:"company_tags"`
	Country                        string   `json:"country"`
	CountryCode                    string   `json:"country_code"`
	Domain                         string   `json:"domain"`
	EmployeeCount                  int      `json:"employee_count"`
	EmployeeCountRange             string   `json:"employee_count_range"`
	FoundedYear                    int      `json:"founded_year"`
	FundingStage                   string   `json:"funding_stage"`
	HasBlurredData                 bool     `json:"has_blurred_data"`
	ID                             string   `json:"id"`
	Industry                       string   `json:"industry"`
	IndustryID                     int      `json:"industry_id"`
	Investors                      []string `json:"investors"`
	IsRecruitingAgency             bool     `json:"is_recruiting_agency"`
	KeywordSlugs                   []string `json:"keyword_slugs"`
	LastFundingRoundAmountReadable string   `json:"last_funding_round_amount_readable"`
	LastFundingRoundDate           string   `json:"last_funding_round_date"`
	LinkedinID                     string   `json:"linkedin_id"`
	LinkedinURL                    string   `json:"linkedin_url"`
	Logo                           string   `json:"logo"`
	LongDescription                string   `json:"long_description"`
	Name                           string   `json:"name"`
	NumBuyingIntentTopics          int      `json:"num_buying_intent_topics"`
	NumJobs                        int      `json:"num_jobs"`
	NumJobsFound                   int      `json:"num_jobs_found"`
	NumJobsLast30Days              int      `json:"num_jobs_last_30_days"`
	NumKeywords                    int      `json:"num_keywords"`
	NumTechnologies                int      `json:"num_technologies"`
	PossibleDomains                []any    `json:"possible_domains"`
	PostalCode                     string   `json:"postal_code"`
	PubliclyTradedExchange         string   `json:"publicly_traded_exchange"`
	PubliclyTradedSymbol           string   `json:"publicly_traded_symbol"`
	SEODescription                 string   `json:"seo_description"`
	TechnologyNames                []string `json:"technology_names"`
	TechnologySlugs                []string `json:"technology_slugs"`
	TotalFundingUSD                float64  `json:"total_funding_usd"`
	URL                            string   `json:"url"`
	URLSource                      string   `json:"url_source"`
	YCBatch                        string   `json:"yc_batch"`
}

type TheirstackHiringTeamMember struct {
	FirstName    string `json:"first_name"`
	FullName     string `json:"full_name"`
	ImageURL     string `json:"image_url"`
	LinkedinURL  string `json:"linkedin_url"`
	Role         string `json:"role"`
	ThumbnailURL string `json:"thumbnail_url"`
}

type TheirstackLocation struct {
	Admin1Code  string  `json:"admin1_code"`
	Admin1Name  string  `json:"admin1_name"`
	Admin2Code  string  `json:"admin2_code"`
	Admin2Name  string  `json:"admin2_name"`
	Continent   string  `json:"continent"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	DisplayName string  `json:"display_name"`
	FeatureCode string  `json:"feature_code"`
	ID          int     `json:"id"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Name        string  `json:"name"`
	State       string  `json:"state"`
	StateCode   string  `json:"state_code"`
	Type        string  `json:"type"`
}

type TheirstackJob struct {
	AvgAnnualSalaryUSD *float64                     `json:"avg_annual_salary_usd"`
	Cities             []string                     `json:"cities"`
	ClosedAt           *string                      `json:"closed_at"`
	Company            string                       `json:"company"`
	CompanyDomain      string                       `json:"company_domain"`
	CompanyObject      TheirstackCompany            `json:"company_object"`
	Continents         []string                     `json:"continents"`
	Countries          []string                     `json:"countries"`
	Country            string                       `json:"country"`
	CountryCode        string                       `json:"country_code"`
	CountryCodes       []string                     `json:"country_codes"`
	DatePosted         string                       `json:"date_posted"`
	DateReposted       string                       `json:"date_reposted"`
	Description        string                       `json:"description"`
	DiscoveredAt       string                       `json:"discovered_at"`
	EasyApply          bool                         `json:"easy_apply"`
	EmploymentStatuses []string                     `json:"employment_statuses"`
	FinalURL           string                       `json:"final_url"`
	HasBlurredData     bool                         `json:"has_blurred_data"`
	HiringTeam         []TheirstackHiringTeamMember `json:"hiring_team"`
	Hybrid             bool                         `json:"hybrid"`
	ID                 int                          `json:"id"`
	JobTitle           string                       `json:"job_title"`
	KeywordSlugs       []string                     `json:"keyword_slugs"`
	Latitude           *float64                     `json:"latitude"`
	Location           string                       `json:"location"`
	Locations          []TheirstackLocation         `json:"locations"`
	LongLocation       string                       `json:"long_location"`
	Longitude          *float64                     `json:"longitude"`
	ManagerRoles       []string                     `json:"manager_roles"`
	MatchingPhrases    []any                        `json:"matching_phrases"`
	MatchingWords      []any                        `json:"matching_words"`
	MaxAnnualSalary    *float64                     `json:"max_annual_salary"`
	MaxAnnualSalaryUSD *float64                     `json:"max_annual_salary_usd"`
	MinAnnualSalary    *float64                     `json:"min_annual_salary"`
	MinAnnualSalaryUSD *float64                     `json:"min_annual_salary_usd"`
	NormalizedTitle    string                       `json:"normalized_title"`
	PostalCode         string                       `json:"postal_code"`
	Remote             bool                         `json:"remote"`
	Reposted           bool                         `json:"reposted"`
	SalaryCurrency     string                       `json:"salary_currency"`
	SalaryString       string                       `json:"salary_string"`
	Seniority          string                       `json:"seniority"`
	ShortLocation      string                       `json:"short_location"`
	SourceURL          string                       `json:"source_url"`
	StateCode          string                       `json:"state_code"`
	TechnologySlugs    []string                     `json:"technology_slugs"`
	URL                string                       `json:"url"`
}

type TheirstackMetadata struct {
	TotalCompanies     int `json:"total_companies"`
	TotalResults       int `json:"total_results"`
	TruncatedCompanies int `json:"truncated_companies"`
	TruncatedResults   int `json:"truncated_results"`
}

type TheirstackResponse struct {
	Data     []TheirstackJob    `json:"data"`
	Metadata TheirstackMetadata `json:"metadata"`
}

type TheirstackConfig struct {
	THEIRSTACK_API_URL        string                   `yaml:"theirstack_api_url"`
	THEIRSTACK_API_HEADERS    map[string]string        `yaml:"theirstack_api_headers"`
	THEIRSTACK_API_URL_PARAMS TheirstackAPIRequestBody `yaml:"theirstack_api_params"`
}

type TheirstackJobLocationFilter struct {
	ID   int    `json:"id,omitempty"   yaml:"id"`
	Type string `json:"type,omitempty" yaml:"type"`
}

type TheirstackColumnSort struct {
	Field string `json:"field"          yaml:"field"`
	Desc  *bool  `json:"desc,omitempty" yaml:"desc"`
}

type TheirstackAPIRequestBody struct {
	BlurCompanyData                              *bool                          `json:"blur_company_data,omitempty"                                yaml:"blur_company_data"`
	ClosedAtGte                                  *string                        `json:"closed_at_gte,omitempty"                                    yaml:"closed_at_gte"`
	ClosedAtLte                                  *string                        `json:"closed_at_lte,omitempty"                                    yaml:"closed_at_lte"`
	CompanyCountryCodeNot                        []string                       `json:"company_country_code_not,omitempty"                         yaml:"company_country_code_not"`
	CompanyCountryCodeOr                         []string                       `json:"company_country_code_or,omitempty"                          yaml:"company_country_code_or"`
	CompanyDescriptionPatternAccentInsensitive   *bool                          `json:"company_description_pattern_accent_insensitive,omitempty"   yaml:"company_description_pattern_accent_insensitive"`
	CompanyDescriptionPatternNot                 []string                       `json:"company_description_pattern_not,omitempty"                  yaml:"company_description_pattern_not"`
	CompanyDescriptionPatternOr                  []string                       `json:"company_description_pattern_or,omitempty"                   yaml:"company_description_pattern_or"`
	CompanyDomainNot                             []string                       `json:"company_domain_not,omitempty"                               yaml:"company_domain_not"`
	CompanyDomainOr                              []string                       `json:"company_domain_or,omitempty"                                yaml:"company_domain_or"`
	CompanyIDNot                                 []string                       `json:"company_id_not,omitempty"                                   yaml:"company_id_not"`
	CompanyIDOr                                  []string                       `json:"company_id_or,omitempty"                                    yaml:"company_id_or"`
	CompanyInvestorsOr                           []string                       `json:"company_investors_or,omitempty"                             yaml:"company_investors_or"`
	CompanyInvestorsPartialMatchOr               []string                       `json:"company_investors_partial_match_or,omitempty"               yaml:"company_investors_partial_match_or"`
	CompanyKeywordSlugAnd                        []string                       `json:"company_keyword_slug_and,omitempty"                         yaml:"company_keyword_slug_and"`
	CompanyKeywordSlugNot                        []string                       `json:"company_keyword_slug_not,omitempty"                         yaml:"company_keyword_slug_not"`
	CompanyKeywordSlugOr                         []string                       `json:"company_keyword_slug_or,omitempty"                          yaml:"company_keyword_slug_or"`
	CompanyLinkedinURLOr                         []string                       `json:"company_linkedin_url_or,omitempty"                          yaml:"company_linkedin_url_or"`
	CompanyListIDNot                             []int                          `json:"company_list_id_not,omitempty"                              yaml:"company_list_id_not"`
	CompanyListIDOr                              []int                          `json:"company_list_id_or,omitempty"                               yaml:"company_list_id_or"`
	CompanyLocationPatternOr                     []string                       `json:"company_location_pattern_or,omitempty"                      yaml:"company_location_pattern_or"`
	CompanyNameCaseInsensitiveOr                 []string                       `json:"company_name_case_insensitive_or,omitempty"                 yaml:"company_name_case_insensitive_or"`
	CompanyNameNot                               []string                       `json:"company_name_not,omitempty"                                 yaml:"company_name_not"`
	CompanyNameOr                                []string                       `json:"company_name_or,omitempty"                                  yaml:"company_name_or"`
	CompanyNamePartialMatchNot                   []string                       `json:"company_name_partial_match_not,omitempty"                   yaml:"company_name_partial_match_not"`
	CompanyNamePartialMatchOr                    []string                       `json:"company_name_partial_match_or,omitempty"                    yaml:"company_name_partial_match_or"`
	CompanyTagsOr                                []string                       `json:"company_tags_or,omitempty"                                  yaml:"company_tags_or"`
	CompanyTechnologySlugAnd                     []string                       `json:"company_technology_slug_and,omitempty"                      yaml:"company_technology_slug_and"`
	CompanyTechnologySlugNot                     []string                       `json:"company_technology_slug_not,omitempty"                      yaml:"company_technology_slug_not"`
	CompanyTechnologySlugOr                      []string                       `json:"company_technology_slug_or,omitempty"                       yaml:"company_technology_slug_or"`
	CompanyType                                  *string                        `json:"company_type,omitempty"                                     yaml:"company_type"`
	Cursor                                       *string                        `json:"cursor,omitempty"                                           yaml:"cursor"`
	DiscoveredAtGte                              *string                        `json:"discovered_at_gte,omitempty"                                yaml:"discovered_at_gte"`
	DiscoveredAtLte                              *string                        `json:"discovered_at_lte,omitempty"                                yaml:"discovered_at_lte"`
	DiscoveredAtMaxAgeDays                       *int                           `json:"discovered_at_max_age_days,omitempty"                       yaml:"discovered_at_max_age_days"`
	DiscoveredAtMinAgeDays                       *int                           `json:"discovered_at_min_age_days,omitempty"                       yaml:"discovered_at_min_age_days"`
	EasyApply                                    *bool                          `json:"easy_apply,omitempty"                                       yaml:"easy_apply"`
	EmploymentStatusesOr                         []string                       `json:"employment_statuses_or,omitempty"                           yaml:"employment_statuses_or"`
	FundingStageOr                               []string                       `json:"funding_stage_or,omitempty"                                 yaml:"funding_stage_or"`
	IncludeTotalResults                          *bool                          `json:"include_total_results,omitempty"                            yaml:"include_total_results"`
	IndustryIDNot                                []int                          `json:"industry_id_not,omitempty"                                  yaml:"industry_id_not"`
	IndustryIDOr                                 []int                          `json:"industry_id_or,omitempty"                                   yaml:"industry_id_or"`
	IsClosed                                     *bool                          `json:"is_closed,omitempty"                                        yaml:"is_closed"`
	JobCountryCodeNot                            []string                       `json:"job_country_code_not,omitempty"                             yaml:"job_country_code_not"`
	JobCountryCodeOr                             []string                       `json:"job_country_code_or,omitempty"                              yaml:"job_country_code_or"`
	JobDescriptionContainsNot                    []string                       `json:"job_description_contains_not,omitempty"                     yaml:"job_description_contains_not"`
	JobDescriptionContainsOr                     []string                       `json:"job_description_contains_or,omitempty"                      yaml:"job_description_contains_or"`
	JobDescriptionPatternAnd                     []string                       `json:"job_description_pattern_and,omitempty"                      yaml:"job_description_pattern_and"`
	JobDescriptionPatternNot                     []string                       `json:"job_description_pattern_not,omitempty"                      yaml:"job_description_pattern_not"`
	JobDescriptionPatternOr                      []string                       `json:"job_description_pattern_or,omitempty"                       yaml:"job_description_pattern_or"`
	JobIDNot                                     []int                          `json:"job_id_not,omitempty"                                       yaml:"job_id_not"`
	JobIDOr                                      []int                          `json:"job_id_or,omitempty"                                        yaml:"job_id_or"`
	JobKeywordSlugAnd                            []string                       `json:"job_keyword_slug_and,omitempty"                             yaml:"job_keyword_slug_and"`
	JobKeywordSlugNot                            []string                       `json:"job_keyword_slug_not,omitempty"                             yaml:"job_keyword_slug_not"`
	JobKeywordSlugOr                             []string                       `json:"job_keyword_slug_or,omitempty"                              yaml:"job_keyword_slug_or"`
	JobLocationNot                               []TheirstackJobLocationFilter   `json:"job_location_not,omitempty"                                 yaml:"job_location_not"`
	JobLocationOr                                []TheirstackJobLocationFilter   `json:"job_location_or,omitempty"                                  yaml:"job_location_or"`
	JobSeniorityOr                               []string                       `json:"job_seniority_or,omitempty"                                 yaml:"job_seniority_or"`
	JobTechnologySlugAnd                         []string                       `json:"job_technology_slug_and,omitempty"                          yaml:"job_technology_slug_and"`
	JobTechnologySlugNot                         []string                       `json:"job_technology_slug_not,omitempty"                          yaml:"job_technology_slug_not"`
	JobTechnologySlugOr                          []string                       `json:"job_technology_slug_or,omitempty"                           yaml:"job_technology_slug_or"`
	JobTitleNot                                  []string                       `json:"job_title_not,omitempty"                                    yaml:"job_title_not"`
	JobTitleOr                                   []string                       `json:"job_title_or,omitempty"                                     yaml:"job_title_or"`
	JobTitlePatternAnd                           []string                       `json:"job_title_pattern_and,omitempty"                            yaml:"job_title_pattern_and"`
	JobTitlePatternNot                           []string                       `json:"job_title_pattern_not,omitempty"                            yaml:"job_title_pattern_not"`
	JobTitlePatternOr                            []string                       `json:"job_title_pattern_or,omitempty"                             yaml:"job_title_pattern_or"`
	LastFundingRoundDateGte                      *string                        `json:"last_funding_round_date_gte,omitempty"                      yaml:"last_funding_round_date_gte"`
	LastFundingRoundDateLte                      *string                        `json:"last_funding_round_date_lte,omitempty"                      yaml:"last_funding_round_date_lte"`
	Limit                                        *int                           `json:"limit,omitempty"                                            yaml:"limit"`
	MaxEmployeeCount                             *int                           `json:"max_employee_count,omitempty"                               yaml:"max_employee_count"`
	MaxEmployeeCountOrNull                       *int                           `json:"max_employee_count_or_null,omitempty"                       yaml:"max_employee_count_or_null"`
	MaxFundingUSD                                *int                           `json:"max_funding_usd,omitempty"                                  yaml:"max_funding_usd"`
	MaxRevenueUSD                                *int                           `json:"max_revenue_usd,omitempty"                                  yaml:"max_revenue_usd"`
	MaxSalaryUSD                                 *float64                       `json:"max_salary_usd,omitempty"                                   yaml:"max_salary_usd"`
	MinEmployeeCount                             *int                           `json:"min_employee_count,omitempty"                               yaml:"min_employee_count"`
	MinEmployeeCountOrNull                       *int                           `json:"min_employee_count_or_null,omitempty"                       yaml:"min_employee_count_or_null"`
	MinFundingUSD                                *int                           `json:"min_funding_usd,omitempty"                                  yaml:"min_funding_usd"`
	MinRevenueUSD                                *int                           `json:"min_revenue_usd,omitempty"                                  yaml:"min_revenue_usd"`
	MinSalaryUSD                                 *float64                       `json:"min_salary_usd,omitempty"                                   yaml:"min_salary_usd"`
	Offset                                       *int                           `json:"offset,omitempty"                                           yaml:"offset"`
	OnlyYCCompanies                              *bool                          `json:"only_yc_companies,omitempty"                                yaml:"only_yc_companies"`
	OrderBy                                      []TheirstackColumnSort         `json:"order_by,omitempty"                                         yaml:"order_by"`
	Page                                         *int                           `json:"page,omitempty"                                             yaml:"page"`
	PostedAtGte                                  *string                        `json:"posted_at_gte,omitempty"                                    yaml:"posted_at_gte"`
	PostedAtLte                                  *string                        `json:"posted_at_lte,omitempty"                                    yaml:"posted_at_lte"`
	PostedAtMaxAgeDays                           *int                           `json:"posted_at_max_age_days,omitempty"                           yaml:"posted_at_max_age_days"`
	PropertyExistsAnd                            []string                       `json:"property_exists_and,omitempty"                              yaml:"property_exists_and"`
	PropertyExistsOr                             []string                       `json:"property_exists_or,omitempty"                               yaml:"property_exists_or"`
	Remote                                       *bool                          `json:"remote,omitempty"                                           yaml:"remote"`
	URLDomainNot                                 []string                       `json:"url_domain_not,omitempty"                                   yaml:"url_domain_not"`
	URLDomainOr                                  []string                       `json:"url_domain_or,omitempty"                                    yaml:"url_domain_or"`
}

func Theirstack(file string) ([]types.ProviderOutput, error) {
	fb, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var conf TheirstackConfig
	if err = yaml.Unmarshal(fb, &conf); err != nil {
		return nil, err
	}
	data, err := fetchTheirstack(conf)
	if err != nil {
		return nil, err
	}
	resp, err := parseTheirstack(data)
	if err != nil {
		return nil, err
	}
	pData := normalizeTheirstack(resp)
	return pData, nil
}

func fetchTheirstack(conf TheirstackConfig) ([]byte, error) {

	jsonData, err := json.Marshal(conf.THEIRSTACK_API_URL_PARAMS)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", conf.THEIRSTACK_API_URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	for key, val := range conf.THEIRSTACK_API_HEADERS {
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

func parseTheirstack(body []byte) (TheirstackResponse, error) {
	var resp TheirstackResponse
	err := json.Unmarshal(body, &resp)
	return resp, err
}

func normalizeTheirstack(t TheirstackResponse) (pOut []types.ProviderOutput) {

	getEmploymentType := func(remote bool, hybrid bool) string {
		if remote == true {
			return "Remote"
		}
		if hybrid == true {
			return "Hybrid"
		}
		return "In Person"
	}

	for _, job := range t.Data {
		var newP types.ProviderOutput
		newP.JobTitle = job.JobTitle
		newP.EmployerName = job.Company
		newP.EmployerWebsite = job.CompanyObject.URL
		newP.JobPublisher = job.SourceURL
		newP.JobEmploymentType = getEmploymentType(job.Remote, job.Hybrid)
		newP.JobApplyLink = &job.CompanyObject.LinkedinURL
		newP.JobApplyIsDirect = &job.EasyApply
		newP.JobDescription = job.Description
		newP.JobPostedAtDatetimeUTC = &job.DatePosted
		newP.JobLocation = job.Location
		newP.JobBenefits = &job.CompanyObject.CompanyKeywords
		pOut = append(pOut, newP)
	}
	return pOut
}
