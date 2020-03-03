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

package main

import (
	"github.com/codingfinest/neo4j-go-ogm-university/app/controllers"
	"github.com/labstack/echo/v4"
)

var entityController = controllers.Entity{}

func main() {
	e := echo.New()

	e.GET("api/reload", controllers.Import)

	studyBuddiesController := controllers.Mapping["studyBuddies"].(*controllers.StudyBuddy)
	e.GET("api/studyBuddies/popular", studyBuddiesController.GetStudyBuddiesByPopularity)

	e.GET("api/:entity", entityController.FindAll)
	e.GET("api/:entity/:id", entityController.Find)
	e.DELETE("api/:entity/:id", entityController.Delete)
	e.POST("api/:entity", entityController.CreateOrUpdate)

	e.Static("/app", "assets/javascripts/app")
	e.Static("/components", "assets/javascripts/components")

	e.Static("/jquery", "assets/libs/jquery")
	e.Static("/angular", "assets/libs/angular")
	e.Static("/angular-ui-router", "assets/libs/angular-ui-router")
	e.Static("/angular-resource", "assets/libs/angular-resource")
	e.Static("/angular-sanitize", "assets/libs/angular-sanitize")
	e.Static("/jsog", "assets/libs/jsog")
	e.Static("/bootstrap", "assets/libs/bootstrap")
	e.Static("/checklist-model", "assets/libs/checklist-model")

	e.Static("/stylesheets", "assets/stylesheets")
	e.Static("/img", "assets/images")

	e.Static("/html", "assets/html")

	e.Static("/", "app/views/public")

	e.Logger.Fatal(e.Start(":9000"))
}
