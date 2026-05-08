package errors_tec

import "errors"

var ErrTaskAlreadyExist = errors.New("Такая задача уже присутствует в списке")
var ErrTaskDoesntExist = errors.New("Такой задачи не существует")
var ErrTaskAlreadyCompleted = errors.New("Эта задача уже выполнена")
