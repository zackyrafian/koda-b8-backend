package libs

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const ( 
  memory uint32 = 64 * 1024 
  interations uint32 = 3 
  parallelism uint8 = 2
  keyLength uint32 = 32 
  saltLength uint32 = 16
)

func HashPassword(password string) (string, error) { 
  salt := make([]byte, saltLength)
  if _, err := rand.Read(salt); err != nil { 
    return "", err
  }

  hash := argon2.IDKey(
    []byte(password), 
    salt, 
    interations, 
    memory, 
    parallelism, 
    keyLength,
  )

  b64Salt := base64.RawStdEncoding.EncodeToString(salt)
  b64Hash := base64.RawStdEncoding.EncodeToString(hash)

  encoded := fmt.Sprintf(
    "$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
    memory, 
    interations,
    parallelism, 
    b64Salt,
    b64Hash,
  )

  return encoded, nil
}

func Verify(password, encodedHash string) (bool, error) { 
  parts := strings.Split(encodedHash, "$")
  if len(parts) != 6 { 
    return false, fmt.Errorf("invalid hash")
  }

  var mem uint32 
  var iter uint32
  var par uint8

  _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &mem, &iter, &par)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	other := argon2.IDKey(
		[]byte(password),
		salt,
		iter,
		mem,
		par,
		uint32(len(hash)),
	)
	return subtle.ConstantTimeCompare(hash, other) == 1, nil
}