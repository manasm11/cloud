package domain

// Step represents an implementation step in the project plan.
type Step struct {
	Number             int      `json:"number" yaml:"number"`
	Phase              int      `json:"phase" yaml:"phase"`
	Title              string   `json:"title" yaml:"title"`
	Objective          string   `json:"objective" yaml:"objective"`
	Tasks              []string `json:"tasks" yaml:"tasks"`
	Tests              []string `json:"tests" yaml:"tests"`
	TDDApproach        string   `json:"tdd_approach" yaml:"tdd_approach"`
	AcceptanceCriteria []string `json:"acceptance_criteria" yaml:"acceptance_criteria"`
	FilesToCreate      []string `json:"files_to_create" yaml:"files_to_create"`
	FilesToModify      []string `json:"files_to_modify" yaml:"files_to_modify"`
	DependsOn          []int    `json:"depends_on" yaml:"depends_on"`
}

// Phase represents a phase of implementation.
type Phase struct {
	Number      int    `json:"number" yaml:"number"`
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Steps       []Step `json:"steps" yaml:"steps"`
}

// NewStep creates a new implementation step.
func NewStep(number, phase int, title, objective string) *Step {
	return &Step{
		Number:             number,
		Phase:              phase,
		Title:              title,
		Objective:          objective,
		Tasks:              make([]string, 0),
		Tests:              make([]string, 0),
		AcceptanceCriteria: make([]string, 0),
		FilesToCreate:      make([]string, 0),
		FilesToModify:      make([]string, 0),
		DependsOn:          make([]int, 0),
	}
}

// AddTask adds a task to the step.
func (s *Step) AddTask(task string) {
	s.Tasks = append(s.Tasks, task)
}

// AddTest adds a test requirement to the step.
func (s *Step) AddTest(test string) {
	s.Tests = append(s.Tests, test)
}

// AddAcceptanceCriteria adds an acceptance criterion to the step.
func (s *Step) AddAcceptanceCriteria(criteria string) {
	s.AcceptanceCriteria = append(s.AcceptanceCriteria, criteria)
}

// AddFile adds a file to create to the step.
func (s *Step) AddFile(path string) {
	s.FilesToCreate = append(s.FilesToCreate, path)
}

// AddDependency adds a dependency on another step.
func (s *Step) AddDependency(stepNumber int) {
	s.DependsOn = append(s.DependsOn, stepNumber)
}

// StepBuilder helps build implementation steps.
type StepBuilder struct {
	currentPhase int
	stepCounter  int
	phases       []Phase
	steps        []Step
}

// NewStepBuilder creates a new step builder.
func NewStepBuilder() *StepBuilder {
	return &StepBuilder{
		currentPhase: 1,
		stepCounter:  0,
		phases:       make([]Phase, 0),
		steps:        make([]Step, 0),
	}
}

// StartPhase begins a new phase.
func (b *StepBuilder) StartPhase(name, description string) {
	b.phases = append(b.phases, Phase{
		Number:      b.currentPhase,
		Name:        name,
		Description: description,
		Steps:       make([]Step, 0),
	})
	b.currentPhase++
}

// AddStep adds a step to the current phase.
func (b *StepBuilder) AddStep(title, objective string) *Step {
	b.stepCounter++
	step := NewStep(b.stepCounter, b.currentPhase-1, title, objective)
	b.steps = append(b.steps, *step)
	return step
}

// Build returns all phases and steps.
func (b *StepBuilder) Build() ([]Phase, []Step) {
	// Associate steps with phases
	for i := range b.phases {
		for _, step := range b.steps {
			if step.Phase == b.phases[i].Number {
				b.phases[i].Steps = append(b.phases[i].Steps, step)
			}
		}
	}
	return b.phases, b.steps
}
