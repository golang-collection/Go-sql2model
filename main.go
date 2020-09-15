package main

import (
	"Go-sql2model/cmd"
	"log"
)

/**
* @Author: super
* @Date: 2020-09-15 08:33
* @Description:
**/

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}