# MigMig
Simple HTTP Client for Golang

## Features

* Get/Post/Put/Head/Delete/Patch/Options
* Clean and Minimal API
* Set Request Headers
* Set URL Query Parameters
* JSON Request Body
* Creating an Instance with Default Config

## Usage
    
Install MigMig using `go get`
```bash
$ go get github.com/majidsajadi/migmig
```

Import MigMig into your code and refer it as `migmig`.
```go
import "github.com/majidsajadi/migmig"
```

### Get Request
Simple get request
```go
client := migmig.New()

resp, err := client.Get("https://httpbin.org/get", nil)
if err != nil {
    panic(err)
}
defer resp.Body.Close()

body, err := ioutil.ReadAll(resp.Body)
if err != nil {
    panic(err)
}

fmt.Println(string(body))
```

### Post Request
Simple post request with headers and JSON body
```go
resp, err := migmig.New().Post("https://httpbin.org/post", &migmig.Config{
    Headers: map[string]string{
        "Content-Type": "application/json",
    },
    Body: map[string]interface{} {
        "foo": "bar",
    },
})
```

### Creating an Instance
You can create a migmig instance with default config
```go
client := migmig.Create(migmig.Config{
    BaseURL: "https://api.github.com",
    Headers: map[string]string{
        "accept": "application/vnd.github.v3+json",
    },
})

resp, err := client.Get("/users",
    &migmig.Config{
        QueryParams: map[string]string{
            "per_page": "25",
            "since":    "300",
        },
    })
if err != nil {
    panic(err)
}
defer resp.Body.Close()

users, err := ioutil.ReadAll(resp.Body)
if err != nil {
    panic(err)
}

fmt.Println(string(users))

resp, err = client.Get("/users/majidsajadi/repos",
    &migmig.Config{
        QueryParams: map[string]string{
            "per_page": "25",
        },
    })
if err != nil {
    panic(err)
}
defer resp.Body.Close()

repos, err := ioutil.ReadAll(resp.Body)
if err != nil {
    panic(err)
}

fmt.Println(string(repos))
```


### Query Parameter
```go
resp, err := migmig.New().Get("https://httpbin.org/get", &migmig.Config{
    QueryParams: map[string]string{
        "foo": "bar",
    },
})
```

### Request Header
```go
resp, err := migmig.New().Post("https://httpbin.org/post", &migmig.Config{
    Headers: map[string]string {
        "foo": "bar",
    },
})
```

### Delete Request
```go
resp, err := migmig.New().Delete("https://httpbin.org/delete", nil)
```

### Put Request
```go

resp, err := migmig.New().Put("https://httpbin.org/put", &migmig.Config{
    Body: map[string]interface{} {
        "foo": "bar",
    },
})
```

### Patch Request
```go
resp, err := migmig.New().Patch("https://httpbin.org/patch", &migmig.Config{
    Body: map[string]interface{} {
        "foo": "bar",
    },
})
```

## Contirbuting
Any contribution, pull requests, issue and feedbacks would be greatly appreciated.
If you have any idea about the MigMig, such as, feature requests, refactoring, API changes, etc, feel free to open an issue.

## License 
MigMig is open source software licensed as [MIT](LICENSE). 