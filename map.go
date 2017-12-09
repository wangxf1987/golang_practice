package main

import (
    "fmt"
)

type cnmNetworkAllocator struct {
    networks map[string]*network
}
type network struct {
    TaskIDMap map[string]string
}

func main() {
    na := &cnmNetworkAllocator{
        networks: make(map[string]*network),
    }
    nw := &network{
        TaskIDMap: make(map[string]string),
    }
    nw.TaskIDMap["netid01"] = "netid01"
    nw.TaskIDMap["ip01"] = "ip01"
    nw.TaskIDMap["netid02"] = "netid02"
    nw.TaskIDMap["ip02"] = "ip02"
    na.networks["taskid01"] = nw
    nw.TaskIDMap["netid03"] = "netid03"
    nw.TaskIDMap["ip03"] = "ip03"
    nw.TaskIDMap["netid04"] = "netid04"
    nw.TaskIDMap["ip04"] = "ip04"
    na.networks["taskid02"] = nw
    fmt.Println(na.networks)
    for k,v := range na.networks {
        fmt.Println(k)
        fmt.Println(v)
    }
}
