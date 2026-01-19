package questions

import (
	"github.com/manas/cloud/internal/interview/validation"
)

// BasicsResult contains the results from the basics questions.
type BasicsResult struct {
	Name         string
	Description  string
	Problem      string
	TargetUsers  string
	IsNewProject bool
}

// Validate validates the basics result.
func (r *BasicsResult) Validate() error {
	if err := validation.ValidateProjectName(r.Name); err != nil {
		return err
	}
	if err := validation.ValidateNotEmpty(r.Description, "description"); err != nil {
		return err
	}
	return nil
}

// BasicsQuestions handles the project basics interview questions.
type BasicsQuestions struct {
	prompter Prompter
}

// NewBasicsQuestions creates a new BasicsQuestions instance.
func NewBasicsQuestions(prompter Prompter) *BasicsQuestions {
	return &BasicsQuestions{
		prompter: prompter,
	}
}

// Ask runs the basics questions interview.
func (q *BasicsQuestions) Ask() (*BasicsResult, error) {
	result := &BasicsResult{}

	// Project name
	name, err := q.prompter.Text(
		"Project Name",
		"Enter a name for your project (lowercase, numbers, hyphens only)",
		"my-awesome-app",
		validation.ValidateProjectName,
	)
	if err != nil {
		return nil, err
	}
	result.Name = name

	// Project description
	description, err := q.prompter.Text(
		"Project Description",
		"Describe your project in one sentence",
		"A tool that helps developers...",
		func(s string) error {
			return validation.ValidateNotEmpty(s, "description")
		},
	)
	if err != nil {
		return nil, err
	}
	result.Description = description

	// Problem statement
	problem, err := q.prompter.Text(
		"Problem Statement",
		"What problem does this project solve?",
		"Currently, developers struggle with...",
		nil, // Optional
	)
	if err != nil {
		return nil, err
	}
	result.Problem = problem

	// Target users
	targetUsers, err := q.prompter.Text(
		"Target Users",
		"Who are the target users of this application?",
		"Developers, DevOps engineers, etc.",
		nil, // Optional
	)
	if err != nil {
		return nil, err
	}
	result.TargetUsers = targetUsers

	// New or existing project
	isNew, err := q.prompter.Confirm(
		"Is this a new project?",
		"Select 'No' if you're modernizing an existing system",
		true,
	)
	if err != nil {
		return nil, err
	}
	result.IsNewProject = isNew

	return result, nil
}
