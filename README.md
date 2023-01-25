...

# Useful commands

## Watch mode

Install CompileDaemon

```sh
go install github.com/githubnemo/CompileDaemon
```

Running CompileDaemon

```sh
CompileDaemon -command="./goau-biat.exe"
```

## Running tests

```sh
go clean -testcache && go test ./... -p 1
```

## Add collors to test output

```sh
go test -v ./... -p 1 |
grep -v RUN |
grep -v time= |
sed ''/^PASS/s//$(printf "")/'' |
sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' |
sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
```
