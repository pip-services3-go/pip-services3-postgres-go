package persistence

///* @module persistence */
// const _ = require('lodash');
// const async = require('async');

// import { IReferenceable } from 'pip-services3-commons-node';
// import { IUnreferenceable } from 'pip-services3-commons-node';
// import { IReferences } from 'pip-services3-commons-node';
// import { IConfigurable } from 'pip-services3-commons-node';
// import { IOpenable } from 'pip-services3-commons-node';
// import { ICleanable } from 'pip-services3-commons-node';
// import { ConfigParams } from 'pip-services3-commons-node';
// import { PagingParams } from 'pip-services3-commons-node';
// import { DataPage } from 'pip-services3-commons-node';
// import { ConnectionException } from 'pip-services3-commons-node';
// import { InvalidStateException } from 'pip-services3-commons-node';
// import { DependencyResolver } from 'pip-services3-commons-node';
// import { LongConverter } from 'pip-services3-commons-node';
// import { CompositeLogger } from 'pip-services3-components-node';

// import { PostgresConnection } from './PostgresConnection';

///*
//  * Abstract persistence component that stores data in PostgreSQL using plain driver.
//  *
//  * This is the most basic persistence component that is only
//  * able to store data items of any type. Specific CRUD operations
//  * over the data items must be implemented in child classes by
//  * accessing this._db or this._collection properties.
//  *
//  * ### Configuration parameters ###
//  *
//  * - collection:                  (optional) PostgreSQL collection name
//  * - connection(s):
//  *   - discovery_key:             (optional) a key to retrieve the connection from [[https://rawgit.com/pip-services-node/pip-services3-components-node/master/doc/api/interfaces/connect.idiscovery.html IDiscovery]]
//  *   - host:                      host name or IP address
//  *   - port:                      port number (default: 27017)
//  *   - uri:                       resource URI or connection string with all parameters in it
//  * - credential(s):
//  *   - store_key:                 (optional) a key to retrieve the credentials from [[https://rawgit.com/pip-services-node/pip-services3-components-node/master/doc/api/interfaces/auth.icredentialstore.html ICredentialStore]]
//  *   - username:                  (optional) user name
//  *   - password:                  (optional) user password
//  * - options:
//  *   - connect_timeout:      (optional) number of milliseconds to wait before timing out when connecting a new client (default: 0)
//  *   - idle_timeout:         (optional) number of milliseconds a client must sit idle in the pool and not be checked out (default: 10000)
//  *   - max_pool_size:        (optional) maximum number of clients the pool should contain (default: 10)
//  *
//  * ### References ###
//  *
//  * - \*:logger:\*:\*:1.0           (optional) [[https://rawgit.com/pip-services-node/pip-services3-components-node/master/doc/api/interfaces/log.ilogger.html ILogger]] components to pass log messages
//  * - \*:discovery:\*:\*:1.0        (optional) [[https://rawgit.com/pip-services-node/pip-services3-components-node/master/doc/api/interfaces/connect.idiscovery.html IDiscovery]] services
//  * - \*:credential-store:\*:\*:1.0 (optional) Credential stores to resolve credentials
//  *
//  * ### Example ###
//  *
//  *     class MyPostgresPersistence extends PostgresPersistence<MyData> {
//  *
//  *       public constructor() {
//  *           base("mydata");
//  *       }
//  *
//  *       public getByName(correlationId: string, name: string, callback: (err, item) => void): void {
//  *         let criteria = { name: name };
//  *         this._model.findOne(criteria, callback);
//  *       });
//  *
//  *       public set(correlatonId: string, item: MyData, callback: (err) => void): void {
//  *         let criteria = { name: item.name };
//  *         let options = { upsert: true, new: true };
//  *         this._model.findOneAndUpdate(criteria, item, options, callback);
//  *       }
//  *
//  *     }
//  *
//  *     let persistence = new MyPostgresPersistence();
//  *     persistence.configure(ConfigParams.fromTuples(
//  *         "host", "localhost",
//  *         "port", 27017
//  *     ));
//  *
//  *     persitence.open("123", (err) => {
//  *          ...
//  *     });
//  *
//  *     persistence.set("123", { name: "ABC" }, (err) => {
//  *         persistence.getByName("123", "ABC", (err, item) => {
//  *             console.log(item);                   // Result: { name: "ABC" }
//  *         });
//  *     });
//  */
// export class PostgresPersistence<T> implements IReferenceable, IUnreferenceable, IConfigurable, IOpenable, ICleanable {

