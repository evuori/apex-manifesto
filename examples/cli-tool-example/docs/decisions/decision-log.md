# Decision Log: Report Generator

This log tracks significant architecture, design, and tradeoff decisions. Each entry includes:
- What decision was made
- Why (context and constraints)
- Alternatives considered
- Owner (who decided)
- Date

---

## Language: Go

**Decision:** Implement Report Generator in Go

**Owner:** Jake (Technical Lead)

**Date:** 2026-02-15

**Context:**
- Need to distribute as single binary (single command, no Python/Node runtime required)
- Must run on analysts' laptops (macOS, Linux, Windows)
- Performance matters for typical queries (8-12 seconds execution)

**Alternatives Considered:**
1. **Python:** Easy to write, but requires Python runtime. Harder to distribute single binary.
2. **Rust:** Excellent binary story, but steep learning curve. Go is sufficient and faster to deliver.
3. **Node/TypeScript:** Better for UI/web, not ideal for CLI. Same distribution challenge as Python.

**Decision:** Go
- ✅ Compiles to single, stateless binary
- ✅ Cross-platform (macOS, Linux, Windows) with minimal effort
- ✅ Fast execution, good for query runner
- ✅ Standard library covers everything we need (CLI, HTTP, CSV)
- ✅ Team has Go expertise

---

## Distribution: Single Binary

**Decision:** Ship as compiled binary (no source code distribution, no build from source)

**Owner:** Alex (Product)

**Date:** 2026-02-15

**Context:**
- Users are analysts, not developers. They should not need to install Go or build from source.
- Single binary is easiest to install: download, make executable, run.
- Alternative: Docker container, but requires Docker and is overkill for a CLI tool.

**Alternatives Considered:**
1. **Docker container:** Would ensure consistency, but requires Docker installation and adds complexity.
2. **Shell script wrapper around Python:** Still requires Python runtime and dependencies.
3. **Homebrew/apt packages:** Good long-term, but more distribution infrastructure.

**Decision:** Single binary, distributed via GitHub releases
- ✅ Download one file, no installation required
- ✅ No dependencies to install
- ✅ Works on any macOS, Linux, Windows machine with no setup
- **Future:** Can add Homebrew formula later if adoption grows

---

## Configuration Storage: YAML in Home Directory

**Decision:** Store config in `~/.report-gen/config.yaml` (YAML format, home directory)

**Owner:** Jake

**Date:** 2026-02-20

**Context:**
- Need to store data warehouse connection details (host, port, database, user)
- Should work on personal laptops, no central config server
- Users may have multiple reports, all sharing same warehouse connection
- Analysts are comfortable with editing text files (they edit SQL)

**Alternatives Considered:**
1. **Environment variables only:** Cleaner, but hard for users to manage multiple vars. Doesn't scale to structured config.
2. **JSON:** Fine, but YAML is more readable and forgiving of user edits.
3. **In-memory only:** Would need re-entry on every run. Bad UX.
4. **Encrypted file:** Safer, but adds complexity. v1 assumes local machine is already secure.

**Decision:** `~/.report-gen/config.yaml`
- ✅ Standard location (home directory)
- ✅ Human-readable format (analysts understand YAML)
- ✅ Shared across all reports
- ✅ Easy to backup or migrate

**Note:** Password stored in memory only, not in config file. Can be passed via env var (`WAREHOUSE_PASSWORD`) to avoid interactive prompt.

---

## Password Handling: Environment Variable or Interactive Prompt

**Decision:** Password via `WAREHOUSE_PASSWORD` env var or interactive prompt if not set

**Owner:** Jake

**Date:** 2026-02-20

**Context:**
- Cannot store password in plaintext config file (security)
- Cannot use username/password in config and require env var on every run (bad UX)
- Some users may want to automate (CI/scripts), some may want interactive

**Alternatives Considered:**
1. **Always interactive prompt:** Safest, but annoying for scripting or frequent use.
2. **Only env var:** Requires users to know about env vars and set them correctly.
3. **OAuth/tokens:** More secure long-term, but infrastructure overhead for v1.
4. **Keychain/credential storage:** Secure, but platform-specific (macOS Keychain, Windows Credential Manager, Linux Secret Service).

**Decision:** Env var first, interactive prompt as fallback
- Supports automated workflows and scripts (env var)
- Supports one-off runs without setup (interactive prompt)
- Secure: password never written to disk
- **Future:** Can upgrade to system keychain if adoption grows

---

## Date Format: YYYY-MM (Month Level Only)

**Decision:** Support month-level granularity only. Date format: `YYYY-MM` (e.g., `2026-03`).

