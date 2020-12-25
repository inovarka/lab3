package balancers

import (
	"database/sql"
)

type Balancer struct {
	Id                 int64    `json:"id"`
	UsedMachines       []string `json:"usedMachines"`
	TotalMachinesCount []string `json:"totalMachinesCount"`
}

type Machine struct {
	Id        int64 `json:"id"`
	IsWorking bool  `json:"isWorking"`
}

type Balancers struct {
	BalancersArr []*Balancer `json:"Balancers"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

//ListBalancers returns a list of all balancers
func (s *Store) ListBalancers() (*Balancers, error) {
	rows, err := s.Db.Query("SELECT id FROM Balancer")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Balancer
	for rows.Next() {
		var b Balancer
		if err := rows.Scan(&b.Id); err != nil {
			return nil, err
		}
		res = append(res, &b)
	}

	var fullBalancers []*Balancer
	if res == nil {
		fullBalancers = make([]*Balancer, 0)
	} else {
		for i := 0; i < len(res); i++ {
			machines, err := s.GetWorkingMachineByID(res[i].Id)
			if err != nil {
				return nil, err
			}
			machinesCount, err := s.GetMachineCountByID(res[i].Id)
			if err != nil {
				return nil, err
			}
			fullBalancer := Balancer{
				Id:                 res[i].Id,
				UsedMachines:       machines,
				TotalMachinesCount: machinesCount}
			fullBalancers = append(fullBalancers, &fullBalancer)
		}

	}

	result := &Balancers{fullBalancers}
	return result, err
}

func (s *Store) GetMachineCountByID(id int64) ([]string, error) {
	rows, err := s.Db.Query(`select count(*) from balancer b 
	join VirtualMachine vm on vm.BalancerID = b.ID
	where b.ID = ?`, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []string
	for rows.Next() {
		var u string
		if err := rows.Scan(&u); err != nil {
			return nil, err
		}
		if u != "" {
			res = append(res, u)
		}
	}
	if res == nil {
		res = make([]string, 0)
	}

	return res, nil
}

func (s *Store) GetWorkingMachineByID(id int64) ([]string, error) {
	rows, err := s.Db.Query(`select vm.ID from balancer b 
	join VirtualMachine vm on vm.BalancerID = b.ID
	where b.ID = ? and vm.IsWorking = true`, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []string
	for rows.Next() {
		var u string
		if err := rows.Scan(&u); err != nil {
			return nil, err
		}
		if u != "" {
			res = append(res, u)
		}
	}
	if res == nil {
		res = make([]string, 0)
	}

	return res, nil
}

//UpdateMachine updates a machine in DB
func (s *Store) UpdateMachine(id int64, isWorking bool) error {
	_, err := s.Db.Exec("update VirtualMachine set isWorking=$1 where id=$2", isWorking, id)
	return err
}
