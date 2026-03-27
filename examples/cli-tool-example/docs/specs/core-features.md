# Specification: Report Generator Core Features

## 1. Init Command: Create a New Report Template

### User Workflow
```
$ report-gen init revenue-report
Created template: ~/.report-gen/templates/revenue-report.sql
Edit the file with your SQL query, then run: report-gen run revenue-report --month 2026-03
```

### Behavior
The `init` command creates a new SQL template file for an analyst to customize.

**Input:**
- Template name (e.g., `revenue-report`)
  - Required
  - Must be lowercase letters, numbers, hyphens only (no spaces, no special chars)
  - Must not already exist

**Output:**
- Creates file at `~/.report-gen/templates/{name}.sql`
- File contains a template SQL query with a `{month}` placeholder
- Shows success message with path and next step
- Success exit code (0)

**Error Cases:**
| Situation | Error Message | Exit Code |
|-----------|---------------|-----------|
| Template name already exists | `Template "revenue-report" already exists at ~/.report-gen/templates/revenue-report.sql` | 1 |
| Invalid name (spaces, special chars) | `Template name must be lowercase letters, numbers, and hyphens only. Example: my-report` | 1 |
| Cannot write to ~/.report-gen (permission denied) | `Cannot create template directory: Permission denied at ~/.report-gen/` | 1 |
| No template name provided | `Usage: report-gen init <template-name>` | 1 |

**Example Template Content:**
```sql
-- Report: revenue-report
-- Run with: report-gen run revenue-report --month 2026-03
-- Replace this query with your SQL:

SELECT
  region,
  COUNT(*) as transaction_count,
  SUM(amount) as total_revenue
FROM sales
WHERE DATE_TRUNC('month', created_at) = '{month}'
GROUP BY region
ORDER BY total_revenue DESC
```

### Success Criteria
- ✅ Analyst runs `init`, edits the SQL, and runs a report with no additional setup
- ✅ Template directory is created automatically if it doesn't exist
- ✅ Error messages tell the analyst exactly what's wrong and how to fix it
- ✅ No SQL knowledge required to understand the template

---

## 2. Run Command: Execute a Report

### User Workflow
```
$ report-gen run revenue-report --month 2026-03
Executing query for revenue-report (2026-03)...
✓ Query completed in 8.2 seconds
✓ Results: 12 rows, 3 columns
→ Output: revenue-report_2026-03.csv
```

### Behavior
The `run` command executes a SQL template for a specific month and exports results to CSV.

**Input:**
- Template name (required, positional): name of template to run
- Month (required, flag `--month`): format `YYYY-MM` (e.g., `2026-03`)
  - Validates format
  - Converts to date range for WHERE clause

**Process:**
1. Load template from `~/.report-gen/templates/{name}.sql`
2. Replace `{month}` placeholder with date range: `2026-03-01` to `2026-03-31`
3. Connect to data warehouse
4. Execute query
5. Export results to CSV at `~/{name}_{month}.csv` in current directory

**Output:**
- CSV file in current directory named `{template}_{month}.csv`
- Console shows:
  - Template name and month being executed
  - Progress indicator (optional: "Executing...")
  - Query execution time
  - Result summary (row count, column count)
  - Output file path
  - Success exit code (0)

**Error Cases:**
| Situation | Error Message | Exit Code |
|-----------|---------------|-----------|
| Template doesn't exist | `Template "not-found" not found. Use report-gen list to see available templates.` | 1 |
| Invalid month format | `Invalid month format. Use YYYY-MM (example: 2026-03)` | 1 |
| Month in future | `Cannot query future month: 2026-05 is in the future` | 1 |
| Query times out | `Query timed out after 5 minutes. Check your warehouse or try a simpler query.` | 1 |
| Query has syntax error | `Query error: [actual DW error]. Check your SQL template and try again.` | 1 |
| Warehouse connection fails | `Cannot connect to data warehouse. Check your credentials and network.` | 1 |
| No rows returned | `✓ Query completed in 2.1 seconds. Result: 0 rows (CSV will be empty)` | 0 |
| Cannot write CSV | `Cannot write output file: Permission denied at ./revenue-report_2026-03.csv` | 1 |

