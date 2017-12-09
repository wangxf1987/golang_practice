package main

import (
    "fmt"
    "strconv"
    "reflect"
)

type NetIDMap struct {
    ID string `json: "Network ID"`
    IP string `json: "IP address for contianer"`
}

type network struct {
    TaskID string `json: "Task ID"`
    NetIDMap []NetIDMap `json: "Netwroks"`
}


func main() {
    var Ns []network
    //Add
    for j := 1; j <= 2; j++ {
        var N network
        var temp_NetIDMap []NetIDMap
        for i := 1; i <= 3; i++ {
            temp_NetIDMap = append(temp_NetIDMap, NetIDMap{
                      ID: strconv.Itoa(i),
                      IP: strconv.Itoa(i),})
        }
        TaskID := strconv.Itoa(j)
        N.TaskID = TaskID
        N.NetIDMap = temp_NetIDMap
        Ns = append(Ns, N)
    }
    //fmt.Println(Ns)
    //research
    for m := 0; m < len(Ns); m++ {
        //fmt.Println(Ns[m])
        tmp_m := Ns[m]
        //fmt.Println(tmp_m)
       t := reflect.TypeOf(tmp_m)
       v := reflect.ValueOf(tmp_m)
       for k := 0; k < t.NumField(); k++ {
           fmt.Printf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())   
           tmp_v := v.Field(k).Interface()
           fmt.Println(tmp_v)
           //for p := 0; p < m.NumField(); p++ {
           //    fmt.Printf("%s -- %v \n", m.Field(p).Name, n.Field(p).Interface())
           //}
           //for m := 0; m < len(tmp_v); m++ {
           //    fmt.Printf(tmp_v[m])
           //}
        }
    }
}

