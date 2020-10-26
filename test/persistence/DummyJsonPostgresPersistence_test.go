package test

// let process = require('process');

// import { ConfigParams } from 'pip-services3-commons-node';
// import { DummyPersistenceFixture } from '../fixtures/DummyPersistenceFixture';
// import { DummyJsonPostgresPersistence } from './DummyJsonPostgresPersistence';

// suite('DummyJsonPostgresPersistence', ()=> {
//     let persistence: DummyJsonPostgresPersistence;
//     let fixture: DummyPersistenceFixture;

//     let postgresUri = process.env['POSTGRES_URI'];
//     let postgresHost = process.env['POSTGRES_HOST'] || 'localhost';
//     let postgresPort = process.env['POSTGRES_PORT'] || 5432;
//     let postgresDatabase = process.env['POSTGRES_DB'] || 'test';
//     let postgresUser = process.env['POSTGRES_USER'] || 'postgres';
//     let postgresPassword = process.env['POSTGRES_PASSWORD'] || 'postgres';
//     if (postgresUri == null && postgresHost == null)
//         return;

//     setup((done) => {
//         let dbConfig = ConfigParams.fromTuples(
//             'connection.uri', postgresUri,
//             'connection.host', postgresHost,
//             'connection.port', postgresPort,
//             'connection.database', postgresDatabase,
//             'credential.username', postgresUser,
//             'credential.password', postgresPassword
//         );

//         persistence = new DummyJsonPostgresPersistence();
//         persistence.configure(dbConfig);

//         fixture = new DummyPersistenceFixture(persistence);

//         persistence.open(null, (err: any) => {
//             if (err) {
//                 done(err);
//                 return;
//             }

//             persistence.clear(null, (err) => {
//                 done(err);
//             });
//         });
//     });

//     teardown((done) => {
//         persistence.close(null, done);
//     });

//     test('Crud Operations', (done) => {
//         fixture.testCrudOperations(done);
//     });

//     test('Batch Operations', (done) => {
//         fixture.testBatchOperations(done);
//     });
// });
