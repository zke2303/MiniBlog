# Agent Guidelines: MiniBlog

This document provides essential information for agentic coding agents operating in the MiniBlog repository.
Adhere strictly to these guidelines to ensure consistency and maintainability.

## 1. Build, Lint, and Test Commands

### Backend (`apps/` directory)
**Working Directory:** All backend commands **must** be run from the `apps/` directory.

- **Build:**
  ```bash
  go build -o bin/serve cmd/serve/main.go
  ```
- **Run (Dev):**
  ```bash
  go run cmd/serve/main.go
  ```
- **Test All:**
  ```bash
  go test -v ./...
  ```
- **Test Single Package:**
  ```bash
  go test -v ./internal/service/...
  ```
- **Test Single Function (Targeted):**
  ```bash
  # Syntax: go test -v -run ^TestName$ <package_path>
  go test -v -run ^TestUserCreate$ ./internal/service
  ```
- **Lint/Tidy:**
  ```bash
  go mod tidy
  # If golangci-lint is available:
  # golangci-lint run ./...
  ```
- **Swagger Generation:**
  ```bash
  # If modifying controller annotations
  swag init -g cmd/serve/main.go -o docs
  ```

### Frontend (`ui/` directory)
**Working Directory:** All frontend commands **must** be run from the `ui/` directory.

- **Install Dependencies:**
  ```bash
  npm install
  ```
- **Start Dev Server:**
  ```bash
  npm run serve
  ```
- **Build for Production:**
  ```bash
  npm run build
  ```
- **Lint and Fix:**
  ```bash
  npm run lint
  ```

---

## 2. Code Style & Conventions

### Go (Backend)

**General:**
-   **Formatting:** Standard `gofmt` or `goimports`. 4-space tabs for indentation.
-   **Comments:** Business logic comments are often in **Chinese**. Maintain this style for high-level descriptions.
    -   Example: `// UserService 用户业务逻辑对象`
-   **Imports:** Group imports into three distinct blocks separated by newlines:
    1.  Standard Library
    2.  Internal Project Packages (`mini-blog/...`)
    3.  External/Third-party Libraries (`github.com/...`)

**Naming:**
-   **Exported:** PascalCase (e.g., `UserService`, `Create`).
-   **Unexported/Local:** camelCase (e.g., `db`, `repo`, `user`).
-   **Interfaces:** Should usually match the implementation name + `I` prefix if explicitly defined, or suffix `Repository`/`Service`.
    -   Current usage: `IUserRepository`.

**Error Handling:**
-   **Immediate Check:** Check `if err != nil` immediately after the call.
-   **Business Errors:** Use `mini-blog/internal/dto/errmsg` for typed errors.
    -   Wrapping: `errmsg.InternalErr.Wrap(err)`
    -   Creation: `errmsg.New(errmsg.CodeInternalErr, "Description", err)`
-   **Controller:** Use `c.Error(err)` to pass errors to the middleware for centralized handling.
    -   Do not manually construct error JSON responses if `c.Error` is used.

**Models (GORM):**
-   **IDs:** Use `uuid.UUID` (`github.com/google/uuid`).
-   **Tags:** Explicitly define `json` and `gorm` tags.
    -   `json:"-"` for sensitive fields (Password) or internal fields (DeleteAt).
    -   `gorm:"column:create_time"` mapping is common.
-   **Soft Delete:** Use `gorm.DeletedAt`.

**Architecture:**
-   **Pattern:** Controller -> Service -> Repository.
-   **Context:** Pass `context.Context` (usually `c.Request.Context()`) down to Service and Repository layers.
-   **Transactions:** Use `svc.db.Transaction(func(tx *gorm.DB) error { ... })` in the Service layer.

### Vue.js (Frontend)

**General:**
-   **Version:** Vue 2.x options API.
-   **Formatting:** 2-space indentation. Adhere to `eslint:recommended` and `plugin:vue/essential`.
-   **Structure:**
    ```vue
    <template>
      <div id="app">...</div>
    </template>
    <script>
    export default {
      name: 'ComponentName',
      data() { return { ... } }, // Data must be a function
      methods: { ... }
    }
    </script>
    <style scoped>...</style>
    ```

**Naming:**
-   **Files:** PascalCase (e.g., `UserList.vue`, `App.vue`).
-   **Components:** PascalCase `name` property.
-   **Props/Data:** camelCase.

---

## 3. Architecture & Design Patterns

**API Design:**
-   **Versioning:** Prefix routes with `/api/v1`.
-   **Documentation:** Use `swag` annotations (go-swagger) on controller methods.
    -   Example: `// @router /api/v1/auth/profile [get]`
-   **Response Format:**
    All successful API responses return a standardized JSON structure:
    ```json
    {
      "code": 200,
      "msg": "success",
      "data": { ... }
    }
    ```
    Use the helper: `response.Success(c, data)`.

**Request Validation:**
-   Use `gin` binding and `validator/v10` tags in DTO structs.
    -   Example: `binding:"required,min=6"`

**Security:**
-   **Passwords:** NEVER store plain text. Use `bcrypt` (cost: default) before saving.
-   **Secrets:** Configs (`config.yml`) should not contain real production secrets in git.
-   **Context:** User ID is often stored in the Gin context (e.g., `c.GetString("userID")`).

## 4. Directory Structure Reference

-   `apps/cmd/serve`: Entry point (main.go).
-   `apps/internal/controller`: HTTP handlers (Gin).
-   `apps/internal/service`: Business logic, transactions.
-   `apps/internal/repository`: Database access (GORM).
-   `apps/internal/dto`: Data Transfer Objects.
    -   `request`: Request binding structs.
    -   `response`: Response structs.
    -   `errmsg`: Error codes and definitions.
-   `apps/internal/model`: DB entities.
-   `ui/src/views`: Page-level components (mapped to routes).
-   `ui/src/components`: Reusable UI widgets.
-   `ui/src/api`: Axios wrapper and API calls (if present).

---
*Generated for agentic use. adhere strictly to these patterns.*
