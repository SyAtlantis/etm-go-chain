package models

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"etm-go-chain/utils"
	"fmt"
	"reflect"
)

type iSign interface {
	IsEmpty() bool
	Create([]utils.Keypair, Block) error
	GetBytes() ([]byte, error)
	GetHash() ([32]byte, error)
	GetId() (string, error)
	GetSignature(utils.Keypair) (string, error)
	VerifySignature() (bool error)
	HasEnoughSigns() bool
}
type Sign struct {
	Id         string
	Height     int64
	Timestamp  int64
	Signatures []map[string]string
}

func (s *Sign) IsEmpty() bool {
	return s == nil || reflect.DeepEqual(s, Sign{})
}

func (s *Sign) Create(keypairs []utils.Keypair, block Block) error {
	s.Id = block.Id
	s.Height = block.Height
	s.Timestamp = block.Timestamp

	for _, v := range keypairs {
		sign := make(map[string]string)
		publicKey := fmt.Sprintf("%x", v.PublicKey)
		var err error
		if sign[publicKey], err = s.GetSignature(v); err != nil {
			return err
		}
		s.Signatures = append(s.Signatures, sign)
	}

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
	return fmt.Sprintf("%x", sign), err
}

func (s *Sign) VerifySignature() (bool error) {
	panic("implement me")
}

func (s *Sign) HasEnoughSigns() bool {
	panic("implement me")
}
