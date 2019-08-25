package goconfig

import (
	"time"

	"math/rand"

	"github.com/sirupsen/logrus"
)

var (
	// TimeFormat configuration for date/time format outputting.
	TimeFormat = "2006-01-02 15:04:05"
	// TimeZone configuration for timezone.
	TimeZone, _ = time.LoadLocation("Local")

	// LetterSet character set for random string functions contaning letters.
	LetterSet = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// NumberSet character set for random string functions contaning Numbers.
	NumberSet = "0123456789"
	// AlphaNumericSet character set for random string functions contaning basic alphanumberic characters.
	AlphaNumericSet = LetterSet + NumberSet
	// RandomSeed create a random seed for random functions
	RandomSeed = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Log sets the default logging facility for output.
	Log = logrus.New()

	// Dynamic configuration holder for miscellaneous package configuration.
	Dynamic map[string]interface{}
)
