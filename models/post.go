package models

import "time"

type Post struct {
	Id      int64
	Title   string `sql:"size:255"`
	Author  string `sql:"size:255"`
	Body    string `sql:"size:65535"`
	Written time.Time
}

func (this *Post) String() string {
	return "title=" + this.Title + "|author=" + this.Author + "|body=" + this.Body
}
