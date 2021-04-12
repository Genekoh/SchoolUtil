package schoolutil

type Student struct {
	FirstName string
	LastName string
}

func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

func NewStudent(params ...string) Student {
	s := Student{}
	if len(params) == 0 {
		return s
	}

	s.FirstName = params[0]
	s.LastName = params[1]
	return s
}