package app

import (
	"bufio"
	"decomposition/model"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* есть модель данных, с которой работаем, соответственно мы можем описать интерфейс, потому что мы описываем методы с помощью которых будут взаимодействовать наши слои приложения и нам нужно будет указать как эти методы получают и возвращают данные */

type Storage interface {
	GetAll() ([]*model.Student, error) // получит всё - возвращаем срез из наших студентов
	Put(student *model.Student) error // создать или сохранит студента в хранилище либо вернуть ошибку, если 
                                    // что-то пошло не так
}

type App struct { // объявление структуры app, чтобы был указатель на что-то, что реализует интерфейс 
                  // хранилища, что-то, у чего есть методы интерфейса, на что конкретно нам неважно, 
                  // соотвественно я могу взаимодействовать, а структура с методами создается для того, 
                  // чтобы просто хранить указатель 
	storage Storage // конфликта имен не будет, слово куратора))
                  // pointer implement interface - указатель, что реализует интерфейс
}

func New(storage Storage) *App { // функция new играет роль конструкторов для того, чтобы создать новый 
                                 // экземпляр, в их задачу входит вернуть инициализированный экземпляр 
                                 // нашей структуры, поля которой уже будут заполнены, возвращает 
                                // указатель на созданную структуру, по сути это альтернатива конструкторам
	return &App{
		storage: storage,
	}
}

/* в решении можно создать заглушку, здесь сразу пишу код,
   нужно место, где будет запускаться наше приложение, поэтому сделаем единственный экспортируемый метод Run */

func (app *App) Run() {  // то место откуда будет запускаться приложение
	app.enteringStudents()
	app.printStudents()
}

func (app *App) enteringStudents() { // метод для ввода студентов
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() { // цикл, используем сканер 
    fmt.Println()
		fmt.Println("Enter student's name, age, and grade separated by spaces or type 'exit' to quit:")
		
    line := scanner.Text() // получаем строку
    if line == "exit" {
      app.printStudents()
      fmt.Println("Final students data")
    }

		student, err := populate(line) // вызываем вспомогательную функция, чтобы преобразовать строку в структуру студента
		if err != nil {
			fmt.Println("Неверные данные", err)
			continue // что-то пошло не так, начинаем заново
		}

		err = app.storage.Put(student) // если все хорошо обращаемся к репозиторию и говорим, что надо создать новую запись студента в репозитории
		if err != nil {
			fmt.Println("Не удалось зачислить студента на курс:", err)
			continue
		}
		fmt.Println("Зачисленстудент:", student)

	}

	if err := scanner.Err(); err != nil { // проверка по какой причине был прерван цикл, если ошибки нет, значит завершился ввод
		fmt.Println("Ошибка ввода вывода")
	}
}

// реализуем метод, который напечатает на экране всех студентов, простая программа, запрашивает, потом выводит

func (app *App) printStudents() { // метод распечатать студентов, когда запрашиваем студентов, получаем срез, а как данные хранит структура репозитория это его дело,и будем ключи использовать там, потому что эта часть программы, которая взаимодействует с пользователем не должна меняться от места хранения данных в файле или в хешкарте, каждый занимается своим делом, принцип разделения задач и ответственности, метод может работать с заглушкой без метода ввода.
	students, err := app.storage.GetAll() // обращаемся к app, там есть указатель на хранилище, у хранилища есть метод получить всех студентов
	if err != nil {
		fmt.Println("Ошибка при получении данных:", err)
		return
	}
// если все ок, в цикле обходим срез списка занесенных студентов
	for _, v := range students {
		fmt.Println(v)
	}
}

func populate(line string) (*model.Student, error) { // вспомогательная функция, которая введёную строку преобразует в структуру, получает на вход строку, возвращает модель/структуру студента либо ошибку
	arr := strings.Split(line, " ") // проверяем строку, режем по пробелу, поэтому имена вводятся без фамилии

	if len(arr) != 3 { // проверяем что у нас 3 элемента получилось после разрезания 
		return nil, errors.New("неправильный ввод данных")
	}
// получаем данные в переменные
	name := arr[0] // имя обычная строка

	age, err := strconv.Atoi(arr[1]) // возраст из строки преобразуем с 1го элемента среза в число
	if err != nil {
		return nil, err
	}

	grade, err := strconv.Atoi(arr[1]) // преобразуем оценку
	if err != nil { // если что-то идет не так возвращаем ошибку
		return nil, err 
	}
// если всё нормально, всё преобразовалось, создаём экземпляр студента, создаем именно указатель
	student := &model.Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
	return student, nil // возвращаем указатель на структуру студента
}