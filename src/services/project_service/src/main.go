package main

import (
	"fmt"
	"flag"
	"os"

	_ "github.com/amthorn/task_wizard_common"
)

func main() {
	flag.Parse()
	server := &Server{}
	server.serve(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
}