/*
  Copyright (C) 2019 - 2021 MWSOFT
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

	"github.com/superhero-match/consumer-update-media/internal/es/model"
)

const mainProfilePicturePosition = int64(0)

// UpdateProfilePicture updates profile picture.
func (es *es) UpdateProfilePicture(superheroID string, pp model.ProfilePicture) error {
	superhero, err := es.GetSuperhero(superheroID)
	if err != nil {
		return err
	}

	sourceID, err := es.GetDocumentID(superheroID)
	if err != nil {
		return err
	}

	if pp.Position == mainProfilePicturePosition {
		return updateMainProfilePic(es, sourceID, pp)
	}

	if len(superhero.ProfilePictures) == 0 {
		pps := make([]model.ProfilePicture, 0)
		pps = append(pps, pp)

		return updateProfilePics(es, sourceID, pps, pp)
	}

	for i := 0; i < len(superhero.ProfilePictures); i++ {
		if pp.Position == superhero.ProfilePictures[i].Position {
			superhero.ProfilePictures = append(superhero.ProfilePictures[:i], superhero.ProfilePictures[i+1:]...)
		}
	}

	superhero.ProfilePictures = append(superhero.ProfilePictures, pp)

	return updateProfilePics(es, sourceID, superhero.ProfilePictures, pp)
}

func updateProfilePics(es *es, sourceID string, pps []model.ProfilePicture, pp model.ProfilePicture) error {
	_, err := es.Client.Update().
		Index(es.Index).
		Id(sourceID).
		Doc(map[string]interface{}{
			"profile_pics": pps,
			"updated_at":   pp.CreatedAt,
		}).
		Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func updateMainProfilePic(es *es, sourceID string, pp model.ProfilePicture) error {
	_, err := es.Client.Update().
		Index(es.Index).
		Id(sourceID).
		Doc(map[string]interface{}{
			"main_profile_pic_url": pp.URL,
			"updated_at":           pp.CreatedAt,
		}).
		Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