**Owner:** Alex (with Maya's input)

**Date:** 2026-02-25

**Context:**
- Maya's reports are monthly (revenue by region, monthly cohort analysis, etc.)
- Daily or hourly granularity not needed for v1
- Simpler implementation: fewer edge cases (leap years, day boundaries)
- Analyst can always run multiple months separately if needed

**Alternatives Considered:**
1. **ISO 8601 date format (YYYY-MM-DD):** Flexible, but overly complex for monthly reports.
2. **Week format (YYYY-W##):** Less intuitive for business analysts.
3. **Quarter format (2026-Q1):** Could support, but adds complexity for little benefit.
4. **Date range (2026-03-01:2026-03-31):** Verbose and confusing.

**Decision:** YYYY-MM (month only)
- ✅ Matches Maya's reporting cadence
- ✅ Clear and unambiguous (no confusion about day boundaries)
- ✅ Simpler validation and implementation
- **Future:** Can add daily granularity if use case emerges

---

## Error Messages: User-Friendly, Non-Technical

**Decision:** Error messages should be actionable and non-technical. No stack traces.

**Owner:** Alex (with Maya's input)

**Date:** 2026-02-28

**Context:**
- Users are analysts, not engineers
- Stack traces are noise and scary
- Good errors explain what went wrong and how to fix it
- Bad errors: "java.sql.SQLException: Connection timeout"
- Good errors: "Cannot connect to data warehouse. Check your credentials and network."

**Alternatives Considered:**
1. **Show stack traces:** Gives full details, but intimidates non-technical users.
2. **Generic errors:** Safe, but unhelpful ("An error occurred").
3. **Logging to file:** Can log details while showing clean message to user.

**Decision:** User-friendly messages with logging
- Show user: Simple, actionable message
- Log to file (`~/.report-gen/report-gen.log`): Full details including stack trace
- If user reports an issue, they can provide the log file
- **Benefits:** Users never see scary stack traces, but we have details for debugging

---

## CSV Export: No Quoting Unless Necessary

**Decision:** CSV export uses minimal quoting (only quote fields that contain comma, newline, or quote char).

**Owner:** Morgan

**Date:** 2026-03-01

**Context:**
- CSV must be directly importable to Excel
- Excel handles both quoted and unquoted fields
- Over-quoting (every field) makes CSV harder to read
- Some data may contain special characters (commas in names, newlines in descriptions)

**Alternatives Considered:**
1. **Always quote every field:** Safe but verbose, harder to read
2. **Never quote:** Fails if data contains commas or newlines
3. **Quote only if necessary:** Clean, Excel-compatible

**Decision:** Quote only if necessary (RFC 4180 compliant)
- ✅ Readable (most fields unquoted)
- ✅ Excel-compatible (handles quoted fields)
- ✅ Handles edge cases (commas, newlines, quotes)
- ✅ Standard CSV behavior

---

## Query Timeout: 5 Minutes

**Decision:** Queries timeout after 5 minutes and fail fast.

**Owner:** Jake

**Date:** 2026-03-01

**Context:**
- Analyst should know quickly if something is wrong (warehouse down, query inefficient)
- 5 minutes is long enough for typical monthly reports but short enough to fail fast
- Longer timeouts mean analyst waits without feedback
- Shorter timeouts may prematurely kill legitimate slow queries

**Alternatives Considered:**
1. **30 seconds:** Too aggressive for complex queries.
2. **10 minutes:** Long, analyst may think tool is hung.
3. **No timeout:** Risk of indefinite hangs, user can't recover without killing process.
4. **Configurable timeout:** Added complexity, most users won't tune it.

**Decision:** 5 minute timeout
- ✅ Long enough for complex monthly reports
- ✅ Short enough to fail fast on actual problems
- ✅ User can retry with simpler query if needed
- **Future:** Can make configurable if use cases require longer times

---

## Logging: Simple File-Based, Append-Only

**Decision:** Log to `~/.report-gen/report-gen.log` (append-only, no rotation in v1).

**Owner:** Jake

**Date:** 2026-03-02

**Context:**
- Need to capture errors and details for troubleshooting
- Users are on local machines, not servers
- Don't need sophisticated log management in v1
- Simple append-only is sufficient for now

**Alternatives Considered:**
1. **No logging:** Hard to debug user issues.
2. **stdout/stderr only:** Doesn't persist, hard to review later.
3. **Rotating logs (daily, size-based):** Overkill for v1, added complexity.
4. **Structured logging (JSON):** Good for infrastructure, not needed for CLI yet.

**Decision:** Simple append-only text log
- ✅ Easy to implement
- ✅ Users can provide log file when reporting issues
- ✅ No log management overhead
- **Future:** Can add log rotation if file grows too large

---

## Backward Compatibility: Not Required for v1

**Decision:** No guarantee of backward compatibility for v1. Users should expect breaking changes.

**Owner:** Alex

**Date:** 2026-03-05

**Context:**
- This is v1, not a stable API
- Small group of internal users (Maya and colleagues)
- Can fix specs and binaries if we learn something
- Formal backward compatibility commitment would slow iteration

**Alternatives Considered:**
1. **Semantic versioning with stability promises:** Safe, but locks us in too early.
2. **Rapid iteration with breaking changes:** Fast, but frustrating for users.

**Decision:** v1 is experimental, no backward compatibility guarantee
- Users are warned in README that specs and commands may change
- When we ship v2, we'll provide migration guide
- **When:** Once 10+ analysts are using it, we'll consider backward compatibility guarantee

---

## When to Update This Log

Add a new entry when:
- ✅ Architecture decision (language, framework, database)
- ✅ Design tradeoff (feature vs. performance, complexity vs. simplicity)
- ✅ Spec change (if requirement changed after approval)
- ✅ Significant code change with rationale (not "refactored function X")

Don't add entries for:
- ❌ Bug fixes
- ❌ Performance optimizations
- ❌ Code style changes
