package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Employee struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}

func main() {
	file, err := os.Open("employees.json")
	if err != nil {
		log.Fatal()
	}
	defer file.Close()

	var employees []Employee
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&employees)
	if err != nil {
		log.Fatal()
	}

	for i, employee := range employees {
		if employee.Id ==3{
			employee.Age = 50
			employees [i ] = employee
		}	
	}

	n := Employee {
		Id:       6,
		Name:     "Shaxboz",
		Age:      20,
		Position: "Golang Junior",
	
	}
	employees = append(employees, n)

    f, err := os.OpenFile("employees_new.json", os.O_RDWR |os.O_CREATE| os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal()
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(employees)
	if err != nil {
		log.Fatal()
	}

	fmt.Println("Employees updated successfully")
}
