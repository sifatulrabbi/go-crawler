package test

import (
	"testing"

	"crawler/internals/sshmanager"
)

func TestReadSSHKeys(t *testing.T) {
	if err := sshmanager.AddToAuthorizedKeysList(); err != nil {
		t.Error(err)
	}
}
