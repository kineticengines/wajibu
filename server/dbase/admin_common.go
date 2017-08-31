/*
Wajibu is an online web app that collects,analyses, aggregates and visualizes sentiments
from the public pertaining the government of a nation. This tool allows citizens to contribute
to the governance talk by airing out their honest views about the state of the nation and in
particular the people placed in government or leadership positions.

Copyright (C) 2017
David 'Dexter' Mwangi
dmwangimail@gmail.com
https://github.com/daviddexter
https://github.com/daviddexter/wajibu

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

package dbase

import (
	"log"
	"sync"
	"time"

	cfg "github.com/daviddexter/wajibu/configure"
	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"
)

func GetTitlesFromDB() *types.TitleData {
	var r types.TitleData
	_, err := DB.Exec(`DESCRIBE ` + cfg.Loader().TitlesTable)
	switch err {
	case nil:
		//get titles
		rows, err := DB.Query(`SELECT titlename FROM ` + cfg.Loader().TitlesTable)
		defer rows.Close()
		if err == nil {
			for rows.Next() {
				var title string
				err = rows.Scan(&title)
				if err == nil {
					r.Title = append(r.Title, title)
				} else {
					r.Error = err
				}
			}
		} else {
			r.Error = err
		}
	default:
		r.Error = err
	}
	return &r
}

func GetPillarsFromDB() *types.PillarData {
	var r types.PillarData
	_, err := DB.Exec(`DESCRIBE ` + cfg.Loader().PillarsTable)
	switch err {
	case nil:
		//get pillars
		rows, err := DB.Query(`SELECT pillar,fortitle FROM ` + cfg.Loader().PillarsTable)
		defer rows.Close()
		if err == nil {
			for rows.Next() {
				var pillar string
				var fortitle string
				var p types.Pillar

				err := rows.Scan(&pillar, &fortitle)

				p.Pillar = pillar
				p.Fortitle = fortitle

				if err == nil {
					r.Pillars = append(r.Pillars, p)
				} else {
					r.Error = err
				}
			}
		} else {
			r.Error = err
		}
	default:
		r.Error = err
	}
	return &r
}

func GetHousesFromDB() *types.HouseData {
	var r types.HouseData
	_, err := DB.Exec(`DESCRIBE ` + cfg.Loader().HouseLevelTable)
	switch err {
	case nil:
		//get houses
		rows, err := DB.Query(`SELECT DISTINCT(housename) FROM ` + cfg.Loader().HouseLevelTable)
		defer rows.Close()
		if err == nil {
			for rows.Next() {
				var house string
				err = rows.Scan(&house)
				if err == nil {
					r.House = append(r.House, house)
				} else {
					r.Error = err
				}
			}
		} else {
			r.Error = err
		}
	default:
		r.Error = err
	}
	return &r
}

func GetSubGovsFromDB() *struct {
	Designation string
	Data        types.HouseData
} {
	var r struct {
		Designation string
		Data        types.HouseData
	}
	_, err := DB.Exec(`DESCRIBE ` + cfg.Loader().SubGovLevelTable)
	switch err {
	case nil:
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			//get designation
			defer wg.Done()
			rows, err := DB.Query(`SELECT DISTINCT(slotdesignation) FROM ` + cfg.Loader().SubGovLevelTable)
			defer rows.Close()
			if err == nil {
				for rows.Next() {
					var d string
					err = rows.Scan(&d)
					if err == nil {
						r.Designation = d
					} else {
						r.Data.Error = err
					}

				}
			}
		}()

		go func() {
			defer wg.Done()
			//get slotnames
			rows, err := DB.Query(`SELECT DISTINCT(slotname) FROM ` + cfg.Loader().SubGovLevelTable)
			defer rows.Close()
			if err == nil {
				for rows.Next() {
					var house string
					err = rows.Scan(&house)
					if err == nil {
						r.Data.House = append(r.Data.House, house)
					} else {
						r.Data.Error = err
					}
				}
			} else {
				r.Data.Error = err
			}
		}()

		wg.Wait()
	default:
		r.Data.Error = err
	}
	return &r
}

func GetPillarsFor(level interface{}) *[]string {
	var r []string
	all := "All"
	switch dd := level.(type) {
	case string:
		switch dd {
		case "president":
			_, err := DB.Exec(`DESCRIBE ` + cfg.Loader().PillarsTable)
			switch err {
			case nil:
				//get pillars
				rows, err := DB.Query(`SELECT pillar FROM `+cfg.Loader().PillarsTable+` WHERE fortitle=? OR fortitle=?`, level, all)
				defer rows.Close()
				if err == nil {
					for rows.Next() {
						var pillar string
						err := rows.Scan(&pillar)
						if err == nil {
							r = append(r, pillar)
						}
					}
				}
			}
		case "deputy president":
			_, err := DB.Exec(`DESCRIBE ` + cfg.Loader().PillarsTable)
			switch err {
			case nil:
				//get pillars
				rows, err := DB.Query(`SELECT pillar FROM `+cfg.Loader().PillarsTable+` WHERE fortitle=? OR fortitle=?`, level, all)
				defer rows.Close()
				if err == nil {
					for rows.Next() {
						var pillar string
						err := rows.Scan(&pillar)
						if err == nil {
							r = append(r, pillar)
						}
					}
				}
			}
		}
	case struct {
		Designation string
		Type        string
		Data        struct {
			SlotName string
		}
	}:
		switch dd.Type {
		case "houseslot":
			var title string
			row, err := DB.Query(`SELECT DISTINCT(title) FROM `+cfg.Loader().HouseLevelTable+` WHERE slotdesignation=? AND slotname=?`, dd.Designation, dd.Data.SlotName)
			defer row.Close()
			if err == nil {
				for row.Next() {
					_ = row.Scan(&title)
				}
			}
			//configure for houselevel
			_, errT := DB.Exec(`DESCRIBE ` + cfg.Loader().PillarsTable)
			switch errT {
			case nil:
				//get pillars
				rows, err := DB.Query(`SELECT pillar FROM `+cfg.Loader().PillarsTable+` WHERE fortitle=? OR fortitle=?`, title, all)
				defer rows.Close()
				if err == nil {
					for rows.Next() {
						var pillar string
						err := rows.Scan(&pillar)
						if err == nil {
							r = append(r, pillar)
						}
					}
				}
			}
		case "rootslot":
			var title string
			log.Println(dd.Designation)
			log.Println(dd.Data.SlotName)
			row, err := DB.Query(`SELECT DISTINCT(title) FROM `+cfg.Loader().GrassRootLevelTable+` WHERE slotdesignation=? AND slotname=?`, dd.Designation, dd.Data.SlotName)
			defer row.Close()
			if err == nil {
				for row.Next() {
					_ = row.Scan(&title)
				}
			}
			//configure for rootlevel
			_, errT := DB.Exec(`DESCRIBE ` + cfg.Loader().PillarsTable)
			switch errT {
			case nil:
				//get pillars
				rows, err := DB.Query(`SELECT pillar FROM `+cfg.Loader().PillarsTable+` WHERE fortitle=? OR fortitle=?`, title, all)
				defer rows.Close()
				if err == nil {
					for rows.Next() {
						var pillar string
						err := rows.Scan(&pillar)
						if err == nil {
							r = append(r, pillar)
							//log.Println(pillar)
						}
					}
				}
			}
		}
	case struct{ GovName string }:
		//configure for subgovlevel
		var title string
		row, err := DB.Query(`SELECT DISTINCT(position) FROM `+cfg.Loader().SubGovLevelTable+` WHERE slotname=?`, dd.GovName)
		defer row.Close()
		if err == nil {
			for row.Next() {
				_ = row.Scan(&title)
			}
		}
		//get pillars
		rows, err := DB.Query(`SELECT pillar FROM `+cfg.Loader().PillarsTable+` WHERE fortitle=? OR fortitle=?`, title, all)
		defer rows.Close()
		if err == nil {
			for rows.Next() {
				var pillar string
				err := rows.Scan(&pillar)
				if err == nil {
					r = append(r, pillar)
				}
			}
		}

	}
	return &r
}

func NewPillar(d string, e string) *bool {
	var r bool
	table := cfg.Loader().PillarsTable
	_, errDB := DB.Exec(`CREATE TABLE IF NOT EXISTS ` + table + `(
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		pillar VARCHAR(255) NOT NULL,
		fortitle VARCHAR(255) NOT NULL,
		createdAt VARCHAR(255) NOT NULL,
		PRIMARY KEY (id),
		UNIQUE KEY(pillar)
	)`)
	report.ErrLogger(errDB)
	// insert
	stmt, errI := DB.Prepare(`INSERT IGNORE INTO ` + table + ` (pillar,fortitle,createdAt) values(?,?,?)`)
	report.ErrLogger(errI)
	res, _ := stmt.Exec(d, e, time.Now())
	_, errL := res.LastInsertId()
	if errL == nil {
		r = true
	} else {
		r = false
	}
	return &r
}

func RemovePillar(d string) *bool {
	var r bool
	table := cfg.Loader().PillarsTable
	stmt, err := DB.Prepare(`DELETE FROM ` + table + ` WHERE pillar=? `)
	report.ErrLogger(err)
	res, _ := stmt.Exec(d)
	_, errL := res.LastInsertId()
	if errL == nil {
		r = true
	} else {
		r = false
	}
	return &r

}

func GetRepSlots() *map[string][]string {
	table := cfg.Loader().SlotsTable
	var d []string
	Slots := make(map[string][]string)
	rows, err := DB.Query(`SELECT DISTINCT designation FROM ` + table)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var slot types.Slot
			err := rows.Scan(&slot.Designation)
			if err == nil {
				d = append(d, slot.Designation)
			}
		}
	}
	for i := range d {
		key := d[i]
		var slotsOfKey []string
		//get the slotnames for the key
		rows, err := DB.Query(`SELECT DISTINCT slotname FROM slotstable WHERE designation=?`, key)
		defer rows.Close()
		if err == nil {
			for rows.Next() {
				var slotname string
				err := rows.Scan(&slotname)
				if err == nil {
					slotsOfKey = append(slotsOfKey, slotname)
				}
			}
		}
		Slots[key] = slotsOfKey
	}
	return &Slots
}

func GetRepSlotDesignationForSubGov(s string) *string {
	var t string
	row, err := DB.Query(`SELECT DISTINCT(slotdesignation) FROM `+cfg.Loader().SubGovLevelTable+` WHERE slotname=?`, s)
	defer row.Close()
	if err == nil {
		for row.Next() {
			_ = row.Scan(&t)
		}
	}
	return &t
}

func GetAPIofForLevelAndImage(d interface{}) *map[string]string {
	var theAPI string
	var theImage string
	m := make(map[string]string)
	switch dd := d.(type) {
	case string:
		table := cfg.Loader().TopLevelTable
		rows, err := DB.Query(`SELECT api,imageurl FROM `+table+` WHERE position=?`, dd)
		defer rows.Close()
		if err == nil {
			for rows.Next() {
				err := rows.Scan(&theAPI, &theImage)
				report.ErrLogger(err)
			}
		}
		m["api"] = theAPI
		m["image"] = theImage
	case struct {
		Designation string
		Type        string
		Data        struct {
			SlotName string
		}
	}:
		switch dd.Type {
		case "houseslot":
			table := cfg.Loader().HouseLevelTable
			rows, err := DB.Query(`SELECT api,imageurl FROM `+table+` WHERE slotdesignation =? AND slotname=?`, dd.Designation, dd.Data.SlotName)
			defer rows.Close()
			if err == nil {
				for rows.Next() {
					err := rows.Scan(&theAPI, &theImage)
					report.ErrLogger(err)
				}
			}
			m["api"] = theAPI
			m["image"] = theImage
		case "rootslot":
			table := cfg.Loader().GrassRootLevelTable
			rows, err := DB.Query(`SELECT api,imageurl FROM `+table+` WHERE slotdesignation =? AND slotname=?`, dd.Designation, dd.Data.SlotName)
			defer rows.Close()
			if err == nil {
				for rows.Next() {
					err := rows.Scan(&theAPI, &theImage)
					report.ErrLogger(err)
				}
			}
			m["api"] = theAPI
			m["image"] = theImage
		}

	case struct{ GovName string }:
		table := cfg.Loader().SubGovLevelTable
		rows, err := DB.Query(`SELECT api,imageurl FROM `+table+` WHERE slotname=?`, dd.GovName)
		defer rows.Close()
		if err == nil {
			for rows.Next() {
				err := rows.Scan(&theAPI, &theImage)
				report.ErrLogger(err)
			}
		}
		m["api"] = theAPI
		m["image"] = theImage
	}

	return &m
}

func GetRepSlotsForHouse(house string) *[]string {
	var rAll []string
	table := cfg.Loader().HouseLevelTable
	rows, err := DB.Query(`SELECT DISTINCT slotname FROM `+table+` WHERE housename=?`, house)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var s string
			err := rows.Scan(&s)
			report.ErrLogger(err)
			rAll = append(rAll, s)
		}
	}
	return &rAll
}

func GetRepSlotsForHouseDesignation(house string) *string {
	var rAll string
	table := cfg.Loader().HouseLevelTable
	rows, err := DB.Query(`SELECT DISTINCT slotdesignation FROM `+table+` WHERE housename=?`, house)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			err := rows.Scan(&rAll)
			report.ErrLogger(err)
		}
	}
	return &rAll

}

func GetRepSlotsForRoot(gov string) *[]string {
	var rAll []string
	table := cfg.Loader().GrassRootLevelTable
	rows, err := DB.Query(`SELECT DISTINCT slotname FROM `+table+` WHERE legof=?`, gov)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var s string
			err := rows.Scan(&s)
			report.ErrLogger(err)
			rAll = append(rAll, s)
		}
	}
	return &rAll

}

func GetRepSlotsForRootDesignation(gov string) *string {
	var rAll string
	table := cfg.Loader().GrassRootLevelTable
	rows, err := DB.Query(`SELECT DISTINCT slotdesignation FROM `+table+` WHERE legof=?`, gov)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			err := rows.Scan(&rAll)
			report.ErrLogger(err)
		}
	}
	return &rAll
}
