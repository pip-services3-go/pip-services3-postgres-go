package persistence

///* @module persistence */
///* @hidden */
// const _ = require('lodash');

// import { AnyValueMap } from 'pip-services3-commons-node';
// import { IIdentifiable } from 'pip-services3-commons-node';
// import { IdGenerator } from 'pip-services3-commons-node';

// import { IdentifiablePostgresPersistence } from './IdentifiablePostgresPersistence';

///*
//  * Abstract persistence component that stores data in PostgreSQL in JSON or JSONB fields
//  * and implements a number of CRUD operations over data items with unique ids.
//  * The data items must implement IIdentifiable interface.
//  *
//  * The JSON table has only two fields: id and data.
//  *
//  * In basic scenarios child classes shall only override [[getPageByFilter]],
//  * [[getListByFilter]] or [[deleteByFilter]] operations with specific filter function.
//  * All other operations can be used out of the box.
//  *
//  * In complex scenarios child classes can implement additional operations by
//  * accessing this._collection and this._model properties.

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
//  * - \*:logger:\*:\*:1.0           (optional) [[https://rawgit.com/pip-services-node/pip-services3-components-node/master/doc/api/interfaces/log.ilogger.html ILogger]] components to pass log messages components to pass log messages
//  * - \*:discovery:\*:\*:1.0        (optional) [[https://rawgit.com/pip-services-node/pip-services3-components-node/master/doc/api/interfaces/connect.idiscovery.html IDiscovery]] services
//  * - \*:credential-store:\*:\*:1.0 (optional) Credential stores to resolve credentials
//  *
//  * ### Example ###
//  *
//  *     class MyPostgresPersistence extends IdentifiablePostgresJsonPersistence<MyData, string> {
//  *
//  *     public constructor() {
//  *         base("mydata", new MyDataPostgresSchema());
//  *     }
//  *
//  *     private composeFilter(filter: FilterParams): any {
//  *         filter = filter || new FilterParams();
//  *         let criteria = [];
//  *         let name = filter.getAsNullableString('name');
//  *         if (name != null)
//  *             criteria.push({ name: name });
//  *         return criteria.length > 0 ? { $and: criteria } : null;
//  *     }
//  *
//  *     public getPageByFilter(correlationId: string, filter: FilterParams, paging: PagingParams,
//  *         callback: (err: any, page: DataPage<MyData>) => void): void {
//  *         base.getPageByFilter(correlationId, this.composeFilter(filter), paging, null, null, callback);
//  *     }
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
//  *         ...
//  *     });
//  *
//  *     persistence.create("123", { id: "1", name: "ABC" }, (err, item) => {
//  *         persistence.getPageByFilter(
//  *             "123",
//  *             FilterParams.fromTuples("name", "ABC"),
//  *             null,
//  *             (err, page) => {
//  *                 console.log(page.data);          // Result: { id: "1", name: "ABC" }
//  *
//  *                 persistence.deleteById("123", "1", (err, item) => {
//  *                    ...
//  *                 });
//  *             }
//  *         )
//  *     });
//  */
// export class IdentifiableJsonPostgresPersistence<T extends IIdentifiable<K>, K> extends IdentifiablePostgresPersistence<T, K> {
//    /*
//      * Creates a new instance of the persistence component.
//      *
//      * - collection    (optional) a collection name.
//      */
//     public constructor(tableName: string) {
//         super(tableName);
//     }

//    /*
//      * Adds DML statement to automatically create JSON(B) table
//      *
//      * - idType type of the id column (default: TEXT)
//      * - dataType type of the data column (default: JSONB)
//      */
//     protected ensureTable(idType: string = 'TEXT', dataType: string = 'JSONB') {
//         let query = "CREATE TABLE IF NOT EXISTS " + this.quoteIdentifier(this._tableName)
//             + " (\"id\" " + idType + " PRIMARY KEY, \"data\" " + dataType + ")";
//         this.autoCreateObject(query);
//     }

//    /*
//      * Converts object value from internal to public format.
//      *
//      * - value     an object in internal format to convert.
//      * Returns converted object in public format.
//      */
//     protected convertToPublic(value: any): any {
//         if (value == null) return null;
//          return value.data;
//     }

//    /*
//      * Convert object value from public to internal format.
//      *
//      * - value     an object in public format to convert.
//      * Returns converted object in internal format.
//      */
//     protected convertFromPublic(value: any): any {
//         if (value == null) return null;
//         let result: any = {
//             id: value.id,
//             data: value
//         };
//         return result;
//     }

//    /*
//      * Updates only few selected fields in a data item.
//      *
//      * - correlation_id    (optional) transaction id to trace execution through call chain.
//      * - id                an id of data item to be updated.
//      * - data              a map with fields to be updated.
//      * - callback          callback function that receives updated item or error.
//      */
//     public updatePartially(correlationId: string, id: K, data: AnyValueMap,
//         callback?: (err: any, item: T) => void): void {

//         if (data == null || id == null) {
//             if (callback) callback(null, null);
//             return;
//         }

//         let query = "UPDATE " + this.quoteIdentifier(this._tableName) + " SET \"data\"=\"data\"||$2 WHERE \"id\"=$1 RETURNING *";
//         let values = [id, data.getAsObject()];

//         this._client.query(query, values, (err, result) => {
//             err = err || null;
//             if (!err)
//                 this._logger.trace(correlationId, "Updated partially in %s with id = %s", this._tableName, id);

//             let newItem = result && result.rows && result.rows.length == 1
//                 ? this.convertToPublic(result.rows[0]) : null;

//             if (callback) callback(err, newItem);
//         });
//     }

// }
