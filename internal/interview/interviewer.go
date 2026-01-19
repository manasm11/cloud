package interview

import (
	"fmt"

	"github.com/manas/cloud/internal/domain"
	"github.com/manas/cloud/internal/interview/questions"
)

// Interviewer orchestrates the interview process.
type Interviewer struct {
	session  *Session
	prompter questions.Prompter
	outputDir string
}

// NewInterviewer creates a new Interviewer.
func NewInterviewer(prompter questions.Prompter, outputDir string) *Interviewer {
	return &Interviewer{
		session:   NewSession(),
		prompter:  prompter,
		outputDir: outputDir,
	}
}

// NewInterviewerWithSession creates an Interviewer with an existing session.
func NewInterviewerWithSession(session *Session, prompter questions.Prompter, outputDir string) *Interviewer {
	return &Interviewer{
		session:   session,
		prompter:  prompter,
		outputDir: outputDir,
	}
}

// Session returns the current session.
func (i *Interviewer) Session() *Session {
	return i.session
}

// Run executes the interview from the current step.
func (i *Interviewer) Run() error {
	for !i.session.IsComplete {
		if err := i.runCurrentStep(); err != nil {
			// Save session on error so user can resume
			if saveErr := i.SaveSession(); saveErr != nil {
				fmt.Printf("Warning: failed to save session: %v\n", saveErr)
			}
			return err
		}
		i.session.AdvanceStep()
	}
	return nil
}

// RunStep executes a single interview step.
func (i *Interviewer) runCurrentStep() error {
	switch i.session.CurrentStep {
	case StepBasics:
		return i.runBasics()
	case StepAppType:
		return i.runAppType()
	case StepFeatures:
		return i.runFeatures()
	case StepUserRoles:
		return i.runUserRoles()
	case StepData:
		return i.runData()
	case StepIntegrations:
		return i.runIntegrations()
	case StepPerformance:
		return i.runPerformance()
	case StepSecurity:
		return i.runSecurity()
	case StepDeployment:
		return i.runDeployment()
	case StepDevPractices:
		return i.runDevPractices()
	case StepTechStack:
		return i.runTechStack()
	}
	return nil
}

func (i *Interviewer) runBasics() error {
	q := questions.NewBasicsQuestions(i.prompter)
	result, err := q.Ask()
	if err != nil {
		return fmt.Errorf("basics questions: %w", err)
	}

	i.session.SetProjectName(result.Name)
	i.session.SetDescription(result.Description)
	i.session.SetProblem(result.Problem)
	i.session.SetTargetUsers(result.TargetUsers)
	i.session.Project.IsNewProject = result.IsNewProject

	return nil
}

func (i *Interviewer) runAppType() error {
	q := questions.NewAppTypeQuestions(i.prompter)
	result, err := q.Ask()
	if err != nil {
		return fmt.Errorf("app type questions: %w", err)
	}

	i.session.SetAppType(result.AppType)
	i.session.Project.AppTypeDetails = result.Details

	return nil
}

func (i *Interviewer) runFeatures() error {
	// For now, a simplified version that asks for features one by one
	fmt.Println("\n📋 Let's list the features your application needs.")
	fmt.Println("   I'll keep asking until you say 'done'.")

	for {
		featureName, err := i.prompter.Text(
			"Feature Name",
			"Enter a feature name (or 'done' to finish)",
			"User Authentication",
			nil,
		)
		if err != nil {
			return err
		}

		if featureName == "done" || featureName == "" {
			break
		}

		description, err := i.prompter.Text(
			"Feature Description",
			"Describe what this feature does",
			"Users can log in and out of the application",
			nil,
		)
		if err != nil {
			return err
		}

		isMVP, err := i.prompter.Confirm(
			"Is this feature required for MVP?",
			"Select 'No' if it can come in a later phase",
			true,
		)
		if err != nil {
			return err
		}

		i.session.AddFeature(domain.Feature{
			Name:        featureName,
			Description: description,
			IsMVP:       isMVP,
		})

		fmt.Printf("   ✓ Added feature: %s\n", featureName)
	}

	fmt.Printf("\n   Total features: %d\n", len(i.session.Project.Features))
	return nil
}

