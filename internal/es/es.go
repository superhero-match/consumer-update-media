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
	elastic "github.com/olivere/elastic/v7"

	"github.com/superhero-match/consumer-update-media/internal/es/model"
)

// ES interface defines es methods.
type ES interface {
	GetSuperhero(superheroID string) (s *model.Superhero, err error)
	GetDocumentID(superheroID string) (string, error)
	UpdateProfilePicture(superheroID string, pp model.ProfilePicture) error
}

// es holds all the Elasticsearch client relevant data.
type es struct {
	Client *elastic.Client
	Index  string
}

// New returns value of type es which implements all ES methods.
func New(client *elastic.Client, index string) ES {
	return &es{
		Client: client,
		Index:  index,
	}
}