//     private static _defaultConfig: ConfigParams = ConfigParams.fromTuples(
//         "collection", null,
//         "dependencies.connection", "*:connection:postgres:*:1.0",

//         // connections.*
//         // credential.*

//         "options.max_pool_size", 2,
//         "options.keep_alive", 1,
//         "options.connect_timeout", 5000,
//         "options.auto_reconnect", true,
//         "options.max_page_size", 100,
//         "options.debug", true
//     );

//     private _config: ConfigParams;
//     private _references: IReferences;
//     private _opened: boolean;
//     private _localConnection: boolean;
//     private _autoObjects: string[] = [];

//    /*
//      * The dependency resolver.
//      */
//     protected _dependencyResolver: DependencyResolver = new DependencyResolver(PostgresPersistence._defaultConfig);
//    /*
//      * The logger.
//      */
//     protected _logger: CompositeLogger = new CompositeLogger();

//    /*
//      * The PostgreSQL connection component.
//      */
//     protected _connection: PostgresConnection;

//    /*
//      * The PostgreSQL connection pool object.
//      */
//     protected _client: any;
//    /*
//      * The PostgreSQL database name.
//      */
//     protected _databaseName: string;
//    /*
//      * The PostgreSQL table object.
//      */
//     protected _tableName: string;

//     protected _maxPageSize: number = 100;

//    /*
//      * Creates a new instance of the persistence component.
//      *
//      * - tableName    (optional) a table name.
//      */
//     public constructor(tableName?: string) {
//         this._tableName = tableName;
//     }

//    /*
//      * Configures component by passing configuration parameters.
//      *
//      * - config    configuration parameters to be set.
//      */
//     public configure(config: ConfigParams): void {
//         config = config.setDefaults(PostgresPersistence._defaultConfig);
//         this._config = config;

//         this._dependencyResolver.configure(config);

//         this._tableName = config.getAsStringWithDefault("collection", this._tableName);
//         this._tableName = config.getAsStringWithDefault("table", this._tableName);
//         this._maxPageSize = config.getAsIntegerWithDefault("options.max_page_size", this._maxPageSize);
//     }

//    /*
// 	 * Sets references to dependent components.
// 	 *
// 	 * - references 	references to locate the component dependencies.
//      */
//     public setReferences(references: IReferences): void {
//         this._references = references;
//         this._logger.setReferences(references);

//         // Get connection
//         this._dependencyResolver.setReferences(references);
//         this._connection = this._dependencyResolver.getOneOptional('connection');
//         // Or create a local one
//         if (this._connection == null) {
//             this._connection = this.createConnection();
//             this._localConnection = true;
//         } else {
//             this._localConnection = false;
//         }
//     }

//    /*
// 	 * Unsets (clears) previously set references to dependent components.
//      */
//     public unsetReferences(): void {
//         this._connection = null;
//     }

//     private createConnection(): PostgresConnection {
//         let connection = new PostgresConnection();

//         if (this._config)
//             connection.configure(this._config);

//         if (this._references)
//             connection.setReferences(this._references);

//         return connection;
//     }

//    /*
//      * Adds index definition to create it on opening
//      * - keys index keys (fields)
//      * - options index options
//      */
//     protected ensureIndex(name: string, keys: any, options?: any): void {
//         let builder = "CREATE";
//         options = options || {};

//         if (options.unique) {
//             builder += " UNIQUE";
//         }

//         builder += " INDEX IF NOT EXISTS " + name + " ON " + this.quoteIdentifier(this._tableName);

