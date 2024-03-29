package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/domain/event"
)



type EventController struct {
	service *event.Service
}

func NewEventController(s *event.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s", err)
			}
			return
		}

		err = success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}
	}
}

func (c *EventController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		event, err := (*c.service).FindOne(id)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}

		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}

func (c *EventController) Del() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.Del(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Del(): %s", err)
			}
			return
		}
		err = (*c.service).Del(id)
		if err != nil {
			fmt.Printf("EventController.Del(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Del(): %s", err)
			}
			return
		}
		err = success(w, "Event is deleting")
		if err != nil {
			fmt.Printf("EventController.Del(): %s", err)
		}
	}
}

func (c *EventController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event event.Event
		json.NewDecoder(r.Body).Decode(&event)
		err := (*c.service).Update(event.Id, event.Name)
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Update(): %s", err)
			}
			return
		}
		err = success(w, "Event is updating")
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
		}
	}
}

func (c *EventController) Create() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request)  {
		var event event.Event
		json.NewDecoder(r.Body).Decode(&event)
		err := (*c.service).Create(&event)
		if err != nil {
			fmt.Printf("EventController.Create(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Create(): %s", err)
			}
			return
		}
		err = success(w, "Event is creating")
		if err != nil {
			fmt.Printf("EventController.Create(): %s", err)
		}
	}
}