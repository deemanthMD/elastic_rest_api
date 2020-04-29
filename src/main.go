package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
	//"github.com/dheemanth/container/entities"
	//elasticsearch "github.com/elastic/go-elasticsearch"
)

const (
	Url       = "http://localhost:9200/"
	Index     = "employees/"
	post_data = "_doc"
	search    = "_search"
)

type Job_history struct {
	Company string `json:"Company"`
	Role    string `json:"Role"`
}

type Employee struct {
	Id          string      `json:"Emp_ID,omitempty"`
	Name        string      `json:"Name"`
	Age         int         `json:"Age"`
	Designation string      `json:"Designation"`
	Email_id    string      `json:"Email_id"`
	Experience  float32     `json:"Experience"`
	Job_history Job_history `json:"Job_history"`
}

type coredata struct {
	Index  string   `json:"_index"`
	Type   string   `json:"_type"`
	Id     string   `json:"_id"`
	Score  float32  `json:"_score"`
	Source Employee `json:"_source"`
}

type hits struct {
	Total     int        `json:"total"`
	Max_score float32    `json:"max_score"`
	Hits      []coredata `json:"hits"`
}

type ESResponse struct {
	Took      int    `json:"took"`
	Timed_out bool   `json:"timed_out"`
	Shards    Shards `json:"_shards"`
	Hits      hits   `json:"hits"`
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type Ids struct {
	ID []string `json:"_id"`
}

type terms struct {
	Terms Ids `json:"terms"`
}

type Query struct {
	Query terms `json:"query"`
}

func getById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Performing GET_BY_ID OPERATIONS")
	url_list := strings.Split(r.URL.String(), "/")
	ids_list := strings.Split(url_list[len(url_list)-1], ",")

	query := Query{
		Query: terms{
			Terms: Ids{
				ID: ids_list,
			},
		},
	}
	request_url := Url + Index + search
	json_query, _ := json.Marshal(query)

	client := &http.Client{}
	a := strings.NewReader(string(json_query))
	req, _ := http.NewRequest("GET", request_url, a)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error occured ", err)
		return
	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	//unmasrhalling from response
	var emps ESResponse
	error := json.Unmarshal(resp_body, &emps)
	if error != nil {
		fmt.Println("Error is ", error)
	} else {
		data := emps.Hits.Hits
		var emp_list []Employee
		for _, v := range data {
			emp_list = append(emp_list, v.Source)
		}
		fmt.Println(emp_list)
		fmt.Printf("%T\n", emp_list)
		json.NewEncoder(w).Encode(emp_list)
	}
	return
}

func getAll(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Performing GET_ALL OPERATIONS")
	request_url := Url + Index + search
	client := &http.Client{}

	req, _ := http.NewRequest("GET", request_url, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error is ", err)
		return
	}
	defer resp.Body.Close()

	response_body, _ := ioutil.ReadAll(resp.Body)

	//unmasrhalling from response
	var emps ESResponse
	error := json.Unmarshal(response_body, &emps)
	if error != nil {
		fmt.Println("Error is ", error)
	} else {
		data := emps.Hits.Hits
		var emp_list []Employee
		for _, v := range data {
			v.Source.Id = v.Id
			emp_list = append(emp_list, v.Source)
		}
		fmt.Println(emp_list)
		fmt.Printf("%T\n", emp_list)
		json.NewEncoder(w).Encode(emp_list)
	}
	return
}

func postData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Performing POST operation")
	data_to_post := r.Body
	decoder := json.NewDecoder(data_to_post)
	var emps Employee
	if err := decoder.Decode(&emps); err != nil {
		panic(err)
	}
	bs, err2 := json.Marshal(emps)
	if err2 != nil {
		panic(err2)
	}

	request_url := Url + Index + post_data
	fmt.Printf(request_url)
	client := &http.Client{}
	data_to := strings.NewReader(string(bs))
	fmt.Println(data_to)
	req, _ := http.NewRequest("POST", request_url, data_to)
	req.Header.Add("Content-Type", "application/json")
	res, err3 := client.Do(req)
	if err3 != nil {
		panic(err3)
	} else {
		fmt.Println(res.Status)
		json.NewEncoder(w).Encode("Data has been Posted")
	}
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Performing PUT operation")
	url_list := strings.Split(r.URL.String(), "/")
	fmt.Println(url_list)
	data_to_post := r.Body
	decoder := json.NewDecoder(data_to_post)
	var emps Employee
	if err := decoder.Decode(&emps); err != nil {
		panic(err)
	}
	bs, err2 := json.Marshal(emps)
	if err2 != nil {
		panic(err2)
	}

	request_url := Url + Index + post_data + "/"+ string(url_list[len(url_list)-1])
	fmt.Printf(request_url)
	client := &http.Client{}
	data_to := strings.NewReader(string(bs))
	fmt.Println(data_to)
	req, _ := http.NewRequest("PUT", request_url, data_to)
	req.Header.Add("Content-Type", "application/json")
	res, err3 := client.Do(req)
	if err3 != nil {
		panic(err3)
	} else {
		fmt.Println(res.Status)
		json.NewEncoder(w).Encode("Data has been Posted")
	}
}

func main() {

	// Starting a server using mux
	server := mux.NewRouter()

	//Function Handlers
	server.HandleFunc("/employees/{ids}", getById).Methods("GET")
	server.HandleFunc("/employees/", getAll).Methods("GET")
	server.HandleFunc("/employees/", postData).Methods("POST")
	server.HandleFunc("/employees/{id}", UpdateEmployee).Methods("PUT")
	http.ListenAndServe(":7000", server)

}

// ids, ok := r.URL.Query()["ids"]
// if !ok || ids == nil {
// 	panic("Missing Ids")
// }
// fmt.Println(ids)
