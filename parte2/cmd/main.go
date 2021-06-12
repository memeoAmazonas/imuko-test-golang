package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

/*type Employee struct {
	ID             string `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary string `json:"employee_salary"`
	EmployeeAge    string `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}*/
type User struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}
type Organization struct {
	Organization string `json:"organizacion"`
	Users        []User `json:"users"`
}

func main() {
	file, err := os.Open("parte2/file/data.csv")
	if err != nil {
		log.Fatal("A ocurrido un error leyendo el archivo")
		os.Exit(2)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	columns, err := reader.ReadAll()
	if err != nil {
		log.Fatal("A ocurrido un error leyendo el archivo")
		os.Exit(1)
	}
	var (
		org  Organization
		orgs []Organization
		us   User
	)
	values := make(map[string]Organization)
	usernames := make(map[string][]string)
	for _, item := range columns[1:] {
		if values[item[0]].Organization == "" {
			org.Organization = item[0]
			us.Username = item[1]
			us.Roles = []string{item[2]}
			org.Users = []User{us}
			values[item[0]] = org
			usernames[item[0]] = append(usernames[item[0]], us.Username)
		} else {
			or := values[item[0]]
			fmt.Println(or)
			if contains(usernames[item[0]], item[1]) {
				for u, j := range or.Users {
					if j.Username == item[1] {
						if !contains(j.Roles, item[2]) {
							or.Users[u].Roles = append(or.Users[u].Roles, item[2])
						}
						break
					}
				}
			} else {
				usernames[item[0]] = append(usernames[item[0]], item[1])
				us.Username = item[1]
				us.Roles = []string{item[2]}
				or.Users = append(or.Users, us)
				values[item[0]] = or
			}
		}
	}
	keys := reflect.ValueOf(values).MapKeys()
	for _, key := range keys {
		orgs = append(orgs, values[key.String()])
	}
	json_data, err := json.Marshal(orgs)
	if err != nil {
		log.Fatal("A ocurrido un error parseando el archivo a json")
		os.Exit(1)
	}
	fmt.Println(string(json_data))

	json_file, err := os.Create("parte2/file/sample.json")
	if err != nil {
		log.Fatal("A ocurrido un error creando el archivo a json")
		os.Exit(2)
	}
	defer json_file.Close()

	json_file.Write(json_data)
	json_file.Close()

}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
