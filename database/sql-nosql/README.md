# SQL vs NoSQL

### SQL

- Structured query language (SQL) is a standard language for database creation and manipulation.
- SQL database is a structured approach to storing data and performing operations using structured query language for data creation, modification.
- MySQL, MSSQL, Oracle, PostgreSQL, SQLite are a relational database program that uses SQL queries.

#### Points

- Tabular data structure format
- Relational
- Pre-determine schema
- Better for multi-row transactions
- Relations b/w different tables
- Support JOIN and Complex queries
- Centralised ( _suppose we have multiples tables then we have to put all tables into one DB [ difficult to have table-1 in DB-1 and table-2 in DB-2 etc ... ]_ )
- Scalability - Vertical Supported ( _Horizontal is not well supported_ )
- Properties - ACID ( _Atomicity, Consistency, Isolation, Durability_ )


---

### NoSQL

- NoSQL databases (aka "not only SQL") are non-tabular databases and store data differently than relational tables.
- NoSQL is a type of database management system (DBMS) that is designed to handle and store large volumes of unstructured and semi-structured data.
- NoSQL is also type of distributed database, which means that information is copied and stored on various servers, which can be remote or local.
- It ensures availability and reliability of data.


#### Types of NoSQL databases

- Key-value store
  - Amazon DynamoDB
  - Redis
  - Aerospike
  - Amazon ElastiCache
- Document store
  - MongoDB
  - Amazon DocumentDB
- Wide-column store
  - Cassandra 
- Graph store
  - Neo4j
  

#### Points

- Unstructured data
- Distributed nature
- Dynamic schemas for unstructured data
- Better for unstructured data like documents or JSON
- Scalability - Horizontal (easily)
- Not ACID, it follow BASE

##### BASE

- BA - Basically Available (Highly available)
- S - Safe state
- E - Eventual Consistency ( _Might get the stale data_ ) ( _when data will be refresh then get the updated data_ )
