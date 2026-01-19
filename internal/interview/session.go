// Package interview manages the interactive interview process for gathering project requirements.
package interview

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/manas/cloud/internal/domain"
)

// Step represents a step in the interview process.
type Step int

const (
	StepBasics Step = iota
	StepAppType
	StepFeatures
	StepUserRoles
	StepData
	StepIntegrations
	StepPerformance
	StepSecurity
	StepDeployment
	StepDevPractices
	StepTechStack
	StepComplete
)

// String returns the human-readable name of the step.
func (s Step) String() string {
	names := map[Step]string{
		StepBasics:       "Project Basics",
		StepAppType:      "Application Type",
		StepFeatures:     "Features",
		StepUserRoles:    "User Roles",
		StepData:         "Data Requirements",
		StepIntegrations: "Integrations",
		StepPerformance:  "Performance",
		StepSecurity:     "Security",
		StepDeployment:   "Deployment",
		StepDevPractices: "Development Practices",
		StepTechStack:    "Tech Stack",
		StepComplete:     "Complete",
	}
	if name, ok := names[s]; ok {
		return name
	}
	return "Unknown"
}

// TotalSteps is the number of interview steps (excluding Complete).
const TotalSteps = 11

// Session represents the current state of an interview session.
type Session struct {
	Project     *domain.Project `json:"project"`
	CurrentStep Step            `json:"current_step"`
	IsComplete  bool            `json:"is_complete"`

	// Metadata
	Version   string `json:"version"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Progress represents the progress of the interview.
type Progress struct {
	CurrentStep int
	TotalSteps  int
	Percentage  float64
	StepName    string
}

// NewSession creates a new interview session.
func NewSession() *Session {
	return &Session{
		Project:     domain.NewProject("", ""),
		CurrentStep: StepBasics,
		IsComplete:  false,
		Version:     "1.0",
	}
}

// SetProjectName sets the project name.
func (s *Session) SetProjectName(name string) {
	s.Project.Name = name
}

// SetDescription sets the project description.
func (s *Session) SetDescription(description string) {
	s.Project.Description = description
}

// SetProblem sets the problem statement.
func (s *Session) SetProblem(problem string) {
	s.Project.Problem = problem
}

// SetTargetUsers sets the target users.
func (s *Session) SetTargetUsers(users string) {
	s.Project.TargetUsers = users
}

// SetAppType sets the application type.
func (s *Session) SetAppType(appType domain.AppType) {
	s.Project.AppType = appType
}

// AddFeature adds a feature to the project.
func (s *Session) AddFeature(feature domain.Feature) {
	s.Project.Features = append(s.Project.Features, feature)
}

// RemoveFeature removes a feature by index.
func (s *Session) RemoveFeature(index int) error {
	if index < 0 || index >= len(s.Project.Features) {
		return fmt.Errorf("invalid feature index: %d", index)
	}
	s.Project.Features = append(s.Project.Features[:index], s.Project.Features[index+1:]...)
	return nil
}

// UpdateFeature updates a feature at the given index.
func (s *Session) UpdateFeature(index int, feature domain.Feature) error {
	if index < 0 || index >= len(s.Project.Features) {
		return fmt.Errorf("invalid feature index: %d", index)
	}
	s.Project.Features[index] = feature
	return nil
}

// AddUserRole adds a user role to the project.
func (s *Session) AddUserRole(role domain.UserRole) {
	s.Project.UserRoles = append(s.Project.UserRoles, role)
}

// AddEntity adds an entity to the project.
func (s *Session) AddEntity(entity domain.Entity) {
	s.Project.Entities = append(s.Project.Entities, entity)
}

// AddIntegration adds an integration to the project.
func (s *Session) AddIntegration(integration domain.Integration) {
	s.Project.Integrations = append(s.Project.Integrations, integration)
}

// AdvanceStep moves to the next step in the interview.
func (s *Session) AdvanceStep() {
	if s.CurrentStep < StepComplete {
		s.CurrentStep++
		if s.CurrentStep == StepComplete {
			s.IsComplete = true
		}
	}
}

// GoBack moves to the previous step.
func (s *Session) GoBack() bool {
	if s.CurrentStep > StepBasics {
		s.CurrentStep--
		s.IsComplete = false
		return true
	}
	return false
}

// CanGoBack returns true if we can go back to a previous step.
func (s *Session) CanGoBack() bool {
	return s.CurrentStep > StepBasics
}

// Progress returns the current progress of the interview.
func (s *Session) Progress() Progress {
	percentage := float64(s.CurrentStep) / float64(TotalSteps) * 100
	return Progress{
		CurrentStep: int(s.CurrentStep),
		TotalSteps:  TotalSteps,
		Percentage:  percentage,
		StepName:    s.CurrentStep.String(),
	}
}

// Save saves the session to a file.
func (s *Session) Save(path string) error {
	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Marshal to JSON
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}

	// Write to file
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write session file: %w", err)
	}

	return nil
}

// LoadSession loads a session from a file.
func LoadSession(path string) (*Session, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read session file: %w", err)
	}

	var session Session
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session: %w", err)
	}

	return &session, nil
}

// SessionPath returns the default session file path.
func SessionPath(projectDir string) string {
	return filepath.Join(projectDir, ".cloud", "session.json")
}

// HasExistingSession checks if a session file exists.
func HasExistingSession(projectDir string) bool {
	path := SessionPath(projectDir)
	_, err := os.Stat(path)
	return err == nil
}

// DeleteSession removes the session file.
func DeleteSession(projectDir string) error {
	path := SessionPath(projectDir)
	return os.Remove(path)
}
