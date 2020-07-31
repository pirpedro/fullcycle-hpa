package main

import (
	"fmt"
	"net/http"
	"math"
)

func main() {
		http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
			for i := 0.0001; i <= 10; i++ {
					i += sqrt(i) 
			}
			fmt.Fprintf(w, "OK")
		})
		http.ListenAndServe(":8000", nil)
}

func sqrt(s float64) float64{
	return math.Sqrt(s)
}
