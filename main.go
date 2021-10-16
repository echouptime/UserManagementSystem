package main

import (
	"UserManagementSystem/routers"
	"fmt"
	"net/http"
)

func main() {
	Addr := ":8082"
	routers.Register()
	if err := http.ListenAndServe(Addr, nil); err != nil {
		fmt.Println(err)
	}

}
