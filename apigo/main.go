http://146.59.204.94:8123http://146.59.204.94:8123package main

import (
  //  "bytes"
    //"encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

func main() {
    time.Sleep(10 * time.Second)
    fmt.Println("Starting the application...")
    response, err := http.Get("http://146.59.204.94:8123/ping")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
   // jsonData := map[string]string{"Age", "Income"}
    //jsonData = string
    //jsonValue, _ := json.Marshal(jsonData)
    //response, err = http.Post("http://127.0.0.1:8123/?query=SELECT%20Age,Income%20FROM%20datasets.visits_v1%20limit%205;&user=default&password=&database=datasets", "application/json", bytes.NewBuffer(jsonValue))
//response, err = http.Get("http://146.59.204.94:8123/?query=SELECT%20Age,Income%20FROM%20datasets.visits_v1%20limit%205;&user=default&password=&database=datasets")

response, err = http.Get("http://146.59.204.94:8123/?query=SELECT%20name,is_internal%20FROM%20system.aggregate_function_combinators%20limit%205;&user=default&password=&database=system")

    if err != nil {

        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
    fmt.Println("Terminating the application...")
	
	}
