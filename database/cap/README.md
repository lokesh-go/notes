# CAP Theorem

- It is desirable property of distributed system with replicated data.
- CAP Theorem is a concept that a distributed database system can only have 2 of the 3: _Consistency, Availability and Partition Tolerance_

<p align="center"><img src="https://github-production-user-asset-6210df.s3.amazonaws.com/31778886/241996687-b520dad2-82ca-4fe4-a1c0-f1fbc9e19ef2.png" alt="CAP_Theorem" width="300px"/></p>


---


#### Consistency

- A system is said to be consistent if all nodes see the same data at the same time.
- Simply, if we perform a read operation on a consistent system, it should return the value of the most recent write operation.
- It means that, the read should cause all nodes to return the same data, i.e., the value of the most recent write.


#### Availability

- All nodes should be available. All node should be response.
- Availability in a distributed system ensures that the system remains operational 100% of the time.
- Every request gets a (non-error) response regardless of the _individual state_ of a node. ( _Note: this does not guarantee that the response contains the most recent write._ )


#### Partition Tolerance

- A partition is a communications break within a distributed system—a lost or temporarily delayed connection between two nodes.
- Partition tolerance means that the cluster must continue to work despite any number of communication breakdowns between nodes in the system.


---

#### Examples

- MongoDB is a CP data store — it resolves network partitions by maintaining consistency, while compromising on availability.
- Cassandra is an AP database — it delivers availability and partition tolerance but can't deliver consistency all the time. Because Cassandra doesn't have a master node, all the nodes must be available continuously. However, Cassandra provides eventual consistency by allowing clients to write to any nodes at any time and reconciling inconsistencies as quickly as possible.


---

#### Why we can't use CAP together?

Suppose, A, B & C are three nodes.
- A -> Application
- B -> DB node-1
- C -> DB node-2

###### Check for Consistency:

Here, B node is sync to C node.
So, If a=5 at B node then a=5 at C node also. ( _because bothe nodes are sync with each others_ )

Hence, Consistency is possible.

###### Check for Availability:

If A query B node or C node then DB should be respond. Yes It is possible here.

Hence, Availability is possible.


###### Check for Partition Tolerance:

Now, let's say partition happens between B and C nodes for any reason. Then data sync is not happening.

If a value has been updated to 6 at B node then C nodes will not sync due to partition.

So, here data is not consistent. 

Here we find that, If Partition happens then data is not Consistent.

So possilbe senario for Partition Tolerance is:  AP ( _Availability and Partition Tolerance_ ).


---

#### Possible cases

- CA - If we want to achieve Consistency & Availability then can’t do a partition between any two nodes in the system.
- CP - If we want to achieve Consistency & Partition tolerance then will have give up Availability. When a partition occurs between any two nodes, the system has to shut down the non-consistent node (i.e., make it unavailable) until the partition is resolved.
- AP - If we want to achieve Availability & Partition tolerance then will have give up Consistency. When a partition occurs, all nodes remain available but those at the wrong end of a partition might return an older version of data than others. (When the partition is resolved, the AP databases typically resync the nodes to repair all inconsistencies in the system.)


---

#### What we should consider

In distributed system, Partition tolerance is common problem.
So we should consider Partition Tolerance.

And we can choose one of them ( _Consistency or Availability_ ). Either CP or AP.