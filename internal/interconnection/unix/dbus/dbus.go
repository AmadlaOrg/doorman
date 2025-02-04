package dbus

import (
	"fmt"
	"github.com/AmadlaOrg/LibraryUtils/encryption"
	"github.com/godbus/dbus/v5"
	"log"
)

type IDBus interface{}

type SDBus struct{}

const secretKey = "32byte-long-secret-key!!!!!" // Ensure key matches Clerk-AWS

func (s *SDBus) Connect() {
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Fatal("Failed to connect to D-Bus:", err)
	}

	obj := conn.Object("com.clerk.aws", "/com/clerk/aws")
	var encryptedJWT string

	// Call GetJWT method
	err = obj.Call("com.clerk.aws.GetJWT", 0).Store(&encryptedJWT)
	if err != nil {
		log.Fatal("Error calling Clerk-AWS:", err)
	}

	// Decrypt received JWT
	decryptedJWT, err := encryption.Decrypt(encryptedJWT, secretKey)
	if err != nil {
		log.Fatal("Decryption error:", err)
	}

	fmt.Println("Retrieved JWT:", decryptedJWT)
}
