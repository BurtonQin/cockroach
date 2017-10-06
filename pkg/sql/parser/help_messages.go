// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 968
	`ALTER`: {
		//line sql.y: 969
		Category: hGroup,
		//line sql.y: 970
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER DATABASE
`,
	},
	//line sql.y: 979
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 980
		Category: hDDL,
		//line sql.y: 981
		Text: `
ALTER TABLE [IF EXISTS] <tablename> <command> [, ...]

Commands:
  ALTER TABLE ... ADD [COLUMN] [IF NOT EXISTS] <colname> <type> [<qualifiers...>]
  ALTER TABLE ... ADD <constraint>
  ALTER TABLE ... DROP [COLUMN] [IF EXISTS] <colname> [RESTRICT | CASCADE]
  ALTER TABLE ... DROP CONSTRAINT [IF EXISTS] <constraintname> [RESTRICT | CASCADE]
  ALTER TABLE ... ALTER [COLUMN] <colname> {SET DEFAULT <expr> | DROP DEFAULT}
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP NOT NULL
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause>
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

`,
		//line sql.y: 1003
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1015
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1016
		Category: hDDL,
		//line sql.y: 1017
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1019
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1026
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1027
		Category: hDDL,
		//line sql.y: 1028
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1030
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1041
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1042
		Category: hDDL,
		//line sql.y: 1043
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause>
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

`,
		//line sql.y: 1051
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1291
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1292
		Category: hCCL,
		//line sql.y: 1293
		Text: `
BACKUP <targets...> TO <location...>
       [ AS OF SYSTEM TIME <expr> ]
       [ INCREMENTAL FROM <location...> ]
       [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Location:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1310
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1318
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1319
		Category: hCCL,
		//line sql.y: 1320
		Text: `
RESTORE <targets...> FROM <location...>
        [ AS OF SYSTEM TIME <expr> ]
        [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Locations:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1336
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1350
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1351
		Category: hCCL,
		//line sql.y: 1352
		Text: `
IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV

Options:
   distributed = '...'
   sstsize = '...'
   temp = '...'
   comma = '...'          [CSV-specific]
   comment = '...'        [CSV-specific]
   nullif = '...'         [CSV-specific]

`,
		//line sql.y: 1370
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1465
	`CANCEL`: {
		//line sql.y: 1466
		Category: hGroup,
		//line sql.y: 1467
		Text: `CANCEL JOB, CANCEL QUERY
`,
	},
	//line sql.y: 1473
	`CANCEL JOB`: {
		ShortDescription: `cancel a background job`,
		//line sql.y: 1474
		Category: hMisc,
		//line sql.y: 1475
		Text: `CANCEL JOB <jobid>
`,
		//line sql.y: 1476
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOB
`,
	},
	//line sql.y: 1484
	`CANCEL QUERY`: {
		ShortDescription: `cancel a running query`,
		//line sql.y: 1485
		Category: hMisc,
		//line sql.y: 1486
		Text: `CANCEL QUERY <queryid>
`,
		//line sql.y: 1487
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 1495
	`CREATE`: {
		//line sql.y: 1496
		Category: hGroup,
		//line sql.y: 1497
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW
`,
	},
	//line sql.y: 1515
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 1516
		Category: hDML,
		//line sql.y: 1517
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 1520
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 1533
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 1534
		Category: hCfg,
		//line sql.y: 1535
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 1547
	`DROP`: {
		//line sql.y: 1548
		Category: hGroup,
		//line sql.y: 1549
		Text: `DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP USER
`,
	},
	//line sql.y: 1561
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 1562
		Category: hDDL,
		//line sql.y: 1563
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1564
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 1576
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 1577
		Category: hDDL,
		//line sql.y: 1578
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1579
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 1591
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 1592
		Category: hDDL,
		//line sql.y: 1593
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1594
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 1614
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 1615
		Category: hDDL,
		//line sql.y: 1616
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 1617
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 1637
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 1638
		Category: hPriv,
		//line sql.y: 1639
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 1640
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 1682
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 1683
		Category: hMisc,
		//line sql.y: 1684
		Text: `
EXPLAIN <statement>
EXPLAIN [( [PLAN ,] <planoptions...> )] <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN, EXECUTE

Plan options:
    TYPES, EXPRS, METADATA, QUALIFY, INDENT, VERBOSE, DIST_SQL

`,
		//line sql.y: 1695
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 1755
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 1756
		Category: hMisc,
		//line sql.y: 1757
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 1758
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 1780
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 1781
		Category: hMisc,
		//line sql.y: 1782
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 1783
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 1806
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 1807
		Category: hMisc,
		//line sql.y: 1808
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 1809
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 1829
	`GRANT`: {
		ShortDescription: `define access privileges`,
		//line sql.y: 1830
		Category: hPriv,
		//line sql.y: 1831
		Text: `
GRANT {ALL | <privileges...> } ON <targets...> TO <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, ...]
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 1841
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 1849
	`REVOKE`: {
		ShortDescription: `remove access privileges`,
		//line sql.y: 1850
		Category: hPriv,
		//line sql.y: 1851
		Text: `
REVOKE {ALL | <privileges...> } ON <targets...> FROM <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, <databasename>]...
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 1861
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 1948
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 1949
		Category: hCfg,
		//line sql.y: 1950
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 1951
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 1963
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 1964
		Category: hCfg,
		//line sql.y: 1965
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 1966
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 1996
	`SCRUB TABLE`: {
		ShortDescription: `run a scrub check on a table`,
		//line sql.y: 1997
		Category: hMisc,
		//line sql.y: 1998
		Text: `
SCRUB TABLE <tablename> [WITH <option> [, ...]]

Options:
  SCRUB TABLE ... WITH OPTIONS INDEX ALL
  SCRUB TABLE ... WITH OPTIONS INDEX (<index>...)

`,
	},
	//line sql.y: 2035
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 2036
		Category: hCfg,
		//line sql.y: 2037
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 2038
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2059
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 2060
		Category: hCfg,
		//line sql.y: 2061
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }

`,
		//line sql.y: 2066
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2083
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 2084
		Category: hTxn,
		//line sql.y: 2085
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 2092
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 2231
	`SHOW`: {
		//line sql.y: 2232
		Category: hGroup,
		//line sql.y: 2233
		Text: `
SHOW SESSION, SHOW CLUSTER SETTING, SHOW DATABASES, SHOW TABLES, SHOW COLUMNS, SHOW INDEXES,
SHOW CONSTRAINTS, SHOW CREATE TABLE, SHOW CREATE VIEW, SHOW USERS, SHOW TRANSACTION, SHOW BACKUP,
SHOW JOBS, SHOW QUERIES, SHOW SESSIONS, SHOW TRACE
`,
	},
	//line sql.y: 2259
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 2260
		Category: hCfg,
		//line sql.y: 2261
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 2262
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 2283
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 2284
		Category: hCCL,
		//line sql.y: 2285
		Text: `SHOW BACKUP <location>
`,
		//line sql.y: 2286
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 2294
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 2295
		Category: hCfg,
		//line sql.y: 2296
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 2299
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2316
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 2317
		Category: hDDL,
		//line sql.y: 2318
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 2319
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 2327
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 2328
		Category: hDDL,
		//line sql.y: 2329
		Text: `SHOW DATABASES
`,
		//line sql.y: 2330
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 2338
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 2339
		Category: hPriv,
		//line sql.y: 2340
		Text: `SHOW GRANTS [ON <targets...>] [FOR <users...>]
`,
		//line sql.y: 2341
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 2349
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 2350
		Category: hDDL,
		//line sql.y: 2351
		Text: `SHOW INDEXES FROM <tablename>
`,
		//line sql.y: 2352
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 2370
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 2371
		Category: hDDL,
		//line sql.y: 2372
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 2373
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 2386
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 2387
		Category: hMisc,
		//line sql.y: 2388
		Text: `SHOW [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 2389
		SeeAlso: `CANCEL QUERY
`,
	},
	//line sql.y: 2405
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 2406
		Category: hMisc,
		//line sql.y: 2407
		Text: `SHOW JOBS
`,
		//line sql.y: 2408
		SeeAlso: `CANCEL JOB, PAUSE JOB, RESUME JOB
`,
	},
	//line sql.y: 2416
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 2417
		Category: hMisc,
		//line sql.y: 2418
		Text: `
SHOW [KV] TRACE FOR SESSION
SHOW [KV] TRACE FOR <statement>
`,
		//line sql.y: 2421
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 2442
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 2443
		Category: hMisc,
		//line sql.y: 2444
		Text: `SHOW [CLUSTER | LOCAL] SESSIONS
`,
	},
	//line sql.y: 2460
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 2461
		Category: hDDL,
		//line sql.y: 2462
		Text: `SHOW TABLES [FROM <databasename>]
`,
		//line sql.y: 2463
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 2475
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 2476
		Category: hCfg,
		//line sql.y: 2477
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 2478
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 2497
	`SHOW CREATE TABLE`: {
		ShortDescription: `display the CREATE TABLE statement for a table`,
		//line sql.y: 2498
		Category: hDDL,
		//line sql.y: 2499
		Text: `SHOW CREATE TABLE <tablename>
`,
		//line sql.y: 2500
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 2508
	`SHOW CREATE VIEW`: {
		ShortDescription: `display the CREATE VIEW statement for a view`,
		//line sql.y: 2509
		Category: hDDL,
		//line sql.y: 2510
		Text: `SHOW CREATE VIEW <viewname>
`,
		//line sql.y: 2511
		SeeAlso: `WEBDOCS/show-create-view.html
`,
	},
	//line sql.y: 2519
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 2520
		Category: hPriv,
		//line sql.y: 2521
		Text: `SHOW USERS
`,
		//line sql.y: 2522
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 2595
	`PAUSE JOB`: {
		ShortDescription: `pause a background job`,
		//line sql.y: 2596
		Category: hMisc,
		//line sql.y: 2597
		Text: `PAUSE JOB <jobid>
`,
		//line sql.y: 2598
		SeeAlso: `SHOW JOBS, CANCEL JOB, RESUME JOB
`,
	},
	//line sql.y: 2606
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 2607
		Category: hDDL,
		//line sql.y: 2608
		Text: `
CREATE TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> )
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
   UNIQUE ( <colnames... ) [STORING ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
  COLLATE <collationname>

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 2634
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE TABLE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 3114
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 3115
		Category: hDML,
		//line sql.y: 3116
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 3117
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 3125
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 3126
		Category: hPriv,
		//line sql.y: 3127
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 3128
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 3150
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 3151
		Category: hDDL,
		//line sql.y: 3152
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 3153
		SeeAlso: `CREATE TABLE, SHOW CREATE VIEW, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 3167
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 3168
		Category: hDDL,
		//line sql.y: 3169
		Text: `
CREATE [UNIQUE] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 3177
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE INDEX,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 3316
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 3317
		Category: hTxn,
		//line sql.y: 3318
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 3319
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 3327
	`RESUME JOB`: {
		ShortDescription: `resume a background job`,
		//line sql.y: 3328
		Category: hMisc,
		//line sql.y: 3329
		Text: `RESUME JOB <jobid>
`,
		//line sql.y: 3330
		SeeAlso: `SHOW JOBS, CANCEL JOB, PAUSE JOB
`,
	},
	//line sql.y: 3338
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 3339
		Category: hTxn,
		//line sql.y: 3340
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 3341
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 3355
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 3356
		Category: hTxn,
		//line sql.y: 3357
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 3365
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 3378
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 3379
		Category: hTxn,
		//line sql.y: 3380
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 3383
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 3396
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 3397
		Category: hTxn,
		//line sql.y: 3398
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 3399
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 3512
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 3513
		Category: hDDL,
		//line sql.y: 3514
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 3515
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 3584
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 3585
		Category: hDML,
		//line sql.y: 3586
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 3591
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 3608
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 3609
		Category: hDML,
		//line sql.y: 3610
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 3614
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 3690
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 3691
		Category: hDML,
		//line sql.y: 3692
		Text: `UPDATE <tablename> [[AS] <name>] SET ... [WHERE <expr>] [RETURNING <exprs...>]
`,
		//line sql.y: 3693
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 3859
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 3860
		Category: hDML,
		//line sql.y: 3861
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 3872
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 3873
		Category: hDML,
		//line sql.y: 3874
		Text: `
SELECT [DISTINCT]
       { <expr> [[AS] <name>] | [ [<dbname>.] <tablename>. ] * } [, ...]
       [ FROM <source> ]
       [ WHERE <expr> ]
       [ GROUP BY <expr> [ , ... ] ]
       [ HAVING <expr> ]
       [ WINDOW <name> AS ( <definition> ) ]
       [ { UNION | INTERSECT | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
       [ ORDER BY <expr> [ ASC | DESC ] [, ...] ]
       [ LIMIT { <expr> | ALL } ]
       [ OFFSET <expr> [ ROW | ROWS ] ]
       [ FOR UPDATE ]
`,
		//line sql.y: 3887
		SeeAlso: `WEBDOCS/select.html
`,
	},
	//line sql.y: 3947
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 3948
		Category: hDML,
		//line sql.y: 3949
		Text: `TABLE <tablename>
`,
		//line sql.y: 3950
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 4213
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 4214
		Category: hDML,
		//line sql.y: 4215
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 4216
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 4321
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 4322
		Category: hDML,
		//line sql.y: 4323
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexhint> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> ON <expr>
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> USING ( <colnames...> )
  <source> NATURAL { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index hints:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'

`,
		//line sql.y: 4341
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
