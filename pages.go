package main

import (
	//"fmt"
	"bytes"
	"html/template"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type HeaderVars struct {
	NoticeList  []string
	Scripts     []string
	Stylesheets []string
	Widgets     PageWidgets
	Site        *Site
	Settings    map[string]interface{}
	Themes      map[string]Theme // TODO: Use a slice containing every theme instead of the main map for speed
	ThemeName   string
	ExtData     ExtData
}

// TODO: Add this to routes which don't use templates. E.g. Json APIs.
type HeaderLite struct {
	Site     *Site
	Settings SettingBox
	ExtData  ExtData
}

type PageWidgets struct {
	LeftSidebar  template.HTML
	RightSidebar template.HTML
}

// TODO: Add a ExtDataHolder interface with methods for manipulating the contents?
// ? - Could we use a sync.Map instead?
type ExtData struct {
	items map[string]interface{} // Key: pluginname
	sync.RWMutex
}

type Page struct {
	Title       string
	CurrentUser User
	Header      *HeaderVars
	ItemList    []interface{}
	Something   interface{}
}

type TopicPage struct {
	Title       string
	CurrentUser User
	Header      *HeaderVars
	ItemList    []ReplyUser
	Topic       TopicUser
	Page        int
	LastPage    int
}

type TopicsPage struct {
	Title        string
	CurrentUser  User
	Header       *HeaderVars
	TopicList    []*TopicsRow
	ForumList    []Forum
	DefaultForum int
}

type ForumPage struct {
	Title       string
	CurrentUser User
	Header      *HeaderVars
	ItemList    []*TopicsRow
	Forum       *Forum
	Page        int
	LastPage    int
}

type ForumsPage struct {
	Title       string
	CurrentUser User
	Header      *HeaderVars
	ItemList    []Forum
}

type ProfilePage struct {
	Title        string
	CurrentUser  User
	Header       *HeaderVars
	ItemList     []ReplyUser
	ProfileOwner User
}

type CreateTopicPage struct {
	Title       string
	CurrentUser User
	Header      *HeaderVars
	ItemList    []Forum
	FID         int
}

type IPSearchPage struct {
	Title       string
	CurrentUser User
	Header      *HeaderVars
	ItemList    map[int]*User
	IP          string
}

type PanelStats struct {
	Users       int
	Groups      int
	Forums      int
	Settings    int
	WordFilters int
	Themes      int
	Reports     int
}

type PanelPage struct {
	Title       string
	CurrentUser User
	Header      *HeaderVars
	Stats       PanelStats
	ItemList    []interface{}
	Something   interface{}
}

type GridElement struct {
	ID         string
	Body       string
	Order      int // For future use
	Class      string
	Background string
	TextColour string
	Note       string
}

type PanelDashboardPage struct {
	Title       string
	CurrentUser User
	Header      *HeaderVars
	Stats       PanelStats
	GridItems   []GridElement
}

type PanelThemesPage struct {
	Title         string
	CurrentUser   User
	Header        *HeaderVars
	Stats         PanelStats
	PrimaryThemes []Theme
	VariantThemes []Theme
}

type PanelUserPage struct {
	Title       string
	CurrentUser User
	Header      *HeaderVars
	Stats       PanelStats
	ItemList    []User
	PageList    []int
	Page        int
	LastPage    int
}

type PanelGroupPage struct {
	Title       string
	CurrentUser User
	Header      *HeaderVars
	Stats       PanelStats
	ItemList    []GroupAdmin
	PageList    []int
	Page        int
	LastPage    int
}

type PanelEditGroupPage struct {
	Title       string
	CurrentUser User
	Header      *HeaderVars
	Stats       PanelStats
	ID          int
	Name        string
	Tag         string
	Rank        string
	DisableRank bool
}

type GroupForumPermPreset struct {
	Group  *Group
	Preset string
}

type PanelEditForumPage struct {
	Title       string
	CurrentUser User
	Header      *HeaderVars
	Stats       PanelStats
	ID          int
	Name        string
	Desc        string
	Active      bool
	Preset      string
	Groups      []GroupForumPermPreset
}

/*type NameLangPair struct {
	Name    string
	LangStr string
}*/

type NameLangToggle struct {
	Name    string
	LangStr string
	Toggle  bool
}

type PanelEditGroupPermsPage struct {
	Title       string
	CurrentUser User
	Header      *HeaderVars
	Stats       PanelStats
	ID          int
	Name        string
	LocalPerms  []NameLangToggle
	GlobalPerms []NameLangToggle
}

type backupItem struct {
	SQLURL string

	// TODO: Add an easier to parse format here for Gosora to be able to more easily reimport portions of the dump and to strip unneccesary data (e.g. table defs and parsed post data)

	Timestamp time.Time
}

type PanelBackupPage struct {
	Title       string
	CurrentUser User
	Header      *HeaderVars
	Stats       PanelStats
	Backups     []backupItem
}

type logItem struct {
	Action    template.HTML
	IPAddress string
	DoneAt    string
}

type PanelLogsPage struct {
	Title       string
	CurrentUser User
	Header      *HeaderVars
	Stats       PanelStats
	Logs        []logItem
	PageList    []int
	Page        int
	LastPage    int
}

type PanelDebugPage struct {
	Title       string
	CurrentUser User
	Header      *HeaderVars
	Stats       PanelStats
	Uptime      string
	OpenConns   int
	DBAdapter   string
}

type PageSimple struct {
	Title     string
	Something interface{}
}

type AreYouSure struct {
	URL     string
	Message string
}

var spaceGap = []byte("          ")
var httpProtBytes = []byte("http://")
var invalidURL = []byte("<span style='color: red;'>[Invalid URL]</span>")
var invalidTopic = []byte("<span style='color: red;'>[Invalid Topic]</span>")
var invalidProfile = []byte("<span style='color: red;'>[Invalid Profile]</span>")
var invalidForum = []byte("<span style='color: red;'>[Invalid Forum]</span>")
var unknownMedia = []byte("<span style='color: red;'>[Unknown Media]</span>")
var urlOpen = []byte("<a href='")
var urlOpen2 = []byte("'>")
var bytesSinglequote = []byte("'")
var bytesGreaterthan = []byte(">")
var urlMention = []byte(" class='mention'")
var urlClose = []byte("</a>")
var imageOpen = []byte("<a href=\"")
var imageOpen2 = []byte("\"><img src='")
var imageClose = []byte("' class='postImage' /></a>")
var urlpattern = `(?s)([ {1}])((http|https|ftp|mailto)*)(:{??)\/\/([\.a-zA-Z\/]+)([ {1}])`
var urlReg *regexp.Regexp

func init() {
	urlReg = regexp.MustCompile(urlpattern)
}

// TODO: Write a test for this
func shortcodeToUnicode(msg string) string {
	//re := regexp.MustCompile(":(.):")
	msg = strings.Replace(msg, ":grinning:", "😀", -1)
	msg = strings.Replace(msg, ":grin:", "😁", -1)
	msg = strings.Replace(msg, ":joy:", "😂", -1)
	msg = strings.Replace(msg, ":rofl:", "🤣", -1)
	msg = strings.Replace(msg, ":smiley:", "😃", -1)
	msg = strings.Replace(msg, ":smile:", "😄", -1)
	msg = strings.Replace(msg, ":sweat_smile:", "😅", -1)
	msg = strings.Replace(msg, ":laughing:", "😆", -1)
	msg = strings.Replace(msg, ":satisfied:", "😆", -1)
	msg = strings.Replace(msg, ":wink:", "😉", -1)
	msg = strings.Replace(msg, ":blush:", "😊", -1)
	msg = strings.Replace(msg, ":yum:", "😋", -1)
	msg = strings.Replace(msg, ":sunglasses:", "😎", -1)
	msg = strings.Replace(msg, ":heart_eyes:", "😍", -1)
	msg = strings.Replace(msg, ":kissing_heart:", "😘", -1)
	msg = strings.Replace(msg, ":kissing:", "😗", -1)
	msg = strings.Replace(msg, ":kissing_smiling_eyes:", "😙", -1)
	msg = strings.Replace(msg, ":kissing_closed_eyes:", "😚", -1)
	msg = strings.Replace(msg, ":relaxed:", "☺️", -1)
	msg = strings.Replace(msg, ":slight_smile:", "🙂", -1)
	msg = strings.Replace(msg, ":hugging:", "🤗", -1)
	msg = strings.Replace(msg, ":thinking:", "🤔", -1)
	msg = strings.Replace(msg, ":neutral_face:", "😐", -1)
	msg = strings.Replace(msg, ":expressionless:", "😑", -1)
	msg = strings.Replace(msg, ":no_mouth:", "😶", -1)
	msg = strings.Replace(msg, ":rolling_eyes:", "🙄", -1)
	msg = strings.Replace(msg, ":smirk:", "😏", -1)
	msg = strings.Replace(msg, ":persevere:", "😣", -1)
	msg = strings.Replace(msg, ":disappointed_relieved:", "😥", -1)
	msg = strings.Replace(msg, ":open_mouth:", "😮", -1)
	msg = strings.Replace(msg, ":zipper_mouth:", "🤐", -1)
	msg = strings.Replace(msg, ":hushed:", "😯", -1)
	msg = strings.Replace(msg, ":sleepy:", "😪", -1)
	msg = strings.Replace(msg, ":tired_face:", "😫", -1)
	msg = strings.Replace(msg, ":sleeping:", "😴", -1)
	msg = strings.Replace(msg, ":relieved:", "😌", -1)
	msg = strings.Replace(msg, ":nerd:", "🤓", -1)
	msg = strings.Replace(msg, ":stuck_out_tongue:", "😛", -1)
	msg = strings.Replace(msg, ":worried:", "😟", -1)
	msg = strings.Replace(msg, ":drooling_face:", "🤤", -1)
	msg = strings.Replace(msg, ":disappointed:", "😞", -1)
	msg = strings.Replace(msg, ":astonished:", "😲", -1)
	msg = strings.Replace(msg, ":slight_frown:", "🙁", -1)
	msg = strings.Replace(msg, ":skull_crossbones:", "☠️", -1)
	msg = strings.Replace(msg, ":skull:", "💀", -1)
	msg = strings.Replace(msg, ":point_up:", "☝️", -1)
	msg = strings.Replace(msg, ":v:", "✌️️", -1)
	msg = strings.Replace(msg, ":writing_hand:", "✍️", -1)
	msg = strings.Replace(msg, ":heart:", "❤️️", -1)
	msg = strings.Replace(msg, ":heart_exclamation:", "❣️", -1)
	msg = strings.Replace(msg, ":hotsprings:", "♨️", -1)
	msg = strings.Replace(msg, ":airplane:", "✈️️", -1)
	msg = strings.Replace(msg, ":hourglass:", "⌛", -1)
	msg = strings.Replace(msg, ":watch:", "⌚", -1)
	msg = strings.Replace(msg, ":comet:", "☄️", -1)
	msg = strings.Replace(msg, ":snowflake:", "❄️", -1)
	msg = strings.Replace(msg, ":cloud:", "☁️", -1)
	msg = strings.Replace(msg, ":sunny:", "☀️", -1)
	msg = strings.Replace(msg, ":spades:", "♠️", -1)
	msg = strings.Replace(msg, ":hearts:", "♥️️", -1)
	msg = strings.Replace(msg, ":diamonds:", "♦️", -1)
	msg = strings.Replace(msg, ":clubs:", "♣️", -1)
	msg = strings.Replace(msg, ":phone:", "☎️", -1)
	msg = strings.Replace(msg, ":telephone:", "☎️", -1)
	msg = strings.Replace(msg, ":biohazard:", "☣️", -1)
	msg = strings.Replace(msg, ":radioactive:", "☢️", -1)
	msg = strings.Replace(msg, ":scissors:", "✂️", -1)
	msg = strings.Replace(msg, ":arrow_upper_right:", "↗️", -1)
	msg = strings.Replace(msg, ":arrow_right:", "➡️", -1)
	msg = strings.Replace(msg, ":arrow_lower_right:", "↘️", -1)
	msg = strings.Replace(msg, ":arrow_lower_left:", "↙️", -1)
	msg = strings.Replace(msg, ":arrow_upper_left:", "↖️", -1)
	msg = strings.Replace(msg, ":arrow_up_down:", "↕️", -1)
	msg = strings.Replace(msg, ":left_right_arrow:", "↔️", -1)
	msg = strings.Replace(msg, ":leftwards_arrow_with_hook:", "↩️", -1)
	msg = strings.Replace(msg, ":arrow_right_hook:", "↪️", -1)
	msg = strings.Replace(msg, ":arrow_forward:", "▶️", -1)
	msg = strings.Replace(msg, ":arrow_backward:", "◀️", -1)
	msg = strings.Replace(msg, ":female:", "♀️", -1)
	msg = strings.Replace(msg, ":male:", "♂️", -1)
	msg = strings.Replace(msg, ":ballot_box_with_check:", "☑️", -1)
	msg = strings.Replace(msg, ":heavy_check_mark:", "✔️️", -1)
	msg = strings.Replace(msg, ":heavy_multiplication_x:", "✖️", -1)
	msg = strings.Replace(msg, ":pisces:", "♓", -1)
	msg = strings.Replace(msg, ":aquarius:", "♒", -1)
	msg = strings.Replace(msg, ":capricorn:", "♑", -1)
	msg = strings.Replace(msg, ":sagittarius:", "♐", -1)
	msg = strings.Replace(msg, ":scorpius:", "♏", -1)
	msg = strings.Replace(msg, ":libra:", "♎", -1)
	msg = strings.Replace(msg, ":virgo:", "♍", -1)
	msg = strings.Replace(msg, ":leo:", "♌", -1)
	msg = strings.Replace(msg, ":cancer:", "♋", -1)
	msg = strings.Replace(msg, ":gemini:", "♊", -1)
	msg = strings.Replace(msg, ":taurus:", "♉", -1)
	msg = strings.Replace(msg, ":aries:", "♈", -1)
	msg = strings.Replace(msg, ":peace:", "☮️", -1)
	msg = strings.Replace(msg, ":eight_spoked_asterisk:", "✳️", -1)
	msg = strings.Replace(msg, ":eight_pointed_black_star:", "✴️", -1)
	msg = strings.Replace(msg, ":snowman2:", "☃️", -1)
	msg = strings.Replace(msg, ":umbrella2:", "☂️", -1)
	msg = strings.Replace(msg, ":pencil2:", "✏️", -1)
	msg = strings.Replace(msg, ":black_nib:", "✒️", -1)
	msg = strings.Replace(msg, ":email:", "✉️", -1)
	msg = strings.Replace(msg, ":envelope:", "✉️", -1)
	msg = strings.Replace(msg, ":keyboard:", "⌨️", -1)
	msg = strings.Replace(msg, ":white_small_square:", "▫️", -1)
	msg = strings.Replace(msg, ":black_small_square:", "▪️", -1)
	msg = strings.Replace(msg, ":secret:", "㊙️", -1)
	msg = strings.Replace(msg, ":congratulations:", "㊗️", -1)
	msg = strings.Replace(msg, ":m:", "Ⓜ️", -1)
	msg = strings.Replace(msg, ":tm:", "™️️", -1)
	msg = strings.Replace(msg, ":registered:", "®️", -1)
	msg = strings.Replace(msg, ":copyright:", "©️", -1)
	msg = strings.Replace(msg, ":wavy_dash:", "〰️", -1)
	msg = strings.Replace(msg, ":bangbang:", "‼️", -1)
	msg = strings.Replace(msg, ":sparkle:", "❇️", -1)
	msg = strings.Replace(msg, ":star_of_david:", "✡️", -1)
	msg = strings.Replace(msg, ":wheel_of_dharma:", "☸️", -1)
	msg = strings.Replace(msg, ":yin_yang:", "☯️", -1)
	msg = strings.Replace(msg, ":cross:", "✝️", -1)
	msg = strings.Replace(msg, ":orthodox_cross:", "☦️", -1)
	msg = strings.Replace(msg, ":star_and_crescent:", "☪️", -1)
	msg = strings.Replace(msg, ":frowning2:", "☹️", -1)
	msg = strings.Replace(msg, ":information_source:", "ℹ️", -1)
	msg = strings.Replace(msg, ":interrobang:", "⁉️", -1)

	return msg
}

func preparseMessage(msg string) string {
	if sshooks["preparse_preassign"] != nil {
		msg = runSshook("preparse_preassign", msg)
	}
	return shortcodeToUnicode(msg)
}

// TODO: Write a test for this
// TODO: We need a lot more hooks here. E.g. To add custom media types and handlers.
func parseMessage(msg string, sectionID int, sectionType string /*, user User*/) string {
	msg = strings.Replace(msg, ":)", "😀", -1)
	msg = strings.Replace(msg, ":(", "😞", -1)
	msg = strings.Replace(msg, ":D", "😃", -1)
	msg = strings.Replace(msg, ":P", "😛", -1)
	msg = strings.Replace(msg, ":O", "😲", -1)
	msg = strings.Replace(msg, ":p", "😛", -1)
	msg = strings.Replace(msg, ":o", "😲", -1)
	msg = strings.Replace(msg, ";)", "😉", -1)
	//msg = url_reg.ReplaceAllString(msg,"<a href=\"$2$3//$4\" rel=\"nofollow\">$2$3//$4</a>")

	// Word filter list. E.g. Swear words and other things the admins don't like
	wordFilters := wordFilterBox.Load().(WordFilterBox)
	for _, filter := range wordFilters {
		msg = strings.Replace(msg, filter.Find, filter.Replacement, -1)
	}

	// Search for URLs, mentions and hashlinks in the messages...
	//log.Print("Parser Loop!")
	var msgbytes = []byte(msg)
	var outbytes []byte
	msgbytes = append(msgbytes, spaceGap...)
	//log.Printf("string(msgbytes) %+v\n", `"`+string(msgbytes)+`"`)
	var lastItem = 0
	var i = 0
	for ; len(msgbytes) > (i + 1); i++ {
		//log.Print("Index: ",i)
		//log.Print("Index Item: ",msgbytes[i])
		//log.Print("string(msgbytes[i]): ",string(msgbytes[i]))
		//log.Print("End Index")
		if (i == 0 && (msgbytes[0] > 32)) || ((msgbytes[i] < 33) && (msgbytes[i+1] > 32)) {
			//log.Print("IN ",msgbytes[i])
			if (i != 0) || msgbytes[i] < 33 {
				i++
			}

			if msgbytes[i] == '#' {
				//log.Print("IN #")
				if bytes.Equal(msgbytes[i+1:i+5], []byte("tid-")) {
					outbytes = append(outbytes, msgbytes[lastItem:i]...)
					i += 5
					start := i
					tid, intLen := coerceIntBytes(msgbytes[start:])
					i += intLen

					topic, err := topics.Get(tid)
					if err != nil || !fstore.Exists(topic.ParentID) {
						outbytes = append(outbytes, invalidTopic...)
						lastItem = i
						continue
					}

					outbytes = append(outbytes, urlOpen...)
					var urlBit = []byte(buildTopicURL("", tid))
					outbytes = append(outbytes, urlBit...)
					outbytes = append(outbytes, urlOpen2...)
					var tidBit = []byte("#tid-" + strconv.Itoa(tid))
					outbytes = append(outbytes, tidBit...)
					outbytes = append(outbytes, urlClose...)
					lastItem = i

					//log.Print("string(msgbytes): ",string(msgbytes))
					//log.Print("msgbytes: ",msgbytes)
					//log.Print("msgbytes[lastItem - 1]: ",msgbytes[lastItem - 1])
					//log.Print("lastItem - 1: ",lastItem - 1)
					//log.Print("msgbytes[lastItem]: ",msgbytes[lastItem])
					//log.Print("lastItem: ",lastItem)
				} else if bytes.Equal(msgbytes[i+1:i+5], []byte("rid-")) {
					outbytes = append(outbytes, msgbytes[lastItem:i]...)
					i += 5
					start := i
					rid, intLen := coerceIntBytes(msgbytes[start:])
					i += intLen

					topic, err := getTopicByReply(rid)
					if err != nil || !fstore.Exists(topic.ParentID) {
						outbytes = append(outbytes, invalidTopic...)
						lastItem = i
						continue
					}

					outbytes = append(outbytes, urlOpen...)
					var urlBit = []byte(buildTopicURL("", topic.ID))
					outbytes = append(outbytes, urlBit...)
					outbytes = append(outbytes, urlOpen2...)
					var ridBit = []byte("#rid-" + strconv.Itoa(rid))
					outbytes = append(outbytes, ridBit...)
					outbytes = append(outbytes, urlClose...)
					lastItem = i
				} else if bytes.Equal(msgbytes[i+1:i+5], []byte("fid-")) {
					outbytes = append(outbytes, msgbytes[lastItem:i]...)
					i += 5
					start := i
					fid, intLen := coerceIntBytes(msgbytes[start:])
					i += intLen

					if !fstore.Exists(fid) {
						outbytes = append(outbytes, invalidForum...)
						lastItem = i
						continue
					}

					outbytes = append(outbytes, urlOpen...)
					var urlBit = []byte(buildForumURL("", fid))
					outbytes = append(outbytes, urlBit...)
					outbytes = append(outbytes, urlOpen2...)
					var fidBit = []byte("#fid-" + strconv.Itoa(fid))
					outbytes = append(outbytes, fidBit...)
					outbytes = append(outbytes, urlClose...)
					lastItem = i
				} else {
					// TODO: Forum Shortcode Link
				}
			} else if msgbytes[i] == '@' {
				//log.Print("IN @")
				outbytes = append(outbytes, msgbytes[lastItem:i]...)
				i++
				start := i
				uid, intLen := coerceIntBytes(msgbytes[start:])
				i += intLen

				menUser, err := users.Get(uid)
				if err != nil {
					outbytes = append(outbytes, invalidProfile...)
					lastItem = i
					continue
				}

				outbytes = append(outbytes, urlOpen...)
				var urlBit = []byte(menUser.Link)
				outbytes = append(outbytes, urlBit...)
				outbytes = append(outbytes, bytesSinglequote...)
				outbytes = append(outbytes, urlMention...)
				outbytes = append(outbytes, bytesGreaterthan...)
				var uidBit = []byte("@" + menUser.Name)
				outbytes = append(outbytes, uidBit...)
				outbytes = append(outbytes, urlClose...)
				lastItem = i

				//log.Print(string(msgbytes))
				//log.Print(msgbytes)
				//log.Print("msgbytes[lastItem - 1]: ", msgbytes[lastItem - 1])
				//log.Print("lastItem - 1: ", lastItem - 1)
				//log.Print("msgbytes[lastItem]: ", msgbytes[lastItem])
				//log.Print("lastItem: ", lastItem)
			} else if msgbytes[i] == 'h' || msgbytes[i] == 'f' || msgbytes[i] == 'g' {
				//log.Print("IN hfg")
				if msgbytes[i+1] == 't' && msgbytes[i+2] == 't' && msgbytes[i+3] == 'p' {
					if msgbytes[i+4] == 's' && msgbytes[i+5] == ':' && msgbytes[i+6] == '/' && msgbytes[i+7] == '/' {
						// Do nothing
					} else if msgbytes[i+4] == ':' && msgbytes[i+5] == '/' && msgbytes[i+6] == '/' {
						// Do nothing
					} else {
						continue
					}
				} else if msgbytes[i+1] == 't' && msgbytes[i+2] == 'p' && msgbytes[i+3] == ':' && msgbytes[i+4] == '/' && msgbytes[i+5] == '/' {
					// Do nothing
				} else if msgbytes[i+1] == 'i' && msgbytes[i+2] == 't' && msgbytes[i+3] == ':' && msgbytes[i+4] == '/' && msgbytes[i+5] == '/' {
					// Do nothing
				} else {
					continue
				}

				//log.Print("Normal URL")
				outbytes = append(outbytes, msgbytes[lastItem:i]...)
				urlLen := partialURLBytesLen(msgbytes[i:])
				if msgbytes[i+urlLen] > 32 { // space and invisibles
					//log.Print("INVALID URL")
					//log.Print("msgbytes[i+urlLen]: ", msgbytes[i+urlLen])
					//log.Print("string(msgbytes[i+urlLen]): ", string(msgbytes[i+urlLen]))
					//log.Print("msgbytes[i:i+urlLen]: ", msgbytes[i:i+urlLen])
					//log.Print("string(msgbytes[i:i+urlLen]): ", string(msgbytes[i:i+urlLen]))
					outbytes = append(outbytes, invalidURL...)
					i += urlLen
					continue
				}

				media, ok := parseMediaBytes(msgbytes[i : i+urlLen])
				if !ok {
					outbytes = append(outbytes, invalidURL...)
					i += urlLen
					continue
				}

				if media.Type == "attach" {
					outbytes = append(outbytes, imageOpen...)
					outbytes = append(outbytes, []byte(media.URL+"?sectionID="+strconv.Itoa(sectionID)+"&sectionType="+sectionType)...)
					outbytes = append(outbytes, imageOpen2...)
					outbytes = append(outbytes, []byte(media.URL+"?sectionID="+strconv.Itoa(sectionID)+"&sectionType="+sectionType)...)
					outbytes = append(outbytes, imageClose...)
					i += urlLen
					lastItem = i
					continue
				} else if media.Type == "image" {
					outbytes = append(outbytes, imageOpen...)
					outbytes = append(outbytes, []byte(media.URL)...)
					outbytes = append(outbytes, imageOpen2...)
					outbytes = append(outbytes, []byte(media.URL)...)
					outbytes = append(outbytes, imageClose...)
					i += urlLen
					lastItem = i
					continue
				} else if media.Type == "raw" {
					outbytes = append(outbytes, []byte(media.Body)...)
					i += urlLen
					lastItem = i
					continue
				} else if media.Type != "" {
					outbytes = append(outbytes, unknownMedia...)
					i += urlLen
					continue
				}

				outbytes = append(outbytes, urlOpen...)
				outbytes = append(outbytes, msgbytes[i:i+urlLen]...)
				outbytes = append(outbytes, urlOpen2...)
				outbytes = append(outbytes, msgbytes[i:i+urlLen]...)
				outbytes = append(outbytes, urlClose...)
				i += urlLen
				lastItem = i
			} else if msgbytes[i] == '/' && msgbytes[i+1] == '/' {
				outbytes = append(outbytes, msgbytes[lastItem:i]...)
				urlLen := partialURLBytesLen(msgbytes[i:])
				if msgbytes[i+urlLen] > 32 { // space and invisibles
					//log.Print("INVALID URL")
					//log.Print("msgbytes[i+urlLen]: ", msgbytes[i+urlLen])
					//log.Print("string(msgbytes[i+urlLen]): ", string(msgbytes[i+urlLen]))
					//log.Print("msgbytes[i:i+urlLen]: ", msgbytes[i:i+urlLen])
					//log.Print("string(msgbytes[i:i+urlLen]): ", string(msgbytes[i:i+urlLen]))
					outbytes = append(outbytes, invalidURL...)
					i += urlLen
					continue
				}

				//log.Print("VALID URL")
				//log.Print("msgbytes[i:i+urlLen]: ", msgbytes[i:i+urlLen])
				//log.Print("string(msgbytes[i:i+urlLen]): ", string(msgbytes[i:i+urlLen]))
				media, ok := parseMediaBytes(msgbytes[i : i+urlLen])
				if !ok {
					outbytes = append(outbytes, invalidURL...)
					i += urlLen
					continue
				}

				if media.Type == "attach" {
					outbytes = append(outbytes, imageOpen...)
					outbytes = append(outbytes, []byte(media.URL+"?sectionID="+strconv.Itoa(sectionID)+"&sectionType="+sectionType)...)
					outbytes = append(outbytes, imageOpen2...)
					outbytes = append(outbytes, []byte(media.URL+"?sectionID="+strconv.Itoa(sectionID)+"&sectionType="+sectionType)...)
					outbytes = append(outbytes, imageClose...)
					i += urlLen
					lastItem = i
					continue
				} else if media.Type == "image" {
					outbytes = append(outbytes, imageOpen...)
					outbytes = append(outbytes, []byte(media.URL)...)
					outbytes = append(outbytes, imageOpen2...)
					outbytes = append(outbytes, []byte(media.URL)...)
					outbytes = append(outbytes, imageClose...)
					i += urlLen
					lastItem = i
					continue
				} else if media.Type == "raw" {
					outbytes = append(outbytes, []byte(media.Body)...)
					i += urlLen
					lastItem = i
					continue
				} else if media.Type != "" {
					outbytes = append(outbytes, unknownMedia...)
					i += urlLen
					continue
				}

				outbytes = append(outbytes, urlOpen...)
				outbytes = append(outbytes, msgbytes[i:i+urlLen]...)
				outbytes = append(outbytes, urlOpen2...)
				outbytes = append(outbytes, msgbytes[i:i+urlLen]...)
				outbytes = append(outbytes, urlClose...)
				i += urlLen
				lastItem = i
			}
		}
	}

	if lastItem != i && len(outbytes) != 0 {
		//log.Print("lastItem: ", msgbytes[lastItem])
		//log.Print("lastItem index: ", lastItem)
		//log.Print("i: ", i)
		//log.Print("lastItem to end: ", msgbytes[lastItem:])
		//log.Print("-----")
		calclen := len(msgbytes) - 10
		if calclen <= lastItem {
			calclen = lastItem
		}
		outbytes = append(outbytes, msgbytes[lastItem:calclen]...)
		msg = string(outbytes)
	}
	//log.Print(`"`+string(outbytes)+`"`)
	//log.Print("msg",`"`+msg+`"`)

	msg = strings.Replace(msg, "\n", "<br>", -1)
	if sshooks["parse_assign"] != nil {
		msg = runSshook("parse_assign", msg)
	}
	return msg
}

// TODO: Write a test for this
/*func regexParseMessage(msg string) string {
	msg = strings.Replace(msg, ":)", "😀", -1)
	msg = strings.Replace(msg, ":D", "😃", -1)
	msg = strings.Replace(msg, ":P", "😛", -1)
	msg = urlReg.ReplaceAllString(msg, "<a href=\"$2$3//$4\" rel=\"nofollow\">$2$3//$4</a>")
	msg = strings.Replace(msg, "\n", "<br>", -1)
	if sshooks["parse_assign"] != nil {
		msg = runSshook("parse_assign", msg)
	}
	return msg
}*/

// 6, 7, 8, 6, 2, 7
// ftp://, http://, https:// git://, //, mailto: (not a URL, just here for length comparison purposes)
// TODO: Write a test for this
func validateURLBytes(data []byte) bool {
	datalen := len(data)
	i := 0

	if datalen >= 6 {
		if bytes.Equal(data[0:6], []byte("ftp://")) || bytes.Equal(data[0:6], []byte("git://")) {
			i = 6
		} else if datalen >= 7 && bytes.Equal(data[0:7], httpProtBytes) {
			i = 7
		} else if datalen >= 8 && bytes.Equal(data[0:8], []byte("https://")) {
			i = 8
		}
	} else if datalen >= 2 && data[0] == '/' && data[1] == '/' {
		i = 2
	}

	// ? - There should only be one : and that's only if the URL is on a non-standard port. Same for ?s.
	for ; datalen > i; i++ {
		if data[i] != '\\' && data[i] != '_' && data[i] != ':' && data[i] != '?' && data[i] != '&' && data[i] != '=' && data[i] != ';' && data[i] != '@' && !(data[i] > 44 && data[i] < 58) && !(data[i] > 64 && data[i] < 91) && !(data[i] > 96 && data[i] < 123) {
			return false
		}
	}
	return true
}

// TODO: Write a test for this
func validatedURLBytes(data []byte) (url []byte) {
	datalen := len(data)
	i := 0

	if datalen >= 6 {
		if bytes.Equal(data[0:6], []byte("ftp://")) || bytes.Equal(data[0:6], []byte("git://")) {
			i = 6
		} else if datalen >= 7 && bytes.Equal(data[0:7], httpProtBytes) {
			i = 7
		} else if datalen >= 8 && bytes.Equal(data[0:8], []byte("https://")) {
			i = 8
		}
	} else if datalen >= 2 && data[0] == '/' && data[1] == '/' {
		i = 2
	}

	// ? - There should only be one : and that's only if the URL is on a non-standard port. Same for ?s.
	for ; datalen > i; i++ {
		if data[i] != '\\' && data[i] != '_' && data[i] != ':' && data[i] != '?' && data[i] != '&' && data[i] != '=' && data[i] != ';' && data[i] != '@' && !(data[i] > 44 && data[i] < 58) && !(data[i] > 64 && data[i] < 91) && !(data[i] > 96 && data[i] < 123) {
			return invalidURL
		}
	}

	url = append(url, data...)
	return url
}

// TODO: Write a test for this
func partialURLBytes(data []byte) (url []byte) {
	datalen := len(data)
	i := 0
	end := datalen - 1

	if datalen >= 6 {
		if bytes.Equal(data[0:6], []byte("ftp://")) || bytes.Equal(data[0:6], []byte("git://")) {
			i = 6
		} else if datalen >= 7 && bytes.Equal(data[0:7], httpProtBytes) {
			i = 7
		} else if datalen >= 8 && bytes.Equal(data[0:8], []byte("https://")) {
			i = 8
		}
	} else if datalen >= 2 && data[0] == '/' && data[1] == '/' {
		i = 2
	}

	// ? - There should only be one : and that's only if the URL is on a non-standard port. Same for ?s.
	for ; end >= i; i++ {
		if data[i] != '\\' && data[i] != '_' && data[i] != ':' && data[i] != '?' && data[i] != '&' && data[i] != '=' && data[i] != ';' && data[i] != '@' && !(data[i] > 44 && data[i] < 58) && !(data[i] > 64 && data[i] < 91) && !(data[i] > 96 && data[i] < 123) {
			end = i
		}
	}

	url = append(url, data[0:end]...)
	return url
}

// TODO: Write a test for this
func partialURLBytesLen(data []byte) int {
	datalen := len(data)
	i := 0

	if datalen >= 6 {
		//log.Print(string(data[0:5]))
		if bytes.Equal(data[0:6], []byte("ftp://")) || bytes.Equal(data[0:6], []byte("git://")) {
			i = 6
		} else if datalen >= 7 && bytes.Equal(data[0:7], httpProtBytes) {
			i = 7
		} else if datalen >= 8 && bytes.Equal(data[0:8], []byte("https://")) {
			i = 8
		}
	} else if datalen >= 2 && data[0] == '/' && data[1] == '/' {
		i = 2
	}

	// ? - There should only be one : and that's only if the URL is on a non-standard port. Same for ?s.
	for ; datalen > i; i++ {
		if data[i] != '\\' && data[i] != '_' && data[i] != ':' && data[i] != '?' && data[i] != '&' && data[i] != '=' && data[i] != ';' && data[i] != '@' && !(data[i] > 44 && data[i] < 58) && !(data[i] > 64 && data[i] < 91) && !(data[i] > 96 && data[i] < 123) {
			//log.Print("Bad Character: ", data[i])
			return i
		}
	}
	//log.Print("Data Length: ",datalen)
	return datalen
}

type MediaEmbed struct {
	Type string //image
	URL  string
	Body string
}

// TODO: Write a test for this
func parseMediaBytes(data []byte) (media MediaEmbed, ok bool) {
	if !validateURLBytes(data) {
		return media, false
	}
	url, err := parseURL(data)
	if err != nil {
		return media, false
	}

	//log.Print("url ", url)
	hostname := url.Hostname()
	scheme := url.Scheme
	port := url.Port()
	//log.Print("hostname ", hostname)
	//log.Print("scheme ", scheme)
	query := url.Query()
	//log.Printf("query %+v\n", query)

	var samesite = hostname == "localhost" || hostname == site.URL
	if samesite {
		//log.Print("samesite")
		hostname = strings.Split(site.URL, ":")[0]
		// ?- Test this as I'm not sure it'll do what it should. If someone's running SSL on port 80 or non-SSL on port 443 then... Well... They're in far worse trouble than this...
		port = site.Port
		if scheme == "" && site.EnableSsl {
			scheme = "https"
		}
	}
	if scheme == "" {
		scheme = "http"
	}

	path := url.EscapedPath()
	//log.Print("path", path)
	pathFrags := strings.Split(path, "/")
	//log.Printf("pathFrags %+v\n", pathFrags)
	//log.Print("scheme ", scheme)
	//log.Print("hostname ", hostname)
	if len(pathFrags) >= 2 {
		if samesite && pathFrags[1] == "attachs" && (scheme == "http" || scheme == "https") {
			//log.Print("Attachment")
			media.Type = "attach"
			var sport string
			// ? - Assumes the sysadmin hasn't mixed up the two standard ports
			if port != "443" && port != "80" {
				sport = ":" + port
			}
			media.URL = scheme + "://" + hostname + sport + path
			return media, true
		}
	}

	// ? - I don't think this hostname will hit every YT domain
	// TODO: Make this a more customisable handler rather than hard-coding it in here
	if hostname == "www.youtube.com" && path == "/watch" {
		video, ok := query["v"]
		if ok && len(video) >= 1 && video[0] != "" {
			media.Type = "raw"
			// TODO: Filter the URL to make sure no nasties end up in there
			media.Body = "<iframe class='postIframe' src='https://www.youtube-nocookie.com/embed/" + video[0] + "' frameborder='0' allowfullscreen></iframe>"
			return media, true
		}
	}

	lastFrag := pathFrags[len(pathFrags)-1]
	if lastFrag != "" {
		// TODO: Write a function for getting the file extension of a string
		extarr := strings.Split(lastFrag, ".")
		if len(extarr) >= 2 {
			ext := extarr[len(extarr)-1]
			if imageFileExts.Contains(ext) {
				media.Type = "image"
				var sport string
				if port != "443" && port != "80" {
					sport = ":" + port
				}
				media.URL = scheme + "://" + hostname + sport + path
				return media, true
			}
		}
	}

	return media, true
}

func parseURL(data []byte) (*url.URL, error) {
	return url.Parse(string(data))
}

// TODO: Write a test for this
func coerceIntBytes(data []byte) (res int, length int) {
	if !(data[0] > 47 && data[0] < 58) {
		return 0, 1
	}

	i := 0
	for ; len(data) > i; i++ {
		if !(data[i] > 47 && data[i] < 58) {
			conv, err := strconv.Atoi(string(data[0:i]))
			if err != nil {
				return 0, i
			}
			return conv, i
		}
	}

	conv, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, i
	}
	return conv, i
}

// TODO: Write tests for this
func paginate(count int, perPage int, maxPages int) []int {
	if count < perPage {
		return []int{1}
	}
	var page int
	var out []int
	for current := 0; current < count; current += perPage {
		page++
		out = append(out, page)
		if len(out) >= maxPages {
			break
		}
	}
	return out
}

// TODO: Write tests for this
func pageOffset(count int, page int, perPage int) (int, int, int) {
	var offset int
	lastPage := (count / perPage) + 1
	if page > 1 {
		offset = (perPage * page) - perPage
	} else if page == -1 {
		page = lastPage
		offset = (perPage * page) - perPage
	} else {
		page = 1
	}

	// We don't want the offset to overflow the slices, if everything's in memory
	if offset >= (count - 1) {
		offset = 0
	}
	return offset, page, lastPage
}
