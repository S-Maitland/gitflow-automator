package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/s-maitland/gitflow-automator/internal/git"
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
		fmt.Printf("Unknown command: %s\n, command")
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
	if len(args) == 0 {
		fmt.Println("Error: Feature name is required")
		fmt.Println("Usage: gwa feature <name>")
		os.Exit(1)
	}

	branchName := strings.Join(args, "-")
	branchName = strings.ToLower(branchName)
	branchName = strings.ReplaceAll(branchName, " ", "-")
	fullBranchName := "feature/" + branchName

	fmt.Printf("Creating branch: %s\n", fullBranchName)

	if err := git.CreateBranch(fullBranchName); err != nil {
		fmt.Println("Error: Failed to create branch")
		os.Exit(1)
	}

	fmt.Printf("✓ Created and switched to branch: %s\n", fullBranchName)
}

func handleCommit(args []string) {
	fmt.Println("Running commit command...")
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

	fmt.Println("═══════════════════════════════════════")
	fmt.Printf("📍 Current branch: %s\n", currentBranch)
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("\n📝 Status:")
	if status == "" {
		fmt.Println("  ✓ Clean - no changes")
	} else {
		fmt.Println(status)
	}
	fmt.Println("═══════════════════════════════════════")
}