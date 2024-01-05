# Release notes

Always check these release notes before updating.

### v1.21.0
- Added log output to terminal as a missing or empty value for key on .env file.

### v1.20.0
#### New Stuff
- Gin cache middleware added

### v1.19.2
- APM package version bumped to 2 

### v1.19.1
- Added NewLoggerInstance method to reduce BC breaking change in 1.16.0. It will return the ZipLogger instead of the interface. 

### v1.19.0
- New Price To Cents function added
- Dependencies updated

### v1.18.0
- Added a retry on DB connections on setup. The application attempts to create a DB connection for 20 seconds with a 2 second interval. This is done for race conditions when the DB takes longer to launch.

### v1.17.0
- Added a rawQuery adapter to the bun DB interfaces. This solves an issue with ScanRows not returning an `sql.ErrNoRows`.
- The `request` and `getAPMSpanHandlerName` are made public, so they can be used when using a custom `GetHandlerFunc`.
---

### v1.16.0
- Request response logging.
- Response timeout 15 to 60 seconds.
- Made logger mockable. (update to 1.19.1 with NewLoggerIntance to get the original return type)
---

### v1.15.0
- Added Bun feature.
### v1.14.0
#### Golang version upgraded
- go version upgraded to 1.19
---

### v1.13.0
#### Logger package improvements
- EnvironmentHelper dependency removed from logger.
- Singleton pattern added to logger package.
---

### v1.12.0
#### Price convert package converted to generics
- Go version upgraded to 1.18
- Old price methods removed and new convert methods added with using generic.
---

### v1.11.0
#### New Stuff
- Added InterfaceGetterSetter for setting cache values
- Extended InterfaceTimeHelper with `NowIn(locationName, fallbackLocationName string) time.Time`

---

### v1.10.0
#### New Stuff
- Added APM logging for cache

---
### v1.9.0
#### Breaking changes
- Renamed the `LANGUAGE_KEY` const to `LanguageKey` to fit Golang standards.
#### New Stuff
- Added a mock helper which is used for testing in microservices.
- Added caching support with `NewLoadableCache`
- Removed the last references to LIBGOAZ. It is now deprecated!
- Extended HTTPMocker to record request and response

---

### v1.8.1

*Bugfix*

- Fixed `file.GetRootDir()`

---

### v1.8.0
#### New Stuff
- Added JSON:API error based responses to the JSON Presenter.
- Added Generic HTTP Status errors and a mechanism to introduce application specific custom errors.
- Added new Panic handler, panics no longer result in empty response bodies and prints a stacktrace in console.

#### Breaking changes
- pg.NoRowsErr no longer results in a 404 response but should be handled by the application itself.
- Request parsing errors result in 400 errors now
- Error response syntax is different now, data no longer contains the error string

---
### v1.7.3

*Bugfix*

- Fixed `file.GetRootDir()`

### v1.7.0

#### (Breaking ) changes
- Updated all dependencies
- Removed the deprecated PriceHelper. Use convert package instead.

#### New Stuff
- New helpers in the convert package: `NewFloat64` `NewUint`. They return pointer versions of the given input.
- Added uint support to price helpers: `Float32ToCentsUint` `Float32ToCentsUint`
- Added a Slice helper with a `CombineOnlyMatchingValues` function.
- Extended the timeHelper with a handy `NewDate` utility.

### v1.6.0

#### New Stuff

- Moved part of the file package from Libgoaz to Capila
- Added a mock Logger and logger interface

#### changes
- improved boilerplate templates for faster development

### v1.5.2

#### New Stuff

- Deprecated the priceHelper ( this will be removed in 1.6.0 ~)
- Added the price helper functions into the convert package.

### v1.5.1

*Bugfix*

- Fixed migrate drop/fresh commands

### v1.5.0

#### Breaking changes

