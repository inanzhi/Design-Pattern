package null_object_pattern

type ScienceDepartment struct {
	departmentName string
	professors     int
}

func NewScienceDepartment(departmentName string, professors int) *ScienceDepartment {
	return &ScienceDepartment{departmentName, professors}
}
func (md ScienceDepartment) GetName() string {
	return md.departmentName
}

func (md ScienceDepartment) GetNumberOfProfessors() int {
	return md.professors
}
