
insert into Balancer(Name)
values
("balancer1"),
("balancer2"),
("balancer3");

insert into VirtualMachine(BalancerID,isWorking)
values
(1,1),(1,0),(2,1),(3,1),(1,1),(3,1),(1,1),(3,1),
(2,1),(1,1),(2,1),(1,1),(2,0),(3,1),(1,0),(2,1),
(3,1),(2,0),(1,1),(3,0),(1,1),(2,1),(1,1),(3,1);