- The following packages have been moved into an `http` folder to improve readability for the Capila folders.
  ```text
  context
  handlers
  middleware
  presenters
  response
  web
  ```
  Please change your imports from:
  ```go
  "github.com/safciplak/capila/src/context" to "github.com/safciplak/capila/src/http/context"
  "github.com/safciplak/capila/src/handlers" to "github.com/safciplak/capila/src/http/handlers"
  "github.com/safciplak/capila/src/middleware" to "github.com/safciplak/capila/src/http/middleware"
  "github.com/safciplak/capila/src/presenters" to "github.com/safciplak/capila/src/http/presenters"
  "github.com/safciplak/capila/src/response" to "github.com/safciplak/capila/src/http/response"
  "github.com/safciplak/capila/src/web" to "github.com/safciplak/capila/src/http/web"
  ```

- The context `GetLanguage` function has been renamed to `GetTwoLetterLanguageCode` because it now supports `nl-NL` as
  well.
- The config package that contained a `func Loader` for `github.com/joho/godotenv` has been removed due to using Docker
  ENV's and Rancher secrets which do not require a `.env` file. We found the log message this outputs only confuses
  people. If you use this package please add it to your project instead.

#### New Stuff

- The logger now supports log levels based on an ENV variable. This means you can now debug API calls with debug
  logging. Make sure to restart your service because it only loads on restarts. Add `LOG_LEVEL` to your environment with
  on of the following values:
    - debug
    - info
    - warn
- Added support for four letter language codes by adding `GetFourLetterLanguageCode` to the context package.
- Added `priceHelper` which convert prices to cents.
- Added `requestHelper` which allows you to easily make requests to external URL's.
- Added `timeHelper` for mocking the Go `time.Time`.
- Added `NewString` which is a handy utility for turning a string into a pointer string.
- Removed some libgoaz asserts.
- Moved some conversion helpers from libgoaz to Capila.

### v1.4.0

#### New Stuff

- Added mocking interfaces/adapters for the pg transaction and statement which makes them mockable

### v1.3.5

*Bugfix*

- Fixed some small bugs within the Boilerplate templates
- Added internal project imports in the Boilerplate templates

### v1.3.4

*Bugfix*

- GetHandlerFunc now handles the context correctly
- Boilerplate templates are fixed to work with the new GetHandlerFunc

### v1.3.1

*Bugfix*

- NewEnvironmentHelper should return an interface for mocking.
- EnvironmentHelper didn't expose the error anymore. The new function `Error()` now returns it ( or nil ).
- Fixed a bug with the HTTPMocker where it used the Capila path as base path.

### v1.3.0

#### Breaking changes

- Everything was moved a directory deeper. Change your imports from:

```go
import (
"github.com/safciplak/capila/apm"
"github.com/safciplak/capila/database"
)
```

to:

```go
import (
"github.com/safciplak/capila/src/apm"
"github.com/safciplak/capila/src/database"
)
```

- The Environment helper has been refactored. All methods are now receiver functions so they're mockable.

```go
GetEnvironmentString(variable) -> environmentHelper.GetString(variable)
GetEnvironmentInteger(variable) -> environmentHelper.GetInteger(variable)
GetEnvironmentBoolean(variable) -> environmentHelper.GetBoolean(variable)
```

#### New Stuff

- Installed the new watcher and updated the folder structure.
- Improved JSON presenter with JSON HAL specifications. See https://tools.ietf.org/html/draft-kelly-json-hal-08.
- Added a Gin middleware that extracts the `accept-language` from the headers and puts it in the context. This way its
  easy to retrieve the language the user wants by calling `capilaContext.GetLanguage(ctx)`. This function returns a
  lowercase string that defaults to "en".
- Added Zap logger which improves log output.
- Added a default router that contains the most used middleware for easier setups.
- httpclientmock was added to connect the go tests with the Mock package that is used in most microservices. This allows
  you to use the JSON files from the mock.

### v1.2.1 - Remove goimports, fix boilerplate tests

*Bugfix*

