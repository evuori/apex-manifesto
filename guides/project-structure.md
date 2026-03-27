# Project Structure for APEX + SpecDriven Development

## Principle: Structure Serves Documentation

In APEX, how you organize a project should make it obvious:
- Where the user requirements live
- What decisions have been made and why
- What work was completed and in what order
- What code corresponds to what specification

This guide shows patterns that work. Adapt them to your project type—there is no single "right" structure, only structures that serve understanding.

---

## Core APEX Project Structure

This is the baseline structure that works for most projects:

```
project-name/
├── README.md                    # Project entry point
├── docs/
│   ├── project-manifesto.md    # User, problem, scope, agreements
│   ├── specs/
│   │   ├── feature-1.md        # One spec per feature
│   │   ├── feature-2.md
│   │   └── README.md           # Index of specs
│   ├── decisions/
│   │   └── decision-log.md     # Architecture, tech choices, tradeoffs
│   └── architecture.md         # System design (if complex)
├── src/
│   ├── [code organized by function]
│   └── tests/
├── IMPLEMENTATION.md            # Running notes: what's done, what's in progress
└── .github/
    └── CODEOWNERS              # Who owns what (connects to specs)
```

### What Goes Where

**`README.md`** — Project overview at a glance
- What this project does
- How to get started (local dev, running tests, building)
- Key decisions (see decision log for details)
- Links to docs and specs

**`docs/project-manifesto.md`** — User, problem, scope, and team agreement
- Who is this for?
- What problem does it solve?
- What are we building and what are we explicitly not building?
- Key constraints and assumptions
- Who owns what decisions?

**`docs/specs/`** — One spec per user-facing feature or major component
- Each spec follows the SpecDriven structure from [guides/spec-driven-development.md](spec-driven-development.md)
- Use consistent naming: `feature-name.md` or `api-endpoints.md`
- Include a `README.md` in this directory listing all specs with status (planned, in progress, complete)

**`docs/decisions/decision-log.md`** — Significant technical decisions
- Architecture choices (monolith vs. microservices, language selection, database choice)
- Design decisions that trade off different concerns
- When specs changed and why
- Backwards compatibility decisions
- Format: decision name, date, owner, context, decision, alternatives considered, outcome

**`src/`** — Code organized by your project's logic, not by pattern
- For a web service: by domain (users, payments, reports)
- For a library: by exported API (authentication, caching, logging)
- For a CLI: by command (init, run, publish)
- **Not** organized by layer (controllers, models, services) — that's an implementation detail

**`IMPLEMENTATION.md`** — Running notes (optional but valuable)
- What's complete and working
- What's in progress and why
- Known limitations or tech debt
- Next priorities
- Updated as work progresses (unlike specs, which are stable contracts)

---

## Project Type Examples

### Web Service / API

```
api-service/
├── README.md
├── docs/
│   ├── project-manifesto.md
│   ├── specs/
│   │   ├── README.md
│   │   ├── authentication.md      # How users log in/authorize
│   │   ├── user-api.md            # POST/GET/DELETE /users endpoints
│   │   ├── data-export.md         # Bulk export feature
│   │   └── rate-limiting.md       # Quota/rate limit behavior
│   ├── decisions/
│   │   └── decision-log.md
│   └── api-schema.yaml            # OpenAPI/full contract (generated from specs)
├── src/
│   ├── auth/
│   │   ├── middleware.go
│   │   ├── jwt.go
│   │   └── test/
│   ├── users/
│   │   ├── api.go
│   │   ├── model.go
│   │   ├── repository.go
│   │   └── test/
│   ├── export/
│   │   ├── handler.go
│   │   ├── formats.go
│   │   └── test/
│   ├── main.go
│   └── config.go
├── tests/
│   ├── integration/
│   ├── load/
│   └── e2e/
├── Dockerfile
├── go.mod
├── IMPLEMENTATION.md
└── .github/
    └── CODEOWNERS              # auth/ team, users/ team, export/ team
```

