package cache

import (
	"crypto/rand"
	"log"
	"os"
	"runtime"
	"time"
	"unsafe"

	"github.com/dgraph-io/ristretto"
	"golang.org/x/sys/windows"
)

type ICache interface{}
type SCache struct {
	cache *ristretto.Cache
}

func (s *SCache) Init() {

}

// Encrypt JWT (Linux: TPM, Windows: DPAPI)
func encryptJWT(jwt string) ([]byte, error) {
	if runtime.GOOS == "windows" {
		return encryptWithDPAPI(jwt)
	}
	return encryptWithTPM(jwt)
}

// Decrypt JWT
func decryptJWT(encrypted []byte) (string, error) {
	if runtime.GOOS == "windows" {
		return decryptWithDPAPI(encrypted)
	}
	return decryptWithTPM(encrypted)
}

// =======================
// Linux TPM Encryption
// =======================
func encryptWithTPM(jwt string) ([]byte, error) {
	tpmDevice, err := os.Open("/dev/tpm0")
	if err != nil {
		return nil, err
	}
	defer func(tpmDevice *os.File) {
		err := tpmDevice.Close()
		if err != nil {
			// TODO:
		}
	}(tpmDevice)

	// Generate a TPM-based key (mocked here)
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		return nil, err
	}

	// Encrypt using XOR (Replace with AES-GCM in production)
	encrypted := make([]byte, len(jwt))
	for i := range jwt {
		encrypted[i] = jwt[i] ^ key[i%len(key)]
	}
	return encrypted, nil
}

// Decrypt with TPM
func decryptWithTPM(encrypted []byte) (string, error) {
	// Use same TPM-based key
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	decrypted := make([]byte, len(encrypted))
	for i := range encrypted {
		decrypted[i] = encrypted[i] ^ key[i%len(key)]
	}
	return string(decrypted), nil
}

// =======================
// Windows DPAPI Encryption
// =======================
func encryptWithDPAPI(jwt string) ([]byte, error) {
	data := windows.DataBlob{Size: uint32(len(jwt)), Data: (*byte)(unsafe.Pointer(&[]byte(jwt)[0]))}

	var out windows.DataBlob
	err := windows.CryptProtectData(&data, nil, nil, nil, nil, windows.CRYPTPROTECT_LOCAL_MACHINE, &out)
	if err != nil {
		return nil, err
	}

	return unsafe.Slice(out.Data, out.Size), nil
}

// Decrypt with DPAPI
func decryptWithDPAPI(encrypted []byte) (string, error) {
	data := windows.DataBlob{Size: uint32(len(encrypted)), Data: (*byte)(unsafe.Pointer(&encrypted[0]))}

	var out windows.DataBlob
	err := windows.CryptUnprotectData(&data, nil, nil, nil, nil, 0, &out)
	if err != nil {
		return "", err
	}

	return string(unsafe.Slice(out.Data, out.Size)), nil
}

// =======================
// JWT Storage in Memory
// =======================
func storeJWT(jwt string) {
	encrypted, err := encryptJWT(jwt)
	if err != nil {
		log.Fatal("Encryption failed:", err)
	}
	s.cache.SetWithTTL("jwt", encrypted, 1, 10*time.Minute) // Auto expires
}

func getJWT() string {
	value, found := s.cache.Get("jwt")
	if !found {
		return "JWT expired or not found"
	}

	encrypted := value.([]byte)
	jwt, err := decryptJWT(encrypted)
	if err != nil {
		log.Fatal("Decryption failed:", err)
	}

	return jwt
}