- Remove goimports, because it did not solve imports automatically.
- Fix generated repository boilerplate tests, used wrong type.

### v1.2.0 - Add boilterplate, move handlers

*New stuff*

- Boilerplate
- New handlers package, used in microservice-starter for simpler handlers.

### v1.1.0 - Database bugfix

*Breaking changes*

- interfacePGConn is renamed to interfacePGDB, this will have an impact on mocks used in unit tests
  *New stuff*
- Go-pg is refactored so the threadpool is used instead of a single connection. This fixes an issue where a broken
  connection was kept being used after a connection time out occurred, causing all the subsequent calls to fail.

### V1.0.0 - Major update

*Breaking changes*

- Context can no longer be set as a global. Setting it as a global causes the context cancelled issue when there are
  multiple http calls at the same time. It should now be passed as a parameter.
    - `StartSpan` has been removed. Use Start instead
    - `EndSpan` has been removed. Use End instead
- `sqlx` is removed and replaced by `go-pg`
    - The sqlx connector was replaced by a `go-pg` connector.
- The log packages has been removed. Debugging can be done by setting `DB_DEBUG=true` in your project.
- All test asserts have been replaced with stretchr/testify. This package contains more and better assertions. It should
  make testing easier in the future.
- Old httprouter middleware has been removed because of default Gin Middleware
- The presenters have been changed so they are easier to use
    - New Presenters should implement the new PresenterInterface
- Cleaned up the Makefile. Added Mockery and fixed the test command.
- Validate package has been removed. Use Gin instead.
- Removed some unused Errors in the error package. If you miss one for your project please add it again.
  *New stuff*
- Linter update
- Pipeline update

### V0.10.0

*New stuff*

- Added shorter APM span methods: Start & End

### V0.9.1

*New stuff*

- TraceError now allow nil's

### V0.9.0

*New stuff*

- Added changelog.
- Added a way to add `items` to the `link` element for the JSON presenter.

The validator library got a small update

---

### V0.8.0

*New stuff*

- Added APM package to Capila
  Important: Renamed `ApmContext` to `ContextMiddleware`

---

### V0.7.0

*Breaking changes*
db.GetChangedDBFields now return a slice of strings with all the changed fields instead of the part of a query string
with those fields. - because at injection time we sometimes need the tablename in front of the fieldname

db.InjectFields is changed into SQLInject which has different parameters. - it now works with options, which it more
flexible for extension.

The following functions all work with options. - db.MoreRecordsQuery - db.OneRecordQuery - db.FullRecordQuery -
db.InsertRecordQuery - db.UpdateRecordQuery

*New stuff*
To make joins, there are 2 extra options in the db package: - db.MoreRecordsJoinQuery - db.OneRecordJoinQuery

*Other changes*
The validator library got a small update

---

### V0.6.0

*Breaking changes*

Moved functionality to libGOAZ

> helpers.bodyToJSON -> libgoazConv.bodyToJSON

> helpers.FillStruct -> libgoazStructs.FillStruct

> helpers.StructToMap -> libgoazStructs.StructToMap

> structhelpers.GetTags -> libgoazStructs.GetTags

> helpers.JSONToMap -> libgoazConv.JSONToMap

> helpers.JoinErrors -> libgoazErrors.JoinErrors

Removed functionality

> helpers.HashHelper

- was already depricated in favor of keyhasher

Moved functionality internally

> structhelpers.GetChangedDbFields -> db.GetChangedDbFields

---

### V0.4.0

*New stuff*

The webtest have an extra helper function:

> InitTest

- used for testing a handler more directly.

This means, there is no call to a webserver, no router needed, just a handler-function you want to test.

As example:

```
	asserts := libgoazAsserts.New(t)
	queryParameters := map[string]string{"OrderBy": "wrong"}
	resp, req, params := web.InitTest(t, queryParameters)

	availabilityHandler.GetMoreRecords(resp, req, params)
	result := strings.TrimSpace(resp.Body.String())
	asserts.Contains(`invalid OrderBy 'wrong' is used`, result)
```

