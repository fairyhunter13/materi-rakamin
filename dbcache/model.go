package main

import "time"

type Book struct {
	Title         string    `json:"title"`
	Topic         string    `json:"topic"`
	Author        string    `json:"author"`
	DatePublished time.Time `json:"date_published"`
}

type Publisher struct {
	Name              string    `json:"name"`
	FoundedIn         time.Time `json:"founded_in"`
	NumberOfEmployees int       `json:"number_of_employees"`
}
