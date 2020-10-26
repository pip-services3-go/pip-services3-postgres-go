package test

// const process = require('process');
// const assert = require('chai').assert;

// import { ConfigParams } from 'pip-services3-commons-node';
// import { Descriptor } from 'pip-services3-commons-node';
// import { References } from 'pip-services3-commons-node';
// import { PostgresConnection } from '../../src/persistence/PostgresConnection';
// import { DummyPersistenceFixture } from '../fixtures/DummyPersistenceFixture';
// import { DummyPostgresPersistence } from './DummyPostgresPersistence';

// suite('DummyPostgresConnection', ()=> {
//     let connection: PostgresConnection;
//     let persistence: DummyPostgresPersistence;
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

//         connection = new PostgresConnection();
//         connection.configure(dbConfig);

//         persistence = new DummyPostgresPersistence();
//         persistence.setReferences(References.fromTuples(
//             new Descriptor("pip-services", "connection", "postgres", "default", "1.0"), connection
//         ));

//         fixture = new DummyPersistenceFixture(persistence);

//         connection.open(null, (err: any) => {
//             if (err) {
//                 done(err);
//                 return;
//             }

//             persistence.open(null, (err: any) => {
//                 if (err) {
//                     done(err);
//                     return;
//                 }

//                 persistence.clear(null, (err) => {
//                     done(err);
//                 });
//             });
//         });
//     });

//     teardown((done) => {
//         connection.close(null, (err) => {
//             persistence.close(null, done);
//         });
//     });

//     test('Connection', (done) => {
//         assert.isObject(connection.getConnection());
//         assert.isString(connection.getDatabaseName());

//         done();
//     });

//     test('Crud Operations', (done) => {
//         fixture.testCrudOperations(done);
//     });

//     test('Batch Operations', (done) => {
//         fixture.testBatchOperations(done);
//     });
// });
