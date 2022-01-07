package main

import (
	"log"
)

func init() {
	log.SetPrefix("【日志】")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {

}
