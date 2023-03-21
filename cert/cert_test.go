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
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2023-03-21")
	if err == nil {
		t.Errorf("Error should be returned on empty course string")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "asasdhsjdhflauhdsfgliuehgliuhergliheqlrigbnlqkjbgljkabfdl"
	_, err := New(course, "Bob", "2023-03-21")
	if err == nil {
		t.Errorf("Error should be returned on a too long name (course=%s)", course)
	}
}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("Golang", "", "2023-03-21")
	if err == nil {
		t.Errorf("Error should be returned on empty name string")
	}
}

func TestNameTooLong(t *testing.T) {
	name := "asasdhsjdhflauhdsfgliuehgliuhergliheqlrigbnlqkjbgff"
	_, err := New("Golang", name, "2023-03-21")
	if err == nil {
		t.Errorf("Error should be returned on a too long name (name=%s)", name)
	}
}