func (i *Interviewer) runUserRoles() error {
	// Simplified user roles collection
	fmt.Println("\n👥 Let's define user roles for your application.")

	// Common role presets
	presetRoles := []questions.Option{
		{Key: "anonymous", Title: "Anonymous Users", Description: "Not logged in visitors"},
		{Key: "user", Title: "Registered Users", Description: "Standard user accounts"},
		{Key: "premium", Title: "Premium Users", Description: "Paid/upgraded accounts"},
		{Key: "admin", Title: "Administrators", Description: "Full system access"},
		{Key: "api", Title: "API Consumers", Description: "External systems"},
		{Key: "custom", Title: "Add Custom Role", Description: "Define your own role"},
	}

	selectedRoles, err := i.prompter.MultiSelect(
		"Select User Roles",
		"Choose all roles that apply to your application",
		presetRoles,
	)
	if err != nil {
		return err
	}

	for _, roleKey := range selectedRoles {
		role := domain.UserRole{Name: roleKey}
		switch roleKey {
		case "anonymous":
			role.Description = "Visitors who have not logged in"
		case "user":
			role.Description = "Registered user with standard access"
		case "premium":
			role.Description = "User with premium/paid features"
		case "admin":
			role.Description = "Administrator with full access"
		case "api":
			role.Description = "External API consumer"
		}
		i.session.AddUserRole(role)
	}

	return nil
}

func (i *Interviewer) runData() error {
	// Simplified data requirements
	fmt.Println("\n💾 Let's understand your data requirements.")

	// Ask about data characteristics
	structure, err := i.prompter.Select(
		"Data Structure",
		"What type of data will you primarily work with?",
		[]questions.Option{
			{Key: "structured", Title: "Structured", Description: "Tables with defined schemas (SQL)"},
			{Key: "semi", Title: "Semi-structured", Description: "JSON/documents with flexible schemas"},
			{Key: "unstructured", Title: "Unstructured", Description: "Files, images, videos"},
			{Key: "mixed", Title: "Mixed", Description: "Combination of the above"},
		},
	)
	if err != nil {
		return err
	}
	i.session.Project.DataCharacteristics.Structure = structure

	// Ask about special requirements
	needsSearch, err := i.prompter.Confirm(
		"Full-text Search",
		"Do you need full-text search capabilities?",
		false,
	)
	if err != nil {
		return err
	}
	i.session.Project.DataCharacteristics.FullTextSearch = needsSearch

	needsAudit, err := i.prompter.Confirm(
		"Audit Trail",
		"Do you need to track changes to data (audit log)?",
		false,
	)
	if err != nil {
		return err
	}
	i.session.Project.DataCharacteristics.AuditTrail = needsAudit

	return nil
}

func (i *Interviewer) runIntegrations() error {
	fmt.Println("\n🔌 Let's identify external integrations.")

	integrationTypes := []questions.Option{
		{Key: "auth", Title: "Authentication Providers", Description: "Google, GitHub, SAML, etc."},
		{Key: "payment", Title: "Payment Processors", Description: "Stripe, PayPal, etc."},
		{Key: "email", Title: "Email Services", Description: "SendGrid, SES, etc."},
		{Key: "storage", Title: "Cloud Storage", Description: "S3, GCS, etc."},
		{Key: "analytics", Title: "Analytics", Description: "GA, Mixpanel, etc."},
		{Key: "monitoring", Title: "Monitoring", Description: "DataDog, Sentry, etc."},
		{Key: "none", Title: "None", Description: "No external integrations needed"},
	}

	selected, err := i.prompter.MultiSelect(
		"Required Integrations",
		"Select all integrations your application needs",
		integrationTypes,
	)
	if err != nil {
		return err
	}

	for _, intType := range selected {
		if intType != "none" {
			i.session.AddIntegration(domain.Integration{
				Type:  intType,
				IsMVP: true,
			})
		}
	}

	return nil
}

func (i *Interviewer) runPerformance() error {
	fmt.Println("\n⚡ Let's understand your performance requirements.")

	users, err := i.prompter.Select(
		"Expected Concurrent Users",
		"How many users will use the application simultaneously?",
		[]questions.Option{
			{Key: "small", Title: "< 100", Description: "Small internal tool"},
			{Key: "medium", Title: "100 - 1,000", Description: "Small to medium business"},
			{Key: "large", Title: "1,000 - 10,000", Description: "Medium scale"},
			{Key: "xlarge", Title: "10,000+", Description: "Large scale"},
		},
	)
	if err != nil {
		return err
	}
	i.session.Project.Performance.ConcurrentUsers = users

	responseTime, err := i.prompter.Select(
		"Response Time Requirements",
		"What's the acceptable response time for typical operations?",
		[]questions.Option{
			{Key: "realtime", Title: "< 100ms", Description: "Real-time critical"},
			{Key: "interactive", Title: "< 500ms", Description: "Interactive applications"},
			{Key: "standard", Title: "< 2 seconds", Description: "Standard web applications"},
			{Key: "background", Title: "> 2 seconds OK", Description: "Background processing acceptable"},
		},
	)
	if err != nil {
		return err
	}
	i.session.Project.Performance.ResponseTime = responseTime

	return nil
}

