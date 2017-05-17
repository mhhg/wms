package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Task jobs which should be done based on the contract
	Task struct {
		ID          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
		Name        string        `json:"name,omitempty"`
		No          int           `json:"no,omitempty"`
		Description string        `json:"description,omitempty"`

		FlowID   bson.ObjectId `json:"flow_id,omitempty"`
		PoID     bson.ObjectId `json:"po_id,omitempty"`
		ParentID bson.ObjectId `json:"parent_id,omitempty"`
		SiteID   bson.ObjectId `json:"site_id,omitempty"`

		Type       string `json:"type,omitempty"`
		Status     string `json:"status,omitempty"`
		Value      int    `json:"value,omitempty"`
		Priority   string `json:"priority,omitempty"`
		AssignType string `json:"assign_type,omitempty"`

		Quantity    string  `json:"quantity,omitempty"`
		Line        int     `json:"line,omitempty"`
		Item        string  `json:"item,omitempty"`
		ShipmentNum string  `json:"shipment_num,omitempty"`
		Unit        string  `json:"unit,omitempty"`
		SubTotal    float64 `json:"sub_total,omitempty"`

		StartDate time.Time `json:"start_date,omitempty"`
		EndDate   time.Time `json:"end_date,omitempty"`
		Due       time.Time `json:"due,omitempty"`
	}
	// Flow project or process
	Flow struct {
		ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
		VendorID bson.ObjectId
		No       string
		Name     string
		Parent   bson.ObjectId
		Status   string
		Type     string
	}
	// PO Purchase Order between two company
	PO struct {
		ID          bson.ObjectId `json:"id,omitempty"`
		Name        string
		Description string

		SubconNo  string
		No        string
		ProjectNo string
		Version   string
		Value     string
		Supplier  string
	}
	// Site telecommunications device
	Site struct {
		ID          bson.ObjectId
		Name        string
		Description string
		City        string
		Province    string
		Latitude    float32
		Longitude   float32
		Tech2G      bool
		Tech3G      bool
		Tech4G      bool
		CityID      bson.ObjectId
		ProvinceID  bson.ObjectId
	}
	// Vendor business company
	Vendor struct {
		ID   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
		Name string        `json:"name,omitempty"`
		No   string        `json:"no,omitempty"`
	}
	// User user specification of the app
	User struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Username  string
		Password  string
		Firstname string
		Lastname  string
		Email     string
		Role      string
		Created   time.Time
	}
	Role struct {
		ID bson.ObjectId `bson:"_id,omitempty" json:"id"`
	}
	Permission struct {
		ID bson.ObjectId `bson:"_id,omitempty" json:"id"`
	}
	// Note content which is saved by user on task or flow or po
	Note struct {
		ID           bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
		Content      string        `json:"content,omitempty"`
		Type         string        `json:"type,omitempty"`
		UserID       bson.ObjectId `json:"user_id,omitempty"`
		RefID        bson.ObjectId `json:"ref_id,omitempty"`
		Origin       string        `json:"origin,omitempty"` // task, flow, po, .. : collection name
		Availability bool          `json:"availability,omitempty"`
		Date         time.Time     `json:"date,omitempty"`
	}
	// History log of datetime of create/update/delete a collection and the user who done that
	History struct {
		ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
		UserID bson.ObjectId
		RefID  bson.ObjectId
		Origin string
		Type   string // create, update, delete
		On     time.Time
	}
	Event struct {
		ID bson.ObjectId `bson:"_id,omitempty" json:"id"`
	}
)
