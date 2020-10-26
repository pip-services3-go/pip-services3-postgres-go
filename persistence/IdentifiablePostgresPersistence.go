package persistence

///* @module persistence */
///* @hidden */
// const _ = require('lodash');

// import { AnyValueMap } from 'pip-services3-commons-node';
// import { IIdentifiable } from 'pip-services3-commons-node';
// import { IdGenerator } from 'pip-services3-commons-node';

// import { IWriter } from 'pip-services3-data-node';
// import { IGetter } from 'pip-services3-data-node';
// import { ISetter } from 'pip-services3-data-node';

// import { PostgresPersistence } from './PostgresPersistence';

///*
//  * Abstract persistence component that stores data in PostgreSQL
//  * and implements a number of CRUD operations over data items with unique ids.
//  * The data items must implement IIdentifiable interface.
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
//  *     class MyPostgresPersistence extends IdentifiablePostgresPersistence<MyData, string> {
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
// export class IdentifiablePostgresPersistence<T extends IIdentifiable<K>, K> extends PostgresPersistence<T>
//     implements IWriter<T, K>, IGetter<T, K>, ISetter<T> {

//    /*
//      * Creates a new instance of the persistence component.
//      *
//      * - collection    (optional) a collection name.
//      */
//     public constructor(tableName: string) {
//         super(tableName);

//         if (tableName == null)
//             throw new Error("Table name could not be null");
//     }

//    /*
//      * Converts the given object from the public partial format.
//      *
//      * - value     the object to convert from the public partial format.
//      * Returns the initial object.
//      */
//     protected convertFromPublicPartial(value: any): any {
//         return this.convertFromPublic(value);
//     }

//    /*
//      * Gets a list of data items retrieved by given unique ids.
//      *
//      * - correlationId     (optional) transaction id to trace execution through call chain.
//      * - ids               ids of data items to be retrieved
//      * - callback         callback function that receives a data list or error.
//      */
//     public getListByIds(correlationId: string, ids: K[],
//         callback: (err: any, items: T[]) => void): void {

//         let params = this.generateParameters(ids);
//         let query = "SELECT * FROM " + this.quoteIdentifier(this._tableName) + " WHERE \"id\" IN(" + params + ")";

//         this._client.query(query, ids, (err, result) => {
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
//      * Gets a data item by its unique id.
//      *
//      * - correlationId     (optional) transaction id to trace execution through call chain.
//      * - id                an id of data item to be retrieved.
//      * - callback          callback function that receives data item or error.
//      */
//     public getOneById(correlationId: string, id: K, callback: (err: any, item: T) => void): void {
//         let query = "SELECT * FROM " + this.quoteIdentifier(this._tableName) + " WHERE \"id\"=$1";
//         let params = [ id ];

//         this._client.query(query, params, (err, result) => {
//             err = err || null;

//             let item = result && result.rows ? result.rows[0] || null : null;

//             if (item == null)
//                 this._logger.trace(correlationId, "Nothing found from %s with id = %s", this._tableName, id);
//             else
//                 this._logger.trace(correlationId, "Retrieved from %s with id = %s", this._tableName, id);

//             item = this.convertToPublic(item);
//             callback(err, item);
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

//         // Assign unique id
//         let newItem: any = item;
//         if (newItem.id == null) {
//             newItem = _.clone(newItem);
//             newItem.id = item.id || IdGenerator.nextLong();
//         }

//         super.create(correlationId, newItem, callback);
//     }

//    /*
//      * Sets a data item. If the data item exists it updates it,
//      * otherwise it create a new data item.
//      *
//      * - correlation_id    (optional) transaction id to trace execution through call chain.
//      * - item              a item to be set.
//      * - callback          (optional) callback function that receives updated item or error.
//      */
//     public set(correlationId: string, item: T, callback?: (err: any, item: T) => void): void {
//         if (item == null) {
//             if (callback) callback(null, null);
//             return;
//         }

//         // Assign unique id
//         if (item.id == null) {
//             item = _.clone(item);
//             item.id = <any>IdGenerator.nextLong();
//         }

//         let row = this.convertFromPublic(item);
//         let columns = this.generateColumns(row);
//         let params = this.generateParameters(row);
//         let setParams = this.generateSetParameters(row);
//         let values = this.generateValues(row);

