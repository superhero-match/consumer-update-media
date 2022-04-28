/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package es

import (
	"net/http"
	"net/http/httptest"

	"github.com/olivere/elastic/v7"
)

type DummyHttpClient struct {
	responseMock string
}

func (c *DummyHttpClient) Do(r *http.Request) (*http.Response, error) {
	recorder := httptest.NewRecorder()
	recorder.Write([]byte(c.responseMock))
	recorder.Header().Set("Content-Type", "application/json")

	return recorder.Result(), nil
}

func MockHttpClient(responseMock string) *DummyHttpClient {
	return &DummyHttpClient{responseMock}
}

func MockElasticSearchClient(endpoint string, responseMock string) (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(endpoint),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetHttpClient(MockHttpClient(responseMock)))

	return client, err
}
