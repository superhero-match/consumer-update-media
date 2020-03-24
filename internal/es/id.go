/*
  Copyright (C) 2019 - 2020 MWSOFT
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
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
)

// GetDocumentID fetches document source id, it is needed for update function.
func (es *ES) GetDocumentID(superheroID string) (string, error) {
	q := elastic.NewTermQuery("superhero_id", superheroID)

	searchResult, err := es.Client.Search().
		Index(es.Index).
		Query(q).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return "", err
	}

	if searchResult.TotalHits() > 0 {
		for _, hit := range searchResult.Hits.Hits {
			return hit.Id, nil
		}
	}

	return "", errors.New("no result")
}
