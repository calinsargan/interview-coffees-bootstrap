package main

import (
	"github.com/calinsargan/interview-coffees-bootstrap/internal/coffee"
)

var CoffeeService = coffee.NewService()

// todo: #1 define handlers

/*
POST /coffee
payload: coffee.Coffee (json)
response: coffee.Coffee (json) with generated ID
*/

/*
GET /coffee/{id}
payload: N\A
response: coffee.Coffee (json)
*/

/*
PATCH /coffee/{id}
payload: coffee.Coffee (json)
response: coffee.Coffee (json) with updated fields
*/

/*
DELETE /coffee/{id}
payload: N\A
response: N\A
*/
