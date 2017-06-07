// Code generated by. DO NOT EDIT.
/* This file was generated by Gosora's Query Generator. The thing above is to tell GH this file is generated. */
// +build !pgsql !sqlite !mssql
package main

import "log"
import "database/sql"

var get_user_stmt *sql.Stmt
var get_full_user_stmt *sql.Stmt
var get_topic_stmt *sql.Stmt
var get_reply_stmt *sql.Stmt
var login_stmt *sql.Stmt
var get_password_stmt *sql.Stmt
var username_exists_stmt *sql.Stmt
var get_settings_stmt *sql.Stmt
var get_setting_stmt *sql.Stmt
var get_full_setting_stmt *sql.Stmt
var is_plugin_active_stmt *sql.Stmt
var get_users_stmt *sql.Stmt
var is_theme_default_stmt *sql.Stmt
var get_modlogs_stmt *sql.Stmt
var get_reply_tid_stmt *sql.Stmt
var get_topic_fid_stmt *sql.Stmt
var get_user_reply_uid_stmt *sql.Stmt
var has_liked_topic_stmt *sql.Stmt
var has_liked_reply_stmt *sql.Stmt
var get_user_name_stmt *sql.Stmt
var get_emails_by_user_stmt *sql.Stmt
var get_topic_basic_stmt *sql.Stmt
var get_topic_list_stmt *sql.Stmt
var get_topic_user_stmt *sql.Stmt
var get_topic_by_reply_stmt *sql.Stmt
var get_topic_replies_stmt *sql.Stmt
var get_forum_topics_stmt *sql.Stmt
var get_profile_replies_stmt *sql.Stmt
var create_topic_stmt *sql.Stmt
var create_report_stmt *sql.Stmt
var create_reply_stmt *sql.Stmt
var create_action_reply_stmt *sql.Stmt
var create_like_stmt *sql.Stmt
var add_activity_stmt *sql.Stmt
var register_stmt *sql.Stmt
var add_replies_to_topic_stmt *sql.Stmt
var remove_replies_from_topic_stmt *sql.Stmt
var add_topics_to_forum_stmt *sql.Stmt
var remove_topics_from_forum_stmt *sql.Stmt
var update_forum_cache_stmt *sql.Stmt
var add_likes_to_topic_stmt *sql.Stmt
var add_likes_to_reply_stmt *sql.Stmt
var edit_topic_stmt *sql.Stmt
var edit_reply_stmt *sql.Stmt
var stick_topic_stmt *sql.Stmt
var unstick_topic_stmt *sql.Stmt
var update_last_ip_stmt *sql.Stmt
var update_session_stmt *sql.Stmt
var logout_stmt *sql.Stmt
var set_password_stmt *sql.Stmt
var set_avatar_stmt *sql.Stmt
var set_username_stmt *sql.Stmt
var change_group_stmt *sql.Stmt
var activate_user_stmt *sql.Stmt

