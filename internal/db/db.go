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
	_ "github.com/go-sql-driver/mysql" // MySQL driver.
	"github.com/jmoiron/sqlx"

	"github.com/superhero-match/consumer-update-media/internal/db/model"
)

// DB interface defines database methods.
type DB interface {
	StoreProfilePicture(pp model.ProfilePicture) error
}

// db holds the database connection.
type db struct {
	dtbs                        *sqlx.DB
	stmtInsertNewProfilePicture *sqlx.Stmt
}

// New returns database.
func New(dtbs *sqlx.DB) (dbs DB, err error) {
	stmtIns, err := dtbs.Preparex("call insert_new_profile_picture(?,?,?,?)")
	if err != nil {
		return nil, err
	}

	return &db{
		dtbs:                        dtbs,
		stmtInsertNewProfilePicture: stmtIns,
	}, nil
}
