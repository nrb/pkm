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

func init() {
	rootCmd.AddCommand(gatherCmd)
	rootCmd.AddCommand(reportCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runGather(cmd *cobra.Command, args []string) error {
	fmt.Println("Gathering data...")

	execDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	scripts := []string{
		"gh_reviews.sh",
		"list_jira_issues.sh",
		"output_todays_commits.sh",
	}

	for _, script := range scripts {
		scriptPath := filepath.Join(execDir, script)
		fmt.Printf("Running %s...\n", script)

		shellCmd := exec.Command("bash", scriptPath)
		shellCmd.Dir = execDir
		shellCmd.Stdout = os.Stdout
		shellCmd.Stderr = os.Stderr

		if err := shellCmd.Run(); err != nil {
			return fmt.Errorf("failed to run %s: %w", script, err)
		}
	}

	fmt.Println("\nData gathering complete!")
	return nil
}

func runReport(cmd *cobra.Command, args []string) error {
	fmt.Println("Generating report...")

	execDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	var reviews []GitHubReview
	if err := readJSONFile(filepath.Join(execDir, "gh_reviews.json"), &reviews); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not read gh_reviews.json: %v\n", err)
	}

	var issues []JiraIssue
	if err := readJSONFile(filepath.Join(execDir, "jira_issues.json"), &issues); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not read jira_issues.json: %v\n", err)
	}

	var commits []GitCommit
	if err := readJSONFile(filepath.Join(execDir, "todays_commits.json"), &commits); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not read todays_commits.json: %v\n", err)
	}

	printGitHubReviews(reviews)
	printJiraIssues(issues)
	printGitCommits(commits)

	return nil
}

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
