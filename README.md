# wanon
It's a bot for telegram in go to add and retrieve quotes.

## Developer notes

* Remember to run tests with:
  * From the project folder: `go test ./...`
  * From outside: `go test github.com/graffic/wanon...`

### Mocks

For mocks this project uses `testify/mock` and `mockery`. Generated mocks are 
a PITA, or at least the way I like them:

```
mockery -dir telegram -name Request -inpkg -testonly
```

This creates:

* A mock for the Request interface
* Found in the telegram package
* Creates a file in the package
* appends a _test on it
