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

	"github.com/codingfinest/neo4j-go-ogm-university/entities"
	"github.com/labstack/echo/v4"
)

//StudyBuddy is StudyBuddy domain entity
type StudyBuddy struct {
	baseEntity
}

//GetStudyBuddiesByPopularity returns the buddies by popularity
func (studyBuddy StudyBuddy) GetStudyBuddiesByPopularity(c echo.Context) error {

	query := "MATCH (s:StudyBuddy)<-[:BUDDY]-(p:Student) return p, count(s) as buddies ORDER BY buddies DESC"

	var student *entities.Student
	result, err := store.Query(query, nil, &student)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
