# <img src="https://uploads-ssl.webflow.com/5ea5d3315186cf5ec60c3ee4/5edf1c94ce4c859f2b188094_logo.svg" alt="Pip.Services Logo" width="200"> <br/> PostgreSQL components for Golang Changelog

## <a name="1.2.10"></a> 1.2.10 (2022-09-26)
### Bug fixing
* Fixed EnsureIndex

## <a name="1.2.9"></a> 1.2.9 (2022-06-16)
### Bug fixing
* Fixed query builder for total calculation in GetPageByFilter method

## <a name="1.2.8"></a> 1.2.8 (2022-06-02)
### Bug fixing
* Fixed return total value

## <a name="1.2.7"></a> 1.2.7 (2021-07-01)
### Features
* Change method naming QuotedTableNameWithSchema -> QuotedTableName
## <a name="1.2.7"></a> 1.2.7 (2021-05-19)
### Bug fixing
* Fix GetOneRandom method

## <a name="1.2.5"></a> 1.2.5 (2021-04-27)
### Bug fixing
* Fix parameter index converting in GenerateSetParameters and GenerateParameters

## <a name="1.2.4"></a> 1.2.4 (2021-04-27)

### Features
* Add ability to use custom PostgreSQL schema

## <a name="1.2.3"></a> 1.2.3 (2021-04-16)

### Bug fixing
* Update dependencies for fix errors in clone object

## <a name="1.2.2"></a> 1.2.2 (2021-04-15) 

### Bug fixing
* Fixed  composeUri in PostgresConnectionResolver

## <a name="1.2.1"></a> 1.2.1 (2021-04-12) 

### Bug fixing
* Fixed catching parsing config error in Open method in PostgresConnection

## <a name="1.2.0"></a> 1.2.0 (2021-04-03) 

### Features
* Moved PostgresConnection to connect package
* Added IPostgresPersistenceOverride interface to overload virtual methods

### Breaking changes
* Method autoCreateObject is deprecated and shall be renamed to ensureSchema

## <a name="1.1.0"></a> 1.1.0 (2021-02-18) 

### Features
* Renamed autoCreateObject to ensureSchema
* Added defineSchema method that shall be overriden in child classes
* Added clearSchema method

### Breaking changes
* Method autoCreateObject is deprecated and shall be renamed to ensureSchema

## <a name="1.0.2"></a> 1.0.2 (2020-12-11) 

### Features
* Update dependencies

## <a name="1.0.1"></a> 1.0.1 (2020-11-12) 

### Features
* Changed convert data methods

## <a name="1.0.0"></a> 1.0.0 (2020-11-06) 

Initial public release

### Features
* **build** standard factory for constructing components
* **connect** instruments for configuring connections to the database
* **persistence** abstract classes for working with the database

