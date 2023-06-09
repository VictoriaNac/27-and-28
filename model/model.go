package model

import "fmt"

type Student struct {
	Name  string  // для модели используют экспортируемые поля, то есть с большой буквы
	Age   int
	Grade int
}

// использование реализации стандартных интерфейсов, в библиотеке Го есть масса интерфейсов, в частности в пакете fmt есть интерфейс String, который говорит нам о том, что если у структуры реализован интерфейс String, есть метод String, то именно он будет использоваться при выводе структуры, когда мы используем команду fmt.print, если этого метода String не будет, то fmt.print сама будет решать, как ей вывести информацию о нашей структуре, если мы ей подсуним этот метод, функция fmt.print будет вызывать этот метод и выведет информацию о структуре в том виде, в каком мы здесь опишем. Таких интерфейсов достаточно много. Когда у нас есть модель данных, мы можем продолжить разработку и важно как //

func (s *Student) String() string {
	return fmt.Sprintf("Имя %v, возраст %v, оценка %v", s.Name, s.Age, s.Grade)
}

// для того, чтобы описать код, который взаимодействует с пользователем, либо ту часть, которая будет выполнять хранилище данных нам нужно определиться, как они будут между собой взаимодействовать, для этого как раз используется интерфейс, мы описываем интерфейс взаимодействия (созранение данных и получение), интерфейс не содержит кода, а лишь описывает, как они должны выглядеть. Создаем файл app в пакете app и описываем интерфейс