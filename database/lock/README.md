# LOCKing Concept

- Lock Based Concept in DBMS is a mechanism in which a transaction cannot Read or Write the data until it acquires an appropriate lock.
- Lock based concept help to eliminate the concurrency problem in DBMS for simultaneous transactions by locking or isolating a particular transaction to a single user.
- The lock is a mechanism associated with a table for restricting unauthorized access to the data. It is mainly used to solve the concurrency problem in transactions. We can apply a lock on row level, database level, table level, and page level.
- It is better to understand that _"A lock is designed to ensure data integrity and consistency while enabling concurrent data access, as it forces each transaction to pass the ACID test. When several users access a database to alter its data at the same time, it implements concurrency control."_
- We know that multiple users can access databases at the same time. As a result, locking is essential for a successful transaction and protects data from being corrupted or invalidated when several users attempt to read, write, or update a database. Usually, the lock is an in-memory structure with owners, types, and the hash of the resource they are supposed to protect. As an in-memory structure, the size of a lock is 96 bytes.
  