package main

import "net/http"

func App(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Consumer API Gateway"))
}

func main()  {
	println("Gateway running on: localhost:8080")

	http.HandleFunc("/", App)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
