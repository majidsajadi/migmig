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
```go
resp, err := migmig.New().Post("https://httpbin.org/post", &migmig.Config{
    Body: map[string]interface{} {
        "Foo": "Bar",
    },
})
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


### Query Parameter
```go
resp, err := migmig.New().Get("https://httpbin.org/get", &migmig.Config{
    QueryParams: map[string]string{
        "foo": "bar",
    },
})
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

### Request Header
```go
client := migmig.New()

resp, err := client.Post("https://httpbin.org/post", &migmig.Config{
    Headers: map[string]interface{} {
        "Foo": "Bar",
    },
})
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

### Delete Request
```go
resp, err := migmig.New().Delete("https://httpbin.org/delete", nil)
if err != nil {
    panic(err)
}
defer resp.Body.Close()

body, err := ioutil.ReadAll(resp.Body)
if err != nil {
    panic(err)
}
```

### Put Request
```go

resp, err := migmig.New().Put("https://httpbin.org/put", &migmig.Config{
    Body: map[string]interface{} {
        "Foo": "Bar",
    },
})
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

### Patch Request
```go

resp, err := migmig.New().Patch("https://httpbin.org/patch", &migmig.Config{
    Body: map[string]interface{} {
        "Foo": "Bar",
    },
})
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





