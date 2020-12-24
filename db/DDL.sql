
create table Balancer(
	ID int auto_increment,
    Name varchar(10),
    Primary key(ID)
);

create table VirtualMachine(
	ID int auto_increment,
    BalancerID int not null,
    isWorking bool not null default 1,
    Primary key(ID),
    Foreign key(BalancerID) references Balancer(ID)
);

