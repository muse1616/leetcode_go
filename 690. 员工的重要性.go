package leetcode

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	for _, v := range employees {
		if v.Id == id {
			if len(v.Subordinates) == 0 {
				return v.Importance
			} else {
				for _, v2 := range v.Subordinates {
					v.Importance += getImportance(employees, v2)
				}
				return v.Importance
			}
		}
	}
	return 0
}
