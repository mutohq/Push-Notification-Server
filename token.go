package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"github.com/muto_gcm"
	"github.com/parnurzeal/gorequest"
)

var db *pgx.ConnPool
var db_err error

func init() {
	db, db_err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "localhost",
			Database: "pmp",
			User:     "shivam_mac",
			Password: "",
			Port:     5432,
		},
		MaxConnections: 10,
	})

	if db_err != nil {
		log.Fatal(db_err)
	}
}

func main() {

	r := gin.Default()

	r.POST("/token", func(c *gin.Context) {

		var request muto_gcm.Request
		var combined muto_gcm.Combined
		c.BindJSON(&request)
		fmt.Println("\n\n\n Received : %#v\n\n\n", request)

		rows, err := db.Query(`
                SELECT usersdescription.deviceid, usersdescription.platform
                FROM users, usersdescription
                WHERE usersdescription.userid = users.userid AND username = $1 AND email = $2
    `, request.Username, request.Email)

		if err != nil {
			responseBack := make(map[string]error)
			responseBack["Error : "] = err
			c.JSON(500, responseBack)
			log.Fatal(err)
		}

		defer rows.Close()

		var deviceID, platform string
		var device muto_gcm.Device

		for rows.Next() {
			if err := rows.Scan(&deviceID, &platform); err != nil {
				responseBack := make(map[string]error)
				responseBack["Error : "] = err
				c.JSON(500, responseBack)
				log.Fatal(err)
			}
			device.DeviceID = deviceID
			device.Platform = platform
			combined.DevicesList = append(combined.DevicesList, device)
		}

		b, _ := json.Marshal(combined)

		requestHTTP := gorequest.New()
		resp, body, errs := requestHTTP.Post("http://localhost:9011/send").
			Set("Content-Type", "application/json").
			Send(string(b)).
			End()
		fmt.Println(resp, body, errs)

	})

	fmt.Println("Going Live on : 9010\n")
	r.Run(":9010")
}
