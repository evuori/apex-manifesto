package init

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"report-gen/src/errors"
	"report-gen/src/logging"
)

/*
Implements Spec: docs/specs/core-features.md#1-init-command-create-a-new-report-template

This command creates a new SQL template file that analysts can customize.
The spec defines:
- Input validation (lowercase, no spaces, no special chars)
- Template location (~/.report-gen/templates/)
- Error cases (already exists, invalid name, permission denied)
- Output: template with {month} placeholder and helpful comments

All spec requirements are implemented here.
*/

const templateContent = `-- Report: %s
-- Run with: report-gen run %s --month 2026-03
-- Replace this query with your SQL:
-- Placeholder {month} will be replaced with the actual month range (e.g., 2026-03-01 to 2026-03-31)

SELECT
  region,
  COUNT(*) as transaction_count,
  SUM(amount) as total_revenue
FROM sales
WHERE DATE_TRUNC('month', created_at) >= '{month}-01'
  AND DATE_TRUNC('month', created_at) < DATE_TRUNC('month', '{month}-01'::date + interval '1 month')
GROUP BY region
ORDER BY total_revenue DESC
`

// Execute runs the init command with the given template name
func Execute(templateName string) error {
	// Validate template name per spec requirements:
	// - Lowercase letters, numbers, hyphens only
	// - No spaces or special characters
	if err := validateTemplateName(templateName); err != nil {
		return err
	}

	// Get template directory path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.NewError("Cannot determine home directory: " + err.Error())
	}

	templatesDir := filepath.Join(homeDir, ".report-gen", "templates")
	templatePath := filepath.Join(templatesDir, templateName+".sql")

	// Check if template already exists (spec requirement: error if exists)
	if _, err := os.Stat(templatePath); err == nil {
		return errors.AlreadyExistsError(
			fmt.Sprintf(`Template "%s" already exists at %s`, templateName, templatePath),
		)
	}

	// Create templates directory if it doesn't exist
	if err := os.MkdirAll(templatesDir, 0700); err != nil {
		return errors.NewError(
			fmt.Sprintf("Cannot create template directory: Permission denied at %s", templatesDir),
		)
	}

	// Write template file
	content := fmt.Sprintf(templateContent, templateName, templateName)
	if err := os.WriteFile(templatePath, []byte(content), 0600); err != nil {
		return errors.NewError(
			fmt.Sprintf("Cannot write template: %v", err),
		)
	}

	// Log successful creation
	logging.Info(fmt.Sprintf("Created template: %s", templatePath))

	// User-friendly output (per spec: show success message with path and next step)
	fmt.Printf("✓ Created template: %s\n", templatePath)
	fmt.Printf("→ Edit the file with your SQL query, then run:\n")
	fmt.Printf("  report-gen run %s --month 2026-03\n", templateName)

	return nil
}

// validateTemplateName checks that template name meets spec requirements
func validateTemplateName(name string) error {
	if name == "" {
		return errors.UsageError("Usage: report-gen init <template-name>")
	}

	// Spec requirement: lowercase letters, numbers, hyphens only
	// Pattern: ^[a-z0-9-]+$
	pattern := regexp.MustCompile(`^[a-z0-9-]+$`)
	if !pattern.MatchString(name) {
		return errors.ValidationError(
			"Template name must be lowercase letters, numbers, and hyphens only. Example: my-report",
		)
	}

	return nil
}
