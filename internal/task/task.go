package task

import "time"

type Task struct {
	Messange  string
	DateAlert time.Time
}

// CreateTask принимает Сообщение которое будет нести основную суть, второй аргумент дата оповещение в виде строки
// формата "ДД-ММ-ГГГ ЧЧ:ММ", если дата спарсена не правильно то функция выкидывает ошибку. (nil, error)
func CreateTask(messange string, dateAlert string) (*Task, error) {

	dateParesed, err := time.Parse(time.RFC3339, dateAlert)
	if err != nil {
		return nil, err
	}

	return &Task{
		Messange:  messange,
		DateAlert: dateParesed,
	}, nil
}
