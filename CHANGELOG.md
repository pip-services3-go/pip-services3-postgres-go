# PostgreSQL components for Pip.Services in Node.js Changelog

## <a name="3.4.0"></a> 3.4.0 (2020-06-09) 

### Features
* Moved some CRUD operations from IdentifiablePostgresPersistence to PostgresPersistence

## <a name="3.3.0"></a> 3.3.0 (2020-05-18) 

### Features
* Added getCountByFilter

## <a name="3.2.6"></a> 3.2.6 (2020-04-10) 

### Bug fixes
* Fixed logging message of the collection name in IdentifiablePostgresPersistence

## <a name="3.2.0"></a> 3.2.0 (2019-11-05) 

### Features
* Added PostgresConnection
* Added PostgresPersistence.ensureIndex()

## <a name="3.1.0"></a> 3.1.0 (2019-05-20) 

### Breaking Changes
* Reimplemented persistence using plain PostgreSQL driver
* Mongoose-based persistence moved to a separate package

## <a name="3.0.0"></a> 3.0.0 (2018-08-21) 

### Breaking Changes
* Moved to a separate package

## <a name="1.2.0"></a> 1.2.0 (2018-08-10) 

### Features
* Added RedisCache
* Added RedisLock

## <a name="1.1.0"></a> 1.1.0 (2018-03-26) 

### Features
* Added PrometheusCounters and PrometheusMetricsService
* Added labels to PrometheusCounters and PrometheusMetricsService

## <a name="1.0.0"></a> 1.0.0 (2018-03-20) 

### Features
* **memcached** Added MemcachedCache
* **memcached** Added MemcachedLock
* **fluentd** Added FluentdLogger
* **elasticsearch** Added ElasticSearchLogger

