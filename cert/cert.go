package cert

import (
	"fmt"
	"strings"
	"time"
)

var CourseMaxLength = 20
var NameMaxLength = 30

type Cert struct {
	Course string
	Name   string
	Date   time.Time

	LabelTitle         string
	LabelCompletion    string
	LabelPresented     string
	LabelParticipation string
	LabelDate          string
}

type Saver interface {
	Save(c Cert) error
}

func New(course, name, date string) (*Cert, error) {
	c, err := validateCourse(course)
	if err != nil {
		return nil, err
	}
	n, err := validateName(name)
	if err != nil {
		return nil, err
	}
	d, err := parseDate(date)
	if err != nil {
		return nil, err
	}

	cert := &Cert{
		Course:             c,
		Name:               n,
		Date:               d,
		LabelTitle:         fmt.Sprintf("%v Certificate - %v", c, n),
		LabelCompletion:    "Certificate of Completion",
		LabelPresented:     "This Certificate is Presented to",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date: %v", d.Format("02/01/2006")),
	}
	return cert, nil
}

func validateCourse(course string) (string, error) {
	c, err := validateStr(course, CourseMaxLength)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(c, " course") {
		c = c + " course"
	}
	return strings.ToTitle(c), nil
}

func validateStr(str string, maxLength int) (string, error) {
	c := strings.TrimSpace(str)
	if len(c) <= 0 || len(c) >= maxLength {
		return c, fmt.Errorf("invalid string got=%s, len=%d", c, len(c))
	}
	return c, nil
}

func validateName(name string) (string, error) {
	n, err := validateStr(name, NameMaxLength)
	if err != nil {
		return "", err
	}
	return strings.ToTitle(n), nil
}

func parseDate(date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, err
	}
	return t, nil
}