**Example CSV Output:**
```
region,transaction_count,total_revenue
North,145,52300
South,98,41200
West,67,28900
```

### Success Criteria
- ✅ Analyst can run a query without knowing SQL syntax for date filtering
- ✅ Errors are clear enough that analyst understands what went wrong
- ✅ Query timeout is long enough for typical queries (5 min) but short enough that analyst doesn't wait forever
- ✅ CSV is directly importable to Excel without cleanup
- ✅ Execution time is acceptable (< 30 seconds for typical reports)

### Performance Requirements
- Typical query: < 30 seconds
- Worst case (large table, complex query): < 5 minutes before timeout
- CSV write: < 5 seconds for typical result sizes (< 100k rows)

---

## 3. Configuration & Setup

### Configuration File
Location: `~/.report-gen/config.yaml`

**Format:**
```yaml
# Data Warehouse Connection
warehouse:
  host: data-warehouse.internal
  port: 5432
  database: analytics
  user: maya@company.com
  # Password stored securely (prompt on first run, don't store in config)

# Output Settings
output:
  directory: ./              # Where to write CSV files
  format: csv               # Only CSV supported in v1
```

### Initial Setup
On first run, if config doesn't exist:
1. Check if `WAREHOUSE_PASSWORD` env var is set
2. If not, prompt user: "Enter data warehouse password:"
3. Store connection (password in memory only, not in config file)
4. Create default config.yaml
5. Continue with command

**Do not:**
- Store password in config file
- Ask for password on every run if password env var is set
- Try to connect without credentials

### Success Criteria
- ✅ Analyst can get started in < 2 minutes (one-time setup)
- ✅ Credentials are not stored in plain text on disk
- ✅ Analyst is not prompted for password on every run if env var is set

---

## 4. List Command: Show Available Templates

### User Workflow
```
$ report-gen list
Available templates:
  revenue-report
  user-engagement
  churn-analysis
```

### Behavior
Shows all templates available in `~/.report-gen/templates/`

**Output:**
- List of template names (alphabetical)
- If no templates exist: `No templates found. Use report-gen init <name> to create one.`
- Exit code: 0

---

## 5. Edge Cases & Limitations

### Handled in This Spec
- ✅ Empty result sets (report success with 0 rows)
- ✅ Malformed SQL (show DW error clearly)
- ✅ Missing templates (show available templates)
- ✅ Invalid month formats (show expected format)
- ✅ Warehouse unavailable (fail fast, don't retry)
- ✅ Query timeout (configurable, default 5 min)

### Explicitly Not Handled (v1)
- ❌ Very large result sets (> 100k rows) — assume analyst writes efficient queries
- ❌ Network retries — fail fast, analyst can retry manually
- ❌ Scheduling/automation — manual runs only
- ❌ Multi-month queries in one run — one month at a time
- ❌ Report versioning — overwrites previous CSV with same month
- ❌ Access control — anyone with this binary can run queries

---

## 6. Open Questions (For User Confirmation)

1. **Date range per month:** Should `2026-03` mean the entire month (03-01 to 03-31) or just the month number?
   - **Answer:** Entire month (03-01 to 03-31 inclusive)

2. **Output location:** Should output CSV go to current directory or a fixed reports directory?
   - **Answer:** Current directory (analyst controls where files go)

3. **Column order in CSV:** Should CSV columns match query SELECT order or be alphabetized?
   - **Answer:** Match query SELECT order (analyst controls it)

4. **Empty result sets:** If a query returns 0 rows, should we still create the CSV file?
   - **Answer:** Yes, create empty CSV (columns but no data rows) so analyst knows the query ran
