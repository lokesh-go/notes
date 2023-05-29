# ACID

In the transaction processing context, the ACID acronym refers to the four primary and essential properties of a transaction.

These are:
 - Atomicity
 - Consistency
 - Isolation
 - Durability

---

### Atomicity

It implies that all commands within a transaction are treated as a single unit.
Either all will be successful, or none of them will be. 
There won’t be cases of some successful commands and some failed commands.

---

### Consistency

It ensures that all the changes made in the database during the transaction are consistent with the constraints of the database.

---

### Isolation

It guarantees that concurrent transactions, i.e. transactions executing simultaneously, will not affect each other.
The outcomes of one transaction will be unaffected by the other.
Therefore, transactions that run concurrently appear to be running sequentially.

---

### Durability

It takes care of ensuring that once the transaction is complete, the data will be persistent even in case of a system failure, i.e., the data is also written to a backing store.