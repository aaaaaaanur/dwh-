package task_11

import (
	"reflect"
	"testing"
)

func TestNewClass(t *testing.T) {
	c := Class{
		ID:       1,
		Number:   "3г",
		students: make([]Student, 0),
	}

	class := NewClass(c.Number)

	if c.Number != class.Number {
		t.Error("number is incorrect")
		t.FailNow()
	}
}

func TestNewTeacher(t *testing.T) {
	teach := Teacher{
		ID:      1,
		Name:    "Иванов Иван Иванович",
		Age:     45,
		Subject: "Физкультура",
	}

	teacher := NewTeacher(teach.Name, teach.Age, teach.Subject)

	if teach.Name != teacher.Name {
		t.Error("name is incorrect")
		t.FailNow()
	}

	if teach.Age != teacher.Age {
		t.Error("age is incorrect")
		t.FailNow()
	}

	if teach.Subject != teacher.Subject {
		t.Error("subject is incorrect")
		t.FailNow()
	}
}

func TestNewStudent(t *testing.T) {
	s := Student{
		ID:   1,
		Name: "Пушкин Александр Сергеевич",
		Age:  15,
	}

	student := NewStudent(s.Name, s.Age)

	if student.Name != s.Name {
		t.Error("name is incorrect")
		t.FailNow()
	}

	if student.Age != s.Age {
		t.Error("age is incorrect")
		t.FailNow()
	}
}

func TestSchool_AddClass(t *testing.T) {
	s := School{
		classes:  []Class{},
		teachers: []Teacher{},
	}

	c1 := NewClass("3г")
	c2 := NewClass("6а")

	s.AddClass(c1)
	s.AddClass(c2)

	if len(s.classes) != 2 {
		t.Error("ошибка с добавлением класса")
		t.FailNow()
	}
}

func TestSchool_AddTeacher(t *testing.T) {
	s := School{
		classes:  []Class{},
		teachers: []Teacher{},
	}

	teach1 := NewTeacher("Иванов Иван Иванович", 34, "ОБЖ")
	teach2 := NewTeacher("Сергеев Сергей Сергеевич", 38, "Физика")

	s.AddTeacher(teach1)
	s.AddTeacher(teach2)

	if len(s.teachers) != 2 {
		t.Error("ошибка с добавлением учителя")
		t.FailNow()
	}
}

func TestClass_AddStudent(t *testing.T) {
	c := Class{
		ID:       1,
		Number:   "3г",
		students: []Student{},
	}

	s1 := NewStudent("Тютчев Фёдор Иванович", 14)
	s2 := NewStudent("Пушкин Александр Сергеевич", 15)

	c.AddStudent(s1)
	c.AddStudent(s2)

	if len(c.students) != 2 {
		t.Error("ошибка с добавлением ученика")
		t.FailNow()
	}
}

func TestSchool_UpdateClass(t *testing.T) {
	s := School{
		classes:  []Class{},
		teachers: []Teacher{},
	}

	c1 := NewClass("3г")
	c2 := NewClass("6а")

	s.AddClass(c1)
	s.AddClass(c2)

	c2.Number = "6в"
	s.UpdateClass(c2)

	for _, val := range s.classes {
		if val.ID == c2.ID {
			if reflect.DeepEqual(val, c2) == false {
				t.Error("ошибка обновления данных класса")
			}
		}
	}
}

func TestSchool_UpdateTeacher(t *testing.T) {
	s := School{
		classes:  []Class{},
		teachers: []Teacher{},
	}

	teach1 := NewTeacher("Иванов Иван Иванович", 34, "ОБЖ")
	teach2 := NewTeacher("Сергеев Сергей Сергеевич", 38, "Информатика")

	s.AddTeacher(teach1)
	s.AddTeacher(teach2)

	teach1.Age = 30
	s.UpdateTeacher(teach1)

	for _, val := range s.teachers {
		if val.ID == teach1.ID {
			if reflect.DeepEqual(val, teach1) == false {
				t.Error("ошибка обновления данных учителя")
			}
		}
	}
}

func TestClass_UpdateStudent(t *testing.T) {
	c := Class{
		ID:       1,
		Number:   "3г",
		students: []Student{},
	}

	s1 := NewStudent("Тютчев Фёдор Иванович", 14)
	s2 := NewStudent("Пушкин Александр Сергеевич", 15)

	c.AddStudent(s1)
	c.AddStudent(s2)

	s1.Age = 24
	c.UpdateStudent(s1)

	for _, val := range c.students {
		if val.ID == s1.ID {
			if reflect.DeepEqual(val, s1) == false {
				t.Error("ошибка обновления данных ученика")
			}
		}
	}
}

func TestSchool_DeleteClass(t *testing.T) {
	s := School{
		classes:  []Class{},
		teachers: []Teacher{},
	}

	c1 := NewClass("3г")
	c2 := NewClass("6а")

	s.AddClass(c1)
	s.AddClass(c2)

	s.DeleteClass(c2)

	if len(s.classes) != 1 {
		t.Error("ошибка с удалением данных класса")
		t.FailNow()
	}
}

