// +build mssql

// This file was generated by Gosora's Query Generator. Please try to avoid modifying this file, as it might change at any time.
package main

import "log"
import "database/sql"

// nolint
var getUserStmt *sql.Stmt
var getReplyStmt *sql.Stmt
var getUserReplyStmt *sql.Stmt
var getPasswordStmt *sql.Stmt
var getSettingsStmt *sql.Stmt
var getSettingStmt *sql.Stmt
var getFullSettingStmt *sql.Stmt
var getFullSettingsStmt *sql.Stmt
var getGroupsStmt *sql.Stmt
var getForumsStmt *sql.Stmt
var getForumsPermissionsStmt *sql.Stmt
var getPluginsStmt *sql.Stmt
var getThemesStmt *sql.Stmt
var getWidgetsStmt *sql.Stmt
var isPluginActiveStmt *sql.Stmt
var getUsersStmt *sql.Stmt
var getUsersOffsetStmt *sql.Stmt
var getWordFiltersStmt *sql.Stmt
var isThemeDefaultStmt *sql.Stmt
var getModlogsStmt *sql.Stmt
var getModlogsOffsetStmt *sql.Stmt
var getReplyTIDStmt *sql.Stmt
var getTopicFIDStmt *sql.Stmt
var getUserReplyUIDStmt *sql.Stmt
var hasLikedTopicStmt *sql.Stmt
var hasLikedReplyStmt *sql.Stmt
var getUserNameStmt *sql.Stmt
var getEmailsByUserStmt *sql.Stmt
var getTopicBasicStmt *sql.Stmt
var getActivityEntryStmt *sql.Stmt
var forumEntryExistsStmt *sql.Stmt
var groupEntryExistsStmt *sql.Stmt
var getForumTopicsOffsetStmt *sql.Stmt
var getExpiredScheduledGroupsStmt *sql.Stmt
var getSyncStmt *sql.Stmt
var getAttachmentStmt *sql.Stmt
var getTopicRepliesOffsetStmt *sql.Stmt
var getTopicListStmt *sql.Stmt
var getTopicUserStmt *sql.Stmt
var getTopicByReplyStmt *sql.Stmt
var getTopicRepliesStmt *sql.Stmt
var getForumTopicsStmt *sql.Stmt
var getProfileRepliesStmt *sql.Stmt
var getWatchersStmt *sql.Stmt
var createTopicStmt *sql.Stmt
var createReportStmt *sql.Stmt
var createReplyStmt *sql.Stmt
var createActionReplyStmt *sql.Stmt
var createLikeStmt *sql.Stmt
var addActivityStmt *sql.Stmt
var notifyOneStmt *sql.Stmt
var addEmailStmt *sql.Stmt
var createProfileReplyStmt *sql.Stmt
var addSubscriptionStmt *sql.Stmt
var createForumStmt *sql.Stmt
var addForumPermsToForumStmt *sql.Stmt
var addPluginStmt *sql.Stmt
var addThemeStmt *sql.Stmt
var createGroupStmt *sql.Stmt
var addModlogEntryStmt *sql.Stmt
var addAdminlogEntryStmt *sql.Stmt
var addAttachmentStmt *sql.Stmt
var createWordFilterStmt *sql.Stmt
var addForumPermsToGroupStmt *sql.Stmt
var replaceScheduleGroupStmt *sql.Stmt
var addRepliesToTopicStmt *sql.Stmt
var removeRepliesFromTopicStmt *sql.Stmt
var addTopicsToForumStmt *sql.Stmt
var removeTopicsFromForumStmt *sql.Stmt
var updateForumCacheStmt *sql.Stmt
var addLikesToTopicStmt *sql.Stmt
var addLikesToReplyStmt *sql.Stmt
var editTopicStmt *sql.Stmt
var editReplyStmt *sql.Stmt
var stickTopicStmt *sql.Stmt
var unstickTopicStmt *sql.Stmt
var lockTopicStmt *sql.Stmt
var unlockTopicStmt *sql.Stmt
var updateLastIPStmt *sql.Stmt
var updateSessionStmt *sql.Stmt
var setPasswordStmt *sql.Stmt
var setAvatarStmt *sql.Stmt
var setUsernameStmt *sql.Stmt
var changeGroupStmt *sql.Stmt
var activateUserStmt *sql.Stmt
var updateUserLevelStmt *sql.Stmt
var incrementUserScoreStmt *sql.Stmt
var incrementUserPostsStmt *sql.Stmt
var incrementUserBigpostsStmt *sql.Stmt
var incrementUserMegapostsStmt *sql.Stmt
var incrementUserTopicsStmt *sql.Stmt
var editProfileReplyStmt *sql.Stmt
var updateForumStmt *sql.Stmt
var updateSettingStmt *sql.Stmt
var updatePluginStmt *sql.Stmt
var updatePluginInstallStmt *sql.Stmt
var updateThemeStmt *sql.Stmt
var updateUserStmt *sql.Stmt
var updateGroupPermsStmt *sql.Stmt
var updateGroupRankStmt *sql.Stmt
var updateGroupStmt *sql.Stmt
var updateEmailStmt *sql.Stmt
var verifyEmailStmt *sql.Stmt
var setTempGroupStmt *sql.Stmt
var updateWordFilterStmt *sql.Stmt
var bumpSyncStmt *sql.Stmt
var deleteReplyStmt *sql.Stmt
var deleteProfileReplyStmt *sql.Stmt
var deleteForumPermsByForumStmt *sql.Stmt
var deleteActivityStreamMatchStmt *sql.Stmt
var deleteWordFilterStmt *sql.Stmt
var reportExistsStmt *sql.Stmt
var groupCountStmt *sql.Stmt
var modlogCountStmt *sql.Stmt
var addForumPermsToForumAdminsStmt *sql.Stmt
var addForumPermsToForumStaffStmt *sql.Stmt
var addForumPermsToForumMembersStmt *sql.Stmt
var notifyWatchersStmt *sql.Stmt

