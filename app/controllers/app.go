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
	"github.com/codingfinest/neo4j-go-ogm-university/container"
	"github.com/codingfinest/neo4j-go-ogm-university/events"
	"github.com/codingfinest/neo4j-go-ogm-university/services"
)

var store = services.NewStore()

//Mapping maps entity api endpoints to their controllers
var Mapping = map[string]entity{
	"classes":      &course{},
	"departments":  &department{},
	"enrollments":  &enrollment{},
	"schools":      &school{},
	"students":     &student{},
	"studyBuddies": &StudyBuddy{},
	"subjects":     &subject{},
	"teachers":     &teacher{}}

func init() {
	for name, controller := range Mapping {
		controller.setContainerProvider(container.Mapping[name])
	}

	appEventListener := events.AppEventListener{}
	store.RegisterEventListener(appEventListener)
}
