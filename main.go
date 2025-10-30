package main

import (
	"github.com/Adeshss20/ecommerce-platform-golang.git/internal/router"
)

func main() {
    r := router.SetupRouter()
	r.Run(":8080")
}
