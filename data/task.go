package data

import (
	"gopkg.in/mgo.v2"
	"github.com/mhhg/wms/models"
)

type TaskRepository {
	C *mgo.Collection
}

func (r *TaskRepository) Create(task *models.Task) error {
	
}

func (r *TaskRepository) Update(task *models.Task) error {

}

func (r *TaskRepository) GetAll(task *models.Task) error {

}

func (r *TaskRepository) GetById(task *models.Task) error {

}

func (r *TaskRepository) DeleteById(task *models.Task) error {
	
}