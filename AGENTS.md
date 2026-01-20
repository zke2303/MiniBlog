# Agent Guidelines: MiniBlog

This document provides essential information for agentic coding agents operating in the MiniBlog repository.

## 1. Project Overview
- **Backend:** Go (Gin), GORM (PostgreSQL), Viper. Located in `apps/`.
- **Frontend:** Vue.js 2 (Vue CLI 5), Babel, ESLint. Located in `ui/`.
- **Documentation:** `资料/接口文档/interface.md` contains API specifications.

## 2. Development Commands

### Backend (`apps/` directory)
All commands should be run from the `apps/` directory.

- **Build:** `go build -o bin/serve cmd/serve/main.go`
- **Run:** `go run cmd/serve/main.go`
- **Test all:** `go test ./...`
- **Test single package:** `go test ./internal/service/...`
- **Test single function:** `go test -v -run ^TestFunctionName$ ./internal/service`
- **Tidy dependencies:** `go mod tidy`

### Frontend (`ui/` directory)
All commands should be run from the `ui/` directory.

- **Install dependencies:** `npm install`
- **Start development server:** `npm run serve`
- **Build for production:** `npm run build`
- **Lint and fix files:** `npm run lint`
- **Run tests:** No test runner currently configured. If adding tests, use Jest as per Vue CLI standards.

## 3. Code Style & Conventions

### Go (Backend)
- **Formatting:** Use `gofmt` or `goimports`. Indentation is 4-space tabs.
- **Naming:** 
  - PascalCase for exported symbols (functions, structs, fields).
  - camelCase for unexported symbols and local variables.
  - Interface names should ideally end with `Repository` or `Service` or follow `IAbc` if explicitly preferred (current usage: `IUserRepository`).
- **Imports:** Grouped into three sections separated by newlines:
  1. Standard library
  2. Internal project packages (`mini-blog/...`)
  3. External third-party libraries (`github.com/...`)
- **Error Handling:**
  - Check `if err != nil` immediately after function calls.
  - Use the custom `errmsg` package for business errors: `errmsg.InternalErr.Wrap(err)`.
  - In controllers, use `c.Error(err)` to record errors for middleware processing.
- **Architecture:** Follow the Controller -> Service -> Repository pattern.
  - `dto/request`: Incoming request models with validation tags.
  - `dto/response`: Outgoing response models.
  - `dto/errmsg`: Standardized business error codes and messages.

### Vue.js (Frontend)
- **Formatting:** 2-space indentation. Adhere to `eslint:recommended` and `plugin:vue/essential`.
- **Templates:**
  - Root container should use `#app`.
  - Prefer self-closing tags for components without slots.
- **Scripts:**
  - Always provide a `name` property for components.
  - Use `export default { ... }`.
  - Data should be a function: `data() { return { ... } }`.
- **Styles:** Use `<style scoped>` to prevent global style pollution.
- **Naming:**
  - File names: PascalCase (`UserList.vue`).
  - Variables/Methods: camelCase.

## 4. Architecture & Design Patterns
- **API Versioning:** Prefix all routes with `/api/v1`.
- **Validation:** Use `c.ShouldBindJSON` and Go's `validator/v10` tags (e.g., `binding:"required"`).
- **Responses:** Always return standardized JSON:
  ```json
  {
    "code": 200,
    "msg": "success",
    "data": { ... }
  }
  ```
  Use `response.Success(c, data)` helper.

## 5. Security & Safety
- **Secrets:** Never commit `.env`, `config.yml` with real credentials, or JWT secrets. Use environment variables for sensitive data.
- **Passwords:** Always hash passwords using `bcrypt` before storing in the database.
- **Database:** Use GORM transactions for operations involving multiple writes.

## 6. Directory Structure
- `apps/cmd/serve`: Application entry point.
- `apps/internal/controller`: Request handling and parameter binding.
- `apps/internal/service`: Business logic orchestration.
- `apps/internal/repository`: Data access layer (GORM).
- `apps/internal/dto`: Data Transfer Objects (Requests, Responses, Errors).
- `ui/src/components`: Reusable Vue components.
- `ui/src/views`: Page-level Vue components.

---
*Follow these guidelines to maintain consistency and ensure code quality.*
