# Data Replication

- Data Replication is the process of storing data in more than one site or node. It is useful in improving the availability of data.
- Database replication refers to the process of copying data from a primary database to one or more replica databases in order to improve data accessibility and system fault-tolerance and reliability.

---

#### Challenges for Database Replication Handling

- Inconsistent data - Some of your data may not correctly sync with the rest of your distributed system when you’re copying data between multiple sites at different intervals.
- Lost data - Some of your data may be lost if database objects are incorrectly configured within the source database or if the primary key you use to verify data integrity in the replica is incorrect.

---


#### Mongodb Replica Set


- With MongoDB, replication is achieved through a Replica Set. Writer operations are sent to the primary server (node), which applies the operations across secondary servers, replicating the data.
- If the primary server fails (through a crash or system failure), one of the secondary servers takes over and becomes the new primary node via election. If that server comes back online, it becomes a secondary once it fully recovers, aiding the new primary node.


##### Replica Set Elections

Replica sets use elections to determine which set member will become primary. Replica sets can trigger an election in response to a variety of events, such as:

- Adding a new node to the replica set
- Initiating a replica set
- The secondary members losing connectivity to the primary for more than the configured timeout (10 seconds by default).
- Replica set members send heartbeats (pings) to each other every two seconds. If a heartbeat does not return within 10 seconds, the other members mark the delinquent member as inaccessible.


##### Replica Set Deployment Architectures

- A replica set can have up to 50 members, but only 7 voting members. If the replica set already has 7 voting members, additional members must be non-voting members.
- Ensure that the replica set has an odd number of voting members.
- A replica set can have up to 7 voting members.
-  If you have an even number of voting members, deploy another data bearing voting member or, if constraints prohibit against another data bearing voting member, an arbiter.


##### Three Member Replica Sets

- The minimum number of replica set members needed to obtain the benefits of a replica set is three members. 
- A three member replica set can have either three data-bearing members (Primary-Secondary-Secondary).
- If circumstances (such as cost) prohibit adding a third data bearing member, two data-bearing members and an arbiter (Primary-Secondary-Arbiter).


###### Primary with Two Secondary Members (P-S-S)

- One primary.
- Two secondary members. Both secondaries can become the primary in an election.


<p align="center"><img src="https://github.com/lokesh-go/notes/assets/31778886/17226d5f-4583-42cd-a324-335f17289298" alt="PSS" width="300px"/></p>


- These deployments provide two complete copies of the data set at all times in addition to the primary.
- These replica sets provide additional fault tolerance and high availability. 
- If the primary is unavailable, the replica set elects a secondary to be primary and continues normal operation.
- The old primary rejoins the set when available.


###### Primary with a Secondary and an Arbiter (PSA)

- One primary.
- One secondary member. The secondary can become primary in an election.
- One arbiter. The arbiter only votes in elections.

- For example, in the following replica set with 2 data-bearing members (the primary and a secondary), an arbiter allows the set to have an odd number of votes to break a tie:


<p align="center"><img src="https://github.com/lokesh-go/notes/assets/31778886/b6fc4fe1-416e-43bc-88f4-5463a29a177d" alt="PSA" width="300px"/></p>


- Since the arbiter does not hold a copy of the data, these deployments provides only one complete copy of the data.
- Arbiters require fewer resources, but at the expense of more limited redundancy and fault tolerance.
