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

package db

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"

	"github.com/superhero-match/consumer-update-media/internal/db/model"
)

func TestDb_StoreProfilePicture(t *testing.T) {
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	mock.ExpectPrepare(`call insert_new_profile_picture(?,?,?,?)`).
		ExpectExec().
		WithArgs("123456789", "https://www.test.com/1.jpg", 2, "2022-04-26T12:00:00").
		WillReturnResult(sqlmock.NewResult(3, 1))

	db, err := New(sqlxDB)
	if err != nil {
		t.Fatal(err)
	}

	pp := model.ProfilePicture{
		SuperheroID: "123456789",
		URL:         "https://www.test.com/1.jpg",
		Position:    2,
		CreatedAt:   "2022-04-26T12:00:00",
	}

	err = db.StoreProfilePicture(pp)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDb_StoreProfilePictureFail(t *testing.T) {
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	mock.ExpectPrepare(`call insert_new_profile_picture(?,?,?,?)`).
		ExpectExec().
		WithArgs("123456789", "https://www.test.com/1.jpg", 2, "2022-04-26T12:00:00").
		WillReturnError(fmt.Errorf("testing failure"))

	db, err := New(sqlxDB)
	if err != nil {
		t.Fatal(err)
	}

	pp := model.ProfilePicture{
		SuperheroID: "123456789",
		URL:         "https://www.test.com/1.jpg",
		Position:    2,
		CreatedAt:   "2022-04-26T12:00:00",
	}

	err = db.StoreProfilePicture(pp)
	if err == nil {
		t.Fatal(err)
	}
}
