# TSARKA Tech Task
This API provides two endpoints that accept and return text/plain content-type.
## Find all substrings
Finds the maximal substring containing no repeating characters.

## Email Validation Endpoint

Checks and returns only the valid email addresses from the input string.

## Example Usage

Send a POST request with a text/plain request body to the corresponding endpoint URL.

For example:

```bash

    POST http://localhost:9090/rest/substr/find
    Content-Type: text/plain

    Hello, World!
```
This would return:

```
, World!
```

Similarly, you can use the `/rest/email/check` endpoint to validate and return only the valid email addresses.

## How to start server
Run from binary file
```bash
$ make build

$ ./app
```

Run from main file
```bash
$ make run
```
