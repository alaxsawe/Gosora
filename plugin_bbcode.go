package main

import (
	//"log"
	//"fmt"
	"bytes"

	//"strings"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

var random *rand.Rand
var bbcodeInvalidNumber []byte
var bbcodeNoNegative []byte
var bbcodeMissingTag []byte

var bbcodeBold *regexp.Regexp
var bbcodeItalic *regexp.Regexp
var bbcodeUnderline *regexp.Regexp
var bbcodeStrikethrough *regexp.Regexp
var bbcodeURL *regexp.Regexp
var bbcodeURLLabel *regexp.Regexp
var bbcodeQuotes *regexp.Regexp
var bbcodeCode *regexp.Regexp

func init() {
	plugins["bbcode"] = NewPlugin("bbcode", "BBCode", "Azareal", "http://github.com/Azareal", "", "", "", initBbcode, nil, deactivateBbcode, nil, nil)
}

func initBbcode() error {
	//plugins["bbcode"].AddHook("parse_assign", bbcode_parse_without_code)
	plugins["bbcode"].AddHook("parse_assign", bbcodeFullParse)

	bbcodeInvalidNumber = []byte("<span style='color: red;'>[Invalid Number]</span>")
	bbcodeNoNegative = []byte("<span style='color: red;'>[No Negative Numbers]</span>")
	bbcodeMissingTag = []byte("<span style='color: red;'>[Missing Tag]</span>")

	bbcodeBold = regexp.MustCompile(`(?s)\[b\](.*)\[/b\]`)
	bbcodeItalic = regexp.MustCompile(`(?s)\[i\](.*)\[/i\]`)
	bbcodeUnderline = regexp.MustCompile(`(?s)\[u\](.*)\[/u\]`)
	bbcodeStrikethrough = regexp.MustCompile(`(?s)\[s\](.*)\[/s\]`)
	urlpattern := `(http|https|ftp|mailto*)(:??)\/\/([\.a-zA-Z\/]+)`
	bbcodeURL = regexp.MustCompile(`\[url\]` + urlpattern + `\[/url\]`)
	bbcodeURLLabel = regexp.MustCompile(`(?s)\[url=` + urlpattern + `\](.*)\[/url\]`)
	bbcodeQuotes = regexp.MustCompile(`\[quote\](.*)\[/quote\]`)
	bbcodeCode = regexp.MustCompile(`\[code\](.*)\[/code\]`)

	random = rand.New(rand.NewSource(time.Now().UnixNano()))
	return nil
}

func deactivateBbcode() {
	//plugins["bbcode"].RemoveHook("parse_assign", bbcode_parse_without_code)
	plugins["bbcode"].RemoveHook("parse_assign", bbcodeFullParse)
}

func bbcodeRegexParse(msg string) string {
	msg = bbcodeBold.ReplaceAllString(msg, "<b>$1</b>")
	msg = bbcodeItalic.ReplaceAllString(msg, "<i>$1</i>")
	msg = bbcodeUnderline.ReplaceAllString(msg, "<u>$1</u>")
	msg = bbcodeStrikethrough.ReplaceAllString(msg, "<s>$1</s>")
	msg = bbcodeURL.ReplaceAllString(msg, "<a href=''$1$2//$3' rel='nofollow'>$1$2//$3</i>")
	msg = bbcodeURLLabel.ReplaceAllString(msg, "<a href=''$1$2//$3' rel='nofollow'>$4</i>")
	msg = bbcodeQuotes.ReplaceAllString(msg, "<span class='postQuote'>$1</span>")
	//msg = bbcodeCode.ReplaceAllString(msg,"<span class='codequotes'>$1</span>")
	return msg
}

// Only does the simple BBCode like [u], [b], [i] and [s]
func bbcodeSimpleParse(msg string) string {
	var hasU, hasB, hasI, hasS bool
	msgbytes := []byte(msg)
	for i := 0; (i + 2) < len(msgbytes); i++ {
		if msgbytes[i] == '[' && msgbytes[i+2] == ']' {
			if msgbytes[i+1] == 'b' && !hasB {
				msgbytes[i] = '<'
				msgbytes[i+2] = '>'
				hasB = true
			} else if msgbytes[i+1] == 'i' && !hasI {
				msgbytes[i] = '<'
				msgbytes[i+2] = '>'
				hasI = true
			} else if msgbytes[i+1] == 'u' && !hasU {
				msgbytes[i] = '<'
				msgbytes[i+2] = '>'
				hasU = true
			} else if msgbytes[i+1] == 's' && !hasS {
				msgbytes[i] = '<'
				msgbytes[i+2] = '>'
				hasS = true
			}
			i += 2
		}
	}

	// There's an unclosed tag in there x.x
	if hasI || hasU || hasB || hasS {
		closeUnder := []byte("</u>")
		closeItalic := []byte("</i>")
		closeBold := []byte("</b>")
		closeStrike := []byte("</s>")
		if hasI {
			msgbytes = append(msgbytes, closeItalic...)
		}
		if hasU {
			msgbytes = append(msgbytes, closeUnder...)
		}
		if hasB {
			msgbytes = append(msgbytes, closeBold...)
		}
		if hasS {
			msgbytes = append(msgbytes, closeStrike...)
		}
	}
	return string(msgbytes)
}

// Here for benchmarking purposes. Might add a plugin setting for disabling [code] as it has it's paws everywhere
func bbcodeParseWithoutCode(msg string) string {
	var hasU, hasB, hasI, hasS bool
	var complexBbc bool
	msgbytes := []byte(msg)

	for i := 0; (i + 3) < len(msgbytes); i++ {
		if msgbytes[i] == '[' {
			if msgbytes[i+2] != ']' {
				if msgbytes[i+1] == '/' {
					if msgbytes[i+3] == ']' {
						if msgbytes[i+2] == 'b' {
							msgbytes[i] = '<'
							msgbytes[i+3] = '>'
							hasB = false
						} else if msgbytes[i+2] == 'i' {
							msgbytes[i] = '<'
							msgbytes[i+3] = '>'
							hasI = false
						} else if msgbytes[i+2] == 'u' {
							msgbytes[i] = '<'
							msgbytes[i+3] = '>'
							hasU = false
						} else if msgbytes[i+2] == 's' {
							msgbytes[i] = '<'
							msgbytes[i+3] = '>'
							hasS = false
						}
						i += 3
					} else {
						complexBbc = true
					}
				} else {
					complexBbc = true
				}
			} else {
				if msgbytes[i+1] == 'b' && !hasB {
					msgbytes[i] = '<'
					msgbytes[i+2] = '>'
					hasB = true
				} else if msgbytes[i+1] == 'i' && !hasI {
					msgbytes[i] = '<'
					msgbytes[i+2] = '>'
					hasI = true
				} else if msgbytes[i+1] == 'u' && !hasU {
					msgbytes[i] = '<'
					msgbytes[i+2] = '>'
					hasU = true
				} else if msgbytes[i+1] == 's' && !hasS {
					msgbytes[i] = '<'
					msgbytes[i+2] = '>'
					hasS = true
				}
				i += 2
			}
		}
	}

	// There's an unclosed tag in there x.x
	if hasI || hasU || hasB || hasS {
		closeUnder := []byte("</u>")
		closeItalic := []byte("</i>")
		closeBold := []byte("</b>")
		closeStrike := []byte("</s>")
		if hasI {
			msgbytes = append(bytes.TrimSpace(msgbytes), closeItalic...)
		}
		if hasU {
			msgbytes = append(bytes.TrimSpace(msgbytes), closeUnder...)
		}
		if hasB {
			msgbytes = append(bytes.TrimSpace(msgbytes), closeBold...)
		}
		if hasS {
			msgbytes = append(bytes.TrimSpace(msgbytes), closeStrike...)
		}
	}

	// Copy the new complex parser over once the rough edges have been smoothed over
	if complexBbc {
		msg = string(msgbytes)
		msg = bbcodeURL.ReplaceAllString(msg, "<a href='$1$2//$3' rel='nofollow'>$1$2//$3</i>")
		msg = bbcodeURLLabel.ReplaceAllString(msg, "<a href='$1$2//$3' rel='nofollow'>$4</i>")
		msg = bbcodeQuotes.ReplaceAllString(msg, "<span class='postQuote'>$1</span>")
		return bbcodeCode.ReplaceAllString(msg, "<span class='codequotes'>$1</span>")
	}
	return string(msgbytes)
}

// Does every type of BBCode
func bbcodeFullParse(msg string) string {
	var hasU, hasB, hasI, hasS, hasC bool
	var complexBbc bool

	msgbytes := []byte(msg)
	msgbytes = append(msgbytes, spaceGap...)
	//log.Print("BBCode Simple Pre:","`"+string(msgbytes)+"`")
	//log.Print("----")

	for i := 0; i < len(msgbytes); i++ {
		if msgbytes[i] == '[' {
			if msgbytes[i+2] != ']' {
				if msgbytes[i+1] == '/' {
					if msgbytes[i+3] == ']' {
						if !hasC {
							if msgbytes[i+2] == 'b' {
								msgbytes[i] = '<'
								msgbytes[i+3] = '>'
								hasB = false
							} else if msgbytes[i+2] == 'i' {
								msgbytes[i] = '<'
								msgbytes[i+3] = '>'
								hasI = false
							} else if msgbytes[i+2] == 'u' {
								msgbytes[i] = '<'
								msgbytes[i+3] = '>'
								hasU = false
							} else if msgbytes[i+2] == 's' {
								msgbytes[i] = '<'
								msgbytes[i+3] = '>'
								hasS = false
							}
							i += 3
						}
					} else {
						if msgbytes[i+2] == 'c' && msgbytes[i+3] == 'o' && msgbytes[i+4] == 'd' && msgbytes[i+5] == 'e' && msgbytes[i+6] == ']' {
							hasC = false
							i += 7
						}
						//if msglen >= (i+6) {
						//	log.Print("boo")
						//	log.Print(msglen)
						//	log.Print(i+6)
						//	log.Print(string(msgbytes[i:i+6]))
						//}
						complexBbc = true
					}
				} else {
					if msgbytes[i+1] == 'c' && msgbytes[i+2] == 'o' && msgbytes[i+3] == 'd' && msgbytes[i+4] == 'e' && msgbytes[i+5] == ']' {
						hasC = true
						i += 6
					}
					//if msglen >= (i+5) {
					//	log.Print("boo2")
					//	log.Print(string(msgbytes[i:i+5]))
					//}
					complexBbc = true
				}
			} else if !hasC {
				if msgbytes[i+1] == 'b' && !hasB {
					msgbytes[i] = '<'
					msgbytes[i+2] = '>'
					hasB = true
				} else if msgbytes[i+1] == 'i' && !hasI {
					msgbytes[i] = '<'
					msgbytes[i+2] = '>'
					hasI = true
				} else if msgbytes[i+1] == 'u' && !hasU {
					msgbytes[i] = '<'
					msgbytes[i+2] = '>'
					hasU = true
				} else if msgbytes[i+1] == 's' && !hasS {
					msgbytes[i] = '<'
					msgbytes[i+2] = '>'
					hasS = true
				}
				i += 2
			}
		}
	}

	// There's an unclosed tag in there somewhere x.x
	if hasI || hasU || hasB || hasS {
		closeUnder := []byte("</u>")
		closeItalic := []byte("</i>")
		closeBold := []byte("</b>")
		closeStrike := []byte("</s>")
		if hasI {
			msgbytes = append(bytes.TrimSpace(msgbytes), closeItalic...)
		}
		if hasU {
			msgbytes = append(bytes.TrimSpace(msgbytes), closeUnder...)
		}
		if hasB {
			msgbytes = append(bytes.TrimSpace(msgbytes), closeBold...)
		}
		if hasS {
			msgbytes = append(bytes.TrimSpace(msgbytes), closeStrike...)
		}
		msgbytes = append(msgbytes, spaceGap...)
	}

	if complexBbc {
		i := 0
		var start, lastTag int
		var outbytes []byte
		//log.Print("BBCode Pre:","`"+string(msgbytes)+"`")
		//log.Print("----")
		for ; i < len(msgbytes); i++ {
		MainLoop:
			if msgbytes[i] == '[' {
			OuterComplex:
				if msgbytes[i+1] == 'u' {
					if msgbytes[i+2] == 'r' && msgbytes[i+3] == 'l' && msgbytes[i+4] == ']' {
						start = i + 5
						outbytes = append(outbytes, msgbytes[lastTag:i]...)
						i = start
						i += partialURLBytesLen(msgbytes[start:])
						//log.Print("Partial Bytes:",string(msgbytes[start:]))
						//log.Print("-----")
						if !bytes.Equal(msgbytes[i:i+6], []byte("[/url]")) {
							//log.Print("Invalid Bytes:",string(msgbytes[i:i+6]))
							//log.Print("-----")
							outbytes = append(outbytes, invalidURL...)
							goto MainLoop
						}

						outbytes = append(outbytes, urlOpen...)
						outbytes = append(outbytes, msgbytes[start:i]...)
						outbytes = append(outbytes, urlOpen2...)
						outbytes = append(outbytes, msgbytes[start:i]...)
						outbytes = append(outbytes, urlClose...)
						i += 6
						lastTag = i
					}
				} else if msgbytes[i+1] == 'r' {
					if bytes.Equal(msgbytes[i+2:i+6], []byte("and]")) {
						outbytes = append(outbytes, msgbytes[lastTag:i]...)
						start = i + 6
						i = start
						for ; ; i++ {
							if msgbytes[i] == '[' {
								if !bytes.Equal(msgbytes[i+1:i+7], []byte("/rand]")) {
									outbytes = append(outbytes, bbcodeMissingTag...)
									goto OuterComplex
								}
								break
							} else if (len(msgbytes) - 1) < (i + 10) {
								outbytes = append(outbytes, bbcodeMissingTag...)
								goto OuterComplex
							}
						}

						number, err := strconv.ParseInt(string(msgbytes[start:i]), 10, 64)
						if err != nil {
							outbytes = append(outbytes, bbcodeInvalidNumber...)
							goto MainLoop
						}

						// TODO: Add support for negative numbers?
						if number < 0 {
							outbytes = append(outbytes, bbcodeNoNegative...)
							goto MainLoop
						}

						var dat []byte
						if number == 0 {
							dat = []byte("0")
						} else {
							dat = []byte(strconv.FormatInt((random.Int63n(number)), 10))
						}

						outbytes = append(outbytes, dat...)
						//log.Print("Outputted the random number")
						i += 7
						lastTag = i
					}
				}
			}
		}
		//log.Print("Outbytes:",`"`+string(outbytes)+`"`)
		if lastTag != i {
			outbytes = append(outbytes, msgbytes[lastTag:]...)
		}

		if len(outbytes) != 0 {
			//log.Print("BBCode Post:",`"`+string(outbytes[0:len(outbytes) - 10])+`"`)
			msg = string(outbytes[0 : len(outbytes)-10])
		} else {
			msg = string(msgbytes[0 : len(msgbytes)-10])
		}
		//log.Print("----")

		//msg = bbcode_url.ReplaceAllString(msg,"<a href=\"$1$2//$3\" rel=\"nofollow\">$1$2//$3</i>")
		msg = bbcodeURLLabel.ReplaceAllString(msg, "<a href='$1$2//$3' rel='nofollow'>$4</i>")
		msg = bbcodeQuotes.ReplaceAllString(msg, "<span class='postQuote'>$1</span>")
		msg = bbcodeCode.ReplaceAllString(msg, "<span class='codequotes'>$1</span>")
	} else {
		msg = string(msgbytes[0 : len(msgbytes)-10])
	}

	return msg
}
