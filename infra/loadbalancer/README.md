# Load Balancer

- Load balancing is the practice of distributing computational workloads between two or more computers.
- Load balancing is often employed to divide network traffic among several servers.
- It reduces the strain on each server and makes the servers more efficient, speeding up performance and reducing latency.


#### How does load balancing work?

- Load balancing is handled by a tool or application called a load balancer.
- A load balancer can be either hardware-based or software-based.
- Hardware load balancers require the installation of a dedicated load balancing device; software-based load balancers can run on a server, on a virtual machine, or in the cloud.

###### Working

- When a request arrives from a user, the load balancer assigns the request to a given server, and this process repeats for each request.
- Load balancers determine which server should handle each request based on a number of different algorithms.


#### Categories of Load Balancer

###### Application Load Balancer Type

- At the application load balancer level, data like: headers, session, response etc. are present. So, we can take decision basis of these data to send the request to the server.

###### Network Load Balancer Type

- At the network load balancer level, TCP port / UDP port / IP Address or SRC & DEST. are present. So, based on these field can take decision whether to route which server.

#### Load Balancer Algorithm

##### Static Algorithms

- ###### Round Robin
  - The round-robin load balancing technique is the simplest way to distribute traffic across a group of servers.
  - The load balancer forwards the incoming requests to dedicated servers sequentially (one by one mechanism).
  - Round-robin is a static load balancer, because it does not modify the state of the servers while distributing incoming traffic.
