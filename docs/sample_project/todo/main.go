package main

import (
	"net/http"

	"github.com/uadmin/uadmin"
	"github.com/uadmin/uadmin/docs/sample_project/todo/api"
	"github.com/uadmin/uadmin/docs/sample_project/todo/models"
	"github.com/uadmin/uadmin/docs/sample_project/todo/templates/custom"
)

func main() {
	uadmin.RootURL = "/admin/"
	uadmin.Register(
		models.Todo{},
		models.Category{},
		models.Friend{},
		models.Item{},
	)

	uadmin.RegisterInlines(models.Category{}, map[string]string{
		"Todo": "CategoryID",
	})
	uadmin.RegisterInlines(models.Friend{}, map[string]string{
		"Todo": "FriendID",
	})
	uadmin.RegisterInlines(models.Item{}, map[string]string{
		"Todo": "ItemID",
	})

	// API Handler
	http.HandleFunc("/api/", api.APIHandler)

	// Template Handler
	http.HandleFunc("/template/", templates.TemplateHandler)

	uadmin.Port = 8000
	// uadmin.StartServer()
	uadmin.StartSecureServer("pub.pem", "priv.pem")
}
