// MIT License
//
// Copyright (c) 2020 codingfinest
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package services

import (
	gogm "github.com/codingfinest/neo4j-go-ogm"
)

var findOptions = gogm.NewLoadOptions()
var findAllOptions = gogm.NewLoadOptions()
var saveOptions = gogm.NewSaveOptions()

func init() {
	findAllOptions.Depth = 0
	findOptions.Depth = 1
	saveOptions.Depth = 1
}

//Store provides access to the backend database
type Store struct {
	session gogm.Session
}

//Find finds an entity and stores it in the container if found
func (store Store) Find(container interface{}, id int64) error {
	if err := store.session.Load(container, id, findOptions); err != nil {
		return err
	}
	return nil
}

//FindAll finds all instances of the type of struct contained in the
//container. These instances are stored in containers
func (store Store) FindAll(containers interface{}) error {
	if err := store.session.LoadAll(containers, nil, findAllOptions); err != nil {
		return err
	}
	return nil
}

//CreateOrUpdate creates or updates the entity
func (store Store) CreateOrUpdate(entity interface{}) error {
	if err := store.session.Save(entity, saveOptions); err != nil {
		return err
	}
	return nil
}

//Delete deletes the entity
func (store Store) Delete(entity interface{}) error {
	if err := store.session.Delete(entity); err != nil {
		return err
	}
	return nil
}

//RegisterEventListener registers an OGM event listener
func (store Store) RegisterEventListener(el gogm.EventListener) error {
	return store.session.RegisterEventListener(el)
}

//Query queries the database
func (store Store) Query(cypher string, paramaters map[string]interface{}, queryDepedencies ...interface{}) ([]map[string]interface{}, error) {
	return store.session.Query(cypher, paramaters, queryDepedencies...)
}
