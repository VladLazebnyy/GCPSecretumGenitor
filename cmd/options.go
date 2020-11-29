package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	keyName string

	genGPG       bool
	genSecretGPG bool

	exportGPG       bool
	exportSecretGPG bool

	importGPG string
)

func init() {

	var now time.Time
	now = time.Now()
	ranVal := now.Unix()
	defKeyName := fmt.Sprintf("%s%d", "generatedName", ranVal) // default key name generation
	flag.StringVar(&keyName, "keyname", defKeyName, "The name of gpg key that should b e generated.")
	flag.StringVar(&keyName, "k", defKeyName, "The name of gpg key that should b e generated.")

	// Generation gpg key additional information

	gpgGenCmd := flag.NewFlagSet("generate-key", flag.ContinueOnError)
	gpgGenEnable := gpgGenCmd.Bool("generate-key", true, "generate gpg key")
	genSecretGPG := gpgGenCmd.Bool("generate-armored-key", false, "generate armored gpg key. If this key specified the key `generate-key` ignored")

	flag.StringVar(&importGPG, "import-key", "", "import provided gpg key. Full path should be provided")
	flag.StringVar(&importGPG, "i", "", "import provided gpg key. Full path should be provided")
	flag.Parse()
}