//         if (options.type) {
//             builder += " " + options.type;
//         }

//         let fields = "";
//         for (let key in keys) {
//             if (fields != "") fields += ", ";
//             fields += this.quoteIdentifier(key);
//             let asc = keys[key];
//             if (!asc) fields += " DESC";
//         }

//         builder += "(" + fields + ")";

//         this.autoCreateObject(builder);
//     }

//    /*
//      * Adds index definition to create it on opening
//      * - dmlStatement DML statement to autocreate database object
//      */
//     protected autoCreateObject(dmlStatement: string): void {
//         this._autoObjects.push(dmlStatement);
//     }

//    /*
//      * Converts object value from internal to public format.
//      *
//      * - value     an object in internal format to convert.
//      * Returns converted object in public format.
//      */
//     protected convertToPublic(value: any): any {
//         return value;
//     }

//    /*
//      * Convert object value from public to internal format.
//      *
//      * - value     an object in public format to convert.
//      * Returns converted object in internal format.
//      */
//     protected convertFromPublic(value: any): any {
//         return value;
//     }

//     protected quoteIdentifier(value: string): string {
//         if (value == null || value == "") return value;

//         if (value[0] == '"') return value;

//         return '"' + value + '"';
//     }

//    /*
// 	 * Checks if the component is opened.
// 	 *
// 	 * Returns true if the component has been opened and false otherwise.
//      */
//     public isOpen(): boolean {
//         return this._opened;
//     }

//    /*
// 	 * Opens the component.
// 	 *
// 	 * - correlationId 	(optional) transaction id to trace execution through call chain.
//      * - callback 			callback function that receives error or null no errors occured.
//      */
//     public open(correlationId: string, callback?: (err: any) => void): void {
//     	if (this._opened) {
//             callback(null);
//             return;
//         }

//         if (this._connection == null) {
//             this._connection = this.createConnection();
//             this._localConnection = true;
//         }

//         let openCurl = (err) => {
//             if (err == null && this._connection == null) {
//                 err = new InvalidStateException(correlationId, 'NO_CONNECTION', 'PostgreSQL connection is missing');
//             }

//             if (err == null && !this._connection.isOpen()) {
//                 err = new ConnectionException(correlationId, "CONNECT_FAILED", "PostgreSQL connection is not opened");
//             }

//             this._opened = false;

//             if (err) {
//                 if (callback) callback(err);
//             } else {
//                 this._client = this._connection.getConnection();
//                 this._databaseName = this._connection.getDatabaseName();

//                 // Recreate objects
//                 this.autoCreateObjects(correlationId, (err) => {
//                     if (err) {
//                         this._client == null;
//                         err = new ConnectionException(correlationId, "CONNECT_FAILED", "Connection to postgres failed").withCause(err);
//                     } else {
//                         this._opened = true;
//                         this._logger.debug(correlationId, "Connected to postgres database %s, collection %s", this._databaseName, this.quoteIdentifier(this._tableName));
//                     }

//                     if (callback) callback(err);
//                 });
//             }
//         };

//         if (this._localConnection) {
//             this._connection.open(correlationId, openCurl);
//         } else {
//             openCurl(null);
//         }

//     }

//    /*
// 	 * Closes component and frees used resources.
// 	 *
// 	 * - correlationId 	(optional) transaction id to trace execution through call chain.
//      * - callback 			callback function that receives error or null no errors occured.
//      */
//     public close(correlationId: string, callback?: (err: any) => void): void {
//     	if (!this._opened) {
//             callback(null);
//             return;
//         }

//         if (this._connection == null) {
//             callback(new InvalidStateException(correlationId, 'NO_CONNECTION', 'Postgres connection is missing'));
//             return;
//         }

//         let closeCurl = (err) => {
//             this._opened = false;
//             this._client = null;

//             if (callback) callback(err);
//         }

//         if (this._localConnection) {
//             this._connection.close(correlationId, closeCurl);
//         } else {
//             closeCurl(null);
//         }
//     }