// nolint
func _gen_mssql() (err error) {
	if dev.DebugMode {
		log.Print("Building the generated statements")
	}
	
	log.Print("Preparing getUser statement.")
	getUserStmt, err = db.Prepare("SELECT [name],[group],[is_super_admin],[avatar],[message],[url_prefix],[url_name],[level] FROM [users] WHERE [uid] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [name],[group],[is_super_admin],[avatar],[message],[url_prefix],[url_name],[level] FROM [users] WHERE [uid] = ?1")
		return err
	}
		
	log.Print("Preparing getReply statement.")
	getReplyStmt, err = db.Prepare("SELECT [tid],[content],[createdBy],[createdAt],[lastEdit],[lastEditBy],[ipaddress],[likeCount] FROM [replies] WHERE [rid] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [tid],[content],[createdBy],[createdAt],[lastEdit],[lastEditBy],[ipaddress],[likeCount] FROM [replies] WHERE [rid] = ?1")
		return err
	}
		
	log.Print("Preparing getUserReply statement.")
	getUserReplyStmt, err = db.Prepare("SELECT [uid],[content],[createdBy],[createdAt],[lastEdit],[lastEditBy],[ipaddress] FROM [users_replies] WHERE [rid] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [uid],[content],[createdBy],[createdAt],[lastEdit],[lastEditBy],[ipaddress] FROM [users_replies] WHERE [rid] = ?1")
		return err
	}
		
	log.Print("Preparing getPassword statement.")
	getPasswordStmt, err = db.Prepare("SELECT [password],[salt] FROM [users] WHERE [uid] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [password],[salt] FROM [users] WHERE [uid] = ?1")
		return err
	}
		
	log.Print("Preparing getSettings statement.")
	getSettingsStmt, err = db.Prepare("SELECT [name],[content],[type] FROM [settings]")
	if err != nil {
		log.Print("Bad Query: ","SELECT [name],[content],[type] FROM [settings]")
		return err
	}
		
	log.Print("Preparing getSetting statement.")
	getSettingStmt, err = db.Prepare("SELECT [content],[type] FROM [settings] WHERE [name] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [content],[type] FROM [settings] WHERE [name] = ?1")
		return err
	}
		
	log.Print("Preparing getFullSetting statement.")
	getFullSettingStmt, err = db.Prepare("SELECT [name],[type],[constraints] FROM [settings] WHERE [name] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [name],[type],[constraints] FROM [settings] WHERE [name] = ?1")
		return err
	}
		
	log.Print("Preparing getFullSettings statement.")
	getFullSettingsStmt, err = db.Prepare("SELECT [name],[content],[type],[constraints] FROM [settings]")
	if err != nil {
		log.Print("Bad Query: ","SELECT [name],[content],[type],[constraints] FROM [settings]")
		return err
	}
		
	log.Print("Preparing getGroups statement.")
	getGroupsStmt, err = db.Prepare("SELECT [gid],[name],[permissions],[plugin_perms],[is_mod],[is_admin],[is_banned],[tag] FROM [users_groups]")
	if err != nil {
		log.Print("Bad Query: ","SELECT [gid],[name],[permissions],[plugin_perms],[is_mod],[is_admin],[is_banned],[tag] FROM [users_groups]")
		return err
	}
		
	log.Print("Preparing getForums statement.")
	getForumsStmt, err = db.Prepare("SELECT [fid],[name],[desc],[active],[preset],[parentID],[parentType],[topicCount],[lastTopicID],[lastReplyerID] FROM [forums] ORDER BY fid ASC")
	if err != nil {
		log.Print("Bad Query: ","SELECT [fid],[name],[desc],[active],[preset],[parentID],[parentType],[topicCount],[lastTopicID],[lastReplyerID] FROM [forums] ORDER BY fid ASC")
		return err
	}
		
	log.Print("Preparing getForumsPermissions statement.")
	getForumsPermissionsStmt, err = db.Prepare("SELECT [gid],[fid],[permissions] FROM [forums_permissions] ORDER BY gid ASC,fid ASC")
	if err != nil {
		log.Print("Bad Query: ","SELECT [gid],[fid],[permissions] FROM [forums_permissions] ORDER BY gid ASC,fid ASC")
		return err
	}
		
	log.Print("Preparing getPlugins statement.")
	getPluginsStmt, err = db.Prepare("SELECT [uname],[active],[installed] FROM [plugins]")
	if err != nil {
		log.Print("Bad Query: ","SELECT [uname],[active],[installed] FROM [plugins]")
		return err
	}
		
	log.Print("Preparing getThemes statement.")
	getThemesStmt, err = db.Prepare("SELECT [uname],[default] FROM [themes]")
	if err != nil {
		log.Print("Bad Query: ","SELECT [uname],[default] FROM [themes]")
		return err
	}
		
	log.Print("Preparing getWidgets statement.")
	getWidgetsStmt, err = db.Prepare("SELECT [position],[side],[type],[active],[location],[data] FROM [widgets] ORDER BY position ASC")
	if err != nil {
		log.Print("Bad Query: ","SELECT [position],[side],[type],[active],[location],[data] FROM [widgets] ORDER BY position ASC")
		return err
	}
		
	log.Print("Preparing isPluginActive statement.")
	isPluginActiveStmt, err = db.Prepare("SELECT [active] FROM [plugins] WHERE [uname] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [active] FROM [plugins] WHERE [uname] = ?1")
		return err
	}
		
	log.Print("Preparing getUsers statement.")
	getUsersStmt, err = db.Prepare("SELECT [uid],[name],[group],[active],[is_super_admin],[avatar] FROM [users]")
	if err != nil {
		log.Print("Bad Query: ","SELECT [uid],[name],[group],[active],[is_super_admin],[avatar] FROM [users]")
		return err
	}
		
	log.Print("Preparing getUsersOffset statement.")
	getUsersOffsetStmt, err = db.Prepare("SELECT [uid],[name],[group],[active],[is_super_admin],[avatar] FROM [users] OFFSET ?1 ROWS FETCH NEXT ?2 ROWS ONLY")
	if err != nil {
		log.Print("Bad Query: ","SELECT [uid],[name],[group],[active],[is_super_admin],[avatar] FROM [users] OFFSET ?1 ROWS FETCH NEXT ?2 ROWS ONLY")
		return err
	}
		
	log.Print("Preparing getWordFilters statement.")
	getWordFiltersStmt, err = db.Prepare("SELECT [wfid],[find],[replacement] FROM [word_filters]")
	if err != nil {
		log.Print("Bad Query: ","SELECT [wfid],[find],[replacement] FROM [word_filters]")
		return err
	}
		
	log.Print("Preparing isThemeDefault statement.")
	isThemeDefaultStmt, err = db.Prepare("SELECT [default] FROM [themes] WHERE [uname] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [default] FROM [themes] WHERE [uname] = ?1")
		return err
	}
		
	log.Print("Preparing getModlogs statement.")
	getModlogsStmt, err = db.Prepare("SELECT [action],[elementID],[elementType],[ipaddress],[actorID],[doneAt] FROM [moderation_logs]")
	if err != nil {
		log.Print("Bad Query: ","SELECT [action],[elementID],[elementType],[ipaddress],[actorID],[doneAt] FROM [moderation_logs]")
		return err
	}
		
	log.Print("Preparing getModlogsOffset statement.")
	getModlogsOffsetStmt, err = db.Prepare("SELECT [action],[elementID],[elementType],[ipaddress],[actorID],[doneAt] FROM [moderation_logs] OFFSET ?1 ROWS FETCH NEXT ?2 ROWS ONLY")
	if err != nil {
		log.Print("Bad Query: ","SELECT [action],[elementID],[elementType],[ipaddress],[actorID],[doneAt] FROM [moderation_logs] OFFSET ?1 ROWS FETCH NEXT ?2 ROWS ONLY")
		return err
	}
		
	log.Print("Preparing getReplyTID statement.")
	getReplyTIDStmt, err = db.Prepare("SELECT [tid] FROM [replies] WHERE [rid] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [tid] FROM [replies] WHERE [rid] = ?1")
		return err
	}
		
	log.Print("Preparing getTopicFID statement.")
	getTopicFIDStmt, err = db.Prepare("SELECT [parentID] FROM [topics] WHERE [tid] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [parentID] FROM [topics] WHERE [tid] = ?1")
		return err
	}
		
	log.Print("Preparing getUserReplyUID statement.")
	getUserReplyUIDStmt, err = db.Prepare("SELECT [uid] FROM [users_replies] WHERE [rid] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [uid] FROM [users_replies] WHERE [rid] = ?1")
		return err
	}
		
	log.Print("Preparing hasLikedTopic statement.")
	hasLikedTopicStmt, err = db.Prepare("SELECT [targetItem] FROM [likes] WHERE [sentBy] = ?1 AND [targetItem] = ?2 AND [targetType] = 'topics'")
	if err != nil {
		log.Print("Bad Query: ","SELECT [targetItem] FROM [likes] WHERE [sentBy] = ?1 AND [targetItem] = ?2 AND [targetType] = 'topics'")
		return err
	}
		
	log.Print("Preparing hasLikedReply statement.")
	hasLikedReplyStmt, err = db.Prepare("SELECT [targetItem] FROM [likes] WHERE [sentBy] = ?1 AND [targetItem] = ?2 AND [targetType] = 'replies'")
	if err != nil {
		log.Print("Bad Query: ","SELECT [targetItem] FROM [likes] WHERE [sentBy] = ?1 AND [targetItem] = ?2 AND [targetType] = 'replies'")
		return err
	}
		
	log.Print("Preparing getUserName statement.")
	getUserNameStmt, err = db.Prepare("SELECT [name] FROM [users] WHERE [uid] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [name] FROM [users] WHERE [uid] = ?1")
		return err
	}
		
	log.Print("Preparing getEmailsByUser statement.")
	getEmailsByUserStmt, err = db.Prepare("SELECT [email],[validated],[token] FROM [emails] WHERE [uid] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [email],[validated],[token] FROM [emails] WHERE [uid] = ?1")
		return err
	}
		
	log.Print("Preparing getTopicBasic statement.")
	getTopicBasicStmt, err = db.Prepare("SELECT [title],[content] FROM [topics] WHERE [tid] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [title],[content] FROM [topics] WHERE [tid] = ?1")
		return err
	}
		
	log.Print("Preparing getActivityEntry statement.")
	getActivityEntryStmt, err = db.Prepare("SELECT [actor],[targetUser],[event],[elementType],[elementID] FROM [activity_stream] WHERE [asid] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [actor],[targetUser],[event],[elementType],[elementID] FROM [activity_stream] WHERE [asid] = ?1")
		return err
	}
		
	log.Print("Preparing forumEntryExists statement.")
	forumEntryExistsStmt, err = db.Prepare("SELECT [fid] FROM [forums] WHERE [name] = '' ORDER BY fid ASC OFFSET 0 ROWS FETCH NEXT 1 ROWS ONLY")
	if err != nil {
		log.Print("Bad Query: ","SELECT [fid] FROM [forums] WHERE [name] = '' ORDER BY fid ASC OFFSET 0 ROWS FETCH NEXT 1 ROWS ONLY")
		return err
	}
		
	log.Print("Preparing groupEntryExists statement.")
	groupEntryExistsStmt, err = db.Prepare("SELECT [gid] FROM [users_groups] WHERE [name] = '' ORDER BY gid ASC OFFSET 0 ROWS FETCH NEXT 1 ROWS ONLY")
	if err != nil {
		log.Print("Bad Query: ","SELECT [gid] FROM [users_groups] WHERE [name] = '' ORDER BY gid ASC OFFSET 0 ROWS FETCH NEXT 1 ROWS ONLY")
		return err
	}
		
	log.Print("Preparing getForumTopicsOffset statement.")
	getForumTopicsOffsetStmt, err = db.Prepare("SELECT [tid],[title],[content],[createdBy],[is_closed],[sticky],[createdAt],[lastReplyAt],[lastReplyBy],[parentID],[postCount],[likeCount] FROM [topics] WHERE [parentID] = ?1 ORDER BY sticky DESC,lastReplyAt DESC,createdBy DESC OFFSET ?2 ROWS FETCH NEXT ?3 ROWS ONLY")
	if err != nil {
		log.Print("Bad Query: ","SELECT [tid],[title],[content],[createdBy],[is_closed],[sticky],[createdAt],[lastReplyAt],[lastReplyBy],[parentID],[postCount],[likeCount] FROM [topics] WHERE [parentID] = ?1 ORDER BY sticky DESC,lastReplyAt DESC,createdBy DESC OFFSET ?2 ROWS FETCH NEXT ?3 ROWS ONLY")
		return err
	}
		
	log.Print("Preparing getExpiredScheduledGroups statement.")
	getExpiredScheduledGroupsStmt, err = db.Prepare("SELECT [uid] FROM [users_groups_scheduler] WHERE GETUTCDATE() > [revert_at] AND [temporary] = 1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [uid] FROM [users_groups_scheduler] WHERE GETUTCDATE() > [revert_at] AND [temporary] = 1")
		return err
	}
		
	log.Print("Preparing getSync statement.")
	getSyncStmt, err = db.Prepare("SELECT [last_update] FROM [sync]")
	if err != nil {
		log.Print("Bad Query: ","SELECT [last_update] FROM [sync]")
		return err
	}
		
	log.Print("Preparing getAttachment statement.")
	getAttachmentStmt, err = db.Prepare("SELECT [sectionID],[sectionTable],[originID],[originTable],[uploadedBy],[path] FROM [attachments] WHERE [path] = ?1 AND [sectionID] = ?2 AND [sectionTable] = ?3")
	if err != nil {
		log.Print("Bad Query: ","SELECT [sectionID],[sectionTable],[originID],[originTable],[uploadedBy],[path] FROM [attachments] WHERE [path] = ?1 AND [sectionID] = ?2 AND [sectionTable] = ?3")
		return err
	}
		
	log.Print("Preparing getTopicRepliesOffset statement.")
	getTopicRepliesOffsetStmt, err = db.Prepare("SELECT [replies].[rid],[replies].[content],[replies].[createdBy],[replies].[createdAt],[replies].[lastEdit],[replies].[lastEditBy],[users].[avatar],[users].[name],[users].[group],[users].[url_prefix],[users].[url_name],[users].[level],[replies].[ipaddress],[replies].[likeCount],[replies].[actionType] FROM [replies] LEFT JOIN [users] ON [replies].[createdBy] = [users].[uid]  WHERE [replies].[tid] = ?1 OFFSET ?2 ROWS FETCH NEXT ?3 ROWS ONLY")
	if err != nil {
		log.Print("Bad Query: ","SELECT [replies].[rid],[replies].[content],[replies].[createdBy],[replies].[createdAt],[replies].[lastEdit],[replies].[lastEditBy],[users].[avatar],[users].[name],[users].[group],[users].[url_prefix],[users].[url_name],[users].[level],[replies].[ipaddress],[replies].[likeCount],[replies].[actionType] FROM [replies] LEFT JOIN [users] ON [replies].[createdBy] = [users].[uid]  WHERE [replies].[tid] = ?1 OFFSET ?2 ROWS FETCH NEXT ?3 ROWS ONLY")
		return err
	}
		
	log.Print("Preparing getTopicList statement.")
	getTopicListStmt, err = db.Prepare("SELECT [topics].[tid],[topics].[title],[topics].[content],[topics].[createdBy],[topics].[is_closed],[topics].[sticky],[topics].[createdAt],[topics].[parentID],[users].[name],[users].[avatar] FROM [topics] LEFT JOIN [users] ON [topics].[createdBy] = [users].[uid]  ORDER BY topics.sticky DESC,topics.lastReplyAt DESC,topics.createdBy DESC")
	if err != nil {
		log.Print("Bad Query: ","SELECT [topics].[tid],[topics].[title],[topics].[content],[topics].[createdBy],[topics].[is_closed],[topics].[sticky],[topics].[createdAt],[topics].[parentID],[users].[name],[users].[avatar] FROM [topics] LEFT JOIN [users] ON [topics].[createdBy] = [users].[uid]  ORDER BY topics.sticky DESC,topics.lastReplyAt DESC,topics.createdBy DESC")
		return err
	}
		
	log.Print("Preparing getTopicUser statement.")
	getTopicUserStmt, err = db.Prepare("SELECT [topics].[title],[topics].[content],[topics].[createdBy],[topics].[createdAt],[topics].[is_closed],[topics].[sticky],[topics].[parentID],[topics].[ipaddress],[topics].[postCount],[topics].[likeCount],[users].[name],[users].[avatar],[users].[group],[users].[url_prefix],[users].[url_name],[users].[level] FROM [topics] LEFT JOIN [users] ON [topics].[createdBy] = [users].[uid]  WHERE [tid] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [topics].[title],[topics].[content],[topics].[createdBy],[topics].[createdAt],[topics].[is_closed],[topics].[sticky],[topics].[parentID],[topics].[ipaddress],[topics].[postCount],[topics].[likeCount],[users].[name],[users].[avatar],[users].[group],[users].[url_prefix],[users].[url_name],[users].[level] FROM [topics] LEFT JOIN [users] ON [topics].[createdBy] = [users].[uid]  WHERE [tid] = ?1")
		return err
	}
		
	log.Print("Preparing getTopicByReply statement.")
	getTopicByReplyStmt, err = db.Prepare("SELECT [topics].[tid],[topics].[title],[topics].[content],[topics].[createdBy],[topics].[createdAt],[topics].[is_closed],[topics].[sticky],[topics].[parentID],[topics].[ipaddress],[topics].[postCount],[topics].[likeCount],[topics].[data] FROM [replies] LEFT JOIN [topics] ON [replies].[tid] = [topics].[tid]  WHERE [rid] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [topics].[tid],[topics].[title],[topics].[content],[topics].[createdBy],[topics].[createdAt],[topics].[is_closed],[topics].[sticky],[topics].[parentID],[topics].[ipaddress],[topics].[postCount],[topics].[likeCount],[topics].[data] FROM [replies] LEFT JOIN [topics] ON [replies].[tid] = [topics].[tid]  WHERE [rid] = ?1")
		return err
	}
		
	log.Print("Preparing getTopicReplies statement.")
	getTopicRepliesStmt, err = db.Prepare("SELECT [replies].[rid],[replies].[content],[replies].[createdBy],[replies].[createdAt],[replies].[lastEdit],[replies].[lastEditBy],[users].[avatar],[users].[name],[users].[group],[users].[url_prefix],[users].[url_name],[users].[level],[replies].[ipaddress] FROM [replies] LEFT JOIN [users] ON [replies].[createdBy] = [users].[uid]  WHERE [tid] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [replies].[rid],[replies].[content],[replies].[createdBy],[replies].[createdAt],[replies].[lastEdit],[replies].[lastEditBy],[users].[avatar],[users].[name],[users].[group],[users].[url_prefix],[users].[url_name],[users].[level],[replies].[ipaddress] FROM [replies] LEFT JOIN [users] ON [replies].[createdBy] = [users].[uid]  WHERE [tid] = ?1")
		return err
	}
		
	log.Print("Preparing getForumTopics statement.")
	getForumTopicsStmt, err = db.Prepare("SELECT [topics].[tid],[topics].[title],[topics].[content],[topics].[createdBy],[topics].[is_closed],[topics].[sticky],[topics].[createdAt],[topics].[lastReplyAt],[topics].[parentID],[users].[name],[users].[avatar] FROM [topics] LEFT JOIN [users] ON [topics].[createdBy] = [users].[uid]  WHERE [topics].[parentID] = ?1 ORDER BY topics.sticky DESC,topics.lastReplyAt DESC,topics.createdBy DESC")
	if err != nil {
		log.Print("Bad Query: ","SELECT [topics].[tid],[topics].[title],[topics].[content],[topics].[createdBy],[topics].[is_closed],[topics].[sticky],[topics].[createdAt],[topics].[lastReplyAt],[topics].[parentID],[users].[name],[users].[avatar] FROM [topics] LEFT JOIN [users] ON [topics].[createdBy] = [users].[uid]  WHERE [topics].[parentID] = ?1 ORDER BY topics.sticky DESC,topics.lastReplyAt DESC,topics.createdBy DESC")
		return err
	}
		
	log.Print("Preparing getProfileReplies statement.")
	getProfileRepliesStmt, err = db.Prepare("SELECT [users_replies].[rid],[users_replies].[content],[users_replies].[createdBy],[users_replies].[createdAt],[users_replies].[lastEdit],[users_replies].[lastEditBy],[users].[avatar],[users].[name],[users].[group] FROM [users_replies] LEFT JOIN [users] ON [users_replies].[createdBy] = [users].[uid]  WHERE [users_replies].[uid] = ?1")
	if err != nil {
		log.Print("Bad Query: ","SELECT [users_replies].[rid],[users_replies].[content],[users_replies].[createdBy],[users_replies].[createdAt],[users_replies].[lastEdit],[users_replies].[lastEditBy],[users].[avatar],[users].[name],[users].[group] FROM [users_replies] LEFT JOIN [users] ON [users_replies].[createdBy] = [users].[uid]  WHERE [users_replies].[uid] = ?1")
		return err
	}
		
	log.Print("Preparing getWatchers statement.")
	getWatchersStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing createTopic statement.")
	createTopicStmt, err = db.Prepare("INSERT INTO [topics] ([parentID],[title],[content],[parsed_content],[createdAt],[lastReplyAt],[lastReplyBy],[ipaddress],[words],[createdBy]) VALUES (?,?,?,?,GETUTCDATE(),GETUTCDATE(),?,?,?,?)")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [topics] ([parentID],[title],[content],[parsed_content],[createdAt],[lastReplyAt],[lastReplyBy],[ipaddress],[words],[createdBy]) VALUES (?,?,?,?,GETUTCDATE(),GETUTCDATE(),?,?,?,?)")
		return err
	}
		
	log.Print("Preparing createReport statement.")
	createReportStmt, err = db.Prepare("INSERT INTO [topics] ([title],[content],[parsed_content],[createdAt],[lastReplyAt],[createdBy],[data],[parentID],[css_class]) VALUES (?,?,?,GETUTCDATE(),GETUTCDATE(),?,?,1,'report')")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [topics] ([title],[content],[parsed_content],[createdAt],[lastReplyAt],[createdBy],[data],[parentID],[css_class]) VALUES (?,?,?,GETUTCDATE(),GETUTCDATE(),?,?,1,'report')")
		return err
	}
		
	log.Print("Preparing createReply statement.")
	createReplyStmt, err = db.Prepare("INSERT INTO [replies] ([tid],[content],[parsed_content],[createdAt],[ipaddress],[words],[createdBy]) VALUES (?,?,?,GETUTCDATE(),?,?,?)")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [replies] ([tid],[content],[parsed_content],[createdAt],[ipaddress],[words],[createdBy]) VALUES (?,?,?,GETUTCDATE(),?,?,?)")
		return err
	}
		
	log.Print("Preparing createActionReply statement.")
	createActionReplyStmt, err = db.Prepare("INSERT INTO [replies] ([tid],[actionType],[ipaddress],[createdBy]) VALUES (?,?,?,?)")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [replies] ([tid],[actionType],[ipaddress],[createdBy]) VALUES (?,?,?,?)")
		return err
	}
		
	log.Print("Preparing createLike statement.")
	createLikeStmt, err = db.Prepare("INSERT INTO [likes] ([weight],[targetItem],[targetType],[sentBy]) VALUES (?,?,?,?)")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [likes] ([weight],[targetItem],[targetType],[sentBy]) VALUES (?,?,?,?)")
		return err
	}
		
	log.Print("Preparing addActivity statement.")
	addActivityStmt, err = db.Prepare("INSERT INTO [activity_stream] ([actor],[targetUser],[event],[elementType],[elementID]) VALUES (?,?,?,?,?)")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [activity_stream] ([actor],[targetUser],[event],[elementType],[elementID]) VALUES (?,?,?,?,?)")
		return err
	}
		
	log.Print("Preparing notifyOne statement.")
	notifyOneStmt, err = db.Prepare("INSERT INTO [activity_stream_matches] ([watcher],[asid]) VALUES (?,?)")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [activity_stream_matches] ([watcher],[asid]) VALUES (?,?)")
		return err
	}
		
	log.Print("Preparing addEmail statement.")
	addEmailStmt, err = db.Prepare("INSERT INTO [emails] ([email],[uid],[validated],[token]) VALUES (?,?,?,?)")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [emails] ([email],[uid],[validated],[token]) VALUES (?,?,?,?)")
		return err
	}
		
	log.Print("Preparing createProfileReply statement.")
	createProfileReplyStmt, err = db.Prepare("INSERT INTO [users_replies] ([uid],[content],[parsed_content],[createdAt],[createdBy],[ipaddress]) VALUES (?,?,?,GETUTCDATE(),?,?)")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [users_replies] ([uid],[content],[parsed_content],[createdAt],[createdBy],[ipaddress]) VALUES (?,?,?,GETUTCDATE(),?,?)")
		return err
	}
		
	log.Print("Preparing addSubscription statement.")
	addSubscriptionStmt, err = db.Prepare("INSERT INTO [activity_subscriptions] ([user],[targetID],[targetType],[level]) VALUES (?,?,?,2)")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [activity_subscriptions] ([user],[targetID],[targetType],[level]) VALUES (?,?,?,2)")
		return err
	}
		
	log.Print("Preparing createForum statement.")
	createForumStmt, err = db.Prepare("INSERT INTO [forums] ([name],[desc],[active],[preset]) VALUES (?,?,?,?)")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [forums] ([name],[desc],[active],[preset]) VALUES (?,?,?,?)")
		return err
	}
		
	log.Print("Preparing addForumPermsToForum statement.")
	addForumPermsToForumStmt, err = db.Prepare("INSERT INTO [forums_permissions] ([gid],[fid],[preset],[permissions]) VALUES (?,?,?,?)")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [forums_permissions] ([gid],[fid],[preset],[permissions]) VALUES (?,?,?,?)")
		return err
	}
		
	log.Print("Preparing addPlugin statement.")
	addPluginStmt, err = db.Prepare("INSERT INTO [plugins] ([uname],[active],[installed]) VALUES (?,?,?)")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [plugins] ([uname],[active],[installed]) VALUES (?,?,?)")
		return err
	}
		
	log.Print("Preparing addTheme statement.")
	addThemeStmt, err = db.Prepare("INSERT INTO [themes] ([uname],[default]) VALUES (?,?)")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [themes] ([uname],[default]) VALUES (?,?)")
		return err
	}
		
	log.Print("Preparing createGroup statement.")
	createGroupStmt, err = db.Prepare("INSERT INTO [users_groups] ([name],[tag],[is_admin],[is_mod],[is_banned],[permissions]) VALUES (?,?,?,?,?,?)")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [users_groups] ([name],[tag],[is_admin],[is_mod],[is_banned],[permissions]) VALUES (?,?,?,?,?,?)")
		return err
	}
		
	log.Print("Preparing addModlogEntry statement.")
	addModlogEntryStmt, err = db.Prepare("INSERT INTO [moderation_logs] ([action],[elementID],[elementType],[ipaddress],[actorID],[doneAt]) VALUES (?,?,?,?,?,GETUTCDATE())")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [moderation_logs] ([action],[elementID],[elementType],[ipaddress],[actorID],[doneAt]) VALUES (?,?,?,?,?,GETUTCDATE())")
		return err
	}
		
	log.Print("Preparing addAdminlogEntry statement.")
	addAdminlogEntryStmt, err = db.Prepare("INSERT INTO [administration_logs] ([action],[elementID],[elementType],[ipaddress],[actorID],[doneAt]) VALUES (?,?,?,?,?,GETUTCDATE())")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [administration_logs] ([action],[elementID],[elementType],[ipaddress],[actorID],[doneAt]) VALUES (?,?,?,?,?,GETUTCDATE())")
		return err
	}
		
	log.Print("Preparing addAttachment statement.")
	addAttachmentStmt, err = db.Prepare("INSERT INTO [attachments] ([sectionID],[sectionTable],[originID],[originTable],[uploadedBy],[path]) VALUES (?,?,?,?,?,?)")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [attachments] ([sectionID],[sectionTable],[originID],[originTable],[uploadedBy],[path]) VALUES (?,?,?,?,?,?)")
		return err
	}
		
	log.Print("Preparing createWordFilter statement.")
	createWordFilterStmt, err = db.Prepare("INSERT INTO [word_filters] ([find],[replacement]) VALUES (?,?)")
	if err != nil {
		log.Print("Bad Query: ","INSERT INTO [word_filters] ([find],[replacement]) VALUES (?,?)")
		return err
	}
		
	log.Print("Preparing addForumPermsToGroup statement.")
	addForumPermsToGroupStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing replaceScheduleGroup statement.")
	replaceScheduleGroupStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing addRepliesToTopic statement.")
	addRepliesToTopicStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing removeRepliesFromTopic statement.")
	removeRepliesFromTopicStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing addTopicsToForum statement.")
	addTopicsToForumStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing removeTopicsFromForum statement.")
	removeTopicsFromForumStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing updateForumCache statement.")
	updateForumCacheStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing addLikesToTopic statement.")
	addLikesToTopicStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing addLikesToReply statement.")
	addLikesToReplyStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing editTopic statement.")
	editTopicStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing editReply statement.")
	editReplyStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing stickTopic statement.")
	stickTopicStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing unstickTopic statement.")
	unstickTopicStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing lockTopic statement.")
	lockTopicStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing unlockTopic statement.")
	unlockTopicStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing updateLastIP statement.")
	updateLastIPStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing updateSession statement.")
	updateSessionStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing setPassword statement.")
	setPasswordStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing setAvatar statement.")
	setAvatarStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing setUsername statement.")
	setUsernameStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing changeGroup statement.")
	changeGroupStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing activateUser statement.")
	activateUserStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing updateUserLevel statement.")
	updateUserLevelStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing incrementUserScore statement.")
	incrementUserScoreStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing incrementUserPosts statement.")
	incrementUserPostsStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing incrementUserBigposts statement.")
	incrementUserBigpostsStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing incrementUserMegaposts statement.")
	incrementUserMegapostsStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing incrementUserTopics statement.")
	incrementUserTopicsStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing editProfileReply statement.")
	editProfileReplyStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing updateForum statement.")
	updateForumStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing updateSetting statement.")
	updateSettingStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing updatePlugin statement.")
	updatePluginStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing updatePluginInstall statement.")
	updatePluginInstallStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing updateTheme statement.")
	updateThemeStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing updateUser statement.")
	updateUserStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing updateGroupPerms statement.")
	updateGroupPermsStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing updateGroupRank statement.")
	updateGroupRankStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing updateGroup statement.")
	updateGroupStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing updateEmail statement.")
	updateEmailStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing verifyEmail statement.")
	verifyEmailStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing setTempGroup statement.")
	setTempGroupStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing updateWordFilter statement.")
	updateWordFilterStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing bumpSync statement.")
	bumpSyncStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing deleteReply statement.")
	deleteReplyStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing deleteProfileReply statement.")
	deleteProfileReplyStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing deleteForumPermsByForum statement.")
	deleteForumPermsByForumStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing deleteActivityStreamMatch statement.")
	deleteActivityStreamMatchStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing deleteWordFilter statement.")
	deleteWordFilterStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing reportExists statement.")
	reportExistsStmt, err = db.Prepare("SELECT COUNT(*) AS [count] FROM [topics] WHERE [data] = ? AND [data] != '' AND [parentID] = 1")
	if err != nil {
		log.Print("Bad Query: ","SELECT COUNT(*) AS [count] FROM [topics] WHERE [data] = ? AND [data] != '' AND [parentID] = 1")
		return err
	}
		
	log.Print("Preparing groupCount statement.")
	groupCountStmt, err = db.Prepare("SELECT COUNT(*) AS [count] FROM [users_groups]")
	if err != nil {
		log.Print("Bad Query: ","SELECT COUNT(*) AS [count] FROM [users_groups]")
		return err
	}
		
	log.Print("Preparing modlogCount statement.")
	modlogCountStmt, err = db.Prepare("SELECT COUNT(*) AS [count] FROM [moderation_logs]")
	if err != nil {
		log.Print("Bad Query: ","SELECT COUNT(*) AS [count] FROM [moderation_logs]")
		return err
	}
		
	log.Print("Preparing addForumPermsToForumAdmins statement.")
	addForumPermsToForumAdminsStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing addForumPermsToForumStaff statement.")
	addForumPermsToForumStaffStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing addForumPermsToForumMembers statement.")
	addForumPermsToForumMembersStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
		
	log.Print("Preparing notifyWatchers statement.")
	notifyWatchersStmt, err = db.Prepare("")
	if err != nil {
		log.Print("Bad Query: ","")
		return err
	}
	
	return nil
}