### CLI Tool

```
cli-tool/
├── README.md
├── docs/
│   ├── project-manifesto.md
│   ├── specs/
│   │   ├── README.md
│   │   ├── init-command.md       # `tool init` behavior
│   │   ├── run-command.md        # `tool run` behavior
│   │   ├── config-file.md        # Configuration schema and validation
│   │   └── plugin-system.md      # How plugins extend the tool
│   ├── decisions/
│   │   └── decision-log.md
│   └── user-guide.md             # How-to guide (separate from specs)
├── src/
│   ├── cmd/
│   │   ├── init/
│   │   ├── run/
│   │   ├── config/
│   │   └── root.go
│   ├── config/
│   │   ├── loader.go
│   │   ├── validator.go
│   │   └── test/
│   ├── plugins/
│   │   ├── interface.go
│   │   ├── loader.go
│   │   └── test/
│   ├── main.go
│   └── version.go
├── tests/
│   ├── integration/              # Run actual CLI and check output
│   └── fixtures/                 # Sample configs, plugins
├── IMPLEMENTATION.md
└── .github/
    └── CODEOWNERS
```

### Library / SDK

```
library-name/
├── README.md
├── docs/
│   ├── project-manifesto.md
│   ├── specs/
│   │   ├── README.md
│   │   ├── public-api.md         # What's exported, what's the contract?
│   │   ├── caching.md            # Caching behavior and guarantees
│   │   ├── error-handling.md     # Error types, recovery, logging
│   │   └── backwards-compat.md   # Versioning and deprecation policy
│   ├── decisions/
│   │   └── decision-log.md
│   └── DESIGN.md                 # How it works internally
├── src/
│   ├── public/                   # Exported API
│   │   ├── client.py
│   │   ├── models.py
│   │   └── __init__.py
│   ├── internal/                 # Not exported, implementation details
│   │   ├── cache.py
│   │   ├── transport.py
│   │   └── retry.py
│   └── tests/
│       ├── unit/
│       ├── integration/
│       └── fixtures/
├── examples/
│   ├── basic_usage.py
│   └── advanced_usage.py
├── setup.py
├── IMPLEMENTATION.md
└── .github/
    └── CODEOWNERS
```

### Web Application (Frontend + Backend)

```
web-app/
├── README.md
├── docs/
│   ├── project-manifesto.md
│   ├── specs/
│   │   ├── README.md
│   │   ├── user-flows/
│   │   │   ├── authentication.md
│   │   │   ├── dashboard.md
│   │   │   └── settings.md
│   │   ├── backend/
│   │   │   ├── api-endpoints.md
│   │   │   ├── database-schema.md
│   │   │   └── permissions.md
│   │   └── frontend/
│   │       ├── component-library.md
│   │       └── styling-guide.md
│   ├── decisions/
│   │   └── decision-log.md
│   └── architecture.md
├── backend/
│   ├── src/
│   │   ├── auth/
│   │   ├── api/
│   │   ├── db/
│   │   └── main.py
│   ├── tests/
│   ├── requirements.txt
│   └── IMPLEMENTATION.md
├── frontend/
│   ├── src/
│   │   ├── pages/
│   │   ├── components/
│   │   ├── services/
│   │   └── App.tsx
│   ├── tests/
│   ├── package.json
│   └── IMPLEMENTATION.md
├── docker-compose.yml            # Local dev
├── .github/
│   └── CODEOWNERS                # Who owns what (backend owner, frontend owner)
└── IMPLEMENTATION.md              # Shared implementation status
```

---

## Connecting Specs to Code

When you implement a spec, make the connection explicit:

### Option 1: Code Comments
```python
# Implements spec: docs/specs/user-api.md
# Specifically: POST /users endpoint and success criteria
def create_user(name: str, email: str) -> User:
    ...
```