//    /*
// 	 * Clears component state.
// 	 *
// 	 * - correlationId 	(optional) transaction id to trace execution through call chain.
//      * - callback 			callback function that receives error or null no errors occured.
//      */
//     public clear(correlationId: string, callback?: (err: any) => void): void {
//         // Return error if collection is not set
//         if (this._tableName == null) {
//             if (callback) callback(new Error('Table name is not defined'));
//             return;
//         }

//         let query = "DELETE FROM " + this.quoteIdentifier(this._tableName);

//         this._client.query(query, (err, result) => {
//             if (err) {
//                 err = new ConnectionException(correlationId, "CONNECT_FAILED", "Connection to postgres failed")
//                     .withCause(err);
//             }

//             if (callback) callback(err);
//         });
//     }

//     protected autoCreateObjects(correlationId: string, callback: (err: any) => void): void {
//         if (this._autoObjects == null || this._autoObjects.length == 0) {
//             callback(null);
//             return null;
//         }

//         // Check if table exist to determine weither to auto create objects
//         let query = "SELECT to_regclass('" + this._tableName + "')";
//         this._client.query(query, (err, result) => {
//             if (err) {
//                 callback(err);
//                 return;
//             }

//             // If table already exists then exit
//             if (result.rows && result.rows.length > 0 && result.rows[0].to_regclass != null) {
//                 callback(err);
//                 return;
//             }

//             this._logger.debug(correlationId, 'Table ' + this._tableName + ' does not exist. Creating database objects...');

//             // Run all DML commands
//             async.eachSeries(this._autoObjects, (dml, callback) => {
//                 this._client.query(dml, (err, result) => {
//                     if (err) {
//                         this._logger.error(correlationId, err, 'Failed to autocreate database object');
//                     }
//                     callback(err);
//                 });
//             }, callback);
//         });
//     }

//    /*
//      * Generates a list of column names to use in SQL statements like: "column1,column2,column3"
//      * - values an array with column values or a key-value map
//      * Returns a generated list of column names
//      */
//     protected generateColumns(values: any): string {
//         values = !_.isArray(values) ? _.keys(values) : values;

//         let result = "";
//         for (let value of values) {
//             if (result != "") result += ",";
//             result += this.quoteIdentifier(value);
//         }

//         return result;
//     }

//    /*
//      * Generates a list of value parameters to use in SQL statements like: "$1,$2,$3"
//      * - values an array with values or a key-value map
//      * Returns a generated list of value parameters
//      */
//     protected generateParameters(values: any): string {
//         values = !_.isArray(values) ? _.keys(values) : values;

//         let index = 1;
//         let result = "";
//         for (let value of values) {
//             if (result != "") result += ",";
//             result += "$" + index;
//             index++;
//         }

//         return result;
//     }

//    /*
//      * Generates a list of column sets to use in UPDATE statements like: column1=$1,column2=$2
//      * - values a key-value map with columns and values
//      * Returns a generated list of column sets
//      */
//     protected generateSetParameters(values: any): string {
//         let result = "";
//         let index = 1;
//         for (let column in values) {
//             if (result != "") result += ",";
//             result += this.quoteIdentifier(column) + "=$" + index;
//             index++;
//         }

//         return result;
//     }

//    /*
//      * Generates a list of column parameters
//      * - values a key-value map with columns and values
//      * Returns a generated list of column values
//      */
//     protected generateValues(values: any): any[] {
//         return _.values(values);
//     }

//    /*
//      * Gets a page of data items retrieved by a given filter and sorted according to sort parameters.
//      *
//      * This method shall be called by a public getPageByFilter method from child class that
//      * receives FilterParams and converts them into a filter function.
//      *
//      * - correlationId     (optional) transaction id to trace execution through call chain.
//      * - filter            (optional) a filter JSON object
//      * - paging            (optional) paging parameters
//      * - sort              (optional) sorting JSON object
//      * - select            (optional) projection JSON object
//      * - callback          callback function that receives a data page or error.
//      */
//     protected getPageByFilter(correlationId: string, filter: any, paging: PagingParams,
//         sort: any, select: any, callback: (err: any, items: DataPage<T>) => void): void {

