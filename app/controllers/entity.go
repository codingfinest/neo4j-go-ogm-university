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

package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//Entity is an general controller to handle CRUD requests for entities
type Entity struct {
}

//CreateOrUpdate create's or updates entity
func (e Entity) CreateOrUpdate(c echo.Context) error {
	var (
		entity interface{}
		err    error
	)
	entityController := Mapping[c.Param("entity")]
	if entityController == nil {
		return c.JSON(http.StatusMethodNotAllowed, nil)
	}

	if entity, err = entityController.getContainerProvider().GetContainer(c); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if err := store.CreateOrUpdate(entity); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, entity)
}

//Find find's an entity by id
func (e Entity) Find(c echo.Context) error {
	var (
		err    error
		id     int64
		entity interface{}
	)
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	entityController := Mapping[c.Param("entity")]

	if entityController == nil {
		return c.JSON(http.StatusMethodNotAllowed, nil)
	}

	if entity, err = entityController.getContainerProvider().GetContainer(nil); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if err = store.Find(entity, id); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, entity)
}

//FindAll find's all entity
func (e Entity) FindAll(c echo.Context) error {
	entityController := Mapping[c.Param("entity")]
	if entityController == nil {
		return c.JSON(http.StatusMethodNotAllowed, nil)
	}

	entities := entityController.getContainerProvider().GetContainers()
	if err := store.FindAll(entities); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, entities)
}

//Delete delete's entity by id
func (e Entity) Delete(c echo.Context) error {
	var (
		err    error
		id     int64
		entity interface{}
	)
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	entityController := Mapping[c.Param("entity")]

	if entityController == nil {
		return c.JSON(http.StatusMethodNotAllowed, nil)
	}

	if entity, err = entityController.getContainerProvider().GetContainer(nil); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if err = store.Find(entity, id); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if err = store.Delete(entity); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
