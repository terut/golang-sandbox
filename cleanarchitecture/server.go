package main

import "github.com/terut/golang-sandbox/cleanarchitecture/infrastructure"

func main() {
	infrastructure.Router.Start(":8080")
}
