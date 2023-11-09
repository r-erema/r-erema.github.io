package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("./tpl/layout.html")
	if err != nil {
		log.Panic(err)
	}

	goFile, err := os.Create("go.html")
	if err != nil {
		log.Panic(err)
	}

	devopsFile, err := os.Create("devops.html")
	if err != nil {
		log.Panic(err)
	}

	if err = tpl.Execute(devopsFile, template.FuncMap{
		"DevOps": true,
	}); err != nil {
		panic(err)
	}

	if err = tpl.Execute(goFile, template.FuncMap{
		"DevOps": false,
	}); err != nil {
		panic(err)
	}
}
