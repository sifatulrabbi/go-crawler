package sshmanager

import (
	"fmt"
	"os"
)

const (
	PATH_TO_SSH_FOLDER = "/projects/personal/tmp/.ssh" // just for testing
	// PATH_TO_SSH_FOLDER = "/.ssh"
	DEFAULT_PUB_KEY = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDDpNyaZUfu5MPvevyUZBWBBCJTWoLq4myEtjTFNsRLTn1QMiAoOag63QgOCQPhMxhjT7S/wzuwV4YOTjN4KYKw2JM6N/Lh8zY9SF7Otm86LH2IY1laNyGq9h4y2oXLkIwwRxD4ZduVX3+gO6TUFskL7Bx3V23k2IDJ1ReaU1eLF/Be3ozbVQBJE2Nq0xE8D9zeveH3F7k0KHSzLfXiDUPbF2nyBUOIwjQ39z3SMaIRhxqmkzTfxQ4JzOU6ofrBByZj80k/dWSdBdPum4Ko4bYbqKaPjFLbjEaTCc+WHbtgOe/x3+ZcR8x5GU5ZANpTcsB14ehFxLRGqesTNhOsh0IJ5OKnED002gkJn4GEPMv1XUkW0ePgIL/pVMXXZ45mCG1QWieXbfjznHFmQBxcO8XY7xr5VqOmwUK3jzhijzdPumromFLxC3GlKggBsumqRIj5I+93FRse/QFCPtfidrUE+k9UqXD/3cGLktLSkjWq+IbH+kLfJZJq9ULWqwry4SWquvJUgT38NYUC0FjA8ZDCeb1HIUK2U5yhGNSfG5T84cJvPmyp2Xge96YapV8SMh1RMZbUXS4NpAD7kKecL7cvv+EzjckpPmS5ItHg/6mdOOLRmHTb14Q62ARfzX3+HdTV5Jegw8osbyZeURfaoCFj9a3xLI6M/b03BVvatRDNbQ== sifatulrabbi"
)

func AddToAuthorizedKeysList() error {
	fileEmpty := false
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	filePath := fmt.Sprintf("%s/%s/authorized_keys", homeDir, PATH_TO_SSH_FOLDER)
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := make([]byte, 1024*1024)
	n, err := f.Read(buf)
	if err != nil {
		if err.Error() != "EOF" {
			fmt.Println("error while reading file content:", err)
			return err
		} else {
			fileEmpty = true
		}
	}
	fmt.Println(string(buf[:n]))

	if fileEmpty {
		writeableFile, err := os.OpenFile(filePath, os.O_WRONLY, 0600)
		if err != nil {
			fmt.Println("Error while opening file in writeable mode:", err)
			return err
		}
		defer writeableFile.Close()
		if _, err = writeableFile.Write([]byte(DEFAULT_PUB_KEY)); err != nil {
			fmt.Println("Error while writing to the file:", err)
			return err
		}
	} else {
		// TODO:
	}

	return nil
}