func gen_mysql() (err error) {
	if debug {
		log.Print("Building the generated statements")
	}
	
	log.Print("Preparing get_user statement.")
	get_user_stmt, err = db.Prepare("SELECT `name`,`group`,`is_super_admin`,`avatar`,`message`,`url_prefix`,`url_name`,`level` FROM `users` WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_full_user statement.")
	get_full_user_stmt, err = db.Prepare("SELECT `name`,`group`,`is_super_admin`,`session`,`email`,`avatar`,`message`,`url_prefix`,`url_name`,`level`,`score`,`last_ip` FROM `users` WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_topic statement.")
	get_topic_stmt, err = db.Prepare("SELECT `title`,`content`,`createdBy`,`createdAt`,`is_closed`,`sticky`,`parentID`,`ipaddress`,`postCount`,`likeCount` FROM `topics` WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_reply statement.")
	get_reply_stmt, err = db.Prepare("SELECT `content`,`createdBy`,`createdAt`,`lastEdit`,`lastEditBy`,`ipaddress`,`likeCount` FROM `replies` WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing login statement.")
	login_stmt, err = db.Prepare("SELECT `uid`,`name`,`password`,`salt` FROM `users` WHERE `name` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_password statement.")
	get_password_stmt, err = db.Prepare("SELECT `password`,`salt` FROM `users` WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing username_exists statement.")
	username_exists_stmt, err = db.Prepare("SELECT `name` FROM `users` WHERE `name` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_settings statement.")
	get_settings_stmt, err = db.Prepare("SELECT `name`,`content`,`type` FROM `settings`")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_setting statement.")
	get_setting_stmt, err = db.Prepare("SELECT `content`,`type` FROM `settings` WHERE `name` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_full_setting statement.")
	get_full_setting_stmt, err = db.Prepare("SELECT `name`,`type`,`constraints` FROM `settings` WHERE `name` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing is_plugin_active statement.")
	is_plugin_active_stmt, err = db.Prepare("SELECT `active` FROM `plugins` WHERE `uname` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_users statement.")
	get_users_stmt, err = db.Prepare("SELECT `uid`,`name`,`group`,`active`,`is_super_admin`,`avatar` FROM `users`")
	if err != nil {
		return err
	}
		
	log.Print("Preparing is_theme_default statement.")
	is_theme_default_stmt, err = db.Prepare("SELECT `default` FROM `themes` WHERE `uname` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_modlogs statement.")
	get_modlogs_stmt, err = db.Prepare("SELECT `action`,`elementID`,`elementType`,`ipaddress`,`actorID`,`doneAt` FROM `moderation_logs`")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_reply_tid statement.")
	get_reply_tid_stmt, err = db.Prepare("SELECT `tid` FROM `replies` WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_topic_fid statement.")
	get_topic_fid_stmt, err = db.Prepare("SELECT `parentID` FROM `topics` WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_user_reply_uid statement.")
	get_user_reply_uid_stmt, err = db.Prepare("SELECT `uid` FROM `users_replies` WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing has_liked_topic statement.")
	has_liked_topic_stmt, err = db.Prepare("SELECT `targetItem` FROM `likes` WHERE `sentBy` = ? AND  `targetItem` = ? AND  `targetType` = 'topics'")
	if err != nil {
		return err
	}
		
	log.Print("Preparing has_liked_reply statement.")
	has_liked_reply_stmt, err = db.Prepare("SELECT `targetItem` FROM `likes` WHERE `sentBy` = ? AND  `targetItem` = ? AND  `targetType` = 'replies'")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_user_name statement.")
	get_user_name_stmt, err = db.Prepare("SELECT `name` FROM `users` WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_emails_by_user statement.")
	get_emails_by_user_stmt, err = db.Prepare("SELECT `email`,`validated` FROM `emails` WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_topic_basic statement.")
	get_topic_basic_stmt, err = db.Prepare("SELECT `title`,`content` FROM `topics` WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_topic_list statement.")
	get_topic_list_stmt, err = db.Prepare("SELECT `topics`.`tid`,`topics`.`title`,`topics`.`content`,`topics`.`createdBy`,`topics`.`is_closed`,`topics`.`sticky`,`topics`.`createdAt`,`topics`.`parentID`,`users`.`name`,`users`.`avatar` FROM `topics` LEFT JOIN `users` ON `topics`.`createdBy`=`users`.`uid`  ORDER BY topics.sticky DESC,topics.lastReplyAt DESC,topics.createdBy DESC")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_topic_user statement.")
	get_topic_user_stmt, err = db.Prepare("SELECT `topics`.`title`,`topics`.`content`,`topics`.`createdBy`,`topics`.`createdAt`,`topics`.`is_closed`,`topics`.`sticky`,`topics`.`parentID`,`topics`.`ipaddress`,`topics`.`postCount`,`topics`.`likeCount`,`users`.`name`,`users`.`avatar`,`users`.`group`,`users`.`url_prefix`,`users`.`url_name`,`users`.`level` FROM `topics` LEFT JOIN `users` ON `topics`.`createdBy`=`users`.`uid`  WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_topic_by_reply statement.")
	get_topic_by_reply_stmt, err = db.Prepare("SELECT `topics`.`tid`,`topics`.`title`,`topics`.`content`,`topics`.`createdBy`,`topics`.`createdAt`,`topics`.`is_closed`,`topics`.`sticky`,`topics`.`parentID`,`topics`.`ipaddress`,`topics`.`postCount`,`topics`.`likeCount` FROM `replies` LEFT JOIN `topics` ON `replies`.`tid`=`topics`.`tid`  WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_topic_replies statement.")
	get_topic_replies_stmt, err = db.Prepare("SELECT `replies`.`rid`,`replies`.`content`,`replies`.`createdBy`,`replies`.`createdAt`,`replies`.`lastEdit`,`replies`.`lastEditBy`,`users`.`avatar`,`users`.`name`,`users`.`group`,`users`.`url_prefix`,`users`.`url_name`,`users`.`level`,`replies`.`ipaddress` FROM `replies` LEFT JOIN `users` ON `replies`.`createdBy`=`users`.`uid`  WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_forum_topics statement.")
	get_forum_topics_stmt, err = db.Prepare("SELECT `topics`.`tid`,`topics`.`title`,`topics`.`content`,`topics`.`createdBy`,`topics`.`is_closed`,`topics`.`sticky`,`topics`.`createdAt`,`topics`.`lastReplyAt`,`topics`.`parentID`,`users`.`name`,`users`.`avatar` FROM `topics` LEFT JOIN `users` ON `topics`.`createdBy`=`users`.`uid`  WHERE `topics`.`parentID` = ?  ORDER BY topics.sticky DESC,topics.lastReplyAt DESC,topics.createdBy DESC")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_profile_replies statement.")
	get_profile_replies_stmt, err = db.Prepare("SELECT `users_replies`.`rid`,`users_replies`.`content`,`users_replies`.`createdBy`,`users_replies`.`createdAt`,`users_replies`.`lastEdit`,`users_replies`.`lastEditBy`,`users`.`avatar`,`users`.`name`,`users`.`group` FROM `users_replies` LEFT JOIN `users` ON `users_replies`.`createdBy`=`users`.`uid`  WHERE `users_replies`.`uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing create_topic statement.")
	create_topic_stmt, err = db.Prepare("INSERT INTO `topics`(`parentID`,`title`,`content`,`parsed_content`,`createdAt`,`lastReplyAt`,`ipaddress`,`words`,`createdBy`) VALUES (?,?,?,?,NOW(),NOW(),?,?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing create_report statement.")
	create_report_stmt, err = db.Prepare("INSERT INTO `topics`(`title`,`content`,`parsed_content`,`createdAt`,`lastReplyAt`,`createdBy`,`data`,`parentID`,`css_class`) VALUES (?,?,?,NOW(),NOW(),?,?,1,'report')")
	if err != nil {
		return err
	}
		
	log.Print("Preparing create_reply statement.")
	create_reply_stmt, err = db.Prepare("INSERT INTO `replies`(`tid`,`content`,`parsed_content`,`createdAt`,`ipaddress`,`words`,`createdBy`) VALUES (?,?,?,NOW(),?,?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing create_action_reply statement.")
	create_action_reply_stmt, err = db.Prepare("INSERT INTO `replies`(`tid`,`actionType`,`ipaddress`,`createdBy`) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing create_like statement.")
	create_like_stmt, err = db.Prepare("INSERT INTO `likes`(`weight`,`targetItem`,`targetType`,`sentBy`) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_activity statement.")
	add_activity_stmt, err = db.Prepare("INSERT INTO `activity_stream`(`actor`,`targetUser`,`event`,`elementType`,`elementID`) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing register statement.")
	register_stmt, err = db.Prepare("INSERT INTO `users`(`name`,`email`,`password`,`salt`,`group`,`is_super_admin`,`session`,`active`,`message`) VALUES (?,?,?,?,?,0,?,?,'')")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_replies_to_topic statement.")
	add_replies_to_topic_stmt, err = db.Prepare("UPDATE `topics` SET `postCount` = `postCount` + ?,`lastReplyAt` = NOW() WHERE `tid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing remove_replies_from_topic statement.")
	remove_replies_from_topic_stmt, err = db.Prepare("UPDATE `topics` SET `postCount` = `postCount` - ? WHERE `tid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_topics_to_forum statement.")
	add_topics_to_forum_stmt, err = db.Prepare("UPDATE `forums` SET `topicCount` = `topicCount` + ? WHERE `fid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing remove_topics_from_forum statement.")
	remove_topics_from_forum_stmt, err = db.Prepare("UPDATE `forums` SET `topicCount` = `topicCount` - ? WHERE `fid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_forum_cache statement.")
	update_forum_cache_stmt, err = db.Prepare("UPDATE `forums` SET `lastTopic` = ?,`lastTopicID` = ?,`lastReplyer` = ?,`lastReplyerID` = ?,`lastTopicTime` = NOW() WHERE `fid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_likes_to_topic statement.")
	add_likes_to_topic_stmt, err = db.Prepare("UPDATE `topics` SET `likeCount` = `likeCount` + ? WHERE `tid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_likes_to_reply statement.")
	add_likes_to_reply_stmt, err = db.Prepare("UPDATE `replies` SET `likeCount` = `likeCount` + ? WHERE `rid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing edit_topic statement.")
	edit_topic_stmt, err = db.Prepare("UPDATE `topics` SET `title` = ?,`content` = ?,`parsed_content` = ?,`is_closed` = ? WHERE `tid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing edit_reply statement.")
	edit_reply_stmt, err = db.Prepare("UPDATE `replies` SET `content` = ?,`parsed_content` = ? WHERE `rid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing stick_topic statement.")
	stick_topic_stmt, err = db.Prepare("UPDATE `topics` SET `sticky` = 1 WHERE `tid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing unstick_topic statement.")
	unstick_topic_stmt, err = db.Prepare("UPDATE `topics` SET `sticky` = 0 WHERE `tid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_last_ip statement.")
	update_last_ip_stmt, err = db.Prepare("UPDATE `users` SET `last_ip` = ? WHERE `uid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_session statement.")
	update_session_stmt, err = db.Prepare("UPDATE `users` SET `session` = ? WHERE `uid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing logout statement.")
	logout_stmt, err = db.Prepare("UPDATE `users` SET `session` = '' WHERE `uid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing set_password statement.")
	set_password_stmt, err = db.Prepare("UPDATE `users` SET `password` = ?,`salt` = ? WHERE `uid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing set_avatar statement.")
	set_avatar_stmt, err = db.Prepare("UPDATE `users` SET `avatar` = ? WHERE `uid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing set_username statement.")
	set_username_stmt, err = db.Prepare("UPDATE `users` SET `name` = ? WHERE `uid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing change_group statement.")
	change_group_stmt, err = db.Prepare("UPDATE `users` SET `group` = ? WHERE `uid` = ? ")
	if err != nil {
		return err
	}
		
	log.Print("Preparing activate_user statement.")
	activate_user_stmt, err = db.Prepare("UPDATE `users` SET `active` = 1 WHERE `uid` = ? ")
	if err != nil {
		return err
	}
	
	return nil
}