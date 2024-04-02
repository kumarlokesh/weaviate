//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package api

import (
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/usecases/sharding"
)

type AddClassRequest struct {
	Class *models.Class
	State *sharding.State
}

type UpdateClassRequest struct {
	Class *models.Class
	State *sharding.State
}

type AddPropertyRequest struct {
	Properties []*models.Property
}

type DeleteClassRequest struct {
	Name string
}

type UpdateShardStatusRequest struct {
	Class, Shard, Status string
}

type QueryReadOnlyClassRequest struct {
	Class string
}

type QueryReadOnlyClassResponse struct {
	Class *models.Class
	State *sharding.State
}

type QueryGetSchemaResponse struct {
	Schema models.Schema
}