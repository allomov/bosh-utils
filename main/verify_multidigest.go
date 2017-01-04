package main

import (
	"github.com/jessevdk/go-flags"
	"os"
	boshcrypto "github.com/cloudfoundry/bosh-utils/crypto"
)

type opts struct {
	VerifyMultiDigestCommand MultiDigestCommand `command:"verify-multi-digest"`
}

func main() {
	o := opts{}
	_, err := flags.Parse(&o)

	if err != nil {
		os.Exit(1)
	}
}

type MultiDigestArgs struct {
	File string
	Digest string
}

type MultiDigestCommand struct {
	Args MultiDigestArgs  `positional-args:"yes"`

}

func (m MultiDigestCommand) Execute(args []string) error {
	multipleDigest := boshcrypto.MustParseMultipleDigest(m.Args.Digest)
	file, err := os.Open(m.Args.File)
	if err != nil {
		return err
	}
	return multipleDigest.Verify(file)
}
