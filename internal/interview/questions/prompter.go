// Package questions contains the interview question modules.
package questions

import (
	"github.com/charmbracelet/huh"
)

// Prompter defines the interface for prompting users.
type Prompter interface {
	// Text prompts for text input
	Text(title, description, placeholder string, validate func(string) error) (string, error)

	// Select prompts for single selection
	Select(title, description string, options []Option) (string, error)

	// MultiSelect prompts for multiple selections
	MultiSelect(title, description string, options []Option) ([]string, error)

	// Confirm prompts for yes/no confirmation
	Confirm(title, description string, defaultValue bool) (bool, error)
}

// Option represents a selectable option.
type Option struct {
	Key         string
	Title       string
	Description string
}

// HuhPrompter implements Prompter using charmbracelet/huh.
type HuhPrompter struct {
	theme *huh.Theme
}

// NewHuhPrompter creates a new HuhPrompter.
func NewHuhPrompter() *HuhPrompter {
	return &HuhPrompter{
		theme: huh.ThemeBase(),
	}
}

// Text prompts for text input.
func (p *HuhPrompter) Text(title, description, placeholder string, validate func(string) error) (string, error) {
	var value string

	input := huh.NewInput().
		Title(title).
		Description(description).
		Placeholder(placeholder).
		Value(&value)

	if validate != nil {
		input = input.Validate(validate)
	}

	form := huh.NewForm(huh.NewGroup(input)).WithTheme(p.theme)

	if err := form.Run(); err != nil {
		return "", err
	}

	return value, nil
}

// Select prompts for single selection.
func (p *HuhPrompter) Select(title, description string, options []Option) (string, error) {
	var value string

	huhOptions := make([]huh.Option[string], len(options))
	for i, opt := range options {
		huhOptions[i] = huh.NewOption(opt.Title, opt.Key)
	}

	sel := huh.NewSelect[string]().
		Title(title).
		Description(description).
		Options(huhOptions...).
		Value(&value)

	form := huh.NewForm(huh.NewGroup(sel)).WithTheme(p.theme)

	if err := form.Run(); err != nil {
		return "", err
	}

	return value, nil
}

// MultiSelect prompts for multiple selections.
func (p *HuhPrompter) MultiSelect(title, description string, options []Option) ([]string, error) {
	var values []string

	huhOptions := make([]huh.Option[string], len(options))
	for i, opt := range options {
		huhOptions[i] = huh.NewOption(opt.Title, opt.Key)
	}

	sel := huh.NewMultiSelect[string]().
		Title(title).
		Description(description).
		Options(huhOptions...).
		Value(&values)

	form := huh.NewForm(huh.NewGroup(sel)).WithTheme(p.theme)

	if err := form.Run(); err != nil {
		return nil, err
	}

	return values, nil
}

// Confirm prompts for yes/no confirmation.
func (p *HuhPrompter) Confirm(title, description string, defaultValue bool) (bool, error) {
	var value bool = defaultValue

	confirm := huh.NewConfirm().
		Title(title).
		Description(description).
		Value(&value)

	form := huh.NewForm(huh.NewGroup(confirm)).WithTheme(p.theme)

	if err := form.Run(); err != nil {
		return false, err
	}

	return value, nil
}

// MockPrompter is a test prompter that returns predefined responses.
type MockPrompter struct {
	TextResponses        map[string]string
	SelectResponses      map[string]string
	MultiSelectResponses map[string][]string
	ConfirmResponses     map[string]bool

	// Track calls for verification
	TextCalls        []string
	SelectCalls      []string
	MultiSelectCalls []string
	ConfirmCalls     []string
}

// NewMockPrompter creates a new MockPrompter.
func NewMockPrompter() *MockPrompter {
	return &MockPrompter{
		TextResponses:        make(map[string]string),
		SelectResponses:      make(map[string]string),
		MultiSelectResponses: make(map[string][]string),
		ConfirmResponses:     make(map[string]bool),
		TextCalls:            make([]string, 0),
		SelectCalls:          make([]string, 0),
		MultiSelectCalls:     make([]string, 0),
		ConfirmCalls:         make([]string, 0),
	}
}

// Text returns a predefined text response.
func (m *MockPrompter) Text(title, description, placeholder string, validate func(string) error) (string, error) {
	m.TextCalls = append(m.TextCalls, title)
	if resp, ok := m.TextResponses[title]; ok {
		if validate != nil {
			if err := validate(resp); err != nil {
				return "", err
			}
		}
		return resp, nil
	}
	return "", nil
}

// Select returns a predefined select response.
func (m *MockPrompter) Select(title, description string, options []Option) (string, error) {
	m.SelectCalls = append(m.SelectCalls, title)
	if resp, ok := m.SelectResponses[title]; ok {
		return resp, nil
	}
	if len(options) > 0 {
		return options[0].Key, nil
	}
	return "", nil
}

// MultiSelect returns a predefined multi-select response.
func (m *MockPrompter) MultiSelect(title, description string, options []Option) ([]string, error) {
	m.MultiSelectCalls = append(m.MultiSelectCalls, title)
	if resp, ok := m.MultiSelectResponses[title]; ok {
		return resp, nil
	}
	return []string{}, nil
}

// Confirm returns a predefined confirm response.
func (m *MockPrompter) Confirm(title, description string, defaultValue bool) (bool, error) {
	m.ConfirmCalls = append(m.ConfirmCalls, title)
	if resp, ok := m.ConfirmResponses[title]; ok {
		return resp, nil
	}
	return defaultValue, nil
}
