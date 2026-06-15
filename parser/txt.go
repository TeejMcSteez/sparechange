package parser

import (
	"fmt"
	"os"
	"sparechange/types"
	"strings"
)

func WriteText(filename string, data []types.ProviderOutput) error {
	var sb strings.Builder
	for _, entry := range data {
		fmt.Fprintf(&sb, "Title: %s\n", entry.JobTitle)
		fmt.Fprintf(&sb, "Employer: %s\n", entry.EmployerName)
		fmt.Fprintf(&sb, "Location: %s\n", entry.JobLocation)
		fmt.Fprintf(&sb, "Type: %s\n", entry.JobEmploymentType)
		fmt.Fprintf(&sb, "Publisher: %s\n", entry.JobPublisher)
		if entry.JobApplyLink != nil {
			fmt.Fprintf(&sb, "Apply: %s\n", *entry.JobApplyLink)
		}
		if entry.JobPostedAtDatetimeUTC != nil {
			fmt.Fprintf(&sb, "Posted: %s\n", *entry.JobPostedAtDatetimeUTC)
		}
		if entry.JobBenefits != nil {
			fmt.Fprintf(&sb, "Benefits: %s\n", strings.Join(*entry.JobBenefits, ", "))
		}
		fmt.Fprintf(&sb, "Description:\n%s\n", entry.JobDescription)
		sb.WriteString("\n---\n\n")
	}
	return os.WriteFile(filename, []byte(sb.String()), 0644)
}
