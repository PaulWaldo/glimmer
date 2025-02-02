# CONVENTIONS

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
