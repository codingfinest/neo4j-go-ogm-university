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

package events

import (
	"github.com/codingfinest/neo4j-go-ogm-university/entities"

	gogm "github.com/codingfinest/neo4j-go-ogm"
)

//AppEventListener is the event listener for ogm events
type AppEventListener struct{}

//OnPreSave is called before save
func (e AppEventListener) OnPreSave(event gogm.Event) {
	switch object := event.GetObject().(type) {
	case *entities.StudyBuddy:
		object.Buddies = nil
		if object.BuddyOne != nil {
			object.Buddies = append(object.Buddies, object.BuddyOne)
		}
		if object.BuddyTwo != nil {
			object.Buddies = append(object.Buddies, object.BuddyTwo)
		}
	}
}

//OnPostSave is called after save
func (e AppEventListener) OnPostSave(event gogm.Event) {}

//OnPostLoad is called after load
func (e AppEventListener) OnPostLoad(event gogm.Event) {
	switch object := event.GetObject().(type) {
	case *entities.StudyBuddy:
		object.BuddyOne = nil
		object.BuddyTwo = nil
		if len(object.Buddies) > 0 {
			for index, buddy := range object.Buddies {
				if index == 0 {
					object.BuddyOne = buddy
				} else {
					object.BuddyTwo = buddy
				}
			}
		}
	}
}

//OnPreDelete is called before delete
func (e AppEventListener) OnPreDelete(event gogm.Event) {}

//OnPostDelete is called after delete
func (e AppEventListener) OnPostDelete(event gogm.Event) {}
