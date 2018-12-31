// Copyright 2018 Kuei-chun Chen. All rights reserved.

package mongox

import (
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Collection contains mongo.Collection
type Collection struct {
	collection *mongo.Collection
}

// Find finds docs by given filter
func (c *Collection) Find(filter interface{}) *Session {
	return &Session{filter: filter, collection: c.collection}
}
