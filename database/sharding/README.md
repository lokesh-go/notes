# Data Sharding

- Sharding is a method for distributing a single dataset across multiple databases, which can then be stored on multiple machines.
- Database sharding is the process of making partitions of data in a database or search engine, such that the data is divided into various smaller distinct chunks, or shards.
- This allows for larger datasets to be split into smaller chunks and stored in multiple data nodes, increasing the total storage capacity of the system.
- Data sharding is a common way of implementing horizontal scaling.
- Database sharding divides the table records in a database into smaller portions. Each section is a shard, and is stored on a different server.


<p align="center"><img src="https://github.com/lokesh-go/notes/assets/31778886/e5ed638f-abf8-4ba5-b705-0f4df0c2d9bd" alt="Sharding" width="350px"/></p>


- Some data within the database remains present in all shards (vertical sharding), but some appear only in single shards (horizontal sharding).


<p align="center"><img src="https://github.com/lokesh-go/notes/assets/31778886/b002f634-f58f-4d31-a83a-65d993311b78" alt="Sharding" width="350px"/></p>

- To shard your data, need to decide a key, called a sharding key, to partition data on. The shard key is either an indexed field or indexed compound fields that exist in every document in the collection.


---

#### Sharding Architectures

