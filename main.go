package main

import (
	"fmt"
	"lvs_usermanage_app/routers"
	"net/http"
)

func main() {
	Addr := ":8082"
	routers.Register()
	if err := http.ListenAndServe(Addr, nil); err != nil {
		fmt.Println(err)
	}

}
