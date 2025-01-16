package main

type Task struct {
	ID          int
	Titre       string
	Description string
	Etat        bool
}

var tasks []Task
