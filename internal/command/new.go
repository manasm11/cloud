package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manas/cloud/internal/generator"
	"github.com/manas/cloud/internal/interview"
	"github.com/manas/cloud/internal/interview/questions"
	"github.com/spf13/cobra"
)

// NewNewCommand creates the 'new' command for starting a new project specification.
func NewNewCommand() *cobra.Command {
	var outputDir string
	var projectName string
	var resume bool

	cmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new project specification",
		Long: `Start an interactive interview session to create a comprehensive
project specification (CLAUDE.md) for your new application.

The interview will guide you through:
  - Project basics and goals
  - Application type selection
  - Feature requirements
  - User roles and permissions
  - Data requirements
  - Integration needs
  - Performance expectations
  - Security requirements
  - Deployment preferences
  - Development practices

Upon completion, it generates:
  - CLAUDE.md: Complete project specification
  - Makefile: Build and development commands
  - Dockerfile (optional): Container configuration`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNew(cmd, outputDir, projectName, resume)
		},
	}

	// Flags
	cmd.Flags().StringVarP(&outputDir, "output", "o", ".", "Output directory for generated files")
	cmd.Flags().StringVarP(&projectName, "name", "n", "", "Pre-set project name (will still confirm)")
	cmd.Flags().BoolVarP(&resume, "resume", "r", false, "Resume a previous interrupted session")

	return cmd
}

func runNew(cmd *cobra.Command, outputDir, projectName string, resume bool) error {
	// Resolve output directory
	absOutputDir, err := filepath.Abs(outputDir)
	if err != nil {
		return fmt.Errorf("invalid output directory: %w", err)
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(absOutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create the prompter
	prompter := questions.NewHuhPrompter()

	var interviewer *interview.Interviewer

	// Check for existing session
	if resume && interview.HasExistingSession(absOutputDir) {
		session, err := interview.LoadSession(interview.SessionPath(absOutputDir))
		if err != nil {
			return fmt.Errorf("failed to load existing session: %w", err)
		}
		fmt.Println("📂 Resuming previous session...")
		fmt.Printf("   Current step: %s\n\n", session.CurrentStep.String())
		interviewer = interview.NewInterviewerWithSession(session, prompter, absOutputDir)
	} else if interview.HasExistingSession(absOutputDir) {
		// Ask if they want to resume
		wantResume, err := prompter.Confirm(
			"Existing session found",
			"Would you like to resume the previous session?",
			true,
		)
		if err != nil {
			return err
		}

		if wantResume {
			session, err := interview.LoadSession(interview.SessionPath(absOutputDir))
			if err != nil {
				return fmt.Errorf("failed to load existing session: %w", err)
			}
			fmt.Printf("   Resuming at: %s\n\n", session.CurrentStep.String())
			interviewer = interview.NewInterviewerWithSession(session, prompter, absOutputDir)
		} else {
			// Start fresh
			interviewer = interview.NewInterviewer(prompter, absOutputDir)
		}
	} else {
		// Start fresh
		interviewer = interview.NewInterviewer(prompter, absOutputDir)
	}

	// Print welcome message for new sessions
	if interviewer.Session().CurrentStep == interview.StepBasics {
		printWelcome(cmd)
	}

	// Run the interview
	if err := interviewer.Run(); err != nil {
		return fmt.Errorf("interview failed: %w", err)
	}

	// Interview complete
	fmt.Println("\n✅ Interview complete!")
	fmt.Printf("   Project: %s\n", interviewer.Session().Project.Name)
	fmt.Printf("   Type: %s\n", interviewer.Session().Project.AppType.String())
	fmt.Printf("   Features: %d\n", len(interviewer.Session().Project.Features))

	// Save session
	if err := interviewer.SaveSession(); err != nil {
		fmt.Printf("⚠️  Warning: failed to save session: %v\n", err)
	}

	// Generate files
	fmt.Println("\n📄 Generating project files...")
	fmt.Printf("   Output directory: %s\n", absOutputDir)

	gen := generator.NewFileGenerator()
	if err := gen.GenerateAll(interviewer.Session().Project, absOutputDir); err != nil {
		return fmt.Errorf("failed to generate files: %w", err)
	}

	fmt.Println("   ✓ CLAUDE.md")
	fmt.Println("   ✓ Makefile")
	if interviewer.Session().Project.Deployment.UseDocker {
		fmt.Println("   ✓ Dockerfile")
		fmt.Println("   ✓ docker-compose.yml")
	}

	fmt.Println("\n📝 Next steps:")
	fmt.Println("   1. Review the generated CLAUDE.md")
	fmt.Println("   2. Run 'cloud edit' to modify if needed")
	fmt.Println("   3. Use CLAUDE.md with Claude Code to build your project")

	return nil
}

func printWelcome(cmd *cobra.Command) {
	fmt.Fprintln(cmd.OutOrStdout())
	fmt.Fprintln(cmd.OutOrStdout(), "☁️  Welcome to cloud!")
	fmt.Fprintln(cmd.OutOrStdout())
	fmt.Fprintln(cmd.OutOrStdout(), "I'll help you create a comprehensive project specification by asking")
	fmt.Fprintln(cmd.OutOrStdout(), "questions about your requirements. This specification can then be used")
	fmt.Fprintln(cmd.OutOrStdout(), "with Claude Code to build your application.")
	fmt.Fprintln(cmd.OutOrStdout())
	fmt.Fprintln(cmd.OutOrStdout(), "Philosophy: \"Ask everything, assume nothing, build to last.\"")
	fmt.Fprintln(cmd.OutOrStdout())
	fmt.Fprintln(cmd.OutOrStdout(), "Let's get started!")
	fmt.Fprintln(cmd.OutOrStdout())
}
