package data

import (
	"github.com/mhhg/wms/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TaskRepository mongodb repository for Task
type TaskRepository struct {
	C *mgo.Collection
}

// Create insert a new task document
func (r *TaskRepository) Create(task *models.Task) error {
	task.ID = bson.NewObjectId()
	err := r.C.Insert(
		bson.M{
			"name":         task.Name,
			"no":           task.No,
			"description":  task.Description,
			"flow_id":      task.FlowID,
			"po_id":        task.PoID,
			"type":         task.Type,
			"status":       task.Status,
			"value":        task.Value,
			"priority":     task.Priority,
			"assign_type":  task.AssignType,
			"quantity":     task.Quantity,
			"line":         task.Line,
			"shipment_num": task.ShipmentNum,
			"unit":         task.Unit,
			"sub_total":    task.SubTotal,
			"start_date":   task.StartDate,
			"end_date":     task.EndDate,
		})
	return err
}

// Update a task document
func (r *TaskRepository) Update(task *models.Task) error {
	err := r.C.UpdateId(task.ID,
		bson.M{"$set": bson.M{
			"name":         task.Name,
			"no":           task.No,
			"description":  task.Description,
			"flow_id":      task.FlowID,
			"po_id":        task.PoID,
			"type":         task.Type,
			"status":       task.Status,
			"value":        task.Value,
			"priority":     task.Priority,
			"assign_type":  task.AssignType,
			"quantity":     task.Quantity,
			"line":         task.Line,
			"shipment_num": task.ShipmentNum,
			"unit":         task.Unit,
			"sub_total":    task.SubTotal,
			"start_date":   task.StartDate,
			"end_date":     task.EndDate,
		}})
	return err
}

// Delete a task document by id
func (r *TaskRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// GetAll get all tasks
func (r *TaskRepository) GetAll() []models.Task {
	var tasks []models.Task
	iter := r.C.Find(nil).Iter()
	result := models.Task{}
	for iter.Next(&result) {
		tasks = append(tasks, result)
	}
	return tasks
}

// GetByID return task document by id
func (r *TaskRepository) GetByID(id string) (task models.Task, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&task)
	return
}
