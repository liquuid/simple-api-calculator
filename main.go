//go:generate swagger generate spec
package main

import "fmt"

type Result struct {
	// the result
	//
	// required: true
	result int `json:"result`
}

// swagger:parameters add
type AddParam struct {
	// A param
	//
	// in: path
	// required: true
	// type: integer
	A int `json:"a"`

	// B param
	//
	// in: path
	// required: true
	// type: integer
	B int `json:"b"`
}

func main() {
	fmt.Println("ok")
	fmt.Println(add(1, 2))
}

// return sum of a and b parameters
//
// swagger:route GET /add/{a}/{b} add
//
// Produces:
// - application/json
//
// Consumes:
// - application/json
//
// Schemes: http
//
// Responses:
// default: error
// 200: okResp
//
func add(a int, b int) int {

	return a + b
}
