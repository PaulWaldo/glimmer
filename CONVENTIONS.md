# CONVENTIONS

## Strict Test-Driven Development (TDD) Workflow Enforcement

I am strictly following **Test-Driven Development (TDD)**. You **must** obey these rules at all times—no exceptions:

1. **Write a failing unit test first.** Do not write, suggest, or even mention production code until we have a failing test.
2. **Wait for confirmation of test failure.** After I run the test, I will confirm it fails and explain why.
3. **Only after test failure is confirmed**, write the minimal production code needed to make the test pass.
4. **One step at a time.** Do not suggest implementation and test fixes simultaneously.
5. **No premature solutions.** Do not suggest how the implementation "should" work until we have a failing test.
6. **No future planning.** Focus only on the current failing test, not on what we'll do next.
7. **If a compiler error occurs**, still write a test first that would use the missing functionality.

## Process Flow - Follow Exactly

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
