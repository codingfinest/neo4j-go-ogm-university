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
	"io/ioutil"
	"os"
)

// Reload purges the database and reloads the initial school data
func (store *Store) Reload() error {
	if err := store.session.PurgeDatabase(); err != nil {
		return err
	}

	ogmExampleSchoolCqlPath := os.Getenv("OGM_EAMPLE_SCHOOL_CQL_PATH")
	if len(ogmExampleSchoolCqlPath) == 0 {
		panic("File path for initial database state must be set in `OGM_EAMPLE_SCHOOL_CQL_PATH` env variable. Default is <PROJECT_HOME>/conf/school.cql")
	}
	if _, err := store.session.Query(reloadCypher(ogmExampleSchoolCqlPath), nil); err != nil {
		return err
	}
	return nil
}

func reloadCypher(cqlFile string) string {
	content, _ := ioutil.ReadFile(cqlFile)
	return string(content)
}
