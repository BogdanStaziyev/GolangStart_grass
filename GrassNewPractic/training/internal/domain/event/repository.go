package event

import (
	"fmt"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

var settings = postgresql.ConnectionURL{
	Database: `postgres`,
	Host:     `localhost:54322`,
	User:     `postgres`,
	Password: `password`,
}

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	Del(id int64) error
	Update(id int64, name string) error
	Create(event *Event) error
}

const EventsCount int64 = 10

type repository struct {
	// Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]Event, error) {
	events := make([]Event, EventsCount)
	db, err := postgresql.Open(settings)
	if err != nil{
		log.Fatal("Open: ", err)
	}
	defer db.Close()
	res := db.Collection("Event")
	err = res.Find().All(&events)
	return events, nil
}

func (r *repository) FindOne(id int64) (*Event, error) {
	var event Event
	db, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer db.Close()
	err = db.Collection("Event").Find("id", id).One(&event)
	fmt.Println(err)
	return &event, nil
}

func (r *repository) Del(id int64) error {
	db, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer db.Close()
	err = db.Collection("Event").Find("id", id).Delete()
	fmt.Println(err)
	return nil
}

func (r *repository) Update(id int64, name string) error{
	var event Event
	db, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer db.Close()
	e := db.Collection("Event")
	res := e.Find("id", id)
	err = res.One(&event)
	if err != nil {
		log.Fatal("Find: ", err)
	}
	event.Name = name
	if err := res.Update(event); err != nil{
		log.Fatal("Update: ", name)
	}
	return nil
}
func (r *repository) Create(event *Event) error {
	db, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer db.Close()
	res := db.Collection("Event")
	if _, err := res.Insert(event); err != nil{
		log.Fatal("Insert: ", err)
	}
	return nil
}