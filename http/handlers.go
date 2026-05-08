package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"restapi/errors_tec"
	"restapi/todo"
	"time"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	todoList *todo.Tasks
}

func CreateHTTPHandlers(todoList *todo.Tasks) *HTTPHandlers {
	return &HTTPHandlers{todoList: todoList}
}

func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var info = TaskDTO{}
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		errMessage := ErrDTO{err.Error(), time.Now()}
		http.Error(w, errMessage.ToJSON(), http.StatusBadRequest)
		return
	}
	if err = info.Validate(); err != nil {
		errMessage := ErrDTO{err.Error(), time.Now()}
		http.Error(w, errMessage.ToJSON(), http.StatusBadRequest)
		return
	}
	Task := todo.CreateNewTask(info.Name, info.Description)
	if err = h.todoList.CreateTask(*Task); err != nil {
		errMessage := ErrDTO{err.Error(), time.Now()}
		if errors.Is(err, errors_tec.ErrTaskAlreadyExist) {
			http.Error(w, errMessage.ToJSON(), http.StatusConflict)
		} else {
			http.Error(w, errMessage.ToJSON(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(Task, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(b); err != nil {
		fmt.Println("Не получилось отправить http ответ", err)
	}

}

func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	if err := h.todoList.DeleteTask(title); err != nil {
		if errors.Is(err, errors_tec.ErrTaskDoesntExist) {
			errMessage := ErrDTO{Message: err.Error(), Time: time.Now()}
			http.Error(w, errMessage.ToJSON(), http.StatusNotFound)
		} else {
			errMessage := ErrDTO{Message: err.Error(), Time: time.Now()}
			http.Error(w, errMessage.ToJSON(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPHandlers) HandleMarkAsCompletedTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	var info = CompletedDTO{}
	err := json.NewDecoder(r.Body).Decode(&info)

	if info.Completed == nil {
		errMessage := ErrDTO{Message: "Некорректные данные в теле запроса", Time: time.Now()}
		http.Error(w, errMessage.ToJSON(), http.StatusBadRequest)
		return
	}

	if err != nil {
		errMessage := ErrDTO{Message: err.Error(), Time: time.Now()}
		http.Error(w, errMessage.ToJSON(), http.StatusBadRequest)
		return
	}

	if err := h.todoList.MarkTask(title, *info.Completed); err != nil {
		if errors.Is(err, errors_tec.ErrTaskDoesntExist) {
			errMessage := ErrDTO{Message: err.Error(), Time: time.Now()}
			http.Error(w, errMessage.ToJSON(), http.StatusNotFound)
			return
		} else {
			errMessage := ErrDTO{Message: err.Error(), Time: time.Now()}
			http.Error(w, errMessage.ToJSON(), http.StatusInternalServerError)
			return
		}
	}
	task, _ := h.todoList.GetTask(title)
	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("Не удалось отправить http ответ", err)
	}

}

func (h *HTTPHandlers) HandleGetAllNotCompletedTask(w http.ResponseWriter, r *http.Request) {
	b, err := json.MarshalIndent(h.todoList.GetAllNotCompleted(), "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("Не удалось отправить http ответ", err)
	}

}

func (h *HTTPHandlers) HandleGetAllTask(w http.ResponseWriter, r *http.Request) {
	b, err := json.MarshalIndent(h.todoList.GetAll(), "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("Не удалось отправить http ответ", err)
	}

}

func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	task, err := h.todoList.GetTask(title)
	if err != nil {
		if errors.Is(err, errors_tec.ErrTaskDoesntExist) {
			errMessage := ErrDTO{Message: err.Error(), Time: time.Now()}
			http.Error(w, errMessage.ToJSON(), http.StatusNotFound)
		} else {
			errMessage := ErrDTO{Message: err.Error(), Time: time.Now()}
			http.Error(w, errMessage.ToJSON(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err = w.Write(b); err != nil {
		fmt.Println("Не удалось отправить http ответ", err)
	}

}
