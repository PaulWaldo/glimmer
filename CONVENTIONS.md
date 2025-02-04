# CONVENTIONS

## Strict Test-Driven Development (TDD) Workflow Enforcement

I am strictly following **Test-Driven Development (TDD)**. You **must** obey these rules at all times—no exceptions:

1. **Write a failing unit test first.** Do not write or modify production code until we have a failing test.
2. **If a compiler error occurs, do not fix it** until we have written a failing test that necessitates a fix. TDD **does not allow premature changes** to production code.
3. **Keep tests minimal and focused.** Each test should target a single behavior or requirement.
4. **Run the test and confirm failure.** Do not assume failure—explicitly state why the test is failing.
5. **Only then, write the minimal production code** needed to make the test pass. No extra features or optimizations.
6. **Refactor only after passing tests.** All tests must remain green before refactoring.

At every step, **wait for my confirmation before proceeding**. **If you detect a compiler error, you must still follow the TDD workflow—do not change production code until we have a failing test.**

Now, let’s begin—what feature or function are we testing first?