func (i *Interviewer) runSecurity() error {
	fmt.Println("\n🔒 Let's define security requirements.")

	authMethods := []questions.Option{
		{Key: "password", Title: "Username/Password", Description: "Traditional login"},
		{Key: "social", Title: "Social Login", Description: "Google, GitHub, etc."},
		{Key: "sso", Title: "SSO/SAML", Description: "Enterprise single sign-on"},
		{Key: "apikey", Title: "API Keys", Description: "For API access"},
		{Key: "mfa", Title: "Multi-Factor Auth", Description: "Additional security layer"},
	}

	selectedAuth, err := i.prompter.MultiSelect(
		"Authentication Methods",
		"Select all authentication methods needed",
		authMethods,
	)
	if err != nil {
		return err
	}
	i.session.Project.Security.AuthMethods = selectedAuth

	sensitivity, err := i.prompter.Select(
		"Data Sensitivity",
		"What's the sensitivity level of your data?",
		[]questions.Option{
			{Key: "public", Title: "Public", Description: "No sensitive data"},
			{Key: "internal", Title: "Internal", Description: "Private but not regulated"},
			{Key: "pii", Title: "PII", Description: "Personally identifiable information"},
			{Key: "financial", Title: "Financial", Description: "Payment or financial data"},
			{Key: "health", Title: "Health", Description: "Medical/health information (HIPAA)"},
		},
	)
	if err != nil {
		return err
	}
	i.session.Project.Security.DataSensitivity = sensitivity

	return nil
}

func (i *Interviewer) runDeployment() error {
	fmt.Println("\n🚀 Let's configure deployment settings.")

	hosting, err := i.prompter.Select(
		"Hosting Preference",
		"Where will your application be hosted?",
		[]questions.Option{
			{Key: "aws", Title: "AWS", Description: "Amazon Web Services"},
			{Key: "gcp", Title: "Google Cloud", Description: "Google Cloud Platform"},
			{Key: "azure", Title: "Azure", Description: "Microsoft Azure"},
			{Key: "paas", Title: "PaaS", Description: "Heroku, Railway, Render, Fly.io"},
			{Key: "selfhost", Title: "Self-Hosted", Description: "On-premise servers"},
			{Key: "undecided", Title: "Not Sure", Description: "Help me decide later"},
		},
	)
	if err != nil {
		return err
	}
	i.session.Project.Deployment.CloudProvider = hosting

	useDocker, err := i.prompter.Confirm(
		"Use Docker?",
		"Docker ensures consistent deployments across environments",
		true,
	)
	if err != nil {
		return err
	}
	i.session.Project.Deployment.UseDocker = useDocker

	return nil
}

func (i *Interviewer) runDevPractices() error {
	fmt.Println("\n🛠️ Let's establish development practices.")

	repoHost, err := i.prompter.Select(
		"Repository Host",
		"Where will your code be hosted?",
		[]questions.Option{
			{Key: "github", Title: "GitHub", Description: "GitHub.com"},
			{Key: "gitlab", Title: "GitLab", Description: "GitLab.com or self-hosted"},
			{Key: "bitbucket", Title: "Bitbucket", Description: "Atlassian Bitbucket"},
			{Key: "other", Title: "Other", Description: "Different provider"},
		},
	)
	if err != nil {
		return err
	}
	i.session.Project.DevPractices.RepoHost = repoHost

	testTypes := []questions.Option{
		{Key: "unit", Title: "Unit Tests", Description: "Test individual functions"},
		{Key: "integration", Title: "Integration Tests", Description: "Test component interactions"},
		{Key: "e2e", Title: "End-to-End Tests", Description: "Test full user flows"},
		{Key: "performance", Title: "Performance Tests", Description: "Test under load"},
	}

	selectedTests, err := i.prompter.MultiSelect(
		"Testing Strategy",
		"Select the types of tests you want",
		testTypes,
	)
	if err != nil {
		return err
	}
	i.session.Project.DevPractices.TestingTypes = selectedTests

	return nil
}

func (i *Interviewer) runTechStack() error {
	fmt.Println("\n🔧 Based on your requirements, let me suggest a tech stack.")
	fmt.Println("   (Tech stack suggestions will be implemented in Phase 3)")

	// For now, just confirm we've captured everything
	confirm, err := i.prompter.Confirm(
		"Review Complete",
		"Have we captured all your requirements?",
		true,
	)
	if err != nil {
		return err
	}

	if !confirm {
		fmt.Println("   You can use 'cloud edit' later to modify requirements.")
	}

	return nil
}

// SaveSession saves the current session to disk.
func (i *Interviewer) SaveSession() error {
	path := SessionPath(i.outputDir)
	return i.session.Save(path)
}
