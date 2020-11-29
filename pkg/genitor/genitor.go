package genitor

import (
	"fmt"
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

// generation of the gpg key should satisfy very specific case:
// Case description
func generation() (*openpgp.Entity, error) {
	var newKey *openpgp.Entity
	// func NewEntity(name, comment, email string, config *packet.Config) (*Entity, error)
	// check if name flag was provided
	// check comment
	// check email
	// check config https://github.com/golang/crypto/blob/master/openpgp/packet/config.go
	newKey, err := openpgp.NewEntity("GenitorTimestamp", "auotgenerated keypair", "genitor@dummy.com", nil)
	if err != nil {
		// change to log
		fmt.Println(err)
		return nil, err
	}

	// Sign all the identities
	for _, id := range newKey.Identities {
		err := id.SelfSignature.SignUserId(id.UserId.Id, newKey.PrimaryKey, newKey.PrivateKey, nil)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	w, err := armor.Encode(os.Stdout, openpgp.PublicKeyType, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer w.Close()

	newKey.Serialize(w)

	return newKey, nil

}
