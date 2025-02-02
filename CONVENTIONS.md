# CONVENTIONS

## Test-Driven Development (TDD) Workflow Enforcement

I want to strictly follow **Test-Driven Development (TDD)**. Follow these rules without exception:

1. **Write a failing unit test first.** Do not write any production code until we have a clear, failing test that defines the expected behavior.
2. **Keep tests minimal and focused.** Each test should target a single behavior or requirement.
3. **Run the test and confirm failure.** Do not assume failure—explicitly state why the test is failing.
4. **Only then, write the minimal production code** needed to make the test pass. No extra features or optimizations.
5. **Refactor if necessary.** Only refactor once the test passes, ensuring all tests remain green.

At every step, **confirm with me before proceeding** to the next stage. **Never skip a step.**

## Test Driven Development

* Act as an expert in Test Driven Development
* You always follow the "Red, Green, Refactor" methodology.
* Before you start to think about changing any production code you create a failing unit test.
* You should fully understand the change that is being requested so you can formulate the failing test.  If you are unclear about what the production code should do once implemented, you must ask for more clarification.
* If it is not clear how to create such a test you will interact with USER to refine the requirements.
* Note that the failing test may be a change to an existing test that passes.
* Once the test is created you will run all tests and ensure that the new (or changed) test fails.
* Only then will you suggest changes for production code.
* Once the production code has been changed, you will again run all tests.
* If the test passes, the code change can be accepted.
* If the test fails, continue with code modification.
* When the production code modification has been made and the unit tests pass, commit the changes to git.
* after committing, you should consider refactoring the associated code.  Provide suggestions for making the tested code better.
* Refactoring should never change the behavior, just make the code more efficient, idiomatic, and easier to understand.
* Suggest any refactoring changes as a discussion with USER.  Ask USER whether he wants the refactoring changes. If YES, make the changes and re-run all unit tests.  If NO, simply stop.
