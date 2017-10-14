// +build !pgsql, !sqlite, !mssql

/*
*
*	Gosora MySQL Interface
*	Copyright Azareal 2016 - 2018
*
 */
package main

import "log"

//import "time"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "./query_gen/lib"

var dbCollation = "utf8mb4_general_ci"
var getActivityFeedByWatcherStmt *sql.Stmt
var getActivityCountByWatcherStmt *sql.Stmt
var todaysPostCountStmt *sql.Stmt
var todaysTopicCountStmt *sql.Stmt
var todaysReportCountStmt *sql.Stmt
var todaysNewUserCountStmt *sql.Stmt
var findUsersByIPUsersStmt *sql.Stmt
var findUsersByIPTopicsStmt *sql.Stmt
var findUsersByIPRepliesStmt *sql.Stmt

func init() {
	dbAdapter = "mysql"
	_initDatabase = initMySQL
}

func initMySQL() (err error) {
	var _dbpassword string
	if dbConfig.Password != "" {
		_dbpassword = ":" + dbConfig.Password
	}

	// TODO: Move this bit to the query gen lib
	// Open the database connection
	db, err = sql.Open("mysql", dbConfig.Username+_dbpassword+"@tcp("+dbConfig.Host+":"+dbConfig.Port+")/"+dbConfig.Dbname+"?collation="+dbCollation+"&parseTime=true")
	if err != nil {
		return err
	}

	// Make sure that the connection is alive
	err = db.Ping()
	if err != nil {
		return err
	}

	// Fetch the database version
	db.QueryRow("SELECT VERSION()").Scan(&dbVersion)

	// Set the number of max open connections
	db.SetMaxOpenConns(64)
	db.SetMaxIdleConns(32)

	// Only hold connections open for five seconds to avoid accumulating a large number of stale connections
	//db.SetConnMaxLifetime(5 * time.Second)

	// Build the generated prepared statements, we are going to slowly move the queries over to the query generator rather than writing them all by hand, this'll make it easier for us to implement database adapters for other databases like PostgreSQL, MSSQL, SQlite, etc.
	err = _gen_mysql()
	if err != nil {
		return err
	}

	// Ready the query builder
	qgen.Builder.SetConn(db)
	err = qgen.Builder.SetAdapter("mysql")
	if err != nil {
		return err
	}

	// TODO: Is there a less noisy way of doing this for tests?
	log.Print("Preparing get_activity_feed_by_watcher statement.")
	getActivityFeedByWatcherStmt, err = db.Prepare("SELECT activity_stream_matches.asid, activity_stream.actor, activity_stream.targetUser, activity_stream.event, activity_stream.elementType, activity_stream.elementID FROM `activity_stream_matches` INNER JOIN `activity_stream` ON activity_stream_matches.asid = activity_stream.asid AND activity_stream_matches.watcher != activity_stream.actor WHERE `watcher` = ? ORDER BY activity_stream.asid ASC LIMIT 8")
	if err != nil {
		return err
	}

	log.Print("Preparing get_activity_count_by_watcher statement.")
	getActivityCountByWatcherStmt, err = db.Prepare("SELECT count(*) FROM `activity_stream_matches` INNER JOIN `activity_stream` ON activity_stream_matches.asid = activity_stream.asid AND activity_stream_matches.watcher != activity_stream.actor WHERE `watcher` = ?")
	if err != nil {
		return err
	}

	log.Print("Preparing todays_post_count statement.")
	todaysPostCountStmt, err = db.Prepare("select count(*) from replies where createdAt BETWEEN (utc_timestamp() - interval 1 day) and utc_timestamp()")
	if err != nil {
		return err
	}

	log.Print("Preparing todays_topic_count statement.")
	todaysTopicCountStmt, err = db.Prepare("select count(*) from topics where createdAt BETWEEN (utc_timestamp() - interval 1 day) and utc_timestamp()")
	if err != nil {
		return err
	}

	log.Print("Preparing todays_report_count statement.")
	todaysReportCountStmt, err = db.Prepare("select count(*) from topics where createdAt BETWEEN (utc_timestamp() - interval 1 day) and utc_timestamp() and parentID = 1")
	if err != nil {
		return err
	}

	log.Print("Preparing todays_newuser_count statement.")
	todaysNewUserCountStmt, err = db.Prepare("select count(*) from users where createdAt BETWEEN (utc_timestamp() - interval 1 day) and utc_timestamp()")
	if err != nil {
		return err
	}

	log.Print("Preparing find_users_by_ip_users statement.")
	findUsersByIPUsersStmt, err = db.Prepare("select uid from users where last_ip = ?")
	if err != nil {
		return err
	}

	log.Print("Preparing find_users_by_ip_topics statement.")
	findUsersByIPTopicsStmt, err = db.Prepare("select uid from users where uid in(select createdBy from topics where ipaddress = ?)")
	if err != nil {
		return err
	}

	log.Print("Preparing find_users_by_ip_replies statement.")
	findUsersByIPRepliesStmt, err = db.Prepare("select uid from users where uid in(select createdBy from replies where ipaddress = ?)")
	return err
}
