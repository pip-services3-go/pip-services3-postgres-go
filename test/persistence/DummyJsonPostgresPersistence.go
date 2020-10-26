package test

// import { FilterParams } from 'pip-services3-commons-node';
// import { PagingParams } from 'pip-services3-commons-node';
// import { DataPage } from 'pip-services3-commons-node';

// import { IdentifiableJsonPostgresPersistence } from '../../src/persistence/IdentifiableJsonPostgresPersistence';
// import { Dummy } from '../fixtures/Dummy';
// import { IDummyPersistence } from '../fixtures/IDummyPersistence';

// export class DummyJsonPostgresPersistence
//     extends IdentifiableJsonPostgresPersistence<Dummy, string>
//     implements IDummyPersistence
// {
//     public constructor() {
//         super('dummies_json');
//         this.ensureTable();
//         this.ensureIndex('dummies_json_key', { "(data->>'key')": 1 }, { unique: true });
//     }

//     public getPageByFilter(correlationId: string, filter: FilterParams, paging: PagingParams,
//         callback: (err: any, page: DataPage<Dummy>) => void): void {
//         filter = filter || new FilterParams();
//         let key = filter.getAsNullableString('key');

//         let filterCondition: string = "";
//         if (key != null)
//             filterCondition += "data->key='" + key + "'";

//         super.getPageByFilter(correlationId, filterCondition, paging, null, null, callback);
//     }

//     public getCountByFilter(correlationId: string, filter: FilterParams,
//         callback: (err: any, count: number) => void): void {
//         filter = filter || new FilterParams();
//         let key = filter.getAsNullableString('key');

//         let filterCondition: string = "";
//         if (key != null)
//             filterCondition += "data->key='" + key + "'";

//         super.getCountByFilter(correlationId, filterCondition, callback);
//     }
// }
