package main

import (
	"time"
)

type Status string

const (
	Done  Status = "done"
	Doing Status = "doing"
	Todo  Status = "todo"
)

type Task struct {
	Id        uint32    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Title     string    `db:"title" json:"title"`
	Status    Status    `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
