package main

import (
	"github.com/lottotto/sample-clean-architecture/infrastructure"
)

func main() {
	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	r.Run()
}
