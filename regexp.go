package marker

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"regexp"
)

const (
	weekdaysAbv   = "((Mon)|(Tue)|(Wed)|(Thu)|(Fri)|(Sat)|(Sun))"
	weekdays      = "((Monday)|(Tuesday)|(Wednesday)|(Thursday)|(Friday)|(Saturday)|(Sunday))"
	monthsAbv     = "((Jan)|(Feb)|(Mar)|(Apr)|(May)|(Jun)|(Jul)|(Aug)|(Sep)|(Oct)|(Nov)|(Dec))"
	numericmonths = "(0[1-9]|1[0-2])"
	days          = "((_[1-9])|([1-2][0-9])|3[01])"               // _1 - 31
	daysWithZero  = "((0[1-9])|([1-2][0-9])|3[01])"               // 01 - 31
	hhmmss        = "(([0-1][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9])" // 22:15:02
	hhmm          = "(([0-1][0-9]|2[0-3]):[0-5][0-9])"            // 22:15
	year          = "[0-9]{4}"                                    // 2006
	timezone      = "[A-Z]{3}"                                    // MST
	numericZone   = "-[0-9]{4}"                                   // -0300
	milli         = ".[0-9]{3}"
	micro         = ".[0-9]{6}"
	nano          = ".[0-9]{9}"
)

var (
	ANSICRegexp       = regexp2.MustCompile(fmt.Sprintf("%s %s %s %s %s", weekdaysAbv, monthsAbv, days, hhmmss, year), 0)
	UnixDateRegexp    = regexp2.MustCompile(fmt.Sprintf("%s %s %s %s %s %s", weekdaysAbv, monthsAbv, days, hhmmss, timezone, year), 0)
	RubyDateRegexp    = regexp2.MustCompile(fmt.Sprintf("%s %s %s %s %s %s", weekdaysAbv, monthsAbv, daysWithZero, hhmmss, numericZone, year), 0)
	RFC822Regexp      = regexp2.MustCompile(fmt.Sprintf("%s %s [0-9]{2} %s %s", daysWithZero, monthsAbv, hhmm, timezone), 0)
	RFC822ZRegexp     = regexp2.MustCompile(fmt.Sprintf("%s %s [0-9]{2} %s %s", daysWithZero, monthsAbv, hhmm, numericZone), 0)
	RFC850Regexp      = regexp2.MustCompile(fmt.Sprintf("%s, %s-%s-[0-9]{2} %s %s", weekdays, daysWithZero, monthsAbv, hhmmss, timezone), 0)
	RFC1123Regexp     = regexp2.MustCompile(fmt.Sprintf("%s, %s %s %s %s %s", weekdaysAbv, daysWithZero, monthsAbv, year, hhmmss, timezone), 0)
	RFC1123ZRegexp    = regexp2.MustCompile(fmt.Sprintf("%s, %s %s %s %s %s", weekdaysAbv, daysWithZero, monthsAbv, year, hhmmss, numericZone), 0)
	RFC3339Regexp     = regexp2.MustCompile(fmt.Sprintf("%s-%s-%sT%sZ%s", year, numericmonths, daysWithZero, hhmmss, hhmm), 0)
	RFC3339NanoRegexp = regexp2.MustCompile(fmt.Sprintf("%s-%s-%sT%s%sZ%s", year, numericmonths, daysWithZero, hhmmss, nano, hhmm), 0)
	KitchenRegexp     = regexp2.MustCompile("(([0-1]?[0-9]|2[0-3]):[0-5][0-9][P|A]M)", 0)
	StampRegexp       = regexp2.MustCompile(fmt.Sprintf("%s %s %s(?!(\.[0-9]))", monthsAbv, days, hhmmss), 0)
	StampMilliRegexp  = regexp2.MustCompile(fmt.Sprintf("%s %s %s%s", monthsAbv, days, hhmmss, milli), 0)
	StampMicroRegexp  = regexp2.MustCompile(fmt.Sprintf("%s %s %s%s", monthsAbv, days, hhmmss, micro), 0)
	StampNanoRegexp   = regexp2.MustCompile(fmt.Sprintf("%s %s %s%s", monthsAbv, days, hhmmss, nano), 0)
)

// EmailRegexp is a Regular expression for RFC5322
var EmailRegexp = regexp.MustCompile(`[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])`)
