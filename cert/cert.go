package cert

import (
	"fmt"
	"strings"
	"time"
)

var NameMaxLength = 20

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
	n := name
	d := date

	cert := &Cert{
		Course:             c,
		Name:               n,
		LabelTitle:         fmt.Sprintf("%v Certificate - %v", c, n),
		LabelCompletion:    "Certificate of Completion",
		LabelPresented:     "This Certificate is Presented to",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date: %v", d),
	}
	return cert, nil
}

func validateCourse(course string) (string, error) {
	c, err := validateStr(course, NameMaxLength)
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
		return c, fmt.Errorf("Invalid string got=%s, len=%d", c, len(c))
	}
	return c, nil
}
