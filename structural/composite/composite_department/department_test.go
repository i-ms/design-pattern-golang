package composite_department

import "testing"

func TestCompositeDepartment(t *testing.T) {
	ceo := &Employee{
		Name:     "John CEO",
		Position: "CEO",
		Salary:   100000.0,
	}

	techDirector := &Employee{
		Name:     "Alice Tech Director",
		Position: "Tech Director",
		Salary:   80000.0,
	}

	marketingDirector := &Employee{
		Name:     "Bob Marketing Director",
		Position: "Marketing Director",
		Salary:   75000.0,
	}

	developers := &Department{
		Name: "Developers",
	}

	developer1 := &Employee{
		Name:     "Charlie Developer",
		Position: "Developer",
		Salary:   60000.0,
	}

	developer2 := &Employee{
		Name:     "David Developer",
		Position: "Developer",
		Salary:   60000.0,
	}

	developers.Add(developer1)
	developers.Add(developer2)

	marketingTeam := &Department{
    Name: "Marketing Team",
  }

	marketingTeam.Add(marketingDirector)

	techTeam := &Department{
    Name: "Tech Team",
  }
	techTeam.Add(techDirector)
	techTeam.Add(developers)

	ceoTeam := &Department{
    Name: "CEO's Team",
  }
	ceoTeam.Add(ceo)
	ceoTeam.Add(techTeam)
	ceoTeam.Add(marketingTeam)

	ceoTeam.Display(0)
}
