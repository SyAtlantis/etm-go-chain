package models

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"etm-go-chain/utils"
	"reflect"
)

type iSign interface {
	IsEmpty() bool
	Create([]utils.Keypair, Block) error
	GetBytes() ([]byte, error)
	GetHash() ([32]byte, error)
	GetId() (string, error)
	GetSignature(utils.Keypair) (string, error)
	VerifySignature(map[string]string) (bool, error)
	HasEnoughSigns() bool
}
type Sign struct {
	Id         string
	Height     int64
	Timestamp  int64
	Signatures map[string]string
}

func (s *Sign) IsEmpty() bool {
	return s == nil || reflect.DeepEqual(s, Sign{})
}

func (s *Sign) Create(keypairs []utils.Keypair, block Block) error {
	s.Id = block.Id
	s.Height = block.Height
	s.Timestamp = block.Timestamp

	sign := make(map[string]string)
	for _, v := range keypairs {
		publicKey := hex.EncodeToString(v.PublicKey)
		var err error
		if sign[publicKey], err = s.GetSignature(v); err != nil {
			return err
		}
	}
	s.Signatures = sign

	return nil
}

func (s *Sign) GetBytes() ([]byte, error) {
	bb := bytes.NewBuffer([]byte{})
	err := binary.Write(bb, binary.LittleEndian, uint64(s.Height));
	bb.WriteString(s.Id)
	return bb.Bytes(), err
}

func (s *Sign) GetHash() ([32]byte, error) {
	bs, err := s.GetBytes()
	hash := sha256.Sum256(bs)
	return hash, err
}

func (s *Sign) GetId() (string, error) {
	panic("implement me")
}

func (s *Sign) GetSignature(keypair utils.Keypair) (string, error) {
	hash, err := s.GetHash()
	sign := ed.Sign(hash[:], keypair)
	return hex.EncodeToString(sign), err
}

func (s *Sign) VerifySignature() (bool, error) {
	hash, err := s.GetHash()
	if err != nil {
		return false, err
	}

	sign := s.Signatures
	for k, v := range sign {
		signBytes, err := hex.DecodeString(v)
		if err != nil {
			return false, err
		}
		publicKey, err := hex.DecodeString(k)
		if err != nil {
			return false, err
		}
		ok := ed.Verify(hash[:], signBytes, publicKey)
		if !ok {
			return false, nil
		}
	}
	return true, nil
}

func (s *Sign) HasEnoughSigns() bool {
	slots := utils.NewSlots()
	return len(s.Signatures) > slots.Delegates*2/3
}
