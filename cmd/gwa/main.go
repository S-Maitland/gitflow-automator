package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/s-maitland/gitflow-automator/internal/git"
	"github.com/s-maitland/gitflow-automator/internal/ui"
	"github.com/s-maitland/gitflow-automator/internal/cli"
	"github.com/s-maitland/gitflow-automator/internal/config"
)

func main(){

	if err := git.CheckIfGitRepo(); err != nil {
		fmt.Println("Error: Not a git repository")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "feature":
		handleFeature(os.Args[2:])
	case "commit":
		handleCommit(os.Args[2:])
	case "status":
		handleStatus(os.Args[2:])
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`Git Workflow Automator (GWA)

	Usage:
	gwa <command> [arguments]

	Commands:
	feature <name>    Create a new feature branch
	commit            Create a conventional commit (interactive)
	status            Show enhanced git status
	help              Show this help message

	Examples:
	gwa feature add-user-login
	gwa commit
	gwa status`)
}

func handleFeature(args []string) {
	flags, args := cli.ParseFlags(args)

	if flags.Help {
		fmt.Println("Usage: gwa feature [options] <name>")
		fmt.Println("\nOptions:")
		fmt.Println("  -v, --verbose    Show detailed output")
		fmt.Println("  -d, --dry-run    Show what would be done")
		os.Exit(0)
	}

	if len(args) == 0 {
		ui.PrintRed("Error: Feature name is required")
		fmt.Println("Usage: gwa feature <name>")
		os.Exit(1)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		ui.PrintRed("Error loading config: %v", err)
		os.Exit(1)
	}

	branchName := strings.Join(args, "-")
	branchName = strings.ToLower(branchName)
	branchName = strings.ReplaceAll(branchName, " ", "-")
	fullBranchName := cfg.BranchPrefixes["feature"] + branchName

	if flags.Verbose {
		fmt.Println("Branch name formatting:")
		fmt.Printf("  Input: %v\n", args)
		fmt.Printf("  Output: %s\n", fullBranchName)
		fmt.Println()
	}

	if flags.DryRun {
		ui.PrintYellow("[DRY RUN] Would create branch: %s", fullBranchName)
		return
	}

	ui.PrintCyan("Creating branch: %s\n", fullBranchName)

	if err := git.CreateBranch(fullBranchName); err != nil {
		fmt.Println("Error: Failed to create branch")
		os.Exit(1)
	}

	ui.PrintGreen("✓ Created and switched to branch: %s\n", fullBranchName)
}

func handleCommit(args []string) {
	staged, err := git.GetStagedFiles()
	if err != nil {
		ui.PrintRed("Error: %v", err)
		os.Exit(1)
	}

	if len(staged) == 0 {
		ui.PrintYellow("⚠ No staged changes to commit")
		fmt.Println("\nStage files first with: git add <file>")
		os.Exit(0)
	}

	ui.PrintCyan("Staged files:")
	for _, file := range staged {
		fmt.Printf("  - %s\n", file)
	}
	fmt.Println()

	commitTypes := []string{
		"feat     - A new feature",
		"fix      - A bug fix",
		"docs     - Documentation changes",
		"style    - Code style changes",
		"refactor - Code refactoring",
		"test     - Adding tests",
		"chore    - Maintenance tasks",
	}

	_, selected, err := ui.PromptSelect("Select commit type: ", commitTypes)
	if err != nil {
		ui.PrintRed("Error: %v", err)
		os.Exit(1)
	}

	commitType := strings.Fields(selected)[0]

	scope, _ := ui.PromptString("Scope (optional, press Enter to skip)")

	message, err := ui.PromptString("Commit message: ")
	if err != nil || message == "" {
		ui.PrintRed("Error: commit message required")
		os.Exit(1)
	}

	var fullMessage strings.Builder
	fullMessage.WriteString(commitType)

	if scope != "" {
		fullMessage.WriteString(fmt.Sprintf("(%s)", scope))
	}

	fullMessage.WriteString(": ")
	fullMessage.WriteString(message)

	fmt.Println()
	ui.PrintCyan("Commit message preview:")
	fmt.Println(fullMessage.String())
	fmt.Println()

	if !ui.PromptYesNo("Create commit?") {
		ui.PrintYellow("✗ Commit cancelled")
		os.Exit(0)
	}

	if err := git.Commit(fullMessage.String()); err != nil {
		ui.PrintRed("Error creating commit: %v", err)
		os.Exit(1)
	}

	ui.PrintGreen("✓ Commit created successfully!")
}

func handleStatus(args []string) {
	currentBranch, err := git.GetCurrentBranch()
	if err != nil {
		fmt.Println("Error: %v\n", err)
		os.Exit(1)
	}

	status, err := git.GetStatus()
	if err != nil {
		fmt.Println("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(ui.Cyan("═══════════════════════════════════════"))
	fmt.Printf("📍 Current branch: %s\n", ui.Green(currentBranch))
	fmt.Println(ui.Cyan("═══════════════════════════════════════"))
	fmt.Println("\n📝 Status:")
	if status == "" {
		ui.PrintGreen("  ✓ Clean - no changes")
	} else {
		fmt.Println(status)
	}
	fmt.Println(ui.Cyan("═══════════════════════════════════════"))
}