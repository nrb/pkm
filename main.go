package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pkm",
	Short: "Personal Knowledge Management CLI",
	Long:  `A CLI tool to gather and report on GitHub reviews, Jira issues, and git commits.`,
}

var gatherCmd = &cobra.Command{
	Use:   "gather",
	Short: "Gather data from GitHub, Jira, and git",
	Long:  `Runs scripts to gather data from GitHub pull requests, Jira issues, and today's git commits.`,
	RunE:  runGather,
}

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate a report from gathered data",
	Long:  `Reads the gathered JSON data and generates a formatted report.`,
	RunE:  runReport,
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Display environment available to scripts",
	Long:  `Displays environment variables for scripts executed by pkm.`,
	RunE:  runEnv,
}

var config PKMConfig

func init() {
	readOrCreateConfig()
	rootCmd.AddCommand(gatherCmd)
	rootCmd.AddCommand(reportCmd)
	rootCmd.AddCommand(envCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runGather(cmd *cobra.Command, args []string) error {
	fmt.Println("Gathering data...")

	scripts := []string{
		"gh_reviews.sh",
		"list_jira_issues.sh",
		"output_todays_commits.sh",
	}

	for _, script := range scripts {
		scriptPath := filepath.Join(config.ScriptDir(), script)
		fmt.Printf("Running %s...\n", script)

		shellCmd := exec.Command("zsh", scriptPath)
		shellCmd.Dir = config.WorkDir
		shellCmd.Stdout = os.Stdout
		shellCmd.Stderr = os.Stderr
		shellCmd.Env = append(os.Environ(), config.EnvVars()...)

		if err := shellCmd.Run(); err != nil {
			return fmt.Errorf("failed to run %s: %w", script, err)
		}
	}

	fmt.Println("\nData gathering complete!")
	return nil
}

func runReport(cmd *cobra.Command, args []string) error {
	fmt.Println("Generating report...")

	var reviews []GitHubReview
	if err := readJSONFile(filepath.Join(config.DataDir(), "gh_reviews.json"), &reviews); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not read gh_reviews.json: %v\n", err)
	}

	var issues []JiraIssue
	if err := readJSONFile(filepath.Join(config.DataDir(), "jira_issues.json"), &issues); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not read jira_issues.json: %v\n", err)
	}

	var commits []GitCommit
	if err := readJSONFile(filepath.Join(config.DataDir(), "todays_commits.json"), &commits); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not read todays_commits.json: %v\n", err)
	}

	printGitHubReviews(reviews)
	printJiraIssues(issues)
	printGitCommits(commits)

	return nil
}

func runEnv(cmd *cobra.Command, args []string) error {
	for _, line := range config.EnvVars() {
		fmt.Println(line)
	}
	return nil
}

func readOrCreateConfig() error {
	fileName := "config.json"
	workDir := filepath.Join(os.ExpandEnv("$HOME"), ".pkm")
	filePath := filepath.Join(workDir, fileName)

	config.WorkDir = workDir
	config.GitRoot = filepath.Join(os.ExpandEnv("$HOME"), "projects")

	// File contents will overwrite our coded default.
	if err := readJSONFile(filePath, &config); err == nil {
		return nil
	}

	fmt.Fprintf(os.Stderr, "Could not read %s, creating config file.\n", filePath)
	bytes, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, bytes, 0666)
}

// readJSONFIle reads a JSON file within our working directory
func readJSONFile(filename string, v any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func printGitHubReviews(reviews []GitHubReview) {
	fmt.Println("=== GitHub Pull Request Reviews ===")
	if len(reviews) == 0 {
		fmt.Println("No pending reviews")
	} else {
		for i, review := range reviews {
			fmt.Printf("\n%d. #%d - %s\n", i+1, review.Number, review.Title)
			fmt.Printf("   Repository: %s\n", review.Repository.NameWithOwner)
			fmt.Printf("   Updated: %s\n", review.UpdatedAt.Format("2006-01-02 15:04:05"))
		}
	}
	fmt.Println()
}

func printJiraIssues(issues []JiraIssue) {
	fmt.Println("=== Jira Issues ===")
	if len(issues) == 0 {
		fmt.Println("No issues found")
	} else {
		for i, issue := range issues {
			fmt.Printf("\n%d. %s - %s\n", i+1, issue.Key, issue.Fields.Summary)
			fmt.Printf("   Status: %s\n", issue.Fields.Status.Name)
			if issue.Fields.Assignee.DisplayName != "" {
				fmt.Printf("   Assignee: %s\n", issue.Fields.Assignee.DisplayName)
			}
			fmt.Printf("   Type: %s\n", issue.Fields.IssueType.Name)
			if issue.Fields.Priority.Name != "" {
				fmt.Printf("   Priority: %s\n", issue.Fields.Priority.Name)
			}
		}
	}
	fmt.Println()
}

func printGitCommits(commits []GitCommit) {
	fmt.Println("=== Today's Git Commits ===")
	if len(commits) == 0 {
		fmt.Println("No commits found today")
	} else {
		for i, commit := range commits {
			fmt.Printf("\n%d. [%s] %s\n", i+1, commit.AbbreviatedCommit, commit.Subject)
			fmt.Printf("   Branch: %s\n", commit.Branch)
			fmt.Printf("   Author: %s <%s>\n", commit.Author.Name, commit.Author.Email)
			fmt.Printf("   Date: %s\n", commit.Author.Date)
		}
	}
	fmt.Println()
}

// ScriptDir returns the scripts directory from within our working directory.
func (c PKMConfig) ScriptDir() string {
	return filepath.Join(c.WorkDir, "scripts")
}

// DataDir returns the data directory from within our working directory.
func (c PKMConfig) DataDir() string {
	return filepath.Join(c.WorkDir, "data")
}

// EnvVars returns the environment variables that should be available to data scripts.
func (c PKMConfig) EnvVars() []string {
	return []string{
		fmt.Sprintf("PKM_SCRIPT_DIR=%s", c.ScriptDir()),
		fmt.Sprintf("PKM_DATA_DIR=%s", c.DataDir()),
		fmt.Sprintf("PKM_DIR=%s", c.WorkDir),
		fmt.Sprintf("PKM_GIT_ROOT=%s", c.GitRoot),
	}
}