//         select = select && !_.isEmpty(select) ? select : "*"
//         let query = "SELECT " + select + " FROM " + this.quoteIdentifier(this._tableName);

//         // Adjust max item count based on configuration
//         paging = paging || new PagingParams();
//         let skip = paging.getSkip(-1);
//         let take = paging.getTake(this._maxPageSize);
//         let pagingEnabled = paging.total;

//         if (filter && filter != "")
//             query += " WHERE " + filter;

//         if (sort && !_.isEmpty(sort)) query += " ORDER BY " + sort;

//         if (skip >= 0) query += " OFFSET " + skip;
//         query += " LIMIT " + take;

//         this._client.query(query, (err, result) => {
//             err = err || null;
//             if (err) {
//                 callback(err, null);
//                 return;
//             }

//             let items = result.rows;

//             if (items != null)
//                 this._logger.trace(correlationId, "Retrieved %d from %s", items.length, this._tableName);

//             items = _.map(items, this.convertToPublic);

//             if (pagingEnabled) {
//                 let query = 'SELECT COUNT(*) AS count FROM ' + this.quoteIdentifier(this._tableName);
//                 if (filter != null && filter != "")
//                     query += " WHERE " + filter;

//                 this._client.query(query, (err, result) => {
//                     err = err || null;
//                     if (err) {
//                         callback(err, null);
//                         return;
//                     }

//                     let count = result.rows && result.rows.length == 1
//                         ? LongConverter.toLong(result.rows[0].count) : 0;
//                     let page = new DataPage<T>(items, count);
//                     callback(null, page);
//                 });
//             } else {
//                 let page = new DataPage<T>(items);
//                 callback(null, page);
//             }
//         });
//     }

//    /*
//      * Gets a number of data items retrieved by a given filter.
//      *
//      * This method shall be called by a public getCountByFilter method from child class that
//      * receives FilterParams and converts them into a filter function.
//      *
//      * - correlationId     (optional) transaction id to trace execution through call chain.
//      * - filter            (optional) a filter JSON object
//      * - callback          callback function that receives a data page or error.
//      */
//     protected getCountByFilter(correlationId: string, filter: any,
//         callback: (err: any, count: number) => void): void {

//         let query = 'SELECT COUNT(*) AS count FROM ' + this.quoteIdentifier(this._tableName);
//         if (filter && filter != "")
//             query += " WHERE " + filter;

//         this._client.query(query, (err, result) => {
//             err = err || null;
//             if (err) {
//                 callback(err, null);
//                 return;
//             }

//             let count = result.rows && result.rows.length == 1
//                 ? LongConverter.toLong(result.rows[0].count) : 0;

//             if (count != null)
//                 this._logger.trace(correlationId, "Counted %d items in %s", count, this._tableName);

//             callback(null, count);
//         });
//     }

//    /*
//      * Gets a list of data items retrieved by a given filter and sorted according to sort parameters.
//      *
//      * This method shall be called by a public getListByFilter method from child class that
//      * receives FilterParams and converts them into a filter function.
//      *
//      * - correlationId    (optional) transaction id to trace execution through call chain.
//      * - filter           (optional) a filter JSON object
//      * - paging           (optional) paging parameters
//      * - sort             (optional) sorting JSON object
//      * - select           (optional) projection JSON object
//      * - callback         callback function that receives a data list or error.
//      */
//     protected getListByFilter(correlationId: string, filter: any, sort: any, select: any,
//         callback: (err: any, items: T[]) => void): void {

//         select = select && !_.isEmpty(select) ? select : "*"
//         let query = "SELECT " + select + " FROM " + this.quoteIdentifier(this._tableName);

//         if (filter && filter != "")
//             query += " WHERE " + filter;

//         if (sort && !_.isEmpty(sort)) query += " ORDER BY " + sort;

