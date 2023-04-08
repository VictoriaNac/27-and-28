package repository

import (
	"decomposition/app"
	"decomposition/model"
	"errors"
)

type MemStorage struct { // нужна структура,у которой будет поле в виде хеш карты, ключом будет имя, в 30 задании будет ключом id, которыйсамт будем высчитывать
	studentsByName map[string]*model.Student
}

var _ app.Storage = (*MemStorage)(nil) // чтобы знать, что методы реализованы

func NewMemStorage() *MemStorage { // добавлен конструктор, чтобы не делать дополнительных проверок
	return &MemStorage{
		studentsByName: make(map[string]*model.Student), // инициализируем хешкарту в структуре и возвращаем её готовую
	}                                                   // прячем детализацию, чтобы не знали, как мы храним данные
}                                                     // все взаимоотношения происходят через методы

func (ms *MemStorage) GetAll() ([]*model.Student, error) {
	var v *model.Student
  students := make([]*model.Student, 0, len(ms.studentsByName)) // создаем срез
	for _, v = range ms.studentsByName { // обходим хешкарту, ключ игнорируем, берем данные
		students = append(students, v) // добавляем в срез
	}
	return students, nil
  }

func (ms *MemStorage) Put(student *model.Student) error { // получает указатель на структуру студента на входе
	if _, found := ms.studentsByName[student.Name]; !found { // проверяем естьли такой студент имяв хешкарте или нет, пытаемся достать запись по этому ключу, если ничего не достали, вернулся false, соответственно можем его добавить
		ms.studentsByName[student.Name] = student
		return nil // и возвращаем nil то есть все хорошо
	}
	return errors.New("студент с таким именем уже зачислен")
}