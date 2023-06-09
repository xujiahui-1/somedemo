package main

type Person struct {
	Name   string  `json:"name" param:"name"`
	Age    int64   `json:"-"`
	Weight float64 `json:"weight,omitempty"`
	*dog   `json:"dog,omitempty"`
}
