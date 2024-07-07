package null_object_pattern

type NullDepartment struct {
}

func (c *NullDepartment) GetName() string {
	return "NullDepartment"
}

func (c *NullDepartment) GetNumberOfProfessors() int {
	return 0
}