func TestSchool_DeleteTeacher(t *testing.T) {
	s := School{
		classes:  []Class{},
		teachers: []Teacher{},
	}

	teach1 := NewTeacher("Иванов Иван Иванович", 34, "ОБЖ")
	teach2 := NewTeacher("Сергеев Сергей Сергеевич", 38, "Информатика")

	s.AddTeacher(teach1)
	s.AddTeacher(teach2)

	s.DeleteTeacher(teach2)

	if len(s.teachers) != 1 {
		t.Error("ошибка с удалением данных класса")
		t.FailNow()
	}
}

func TestClass_DeleteStudent(t *testing.T) {
	c := Class{
		ID:       1,
		Number:   "3г",
		students: []Student{},
	}

	s1 := NewStudent("Тютчев Фёдор Иванович", 14)
	s2 := NewStudent("Пушкин Александр Сергеевич", 15)

	c.AddStudent(s1)
	c.AddStudent(s2)

	c.DeleteStudent(s2)

	if len(c.students) != 1 {
		t.Error("ошибка с удалением данных ученика")
		t.FailNow()
	}
}

func TestSchool_FindByClassNumber(t *testing.T) {
	s := School{
		classes:  []Class{},
		teachers: []Teacher{},
	}

	c1 := NewClass("3г")
	c2 := NewClass("6а")

	s.AddClass(c1)
	s.AddClass(c2)

	f := s.FindByClassNumber(c1.Number)

	if !reflect.DeepEqual(c1, f) {
		t.Error("ошибка с поиском класса по номеру")
	}
}

func TestSchool_FindByTeacherName(t *testing.T) {
	s := School{
		classes:  []Class{},
		teachers: []Teacher{},
	}

	teach1 := NewTeacher("Иванов Иван Иванович", 34, "ОБЖ")
	teach2 := NewTeacher("Сергеев Сергей Сергеевич", 38, "Информатика")
	teach3 := NewTeacher("Алёнова Алёна Алёновна", 23, "Информатика")

	s.AddTeacher(teach1)
	s.AddTeacher(teach2)
	s.AddTeacher(teach3)

	f := s.FindByTeacherName(teach1.Name)

	if len(f) != 1 {
		t.Error("ошибка с поиском учителя по имени")
		t.FailNow()
	}

	if !reflect.DeepEqual(f[0], teach1) {
		t.Error("ошибка с поиском учителя по имени")
	}
}

func TestSchool_FindByTeacherSubject(t *testing.T) {
	s := School{
		classes:  []Class{},
		teachers: []Teacher{},
	}

	teach1 := NewTeacher("Иванов Иван Иванович", 34, "ОБЖ")
	teach2 := NewTeacher("Сергеев Сергей Сергеевич", 38, "Информатика")
	teach3 := NewTeacher("Алёнова Алёна Алёновна", 23, "Информатика")

	s.AddTeacher(teach1)
	s.AddTeacher(teach2)
	s.AddTeacher(teach3)

	f := s.FindByTeacherSubject(teach2.Subject)

	if !reflect.DeepEqual(f, []Teacher{teach2, teach3}) {
		t.Error("ошибка с поиском учителя по предмету")

	}
}

func TestClass_FindByStudentName(t *testing.T) {
	c := Class{
		ID:       1,
		Number:   "3г",
		students: []Student{},
	}

	s1 := NewStudent("Тютчев Фёдор Иванович", 14)
	s2 := NewStudent("Пушкин Александр Сергеевич", 15)

	c.AddStudent(s1)
	c.AddStudent(s2)

	f := c.FindByStudentName(s1.Name)

	if len(f) != 1 {
		t.Error("ошибка с поиском ученика по имени")
		t.FailNow()
	}

	if !reflect.DeepEqual(f[0], s1) {
		t.Error("ошибка с поиском ученика по имени")
	}
}

func TestSchool(t *testing.T) {
	class := NewClass("11а")

	student1 := NewStudent("Чеботать Сергей Вячеславович", 38)
	student2 := NewStudent("Нурдавлетова Алёна Зуфаровна", 17)

	class.AddStudent(student1)
	class.AddStudent(student2)

	student1.Age = 18
	class.UpdateStudent(student1)

	for _, val := range class.students {
		if val.ID == student1.ID {
			t.Log(student1.Age)
		}
	}

	student3 := NewStudent("Путин Владимир Владимирович", 71)

	class.AddStudent(student3)
	t.Log(class.students)

	class.DeleteStudent(student3)
	t.Log(class.students)
}

func TestSchool_AddTeacherToClass(t *testing.T) {
	s := School{
		classes:  []Class{},
		teachers: []Teacher{},
	}

	teach1 := NewTeacher("Иванов Иван Иванович", 34, "ОБЖ")
	teach2 := NewTeacher("Сергеев Сергей Сергеевич", 38, "Физика")

	s.AddTeacher(teach1)
	s.AddTeacher(teach2)

	class1 := NewClass("3G")
	class2 := NewClass("5G")

	s.AddClass(class1)
	s.AddClass(class2)

	s.AddTeacherToClass(class1.ID, teach1.ID)
	s.AddTeacherToClass(class1.ID, teach2.ID)
	s.AddTeacherToClass(class2.ID, teach2.ID)

	teachers := s.FindTeachersByClass(class1.Number)

	if !reflect.DeepEqual(teachers, []Teacher{teach1, teach2}) {
		t.Error("Ошибка поиска учителей по классу")
	}

	teachers = s.FindTeachersByClass(class2.Number)

	if !reflect.DeepEqual(teachers, []Teacher{teach2}) {
		t.Error("Ошибка поиска учителей по классу")
	}
}
