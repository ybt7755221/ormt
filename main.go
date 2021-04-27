package main

import (
	"ormt/core"

	"github.com/droundy/goopt"
)

var (
	account  = goopt.String([]string{"-a", "--account"}, "", "mysql account. format - 'username:password'")
	host     = goopt.String([]string{"-h", "--host"}, "", "mysql host. format - 'host:port'")
	database = goopt.String([]string{"-d", "--database"}, "", "mysql database")
)

func init() {
	goopt.Description = func() string {
		return "ormt is automaticlly generate mysql model"
	}
	goopt.Version = "0.1"
	goopt.Summary = `ormt --account account --host host --database dbName`
	goopt.Parse(nil)
}

func main() {
	core.Generate()
}
