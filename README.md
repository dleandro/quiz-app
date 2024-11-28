Quiz CLI



Next steps:

- Improve logging maybe with dependency that has logging formatting and timestamps
- Add a Quiz domain entity 
- Separate each endpoint in its own controller file
- Env variable handling
- Docker compose
- Add ability to manage the available questions (this would require authentication)
- Add tests for each endpoint
- Could have adapters for the dependencies
- Dependency injection
- Extract the Parsing and reading of proto files to infra adapters or aux methods (basically extracting duplicated code)
- Reduce the amount of loops on the server getStats req