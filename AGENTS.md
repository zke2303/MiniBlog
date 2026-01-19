# Agent Guidelines: MiniBlog

This document provides essential information for agentic coding agents operating in the MiniBlog repository.

## 1. Project Overview
- **Backend (Proposed/Documentation):** Gin (Go), GORM, PostgreSQL, Redis, JWT.
- **Frontend:** Vue.js 2 (Vue CLI 5), Babel, ESLint.
- **Current State:** The frontend scaffold exists in `ui/`. Backend documentation is in `资料/接口文档/interface.md`. The `apps/` directory is currently empty.

## 2. Development Commands

### Frontend (`ui/` directory)
All commands should be run from the `ui/` directory.

- **Install dependencies:** `npm install`
- **Start development server:** `npm run serve`
- **Build for production:** `npm run build`
- **Lint and fix files:** `npm run lint`
- **Run single test:** No test runner currently configured in `package.json`. If adding tests, prefer Jest or Mocha as per Vue CLI standards.

### Backend (`apps/` directory)
*Note: Backend is not yet implemented. Follow these conventions if starting implementation.*
- **Build:** `go build -o main .`
- **Run:** `go run main.go`
- **Test all:** `go test ./...`
- **Test single package:** `go test ./path/to/package`
- **Test single function:** `go test -v -run TestFunctionName ./path/to/package`

## 3. Code Style Guidelines

### General
- **Indentation:** 2 spaces for Frontend (JS/Vue), 4 spaces (tabs) for Go.
- **Naming:**
  - **JS/Vue:** camelCase for variables/functions, PascalCase for Components (`HelloWorld.vue`).
  - **Go:** PascalCase for exported symbols, camelCase for unexported.
  - **CSS:** kebab-case for classes.

### Frontend (Vue.js 2)
- **Imports:** 
  - Group standard library/external imports first, then internal assets/components.
  - Use relative paths (e.g., `./components/HelloWorld.vue`).
- **Templates:**
  - Use `id` for root app container (`#app`).
  - Prefer self-closing tags for components with no content.
- **Scripts:**
  - Always provide a `name` property for components.
  - Export components using `export default`.
- **Styles:**
  - Use `scoped` styles in components to prevent global leakage.
- **Formatting:** Adhere to `eslint:recommended` and `plugin:vue/essential`.

### Backend (Go/Gin) - Expected
- **Imports:** Grouped: standard library, blank line, third-party libraries (Gin, GORM).
- **Error Handling:** 
  - Check `if err != nil` immediately.
  - Return uniform JSON errors as per `interface.md`: `{"code": 400, "message": "reason"}`.
- **API Versioning:** Prefix all routes with `/api/v1`.

## 4. Architecture & Integration
- **API Documentation:** Reference `资料/接口文档/interface.md` for endpoint specs, request/response models, and status codes.
- **Authentication:** Use JWT via `Authorization: Bearer <token>`.
- **Database:** PostgreSQL for persistent data, Redis for token blacklisting and caching.

## 5. Security Protocols
- **Secrets:** Never commit `.env` files, JWT secrets, or database credentials.
- **Validation:** Always validate request bodies on the backend using Gin's binding/validation features.

## 6. Directory Structure
- `ui/`: Vue.js frontend.
- `apps/`: Target directory for backend Go services.
- `资料/`: Project documentation and assets.

---
*Follow these guidelines to maintain consistency across the codebase.*