//         this._client.query(query, (err, result) => {
//             err = err || null;
//             if (err) {
//                 callback(err, null);
//                 return;
//             }

//             let items = result.rows;

//             if (items != null)
//                 this._logger.trace(correlationId, "Retrieved %d from %s", items.length, this._tableName);

//             items = _.map(items, this.convertToPublic);
//             callback(null, items);
//         });
//     }

//    /*
//      * Gets a random item from items that match to a given filter.
//      *
//      * This method shall be called by a public getOneRandom method from child class that
//      * receives FilterParams and converts them into a filter function.
//      *
//      * - correlationId     (optional) transaction id to trace execution through call chain.
//      * - filter            (optional) a filter JSON object
//      * - callback          callback function that receives a random item or error.
//      */
//     protected getOneRandom(correlationId: string, filter: any, callback: (err: any, item: T) => void): void {
//         let query = 'SELECT COUNT(*) AS count FROM ' + this.quoteIdentifier(this._tableName);
//         if (filter && filter != "")
//             query += " WHERE " + filter;

//         this._client.query(query, (err, result) => {
//             err = err || null;
//             if (err) {
//                 callback(err, null);
//                 return;
//             }

//             let query = "SELECT * FROM " + this.quoteIdentifier(this._tableName);

//             if (filter && filter != "")
//                 query += " WHERE " + filter;

//             let count = result.rows && result.rows.length == 1 ? result.rows[0].count : 0;
//             let pos = _.random(0, count - 1);
//             query += " OFFSET " + pos + " LIMIT 1";

//             this._client.query(query, (err, result) => {
//                 err = err || null;

//                 let items = result.rows;
//                 let item = (items != null && items.length > 0) ? items[0] : null;

//                 if (item == null)
//                     this._logger.trace(correlationId, "Random item wasn't found from %s", this._tableName);
//                 else
//                     this._logger.trace(correlationId, "Retrieved random item from %s", this._tableName);

//                 item = this.convertToPublic(item);
//                 callback(err, item);
//             });
//         });
//     }

//    /*
//      * Creates a data item.
//      *
//      * - correlation_id    (optional) transaction id to trace execution through call chain.
//      * - item              an item to be created.
//      * - callback          (optional) callback function that receives created item or error.
//      */
//     public create(correlationId: string, item: T, callback?: (err: any, item: T) => void): void {
//         if (item == null) {
//             callback(null, null);
//             return;
//         }

//         let row = this.convertFromPublic(item);
//         let columns = this.generateColumns(row);
//         let params = this.generateParameters(row);
//         let values = this.generateValues(row);

//         let query = "INSERT INTO " + this.quoteIdentifier(this._tableName) + " (" + columns + ") VALUES (" + params + ") RETURNING *";

//         this._client.query(query, values, (err, result) => {
//             err = err || null;
//             if (!err)
//                 this._logger.trace(correlationId, "Created in %s with id = %s", this._tableName, row.id);

//             let newItem = result && result.rows && result.rows.length == 1
//                 ? this.convertToPublic(result.rows[0]) : null;
//             callback(err, newItem);
//         });
//     }

//    /*
//      * Deletes data items that match to a given filter.
//      *
//      * This method shall be called by a public deleteByFilter method from child class that
//      * receives FilterParams and converts them into a filter function.
//      *
//      * - correlationId     (optional) transaction id to trace execution through call chain.
//      * - filter            (optional) a filter JSON object.
//      * - callback          (optional) callback function that receives error or null for success.
//      */
//     public deleteByFilter(correlationId: string, filter: string, callback?: (err: any) => void): void {
//         let query = "DELETE FROM " + this.quoteIdentifier(this._tableName);
//         if (filter != null && filter != "")
//             query += " WHERE " + filter;

//         this._client.query(query, (err, result) => {
//             let count = result ? result.rowCount : 0;

//             err = err || null;
//             if (!err)
//                 this._logger.trace(correlationId, "Deleted %d items from %s", count, this._tableName);

//             if (callback) callback(err);
//         });
//     }

// }
