package repository

import (
	"context"
	"errors"
	"os"
	"restapi/errors_tec"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Database_handler struct {
	conn *pgx.Conn
	ctx  context.Context
}

func Create_database_handler() (*Database_handler, error) {

	ctx := context.Background()
	if err := godotenv.Load("./.env"); err != nil {
		return nil, err
	}
	path := os.Getenv("POSTGRES_PATH")
	conn, err := pgx.Connect(ctx, path)
	if err != nil {
		return nil, err
	}
	d := Database_handler{conn: conn, ctx: ctx}
	if err = d.Create_database(); err != nil {
		return nil, err
	}
	return &d, nil
}

func (d *Database_handler) Create_database() error {
	str := `CREATE TABLE IF NOT EXISTS Todolist(
		Id SERIAL PRIMARY KEY,
		Name VARCHAR(20) NOT NULL,
		IsCompleted BOOL NOT NULL DEFAULT false,
		Description VARCHAR(100) NOT NULL,
		StartedAt TIMESTAMP NOT NULL,
		EndedAt TIMESTAMP
	)`

	if _, err := d.conn.Exec(d.ctx, str); err != nil {
		return err
	}
	return nil

}

func (d *Database_handler) CreateNewTask_database(m Todolist_model) error {
	str := `INSERT INTO Todolist(Name, IsCompleted, Description, StartedAt, EndedAt) 
		   VALUES ($1,$2,$3,$4,$5)`
	if _, err := d.conn.Exec(d.ctx, str, m.Name, m.IsCompleted, m.Description, m.StartedAt, nil); err != nil {
		return err
	}
	return nil
}

func (d *Database_handler) FoundByName_database(name string) (error, Todolist_model) {
	str := `SELECT * FROM Todolist
		  WHERE Name = $1`

	reader := d.conn.QueryRow(d.ctx, str, name)
	var task Todolist_model
	var id int
	err := reader.Scan(&id, &task.Name, &task.IsCompleted, &task.Description, &task.StartedAt, &task.EndedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors_tec.ErrTaskDoesntExist, task
		}
		return err, task
	}
	return nil, task
}

func (d *Database_handler) DeleteByName_database(name string) error {
	str := `DELETE FROM Todolist 
		  WHERE Name = $1`

	if _, err := d.conn.Exec(d.ctx, str, name); err != nil {
		return err
	}
	return nil
}

func (d *Database_handler) MarkTask_database(m Todolist_model) error {
	str := `UPDATE Todolist SET IsCompleted = $1, EndedAt = $2
		   WHERE Name = $3`
	_, err := d.conn.Exec(d.ctx, str, m.IsCompleted, m.EndedAt, m.Name)
	return err
}

func (d *Database_handler) GetAllNotCompleted_database() ([]Todolist_model, error) {
	str := `SELECT * FROM Todolist 
		  WHERE IsCompleted = false`

	reader, err := d.conn.Query(d.ctx, str)
	if err != nil {
		return nil, err
	}
	var fullList []Todolist_model
	for reader.Next() {
		var task Todolist_model
		var id int
		if err = reader.Scan(&id, &task.Name, &task.IsCompleted, &task.Description, &task.StartedAt, &task.EndedAt); err != nil {
			return nil, err
		}
		fullList = append(fullList, task)
	}
	return fullList, nil
}

func (d *Database_handler) GetAll_database() ([]Todolist_model, error) {
	str := `SELECT * FROM Todolist`

	reader, err := d.conn.Query(d.ctx, str)
	if err != nil {
		return nil, err
	}
	var fullList []Todolist_model
	for reader.Next() {
		var task Todolist_model
		var id int
		if err = reader.Scan(&id, &task.Name, &task.IsCompleted, &task.Description, &task.StartedAt, &task.EndedAt); err != nil {
			return nil, err
		}
		fullList = append(fullList, task)
	}
	return fullList, nil
}
