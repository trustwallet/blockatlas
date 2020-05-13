# Checklists & quick info

 - [Bug Reports](https://github.com/trustwallet/blockatlas/blob/master/.github/ISSUE_TEMPLATE/bug_report.md)
 - [Feature Requests](https://github.com/trustwallet/blockatlas/blob/master/.github/ISSUE_TEMPLATE/feature_request.md)
 - [Questions](https://github.com/trustwallet/blockatlas/blob/master/.github/ISSUE_TEMPLATE/question.md**)
 - [Adding a new coin](https://github.com/trustwallet/blockatlas/blob/master/.github/PULL_REQUEST_TEMPLATE/new_blockchain.md)

# Development info

## Tests

### Unit
To run the unit tests: `make test`

### Integration
All integration tests are generated automatically. You only need to set the environment to your coin in the config file.
The tests use a different build constraint, named `integration`.

To run the integration tests: `make integration` 

or you can run manually: `TEST_CONFIG=$(TEST_CONFIG) TEST_COINS=$(TEST_COINS) go test -tags=integration -v ./pkg/integration`

##### Fixtures

 - If you need to change the parameters used in our tests, update the file `pkg/integration/testdata/fixtures.json`
 - To exclude an API from integration tests, you need to add the route inside the file `pkg/integration/testdata/exclude.json`
   E.g.:
```
[
  "/v2/ethereum/collections/:owner",
  "/v2/ethereum/collections/:owner/collection/:collection_id"
]
```

## Error conventions

Use [our error package](https://godoc.org/github.com/trustwallet/blockatlas/pkg/errors) for error handling
and follow Go [best practices](https://blog.golang.org/error-handling-and-go).
All errors thrown in BlockAtlas should use our Error struct:

```
type Error struct {
	Err   error
	Type  Type
	meta  map[string]interface{}
	stack []string
}
```

E is a convenience function for creating errors quickly:

`func E(args ...interface{}) *Error`

Usage: 
 - Wrap a generic error: `errors.E(err)`
 - Wrap error with message: `errors.E(err, "new message to append")`
 - Annotate error with type: `errors.E(err, errors.TypePlatformRequest)`
 - Error with type and metadata:
```
errors.E(err, errors.TypePlatformRequest, errors.Params{
    "coin":   "Ethereum",
    "method": "CurrentBlockNumber",
})
```
 - Any other combinations above


*All fatal errors emitted by logger package already send the error to *

__[List of error types](https://godoc.org/github.com/trustwallet/blockatlas/pkg/errors#Type)__

## Logging Conventions

Use the package `pkg/logger` for logging.
The order of function parameters is arbitrary when using the logger functions.
Examples:

 - Log message: `logger.Info("Loading Observer API")`
 - Log message with params: `logger.Info("Running application", logger.Params{"bind": bind})`
 - Fatal with error: `logger.Fatal("Application failed", err)`
 - Create a simple error log: `logger.Error(err)`
 - Create an error log with a message: `logger.Error("Failed to initialize API", err)`
 - Create an error log, with error, message, and params:
```
p := logger.Params{
	"platform": handle,
	"coin":     platform.Coin(),
}
err := platform.Init()
if err != nil {
	logger.Error("Failed to initialize API", err, p)
}
```
 - Debug log:
   `logger.Debug("Loading Observer API")`
   or 
   `logger.Debug("Loading Observer API", logger.Params{"bind": bind})`
 - Warning log:
   `logger.Warn("Warning", err)`
   or 
   `logger.Warn(err, "Warning")`
   or 
   `logger.Warn("Warning", err, logger.Params{"bind": bind})`
