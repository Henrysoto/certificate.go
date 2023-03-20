package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2023-03-21")
	if err != nil {
		t.Errorf("Should be valid, error: %v", err)
	}
	if c == nil {
		t.Errorf("Cert should be valid ref, error: nil")
	}
	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid, expected='GOLANG COURSE', got c.Course=%v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2023-03-21")
	if err == nil {
		t.Errorf("Error should be returned on empty course string")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "asasdhsjdhflauhdsfgliuehgliuhergliheqlrigbnlqkjbgljkabfdlajsdbfljahsdbflasbfjhd"
	_, err := New(course, "Bob", "2023-03-21")
	if err == nil {
		t.Errorf("Error should be returned on a too long name (course=%s)", course)
	}
}
