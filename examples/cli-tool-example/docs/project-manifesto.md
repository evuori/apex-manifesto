# Project Manifesto: Report Generator

## The User

**Maya, Data Analyst** — She knows SQL, Excel, and business logic. She doesn't know Python, DevOps, or how to deploy code. She has direct access to our data warehouse.

She generates monthly reports (revenue by region, user engagement by cohort, churn analysis) for stakeholders. Each report requires:
1. Writing a SQL query to pull the right data
2. Running the query manually in a query tool
3. Downloading as CSV
4. Opening in Excel
5. Adding pivot tables, charts, formatting

This takes **2-4 hours per report** and requires waiting for someone technical if she has questions.

---

## The Problem

Reports take too long and require engineering help. Maya should be able to generate reports on her own schedule without waiting for a developer.

**Success looks like:** Maya can write a SQL template once, then generate reports for any month in < 5 minutes.

---

## What We're Building

A command-line tool that lets Maya:
1. Create a new report template (once)
2. Run the report for any month without touching code
3. Export results directly to CSV for Excel

**Core workflow:**
```
$ report-gen init revenue-report
# Maya writes SQL in the generated template file
$ report-gen run revenue-report --month 2026-03
# Output: report_2026-03.csv (ready to open in Excel)
```

### In Scope
- ✅ SQL-based templates (Maya's expertise)
- ✅ Date-parameterized queries (monthly reporting)
- ✅ CSV export (directly importable to Excel)
- ✅ Error messages clear enough that Maya understands them
- ✅ Runs on Maya's laptop (no server required)

### Explicitly Out of Scope
- ❌ UI/web interface (Maya is comfortable with CLI)
- ❌ Scheduling or automation (she runs reports on demand)
- ❌ Real-time dashboards (these are static reports)
- ❌ Multi-user features (this is for individual analysts)
- ❌ Permission/access management (assumes Maya has warehouse access)
- ❌ Complex transformations in the tool (SQL handles this)

---

## Constraints & Assumptions

### Constraints
- **Runs locally:** Single binary, deployable to Maya's laptop
- **No database backend:** Configuration in local YAML files
- **No authentication:** Assumes warehouse credentials already configured
- **Performance:** Typical reports complete in < 30 seconds

### Assumptions
- Maya has direct access to the data warehouse
- Maya can write SQL
- Report queries return < 100k rows (no pagination needed)
- Reports are run manually, not scheduled
- Analysts use Excel or similar tools to consume CSV

---

## Team & Ownership

| Component | Owner | Role |
|-----------|-------|------|
| CLI Commands (init, run, config) | Alex | Design, code, testing |
| Data Warehouse Connection | Jake | Integration with DW, query runner |
| CSV Export | Morgan | Format validation, Excel compatibility |
| User Communication | Maya (user) | Feedback, testing, requirements |

---

## Success Criteria

When Report Generator is complete:
1. ✅ A new analyst can create and run their first report in < 15 minutes with only `report-gen --help`
2. ✅ Reports execute in < 30 seconds (typical case)
3. ✅ CSV exports are directly importable to Excel without cleanup
4. ✅ Error messages are clear enough that non-technical users understand what went wrong
5. ✅ Maya has generated at least one real report and confirmed it saves her time

---

## Open Questions (resolved before code starts)

- **Q: Should analysts be able to schedule recurring reports?**
  - A: No, out of scope for v1. Maya runs reports on demand. If automation is needed later, that's a new project.

- **Q: What if the query times out or the DW is unavailable?**
  - A: Show a clear error. Don't retry silently (user should know if DW is down). Timeout < 5 minutes and fail fast.

- **Q: Should we store previous report outputs?**
  - A: No. Each run generates a new CSV with the current timestamp. Maya manages her own archives.

- **Q: How do analysts share reports?**
  - A: Each analyst has their own templates. Sharing is out of scope. They can share CSV files manually.

---

## Non-Requirements

We are **not** building:
- A data warehouse
- A UI or web interface
- A scheduling system
- A data transformation engine
- Access control or multi-tenancy
- A dashboard

This is a thin CLI wrapper around SQL that Maya already knows how to write.

---

## Decision Authority

**Alex** (product owner) owns final decisions on scope and features.
**Jake** owns decisions on data warehouse integration.
**Morgan** owns decisions on export formats.

If scope changes after this manifesto is signed off, we discuss and update this document, then note the change in the decision log.

---

## Definition of Done

The project is done when:
1. All success criteria above are met
2. Maya has tested with real data and confirmed it works
3. Deployment instructions are clear for other analysts
4. Known limitations are documented

Not done means: features work but aren't tested at scale, or error handling isn't production-ready.

---

**Approved by:**
- Product: Alex (date: 2026-03-01)
- Engineer Lead: Jake (date: 2026-03-01)
- End User: Maya (date: 2026-03-01)
