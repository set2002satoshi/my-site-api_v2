package main

import "github.com/set2002satoshi/my-site-api_v2/infrastructure"

func main() {
	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	db.DBInit()
	r.Run()
}