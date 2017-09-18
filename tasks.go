/*
*
*	Gosora Task System
*	Copyright Azareal 2017 - 2018
*
 */
package main

import "time"

var lastSync time.Time

func init() {
	lastSync = time.Now()
}

func handleExpiredScheduledGroups() error {
	rows, err := getExpiredScheduledGroupsStmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	var uid int
	for rows.Next() {
		err := rows.Scan(&uid)
		if err != nil {
			return err
		}
		_, err = replaceScheduleGroupStmt.Exec(uid, 0, 0, time.Now(), false)
		if err != nil {
			return err
		}
		_, err = setTempGroupStmt.Exec(0, uid)
		if err != nil {
			return err
		}
		_ = users.Reload(uid)
	}
	return rows.Err()
}

func handleServerSync() error {
	var lastUpdate time.Time
	var lastUpdateStr string
	err := getSyncStmt.QueryRow().Scan(&lastUpdateStr)
	if err != nil {
		return err
	}

	lastUpdate, err = time.Parse("2006-01-02 15:04:05", lastUpdateStr)
	if err != nil {
		return err
	}

	if lastUpdate.After(lastSync) {
		// TODO: A more granular sync
		err = fstore.LoadForums()
		if err != nil {
			return err
		}
		// TODO: Resync the groups
		// TODO: Resync the permissions
		err = LoadSettings()
		if err != nil {
			return err
		}
		err = LoadWordFilters()
		if err != nil {
			return err
		}
	}
	return nil
}