### Option 2: CODEOWNERS File
```
# Declare who owns what and link to the spec

# API Endpoints: see docs/specs/user-api.md
/src/api/users/ @backend-team
/tests/api/users/ @backend-team

# Authentication: see docs/specs/authentication.md
/src/auth/ @auth-owner
```

### Option 3: Feature Branches
When implementing a feature spec, use a branch name that references it:
```
git checkout -b spec/user-api
```

Or include spec reference in commit messages:
```
commit abc123
    Implement POST /users endpoint

    Implements: docs/specs/user-api.md
    Success criteria: ✓ User creation with validation
                      ✓ Error handling for duplicate email
                      ✓ Response format matches spec
```

---

## Documentation Hierarchy

Your project should have clear documentation at different levels:

### Level 1: User/Problem (what and why)
- **File:** `docs/project-manifesto.md`
- **Reader:** Anyone new to the project
- **Questions answered:** What is this? Who is it for? What problem does it solve?

### Level 2: Features & Behavior (what it does)
- **File:** `docs/specs/*.md`
- **Reader:** Engineers implementing, users understanding behavior
- **Questions answered:** What are the workflows? What are the edge cases? How do I know it works?

### Level 3: Architecture & Decisions (why we built it this way)
- **File:** `docs/decisions/decision-log.md`, `docs/architecture.md`
- **Reader:** Engineers maintaining or extending the system
- **Questions answered:** Why this language? Why this database? What were the tradeoffs?

### Level 4: Implementation (how it works internally)
- **File:** Code comments, design docs
- **Reader:** Engineers modifying the code
- **Questions answered:** How does this module work? What invariants must I maintain?

---

## Specs as Living Documents

Specs start as the contract before building. They stay stable unless reality changes:

- ✅ **Update the spec if:** You discover a requirement was wrong, a constraint changed, or a better approach emerged
- ❌ **Don't update the spec if:** You just chose a different implementation method (that's an implementation detail, not a spec change)

When you update a spec:
1. Document the change in the spec itself (add a note: "Updated [date]: [what changed and why]")
2. Add an entry to the decision log explaining the change and why

---

## Using IMPLEMENTATION.md

Unlike specs (which are contracts), `IMPLEMENTATION.md` is a running log of current state:

```markdown
# Current Status

## Complete ✓
- User authentication (spec: docs/specs/authentication.md)
- Basic CRUD endpoints (spec: docs/specs/user-api.md)
- Error handling with standardized format

## In Progress 🔄
- Data export feature (spec: docs/specs/data-export.md)
  - CSV format working
  - JSON format in progress
  - Quota checking: not yet started

## Known Limitations
- Single-user auth only (multi-user added to backlog, see decision-log.md)
- No caching (added in v2, see decision-log.md)

## Next Priorities
1. Finish data export
2. Add rate limiting
3. Performance testing at scale

## Tech Debt
- Tests for export edge cases need expansion
- Database connection pooling not optimized
```

---

## Organizing Multiple Projects

If you have multiple related projects (monorepo or project suite):

```
org-projects/
├── README.md                     # Overview of all projects
├── projects/
│   ├── service-a/
│   │   ├── docs/
│   │   │   ├── project-manifesto.md
│   │   │   └── specs/
│   │   └── src/
│   ├── service-b/
│   │   ├── docs/
│   │   │   ├── project-manifesto.md
│   │   │   └── specs/
│   │   └── src/
│   └── shared-library/
│       ├── docs/
│       └── src/
├── docs/
│   ├── architecture.md           # How all projects fit together
│   ├── decisions/
│   │   └── decision-log.md       # Org-level decisions
│   └── development.md            # Common development practices
└── .github/
    └── CODEOWNERS
```

---

## Next Steps

1. Pick a project type above that matches yours
2. Adapt the structure to your actual needs (don't blindly copy everything)
3. Create `docs/project-manifesto.md` first — it will clarify what specs you actually need
4. Write one spec for your next feature
5. Reference the spec from your code (comments, CODEOWNERS, or commit messages)
6. Keep `IMPLEMENTATION.md` updated as you work
