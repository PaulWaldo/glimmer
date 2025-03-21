# CONVENTIONS

## ⚠️ STRICT TEST-DRIVEN DEVELOPMENT (TDD) WORKFLOW ⚠️

This project follows **Test-Driven Development (TDD)** with ZERO exceptions. Follow this workflow precisely:

### TDD Cycle - The Only Acceptable Process

1. **RED:** Write a failing test first
   - Write only the test code
   - Do not write or suggest any production code
   - Do not even hint at implementation details

2. **CONFIRMATION:** Wait for test failure confirmation
   - The test must be run and confirmed to fail
   - Understand why it fails before proceeding

3. **GREEN:** Only after confirmed failure, write minimal production code
   - Write only enough code to make the test pass
   - Focus on simplicity, not perfection

4. **REFACTOR:** Only after the test passes, consider refactoring
   - Improve the code without changing behavior
   - All tests must still pass after refactoring

### Critical Rules - No Exceptions

- **✋ NEVER suggest production code before a failing test exists and is confirmed**
- **✋ NEVER combine test and implementation in the same step**
- **✋ NEVER skip the confirmation of test failure**
- **✋ NEVER write more implementation than needed to pass the current test**
- **✋ NEVER plan multiple steps ahead - focus only on the current failing test**

### Process Flow - Follow Exactly

1. Write a single failing test for one specific behavior
2. Wait for me to implement and run the test
3. Wait for me to confirm the test fails
4. Only then, write minimal code to make the test pass
5. Wait for me to implement and confirm the test passes
6. Suggest refactoring if needed
7. Repeat from step 1 for the next behavior

## Coding Style

* All code must be idiomatic go code
* All code should be commented for go-doc
* Before accepting any code, format it using `go fmt`
* Prefer the testify unit test framework