*Other changes*

removed linting errors (long lines) in webtest re-added @tput reset with 2 linters in makefile

---

### V0.3.0

*Breaking changes*

- `[]JSONValidationError` was removed instead a more generic `map[field]string` is now used so the value from the
  validator can be passed through instead of mapped into a specific JSON error. This will help in the future when other
  presenters may be added

*New stuff*

- `DBTracing bool` was added to the DBconfig. This bool decides if APM DB tracing is enabled.

> GetTags

- get []fieldnames from a struct filtered on a specific tag.
- is a variadic function
    - you can give as many remove filters as you like,
      You can even give no remove filter
- See the tests how it can work for you.

> GetChangedDbFields

- gets a string with all fields that have been changed (for an optimal update statement)
- it needs 3 models:
    - first model is as it is stored in the database
    - second is an empty model
    - third is the model which you want to store

> InitDb

- a small improvement in this fix (which has nothing to do with webtests)
- it pings the database to know early when the database isn't up and running.

This gives you the advantage of stopping a failing test early, like this:

```
	err := start.InitDB()
	asserts.EqualsNil(err)
	if err != nil {
		t.FailNow()
	}
```

> validate is moved out of helpers

renamed the 2 functions to avoid stuttering:
ValidateQueryParameters -> QueryParameters ValidateRequestBody -> RequestBody

There was a small bug:
if a model with required fields was validated against empty query parameters, it would pass, but now it it handled as an
error.

#### package db

> Transactions

```
// you give the database connection which you want use for all the underlying database actions
err = capilaDB.Transaction(dbx, func() error {
    // put all database relation actions in here....
    // this can be more than one
})

```

> Prepared Statements -> PrepareSelect

- a prepared select statement to get a slice of records

> Prepared Statements -> PrepareGet is a prepared select statement to get a single record.

- a prepared select statement to get a single record

> Prepared Statements -> PrepareCreate

- a prepared select statement to create a single record

> Prepared Statements -> PrepareUpdate

- a prepared select statement to update a single record

> Log

- you can see what your queries do by using db.Log(true)
- it also dumps the model to understand better what values are used in a query
- you can set it at your database connection initialization, or at some test.
  (Remember, if you turn it on, it stays on)

> InjectFields

- InjectFields injects the model fields into the query.
- you can use this to have more generic queries, because the sql gets "injected" by the correct sql-part

```
SELECT <dbfields> FROM availability where deleted_at IS NULL and id = :id

INSERT INTO availability (<dbfields>) VALUES (<modelfields>) RETURNING id;
```

> Mock

- you need less code in your tests.

```
mock := capilaDB.MockStart()	// created a fake database connection
rows := capilaDB.MockRows(model)	// define rows that have the fields of the model
// fill in your first model
model.FirstName = textPointer("Arie")
model.LastName = textPointer("Aafjes")
model.Email = textPointer("aa@domein.nl")

rowValues := capilaDB.MockRowData(model)
rows.AddRow(rowValues...) // store in in the mocked result-set

// fill in your second model
model.FirstName = textPointer("bbb")
model.LastName = textPointer("bbb")
model.Email = textPointer("bbb@domein.nl")

rowValues = capilaDB.MockRowData(model)
rows.AddRow(rowValues...) // store in in the mocked result-set
rows.AddRow(rowValues...) // store a couple of more (of the same)
rows.AddRow(rowValues...)
rows.AddRow(rowValues...) // we now have 5 records

mock.ExpectPrepare("^SELECT .* where deleted_at IS NULL.*")
// your real code
err := mock.ExpectationsWereMet()
asserts.EqualsNil(err)
```

> StandardSQL -> MoreRecordsQuery

- gets the sql query for more records

> StandardSQL -> OneRecordQuery

- gets the sql query for one record

> StandardSQL -> FullRecordQuery

- gets the sql query for one complete record

> StandardSQL -> InsertRecordQuery

