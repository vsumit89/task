package util

import "log"

func LogRequest(method string, path string, status int){
	log.Println("METHOD:",method,"PATH:", path, "STATUS:",status)
}