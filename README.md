# CAPILA

(C)ommon (API) (L)ibrary (A)an Zee

This is the Common library we will use when developing a specific API for a customer.

The logo is a butterfly which a Capilla is too.
(Artwork by Ben)

## Motivation

Since a lot of our programs can use an API, we need a library which would support to write an API in a quick way.
(This will be a work in progress for a while.)

## Getting Started

* You only have to work on this library if you need enhancements in your specific API project.
* Clone the repository and read the doc/documentatie.

### Developing with Docker
First make sure you have access to `aanzeenl/go-pipeline` in the [Docker Hub repositories](https://cloud.docker.com/u/aanzeenl/repository/list) and you have [Docker](https://www.docker.com) installed on your local machine.

Now simply run the following command from the root:
```bash
make watch
```

The Capila watcher will be started and your changes will automatically check against:
- golangci-lint
- staticcheck
- go tests

## Technical information

### Capila CLI (WIP)

The Capila CLI tool is command line interface which handles:
- migrations
- boiler

and might handle the following in the future:
- seeders?
- cache clearing?

To run it you will need to install Capila in your project. This can be done by
```bash
go get github.com/safciplak/capila
```
Next create a folder in your project with the following Go code for example (src/cmd/capila) :
```go
package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	capilaCLI "github.com/safciplak/capila/cli"
)

// main
func main() {
	err := godotenv.Load()

	if err != nil {
		log.Print("Failed to load the .env")
	}

	capilaCLI.Run(os.Args)
}
```

Now build the Capila CLI
```bash
go build -o capila src/cmd/capila
```

This will generate an executable in your current directory. Run it with:
```bash
./capila
```

To create boilerplate resource run the following comment
```bash
capila boiler create [singular-name] [plural-name]
```

## Caching
Capila exposes a caching provider that will wrap a lookup function 
caching in Redis. It will use the REDIS_HOST environment variable
(no authentication) to decide which host to use. If the environment variable 
doesn't exist, it will not cache. If it is an empty string it will use the 
localhost.

```go
package main

import (
	cache "github.com/safciplak/capila/cache"
)

// main
func main() {
	var cacheOptions = cache.CacheOptions{Ttl: time.Second * 5, Slug: "my-slug"}
	func someLookup(key string) MyStruct {...}
	cache := cache.NewLoadableCache(someLookup, cacheOptions)
	someRecord := cache.Get("some-lookup-id")
}
```

## Test / Run etc.

There is a make command that will show you all the possibilities.

TODO: All cap stuff starting with capAlias


> Teruggeven van Stricted Typed data
- response model
- JSONResponse

Swagger/ OpenApi?
- genereren

> Postman
- token vastleggen dmv javascript (Serge)
- library maken

SonarCube
- inrichten voor Go
- kwaliteits eisen  bepalen

> Rechten
Niet alle velden in een update (rechten)
Rechten om een api endpoint te mogen doen
- db.tabellen op db.user (applicatie) toegang geven

Rechten per applicatie-endpoint: (scoping)
Partial Read (is altijd apart endpoint met apart response model)
Read
Write
Partial Write
Delete (Soft)

> SQL injection vs prepared statements
- SqlService krijgt alleen functies (endpoints)



func GetHotels(postcode string) {

}

> Geen shortvar's?
- ook geen i k v
- geen "hongaarse notatie" dus geen mUser, strUser of userStr
