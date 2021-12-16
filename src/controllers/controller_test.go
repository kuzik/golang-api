package controllers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/url-builder/go-admin/src"
	"gitlab.com/url-builder/go-admin/src/config"
	"gitlab.com/url-builder/go-admin/src/database/connection"
	"gitlab.com/url-builder/go-admin/src/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const adminToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwidXNlcl9pZCI6MSwiZXhwIjoxODkzNDU2MDAwLCJpc3MiOiJnaW4tYmxvZyJ9.z3LuwJwFvtMrs8sD_xUQd9lDck5KaWUZT0aqV5TM47Q"

func before() *gorm.DB {
	configs := config.LoadConfigs("../../.env.test")
	dbConnection, err := gorm.Open(mysql.Open(configs.DB.AsDsnNoDb()), &gorm.Config{})
	rs := dbConnection.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", configs.DB.Name))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Database successfully connected")
	}

	if rs.Error != nil {
		fmt.Println(rs.Error)
		return nil
	}

	dbConnection, _ = gorm.Open(mysql.Open(configs.DB.AsDsn()), &gorm.Config{})

	err = connection.Migrate(dbConnection)
	if err != nil {
		fmt.Println("Migration error")
		return nil
	}

	// Add admin user
	dbConnection.Exec(fmt.Sprintf(
		"INSERT INTO users (username, password) VALUES ('%s', '%s')",
		"akuzo",
		"7f777bccedb2356c3ebab81ddf752e65664ee4248e356e25c3310919860eeb02",
	))

	fmt.Println("Prev done")
	return dbConnection
}

func after(dbConnection *gorm.DB) {
	configs := config.LoadConfigs("../../.env.test")
	rs := dbConnection.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", configs.DB.Name))

	if rs.Error != nil {
		fmt.Println(rs.Error)
		return
	}
}

func TestGetUrlsEmpty(t *testing.T) {
	a := assert.New(t)

	dbConnection := before()
	if dbConnection == nil {
		t.Error("DB connection failure")
	}

	r := src.BootStrap("../../.env.test")

	defer after(dbConnection)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/url/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", adminToken)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	res := make([]models.Url, 0)
	expected := make([]models.Url, 0)

	err := json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		t.Error("Problem with parsing response")
	}

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")
	a.Equal(expected, res, "HTTP response does not match")
}

func TestGetUrls(t *testing.T) {
	a := assert.New(t)

	dbConnection := before()
	if dbConnection == nil {
		t.Error("DB connection failure")
	}

	r := src.BootStrap("../../.env.test")

	defer after(dbConnection)

	dbConnection.Exec("INSERT INTO urls (user_id, label, destination) VALUES ( 1, 'Label', 'dest')")

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/url/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", adminToken)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	res := make([]models.Url, 0)
	expected := []models.Url{
		models.Url{
			Model:       models.Model{ID: 1, CreatedOn: 0, ModifiedOn: 0, DeletedOn: 0},
			UserID:      1,
			Label:       "Label",
			Destination: "dest",
			User:        models.User{ID: 0, Username: "", Password: ""},
			Campaigns:   []models.Campaign(nil),
		},
	}

	err := json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		t.Error("Problem with parsing response")
		t.Error(err)
	}

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")
	a.Equal(expected, res, "HTTP response does not match")
}