- gets the sql query for the insert of one record

> StandardSQL -> IUpdateRecordQuery

- gets the sql query for the update of one record

### 0.2.0 - Thomas Jefferson

*Breaking changes*

- `capilaHelpers.JSONResponse` was moved to the presenters because of the overlap with the JSON presenter. This means
  that the functionality needs to be changed from `capilaHelpers.JSONResponse` to `capilaPresenters.PresentJSON(writer)`
  . For example, this is the new way to use the presenters:

```go
    return func (writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
presenter = *capilaPresenters.PresentJSON(writer)

data, err := handler.Repository.GetAllHotels()
if err != nil {
presenter.OutputError(http.StatusBadRequest, err)
return
}
presenter.Body.Data = responseModel.TransformModel(data)
presenter.OutputJSON()
}
```

Creating a presenter struct gives us more power to change `_links` and `meta` in the JSON response.

- `isDBOpen` was removed from the Init function in the DB package because we want to allow multiple connections and not
  always return the first connection. This occurs when you have a ReadUser and a WriteUser.

- `capilaPresenters.ThrowJSONError` was removed because it was not very different from OutputJSON.
- `capilaPresenters.OutputJSON` will ( probably ) be removed in the next version of Capila. It has been given a '
  deprecation annotation'.
- `HashHelper` is now deprecated and will be removed in a new version of Capila.

*New stuff*

- HashID's were added to Capila.
- `capilaPresenters.OutputError` is a new shorthand function that returns a JSON error response.
- `capilaPresenters.OutputValidationErrors` is a new shorthand function that sets the status code to 400 and returns an
  array of `[]JSONValidationError` errors that consist of a key and a value.

*Other changes*

## branch : validate-struct-in-error-per-field

### Breaking changes

Because the frontend could use more specific errors (like: an error per field), there are 2 calls that have changed.
They also have a better matching name.

#### IsValidParams > ValidateQueryParameters

improvedQueryParameters, err := capilaValidate.IsValidParams(queryParameters, &model)
becomes:
fieldErrors, error := capilaValidate.ValidateQueryParameters(&queryParameters, &model)
attention:
the queryParameters variable is now a pointer (which get changed inside if needed)

#### IsValidModel > ValidateRequestBody

error := capilaValidate.IsValidModel(isNew,requestBody, &model)
becomes:
fieldErrors, error := capilaValidate.ValidateRequestBody(isNew, requestBody, &model)
attention:
the requestBody variable is now a pointer

Both now return an extra value in the vorm of a map[field]errortext. There's also the concatenated error as well.

Validations are part of the tag behind a struct.

> should be a valid email

```
EmailField string `validate:"email"`address
```

> should be a number greater then or equal to 10

```
NumberField int `validate:"gte=10"`
```

> should be a number less then or equal to 10

```
NumberField int `validate:"lte=10"`
```

> you can combine them too
> should be a number greater then or equal to 10, and less then or equal to 15

```
NumberField int `validate:"gte=10,lte=15"`
```

> should be a string with a max. length of 20

```
TextField string `validate:"max=20"`
```

> should be a string with a min. length of 20

```
TextField string `validate:"min=20"`
```

> you can combine them too
> should be a string with a min. length of 5 and a max length of 10

```
PostalCodeField string `validate:"min=5,max=10"`
```

> should be a present and not have its zerovalue

```
SomeTextField string `validate:"required"`
SomeNumberField int `validate:"required"`
```

> Should be empty or a min. length of 10

```
TextField string `validate:"omitempty,min=10"`
```

## branch : feature/improve-db

### Breaking changes

db.Init is needed per database connection. If the application have a database READ user and a WRITE user, you need to
call this twice.

### New stuff

#### structhelpers package has struct related functions:

> GetTags

- get []fieldnames from a struct filtered on a specific tag.
- is a variadic function
    - you can give as many remove filters as you like,
      You can even give no remove filter
- See the tests how it can work for you.

