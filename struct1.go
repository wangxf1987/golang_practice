package main

import (
    "fmt"
    "strconv"
    //"reflect"
)

//type NetIDMap struct {
//    ID string `json: "Network ID"`
//    IP string `json: "IP address for contianer"`
//}
type network struct {
    TaskID map[string]string `json: "Task ID"`
}

func main() {
    ns := &network{
        TaskID: make(map[string]string),
    }
    for i :=0; i <= 3; i++ {
        ns.TaskID["id"] = strconv.Itoa(i)
    }
    fmt.Println(ns)
}
