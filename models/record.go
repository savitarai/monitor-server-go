package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type RecordRequest struct {
	Report RecordReport `json:"report"`
}
type RecordReport struct {
	Seq      int64      `json:"seq"`
	Hostname string     `json:"hostname"`
	Hostip   string     `json:"hostip"`
	Time_min int        `json:"time_min"`
	Commattr []CommAttr `json:"commattr"`
	Statattr []StatAttr `json:"statattr"`
}
type CommAttr struct {
	Attrid int   `json:"attr"`
	Value  int64 `json:"value"`
}

type StatAttr struct {
	Attrid int   `json:"attrid"`
	Min    int64 `json:"min"`
	Max    int64 `json:"max"`
	Count  int64 `json:"count"`
	Sum    int64 `json:"sum"`
}

func (r *RecordRequest) InertDB() error {
	db, err := sql.Open("mysql", "root:jitashan1984@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	return nil
}
