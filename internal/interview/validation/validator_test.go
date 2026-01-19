package validation

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateProjectName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{"valid simple", "myapp", false, ""},
		{"valid with hyphen", "my-app", false, ""},
		{"valid with numbers", "app123", false, ""},
		{"valid with number and hyphen", "app-123", false, ""},
		{"valid min length", "abc", false, ""},
		{"invalid empty", "", true, "required"},
		{"invalid spaces", "my app", true, "lowercase letters, numbers, and hyphens"},
		{"invalid uppercase", "MyApp", true, "lowercase letters, numbers, and hyphens"},
		{"invalid special chars", "my@app", true, "lowercase letters, numbers, and hyphens"},
		{"invalid underscore", "my_app", true, "lowercase letters, numbers, and hyphens"},
		{"too short", "ab", true, "at least 3 characters"},
		{"too long", strings.Repeat("a", 65), true, "at most 64 characters"},
		{"starts with hyphen", "-myapp", true, "cannot start or end with a hyphen"},
		{"ends with hyphen", "myapp-", true, "cannot start or end with a hyphen"},
		{"consecutive hyphens", "my--app", true, "lowercase letters, numbers, and hyphens"},
		{"starts with number", "123app", false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateProjectName(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateNotEmpty(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		fieldName string
		wantErr   bool
	}{
		{"non-empty string", "hello", "description", false},
		{"whitespace only", "   ", "description", true},
		{"empty string", "", "description", true},
		{"tabs only", "\t\t", "description", true},
		{"valid with spaces", "hello world", "description", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateNotEmpty(tt.input, tt.fieldName)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.fieldName)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateMinLength(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		minLen    int
		fieldName string
		wantErr   bool
	}{
		{"meets minimum", "hello", 5, "description", false},
		{"exceeds minimum", "hello world", 5, "description", false},
		{"below minimum", "hi", 5, "description", true},
		{"empty string", "", 1, "description", true},
		{"exact minimum", "abc", 3, "name", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMinLength(tt.input, tt.minLen, tt.fieldName)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.fieldName)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateMaxLength(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		maxLen    int
		fieldName string
		wantErr   bool
	}{
		{"under maximum", "hello", 10, "name", false},
		{"at maximum", "hello", 5, "name", false},
		{"over maximum", "hello world", 5, "name", true},
		{"empty string", "", 10, "name", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMaxLength(tt.input, tt.maxLen, tt.fieldName)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.fieldName)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateSelection(t *testing.T) {
	options := []string{"option1", "option2", "option3"}

	tests := []struct {
		name      string
		input     string
		options   []string
		fieldName string
		wantErr   bool
	}{
		{"valid option", "option1", options, "choice", false},
		{"valid second option", "option2", options, "choice", false},
		{"invalid option", "option4", options, "choice", true},
		{"empty input", "", options, "choice", true},
		{"case sensitive", "Option1", options, "choice", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSelection(tt.input, tt.options, tt.fieldName)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.fieldName)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateURL(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid http", "http://example.com", false},
		{"valid https", "https://example.com", false},
		{"valid with path", "https://example.com/path/to/resource", false},
		{"valid with port", "https://example.com:8080", false},
		{"invalid no scheme", "example.com", true},
		{"invalid empty", "", true},
		{"invalid scheme", "ftp://example.com", true},
		{"valid localhost", "http://localhost:3000", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateURL(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid email", "user@example.com", false},
		{"valid with plus", "user+tag@example.com", false},
		{"valid with subdomain", "user@sub.example.com", false},
		{"invalid no at", "userexample.com", true},
		{"invalid no domain", "user@", true},
		{"invalid empty", "", true},
		{"invalid no local part", "@example.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
