package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Cow holds the structure for the cow collection in mongo
type Cow struct {
	ID      string     `json:"_id" bson:"_id"` // MongoDB ID
	Details CowDetails `json:"cow" bson:"Cow"` // Details
}

// BookDetails holds the checkout details
type BookDetails struct {
	ID        string             `json:"id"        bson:"ID"`        // Generated ID -> cow_id.random_str
	Author    string             `json:"author"    bson:"Author"`    // User who booked
	Devices   []string           `json:"devices"   bson:"Devices"`   // Array of device ID's
	Block     string             `json:"block"     bson:"Block"`     // Block that is booked
	StartDate primitive.DateTime `json:"startdate" bson:"StartDate"` // Date this booking occurs
	EndDate   primitive.DateTime `json:"enddate"   bson:"EndDate"`   // Date this booking ends
}

// CowDetails holds the structure for the inner cow structure as
// defined in the cow collection in mongo
type CowDetails struct {
	Name        string        `json:"name"        bson:"Name"`        // eg. CA-01
	Collection  string        `json:"collection"  bson:"Collection"`  // eg. Laptop, Ipad, etc
	DeviceTotal int           `json:"deviceTotal" bson:"DeviceTotal"` // # of devices in that cart collection
	Bookings    []BookDetails `json:"bookings"    bson:"Bookings"`    // An array of all active bookings (send top 10)
	Devices     []string      `json:"devices"     bson:"Devices"`     // Array of device ID's
}