//         let query = "INSERT INTO " + this.quoteIdentifier(this._tableName) + " (" + columns + ")"
//             + " VALUES (" + params + ")"
//             + " ON CONFLICT (\"id\") DO UPDATE SET " + setParams + " RETURNING *";

//         this._client.query(query, values, (err, result) => {
//             err = err || null;
//             if (!err)
//                 this._logger.trace(correlationId, "Set in %s with id = %s", this._tableName, item.id);

//             let newItem = result && result.rows && result.rows.length == 1
//                 ? this.convertToPublic(result.rows[0]) : null;

//             if (callback) callback(err, newItem);
//         });
//     }

//    /*
//      * Updates a data item.
//      *
//      * - correlation_id    (optional) transaction id to trace execution through call chain.
//      * - item              an item to be updated.
//      * - callback          (optional) callback function that receives updated item or error.
//      */
//     public update(correlationId: string, item: T, callback?: (err: any, item: T) => void): void {
//         if (item == null || item.id == null) {
//             if (callback) callback(null, null);
//             return;
//         }

//         let row = this.convertFromPublic(item);
//         let params = this.generateSetParameters(row);
//         let values = this.generateValues(row);
//         values.push(item.id);

//         let query = "UPDATE " + this.quoteIdentifier(this._tableName)
//             + " SET " + params + " WHERE \"id\"=$" + values.length + " RETURNING *";

//         this._client.query(query, values, (err, result) => {
//             err = err || null;
//             if (!err)
//                 this._logger.trace(correlationId, "Updated in %s with id = %s", this._tableName, item.id);

//             let newItem = result && result.rows && result.rows.length == 1
//                 ? this.convertToPublic(result.rows[0]) : null;

//             if (callback) callback(err, newItem);
//         });
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

//         let row = this.convertFromPublicPartial(data.getAsObject());
//         let params = this.generateSetParameters(row);
//         let values = this.generateValues(row);
//         values.push(id);

//         let query = "UPDATE " + this.quoteIdentifier(this._tableName)
//             + " SET " + params + " WHERE \"id\"=$" + values.length + " RETURNING *";

//         this._client.query(query, values, (err, result) => {
//             err = err || null;
//             if (!err)
//                 this._logger.trace(correlationId, "Updated partially in %s with id = %s", this._tableName, id);

//             let newItem = result && result.rows && result.rows.length == 1
//                 ? this.convertToPublic(result.rows[0]) : null;

//             if (callback) callback(err, newItem);
//         });
//     }

//    /*
//      * Deleted a data item by it's unique id.
//      *
//      * - correlation_id    (optional) transaction id to trace execution through call chain.
//      * - id                an id of the item to be deleted
//      * - callback          (optional) callback function that receives deleted item or error.
//      */
//     public deleteById(correlationId: string, id: K, callback?: (err: any, item: T) => void): void {
//         let values = [ id ];

//         let query = "DELETE FROM " + this.quoteIdentifier(this._tableName) + " WHERE \"id\"=$1 RETURNING *";

//         this._client.query(query, values, (err, result) => {
//             err = err || null;
//             if (!err)
//                 this._logger.trace(correlationId, "Deleted from %s with id = %s", this._tableName, id);

//             let newItem = result && result.rows && result.rows.length == 1
//                 ? this.convertToPublic(result.rows[0]) : null;

//             if (callback) callback(err, newItem);
//         });
//     }

//    /*
//      * Deletes multiple data items by their unique ids.
//      *
//      * - correlationId     (optional) transaction id to trace execution through call chain.
//      * - ids               ids of data items to be deleted.
//      * - callback          (optional) callback function that receives error or null for success.
//      */
//     public deleteByIds(correlationId: string, ids: K[], callback?: (err: any) => void): void {
//         let params = this.generateParameters(ids);
//         let query = "DELETE FROM " + this.quoteIdentifier(this._tableName) + " WHERE \"id\" IN(" + params + ")";

//         this._client.query(query, ids, (err, result) => {
//             let count = result ? result.rowCount : 0;

//             err = err || null;
//             if (!err)
//                 this._logger.trace(correlationId, "Deleted %d items from %s", count, this._tableName);

//             if (callback) callback(err);
//         });
//     }
// }
