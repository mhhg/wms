package controllers

import "github.com/mhhg/wms/models"

// Models for JSON resources
type (
	// TaskResource For Post/Put - /tasks
	// TaskResource For Get - /tasks/:id
	TaskResource struct {
		Data models.Task `json:"data"`
	}
)
