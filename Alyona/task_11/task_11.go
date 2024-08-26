package task_11

var teacherID uint
var classID uint
var studentID uint

type Student struct {
	ID   uint
	Name string
	Age  int
}

type Class struct {
	ID       uint
	Number   string
	students []Student
}

type Teacher struct {
	ID      uint
	Name    string
	Age     int
	Subject string
}

type TeacherClass struct {
	TeacherID uint
	ClassID   uint
}

type School struct {
	classes      []Class
	teachers     []Teacher
	teacherClass []TeacherClass
}

func (s *School) AddTeacherToClass(classID, teacherID uint) {
	x := TeacherClass{
		TeacherID: teacherID,
		ClassID:   classID,
	}
	s.teacherClass = append(s.teacherClass, x)
}

func (s *School) FindTeachersByClass(number string) []Teacher {
	class := s.FindByClassNumber(number)

	teachers := make([]Teacher, 0)
	for _, tc := range s.teacherClass {
		if tc.ClassID == class.ID {
			teacher := s.FindByTeacherID(tc.TeacherID)
			teachers = append(teachers, teacher)
		}
	}
	return teachers
}

func NewClass(class string) Class {
	classID += 1
	var sliceStudents []Student
	C := Class{classID, class, sliceStudents}
	return C
}

func NewTeacher(name string, age int, subject string) Teacher {
	teacherID += 1
	teacher := Teacher{teacherID, name, age, subject}
	return teacher
}

func NewStudent(name string, age int) Student {
	studentID += 1
	student := Student{studentID, name, age}
	return student
}

func (s *School) AddClass(c Class) {
	s.classes = append(s.classes, c)
}

func (s *School) AddTeacher(t Teacher) {
	s.teachers = append(s.teachers, t)
}

func (c *Class) AddStudent(s Student) {
	c.students = append(c.students, s)
}

func (s *School) UpdateClass(c Class) {
	for i, val := range s.classes {
		if val.ID == c.ID {
			s.classes[i] = c
		}
	}
}

func (s *School) UpdateTeacher(t Teacher) {
	for i, val := range s.teachers {
		if val.ID == t.ID {
			s.teachers[i] = t
		}
	}
}

func (c *Class) UpdateStudent(s Student) {
	for i, val := range c.students {
		if val.ID == s.ID {
			c.students[i] = s
		}
	}
}

func (s *School) DeleteClass(c Class) {
	for i, val := range s.classes {
		if val.ID == c.ID && (i+1) < len(s.classes) {
			s.classes = append(s.classes[:i], s.classes[i+1])
		} else if val.ID == c.ID && (i+1) == len(s.classes) {
			s.classes = append(s.classes[:i])
		}
	}
}

func (s *School) DeleteTeacher(t Teacher) {
	for i, val := range s.teachers {
		if val.ID == t.ID && (i+1) < len(s.teachers) {
			s.teachers = append(s.teachers[:i], s.teachers[i+1])
		} else if val.ID == t.ID && (i+1) == len(s.teachers) {
			s.teachers = append(s.teachers[:i])

		}
	}
}

func (c *Class) DeleteStudent(s Student) {
	for i, val := range c.students {
		if val.ID == s.ID && (i+1) < len(c.students) {
			c.students = append(c.students[:i], c.students[i+1])
		} else if val.ID == s.ID && (i+1) == len(c.students) {
			c.students = append(c.students[:i])
		}
	}
}

func (s *School) FindByClassNumber(number string) Class {
	var cn Class
	for _, val := range s.classes {
		if val.Number == number {
			cn = val
		}
	}
	return cn
}

func (s *School) FindByTeacherID(id uint) Teacher {
	var t Teacher
	for _, teacher := range s.teachers {
		if teacher.ID == id {
			t = teacher
		}
	}
	return t
}

func (s *School) FindByTeacherName(name string) []Teacher {
	var tn []Teacher
	for _, val := range s.teachers {
		if val.Name == name {
			tn = append(tn, val)
		}
	}
	return tn
}

func (s *School) FindByTeacherSubject(subject string) []Teacher {
	var ts []Teacher
	for _, val := range s.teachers {
		if val.Subject == subject {
			ts = append(ts, val)
		}
	}
	return ts
}

func (c *Class) FindByStudentName(name string) []Student {
	var sn []Student
	for _, val := range c.students {
		if val.Name == name {
			sn = append(sn, val)
		}
	}
	return sn
}
