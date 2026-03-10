# Recipe Manager

## Project

Go web application for recipe management, meal planning, and shopping list generation. Uses Go stdlib net/http with in-memory storage. Deployed to Vercel.

## Commands

- `go test ./internal/... ./pkg/... -v -race` - Run unit tests
- `go test ./test/integration/... -v -tags integration` - Run integration tests
- `go test ./test/scenarios/... -v -tags scenario` - Run scenario tests
- `go test ./... -v` - Run all tests

# Testing Requirements (AX)

Every feature implementation MUST include tests at all three tiers:

## Test Tiers
1. **Unit tests** -- Test individual functions/methods in isolation. Mock external dependencies.
2. **Integration tests** -- Test component interactions with real services via docker-compose.test.yml.
3. **Scenario tests** -- Test full user workflows end-to-end.

## Test Naming
Use semantic names: `Test<Component>_<Behavior>[_<Condition>]`
- Good: `TestAuthService_LoginWithValidCredentials`, `TestFullCheckoutFlow`
- Bad: `TestShouldWork`, `Test1`, `TestGivenUserWhenLoginThenSuccess`

## Reference
- See `TEST_GUIDE.md` for requirement-to-test mapping
- See `.claude/ax/references/testing-pyramid.md` for full methodology
- Every requirement in ROADMAP.md must map to at least one scenario test
