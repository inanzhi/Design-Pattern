package null_object_pattern

type MusicDepartment struct {
	departmentName string
	professors     int
}

func NewMusicDepartment(departmentName string, professors int) *MusicDepartment {
	return &MusicDepartment{departmentName, professors}
}

func (md MusicDepartment) GetName() string {

	return md.departmentName
}

func (md MusicDepartment) GetNumberOfProfessors() int {
	return md.professors
}