> GetChangedDbFields

- gets a string with all fields that have been changed (for an optimal update statement)
- it needs 3 models:
    - first model is as it is stored in the database
    - second is an empty model
    - third is the model which you want to store

#### package db

> Transactions

```
	// you give the database connection which you want use for all the underlying database actions
	err = capilaDB.Transaction(dbx, func() error {
		// put all database relation actions in here....
		// this can be more than one
	})

```

> Prepared Statements -> PrepareSelect

- a prepared select statement to get a slice of records

> Prepared Statements -> PrepareGet is a prepared select statement to get a single record.

- a prepared select statement to get a single record

> Prepared Statements -> PrepareCreate

- a prepared select statement to create a single record

> Prepared Statements -> PrepareUpdate

- a prepared select statement to update a single record

> Log

- you can see what your queries do by using db.Log(true)
- it also dumps the model to understand better what values are used in a query
- you can set it at your database connection initialization, or at some test.
  (Remember, if you turn it on, it stays on)

> InjectFields

- InjectFields injects the model fields into the query.
- you can use this to have more generic queries, because the sql gets "injected" by the correct sql-part

```
SELECT <dbfields> FROM availability where deleted_at IS NULL and id = :id

INSERT INTO availability (<dbfields>) VALUES (<modelfields>) RETURNING id;
```

> Mock

- you need less code in your tests.

```
	mock := capilaDB.MockStart()	// created a fake database connection
	rows := capilaDB.MockRows(model)	// define rows that have the fields of the model

	// fill in your first model
	model.FirstName = textPointer("Arie")
	model.LastName = textPointer("Aafjes")
	model.Email = textPointer("aa@domein.nl")

	rowValues := capilaDB.MockRowData(model)
	rows.AddRow(rowValues...) // store in in the mocked result-set

	// fill in your second model
	model.FirstName = textPointer("bbb")
	model.LastName = textPointer("bbb")
	model.Email = textPointer("bbb@domein.nl")

	rowValues = capilaDB.MockRowData(model)
	rows.AddRow(rowValues...) // store in in the mocked result-set
	rows.AddRow(rowValues...) // store a couple of more (of the same)
	rows.AddRow(rowValues...)
	rows.AddRow(rowValues...) // we now have 5 records

	mock.ExpectPrepare("^SELECT .* where deleted_at IS NULL.*")
	// your real code
	err := mock.ExpectationsWereMet()
	asserts.EqualsNil(err)
```

> StandardSQL -> MoreRecordsQuery

- gets the sql query for more records

> StandardSQL -> OneRecordQuery

- gets the sql query for one record

> StandardSQL -> FullRecordQuery

- gets the sql query for one complete record

> StandardSQL -> InsertRecordQuery

- gets the sql query for the insert of one record

> StandardSQL -> IUpdateRecordQuery

- gets the sql query for the update of one record

---

### V0.1.0 - John Adams

*Breaking changes*

- The Authenticate middleware that contained JWT authentication `func Authentication` has been renamed
  to `func JWTAuthentication` because of additional Authentication middleware. This caused confusion because
  Authentication is not specific enough.
- InvalidDutchEmail / `var InvalidDutchEmail` was removed from this package as the error was too specific for the Capila
  library. Please refrain from adding errors that will most likely not be used again in other projects. In the future
  more of these errors will most likely be changed into a more generic format.
- Added `DBSchema   string` to the `type DBConnector` it is required to add the field but it can be an empty string.
  Only PostgresSQL is supported now.
- Added `statusCode int` parameter to `func OutputJSON`

*New stuff*

- Capila CLI for migrations. See the README.md on how to implement it. More cool stuff will be added later.
- `func BasicAuthentication` middleware was added.
- `func RateLimit` (Rate limiting) middleware was added.
- `func ThrowJSONError(writer http.ResponseWriter, statusCode int, errorMessage string)` was added for a generic way to
  handle errors and give JSON feedback.
