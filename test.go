package main

import (
	"fmt"
	"ImageServer/src/net"
	_"github.com/gorilla/mux"
	"net/http"
	"flag"
)

/*func main() {
	fmt.Println("主程序执行了")
	//net.UploadFile(nil, nil);
	r := mux.NewRouter()
	r.HandleFunc("/", net.UploadFile).Methods("POST")
	r.HandleFunc("/img",net.LoadImg).Methods("GET")
	//r.HandleFunc("/{path}",net.LoadImg).Methods("GET")
	err := http.ListenAndServe("0.0.0.0:10086", r)
	if err != nil {
		fmt.Println(err)
	}
}*/

func main() {
	realPath := flag.String("path", "", "static resource path")
	fmt.Println(realPath)
	flag.Parse()
	http.HandleFunc("/img/", net.LoadImg)
	http.HandleFunc("/", net.UploadFile)

	err := http.ListenAndServe(":10086", nil)
	if err != nil {
		//log.Fatal("ListenAndServe:", err)
		fmt.Println(err)
	}
}

