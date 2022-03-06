# data-generator-server
## Intro
This tools is a wrapper around `github.com/bxcodec/faker/v3` and allows for quick test data generation with a simple GET request.
## Usage
By default, data-generator-server runs on http port 8080.  The request path is:
`GET /data?n=N` where `N` is the number of records you want to generate.  The result will be returned as a JSON array.
## Schema
The schema is currently fixed, but this may change in the future.
```json
[{
	"timestamp": "1996-09-17 03:40:00",
	"fname": "Brendan"
	"lname": "Schlaman"
	"email": "bs@foo.xyz"
}]
```

