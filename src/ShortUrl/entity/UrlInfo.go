package entity

import "time"

type Urlinfo struct {
	Id         int64
	Surl       string
	Lurl       string
	Views      int
	Createtime time.Time
}
