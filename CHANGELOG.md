# <img src="https://uploads-ssl.webflow.com/5ea5d3315186cf5ec60c3ee4/5edf1c94ce4c859f2b188094_logo.svg" alt="Pip.Services Logo" width="200"> <br/> PostgreSQL components for Golang Changelog

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

