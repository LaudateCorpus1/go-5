//lint:file-ignore S1005 The issue should be fixed in xdrgen. Unfortunately, there's no way to ignore a single file in staticcheck.
//lint:file-ignore U1000 fmtTest is not needed anywhere, should be removed in xdrgen.
// Package xdr is generated from:
//
//  Stellar-SCP.x
//  Stellar-ledger-entries.x
//  Stellar-ledger.x
//  Stellar-overlay.x
//  Stellar-transaction.x
//  Stellar-types.x
//
// DO NOT EDIT or your changes may be overwritten
package xdr

import (
	"bytes"
	"encoding"
	"fmt"
	"io"

	xdr "github.com/stellar/go-xdr/xdr3"
)

// Unmarshal reads an xdr element from `r` into `v`.
func Unmarshal(r io.Reader, v interface{}) (int, error) {
	// delegate to xdr package's Unmarshal
	return xdr.Unmarshal(r, v)
}

// Marshal writes an xdr element `v` into `w`.
func Marshal(w io.Writer, v interface{}) (int, error) {
	// delegate to xdr package's Marshal
	return xdr.Marshal(w, v)
}

// Value is an XDR Typedef defines as:
//
//   typedef opaque Value<>;
//
type Value []byte

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Value) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Value) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Value)(nil)
	_ encoding.BinaryUnmarshaler = (*Value)(nil)
)

// ScpBallot is an XDR Struct defines as:
//
//   struct SCPBallot
//    {
//        uint32 counter; // n
//        Value value;    // x
//    };
//
type ScpBallot struct {
	Counter Uint32
	Value   Value
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ScpBallot) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ScpBallot) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ScpBallot)(nil)
	_ encoding.BinaryUnmarshaler = (*ScpBallot)(nil)
)

// ScpStatementType is an XDR Enum defines as:
//
//   enum SCPStatementType
//    {
//        SCP_ST_PREPARE = 0,
//        SCP_ST_CONFIRM = 1,
//        SCP_ST_EXTERNALIZE = 2,
//        SCP_ST_NOMINATE = 3
//    };
//
type ScpStatementType int32

const (
	ScpStatementTypeScpStPrepare     ScpStatementType = 0
	ScpStatementTypeScpStConfirm     ScpStatementType = 1
	ScpStatementTypeScpStExternalize ScpStatementType = 2
	ScpStatementTypeScpStNominate    ScpStatementType = 3
)

var scpStatementTypeMap = map[int32]string{
	0: "ScpStatementTypeScpStPrepare",
	1: "ScpStatementTypeScpStConfirm",
	2: "ScpStatementTypeScpStExternalize",
	3: "ScpStatementTypeScpStNominate",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ScpStatementType
func (e ScpStatementType) ValidEnum(v int32) bool {
	_, ok := scpStatementTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e ScpStatementType) String() string {
	name, _ := scpStatementTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ScpStatementType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ScpStatementType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ScpStatementType)(nil)
	_ encoding.BinaryUnmarshaler = (*ScpStatementType)(nil)
)

// ScpNomination is an XDR Struct defines as:
//
//   struct SCPNomination
//    {
//        Hash quorumSetHash; // D
//        Value votes<>;      // X
//        Value accepted<>;   // Y
//    };
//
type ScpNomination struct {
	QuorumSetHash Hash
	Votes         []Value
	Accepted      []Value
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ScpNomination) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ScpNomination) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ScpNomination)(nil)
	_ encoding.BinaryUnmarshaler = (*ScpNomination)(nil)
)

// ScpStatementPrepare is an XDR NestedStruct defines as:
//
//   struct
//            {
//                Hash quorumSetHash;       // D
//                SCPBallot ballot;         // b
//                SCPBallot* prepared;      // p
//                SCPBallot* preparedPrime; // p'
//                uint32 nC;                // c.n
//                uint32 nH;                // h.n
//            }
//
type ScpStatementPrepare struct {
	QuorumSetHash Hash
	Ballot        ScpBallot
	Prepared      *ScpBallot
	PreparedPrime *ScpBallot
	NC            Uint32
	NH            Uint32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ScpStatementPrepare) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ScpStatementPrepare) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ScpStatementPrepare)(nil)
	_ encoding.BinaryUnmarshaler = (*ScpStatementPrepare)(nil)
)

// ScpStatementConfirm is an XDR NestedStruct defines as:
//
//   struct
//            {
//                SCPBallot ballot;   // b
//                uint32 nPrepared;   // p.n
//                uint32 nCommit;     // c.n
//                uint32 nH;          // h.n
//                Hash quorumSetHash; // D
//            }
//
type ScpStatementConfirm struct {
	Ballot        ScpBallot
	NPrepared     Uint32
	NCommit       Uint32
	NH            Uint32
	QuorumSetHash Hash
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ScpStatementConfirm) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ScpStatementConfirm) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ScpStatementConfirm)(nil)
	_ encoding.BinaryUnmarshaler = (*ScpStatementConfirm)(nil)
)

// ScpStatementExternalize is an XDR NestedStruct defines as:
//
//   struct
//            {
//                SCPBallot commit;         // c
//                uint32 nH;                // h.n
//                Hash commitQuorumSetHash; // D used before EXTERNALIZE
//            }
//
type ScpStatementExternalize struct {
	Commit              ScpBallot
	NH                  Uint32
	CommitQuorumSetHash Hash
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ScpStatementExternalize) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ScpStatementExternalize) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ScpStatementExternalize)(nil)
	_ encoding.BinaryUnmarshaler = (*ScpStatementExternalize)(nil)
)

// ScpStatementPledges is an XDR NestedUnion defines as:
//
//   union switch (SCPStatementType type)
//        {
//        case SCP_ST_PREPARE:
//            struct
//            {
//                Hash quorumSetHash;       // D
//                SCPBallot ballot;         // b
//                SCPBallot* prepared;      // p
//                SCPBallot* preparedPrime; // p'
//                uint32 nC;                // c.n
//                uint32 nH;                // h.n
//            } prepare;
//        case SCP_ST_CONFIRM:
//            struct
//            {
//                SCPBallot ballot;   // b
//                uint32 nPrepared;   // p.n
//                uint32 nCommit;     // c.n
//                uint32 nH;          // h.n
//                Hash quorumSetHash; // D
//            } confirm;
//        case SCP_ST_EXTERNALIZE:
//            struct
//            {
//                SCPBallot commit;         // c
//                uint32 nH;                // h.n
//                Hash commitQuorumSetHash; // D used before EXTERNALIZE
//            } externalize;
//        case SCP_ST_NOMINATE:
//            SCPNomination nominate;
//        }
//
type ScpStatementPledges struct {
	Type        ScpStatementType
	Prepare     *ScpStatementPrepare
	Confirm     *ScpStatementConfirm
	Externalize *ScpStatementExternalize
	Nominate    *ScpNomination
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ScpStatementPledges) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ScpStatementPledges
func (u ScpStatementPledges) ArmForSwitch(sw int32) (string, bool) {
	switch ScpStatementType(sw) {
	case ScpStatementTypeScpStPrepare:
		return "Prepare", true
	case ScpStatementTypeScpStConfirm:
		return "Confirm", true
	case ScpStatementTypeScpStExternalize:
		return "Externalize", true
	case ScpStatementTypeScpStNominate:
		return "Nominate", true
	}
	return "-", false
}

// NewScpStatementPledges creates a new  ScpStatementPledges.
func NewScpStatementPledges(aType ScpStatementType, value interface{}) (result ScpStatementPledges, err error) {
	result.Type = aType
	switch ScpStatementType(aType) {
	case ScpStatementTypeScpStPrepare:
		tv, ok := value.(ScpStatementPrepare)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpStatementPrepare")
			return
		}
		result.Prepare = &tv
	case ScpStatementTypeScpStConfirm:
		tv, ok := value.(ScpStatementConfirm)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpStatementConfirm")
			return
		}
		result.Confirm = &tv
	case ScpStatementTypeScpStExternalize:
		tv, ok := value.(ScpStatementExternalize)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpStatementExternalize")
			return
		}
		result.Externalize = &tv
	case ScpStatementTypeScpStNominate:
		tv, ok := value.(ScpNomination)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpNomination")
			return
		}
		result.Nominate = &tv
	}
	return
}

// MustPrepare retrieves the Prepare value from the union,
// panicing if the value is not set.
func (u ScpStatementPledges) MustPrepare() ScpStatementPrepare {
	val, ok := u.GetPrepare()

	if !ok {
		panic("arm Prepare is not set")
	}

	return val
}

// GetPrepare retrieves the Prepare value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpStatementPledges) GetPrepare() (result ScpStatementPrepare, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Prepare" {
		result = *u.Prepare
		ok = true
	}

	return
}

// MustConfirm retrieves the Confirm value from the union,
// panicing if the value is not set.
func (u ScpStatementPledges) MustConfirm() ScpStatementConfirm {
	val, ok := u.GetConfirm()

	if !ok {
		panic("arm Confirm is not set")
	}

	return val
}

// GetConfirm retrieves the Confirm value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpStatementPledges) GetConfirm() (result ScpStatementConfirm, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Confirm" {
		result = *u.Confirm
		ok = true
	}

	return
}

// MustExternalize retrieves the Externalize value from the union,
// panicing if the value is not set.
func (u ScpStatementPledges) MustExternalize() ScpStatementExternalize {
	val, ok := u.GetExternalize()

	if !ok {
		panic("arm Externalize is not set")
	}

	return val
}

// GetExternalize retrieves the Externalize value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpStatementPledges) GetExternalize() (result ScpStatementExternalize, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Externalize" {
		result = *u.Externalize
		ok = true
	}

	return
}

// MustNominate retrieves the Nominate value from the union,
// panicing if the value is not set.
func (u ScpStatementPledges) MustNominate() ScpNomination {
	val, ok := u.GetNominate()

	if !ok {
		panic("arm Nominate is not set")
	}

	return val
}

// GetNominate retrieves the Nominate value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpStatementPledges) GetNominate() (result ScpNomination, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Nominate" {
		result = *u.Nominate
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ScpStatementPledges) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ScpStatementPledges) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ScpStatementPledges)(nil)
	_ encoding.BinaryUnmarshaler = (*ScpStatementPledges)(nil)
)

// ScpStatement is an XDR Struct defines as:
//
//   struct SCPStatement
//    {
//        NodeID nodeID;    // v
//        uint64 slotIndex; // i
//
//        union switch (SCPStatementType type)
//        {
//        case SCP_ST_PREPARE:
//            struct
//            {
//                Hash quorumSetHash;       // D
//                SCPBallot ballot;         // b
//                SCPBallot* prepared;      // p
//                SCPBallot* preparedPrime; // p'
//                uint32 nC;                // c.n
//                uint32 nH;                // h.n
//            } prepare;
//        case SCP_ST_CONFIRM:
//            struct
//            {
//                SCPBallot ballot;   // b
//                uint32 nPrepared;   // p.n
//                uint32 nCommit;     // c.n
//                uint32 nH;          // h.n
//                Hash quorumSetHash; // D
//            } confirm;
//        case SCP_ST_EXTERNALIZE:
//            struct
//            {
//                SCPBallot commit;         // c
//                uint32 nH;                // h.n
//                Hash commitQuorumSetHash; // D used before EXTERNALIZE
//            } externalize;
//        case SCP_ST_NOMINATE:
//            SCPNomination nominate;
//        }
//        pledges;
//    };
//
type ScpStatement struct {
	NodeId    NodeId
	SlotIndex Uint64
	Pledges   ScpStatementPledges
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ScpStatement) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ScpStatement) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ScpStatement)(nil)
	_ encoding.BinaryUnmarshaler = (*ScpStatement)(nil)
)

// ScpEnvelope is an XDR Struct defines as:
//
//   struct SCPEnvelope
//    {
//        SCPStatement statement;
//        Signature signature;
//    };
//
type ScpEnvelope struct {
	Statement ScpStatement
	Signature Signature
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ScpEnvelope) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ScpEnvelope) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ScpEnvelope)(nil)
	_ encoding.BinaryUnmarshaler = (*ScpEnvelope)(nil)
)

// ScpQuorumSet is an XDR Struct defines as:
//
//   struct SCPQuorumSet
//    {
//        uint32 threshold;
//        NodeID validators<>;
//        SCPQuorumSet innerSets<>;
//    };
//
type ScpQuorumSet struct {
	Threshold  Uint32
	Validators []NodeId
	InnerSets  []ScpQuorumSet
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ScpQuorumSet) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ScpQuorumSet) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ScpQuorumSet)(nil)
	_ encoding.BinaryUnmarshaler = (*ScpQuorumSet)(nil)
)

// AccountId is an XDR Typedef defines as:
//
//   typedef PublicKey AccountID;
//
type AccountId PublicKey

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountId) SwitchFieldName() string {
	return PublicKey(u).SwitchFieldName()
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PublicKey
func (u AccountId) ArmForSwitch(sw int32) (string, bool) {
	return PublicKey(u).ArmForSwitch(sw)
}

// NewAccountId creates a new  AccountId.
func NewAccountId(aType PublicKeyType, value interface{}) (result AccountId, err error) {
	u, err := NewPublicKey(aType, value)
	result = AccountId(u)
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u AccountId) MustEd25519() Uint256 {
	return PublicKey(u).MustEd25519()
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountId) GetEd25519() (result Uint256, ok bool) {
	return PublicKey(u).GetEd25519()
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountId) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountId) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountId)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountId)(nil)
)

// Thresholds is an XDR Typedef defines as:
//
//   typedef opaque Thresholds[4];
//
type Thresholds [4]byte

// XDRMaxSize implements the Sized interface for Thresholds
func (e Thresholds) XDRMaxSize() int {
	return 4
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Thresholds) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Thresholds) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Thresholds)(nil)
	_ encoding.BinaryUnmarshaler = (*Thresholds)(nil)
)

// String32 is an XDR Typedef defines as:
//
//   typedef string string32<32>;
//
type String32 string

// XDRMaxSize implements the Sized interface for String32
func (e String32) XDRMaxSize() int {
	return 32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s String32) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *String32) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*String32)(nil)
	_ encoding.BinaryUnmarshaler = (*String32)(nil)
)

// String64 is an XDR Typedef defines as:
//
//   typedef string string64<64>;
//
type String64 string

// XDRMaxSize implements the Sized interface for String64
func (e String64) XDRMaxSize() int {
	return 64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s String64) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *String64) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*String64)(nil)
	_ encoding.BinaryUnmarshaler = (*String64)(nil)
)

// SequenceNumber is an XDR Typedef defines as:
//
//   typedef int64 SequenceNumber;
//
type SequenceNumber Int64

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SequenceNumber) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SequenceNumber) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SequenceNumber)(nil)
	_ encoding.BinaryUnmarshaler = (*SequenceNumber)(nil)
)

// TimePoint is an XDR Typedef defines as:
//
//   typedef uint64 TimePoint;
//
type TimePoint Uint64

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TimePoint) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TimePoint) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TimePoint)(nil)
	_ encoding.BinaryUnmarshaler = (*TimePoint)(nil)
)

// DataValue is an XDR Typedef defines as:
//
//   typedef opaque DataValue<64>;
//
type DataValue []byte

// XDRMaxSize implements the Sized interface for DataValue
func (e DataValue) XDRMaxSize() int {
	return 64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s DataValue) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *DataValue) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*DataValue)(nil)
	_ encoding.BinaryUnmarshaler = (*DataValue)(nil)
)

// PoolId is an XDR Typedef defines as:
//
//   typedef Hash PoolID;
//
type PoolId Hash

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PoolId) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PoolId) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PoolId)(nil)
	_ encoding.BinaryUnmarshaler = (*PoolId)(nil)
)

// AssetCode4 is an XDR Typedef defines as:
//
//   typedef opaque AssetCode4[4];
//
type AssetCode4 [4]byte

// XDRMaxSize implements the Sized interface for AssetCode4
func (e AssetCode4) XDRMaxSize() int {
	return 4
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AssetCode4) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AssetCode4) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AssetCode4)(nil)
	_ encoding.BinaryUnmarshaler = (*AssetCode4)(nil)
)

// AssetCode12 is an XDR Typedef defines as:
//
//   typedef opaque AssetCode12[12];
//
type AssetCode12 [12]byte

// XDRMaxSize implements the Sized interface for AssetCode12
func (e AssetCode12) XDRMaxSize() int {
	return 12
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AssetCode12) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AssetCode12) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AssetCode12)(nil)
	_ encoding.BinaryUnmarshaler = (*AssetCode12)(nil)
)

// AssetType is an XDR Enum defines as:
//
//   enum AssetType
//    {
//        ASSET_TYPE_NATIVE = 0,
//        ASSET_TYPE_CREDIT_ALPHANUM4 = 1,
//        ASSET_TYPE_CREDIT_ALPHANUM12 = 2,
//        ASSET_TYPE_POOL_SHARE = 3
//    };
//
type AssetType int32

const (
	AssetTypeAssetTypeNative           AssetType = 0
	AssetTypeAssetTypeCreditAlphanum4  AssetType = 1
	AssetTypeAssetTypeCreditAlphanum12 AssetType = 2
	AssetTypeAssetTypePoolShare        AssetType = 3
)

var assetTypeMap = map[int32]string{
	0: "AssetTypeAssetTypeNative",
	1: "AssetTypeAssetTypeCreditAlphanum4",
	2: "AssetTypeAssetTypeCreditAlphanum12",
	3: "AssetTypeAssetTypePoolShare",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for AssetType
func (e AssetType) ValidEnum(v int32) bool {
	_, ok := assetTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e AssetType) String() string {
	name, _ := assetTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AssetType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AssetType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AssetType)(nil)
	_ encoding.BinaryUnmarshaler = (*AssetType)(nil)
)

// AssetCode is an XDR Union defines as:
//
//   union AssetCode switch (AssetType type)
//    {
//    case ASSET_TYPE_CREDIT_ALPHANUM4:
//        AssetCode4 assetCode4;
//
//    case ASSET_TYPE_CREDIT_ALPHANUM12:
//        AssetCode12 assetCode12;
//
//        // add other asset types here in the future
//    };
//
type AssetCode struct {
	Type        AssetType
	AssetCode4  *AssetCode4
	AssetCode12 *AssetCode12
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AssetCode) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AssetCode
func (u AssetCode) ArmForSwitch(sw int32) (string, bool) {
	switch AssetType(sw) {
	case AssetTypeAssetTypeCreditAlphanum4:
		return "AssetCode4", true
	case AssetTypeAssetTypeCreditAlphanum12:
		return "AssetCode12", true
	}
	return "-", false
}

// NewAssetCode creates a new  AssetCode.
func NewAssetCode(aType AssetType, value interface{}) (result AssetCode, err error) {
	result.Type = aType
	switch AssetType(aType) {
	case AssetTypeAssetTypeCreditAlphanum4:
		tv, ok := value.(AssetCode4)
		if !ok {
			err = fmt.Errorf("invalid value, must be AssetCode4")
			return
		}
		result.AssetCode4 = &tv
	case AssetTypeAssetTypeCreditAlphanum12:
		tv, ok := value.(AssetCode12)
		if !ok {
			err = fmt.Errorf("invalid value, must be AssetCode12")
			return
		}
		result.AssetCode12 = &tv
	}
	return
}

// MustAssetCode4 retrieves the AssetCode4 value from the union,
// panicing if the value is not set.
func (u AssetCode) MustAssetCode4() AssetCode4 {
	val, ok := u.GetAssetCode4()

	if !ok {
		panic("arm AssetCode4 is not set")
	}

	return val
}

// GetAssetCode4 retrieves the AssetCode4 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AssetCode) GetAssetCode4() (result AssetCode4, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AssetCode4" {
		result = *u.AssetCode4
		ok = true
	}

	return
}

// MustAssetCode12 retrieves the AssetCode12 value from the union,
// panicing if the value is not set.
func (u AssetCode) MustAssetCode12() AssetCode12 {
	val, ok := u.GetAssetCode12()

	if !ok {
		panic("arm AssetCode12 is not set")
	}

	return val
}

// GetAssetCode12 retrieves the AssetCode12 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AssetCode) GetAssetCode12() (result AssetCode12, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AssetCode12" {
		result = *u.AssetCode12
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AssetCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AssetCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AssetCode)(nil)
	_ encoding.BinaryUnmarshaler = (*AssetCode)(nil)
)

// AlphaNum4 is an XDR Struct defines as:
//
//   struct AlphaNum4
//    {
//        AssetCode4 assetCode;
//        AccountID issuer;
//    };
//
type AlphaNum4 struct {
	AssetCode AssetCode4
	Issuer    AccountId
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AlphaNum4) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AlphaNum4) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AlphaNum4)(nil)
	_ encoding.BinaryUnmarshaler = (*AlphaNum4)(nil)
)

// AlphaNum12 is an XDR Struct defines as:
//
//   struct AlphaNum12
//    {
//        AssetCode12 assetCode;
//        AccountID issuer;
//    };
//
type AlphaNum12 struct {
	AssetCode AssetCode12
	Issuer    AccountId
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AlphaNum12) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AlphaNum12) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AlphaNum12)(nil)
	_ encoding.BinaryUnmarshaler = (*AlphaNum12)(nil)
)

// Asset is an XDR Union defines as:
//
//   union Asset switch (AssetType type)
//    {
//    case ASSET_TYPE_NATIVE: // Not credit
//        void;
//
//    case ASSET_TYPE_CREDIT_ALPHANUM4:
//        AlphaNum4 alphaNum4;
//
//    case ASSET_TYPE_CREDIT_ALPHANUM12:
//        AlphaNum12 alphaNum12;
//
//        // add other asset types here in the future
//    };
//
type Asset struct {
	Type       AssetType
	AlphaNum4  *AlphaNum4
	AlphaNum12 *AlphaNum12
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u Asset) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of Asset
func (u Asset) ArmForSwitch(sw int32) (string, bool) {
	switch AssetType(sw) {
	case AssetTypeAssetTypeNative:
		return "", true
	case AssetTypeAssetTypeCreditAlphanum4:
		return "AlphaNum4", true
	case AssetTypeAssetTypeCreditAlphanum12:
		return "AlphaNum12", true
	}
	return "-", false
}

// NewAsset creates a new  Asset.
func NewAsset(aType AssetType, value interface{}) (result Asset, err error) {
	result.Type = aType
	switch AssetType(aType) {
	case AssetTypeAssetTypeNative:
		// void
	case AssetTypeAssetTypeCreditAlphanum4:
		tv, ok := value.(AlphaNum4)
		if !ok {
			err = fmt.Errorf("invalid value, must be AlphaNum4")
			return
		}
		result.AlphaNum4 = &tv
	case AssetTypeAssetTypeCreditAlphanum12:
		tv, ok := value.(AlphaNum12)
		if !ok {
			err = fmt.Errorf("invalid value, must be AlphaNum12")
			return
		}
		result.AlphaNum12 = &tv
	}
	return
}

// MustAlphaNum4 retrieves the AlphaNum4 value from the union,
// panicing if the value is not set.
func (u Asset) MustAlphaNum4() AlphaNum4 {
	val, ok := u.GetAlphaNum4()

	if !ok {
		panic("arm AlphaNum4 is not set")
	}

	return val
}

// GetAlphaNum4 retrieves the AlphaNum4 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Asset) GetAlphaNum4() (result AlphaNum4, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AlphaNum4" {
		result = *u.AlphaNum4
		ok = true
	}

	return
}

// MustAlphaNum12 retrieves the AlphaNum12 value from the union,
// panicing if the value is not set.
func (u Asset) MustAlphaNum12() AlphaNum12 {
	val, ok := u.GetAlphaNum12()

	if !ok {
		panic("arm AlphaNum12 is not set")
	}

	return val
}

// GetAlphaNum12 retrieves the AlphaNum12 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Asset) GetAlphaNum12() (result AlphaNum12, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AlphaNum12" {
		result = *u.AlphaNum12
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Asset) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Asset) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Asset)(nil)
	_ encoding.BinaryUnmarshaler = (*Asset)(nil)
)

// Price is an XDR Struct defines as:
//
//   struct Price
//    {
//        int32 n; // numerator
//        int32 d; // denominator
//    };
//
type Price struct {
	N Int32
	D Int32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Price) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Price) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Price)(nil)
	_ encoding.BinaryUnmarshaler = (*Price)(nil)
)

// Liabilities is an XDR Struct defines as:
//
//   struct Liabilities
//    {
//        int64 buying;
//        int64 selling;
//    };
//
type Liabilities struct {
	Buying  Int64
	Selling Int64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Liabilities) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Liabilities) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Liabilities)(nil)
	_ encoding.BinaryUnmarshaler = (*Liabilities)(nil)
)

// ThresholdIndexes is an XDR Enum defines as:
//
//   enum ThresholdIndexes
//    {
//        THRESHOLD_MASTER_WEIGHT = 0,
//        THRESHOLD_LOW = 1,
//        THRESHOLD_MED = 2,
//        THRESHOLD_HIGH = 3
//    };
//
type ThresholdIndexes int32

const (
	ThresholdIndexesThresholdMasterWeight ThresholdIndexes = 0
	ThresholdIndexesThresholdLow          ThresholdIndexes = 1
	ThresholdIndexesThresholdMed          ThresholdIndexes = 2
	ThresholdIndexesThresholdHigh         ThresholdIndexes = 3
)

var thresholdIndexesMap = map[int32]string{
	0: "ThresholdIndexesThresholdMasterWeight",
	1: "ThresholdIndexesThresholdLow",
	2: "ThresholdIndexesThresholdMed",
	3: "ThresholdIndexesThresholdHigh",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ThresholdIndexes
func (e ThresholdIndexes) ValidEnum(v int32) bool {
	_, ok := thresholdIndexesMap[v]
	return ok
}

// String returns the name of `e`
func (e ThresholdIndexes) String() string {
	name, _ := thresholdIndexesMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ThresholdIndexes) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ThresholdIndexes) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ThresholdIndexes)(nil)
	_ encoding.BinaryUnmarshaler = (*ThresholdIndexes)(nil)
)

// LedgerEntryType is an XDR Enum defines as:
//
//   enum LedgerEntryType
//    {
//        ACCOUNT = 0,
//        TRUSTLINE = 1,
//        OFFER = 2,
//        DATA = 3,
//        CLAIMABLE_BALANCE = 4,
//        LIQUIDITY_POOL = 5
//    };
//
type LedgerEntryType int32

const (
	LedgerEntryTypeAccount          LedgerEntryType = 0
	LedgerEntryTypeTrustline        LedgerEntryType = 1
	LedgerEntryTypeOffer            LedgerEntryType = 2
	LedgerEntryTypeData             LedgerEntryType = 3
	LedgerEntryTypeClaimableBalance LedgerEntryType = 4
	LedgerEntryTypeLiquidityPool    LedgerEntryType = 5
)

var ledgerEntryTypeMap = map[int32]string{
	0: "LedgerEntryTypeAccount",
	1: "LedgerEntryTypeTrustline",
	2: "LedgerEntryTypeOffer",
	3: "LedgerEntryTypeData",
	4: "LedgerEntryTypeClaimableBalance",
	5: "LedgerEntryTypeLiquidityPool",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LedgerEntryType
func (e LedgerEntryType) ValidEnum(v int32) bool {
	_, ok := ledgerEntryTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e LedgerEntryType) String() string {
	name, _ := ledgerEntryTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerEntryType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerEntryType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerEntryType)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerEntryType)(nil)
)

// Signer is an XDR Struct defines as:
//
//   struct Signer
//    {
//        SignerKey key;
//        uint32 weight; // really only need 1 byte
//    };
//
type Signer struct {
	Key    SignerKey
	Weight Uint32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Signer) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Signer) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Signer)(nil)
	_ encoding.BinaryUnmarshaler = (*Signer)(nil)
)

// AccountFlags is an XDR Enum defines as:
//
//   enum AccountFlags
//    { // masks for each flag
//
//        // Flags set on issuer accounts
//        // TrustLines are created with authorized set to "false" requiring
//        // the issuer to set it for each TrustLine
//        AUTH_REQUIRED_FLAG = 0x1,
//        // If set, the authorized flag in TrustLines can be cleared
//        // otherwise, authorization cannot be revoked
//        AUTH_REVOCABLE_FLAG = 0x2,
//        // Once set, causes all AUTH_* flags to be read-only
//        AUTH_IMMUTABLE_FLAG = 0x4,
//        // Trustlines are created with clawback enabled set to "true",
//        // and claimable balances created from those trustlines are created
//        // with clawback enabled set to "true"
//        AUTH_CLAWBACK_ENABLED_FLAG = 0x8
//    };
//
type AccountFlags int32

const (
	AccountFlagsAuthRequiredFlag        AccountFlags = 1
	AccountFlagsAuthRevocableFlag       AccountFlags = 2
	AccountFlagsAuthImmutableFlag       AccountFlags = 4
	AccountFlagsAuthClawbackEnabledFlag AccountFlags = 8
)

var accountFlagsMap = map[int32]string{
	1: "AccountFlagsAuthRequiredFlag",
	2: "AccountFlagsAuthRevocableFlag",
	4: "AccountFlagsAuthImmutableFlag",
	8: "AccountFlagsAuthClawbackEnabledFlag",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for AccountFlags
func (e AccountFlags) ValidEnum(v int32) bool {
	_, ok := accountFlagsMap[v]
	return ok
}

// String returns the name of `e`
func (e AccountFlags) String() string {
	name, _ := accountFlagsMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountFlags) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountFlags) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountFlags)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountFlags)(nil)
)

// MaskAccountFlags is an XDR Const defines as:
//
//   const MASK_ACCOUNT_FLAGS = 0x7;
//
const MaskAccountFlags = 0x7

// MaskAccountFlagsV17 is an XDR Const defines as:
//
//   const MASK_ACCOUNT_FLAGS_V17 = 0xF;
//
const MaskAccountFlagsV17 = 0xF

// MaxSigners is an XDR Const defines as:
//
//   const MAX_SIGNERS = 20;
//
const MaxSigners = 20

// SponsorshipDescriptor is an XDR Typedef defines as:
//
//   typedef AccountID* SponsorshipDescriptor;
//
type SponsorshipDescriptor *AccountId

// AccountEntryExtensionV2Ext is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        }
//
type AccountEntryExtensionV2Ext struct {
	V int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountEntryExtensionV2Ext) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AccountEntryExtensionV2Ext
func (u AccountEntryExtensionV2Ext) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	}
	return "-", false
}

// NewAccountEntryExtensionV2Ext creates a new  AccountEntryExtensionV2Ext.
func NewAccountEntryExtensionV2Ext(v int32, value interface{}) (result AccountEntryExtensionV2Ext, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountEntryExtensionV2Ext) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountEntryExtensionV2Ext) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountEntryExtensionV2Ext)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountEntryExtensionV2Ext)(nil)
)

// AccountEntryExtensionV2 is an XDR Struct defines as:
//
//   struct AccountEntryExtensionV2
//    {
//        uint32 numSponsored;
//        uint32 numSponsoring;
//        SponsorshipDescriptor signerSponsoringIDs<MAX_SIGNERS>;
//
//        union switch (int v)
//        {
//        case 0:
//            void;
//        }
//        ext;
//    };
//
type AccountEntryExtensionV2 struct {
	NumSponsored        Uint32
	NumSponsoring       Uint32
	SignerSponsoringIDs []SponsorshipDescriptor `xdrmaxsize:"20"`
	Ext                 AccountEntryExtensionV2Ext
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountEntryExtensionV2) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountEntryExtensionV2) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountEntryExtensionV2)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountEntryExtensionV2)(nil)
)

// AccountEntryExtensionV1Ext is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        case 2:
//            AccountEntryExtensionV2 v2;
//        }
//
type AccountEntryExtensionV1Ext struct {
	V  int32
	V2 *AccountEntryExtensionV2
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountEntryExtensionV1Ext) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AccountEntryExtensionV1Ext
func (u AccountEntryExtensionV1Ext) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	case 2:
		return "V2", true
	}
	return "-", false
}

// NewAccountEntryExtensionV1Ext creates a new  AccountEntryExtensionV1Ext.
func NewAccountEntryExtensionV1Ext(v int32, value interface{}) (result AccountEntryExtensionV1Ext, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	case 2:
		tv, ok := value.(AccountEntryExtensionV2)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountEntryExtensionV2")
			return
		}
		result.V2 = &tv
	}
	return
}

// MustV2 retrieves the V2 value from the union,
// panicing if the value is not set.
func (u AccountEntryExtensionV1Ext) MustV2() AccountEntryExtensionV2 {
	val, ok := u.GetV2()

	if !ok {
		panic("arm V2 is not set")
	}

	return val
}

// GetV2 retrieves the V2 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountEntryExtensionV1Ext) GetV2() (result AccountEntryExtensionV2, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "V2" {
		result = *u.V2
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountEntryExtensionV1Ext) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountEntryExtensionV1Ext) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountEntryExtensionV1Ext)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountEntryExtensionV1Ext)(nil)
)

// AccountEntryExtensionV1 is an XDR Struct defines as:
//
//   struct AccountEntryExtensionV1
//    {
//        Liabilities liabilities;
//
//        union switch (int v)
//        {
//        case 0:
//            void;
//        case 2:
//            AccountEntryExtensionV2 v2;
//        }
//        ext;
//    };
//
type AccountEntryExtensionV1 struct {
	Liabilities Liabilities
	Ext         AccountEntryExtensionV1Ext
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountEntryExtensionV1) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountEntryExtensionV1) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountEntryExtensionV1)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountEntryExtensionV1)(nil)
)

// AccountEntryExt is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        case 1:
//            AccountEntryExtensionV1 v1;
//        }
//
type AccountEntryExt struct {
	V  int32
	V1 *AccountEntryExtensionV1
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AccountEntryExt
func (u AccountEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	case 1:
		return "V1", true
	}
	return "-", false
}

// NewAccountEntryExt creates a new  AccountEntryExt.
func NewAccountEntryExt(v int32, value interface{}) (result AccountEntryExt, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	case 1:
		tv, ok := value.(AccountEntryExtensionV1)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountEntryExtensionV1")
			return
		}
		result.V1 = &tv
	}
	return
}

// MustV1 retrieves the V1 value from the union,
// panicing if the value is not set.
func (u AccountEntryExt) MustV1() AccountEntryExtensionV1 {
	val, ok := u.GetV1()

	if !ok {
		panic("arm V1 is not set")
	}

	return val
}

// GetV1 retrieves the V1 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountEntryExt) GetV1() (result AccountEntryExtensionV1, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "V1" {
		result = *u.V1
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountEntryExt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountEntryExt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountEntryExt)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountEntryExt)(nil)
)

// AccountEntry is an XDR Struct defines as:
//
//   struct AccountEntry
//    {
//        AccountID accountID;      // master public key for this account
//        int64 balance;            // in stroops
//        SequenceNumber seqNum;    // last sequence number used for this account
//        uint32 numSubEntries;     // number of sub-entries this account has
//                                  // drives the reserve
//        AccountID* inflationDest; // Account to vote for during inflation
//        uint32 flags;             // see AccountFlags
//
//        string32 homeDomain; // can be used for reverse federation and memo lookup
//
//        // fields used for signatures
//        // thresholds stores unsigned bytes: [weight of master|low|medium|high]
//        Thresholds thresholds;
//
//        Signer signers<MAX_SIGNERS>; // possible signers for this account
//
//        // reserved for future use
//        union switch (int v)
//        {
//        case 0:
//            void;
//        case 1:
//            AccountEntryExtensionV1 v1;
//        }
//        ext;
//    };
//
type AccountEntry struct {
	AccountId     AccountId
	Balance       Int64
	SeqNum        SequenceNumber
	NumSubEntries Uint32
	InflationDest *AccountId
	Flags         Uint32
	HomeDomain    String32
	Thresholds    Thresholds
	Signers       []Signer `xdrmaxsize:"20"`
	Ext           AccountEntryExt
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountEntry) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountEntry) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountEntry)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountEntry)(nil)
)

// TrustLineFlags is an XDR Enum defines as:
//
//   enum TrustLineFlags
//    {
//        // issuer has authorized account to perform transactions with its credit
//        AUTHORIZED_FLAG = 1,
//        // issuer has authorized account to maintain and reduce liabilities for its
//        // credit
//        AUTHORIZED_TO_MAINTAIN_LIABILITIES_FLAG = 2,
//        // issuer has specified that it may clawback its credit, and that claimable
//        // balances created with its credit may also be clawed back
//        TRUSTLINE_CLAWBACK_ENABLED_FLAG = 4
//    };
//
type TrustLineFlags int32

const (
	TrustLineFlagsAuthorizedFlag                      TrustLineFlags = 1
	TrustLineFlagsAuthorizedToMaintainLiabilitiesFlag TrustLineFlags = 2
	TrustLineFlagsTrustlineClawbackEnabledFlag        TrustLineFlags = 4
)

var trustLineFlagsMap = map[int32]string{
	1: "TrustLineFlagsAuthorizedFlag",
	2: "TrustLineFlagsAuthorizedToMaintainLiabilitiesFlag",
	4: "TrustLineFlagsTrustlineClawbackEnabledFlag",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for TrustLineFlags
func (e TrustLineFlags) ValidEnum(v int32) bool {
	_, ok := trustLineFlagsMap[v]
	return ok
}

// String returns the name of `e`
func (e TrustLineFlags) String() string {
	name, _ := trustLineFlagsMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TrustLineFlags) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TrustLineFlags) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TrustLineFlags)(nil)
	_ encoding.BinaryUnmarshaler = (*TrustLineFlags)(nil)
)

// MaskTrustlineFlags is an XDR Const defines as:
//
//   const MASK_TRUSTLINE_FLAGS = 1;
//
const MaskTrustlineFlags = 1

// MaskTrustlineFlagsV13 is an XDR Const defines as:
//
//   const MASK_TRUSTLINE_FLAGS_V13 = 3;
//
const MaskTrustlineFlagsV13 = 3

// MaskTrustlineFlagsV17 is an XDR Const defines as:
//
//   const MASK_TRUSTLINE_FLAGS_V17 = 7;
//
const MaskTrustlineFlagsV17 = 7

// LiquidityPoolType is an XDR Enum defines as:
//
//   enum LiquidityPoolType
//    {
//        LIQUIDITY_POOL_CONSTANT_PRODUCT = 0
//    };
//
type LiquidityPoolType int32

const (
	LiquidityPoolTypeLiquidityPoolConstantProduct LiquidityPoolType = 0
)

var liquidityPoolTypeMap = map[int32]string{
	0: "LiquidityPoolTypeLiquidityPoolConstantProduct",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LiquidityPoolType
func (e LiquidityPoolType) ValidEnum(v int32) bool {
	_, ok := liquidityPoolTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e LiquidityPoolType) String() string {
	name, _ := liquidityPoolTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LiquidityPoolType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LiquidityPoolType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LiquidityPoolType)(nil)
	_ encoding.BinaryUnmarshaler = (*LiquidityPoolType)(nil)
)

// TrustLineAsset is an XDR Union defines as:
//
//   union TrustLineAsset switch (AssetType type)
//    {
//    case ASSET_TYPE_NATIVE: // Not credit
//        void;
//
//    case ASSET_TYPE_CREDIT_ALPHANUM4:
//        AlphaNum4 alphaNum4;
//
//    case ASSET_TYPE_CREDIT_ALPHANUM12:
//        AlphaNum12 alphaNum12;
//
//    case ASSET_TYPE_POOL_SHARE:
//        PoolID liquidityPoolID;
//
//        // add other asset types here in the future
//    };
//
type TrustLineAsset struct {
	Type            AssetType
	AlphaNum4       *AlphaNum4
	AlphaNum12      *AlphaNum12
	LiquidityPoolId *PoolId
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TrustLineAsset) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TrustLineAsset
func (u TrustLineAsset) ArmForSwitch(sw int32) (string, bool) {
	switch AssetType(sw) {
	case AssetTypeAssetTypeNative:
		return "", true
	case AssetTypeAssetTypeCreditAlphanum4:
		return "AlphaNum4", true
	case AssetTypeAssetTypeCreditAlphanum12:
		return "AlphaNum12", true
	case AssetTypeAssetTypePoolShare:
		return "LiquidityPoolId", true
	}
	return "-", false
}

// NewTrustLineAsset creates a new  TrustLineAsset.
func NewTrustLineAsset(aType AssetType, value interface{}) (result TrustLineAsset, err error) {
	result.Type = aType
	switch AssetType(aType) {
	case AssetTypeAssetTypeNative:
		// void
	case AssetTypeAssetTypeCreditAlphanum4:
		tv, ok := value.(AlphaNum4)
		if !ok {
			err = fmt.Errorf("invalid value, must be AlphaNum4")
			return
		}
		result.AlphaNum4 = &tv
	case AssetTypeAssetTypeCreditAlphanum12:
		tv, ok := value.(AlphaNum12)
		if !ok {
			err = fmt.Errorf("invalid value, must be AlphaNum12")
			return
		}
		result.AlphaNum12 = &tv
	case AssetTypeAssetTypePoolShare:
		tv, ok := value.(PoolId)
		if !ok {
			err = fmt.Errorf("invalid value, must be PoolId")
			return
		}
		result.LiquidityPoolId = &tv
	}
	return
}

// MustAlphaNum4 retrieves the AlphaNum4 value from the union,
// panicing if the value is not set.
func (u TrustLineAsset) MustAlphaNum4() AlphaNum4 {
	val, ok := u.GetAlphaNum4()

	if !ok {
		panic("arm AlphaNum4 is not set")
	}

	return val
}

// GetAlphaNum4 retrieves the AlphaNum4 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TrustLineAsset) GetAlphaNum4() (result AlphaNum4, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AlphaNum4" {
		result = *u.AlphaNum4
		ok = true
	}

	return
}

// MustAlphaNum12 retrieves the AlphaNum12 value from the union,
// panicing if the value is not set.
func (u TrustLineAsset) MustAlphaNum12() AlphaNum12 {
	val, ok := u.GetAlphaNum12()

	if !ok {
		panic("arm AlphaNum12 is not set")
	}

	return val
}

// GetAlphaNum12 retrieves the AlphaNum12 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TrustLineAsset) GetAlphaNum12() (result AlphaNum12, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AlphaNum12" {
		result = *u.AlphaNum12
		ok = true
	}

	return
}

// MustLiquidityPoolId retrieves the LiquidityPoolId value from the union,
// panicing if the value is not set.
func (u TrustLineAsset) MustLiquidityPoolId() PoolId {
	val, ok := u.GetLiquidityPoolId()

	if !ok {
		panic("arm LiquidityPoolId is not set")
	}

	return val
}

// GetLiquidityPoolId retrieves the LiquidityPoolId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TrustLineAsset) GetLiquidityPoolId() (result PoolId, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "LiquidityPoolId" {
		result = *u.LiquidityPoolId
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TrustLineAsset) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TrustLineAsset) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TrustLineAsset)(nil)
	_ encoding.BinaryUnmarshaler = (*TrustLineAsset)(nil)
)

// TrustLineEntryV1Ext is an XDR NestedUnion defines as:
//
//   union switch (int v)
//                {
//                case 0:
//                    void;
//                }
//
type TrustLineEntryV1Ext struct {
	V int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TrustLineEntryV1Ext) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TrustLineEntryV1Ext
func (u TrustLineEntryV1Ext) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	}
	return "-", false
}

// NewTrustLineEntryV1Ext creates a new  TrustLineEntryV1Ext.
func NewTrustLineEntryV1Ext(v int32, value interface{}) (result TrustLineEntryV1Ext, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TrustLineEntryV1Ext) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TrustLineEntryV1Ext) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TrustLineEntryV1Ext)(nil)
	_ encoding.BinaryUnmarshaler = (*TrustLineEntryV1Ext)(nil)
)

// TrustLineEntryV1 is an XDR NestedStruct defines as:
//
//   struct
//            {
//                Liabilities liabilities;
//
//                union switch (int v)
//                {
//                case 0:
//                    void;
//                }
//                ext;
//            }
//
type TrustLineEntryV1 struct {
	Liabilities Liabilities
	Ext         TrustLineEntryV1Ext
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TrustLineEntryV1) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TrustLineEntryV1) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TrustLineEntryV1)(nil)
	_ encoding.BinaryUnmarshaler = (*TrustLineEntryV1)(nil)
)

// TrustLineEntryExt is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        case 1:
//            struct
//            {
//                Liabilities liabilities;
//
//                union switch (int v)
//                {
//                case 0:
//                    void;
//                }
//                ext;
//            } v1;
//        }
//
type TrustLineEntryExt struct {
	V  int32
	V1 *TrustLineEntryV1
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TrustLineEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TrustLineEntryExt
func (u TrustLineEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	case 1:
		return "V1", true
	}
	return "-", false
}

// NewTrustLineEntryExt creates a new  TrustLineEntryExt.
func NewTrustLineEntryExt(v int32, value interface{}) (result TrustLineEntryExt, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	case 1:
		tv, ok := value.(TrustLineEntryV1)
		if !ok {
			err = fmt.Errorf("invalid value, must be TrustLineEntryV1")
			return
		}
		result.V1 = &tv
	}
	return
}

// MustV1 retrieves the V1 value from the union,
// panicing if the value is not set.
func (u TrustLineEntryExt) MustV1() TrustLineEntryV1 {
	val, ok := u.GetV1()

	if !ok {
		panic("arm V1 is not set")
	}

	return val
}

// GetV1 retrieves the V1 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TrustLineEntryExt) GetV1() (result TrustLineEntryV1, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "V1" {
		result = *u.V1
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TrustLineEntryExt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TrustLineEntryExt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TrustLineEntryExt)(nil)
	_ encoding.BinaryUnmarshaler = (*TrustLineEntryExt)(nil)
)

// TrustLineEntry is an XDR Struct defines as:
//
//   struct TrustLineEntry
//    {
//        AccountID accountID; // account this trustline belongs to
//        TrustLineAsset asset;         // type of asset (with issuer)
//        int64 balance;       // how much of this asset the user has.
//                             // Asset defines the unit for this;
//
//        int64 limit;  // balance cannot be above this
//        uint32 flags; // see TrustLineFlags
//
//        // reserved for future use
//        union switch (int v)
//        {
//        case 0:
//            void;
//        case 1:
//            struct
//            {
//                Liabilities liabilities;
//
//                union switch (int v)
//                {
//                case 0:
//                    void;
//                }
//                ext;
//            } v1;
//        }
//        ext;
//    };
//
type TrustLineEntry struct {
	AccountId AccountId
	Asset     TrustLineAsset
	Balance   Int64
	Limit     Int64
	Flags     Uint32
	Ext       TrustLineEntryExt
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TrustLineEntry) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TrustLineEntry) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TrustLineEntry)(nil)
	_ encoding.BinaryUnmarshaler = (*TrustLineEntry)(nil)
)

// OfferEntryFlags is an XDR Enum defines as:
//
//   enum OfferEntryFlags
//    {
//        // issuer has authorized account to perform transactions with its credit
//        PASSIVE_FLAG = 1
//    };
//
type OfferEntryFlags int32

const (
	OfferEntryFlagsPassiveFlag OfferEntryFlags = 1
)

var offerEntryFlagsMap = map[int32]string{
	1: "OfferEntryFlagsPassiveFlag",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for OfferEntryFlags
func (e OfferEntryFlags) ValidEnum(v int32) bool {
	_, ok := offerEntryFlagsMap[v]
	return ok
}

// String returns the name of `e`
func (e OfferEntryFlags) String() string {
	name, _ := offerEntryFlagsMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s OfferEntryFlags) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *OfferEntryFlags) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*OfferEntryFlags)(nil)
	_ encoding.BinaryUnmarshaler = (*OfferEntryFlags)(nil)
)

// MaskOfferentryFlags is an XDR Const defines as:
//
//   const MASK_OFFERENTRY_FLAGS = 1;
//
const MaskOfferentryFlags = 1

// OfferEntryExt is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        }
//
type OfferEntryExt struct {
	V int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OfferEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OfferEntryExt
func (u OfferEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	}
	return "-", false
}

// NewOfferEntryExt creates a new  OfferEntryExt.
func NewOfferEntryExt(v int32, value interface{}) (result OfferEntryExt, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s OfferEntryExt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *OfferEntryExt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*OfferEntryExt)(nil)
	_ encoding.BinaryUnmarshaler = (*OfferEntryExt)(nil)
)

// OfferEntry is an XDR Struct defines as:
//
//   struct OfferEntry
//    {
//        AccountID sellerID;
//        int64 offerID;
//        Asset selling; // A
//        Asset buying;  // B
//        int64 amount;  // amount of A
//
//        /* price for this offer:
//            price of A in terms of B
//            price=AmountB/AmountA=priceNumerator/priceDenominator
//            price is after fees
//        */
//        Price price;
//        uint32 flags; // see OfferEntryFlags
//
//        // reserved for future use
//        union switch (int v)
//        {
//        case 0:
//            void;
//        }
//        ext;
//    };
//
type OfferEntry struct {
	SellerId AccountId
	OfferId  Int64
	Selling  Asset
	Buying   Asset
	Amount   Int64
	Price    Price
	Flags    Uint32
	Ext      OfferEntryExt
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s OfferEntry) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *OfferEntry) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*OfferEntry)(nil)
	_ encoding.BinaryUnmarshaler = (*OfferEntry)(nil)
)

// DataEntryExt is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        }
//
type DataEntryExt struct {
	V int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u DataEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of DataEntryExt
func (u DataEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	}
	return "-", false
}

// NewDataEntryExt creates a new  DataEntryExt.
func NewDataEntryExt(v int32, value interface{}) (result DataEntryExt, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s DataEntryExt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *DataEntryExt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*DataEntryExt)(nil)
	_ encoding.BinaryUnmarshaler = (*DataEntryExt)(nil)
)

// DataEntry is an XDR Struct defines as:
//
//   struct DataEntry
//    {
//        AccountID accountID; // account this data belongs to
//        string64 dataName;
//        DataValue dataValue;
//
//        // reserved for future use
//        union switch (int v)
//        {
//        case 0:
//            void;
//        }
//        ext;
//    };
//
type DataEntry struct {
	AccountId AccountId
	DataName  String64
	DataValue DataValue
	Ext       DataEntryExt
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s DataEntry) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *DataEntry) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*DataEntry)(nil)
	_ encoding.BinaryUnmarshaler = (*DataEntry)(nil)
)

// ClaimPredicateType is an XDR Enum defines as:
//
//   enum ClaimPredicateType
//    {
//        CLAIM_PREDICATE_UNCONDITIONAL = 0,
//        CLAIM_PREDICATE_AND = 1,
//        CLAIM_PREDICATE_OR = 2,
//        CLAIM_PREDICATE_NOT = 3,
//        CLAIM_PREDICATE_BEFORE_ABSOLUTE_TIME = 4,
//        CLAIM_PREDICATE_BEFORE_RELATIVE_TIME = 5
//    };
//
type ClaimPredicateType int32

const (
	ClaimPredicateTypeClaimPredicateUnconditional      ClaimPredicateType = 0
	ClaimPredicateTypeClaimPredicateAnd                ClaimPredicateType = 1
	ClaimPredicateTypeClaimPredicateOr                 ClaimPredicateType = 2
	ClaimPredicateTypeClaimPredicateNot                ClaimPredicateType = 3
	ClaimPredicateTypeClaimPredicateBeforeAbsoluteTime ClaimPredicateType = 4
	ClaimPredicateTypeClaimPredicateBeforeRelativeTime ClaimPredicateType = 5
)

var claimPredicateTypeMap = map[int32]string{
	0: "ClaimPredicateTypeClaimPredicateUnconditional",
	1: "ClaimPredicateTypeClaimPredicateAnd",
	2: "ClaimPredicateTypeClaimPredicateOr",
	3: "ClaimPredicateTypeClaimPredicateNot",
	4: "ClaimPredicateTypeClaimPredicateBeforeAbsoluteTime",
	5: "ClaimPredicateTypeClaimPredicateBeforeRelativeTime",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ClaimPredicateType
func (e ClaimPredicateType) ValidEnum(v int32) bool {
	_, ok := claimPredicateTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e ClaimPredicateType) String() string {
	name, _ := claimPredicateTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimPredicateType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimPredicateType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimPredicateType)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimPredicateType)(nil)
)

// ClaimPredicate is an XDR Union defines as:
//
//   union ClaimPredicate switch (ClaimPredicateType type)
//    {
//    case CLAIM_PREDICATE_UNCONDITIONAL:
//        void;
//    case CLAIM_PREDICATE_AND:
//        ClaimPredicate andPredicates<2>;
//    case CLAIM_PREDICATE_OR:
//        ClaimPredicate orPredicates<2>;
//    case CLAIM_PREDICATE_NOT:
//        ClaimPredicate* notPredicate;
//    case CLAIM_PREDICATE_BEFORE_ABSOLUTE_TIME:
//        int64 absBefore; // Predicate will be true if closeTime < absBefore
//    case CLAIM_PREDICATE_BEFORE_RELATIVE_TIME:
//        int64 relBefore; // Seconds since closeTime of the ledger in which the
//                         // ClaimableBalanceEntry was created
//    };
//
type ClaimPredicate struct {
	Type          ClaimPredicateType
	AndPredicates *[]ClaimPredicate `xdrmaxsize:"2"`
	OrPredicates  *[]ClaimPredicate `xdrmaxsize:"2"`
	NotPredicate  **ClaimPredicate
	AbsBefore     *Int64
	RelBefore     *Int64
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ClaimPredicate) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ClaimPredicate
func (u ClaimPredicate) ArmForSwitch(sw int32) (string, bool) {
	switch ClaimPredicateType(sw) {
	case ClaimPredicateTypeClaimPredicateUnconditional:
		return "", true
	case ClaimPredicateTypeClaimPredicateAnd:
		return "AndPredicates", true
	case ClaimPredicateTypeClaimPredicateOr:
		return "OrPredicates", true
	case ClaimPredicateTypeClaimPredicateNot:
		return "NotPredicate", true
	case ClaimPredicateTypeClaimPredicateBeforeAbsoluteTime:
		return "AbsBefore", true
	case ClaimPredicateTypeClaimPredicateBeforeRelativeTime:
		return "RelBefore", true
	}
	return "-", false
}

// NewClaimPredicate creates a new  ClaimPredicate.
func NewClaimPredicate(aType ClaimPredicateType, value interface{}) (result ClaimPredicate, err error) {
	result.Type = aType
	switch ClaimPredicateType(aType) {
	case ClaimPredicateTypeClaimPredicateUnconditional:
		// void
	case ClaimPredicateTypeClaimPredicateAnd:
		tv, ok := value.([]ClaimPredicate)
		if !ok {
			err = fmt.Errorf("invalid value, must be []ClaimPredicate")
			return
		}
		result.AndPredicates = &tv
	case ClaimPredicateTypeClaimPredicateOr:
		tv, ok := value.([]ClaimPredicate)
		if !ok {
			err = fmt.Errorf("invalid value, must be []ClaimPredicate")
			return
		}
		result.OrPredicates = &tv
	case ClaimPredicateTypeClaimPredicateNot:
		tv, ok := value.(*ClaimPredicate)
		if !ok {
			err = fmt.Errorf("invalid value, must be *ClaimPredicate")
			return
		}
		result.NotPredicate = &tv
	case ClaimPredicateTypeClaimPredicateBeforeAbsoluteTime:
		tv, ok := value.(Int64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Int64")
			return
		}
		result.AbsBefore = &tv
	case ClaimPredicateTypeClaimPredicateBeforeRelativeTime:
		tv, ok := value.(Int64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Int64")
			return
		}
		result.RelBefore = &tv
	}
	return
}

// MustAndPredicates retrieves the AndPredicates value from the union,
// panicing if the value is not set.
func (u ClaimPredicate) MustAndPredicates() []ClaimPredicate {
	val, ok := u.GetAndPredicates()

	if !ok {
		panic("arm AndPredicates is not set")
	}

	return val
}

// GetAndPredicates retrieves the AndPredicates value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ClaimPredicate) GetAndPredicates() (result []ClaimPredicate, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AndPredicates" {
		result = *u.AndPredicates
		ok = true
	}

	return
}

// MustOrPredicates retrieves the OrPredicates value from the union,
// panicing if the value is not set.
func (u ClaimPredicate) MustOrPredicates() []ClaimPredicate {
	val, ok := u.GetOrPredicates()

	if !ok {
		panic("arm OrPredicates is not set")
	}

	return val
}

// GetOrPredicates retrieves the OrPredicates value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ClaimPredicate) GetOrPredicates() (result []ClaimPredicate, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "OrPredicates" {
		result = *u.OrPredicates
		ok = true
	}

	return
}

// MustNotPredicate retrieves the NotPredicate value from the union,
// panicing if the value is not set.
func (u ClaimPredicate) MustNotPredicate() *ClaimPredicate {
	val, ok := u.GetNotPredicate()

	if !ok {
		panic("arm NotPredicate is not set")
	}

	return val
}

// GetNotPredicate retrieves the NotPredicate value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ClaimPredicate) GetNotPredicate() (result *ClaimPredicate, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "NotPredicate" {
		result = *u.NotPredicate
		ok = true
	}

	return
}

// MustAbsBefore retrieves the AbsBefore value from the union,
// panicing if the value is not set.
func (u ClaimPredicate) MustAbsBefore() Int64 {
	val, ok := u.GetAbsBefore()

	if !ok {
		panic("arm AbsBefore is not set")
	}

	return val
}

// GetAbsBefore retrieves the AbsBefore value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ClaimPredicate) GetAbsBefore() (result Int64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AbsBefore" {
		result = *u.AbsBefore
		ok = true
	}

	return
}

// MustRelBefore retrieves the RelBefore value from the union,
// panicing if the value is not set.
func (u ClaimPredicate) MustRelBefore() Int64 {
	val, ok := u.GetRelBefore()

	if !ok {
		panic("arm RelBefore is not set")
	}

	return val
}

// GetRelBefore retrieves the RelBefore value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ClaimPredicate) GetRelBefore() (result Int64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RelBefore" {
		result = *u.RelBefore
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimPredicate) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimPredicate) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimPredicate)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimPredicate)(nil)
)

// ClaimantType is an XDR Enum defines as:
//
//   enum ClaimantType
//    {
//        CLAIMANT_TYPE_V0 = 0
//    };
//
type ClaimantType int32

const (
	ClaimantTypeClaimantTypeV0 ClaimantType = 0
)

var claimantTypeMap = map[int32]string{
	0: "ClaimantTypeClaimantTypeV0",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ClaimantType
func (e ClaimantType) ValidEnum(v int32) bool {
	_, ok := claimantTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e ClaimantType) String() string {
	name, _ := claimantTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimantType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimantType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimantType)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimantType)(nil)
)

// ClaimantV0 is an XDR NestedStruct defines as:
//
//   struct
//        {
//            AccountID destination;    // The account that can use this condition
//            ClaimPredicate predicate; // Claimable if predicate is true
//        }
//
type ClaimantV0 struct {
	Destination AccountId
	Predicate   ClaimPredicate
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimantV0) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimantV0) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimantV0)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimantV0)(nil)
)

// Claimant is an XDR Union defines as:
//
//   union Claimant switch (ClaimantType type)
//    {
//    case CLAIMANT_TYPE_V0:
//        struct
//        {
//            AccountID destination;    // The account that can use this condition
//            ClaimPredicate predicate; // Claimable if predicate is true
//        } v0;
//    };
//
type Claimant struct {
	Type ClaimantType
	V0   *ClaimantV0
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u Claimant) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of Claimant
func (u Claimant) ArmForSwitch(sw int32) (string, bool) {
	switch ClaimantType(sw) {
	case ClaimantTypeClaimantTypeV0:
		return "V0", true
	}
	return "-", false
}

// NewClaimant creates a new  Claimant.
func NewClaimant(aType ClaimantType, value interface{}) (result Claimant, err error) {
	result.Type = aType
	switch ClaimantType(aType) {
	case ClaimantTypeClaimantTypeV0:
		tv, ok := value.(ClaimantV0)
		if !ok {
			err = fmt.Errorf("invalid value, must be ClaimantV0")
			return
		}
		result.V0 = &tv
	}
	return
}

// MustV0 retrieves the V0 value from the union,
// panicing if the value is not set.
func (u Claimant) MustV0() ClaimantV0 {
	val, ok := u.GetV0()

	if !ok {
		panic("arm V0 is not set")
	}

	return val
}

// GetV0 retrieves the V0 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Claimant) GetV0() (result ClaimantV0, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "V0" {
		result = *u.V0
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Claimant) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Claimant) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Claimant)(nil)
	_ encoding.BinaryUnmarshaler = (*Claimant)(nil)
)

// ClaimableBalanceIdType is an XDR Enum defines as:
//
//   enum ClaimableBalanceIDType
//    {
//        CLAIMABLE_BALANCE_ID_TYPE_V0 = 0
//    };
//
type ClaimableBalanceIdType int32

const (
	ClaimableBalanceIdTypeClaimableBalanceIdTypeV0 ClaimableBalanceIdType = 0
)

var claimableBalanceIdTypeMap = map[int32]string{
	0: "ClaimableBalanceIdTypeClaimableBalanceIdTypeV0",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ClaimableBalanceIdType
func (e ClaimableBalanceIdType) ValidEnum(v int32) bool {
	_, ok := claimableBalanceIdTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e ClaimableBalanceIdType) String() string {
	name, _ := claimableBalanceIdTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimableBalanceIdType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimableBalanceIdType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimableBalanceIdType)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimableBalanceIdType)(nil)
)

// ClaimableBalanceId is an XDR Union defines as:
//
//   union ClaimableBalanceID switch (ClaimableBalanceIDType type)
//    {
//    case CLAIMABLE_BALANCE_ID_TYPE_V0:
//        Hash v0;
//    };
//
type ClaimableBalanceId struct {
	Type ClaimableBalanceIdType
	V0   *Hash
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ClaimableBalanceId) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ClaimableBalanceId
func (u ClaimableBalanceId) ArmForSwitch(sw int32) (string, bool) {
	switch ClaimableBalanceIdType(sw) {
	case ClaimableBalanceIdTypeClaimableBalanceIdTypeV0:
		return "V0", true
	}
	return "-", false
}

// NewClaimableBalanceId creates a new  ClaimableBalanceId.
func NewClaimableBalanceId(aType ClaimableBalanceIdType, value interface{}) (result ClaimableBalanceId, err error) {
	result.Type = aType
	switch ClaimableBalanceIdType(aType) {
	case ClaimableBalanceIdTypeClaimableBalanceIdTypeV0:
		tv, ok := value.(Hash)
		if !ok {
			err = fmt.Errorf("invalid value, must be Hash")
			return
		}
		result.V0 = &tv
	}
	return
}

// MustV0 retrieves the V0 value from the union,
// panicing if the value is not set.
func (u ClaimableBalanceId) MustV0() Hash {
	val, ok := u.GetV0()

	if !ok {
		panic("arm V0 is not set")
	}

	return val
}

// GetV0 retrieves the V0 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ClaimableBalanceId) GetV0() (result Hash, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "V0" {
		result = *u.V0
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimableBalanceId) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimableBalanceId) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimableBalanceId)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimableBalanceId)(nil)
)

// ClaimableBalanceFlags is an XDR Enum defines as:
//
//   enum ClaimableBalanceFlags
//    {
//        // If set, the issuer account of the asset held by the claimable balance may
//        // clawback the claimable balance
//        CLAIMABLE_BALANCE_CLAWBACK_ENABLED_FLAG = 0x1
//    };
//
type ClaimableBalanceFlags int32

const (
	ClaimableBalanceFlagsClaimableBalanceClawbackEnabledFlag ClaimableBalanceFlags = 1
)

var claimableBalanceFlagsMap = map[int32]string{
	1: "ClaimableBalanceFlagsClaimableBalanceClawbackEnabledFlag",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ClaimableBalanceFlags
func (e ClaimableBalanceFlags) ValidEnum(v int32) bool {
	_, ok := claimableBalanceFlagsMap[v]
	return ok
}

// String returns the name of `e`
func (e ClaimableBalanceFlags) String() string {
	name, _ := claimableBalanceFlagsMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimableBalanceFlags) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimableBalanceFlags) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimableBalanceFlags)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimableBalanceFlags)(nil)
)

// MaskClaimableBalanceFlags is an XDR Const defines as:
//
//   const MASK_CLAIMABLE_BALANCE_FLAGS = 0x1;
//
const MaskClaimableBalanceFlags = 0x1

// ClaimableBalanceEntryExtensionV1Ext is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        }
//
type ClaimableBalanceEntryExtensionV1Ext struct {
	V int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ClaimableBalanceEntryExtensionV1Ext) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ClaimableBalanceEntryExtensionV1Ext
func (u ClaimableBalanceEntryExtensionV1Ext) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	}
	return "-", false
}

// NewClaimableBalanceEntryExtensionV1Ext creates a new  ClaimableBalanceEntryExtensionV1Ext.
func NewClaimableBalanceEntryExtensionV1Ext(v int32, value interface{}) (result ClaimableBalanceEntryExtensionV1Ext, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimableBalanceEntryExtensionV1Ext) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimableBalanceEntryExtensionV1Ext) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimableBalanceEntryExtensionV1Ext)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimableBalanceEntryExtensionV1Ext)(nil)
)

// ClaimableBalanceEntryExtensionV1 is an XDR Struct defines as:
//
//   struct ClaimableBalanceEntryExtensionV1
//    {
//        union switch (int v)
//        {
//        case 0:
//            void;
//        }
//        ext;
//
//        uint32 flags; // see ClaimableBalanceFlags
//    };
//
type ClaimableBalanceEntryExtensionV1 struct {
	Ext   ClaimableBalanceEntryExtensionV1Ext
	Flags Uint32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimableBalanceEntryExtensionV1) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimableBalanceEntryExtensionV1) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimableBalanceEntryExtensionV1)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimableBalanceEntryExtensionV1)(nil)
)

// ClaimableBalanceEntryExt is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        case 1:
//            ClaimableBalanceEntryExtensionV1 v1;
//        }
//
type ClaimableBalanceEntryExt struct {
	V  int32
	V1 *ClaimableBalanceEntryExtensionV1
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ClaimableBalanceEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ClaimableBalanceEntryExt
func (u ClaimableBalanceEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	case 1:
		return "V1", true
	}
	return "-", false
}

// NewClaimableBalanceEntryExt creates a new  ClaimableBalanceEntryExt.
func NewClaimableBalanceEntryExt(v int32, value interface{}) (result ClaimableBalanceEntryExt, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	case 1:
		tv, ok := value.(ClaimableBalanceEntryExtensionV1)
		if !ok {
			err = fmt.Errorf("invalid value, must be ClaimableBalanceEntryExtensionV1")
			return
		}
		result.V1 = &tv
	}
	return
}

// MustV1 retrieves the V1 value from the union,
// panicing if the value is not set.
func (u ClaimableBalanceEntryExt) MustV1() ClaimableBalanceEntryExtensionV1 {
	val, ok := u.GetV1()

	if !ok {
		panic("arm V1 is not set")
	}

	return val
}

// GetV1 retrieves the V1 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ClaimableBalanceEntryExt) GetV1() (result ClaimableBalanceEntryExtensionV1, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "V1" {
		result = *u.V1
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimableBalanceEntryExt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimableBalanceEntryExt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimableBalanceEntryExt)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimableBalanceEntryExt)(nil)
)

// ClaimableBalanceEntry is an XDR Struct defines as:
//
//   struct ClaimableBalanceEntry
//    {
//        // Unique identifier for this ClaimableBalanceEntry
//        ClaimableBalanceID balanceID;
//
//        // List of claimants with associated predicate
//        Claimant claimants<10>;
//
//        // Any asset including native
//        Asset asset;
//
//        // Amount of asset
//        int64 amount;
//
//        // reserved for future use
//        union switch (int v)
//        {
//        case 0:
//            void;
//        case 1:
//            ClaimableBalanceEntryExtensionV1 v1;
//        }
//        ext;
//    };
//
type ClaimableBalanceEntry struct {
	BalanceId ClaimableBalanceId
	Claimants []Claimant `xdrmaxsize:"10"`
	Asset     Asset
	Amount    Int64
	Ext       ClaimableBalanceEntryExt
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimableBalanceEntry) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimableBalanceEntry) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimableBalanceEntry)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimableBalanceEntry)(nil)
)

// LiquidityPoolConstantProductParameters is an XDR Struct defines as:
//
//   struct LiquidityPoolConstantProductParameters
//    {
//        Asset assetA; // assetA < assetB
//        Asset assetB;
//        int32 fee;    // Fee is in basis points, so the actual rate is (fee/100)%
//    };
//
type LiquidityPoolConstantProductParameters struct {
	AssetA Asset
	AssetB Asset
	Fee    Int32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LiquidityPoolConstantProductParameters) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LiquidityPoolConstantProductParameters) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LiquidityPoolConstantProductParameters)(nil)
	_ encoding.BinaryUnmarshaler = (*LiquidityPoolConstantProductParameters)(nil)
)

// LiquidityPoolEntryConstantProduct is an XDR NestedStruct defines as:
//
//   struct
//            {
//                LiquidityPoolConstantProductParameters params;
//
//                int64 reserveA;        // amount of A in the pool
//                int64 reserveB;        // amount of B in the pool
//                int64 totalPoolShares; // total number of pool shares issued
//                int64 poolSharesTrustLineCount; // number of trust lines for the associated pool shares
//            }
//
type LiquidityPoolEntryConstantProduct struct {
	Params                   LiquidityPoolConstantProductParameters
	ReserveA                 Int64
	ReserveB                 Int64
	TotalPoolShares          Int64
	PoolSharesTrustLineCount Int64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LiquidityPoolEntryConstantProduct) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LiquidityPoolEntryConstantProduct) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LiquidityPoolEntryConstantProduct)(nil)
	_ encoding.BinaryUnmarshaler = (*LiquidityPoolEntryConstantProduct)(nil)
)

// LiquidityPoolEntryBody is an XDR NestedUnion defines as:
//
//   union switch (LiquidityPoolType type)
//        {
//        case LIQUIDITY_POOL_CONSTANT_PRODUCT:
//            struct
//            {
//                LiquidityPoolConstantProductParameters params;
//
//                int64 reserveA;        // amount of A in the pool
//                int64 reserveB;        // amount of B in the pool
//                int64 totalPoolShares; // total number of pool shares issued
//                int64 poolSharesTrustLineCount; // number of trust lines for the associated pool shares
//            } constantProduct;
//        }
//
type LiquidityPoolEntryBody struct {
	Type            LiquidityPoolType
	ConstantProduct *LiquidityPoolEntryConstantProduct
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LiquidityPoolEntryBody) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LiquidityPoolEntryBody
func (u LiquidityPoolEntryBody) ArmForSwitch(sw int32) (string, bool) {
	switch LiquidityPoolType(sw) {
	case LiquidityPoolTypeLiquidityPoolConstantProduct:
		return "ConstantProduct", true
	}
	return "-", false
}

// NewLiquidityPoolEntryBody creates a new  LiquidityPoolEntryBody.
func NewLiquidityPoolEntryBody(aType LiquidityPoolType, value interface{}) (result LiquidityPoolEntryBody, err error) {
	result.Type = aType
	switch LiquidityPoolType(aType) {
	case LiquidityPoolTypeLiquidityPoolConstantProduct:
		tv, ok := value.(LiquidityPoolEntryConstantProduct)
		if !ok {
			err = fmt.Errorf("invalid value, must be LiquidityPoolEntryConstantProduct")
			return
		}
		result.ConstantProduct = &tv
	}
	return
}

// MustConstantProduct retrieves the ConstantProduct value from the union,
// panicing if the value is not set.
func (u LiquidityPoolEntryBody) MustConstantProduct() LiquidityPoolEntryConstantProduct {
	val, ok := u.GetConstantProduct()

	if !ok {
		panic("arm ConstantProduct is not set")
	}

	return val
}

// GetConstantProduct retrieves the ConstantProduct value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LiquidityPoolEntryBody) GetConstantProduct() (result LiquidityPoolEntryConstantProduct, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ConstantProduct" {
		result = *u.ConstantProduct
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LiquidityPoolEntryBody) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LiquidityPoolEntryBody) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LiquidityPoolEntryBody)(nil)
	_ encoding.BinaryUnmarshaler = (*LiquidityPoolEntryBody)(nil)
)

// LiquidityPoolEntry is an XDR Struct defines as:
//
//   struct LiquidityPoolEntry
//    {
//        PoolID liquidityPoolID;
//
//        union switch (LiquidityPoolType type)
//        {
//        case LIQUIDITY_POOL_CONSTANT_PRODUCT:
//            struct
//            {
//                LiquidityPoolConstantProductParameters params;
//
//                int64 reserveA;        // amount of A in the pool
//                int64 reserveB;        // amount of B in the pool
//                int64 totalPoolShares; // total number of pool shares issued
//                int64 poolSharesTrustLineCount; // number of trust lines for the associated pool shares
//            } constantProduct;
//        }
//        body;
//    };
//
type LiquidityPoolEntry struct {
	LiquidityPoolId PoolId
	Body            LiquidityPoolEntryBody
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LiquidityPoolEntry) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LiquidityPoolEntry) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LiquidityPoolEntry)(nil)
	_ encoding.BinaryUnmarshaler = (*LiquidityPoolEntry)(nil)
)

// LedgerEntryExtensionV1Ext is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        }
//
type LedgerEntryExtensionV1Ext struct {
	V int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerEntryExtensionV1Ext) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerEntryExtensionV1Ext
func (u LedgerEntryExtensionV1Ext) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	}
	return "-", false
}

// NewLedgerEntryExtensionV1Ext creates a new  LedgerEntryExtensionV1Ext.
func NewLedgerEntryExtensionV1Ext(v int32, value interface{}) (result LedgerEntryExtensionV1Ext, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerEntryExtensionV1Ext) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerEntryExtensionV1Ext) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerEntryExtensionV1Ext)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerEntryExtensionV1Ext)(nil)
)

// LedgerEntryExtensionV1 is an XDR Struct defines as:
//
//   struct LedgerEntryExtensionV1
//    {
//        SponsorshipDescriptor sponsoringID;
//
//        union switch (int v)
//        {
//        case 0:
//            void;
//        }
//        ext;
//    };
//
type LedgerEntryExtensionV1 struct {
	SponsoringId SponsorshipDescriptor
	Ext          LedgerEntryExtensionV1Ext
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerEntryExtensionV1) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerEntryExtensionV1) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerEntryExtensionV1)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerEntryExtensionV1)(nil)
)

// LedgerEntryData is an XDR NestedUnion defines as:
//
//   union switch (LedgerEntryType type)
//        {
//        case ACCOUNT:
//            AccountEntry account;
//        case TRUSTLINE:
//            TrustLineEntry trustLine;
//        case OFFER:
//            OfferEntry offer;
//        case DATA:
//            DataEntry data;
//        case CLAIMABLE_BALANCE:
//            ClaimableBalanceEntry claimableBalance;
//        case LIQUIDITY_POOL:
//            LiquidityPoolEntry liquidityPool;
//        }
//
type LedgerEntryData struct {
	Type             LedgerEntryType
	Account          *AccountEntry
	TrustLine        *TrustLineEntry
	Offer            *OfferEntry
	Data             *DataEntry
	ClaimableBalance *ClaimableBalanceEntry
	LiquidityPool    *LiquidityPoolEntry
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerEntryData) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerEntryData
func (u LedgerEntryData) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerEntryType(sw) {
	case LedgerEntryTypeAccount:
		return "Account", true
	case LedgerEntryTypeTrustline:
		return "TrustLine", true
	case LedgerEntryTypeOffer:
		return "Offer", true
	case LedgerEntryTypeData:
		return "Data", true
	case LedgerEntryTypeClaimableBalance:
		return "ClaimableBalance", true
	case LedgerEntryTypeLiquidityPool:
		return "LiquidityPool", true
	}
	return "-", false
}

// NewLedgerEntryData creates a new  LedgerEntryData.
func NewLedgerEntryData(aType LedgerEntryType, value interface{}) (result LedgerEntryData, err error) {
	result.Type = aType
	switch LedgerEntryType(aType) {
	case LedgerEntryTypeAccount:
		tv, ok := value.(AccountEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountEntry")
			return
		}
		result.Account = &tv
	case LedgerEntryTypeTrustline:
		tv, ok := value.(TrustLineEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be TrustLineEntry")
			return
		}
		result.TrustLine = &tv
	case LedgerEntryTypeOffer:
		tv, ok := value.(OfferEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be OfferEntry")
			return
		}
		result.Offer = &tv
	case LedgerEntryTypeData:
		tv, ok := value.(DataEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be DataEntry")
			return
		}
		result.Data = &tv
	case LedgerEntryTypeClaimableBalance:
		tv, ok := value.(ClaimableBalanceEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be ClaimableBalanceEntry")
			return
		}
		result.ClaimableBalance = &tv
	case LedgerEntryTypeLiquidityPool:
		tv, ok := value.(LiquidityPoolEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LiquidityPoolEntry")
			return
		}
		result.LiquidityPool = &tv
	}
	return
}

// MustAccount retrieves the Account value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustAccount() AccountEntry {
	val, ok := u.GetAccount()

	if !ok {
		panic("arm Account is not set")
	}

	return val
}

// GetAccount retrieves the Account value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetAccount() (result AccountEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Account" {
		result = *u.Account
		ok = true
	}

	return
}

// MustTrustLine retrieves the TrustLine value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustTrustLine() TrustLineEntry {
	val, ok := u.GetTrustLine()

	if !ok {
		panic("arm TrustLine is not set")
	}

	return val
}

// GetTrustLine retrieves the TrustLine value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetTrustLine() (result TrustLineEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "TrustLine" {
		result = *u.TrustLine
		ok = true
	}

	return
}

// MustOffer retrieves the Offer value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustOffer() OfferEntry {
	val, ok := u.GetOffer()

	if !ok {
		panic("arm Offer is not set")
	}

	return val
}

// GetOffer retrieves the Offer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetOffer() (result OfferEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Offer" {
		result = *u.Offer
		ok = true
	}

	return
}

// MustData retrieves the Data value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustData() DataEntry {
	val, ok := u.GetData()

	if !ok {
		panic("arm Data is not set")
	}

	return val
}

// GetData retrieves the Data value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetData() (result DataEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Data" {
		result = *u.Data
		ok = true
	}

	return
}

// MustClaimableBalance retrieves the ClaimableBalance value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustClaimableBalance() ClaimableBalanceEntry {
	val, ok := u.GetClaimableBalance()

	if !ok {
		panic("arm ClaimableBalance is not set")
	}

	return val
}

// GetClaimableBalance retrieves the ClaimableBalance value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetClaimableBalance() (result ClaimableBalanceEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ClaimableBalance" {
		result = *u.ClaimableBalance
		ok = true
	}

	return
}

// MustLiquidityPool retrieves the LiquidityPool value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustLiquidityPool() LiquidityPoolEntry {
	val, ok := u.GetLiquidityPool()

	if !ok {
		panic("arm LiquidityPool is not set")
	}

	return val
}

// GetLiquidityPool retrieves the LiquidityPool value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetLiquidityPool() (result LiquidityPoolEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "LiquidityPool" {
		result = *u.LiquidityPool
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerEntryData) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerEntryData) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerEntryData)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerEntryData)(nil)
)

// LedgerEntryExt is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        case 1:
//            LedgerEntryExtensionV1 v1;
//        }
//
type LedgerEntryExt struct {
	V  int32
	V1 *LedgerEntryExtensionV1
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerEntryExt
func (u LedgerEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	case 1:
		return "V1", true
	}
	return "-", false
}

// NewLedgerEntryExt creates a new  LedgerEntryExt.
func NewLedgerEntryExt(v int32, value interface{}) (result LedgerEntryExt, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	case 1:
		tv, ok := value.(LedgerEntryExtensionV1)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntryExtensionV1")
			return
		}
		result.V1 = &tv
	}
	return
}

// MustV1 retrieves the V1 value from the union,
// panicing if the value is not set.
func (u LedgerEntryExt) MustV1() LedgerEntryExtensionV1 {
	val, ok := u.GetV1()

	if !ok {
		panic("arm V1 is not set")
	}

	return val
}

// GetV1 retrieves the V1 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryExt) GetV1() (result LedgerEntryExtensionV1, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "V1" {
		result = *u.V1
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerEntryExt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerEntryExt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerEntryExt)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerEntryExt)(nil)
)

// LedgerEntry is an XDR Struct defines as:
//
//   struct LedgerEntry
//    {
//        uint32 lastModifiedLedgerSeq; // ledger the LedgerEntry was last changed
//
//        union switch (LedgerEntryType type)
//        {
//        case ACCOUNT:
//            AccountEntry account;
//        case TRUSTLINE:
//            TrustLineEntry trustLine;
//        case OFFER:
//            OfferEntry offer;
//        case DATA:
//            DataEntry data;
//        case CLAIMABLE_BALANCE:
//            ClaimableBalanceEntry claimableBalance;
//        case LIQUIDITY_POOL:
//            LiquidityPoolEntry liquidityPool;
//        }
//        data;
//
//        // reserved for future use
//        union switch (int v)
//        {
//        case 0:
//            void;
//        case 1:
//            LedgerEntryExtensionV1 v1;
//        }
//        ext;
//    };
//
type LedgerEntry struct {
	LastModifiedLedgerSeq Uint32
	Data                  LedgerEntryData
	Ext                   LedgerEntryExt
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerEntry) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerEntry) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerEntry)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerEntry)(nil)
)

// LedgerKeyAccount is an XDR NestedStruct defines as:
//
//   struct
//        {
//            AccountID accountID;
//        }
//
type LedgerKeyAccount struct {
	AccountId AccountId
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerKeyAccount) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerKeyAccount) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerKeyAccount)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerKeyAccount)(nil)
)

// LedgerKeyTrustLine is an XDR NestedStruct defines as:
//
//   struct
//        {
//            AccountID accountID;
//            TrustLineAsset asset;
//        }
//
type LedgerKeyTrustLine struct {
	AccountId AccountId
	Asset     TrustLineAsset
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerKeyTrustLine) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerKeyTrustLine) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerKeyTrustLine)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerKeyTrustLine)(nil)
)

// LedgerKeyOffer is an XDR NestedStruct defines as:
//
//   struct
//        {
//            AccountID sellerID;
//            int64 offerID;
//        }
//
type LedgerKeyOffer struct {
	SellerId AccountId
	OfferId  Int64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerKeyOffer) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerKeyOffer) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerKeyOffer)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerKeyOffer)(nil)
)

// LedgerKeyData is an XDR NestedStruct defines as:
//
//   struct
//        {
//            AccountID accountID;
//            string64 dataName;
//        }
//
type LedgerKeyData struct {
	AccountId AccountId
	DataName  String64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerKeyData) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerKeyData) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerKeyData)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerKeyData)(nil)
)

// LedgerKeyClaimableBalance is an XDR NestedStruct defines as:
//
//   struct
//        {
//            ClaimableBalanceID balanceID;
//        }
//
type LedgerKeyClaimableBalance struct {
	BalanceId ClaimableBalanceId
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerKeyClaimableBalance) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerKeyClaimableBalance) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerKeyClaimableBalance)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerKeyClaimableBalance)(nil)
)

// LedgerKeyLiquidityPool is an XDR NestedStruct defines as:
//
//   struct
//        {
//            PoolID liquidityPoolID;
//        }
//
type LedgerKeyLiquidityPool struct {
	LiquidityPoolId PoolId
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerKeyLiquidityPool) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerKeyLiquidityPool) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerKeyLiquidityPool)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerKeyLiquidityPool)(nil)
)

// LedgerKey is an XDR Union defines as:
//
//   union LedgerKey switch (LedgerEntryType type)
//    {
//    case ACCOUNT:
//        struct
//        {
//            AccountID accountID;
//        } account;
//
//    case TRUSTLINE:
//        struct
//        {
//            AccountID accountID;
//            TrustLineAsset asset;
//        } trustLine;
//
//    case OFFER:
//        struct
//        {
//            AccountID sellerID;
//            int64 offerID;
//        } offer;
//
//    case DATA:
//        struct
//        {
//            AccountID accountID;
//            string64 dataName;
//        } data;
//
//    case CLAIMABLE_BALANCE:
//        struct
//        {
//            ClaimableBalanceID balanceID;
//        } claimableBalance;
//
//    case LIQUIDITY_POOL:
//        struct
//        {
//            PoolID liquidityPoolID;
//        } liquidityPool;
//    };
//
type LedgerKey struct {
	Type             LedgerEntryType
	Account          *LedgerKeyAccount
	TrustLine        *LedgerKeyTrustLine
	Offer            *LedgerKeyOffer
	Data             *LedgerKeyData
	ClaimableBalance *LedgerKeyClaimableBalance
	LiquidityPool    *LedgerKeyLiquidityPool
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKey) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKey
func (u LedgerKey) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerEntryType(sw) {
	case LedgerEntryTypeAccount:
		return "Account", true
	case LedgerEntryTypeTrustline:
		return "TrustLine", true
	case LedgerEntryTypeOffer:
		return "Offer", true
	case LedgerEntryTypeData:
		return "Data", true
	case LedgerEntryTypeClaimableBalance:
		return "ClaimableBalance", true
	case LedgerEntryTypeLiquidityPool:
		return "LiquidityPool", true
	}
	return "-", false
}

// NewLedgerKey creates a new  LedgerKey.
func NewLedgerKey(aType LedgerEntryType, value interface{}) (result LedgerKey, err error) {
	result.Type = aType
	switch LedgerEntryType(aType) {
	case LedgerEntryTypeAccount:
		tv, ok := value.(LedgerKeyAccount)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyAccount")
			return
		}
		result.Account = &tv
	case LedgerEntryTypeTrustline:
		tv, ok := value.(LedgerKeyTrustLine)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyTrustLine")
			return
		}
		result.TrustLine = &tv
	case LedgerEntryTypeOffer:
		tv, ok := value.(LedgerKeyOffer)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyOffer")
			return
		}
		result.Offer = &tv
	case LedgerEntryTypeData:
		tv, ok := value.(LedgerKeyData)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyData")
			return
		}
		result.Data = &tv
	case LedgerEntryTypeClaimableBalance:
		tv, ok := value.(LedgerKeyClaimableBalance)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyClaimableBalance")
			return
		}
		result.ClaimableBalance = &tv
	case LedgerEntryTypeLiquidityPool:
		tv, ok := value.(LedgerKeyLiquidityPool)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyLiquidityPool")
			return
		}
		result.LiquidityPool = &tv
	}
	return
}

// MustAccount retrieves the Account value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustAccount() LedgerKeyAccount {
	val, ok := u.GetAccount()

	if !ok {
		panic("arm Account is not set")
	}

	return val
}

// GetAccount retrieves the Account value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetAccount() (result LedgerKeyAccount, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Account" {
		result = *u.Account
		ok = true
	}

	return
}

// MustTrustLine retrieves the TrustLine value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustTrustLine() LedgerKeyTrustLine {
	val, ok := u.GetTrustLine()

	if !ok {
		panic("arm TrustLine is not set")
	}

	return val
}

// GetTrustLine retrieves the TrustLine value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetTrustLine() (result LedgerKeyTrustLine, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "TrustLine" {
		result = *u.TrustLine
		ok = true
	}

	return
}

// MustOffer retrieves the Offer value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustOffer() LedgerKeyOffer {
	val, ok := u.GetOffer()

	if !ok {
		panic("arm Offer is not set")
	}

	return val
}

// GetOffer retrieves the Offer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetOffer() (result LedgerKeyOffer, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Offer" {
		result = *u.Offer
		ok = true
	}

	return
}

// MustData retrieves the Data value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustData() LedgerKeyData {
	val, ok := u.GetData()

	if !ok {
		panic("arm Data is not set")
	}

	return val
}

// GetData retrieves the Data value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetData() (result LedgerKeyData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Data" {
		result = *u.Data
		ok = true
	}

	return
}

// MustClaimableBalance retrieves the ClaimableBalance value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustClaimableBalance() LedgerKeyClaimableBalance {
	val, ok := u.GetClaimableBalance()

	if !ok {
		panic("arm ClaimableBalance is not set")
	}

	return val
}

// GetClaimableBalance retrieves the ClaimableBalance value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetClaimableBalance() (result LedgerKeyClaimableBalance, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ClaimableBalance" {
		result = *u.ClaimableBalance
		ok = true
	}

	return
}

// MustLiquidityPool retrieves the LiquidityPool value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustLiquidityPool() LedgerKeyLiquidityPool {
	val, ok := u.GetLiquidityPool()

	if !ok {
		panic("arm LiquidityPool is not set")
	}

	return val
}

// GetLiquidityPool retrieves the LiquidityPool value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetLiquidityPool() (result LedgerKeyLiquidityPool, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "LiquidityPool" {
		result = *u.LiquidityPool
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerKey) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerKey) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerKey)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerKey)(nil)
)

// EnvelopeType is an XDR Enum defines as:
//
//   enum EnvelopeType
//    {
//        ENVELOPE_TYPE_TX_V0 = 0,
//        ENVELOPE_TYPE_SCP = 1,
//        ENVELOPE_TYPE_TX = 2,
//        ENVELOPE_TYPE_AUTH = 3,
//        ENVELOPE_TYPE_SCPVALUE = 4,
//        ENVELOPE_TYPE_TX_FEE_BUMP = 5,
//        ENVELOPE_TYPE_OP_ID = 6
//    };
//
type EnvelopeType int32

const (
	EnvelopeTypeEnvelopeTypeTxV0      EnvelopeType = 0
	EnvelopeTypeEnvelopeTypeScp       EnvelopeType = 1
	EnvelopeTypeEnvelopeTypeTx        EnvelopeType = 2
	EnvelopeTypeEnvelopeTypeAuth      EnvelopeType = 3
	EnvelopeTypeEnvelopeTypeScpvalue  EnvelopeType = 4
	EnvelopeTypeEnvelopeTypeTxFeeBump EnvelopeType = 5
	EnvelopeTypeEnvelopeTypeOpId      EnvelopeType = 6
)

var envelopeTypeMap = map[int32]string{
	0: "EnvelopeTypeEnvelopeTypeTxV0",
	1: "EnvelopeTypeEnvelopeTypeScp",
	2: "EnvelopeTypeEnvelopeTypeTx",
	3: "EnvelopeTypeEnvelopeTypeAuth",
	4: "EnvelopeTypeEnvelopeTypeScpvalue",
	5: "EnvelopeTypeEnvelopeTypeTxFeeBump",
	6: "EnvelopeTypeEnvelopeTypeOpId",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for EnvelopeType
func (e EnvelopeType) ValidEnum(v int32) bool {
	_, ok := envelopeTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e EnvelopeType) String() string {
	name, _ := envelopeTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s EnvelopeType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *EnvelopeType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*EnvelopeType)(nil)
	_ encoding.BinaryUnmarshaler = (*EnvelopeType)(nil)
)

// UpgradeType is an XDR Typedef defines as:
//
//   typedef opaque UpgradeType<128>;
//
type UpgradeType []byte

// XDRMaxSize implements the Sized interface for UpgradeType
func (e UpgradeType) XDRMaxSize() int {
	return 128
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s UpgradeType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *UpgradeType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*UpgradeType)(nil)
	_ encoding.BinaryUnmarshaler = (*UpgradeType)(nil)
)

// StellarValueType is an XDR Enum defines as:
//
//   enum StellarValueType
//    {
//        STELLAR_VALUE_BASIC = 0,
//        STELLAR_VALUE_SIGNED = 1
//    };
//
type StellarValueType int32

const (
	StellarValueTypeStellarValueBasic  StellarValueType = 0
	StellarValueTypeStellarValueSigned StellarValueType = 1
)

var stellarValueTypeMap = map[int32]string{
	0: "StellarValueTypeStellarValueBasic",
	1: "StellarValueTypeStellarValueSigned",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for StellarValueType
func (e StellarValueType) ValidEnum(v int32) bool {
	_, ok := stellarValueTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e StellarValueType) String() string {
	name, _ := stellarValueTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s StellarValueType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *StellarValueType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*StellarValueType)(nil)
	_ encoding.BinaryUnmarshaler = (*StellarValueType)(nil)
)

// LedgerCloseValueSignature is an XDR Struct defines as:
//
//   struct LedgerCloseValueSignature
//    {
//        NodeID nodeID;       // which node introduced the value
//        Signature signature; // nodeID's signature
//    };
//
type LedgerCloseValueSignature struct {
	NodeId    NodeId
	Signature Signature
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerCloseValueSignature) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerCloseValueSignature) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerCloseValueSignature)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerCloseValueSignature)(nil)
)

// StellarValueExt is an XDR NestedUnion defines as:
//
//   union switch (StellarValueType v)
//        {
//        case STELLAR_VALUE_BASIC:
//            void;
//        case STELLAR_VALUE_SIGNED:
//            LedgerCloseValueSignature lcValueSignature;
//        }
//
type StellarValueExt struct {
	V                StellarValueType
	LcValueSignature *LedgerCloseValueSignature
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u StellarValueExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of StellarValueExt
func (u StellarValueExt) ArmForSwitch(sw int32) (string, bool) {
	switch StellarValueType(sw) {
	case StellarValueTypeStellarValueBasic:
		return "", true
	case StellarValueTypeStellarValueSigned:
		return "LcValueSignature", true
	}
	return "-", false
}

// NewStellarValueExt creates a new  StellarValueExt.
func NewStellarValueExt(v StellarValueType, value interface{}) (result StellarValueExt, err error) {
	result.V = v
	switch StellarValueType(v) {
	case StellarValueTypeStellarValueBasic:
		// void
	case StellarValueTypeStellarValueSigned:
		tv, ok := value.(LedgerCloseValueSignature)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerCloseValueSignature")
			return
		}
		result.LcValueSignature = &tv
	}
	return
}

// MustLcValueSignature retrieves the LcValueSignature value from the union,
// panicing if the value is not set.
func (u StellarValueExt) MustLcValueSignature() LedgerCloseValueSignature {
	val, ok := u.GetLcValueSignature()

	if !ok {
		panic("arm LcValueSignature is not set")
	}

	return val
}

// GetLcValueSignature retrieves the LcValueSignature value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarValueExt) GetLcValueSignature() (result LedgerCloseValueSignature, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "LcValueSignature" {
		result = *u.LcValueSignature
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s StellarValueExt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *StellarValueExt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*StellarValueExt)(nil)
	_ encoding.BinaryUnmarshaler = (*StellarValueExt)(nil)
)

// StellarValue is an XDR Struct defines as:
//
//   struct StellarValue
//    {
//        Hash txSetHash;      // transaction set to apply to previous ledger
//        TimePoint closeTime; // network close time
//
//        // upgrades to apply to the previous ledger (usually empty)
//        // this is a vector of encoded 'LedgerUpgrade' so that nodes can drop
//        // unknown steps during consensus if needed.
//        // see notes below on 'LedgerUpgrade' for more detail
//        // max size is dictated by number of upgrade types (+ room for future)
//        UpgradeType upgrades<6>;
//
//        // reserved for future use
//        union switch (StellarValueType v)
//        {
//        case STELLAR_VALUE_BASIC:
//            void;
//        case STELLAR_VALUE_SIGNED:
//            LedgerCloseValueSignature lcValueSignature;
//        }
//        ext;
//    };
//
type StellarValue struct {
	TxSetHash Hash
	CloseTime TimePoint
	Upgrades  []UpgradeType `xdrmaxsize:"6"`
	Ext       StellarValueExt
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s StellarValue) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *StellarValue) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*StellarValue)(nil)
	_ encoding.BinaryUnmarshaler = (*StellarValue)(nil)
)

// LedgerHeaderExt is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        }
//
type LedgerHeaderExt struct {
	V int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerHeaderExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerHeaderExt
func (u LedgerHeaderExt) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	}
	return "-", false
}

// NewLedgerHeaderExt creates a new  LedgerHeaderExt.
func NewLedgerHeaderExt(v int32, value interface{}) (result LedgerHeaderExt, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerHeaderExt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerHeaderExt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerHeaderExt)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerHeaderExt)(nil)
)

// LedgerHeader is an XDR Struct defines as:
//
//   struct LedgerHeader
//    {
//        uint32 ledgerVersion;    // the protocol version of the ledger
//        Hash previousLedgerHash; // hash of the previous ledger header
//        StellarValue scpValue;   // what consensus agreed to
//        Hash txSetResultHash;    // the TransactionResultSet that led to this ledger
//        Hash bucketListHash;     // hash of the ledger state
//
//        uint32 ledgerSeq; // sequence number of this ledger
//
//        int64 totalCoins; // total number of stroops in existence.
//                          // 10,000,000 stroops in 1 XLM
//
//        int64 feePool;       // fees burned since last inflation run
//        uint32 inflationSeq; // inflation sequence number
//
//        uint64 idPool; // last used global ID, used for generating objects
//
//        uint32 baseFee;     // base fee per operation in stroops
//        uint32 baseReserve; // account base reserve in stroops
//
//        uint32 maxTxSetSize; // maximum size a transaction set can be
//
//        Hash skipList[4]; // hashes of ledgers in the past. allows you to jump back
//                          // in time without walking the chain back ledger by ledger
//                          // each slot contains the oldest ledger that is mod of
//                          // either 50  5000  50000 or 500000 depending on index
//                          // skipList[0] mod(50), skipList[1] mod(5000), etc
//
//        // reserved for future use
//        union switch (int v)
//        {
//        case 0:
//            void;
//        }
//        ext;
//    };
//
type LedgerHeader struct {
	LedgerVersion      Uint32
	PreviousLedgerHash Hash
	ScpValue           StellarValue
	TxSetResultHash    Hash
	BucketListHash     Hash
	LedgerSeq          Uint32
	TotalCoins         Int64
	FeePool            Int64
	InflationSeq       Uint32
	IdPool             Uint64
	BaseFee            Uint32
	BaseReserve        Uint32
	MaxTxSetSize       Uint32
	SkipList           [4]Hash
	Ext                LedgerHeaderExt
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerHeader) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerHeader) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerHeader)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerHeader)(nil)
)

// LedgerUpgradeType is an XDR Enum defines as:
//
//   enum LedgerUpgradeType
//    {
//        LEDGER_UPGRADE_VERSION = 1,
//        LEDGER_UPGRADE_BASE_FEE = 2,
//        LEDGER_UPGRADE_MAX_TX_SET_SIZE = 3,
//        LEDGER_UPGRADE_BASE_RESERVE = 4
//    };
//
type LedgerUpgradeType int32

const (
	LedgerUpgradeTypeLedgerUpgradeVersion      LedgerUpgradeType = 1
	LedgerUpgradeTypeLedgerUpgradeBaseFee      LedgerUpgradeType = 2
	LedgerUpgradeTypeLedgerUpgradeMaxTxSetSize LedgerUpgradeType = 3
	LedgerUpgradeTypeLedgerUpgradeBaseReserve  LedgerUpgradeType = 4
)

var ledgerUpgradeTypeMap = map[int32]string{
	1: "LedgerUpgradeTypeLedgerUpgradeVersion",
	2: "LedgerUpgradeTypeLedgerUpgradeBaseFee",
	3: "LedgerUpgradeTypeLedgerUpgradeMaxTxSetSize",
	4: "LedgerUpgradeTypeLedgerUpgradeBaseReserve",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LedgerUpgradeType
func (e LedgerUpgradeType) ValidEnum(v int32) bool {
	_, ok := ledgerUpgradeTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e LedgerUpgradeType) String() string {
	name, _ := ledgerUpgradeTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerUpgradeType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerUpgradeType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerUpgradeType)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerUpgradeType)(nil)
)

// LedgerUpgrade is an XDR Union defines as:
//
//   union LedgerUpgrade switch (LedgerUpgradeType type)
//    {
//    case LEDGER_UPGRADE_VERSION:
//        uint32 newLedgerVersion; // update ledgerVersion
//    case LEDGER_UPGRADE_BASE_FEE:
//        uint32 newBaseFee; // update baseFee
//    case LEDGER_UPGRADE_MAX_TX_SET_SIZE:
//        uint32 newMaxTxSetSize; // update maxTxSetSize
//    case LEDGER_UPGRADE_BASE_RESERVE:
//        uint32 newBaseReserve; // update baseReserve
//    };
//
type LedgerUpgrade struct {
	Type             LedgerUpgradeType
	NewLedgerVersion *Uint32
	NewBaseFee       *Uint32
	NewMaxTxSetSize  *Uint32
	NewBaseReserve   *Uint32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerUpgrade) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerUpgrade
func (u LedgerUpgrade) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerUpgradeType(sw) {
	case LedgerUpgradeTypeLedgerUpgradeVersion:
		return "NewLedgerVersion", true
	case LedgerUpgradeTypeLedgerUpgradeBaseFee:
		return "NewBaseFee", true
	case LedgerUpgradeTypeLedgerUpgradeMaxTxSetSize:
		return "NewMaxTxSetSize", true
	case LedgerUpgradeTypeLedgerUpgradeBaseReserve:
		return "NewBaseReserve", true
	}
	return "-", false
}

// NewLedgerUpgrade creates a new  LedgerUpgrade.
func NewLedgerUpgrade(aType LedgerUpgradeType, value interface{}) (result LedgerUpgrade, err error) {
	result.Type = aType
	switch LedgerUpgradeType(aType) {
	case LedgerUpgradeTypeLedgerUpgradeVersion:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.NewLedgerVersion = &tv
	case LedgerUpgradeTypeLedgerUpgradeBaseFee:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.NewBaseFee = &tv
	case LedgerUpgradeTypeLedgerUpgradeMaxTxSetSize:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.NewMaxTxSetSize = &tv
	case LedgerUpgradeTypeLedgerUpgradeBaseReserve:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.NewBaseReserve = &tv
	}
	return
}

// MustNewLedgerVersion retrieves the NewLedgerVersion value from the union,
// panicing if the value is not set.
func (u LedgerUpgrade) MustNewLedgerVersion() Uint32 {
	val, ok := u.GetNewLedgerVersion()

	if !ok {
		panic("arm NewLedgerVersion is not set")
	}

	return val
}

// GetNewLedgerVersion retrieves the NewLedgerVersion value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerUpgrade) GetNewLedgerVersion() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "NewLedgerVersion" {
		result = *u.NewLedgerVersion
		ok = true
	}

	return
}

// MustNewBaseFee retrieves the NewBaseFee value from the union,
// panicing if the value is not set.
func (u LedgerUpgrade) MustNewBaseFee() Uint32 {
	val, ok := u.GetNewBaseFee()

	if !ok {
		panic("arm NewBaseFee is not set")
	}

	return val
}

// GetNewBaseFee retrieves the NewBaseFee value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerUpgrade) GetNewBaseFee() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "NewBaseFee" {
		result = *u.NewBaseFee
		ok = true
	}

	return
}

// MustNewMaxTxSetSize retrieves the NewMaxTxSetSize value from the union,
// panicing if the value is not set.
func (u LedgerUpgrade) MustNewMaxTxSetSize() Uint32 {
	val, ok := u.GetNewMaxTxSetSize()

	if !ok {
		panic("arm NewMaxTxSetSize is not set")
	}

	return val
}

// GetNewMaxTxSetSize retrieves the NewMaxTxSetSize value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerUpgrade) GetNewMaxTxSetSize() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "NewMaxTxSetSize" {
		result = *u.NewMaxTxSetSize
		ok = true
	}

	return
}

// MustNewBaseReserve retrieves the NewBaseReserve value from the union,
// panicing if the value is not set.
func (u LedgerUpgrade) MustNewBaseReserve() Uint32 {
	val, ok := u.GetNewBaseReserve()

	if !ok {
		panic("arm NewBaseReserve is not set")
	}

	return val
}

// GetNewBaseReserve retrieves the NewBaseReserve value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerUpgrade) GetNewBaseReserve() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "NewBaseReserve" {
		result = *u.NewBaseReserve
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerUpgrade) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerUpgrade) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerUpgrade)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerUpgrade)(nil)
)

// BucketEntryType is an XDR Enum defines as:
//
//   enum BucketEntryType
//    {
//        METAENTRY =
//            -1, // At-and-after protocol 11: bucket metadata, should come first.
//        LIVEENTRY = 0, // Before protocol 11: created-or-updated;
//                       // At-and-after protocol 11: only updated.
//        DEADENTRY = 1,
//        INITENTRY = 2 // At-and-after protocol 11: only created.
//    };
//
type BucketEntryType int32

const (
	BucketEntryTypeMetaentry BucketEntryType = -1
	BucketEntryTypeLiveentry BucketEntryType = 0
	BucketEntryTypeDeadentry BucketEntryType = 1
	BucketEntryTypeInitentry BucketEntryType = 2
)

var bucketEntryTypeMap = map[int32]string{
	-1: "BucketEntryTypeMetaentry",
	0:  "BucketEntryTypeLiveentry",
	1:  "BucketEntryTypeDeadentry",
	2:  "BucketEntryTypeInitentry",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for BucketEntryType
func (e BucketEntryType) ValidEnum(v int32) bool {
	_, ok := bucketEntryTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e BucketEntryType) String() string {
	name, _ := bucketEntryTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BucketEntryType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BucketEntryType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BucketEntryType)(nil)
	_ encoding.BinaryUnmarshaler = (*BucketEntryType)(nil)
)

// BucketMetadataExt is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        }
//
type BucketMetadataExt struct {
	V int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BucketMetadataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of BucketMetadataExt
func (u BucketMetadataExt) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	}
	return "-", false
}

// NewBucketMetadataExt creates a new  BucketMetadataExt.
func NewBucketMetadataExt(v int32, value interface{}) (result BucketMetadataExt, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BucketMetadataExt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BucketMetadataExt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BucketMetadataExt)(nil)
	_ encoding.BinaryUnmarshaler = (*BucketMetadataExt)(nil)
)

// BucketMetadata is an XDR Struct defines as:
//
//   struct BucketMetadata
//    {
//        // Indicates the protocol version used to create / merge this bucket.
//        uint32 ledgerVersion;
//
//        // reserved for future use
//        union switch (int v)
//        {
//        case 0:
//            void;
//        }
//        ext;
//    };
//
type BucketMetadata struct {
	LedgerVersion Uint32
	Ext           BucketMetadataExt
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BucketMetadata) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BucketMetadata) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BucketMetadata)(nil)
	_ encoding.BinaryUnmarshaler = (*BucketMetadata)(nil)
)

// BucketEntry is an XDR Union defines as:
//
//   union BucketEntry switch (BucketEntryType type)
//    {
//    case LIVEENTRY:
//    case INITENTRY:
//        LedgerEntry liveEntry;
//
//    case DEADENTRY:
//        LedgerKey deadEntry;
//    case METAENTRY:
//        BucketMetadata metaEntry;
//    };
//
type BucketEntry struct {
	Type      BucketEntryType
	LiveEntry *LedgerEntry
	DeadEntry *LedgerKey
	MetaEntry *BucketMetadata
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BucketEntry) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of BucketEntry
func (u BucketEntry) ArmForSwitch(sw int32) (string, bool) {
	switch BucketEntryType(sw) {
	case BucketEntryTypeLiveentry:
		return "LiveEntry", true
	case BucketEntryTypeInitentry:
		return "LiveEntry", true
	case BucketEntryTypeDeadentry:
		return "DeadEntry", true
	case BucketEntryTypeMetaentry:
		return "MetaEntry", true
	}
	return "-", false
}

// NewBucketEntry creates a new  BucketEntry.
func NewBucketEntry(aType BucketEntryType, value interface{}) (result BucketEntry, err error) {
	result.Type = aType
	switch BucketEntryType(aType) {
	case BucketEntryTypeLiveentry:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.LiveEntry = &tv
	case BucketEntryTypeInitentry:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.LiveEntry = &tv
	case BucketEntryTypeDeadentry:
		tv, ok := value.(LedgerKey)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKey")
			return
		}
		result.DeadEntry = &tv
	case BucketEntryTypeMetaentry:
		tv, ok := value.(BucketMetadata)
		if !ok {
			err = fmt.Errorf("invalid value, must be BucketMetadata")
			return
		}
		result.MetaEntry = &tv
	}
	return
}

// MustLiveEntry retrieves the LiveEntry value from the union,
// panicing if the value is not set.
func (u BucketEntry) MustLiveEntry() LedgerEntry {
	val, ok := u.GetLiveEntry()

	if !ok {
		panic("arm LiveEntry is not set")
	}

	return val
}

// GetLiveEntry retrieves the LiveEntry value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u BucketEntry) GetLiveEntry() (result LedgerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "LiveEntry" {
		result = *u.LiveEntry
		ok = true
	}

	return
}

// MustDeadEntry retrieves the DeadEntry value from the union,
// panicing if the value is not set.
func (u BucketEntry) MustDeadEntry() LedgerKey {
	val, ok := u.GetDeadEntry()

	if !ok {
		panic("arm DeadEntry is not set")
	}

	return val
}

// GetDeadEntry retrieves the DeadEntry value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u BucketEntry) GetDeadEntry() (result LedgerKey, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "DeadEntry" {
		result = *u.DeadEntry
		ok = true
	}

	return
}

// MustMetaEntry retrieves the MetaEntry value from the union,
// panicing if the value is not set.
func (u BucketEntry) MustMetaEntry() BucketMetadata {
	val, ok := u.GetMetaEntry()

	if !ok {
		panic("arm MetaEntry is not set")
	}

	return val
}

// GetMetaEntry retrieves the MetaEntry value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u BucketEntry) GetMetaEntry() (result BucketMetadata, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "MetaEntry" {
		result = *u.MetaEntry
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BucketEntry) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BucketEntry) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BucketEntry)(nil)
	_ encoding.BinaryUnmarshaler = (*BucketEntry)(nil)
)

// TransactionSet is an XDR Struct defines as:
//
//   struct TransactionSet
//    {
//        Hash previousLedgerHash;
//        TransactionEnvelope txs<>;
//    };
//
type TransactionSet struct {
	PreviousLedgerHash Hash
	Txs                []TransactionEnvelope
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionSet) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionSet) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionSet)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionSet)(nil)
)

// TransactionResultPair is an XDR Struct defines as:
//
//   struct TransactionResultPair
//    {
//        Hash transactionHash;
//        TransactionResult result; // result for the transaction
//    };
//
type TransactionResultPair struct {
	TransactionHash Hash
	Result          TransactionResult
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionResultPair) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionResultPair) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionResultPair)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionResultPair)(nil)
)

// TransactionResultSet is an XDR Struct defines as:
//
//   struct TransactionResultSet
//    {
//        TransactionResultPair results<>;
//    };
//
type TransactionResultSet struct {
	Results []TransactionResultPair
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionResultSet) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionResultSet) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionResultSet)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionResultSet)(nil)
)

// TransactionHistoryEntryExt is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        }
//
type TransactionHistoryEntryExt struct {
	V int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionHistoryEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionHistoryEntryExt
func (u TransactionHistoryEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	}
	return "-", false
}

// NewTransactionHistoryEntryExt creates a new  TransactionHistoryEntryExt.
func NewTransactionHistoryEntryExt(v int32, value interface{}) (result TransactionHistoryEntryExt, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionHistoryEntryExt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionHistoryEntryExt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionHistoryEntryExt)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionHistoryEntryExt)(nil)
)

// TransactionHistoryEntry is an XDR Struct defines as:
//
//   struct TransactionHistoryEntry
//    {
//        uint32 ledgerSeq;
//        TransactionSet txSet;
//
//        // reserved for future use
//        union switch (int v)
//        {
//        case 0:
//            void;
//        }
//        ext;
//    };
//
type TransactionHistoryEntry struct {
	LedgerSeq Uint32
	TxSet     TransactionSet
	Ext       TransactionHistoryEntryExt
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionHistoryEntry) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionHistoryEntry) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionHistoryEntry)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionHistoryEntry)(nil)
)

// TransactionHistoryResultEntryExt is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        }
//
type TransactionHistoryResultEntryExt struct {
	V int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionHistoryResultEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionHistoryResultEntryExt
func (u TransactionHistoryResultEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	}
	return "-", false
}

// NewTransactionHistoryResultEntryExt creates a new  TransactionHistoryResultEntryExt.
func NewTransactionHistoryResultEntryExt(v int32, value interface{}) (result TransactionHistoryResultEntryExt, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionHistoryResultEntryExt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionHistoryResultEntryExt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionHistoryResultEntryExt)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionHistoryResultEntryExt)(nil)
)

// TransactionHistoryResultEntry is an XDR Struct defines as:
//
//   struct TransactionHistoryResultEntry
//    {
//        uint32 ledgerSeq;
//        TransactionResultSet txResultSet;
//
//        // reserved for future use
//        union switch (int v)
//        {
//        case 0:
//            void;
//        }
//        ext;
//    };
//
type TransactionHistoryResultEntry struct {
	LedgerSeq   Uint32
	TxResultSet TransactionResultSet
	Ext         TransactionHistoryResultEntryExt
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionHistoryResultEntry) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionHistoryResultEntry) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionHistoryResultEntry)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionHistoryResultEntry)(nil)
)

// LedgerHeaderHistoryEntryExt is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        }
//
type LedgerHeaderHistoryEntryExt struct {
	V int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerHeaderHistoryEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerHeaderHistoryEntryExt
func (u LedgerHeaderHistoryEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	}
	return "-", false
}

// NewLedgerHeaderHistoryEntryExt creates a new  LedgerHeaderHistoryEntryExt.
func NewLedgerHeaderHistoryEntryExt(v int32, value interface{}) (result LedgerHeaderHistoryEntryExt, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerHeaderHistoryEntryExt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerHeaderHistoryEntryExt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerHeaderHistoryEntryExt)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerHeaderHistoryEntryExt)(nil)
)

// LedgerHeaderHistoryEntry is an XDR Struct defines as:
//
//   struct LedgerHeaderHistoryEntry
//    {
//        Hash hash;
//        LedgerHeader header;
//
//        // reserved for future use
//        union switch (int v)
//        {
//        case 0:
//            void;
//        }
//        ext;
//    };
//
type LedgerHeaderHistoryEntry struct {
	Hash   Hash
	Header LedgerHeader
	Ext    LedgerHeaderHistoryEntryExt
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerHeaderHistoryEntry) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerHeaderHistoryEntry) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerHeaderHistoryEntry)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerHeaderHistoryEntry)(nil)
)

// LedgerScpMessages is an XDR Struct defines as:
//
//   struct LedgerSCPMessages
//    {
//        uint32 ledgerSeq;
//        SCPEnvelope messages<>;
//    };
//
type LedgerScpMessages struct {
	LedgerSeq Uint32
	Messages  []ScpEnvelope
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerScpMessages) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerScpMessages) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerScpMessages)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerScpMessages)(nil)
)

// ScpHistoryEntryV0 is an XDR Struct defines as:
//
//   struct SCPHistoryEntryV0
//    {
//        SCPQuorumSet quorumSets<>; // additional quorum sets used by ledgerMessages
//        LedgerSCPMessages ledgerMessages;
//    };
//
type ScpHistoryEntryV0 struct {
	QuorumSets     []ScpQuorumSet
	LedgerMessages LedgerScpMessages
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ScpHistoryEntryV0) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ScpHistoryEntryV0) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ScpHistoryEntryV0)(nil)
	_ encoding.BinaryUnmarshaler = (*ScpHistoryEntryV0)(nil)
)

// ScpHistoryEntry is an XDR Union defines as:
//
//   union SCPHistoryEntry switch (int v)
//    {
//    case 0:
//        SCPHistoryEntryV0 v0;
//    };
//
type ScpHistoryEntry struct {
	V  int32
	V0 *ScpHistoryEntryV0
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ScpHistoryEntry) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ScpHistoryEntry
func (u ScpHistoryEntry) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "V0", true
	}
	return "-", false
}

// NewScpHistoryEntry creates a new  ScpHistoryEntry.
func NewScpHistoryEntry(v int32, value interface{}) (result ScpHistoryEntry, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		tv, ok := value.(ScpHistoryEntryV0)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpHistoryEntryV0")
			return
		}
		result.V0 = &tv
	}
	return
}

// MustV0 retrieves the V0 value from the union,
// panicing if the value is not set.
func (u ScpHistoryEntry) MustV0() ScpHistoryEntryV0 {
	val, ok := u.GetV0()

	if !ok {
		panic("arm V0 is not set")
	}

	return val
}

// GetV0 retrieves the V0 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpHistoryEntry) GetV0() (result ScpHistoryEntryV0, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "V0" {
		result = *u.V0
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ScpHistoryEntry) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ScpHistoryEntry) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ScpHistoryEntry)(nil)
	_ encoding.BinaryUnmarshaler = (*ScpHistoryEntry)(nil)
)

// LedgerEntryChangeType is an XDR Enum defines as:
//
//   enum LedgerEntryChangeType
//    {
//        LEDGER_ENTRY_CREATED = 0, // entry was added to the ledger
//        LEDGER_ENTRY_UPDATED = 1, // entry was modified in the ledger
//        LEDGER_ENTRY_REMOVED = 2, // entry was removed from the ledger
//        LEDGER_ENTRY_STATE = 3    // value of the entry
//    };
//
type LedgerEntryChangeType int32

const (
	LedgerEntryChangeTypeLedgerEntryCreated LedgerEntryChangeType = 0
	LedgerEntryChangeTypeLedgerEntryUpdated LedgerEntryChangeType = 1
	LedgerEntryChangeTypeLedgerEntryRemoved LedgerEntryChangeType = 2
	LedgerEntryChangeTypeLedgerEntryState   LedgerEntryChangeType = 3
)

var ledgerEntryChangeTypeMap = map[int32]string{
	0: "LedgerEntryChangeTypeLedgerEntryCreated",
	1: "LedgerEntryChangeTypeLedgerEntryUpdated",
	2: "LedgerEntryChangeTypeLedgerEntryRemoved",
	3: "LedgerEntryChangeTypeLedgerEntryState",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LedgerEntryChangeType
func (e LedgerEntryChangeType) ValidEnum(v int32) bool {
	_, ok := ledgerEntryChangeTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e LedgerEntryChangeType) String() string {
	name, _ := ledgerEntryChangeTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerEntryChangeType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerEntryChangeType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerEntryChangeType)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerEntryChangeType)(nil)
)

// LedgerEntryChange is an XDR Union defines as:
//
//   union LedgerEntryChange switch (LedgerEntryChangeType type)
//    {
//    case LEDGER_ENTRY_CREATED:
//        LedgerEntry created;
//    case LEDGER_ENTRY_UPDATED:
//        LedgerEntry updated;
//    case LEDGER_ENTRY_REMOVED:
//        LedgerKey removed;
//    case LEDGER_ENTRY_STATE:
//        LedgerEntry state;
//    };
//
type LedgerEntryChange struct {
	Type    LedgerEntryChangeType
	Created *LedgerEntry
	Updated *LedgerEntry
	Removed *LedgerKey
	State   *LedgerEntry
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerEntryChange) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerEntryChange
func (u LedgerEntryChange) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerEntryChangeType(sw) {
	case LedgerEntryChangeTypeLedgerEntryCreated:
		return "Created", true
	case LedgerEntryChangeTypeLedgerEntryUpdated:
		return "Updated", true
	case LedgerEntryChangeTypeLedgerEntryRemoved:
		return "Removed", true
	case LedgerEntryChangeTypeLedgerEntryState:
		return "State", true
	}
	return "-", false
}

// NewLedgerEntryChange creates a new  LedgerEntryChange.
func NewLedgerEntryChange(aType LedgerEntryChangeType, value interface{}) (result LedgerEntryChange, err error) {
	result.Type = aType
	switch LedgerEntryChangeType(aType) {
	case LedgerEntryChangeTypeLedgerEntryCreated:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.Created = &tv
	case LedgerEntryChangeTypeLedgerEntryUpdated:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.Updated = &tv
	case LedgerEntryChangeTypeLedgerEntryRemoved:
		tv, ok := value.(LedgerKey)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKey")
			return
		}
		result.Removed = &tv
	case LedgerEntryChangeTypeLedgerEntryState:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.State = &tv
	}
	return
}

// MustCreated retrieves the Created value from the union,
// panicing if the value is not set.
func (u LedgerEntryChange) MustCreated() LedgerEntry {
	val, ok := u.GetCreated()

	if !ok {
		panic("arm Created is not set")
	}

	return val
}

// GetCreated retrieves the Created value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryChange) GetCreated() (result LedgerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Created" {
		result = *u.Created
		ok = true
	}

	return
}

// MustUpdated retrieves the Updated value from the union,
// panicing if the value is not set.
func (u LedgerEntryChange) MustUpdated() LedgerEntry {
	val, ok := u.GetUpdated()

	if !ok {
		panic("arm Updated is not set")
	}

	return val
}

// GetUpdated retrieves the Updated value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryChange) GetUpdated() (result LedgerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Updated" {
		result = *u.Updated
		ok = true
	}

	return
}

// MustRemoved retrieves the Removed value from the union,
// panicing if the value is not set.
func (u LedgerEntryChange) MustRemoved() LedgerKey {
	val, ok := u.GetRemoved()

	if !ok {
		panic("arm Removed is not set")
	}

	return val
}

// GetRemoved retrieves the Removed value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryChange) GetRemoved() (result LedgerKey, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Removed" {
		result = *u.Removed
		ok = true
	}

	return
}

// MustState retrieves the State value from the union,
// panicing if the value is not set.
func (u LedgerEntryChange) MustState() LedgerEntry {
	val, ok := u.GetState()

	if !ok {
		panic("arm State is not set")
	}

	return val
}

// GetState retrieves the State value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryChange) GetState() (result LedgerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "State" {
		result = *u.State
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerEntryChange) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerEntryChange) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerEntryChange)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerEntryChange)(nil)
)

// LedgerEntryChanges is an XDR Typedef defines as:
//
//   typedef LedgerEntryChange LedgerEntryChanges<>;
//
type LedgerEntryChanges []LedgerEntryChange

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerEntryChanges) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerEntryChanges) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerEntryChanges)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerEntryChanges)(nil)
)

// OperationMeta is an XDR Struct defines as:
//
//   struct OperationMeta
//    {
//        LedgerEntryChanges changes;
//    };
//
type OperationMeta struct {
	Changes LedgerEntryChanges
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s OperationMeta) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *OperationMeta) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*OperationMeta)(nil)
	_ encoding.BinaryUnmarshaler = (*OperationMeta)(nil)
)

// TransactionMetaV1 is an XDR Struct defines as:
//
//   struct TransactionMetaV1
//    {
//        LedgerEntryChanges txChanges; // tx level changes if any
//        OperationMeta operations<>;   // meta for each operation
//    };
//
type TransactionMetaV1 struct {
	TxChanges  LedgerEntryChanges
	Operations []OperationMeta
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionMetaV1) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionMetaV1) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionMetaV1)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionMetaV1)(nil)
)

// TransactionMetaV2 is an XDR Struct defines as:
//
//   struct TransactionMetaV2
//    {
//        LedgerEntryChanges txChangesBefore; // tx level changes before operations
//                                            // are applied if any
//        OperationMeta operations<>;         // meta for each operation
//        LedgerEntryChanges txChangesAfter;  // tx level changes after operations are
//                                            // applied if any
//    };
//
type TransactionMetaV2 struct {
	TxChangesBefore LedgerEntryChanges
	Operations      []OperationMeta
	TxChangesAfter  LedgerEntryChanges
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionMetaV2) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionMetaV2) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionMetaV2)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionMetaV2)(nil)
)

// TransactionMeta is an XDR Union defines as:
//
//   union TransactionMeta switch (int v)
//    {
//    case 0:
//        OperationMeta operations<>;
//    case 1:
//        TransactionMetaV1 v1;
//    case 2:
//        TransactionMetaV2 v2;
//    };
//
type TransactionMeta struct {
	V          int32
	Operations *[]OperationMeta
	V1         *TransactionMetaV1
	V2         *TransactionMetaV2
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionMeta) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionMeta
func (u TransactionMeta) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "Operations", true
	case 1:
		return "V1", true
	case 2:
		return "V2", true
	}
	return "-", false
}

// NewTransactionMeta creates a new  TransactionMeta.
func NewTransactionMeta(v int32, value interface{}) (result TransactionMeta, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		tv, ok := value.([]OperationMeta)
		if !ok {
			err = fmt.Errorf("invalid value, must be []OperationMeta")
			return
		}
		result.Operations = &tv
	case 1:
		tv, ok := value.(TransactionMetaV1)
		if !ok {
			err = fmt.Errorf("invalid value, must be TransactionMetaV1")
			return
		}
		result.V1 = &tv
	case 2:
		tv, ok := value.(TransactionMetaV2)
		if !ok {
			err = fmt.Errorf("invalid value, must be TransactionMetaV2")
			return
		}
		result.V2 = &tv
	}
	return
}

// MustOperations retrieves the Operations value from the union,
// panicing if the value is not set.
func (u TransactionMeta) MustOperations() []OperationMeta {
	val, ok := u.GetOperations()

	if !ok {
		panic("arm Operations is not set")
	}

	return val
}

// GetOperations retrieves the Operations value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionMeta) GetOperations() (result []OperationMeta, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "Operations" {
		result = *u.Operations
		ok = true
	}

	return
}

// MustV1 retrieves the V1 value from the union,
// panicing if the value is not set.
func (u TransactionMeta) MustV1() TransactionMetaV1 {
	val, ok := u.GetV1()

	if !ok {
		panic("arm V1 is not set")
	}

	return val
}

// GetV1 retrieves the V1 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionMeta) GetV1() (result TransactionMetaV1, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "V1" {
		result = *u.V1
		ok = true
	}

	return
}

// MustV2 retrieves the V2 value from the union,
// panicing if the value is not set.
func (u TransactionMeta) MustV2() TransactionMetaV2 {
	val, ok := u.GetV2()

	if !ok {
		panic("arm V2 is not set")
	}

	return val
}

// GetV2 retrieves the V2 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionMeta) GetV2() (result TransactionMetaV2, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "V2" {
		result = *u.V2
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionMeta) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionMeta) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionMeta)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionMeta)(nil)
)

// TransactionResultMeta is an XDR Struct defines as:
//
//   struct TransactionResultMeta
//    {
//        TransactionResultPair result;
//        LedgerEntryChanges feeProcessing;
//        TransactionMeta txApplyProcessing;
//    };
//
type TransactionResultMeta struct {
	Result            TransactionResultPair
	FeeProcessing     LedgerEntryChanges
	TxApplyProcessing TransactionMeta
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionResultMeta) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionResultMeta) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionResultMeta)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionResultMeta)(nil)
)

// UpgradeEntryMeta is an XDR Struct defines as:
//
//   struct UpgradeEntryMeta
//    {
//        LedgerUpgrade upgrade;
//        LedgerEntryChanges changes;
//    };
//
type UpgradeEntryMeta struct {
	Upgrade LedgerUpgrade
	Changes LedgerEntryChanges
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s UpgradeEntryMeta) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *UpgradeEntryMeta) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*UpgradeEntryMeta)(nil)
	_ encoding.BinaryUnmarshaler = (*UpgradeEntryMeta)(nil)
)

// LedgerCloseMetaV0 is an XDR Struct defines as:
//
//   struct LedgerCloseMetaV0
//    {
//        LedgerHeaderHistoryEntry ledgerHeader;
//        // NB: txSet is sorted in "Hash order"
//        TransactionSet txSet;
//
//        // NB: transactions are sorted in apply order here
//        // fees for all transactions are processed first
//        // followed by applying transactions
//        TransactionResultMeta txProcessing<>;
//
//        // upgrades are applied last
//        UpgradeEntryMeta upgradesProcessing<>;
//
//        // other misc information attached to the ledger close
//        SCPHistoryEntry scpInfo<>;
//    };
//
type LedgerCloseMetaV0 struct {
	LedgerHeader       LedgerHeaderHistoryEntry
	TxSet              TransactionSet
	TxProcessing       []TransactionResultMeta
	UpgradesProcessing []UpgradeEntryMeta
	ScpInfo            []ScpHistoryEntry
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerCloseMetaV0) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerCloseMetaV0) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerCloseMetaV0)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerCloseMetaV0)(nil)
)

// LedgerCloseMeta is an XDR Union defines as:
//
//   union LedgerCloseMeta switch (int v)
//    {
//    case 0:
//        LedgerCloseMetaV0 v0;
//    };
//
type LedgerCloseMeta struct {
	V  int32
	V0 *LedgerCloseMetaV0
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerCloseMeta) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerCloseMeta
func (u LedgerCloseMeta) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "V0", true
	}
	return "-", false
}

// NewLedgerCloseMeta creates a new  LedgerCloseMeta.
func NewLedgerCloseMeta(v int32, value interface{}) (result LedgerCloseMeta, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		tv, ok := value.(LedgerCloseMetaV0)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerCloseMetaV0")
			return
		}
		result.V0 = &tv
	}
	return
}

// MustV0 retrieves the V0 value from the union,
// panicing if the value is not set.
func (u LedgerCloseMeta) MustV0() LedgerCloseMetaV0 {
	val, ok := u.GetV0()

	if !ok {
		panic("arm V0 is not set")
	}

	return val
}

// GetV0 retrieves the V0 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerCloseMeta) GetV0() (result LedgerCloseMetaV0, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "V0" {
		result = *u.V0
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LedgerCloseMeta) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LedgerCloseMeta) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LedgerCloseMeta)(nil)
	_ encoding.BinaryUnmarshaler = (*LedgerCloseMeta)(nil)
)

// ErrorCode is an XDR Enum defines as:
//
//   enum ErrorCode
//    {
//        ERR_MISC = 0, // Unspecific error
//        ERR_DATA = 1, // Malformed data
//        ERR_CONF = 2, // Misconfiguration error
//        ERR_AUTH = 3, // Authentication failure
//        ERR_LOAD = 4  // System overloaded
//    };
//
type ErrorCode int32

const (
	ErrorCodeErrMisc ErrorCode = 0
	ErrorCodeErrData ErrorCode = 1
	ErrorCodeErrConf ErrorCode = 2
	ErrorCodeErrAuth ErrorCode = 3
	ErrorCodeErrLoad ErrorCode = 4
)

var errorCodeMap = map[int32]string{
	0: "ErrorCodeErrMisc",
	1: "ErrorCodeErrData",
	2: "ErrorCodeErrConf",
	3: "ErrorCodeErrAuth",
	4: "ErrorCodeErrLoad",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ErrorCode
func (e ErrorCode) ValidEnum(v int32) bool {
	_, ok := errorCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e ErrorCode) String() string {
	name, _ := errorCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ErrorCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ErrorCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ErrorCode)(nil)
	_ encoding.BinaryUnmarshaler = (*ErrorCode)(nil)
)

// Error is an XDR Struct defines as:
//
//   struct Error
//    {
//        ErrorCode code;
//        string msg<100>;
//    };
//
type Error struct {
	Code ErrorCode
	Msg  string `xdrmaxsize:"100"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Error) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Error) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Error)(nil)
	_ encoding.BinaryUnmarshaler = (*Error)(nil)
)

// AuthCert is an XDR Struct defines as:
//
//   struct AuthCert
//    {
//        Curve25519Public pubkey;
//        uint64 expiration;
//        Signature sig;
//    };
//
type AuthCert struct {
	Pubkey     Curve25519Public
	Expiration Uint64
	Sig        Signature
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AuthCert) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AuthCert) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AuthCert)(nil)
	_ encoding.BinaryUnmarshaler = (*AuthCert)(nil)
)

// Hello is an XDR Struct defines as:
//
//   struct Hello
//    {
//        uint32 ledgerVersion;
//        uint32 overlayVersion;
//        uint32 overlayMinVersion;
//        Hash networkID;
//        string versionStr<100>;
//        int listeningPort;
//        NodeID peerID;
//        AuthCert cert;
//        uint256 nonce;
//    };
//
type Hello struct {
	LedgerVersion     Uint32
	OverlayVersion    Uint32
	OverlayMinVersion Uint32
	NetworkId         Hash
	VersionStr        string `xdrmaxsize:"100"`
	ListeningPort     int32
	PeerId            NodeId
	Cert              AuthCert
	Nonce             Uint256
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Hello) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Hello) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Hello)(nil)
	_ encoding.BinaryUnmarshaler = (*Hello)(nil)
)

// Auth is an XDR Struct defines as:
//
//   struct Auth
//    {
//        // Empty message, just to confirm
//        // establishment of MAC keys.
//        int unused;
//    };
//
type Auth struct {
	Unused int32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Auth) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Auth) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Auth)(nil)
	_ encoding.BinaryUnmarshaler = (*Auth)(nil)
)

// IpAddrType is an XDR Enum defines as:
//
//   enum IPAddrType
//    {
//        IPv4 = 0,
//        IPv6 = 1
//    };
//
type IpAddrType int32

const (
	IpAddrTypeIPv4 IpAddrType = 0
	IpAddrTypeIPv6 IpAddrType = 1
)

var ipAddrTypeMap = map[int32]string{
	0: "IpAddrTypeIPv4",
	1: "IpAddrTypeIPv6",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for IpAddrType
func (e IpAddrType) ValidEnum(v int32) bool {
	_, ok := ipAddrTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e IpAddrType) String() string {
	name, _ := ipAddrTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s IpAddrType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *IpAddrType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*IpAddrType)(nil)
	_ encoding.BinaryUnmarshaler = (*IpAddrType)(nil)
)

// PeerAddressIp is an XDR NestedUnion defines as:
//
//   union switch (IPAddrType type)
//        {
//        case IPv4:
//            opaque ipv4[4];
//        case IPv6:
//            opaque ipv6[16];
//        }
//
type PeerAddressIp struct {
	Type IpAddrType
	Ipv4 *[4]byte  `xdrmaxsize:"4"`
	Ipv6 *[16]byte `xdrmaxsize:"16"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PeerAddressIp) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PeerAddressIp
func (u PeerAddressIp) ArmForSwitch(sw int32) (string, bool) {
	switch IpAddrType(sw) {
	case IpAddrTypeIPv4:
		return "Ipv4", true
	case IpAddrTypeIPv6:
		return "Ipv6", true
	}
	return "-", false
}

// NewPeerAddressIp creates a new  PeerAddressIp.
func NewPeerAddressIp(aType IpAddrType, value interface{}) (result PeerAddressIp, err error) {
	result.Type = aType
	switch IpAddrType(aType) {
	case IpAddrTypeIPv4:
		tv, ok := value.([4]byte)
		if !ok {
			err = fmt.Errorf("invalid value, must be [4]byte")
			return
		}
		result.Ipv4 = &tv
	case IpAddrTypeIPv6:
		tv, ok := value.([16]byte)
		if !ok {
			err = fmt.Errorf("invalid value, must be [16]byte")
			return
		}
		result.Ipv6 = &tv
	}
	return
}

// MustIpv4 retrieves the Ipv4 value from the union,
// panicing if the value is not set.
func (u PeerAddressIp) MustIpv4() [4]byte {
	val, ok := u.GetIpv4()

	if !ok {
		panic("arm Ipv4 is not set")
	}

	return val
}

// GetIpv4 retrieves the Ipv4 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PeerAddressIp) GetIpv4() (result [4]byte, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ipv4" {
		result = *u.Ipv4
		ok = true
	}

	return
}

// MustIpv6 retrieves the Ipv6 value from the union,
// panicing if the value is not set.
func (u PeerAddressIp) MustIpv6() [16]byte {
	val, ok := u.GetIpv6()

	if !ok {
		panic("arm Ipv6 is not set")
	}

	return val
}

// GetIpv6 retrieves the Ipv6 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PeerAddressIp) GetIpv6() (result [16]byte, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ipv6" {
		result = *u.Ipv6
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PeerAddressIp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PeerAddressIp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PeerAddressIp)(nil)
	_ encoding.BinaryUnmarshaler = (*PeerAddressIp)(nil)
)

// PeerAddress is an XDR Struct defines as:
//
//   struct PeerAddress
//    {
//        union switch (IPAddrType type)
//        {
//        case IPv4:
//            opaque ipv4[4];
//        case IPv6:
//            opaque ipv6[16];
//        }
//        ip;
//        uint32 port;
//        uint32 numFailures;
//    };
//
type PeerAddress struct {
	Ip          PeerAddressIp
	Port        Uint32
	NumFailures Uint32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PeerAddress) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PeerAddress) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PeerAddress)(nil)
	_ encoding.BinaryUnmarshaler = (*PeerAddress)(nil)
)

// MessageType is an XDR Enum defines as:
//
//   enum MessageType
//    {
//        ERROR_MSG = 0,
//        AUTH = 2,
//        DONT_HAVE = 3,
//
//        GET_PEERS = 4, // gets a list of peers this guy knows about
//        PEERS = 5,
//
//        GET_TX_SET = 6, // gets a particular txset by hash
//        TX_SET = 7,
//
//        TRANSACTION = 8, // pass on a tx you have heard about
//
//        // SCP
//        GET_SCP_QUORUMSET = 9,
//        SCP_QUORUMSET = 10,
//        SCP_MESSAGE = 11,
//        GET_SCP_STATE = 12,
//
//        // new messages
//        HELLO = 13,
//
//        SURVEY_REQUEST = 14,
//        SURVEY_RESPONSE = 15
//    };
//
type MessageType int32

const (
	MessageTypeErrorMsg        MessageType = 0
	MessageTypeAuth            MessageType = 2
	MessageTypeDontHave        MessageType = 3
	MessageTypeGetPeers        MessageType = 4
	MessageTypePeers           MessageType = 5
	MessageTypeGetTxSet        MessageType = 6
	MessageTypeTxSet           MessageType = 7
	MessageTypeTransaction     MessageType = 8
	MessageTypeGetScpQuorumset MessageType = 9
	MessageTypeScpQuorumset    MessageType = 10
	MessageTypeScpMessage      MessageType = 11
	MessageTypeGetScpState     MessageType = 12
	MessageTypeHello           MessageType = 13
	MessageTypeSurveyRequest   MessageType = 14
	MessageTypeSurveyResponse  MessageType = 15
)

var messageTypeMap = map[int32]string{
	0:  "MessageTypeErrorMsg",
	2:  "MessageTypeAuth",
	3:  "MessageTypeDontHave",
	4:  "MessageTypeGetPeers",
	5:  "MessageTypePeers",
	6:  "MessageTypeGetTxSet",
	7:  "MessageTypeTxSet",
	8:  "MessageTypeTransaction",
	9:  "MessageTypeGetScpQuorumset",
	10: "MessageTypeScpQuorumset",
	11: "MessageTypeScpMessage",
	12: "MessageTypeGetScpState",
	13: "MessageTypeHello",
	14: "MessageTypeSurveyRequest",
	15: "MessageTypeSurveyResponse",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for MessageType
func (e MessageType) ValidEnum(v int32) bool {
	_, ok := messageTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e MessageType) String() string {
	name, _ := messageTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s MessageType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *MessageType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*MessageType)(nil)
	_ encoding.BinaryUnmarshaler = (*MessageType)(nil)
)

// DontHave is an XDR Struct defines as:
//
//   struct DontHave
//    {
//        MessageType type;
//        uint256 reqHash;
//    };
//
type DontHave struct {
	Type    MessageType
	ReqHash Uint256
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s DontHave) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *DontHave) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*DontHave)(nil)
	_ encoding.BinaryUnmarshaler = (*DontHave)(nil)
)

// SurveyMessageCommandType is an XDR Enum defines as:
//
//   enum SurveyMessageCommandType
//    {
//        SURVEY_TOPOLOGY = 0
//    };
//
type SurveyMessageCommandType int32

const (
	SurveyMessageCommandTypeSurveyTopology SurveyMessageCommandType = 0
)

var surveyMessageCommandTypeMap = map[int32]string{
	0: "SurveyMessageCommandTypeSurveyTopology",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for SurveyMessageCommandType
func (e SurveyMessageCommandType) ValidEnum(v int32) bool {
	_, ok := surveyMessageCommandTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e SurveyMessageCommandType) String() string {
	name, _ := surveyMessageCommandTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SurveyMessageCommandType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SurveyMessageCommandType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SurveyMessageCommandType)(nil)
	_ encoding.BinaryUnmarshaler = (*SurveyMessageCommandType)(nil)
)

// SurveyRequestMessage is an XDR Struct defines as:
//
//   struct SurveyRequestMessage
//    {
//        NodeID surveyorPeerID;
//        NodeID surveyedPeerID;
//        uint32 ledgerNum;
//        Curve25519Public encryptionKey;
//        SurveyMessageCommandType commandType;
//    };
//
type SurveyRequestMessage struct {
	SurveyorPeerId NodeId
	SurveyedPeerId NodeId
	LedgerNum      Uint32
	EncryptionKey  Curve25519Public
	CommandType    SurveyMessageCommandType
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SurveyRequestMessage) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SurveyRequestMessage) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SurveyRequestMessage)(nil)
	_ encoding.BinaryUnmarshaler = (*SurveyRequestMessage)(nil)
)

// SignedSurveyRequestMessage is an XDR Struct defines as:
//
//   struct SignedSurveyRequestMessage
//    {
//        Signature requestSignature;
//        SurveyRequestMessage request;
//    };
//
type SignedSurveyRequestMessage struct {
	RequestSignature Signature
	Request          SurveyRequestMessage
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SignedSurveyRequestMessage) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SignedSurveyRequestMessage) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SignedSurveyRequestMessage)(nil)
	_ encoding.BinaryUnmarshaler = (*SignedSurveyRequestMessage)(nil)
)

// EncryptedBody is an XDR Typedef defines as:
//
//   typedef opaque EncryptedBody<64000>;
//
type EncryptedBody []byte

// XDRMaxSize implements the Sized interface for EncryptedBody
func (e EncryptedBody) XDRMaxSize() int {
	return 64000
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s EncryptedBody) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *EncryptedBody) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*EncryptedBody)(nil)
	_ encoding.BinaryUnmarshaler = (*EncryptedBody)(nil)
)

// SurveyResponseMessage is an XDR Struct defines as:
//
//   struct SurveyResponseMessage
//    {
//        NodeID surveyorPeerID;
//        NodeID surveyedPeerID;
//        uint32 ledgerNum;
//        SurveyMessageCommandType commandType;
//        EncryptedBody encryptedBody;
//    };
//
type SurveyResponseMessage struct {
	SurveyorPeerId NodeId
	SurveyedPeerId NodeId
	LedgerNum      Uint32
	CommandType    SurveyMessageCommandType
	EncryptedBody  EncryptedBody
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SurveyResponseMessage) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SurveyResponseMessage) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SurveyResponseMessage)(nil)
	_ encoding.BinaryUnmarshaler = (*SurveyResponseMessage)(nil)
)

// SignedSurveyResponseMessage is an XDR Struct defines as:
//
//   struct SignedSurveyResponseMessage
//    {
//        Signature responseSignature;
//        SurveyResponseMessage response;
//    };
//
type SignedSurveyResponseMessage struct {
	ResponseSignature Signature
	Response          SurveyResponseMessage
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SignedSurveyResponseMessage) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SignedSurveyResponseMessage) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SignedSurveyResponseMessage)(nil)
	_ encoding.BinaryUnmarshaler = (*SignedSurveyResponseMessage)(nil)
)

// PeerStats is an XDR Struct defines as:
//
//   struct PeerStats
//    {
//        NodeID id;
//        string versionStr<100>;
//        uint64 messagesRead;
//        uint64 messagesWritten;
//        uint64 bytesRead;
//        uint64 bytesWritten;
//        uint64 secondsConnected;
//
//        uint64 uniqueFloodBytesRecv;
//        uint64 duplicateFloodBytesRecv;
//        uint64 uniqueFetchBytesRecv;
//        uint64 duplicateFetchBytesRecv;
//
//        uint64 uniqueFloodMessageRecv;
//        uint64 duplicateFloodMessageRecv;
//        uint64 uniqueFetchMessageRecv;
//        uint64 duplicateFetchMessageRecv;
//    };
//
type PeerStats struct {
	Id                        NodeId
	VersionStr                string `xdrmaxsize:"100"`
	MessagesRead              Uint64
	MessagesWritten           Uint64
	BytesRead                 Uint64
	BytesWritten              Uint64
	SecondsConnected          Uint64
	UniqueFloodBytesRecv      Uint64
	DuplicateFloodBytesRecv   Uint64
	UniqueFetchBytesRecv      Uint64
	DuplicateFetchBytesRecv   Uint64
	UniqueFloodMessageRecv    Uint64
	DuplicateFloodMessageRecv Uint64
	UniqueFetchMessageRecv    Uint64
	DuplicateFetchMessageRecv Uint64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PeerStats) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PeerStats) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PeerStats)(nil)
	_ encoding.BinaryUnmarshaler = (*PeerStats)(nil)
)

// PeerStatList is an XDR Typedef defines as:
//
//   typedef PeerStats PeerStatList<25>;
//
type PeerStatList []PeerStats

// XDRMaxSize implements the Sized interface for PeerStatList
func (e PeerStatList) XDRMaxSize() int {
	return 25
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PeerStatList) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PeerStatList) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PeerStatList)(nil)
	_ encoding.BinaryUnmarshaler = (*PeerStatList)(nil)
)

// TopologyResponseBody is an XDR Struct defines as:
//
//   struct TopologyResponseBody
//    {
//        PeerStatList inboundPeers;
//        PeerStatList outboundPeers;
//
//        uint32 totalInboundPeerCount;
//        uint32 totalOutboundPeerCount;
//    };
//
type TopologyResponseBody struct {
	InboundPeers           PeerStatList
	OutboundPeers          PeerStatList
	TotalInboundPeerCount  Uint32
	TotalOutboundPeerCount Uint32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TopologyResponseBody) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TopologyResponseBody) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TopologyResponseBody)(nil)
	_ encoding.BinaryUnmarshaler = (*TopologyResponseBody)(nil)
)

// SurveyResponseBody is an XDR Union defines as:
//
//   union SurveyResponseBody switch (SurveyMessageCommandType type)
//    {
//    case SURVEY_TOPOLOGY:
//        TopologyResponseBody topologyResponseBody;
//    };
//
type SurveyResponseBody struct {
	Type                 SurveyMessageCommandType
	TopologyResponseBody *TopologyResponseBody
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SurveyResponseBody) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SurveyResponseBody
func (u SurveyResponseBody) ArmForSwitch(sw int32) (string, bool) {
	switch SurveyMessageCommandType(sw) {
	case SurveyMessageCommandTypeSurveyTopology:
		return "TopologyResponseBody", true
	}
	return "-", false
}

// NewSurveyResponseBody creates a new  SurveyResponseBody.
func NewSurveyResponseBody(aType SurveyMessageCommandType, value interface{}) (result SurveyResponseBody, err error) {
	result.Type = aType
	switch SurveyMessageCommandType(aType) {
	case SurveyMessageCommandTypeSurveyTopology:
		tv, ok := value.(TopologyResponseBody)
		if !ok {
			err = fmt.Errorf("invalid value, must be TopologyResponseBody")
			return
		}
		result.TopologyResponseBody = &tv
	}
	return
}

// MustTopologyResponseBody retrieves the TopologyResponseBody value from the union,
// panicing if the value is not set.
func (u SurveyResponseBody) MustTopologyResponseBody() TopologyResponseBody {
	val, ok := u.GetTopologyResponseBody()

	if !ok {
		panic("arm TopologyResponseBody is not set")
	}

	return val
}

// GetTopologyResponseBody retrieves the TopologyResponseBody value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SurveyResponseBody) GetTopologyResponseBody() (result TopologyResponseBody, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "TopologyResponseBody" {
		result = *u.TopologyResponseBody
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SurveyResponseBody) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SurveyResponseBody) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SurveyResponseBody)(nil)
	_ encoding.BinaryUnmarshaler = (*SurveyResponseBody)(nil)
)

// StellarMessage is an XDR Union defines as:
//
//   union StellarMessage switch (MessageType type)
//    {
//    case ERROR_MSG:
//        Error error;
//    case HELLO:
//        Hello hello;
//    case AUTH:
//        Auth auth;
//    case DONT_HAVE:
//        DontHave dontHave;
//    case GET_PEERS:
//        void;
//    case PEERS:
//        PeerAddress peers<100>;
//
//    case GET_TX_SET:
//        uint256 txSetHash;
//    case TX_SET:
//        TransactionSet txSet;
//
//    case TRANSACTION:
//        TransactionEnvelope transaction;
//
//    case SURVEY_REQUEST:
//        SignedSurveyRequestMessage signedSurveyRequestMessage;
//
//    case SURVEY_RESPONSE:
//        SignedSurveyResponseMessage signedSurveyResponseMessage;
//
//    // SCP
//    case GET_SCP_QUORUMSET:
//        uint256 qSetHash;
//    case SCP_QUORUMSET:
//        SCPQuorumSet qSet;
//    case SCP_MESSAGE:
//        SCPEnvelope envelope;
//    case GET_SCP_STATE:
//        uint32 getSCPLedgerSeq; // ledger seq requested ; if 0, requests the latest
//    };
//
type StellarMessage struct {
	Type                        MessageType
	Error                       *Error
	Hello                       *Hello
	Auth                        *Auth
	DontHave                    *DontHave
	Peers                       *[]PeerAddress `xdrmaxsize:"100"`
	TxSetHash                   *Uint256
	TxSet                       *TransactionSet
	Transaction                 *TransactionEnvelope
	SignedSurveyRequestMessage  *SignedSurveyRequestMessage
	SignedSurveyResponseMessage *SignedSurveyResponseMessage
	QSetHash                    *Uint256
	QSet                        *ScpQuorumSet
	Envelope                    *ScpEnvelope
	GetScpLedgerSeq             *Uint32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u StellarMessage) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of StellarMessage
func (u StellarMessage) ArmForSwitch(sw int32) (string, bool) {
	switch MessageType(sw) {
	case MessageTypeErrorMsg:
		return "Error", true
	case MessageTypeHello:
		return "Hello", true
	case MessageTypeAuth:
		return "Auth", true
	case MessageTypeDontHave:
		return "DontHave", true
	case MessageTypeGetPeers:
		return "", true
	case MessageTypePeers:
		return "Peers", true
	case MessageTypeGetTxSet:
		return "TxSetHash", true
	case MessageTypeTxSet:
		return "TxSet", true
	case MessageTypeTransaction:
		return "Transaction", true
	case MessageTypeSurveyRequest:
		return "SignedSurveyRequestMessage", true
	case MessageTypeSurveyResponse:
		return "SignedSurveyResponseMessage", true
	case MessageTypeGetScpQuorumset:
		return "QSetHash", true
	case MessageTypeScpQuorumset:
		return "QSet", true
	case MessageTypeScpMessage:
		return "Envelope", true
	case MessageTypeGetScpState:
		return "GetScpLedgerSeq", true
	}
	return "-", false
}

// NewStellarMessage creates a new  StellarMessage.
func NewStellarMessage(aType MessageType, value interface{}) (result StellarMessage, err error) {
	result.Type = aType
	switch MessageType(aType) {
	case MessageTypeErrorMsg:
		tv, ok := value.(Error)
		if !ok {
			err = fmt.Errorf("invalid value, must be Error")
			return
		}
		result.Error = &tv
	case MessageTypeHello:
		tv, ok := value.(Hello)
		if !ok {
			err = fmt.Errorf("invalid value, must be Hello")
			return
		}
		result.Hello = &tv
	case MessageTypeAuth:
		tv, ok := value.(Auth)
		if !ok {
			err = fmt.Errorf("invalid value, must be Auth")
			return
		}
		result.Auth = &tv
	case MessageTypeDontHave:
		tv, ok := value.(DontHave)
		if !ok {
			err = fmt.Errorf("invalid value, must be DontHave")
			return
		}
		result.DontHave = &tv
	case MessageTypeGetPeers:
		// void
	case MessageTypePeers:
		tv, ok := value.([]PeerAddress)
		if !ok {
			err = fmt.Errorf("invalid value, must be []PeerAddress")
			return
		}
		result.Peers = &tv
	case MessageTypeGetTxSet:
		tv, ok := value.(Uint256)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint256")
			return
		}
		result.TxSetHash = &tv
	case MessageTypeTxSet:
		tv, ok := value.(TransactionSet)
		if !ok {
			err = fmt.Errorf("invalid value, must be TransactionSet")
			return
		}
		result.TxSet = &tv
	case MessageTypeTransaction:
		tv, ok := value.(TransactionEnvelope)
		if !ok {
			err = fmt.Errorf("invalid value, must be TransactionEnvelope")
			return
		}
		result.Transaction = &tv
	case MessageTypeSurveyRequest:
		tv, ok := value.(SignedSurveyRequestMessage)
		if !ok {
			err = fmt.Errorf("invalid value, must be SignedSurveyRequestMessage")
			return
		}
		result.SignedSurveyRequestMessage = &tv
	case MessageTypeSurveyResponse:
		tv, ok := value.(SignedSurveyResponseMessage)
		if !ok {
			err = fmt.Errorf("invalid value, must be SignedSurveyResponseMessage")
			return
		}
		result.SignedSurveyResponseMessage = &tv
	case MessageTypeGetScpQuorumset:
		tv, ok := value.(Uint256)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint256")
			return
		}
		result.QSetHash = &tv
	case MessageTypeScpQuorumset:
		tv, ok := value.(ScpQuorumSet)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpQuorumSet")
			return
		}
		result.QSet = &tv
	case MessageTypeScpMessage:
		tv, ok := value.(ScpEnvelope)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpEnvelope")
			return
		}
		result.Envelope = &tv
	case MessageTypeGetScpState:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.GetScpLedgerSeq = &tv
	}
	return
}

// MustError retrieves the Error value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustError() Error {
	val, ok := u.GetError()

	if !ok {
		panic("arm Error is not set")
	}

	return val
}

// GetError retrieves the Error value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetError() (result Error, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Error" {
		result = *u.Error
		ok = true
	}

	return
}

// MustHello retrieves the Hello value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustHello() Hello {
	val, ok := u.GetHello()

	if !ok {
		panic("arm Hello is not set")
	}

	return val
}

// GetHello retrieves the Hello value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetHello() (result Hello, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Hello" {
		result = *u.Hello
		ok = true
	}

	return
}

// MustAuth retrieves the Auth value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustAuth() Auth {
	val, ok := u.GetAuth()

	if !ok {
		panic("arm Auth is not set")
	}

	return val
}

// GetAuth retrieves the Auth value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetAuth() (result Auth, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Auth" {
		result = *u.Auth
		ok = true
	}

	return
}

// MustDontHave retrieves the DontHave value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustDontHave() DontHave {
	val, ok := u.GetDontHave()

	if !ok {
		panic("arm DontHave is not set")
	}

	return val
}

// GetDontHave retrieves the DontHave value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetDontHave() (result DontHave, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "DontHave" {
		result = *u.DontHave
		ok = true
	}

	return
}

// MustPeers retrieves the Peers value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustPeers() []PeerAddress {
	val, ok := u.GetPeers()

	if !ok {
		panic("arm Peers is not set")
	}

	return val
}

// GetPeers retrieves the Peers value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetPeers() (result []PeerAddress, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Peers" {
		result = *u.Peers
		ok = true
	}

	return
}

// MustTxSetHash retrieves the TxSetHash value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustTxSetHash() Uint256 {
	val, ok := u.GetTxSetHash()

	if !ok {
		panic("arm TxSetHash is not set")
	}

	return val
}

// GetTxSetHash retrieves the TxSetHash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetTxSetHash() (result Uint256, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "TxSetHash" {
		result = *u.TxSetHash
		ok = true
	}

	return
}

// MustTxSet retrieves the TxSet value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustTxSet() TransactionSet {
	val, ok := u.GetTxSet()

	if !ok {
		panic("arm TxSet is not set")
	}

	return val
}

// GetTxSet retrieves the TxSet value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetTxSet() (result TransactionSet, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "TxSet" {
		result = *u.TxSet
		ok = true
	}

	return
}

// MustTransaction retrieves the Transaction value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustTransaction() TransactionEnvelope {
	val, ok := u.GetTransaction()

	if !ok {
		panic("arm Transaction is not set")
	}

	return val
}

// GetTransaction retrieves the Transaction value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetTransaction() (result TransactionEnvelope, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Transaction" {
		result = *u.Transaction
		ok = true
	}

	return
}

// MustSignedSurveyRequestMessage retrieves the SignedSurveyRequestMessage value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustSignedSurveyRequestMessage() SignedSurveyRequestMessage {
	val, ok := u.GetSignedSurveyRequestMessage()

	if !ok {
		panic("arm SignedSurveyRequestMessage is not set")
	}

	return val
}

// GetSignedSurveyRequestMessage retrieves the SignedSurveyRequestMessage value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetSignedSurveyRequestMessage() (result SignedSurveyRequestMessage, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SignedSurveyRequestMessage" {
		result = *u.SignedSurveyRequestMessage
		ok = true
	}

	return
}

// MustSignedSurveyResponseMessage retrieves the SignedSurveyResponseMessage value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustSignedSurveyResponseMessage() SignedSurveyResponseMessage {
	val, ok := u.GetSignedSurveyResponseMessage()

	if !ok {
		panic("arm SignedSurveyResponseMessage is not set")
	}

	return val
}

// GetSignedSurveyResponseMessage retrieves the SignedSurveyResponseMessage value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetSignedSurveyResponseMessage() (result SignedSurveyResponseMessage, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SignedSurveyResponseMessage" {
		result = *u.SignedSurveyResponseMessage
		ok = true
	}

	return
}

// MustQSetHash retrieves the QSetHash value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustQSetHash() Uint256 {
	val, ok := u.GetQSetHash()

	if !ok {
		panic("arm QSetHash is not set")
	}

	return val
}

// GetQSetHash retrieves the QSetHash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetQSetHash() (result Uint256, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "QSetHash" {
		result = *u.QSetHash
		ok = true
	}

	return
}

// MustQSet retrieves the QSet value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustQSet() ScpQuorumSet {
	val, ok := u.GetQSet()

	if !ok {
		panic("arm QSet is not set")
	}

	return val
}

// GetQSet retrieves the QSet value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetQSet() (result ScpQuorumSet, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "QSet" {
		result = *u.QSet
		ok = true
	}

	return
}

// MustEnvelope retrieves the Envelope value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustEnvelope() ScpEnvelope {
	val, ok := u.GetEnvelope()

	if !ok {
		panic("arm Envelope is not set")
	}

	return val
}

// GetEnvelope retrieves the Envelope value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetEnvelope() (result ScpEnvelope, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Envelope" {
		result = *u.Envelope
		ok = true
	}

	return
}

// MustGetScpLedgerSeq retrieves the GetScpLedgerSeq value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustGetScpLedgerSeq() Uint32 {
	val, ok := u.GetGetScpLedgerSeq()

	if !ok {
		panic("arm GetScpLedgerSeq is not set")
	}

	return val
}

// GetGetScpLedgerSeq retrieves the GetScpLedgerSeq value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetGetScpLedgerSeq() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "GetScpLedgerSeq" {
		result = *u.GetScpLedgerSeq
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s StellarMessage) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *StellarMessage) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*StellarMessage)(nil)
	_ encoding.BinaryUnmarshaler = (*StellarMessage)(nil)
)

// AuthenticatedMessageV0 is an XDR NestedStruct defines as:
//
//   struct
//        {
//            uint64 sequence;
//            StellarMessage message;
//            HmacSha256Mac mac;
//        }
//
type AuthenticatedMessageV0 struct {
	Sequence Uint64
	Message  StellarMessage
	Mac      HmacSha256Mac
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AuthenticatedMessageV0) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AuthenticatedMessageV0) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AuthenticatedMessageV0)(nil)
	_ encoding.BinaryUnmarshaler = (*AuthenticatedMessageV0)(nil)
)

// AuthenticatedMessage is an XDR Union defines as:
//
//   union AuthenticatedMessage switch (uint32 v)
//    {
//    case 0:
//        struct
//        {
//            uint64 sequence;
//            StellarMessage message;
//            HmacSha256Mac mac;
//        } v0;
//    };
//
type AuthenticatedMessage struct {
	V  Uint32
	V0 *AuthenticatedMessageV0
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AuthenticatedMessage) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AuthenticatedMessage
func (u AuthenticatedMessage) ArmForSwitch(sw int32) (string, bool) {
	switch Uint32(sw) {
	case 0:
		return "V0", true
	}
	return "-", false
}

// NewAuthenticatedMessage creates a new  AuthenticatedMessage.
func NewAuthenticatedMessage(v Uint32, value interface{}) (result AuthenticatedMessage, err error) {
	result.V = v
	switch Uint32(v) {
	case 0:
		tv, ok := value.(AuthenticatedMessageV0)
		if !ok {
			err = fmt.Errorf("invalid value, must be AuthenticatedMessageV0")
			return
		}
		result.V0 = &tv
	}
	return
}

// MustV0 retrieves the V0 value from the union,
// panicing if the value is not set.
func (u AuthenticatedMessage) MustV0() AuthenticatedMessageV0 {
	val, ok := u.GetV0()

	if !ok {
		panic("arm V0 is not set")
	}

	return val
}

// GetV0 retrieves the V0 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AuthenticatedMessage) GetV0() (result AuthenticatedMessageV0, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "V0" {
		result = *u.V0
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AuthenticatedMessage) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AuthenticatedMessage) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AuthenticatedMessage)(nil)
	_ encoding.BinaryUnmarshaler = (*AuthenticatedMessage)(nil)
)

// LiquidityPoolParameters is an XDR Union defines as:
//
//   union LiquidityPoolParameters switch (LiquidityPoolType type)
//    {
//    case LIQUIDITY_POOL_CONSTANT_PRODUCT:
//        LiquidityPoolConstantProductParameters constantProduct;
//    };
//
type LiquidityPoolParameters struct {
	Type            LiquidityPoolType
	ConstantProduct *LiquidityPoolConstantProductParameters
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LiquidityPoolParameters) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LiquidityPoolParameters
func (u LiquidityPoolParameters) ArmForSwitch(sw int32) (string, bool) {
	switch LiquidityPoolType(sw) {
	case LiquidityPoolTypeLiquidityPoolConstantProduct:
		return "ConstantProduct", true
	}
	return "-", false
}

// NewLiquidityPoolParameters creates a new  LiquidityPoolParameters.
func NewLiquidityPoolParameters(aType LiquidityPoolType, value interface{}) (result LiquidityPoolParameters, err error) {
	result.Type = aType
	switch LiquidityPoolType(aType) {
	case LiquidityPoolTypeLiquidityPoolConstantProduct:
		tv, ok := value.(LiquidityPoolConstantProductParameters)
		if !ok {
			err = fmt.Errorf("invalid value, must be LiquidityPoolConstantProductParameters")
			return
		}
		result.ConstantProduct = &tv
	}
	return
}

// MustConstantProduct retrieves the ConstantProduct value from the union,
// panicing if the value is not set.
func (u LiquidityPoolParameters) MustConstantProduct() LiquidityPoolConstantProductParameters {
	val, ok := u.GetConstantProduct()

	if !ok {
		panic("arm ConstantProduct is not set")
	}

	return val
}

// GetConstantProduct retrieves the ConstantProduct value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LiquidityPoolParameters) GetConstantProduct() (result LiquidityPoolConstantProductParameters, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ConstantProduct" {
		result = *u.ConstantProduct
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s LiquidityPoolParameters) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *LiquidityPoolParameters) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*LiquidityPoolParameters)(nil)
	_ encoding.BinaryUnmarshaler = (*LiquidityPoolParameters)(nil)
)

// MuxedAccountMed25519 is an XDR NestedStruct defines as:
//
//   struct
//        {
//            uint64 id;
//            uint256 ed25519;
//        }
//
type MuxedAccountMed25519 struct {
	Id      Uint64
	Ed25519 Uint256
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s MuxedAccountMed25519) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *MuxedAccountMed25519) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*MuxedAccountMed25519)(nil)
	_ encoding.BinaryUnmarshaler = (*MuxedAccountMed25519)(nil)
)

// MuxedAccount is an XDR Union defines as:
//
//   union MuxedAccount switch (CryptoKeyType type)
//    {
//    case KEY_TYPE_ED25519:
//        uint256 ed25519;
//    case KEY_TYPE_MUXED_ED25519:
//        struct
//        {
//            uint64 id;
//            uint256 ed25519;
//        } med25519;
//    };
//
type MuxedAccount struct {
	Type     CryptoKeyType
	Ed25519  *Uint256
	Med25519 *MuxedAccountMed25519
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u MuxedAccount) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of MuxedAccount
func (u MuxedAccount) ArmForSwitch(sw int32) (string, bool) {
	switch CryptoKeyType(sw) {
	case CryptoKeyTypeKeyTypeEd25519:
		return "Ed25519", true
	case CryptoKeyTypeKeyTypeMuxedEd25519:
		return "Med25519", true
	}
	return "-", false
}

// NewMuxedAccount creates a new  MuxedAccount.
func NewMuxedAccount(aType CryptoKeyType, value interface{}) (result MuxedAccount, err error) {
	result.Type = aType
	switch CryptoKeyType(aType) {
	case CryptoKeyTypeKeyTypeEd25519:
		tv, ok := value.(Uint256)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint256")
			return
		}
		result.Ed25519 = &tv
	case CryptoKeyTypeKeyTypeMuxedEd25519:
		tv, ok := value.(MuxedAccountMed25519)
		if !ok {
			err = fmt.Errorf("invalid value, must be MuxedAccountMed25519")
			return
		}
		result.Med25519 = &tv
	}
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u MuxedAccount) MustEd25519() Uint256 {
	val, ok := u.GetEd25519()

	if !ok {
		panic("arm Ed25519 is not set")
	}

	return val
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u MuxedAccount) GetEd25519() (result Uint256, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ed25519" {
		result = *u.Ed25519
		ok = true
	}

	return
}

// MustMed25519 retrieves the Med25519 value from the union,
// panicing if the value is not set.
func (u MuxedAccount) MustMed25519() MuxedAccountMed25519 {
	val, ok := u.GetMed25519()

	if !ok {
		panic("arm Med25519 is not set")
	}

	return val
}

// GetMed25519 retrieves the Med25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u MuxedAccount) GetMed25519() (result MuxedAccountMed25519, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Med25519" {
		result = *u.Med25519
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s MuxedAccount) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *MuxedAccount) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*MuxedAccount)(nil)
	_ encoding.BinaryUnmarshaler = (*MuxedAccount)(nil)
)

// DecoratedSignature is an XDR Struct defines as:
//
//   struct DecoratedSignature
//    {
//        SignatureHint hint;  // last 4 bytes of the public key, used as a hint
//        Signature signature; // actual signature
//    };
//
type DecoratedSignature struct {
	Hint      SignatureHint
	Signature Signature
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s DecoratedSignature) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *DecoratedSignature) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*DecoratedSignature)(nil)
	_ encoding.BinaryUnmarshaler = (*DecoratedSignature)(nil)
)

// OperationType is an XDR Enum defines as:
//
//   enum OperationType
//    {
//        CREATE_ACCOUNT = 0,
//        PAYMENT = 1,
//        PATH_PAYMENT_STRICT_RECEIVE = 2,
//        MANAGE_SELL_OFFER = 3,
//        CREATE_PASSIVE_SELL_OFFER = 4,
//        SET_OPTIONS = 5,
//        CHANGE_TRUST = 6,
//        ALLOW_TRUST = 7,
//        ACCOUNT_MERGE = 8,
//        INFLATION = 9,
//        MANAGE_DATA = 10,
//        BUMP_SEQUENCE = 11,
//        MANAGE_BUY_OFFER = 12,
//        PATH_PAYMENT_STRICT_SEND = 13,
//        CREATE_CLAIMABLE_BALANCE = 14,
//        CLAIM_CLAIMABLE_BALANCE = 15,
//        BEGIN_SPONSORING_FUTURE_RESERVES = 16,
//        END_SPONSORING_FUTURE_RESERVES = 17,
//        REVOKE_SPONSORSHIP = 18,
//        CLAWBACK = 19,
//        CLAWBACK_CLAIMABLE_BALANCE = 20,
//        SET_TRUST_LINE_FLAGS = 21
//    };
//
type OperationType int32

const (
	OperationTypeCreateAccount                 OperationType = 0
	OperationTypePayment                       OperationType = 1
	OperationTypePathPaymentStrictReceive      OperationType = 2
	OperationTypeManageSellOffer               OperationType = 3
	OperationTypeCreatePassiveSellOffer        OperationType = 4
	OperationTypeSetOptions                    OperationType = 5
	OperationTypeChangeTrust                   OperationType = 6
	OperationTypeAllowTrust                    OperationType = 7
	OperationTypeAccountMerge                  OperationType = 8
	OperationTypeInflation                     OperationType = 9
	OperationTypeManageData                    OperationType = 10
	OperationTypeBumpSequence                  OperationType = 11
	OperationTypeManageBuyOffer                OperationType = 12
	OperationTypePathPaymentStrictSend         OperationType = 13
	OperationTypeCreateClaimableBalance        OperationType = 14
	OperationTypeClaimClaimableBalance         OperationType = 15
	OperationTypeBeginSponsoringFutureReserves OperationType = 16
	OperationTypeEndSponsoringFutureReserves   OperationType = 17
	OperationTypeRevokeSponsorship             OperationType = 18
	OperationTypeClawback                      OperationType = 19
	OperationTypeClawbackClaimableBalance      OperationType = 20
	OperationTypeSetTrustLineFlags             OperationType = 21
)

var operationTypeMap = map[int32]string{
	0:  "OperationTypeCreateAccount",
	1:  "OperationTypePayment",
	2:  "OperationTypePathPaymentStrictReceive",
	3:  "OperationTypeManageSellOffer",
	4:  "OperationTypeCreatePassiveSellOffer",
	5:  "OperationTypeSetOptions",
	6:  "OperationTypeChangeTrust",
	7:  "OperationTypeAllowTrust",
	8:  "OperationTypeAccountMerge",
	9:  "OperationTypeInflation",
	10: "OperationTypeManageData",
	11: "OperationTypeBumpSequence",
	12: "OperationTypeManageBuyOffer",
	13: "OperationTypePathPaymentStrictSend",
	14: "OperationTypeCreateClaimableBalance",
	15: "OperationTypeClaimClaimableBalance",
	16: "OperationTypeBeginSponsoringFutureReserves",
	17: "OperationTypeEndSponsoringFutureReserves",
	18: "OperationTypeRevokeSponsorship",
	19: "OperationTypeClawback",
	20: "OperationTypeClawbackClaimableBalance",
	21: "OperationTypeSetTrustLineFlags",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for OperationType
func (e OperationType) ValidEnum(v int32) bool {
	_, ok := operationTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e OperationType) String() string {
	name, _ := operationTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s OperationType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *OperationType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*OperationType)(nil)
	_ encoding.BinaryUnmarshaler = (*OperationType)(nil)
)

// CreateAccountOp is an XDR Struct defines as:
//
//   struct CreateAccountOp
//    {
//        AccountID destination; // account to create
//        int64 startingBalance; // amount they end up with
//    };
//
type CreateAccountOp struct {
	Destination     AccountId
	StartingBalance Int64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s CreateAccountOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *CreateAccountOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*CreateAccountOp)(nil)
	_ encoding.BinaryUnmarshaler = (*CreateAccountOp)(nil)
)

// PaymentOp is an XDR Struct defines as:
//
//   struct PaymentOp
//    {
//        MuxedAccount destination; // recipient of the payment
//        Asset asset;              // what they end up with
//        int64 amount;             // amount they end up with
//    };
//
type PaymentOp struct {
	Destination MuxedAccount
	Asset       Asset
	Amount      Int64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PaymentOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PaymentOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PaymentOp)(nil)
	_ encoding.BinaryUnmarshaler = (*PaymentOp)(nil)
)

// PathPaymentStrictReceiveOp is an XDR Struct defines as:
//
//   struct PathPaymentStrictReceiveOp
//    {
//        Asset sendAsset; // asset we pay with
//        int64 sendMax;   // the maximum amount of sendAsset to
//                         // send (excluding fees).
//                         // The operation will fail if can't be met
//
//        MuxedAccount destination; // recipient of the payment
//        Asset destAsset;          // what they end up with
//        int64 destAmount;         // amount they end up with
//
//        Asset path<5>; // additional hops it must go through to get there
//    };
//
type PathPaymentStrictReceiveOp struct {
	SendAsset   Asset
	SendMax     Int64
	Destination MuxedAccount
	DestAsset   Asset
	DestAmount  Int64
	Path        []Asset `xdrmaxsize:"5"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PathPaymentStrictReceiveOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PathPaymentStrictReceiveOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PathPaymentStrictReceiveOp)(nil)
	_ encoding.BinaryUnmarshaler = (*PathPaymentStrictReceiveOp)(nil)
)

// PathPaymentStrictSendOp is an XDR Struct defines as:
//
//   struct PathPaymentStrictSendOp
//    {
//        Asset sendAsset;  // asset we pay with
//        int64 sendAmount; // amount of sendAsset to send (excluding fees)
//
//        MuxedAccount destination; // recipient of the payment
//        Asset destAsset;          // what they end up with
//        int64 destMin;            // the minimum amount of dest asset to
//                                  // be received
//                                  // The operation will fail if it can't be met
//
//        Asset path<5>; // additional hops it must go through to get there
//    };
//
type PathPaymentStrictSendOp struct {
	SendAsset   Asset
	SendAmount  Int64
	Destination MuxedAccount
	DestAsset   Asset
	DestMin     Int64
	Path        []Asset `xdrmaxsize:"5"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PathPaymentStrictSendOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PathPaymentStrictSendOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PathPaymentStrictSendOp)(nil)
	_ encoding.BinaryUnmarshaler = (*PathPaymentStrictSendOp)(nil)
)

// ManageSellOfferOp is an XDR Struct defines as:
//
//   struct ManageSellOfferOp
//    {
//        Asset selling;
//        Asset buying;
//        int64 amount; // amount being sold. if set to 0, delete the offer
//        Price price;  // price of thing being sold in terms of what you are buying
//
//        // 0=create a new offer, otherwise edit an existing offer
//        int64 offerID;
//    };
//
type ManageSellOfferOp struct {
	Selling Asset
	Buying  Asset
	Amount  Int64
	Price   Price
	OfferId Int64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ManageSellOfferOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ManageSellOfferOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ManageSellOfferOp)(nil)
	_ encoding.BinaryUnmarshaler = (*ManageSellOfferOp)(nil)
)

// ManageBuyOfferOp is an XDR Struct defines as:
//
//   struct ManageBuyOfferOp
//    {
//        Asset selling;
//        Asset buying;
//        int64 buyAmount; // amount being bought. if set to 0, delete the offer
//        Price price;     // price of thing being bought in terms of what you are
//                         // selling
//
//        // 0=create a new offer, otherwise edit an existing offer
//        int64 offerID;
//    };
//
type ManageBuyOfferOp struct {
	Selling   Asset
	Buying    Asset
	BuyAmount Int64
	Price     Price
	OfferId   Int64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ManageBuyOfferOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ManageBuyOfferOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ManageBuyOfferOp)(nil)
	_ encoding.BinaryUnmarshaler = (*ManageBuyOfferOp)(nil)
)

// CreatePassiveSellOfferOp is an XDR Struct defines as:
//
//   struct CreatePassiveSellOfferOp
//    {
//        Asset selling; // A
//        Asset buying;  // B
//        int64 amount;  // amount taker gets
//        Price price;   // cost of A in terms of B
//    };
//
type CreatePassiveSellOfferOp struct {
	Selling Asset
	Buying  Asset
	Amount  Int64
	Price   Price
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s CreatePassiveSellOfferOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *CreatePassiveSellOfferOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*CreatePassiveSellOfferOp)(nil)
	_ encoding.BinaryUnmarshaler = (*CreatePassiveSellOfferOp)(nil)
)

// SetOptionsOp is an XDR Struct defines as:
//
//   struct SetOptionsOp
//    {
//        AccountID* inflationDest; // sets the inflation destination
//
//        uint32* clearFlags; // which flags to clear
//        uint32* setFlags;   // which flags to set
//
//        // account threshold manipulation
//        uint32* masterWeight; // weight of the master account
//        uint32* lowThreshold;
//        uint32* medThreshold;
//        uint32* highThreshold;
//
//        string32* homeDomain; // sets the home domain
//
//        // Add, update or remove a signer for the account
//        // signer is deleted if the weight is 0
//        Signer* signer;
//    };
//
type SetOptionsOp struct {
	InflationDest *AccountId
	ClearFlags    *Uint32
	SetFlags      *Uint32
	MasterWeight  *Uint32
	LowThreshold  *Uint32
	MedThreshold  *Uint32
	HighThreshold *Uint32
	HomeDomain    *String32
	Signer        *Signer
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SetOptionsOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SetOptionsOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SetOptionsOp)(nil)
	_ encoding.BinaryUnmarshaler = (*SetOptionsOp)(nil)
)

// ChangeTrustAsset is an XDR Union defines as:
//
//   union ChangeTrustAsset switch (AssetType type)
//    {
//    case ASSET_TYPE_NATIVE: // Not credit
//        void;
//
//    case ASSET_TYPE_CREDIT_ALPHANUM4:
//        AlphaNum4 alphaNum4;
//
//    case ASSET_TYPE_CREDIT_ALPHANUM12:
//        AlphaNum12 alphaNum12;
//
//    case ASSET_TYPE_POOL_SHARE:
//        LiquidityPoolParameters liquidityPool;
//
//        // add other asset types here in the future
//    };
//
type ChangeTrustAsset struct {
	Type          AssetType
	AlphaNum4     *AlphaNum4
	AlphaNum12    *AlphaNum12
	LiquidityPool *LiquidityPoolParameters
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ChangeTrustAsset) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ChangeTrustAsset
func (u ChangeTrustAsset) ArmForSwitch(sw int32) (string, bool) {
	switch AssetType(sw) {
	case AssetTypeAssetTypeNative:
		return "", true
	case AssetTypeAssetTypeCreditAlphanum4:
		return "AlphaNum4", true
	case AssetTypeAssetTypeCreditAlphanum12:
		return "AlphaNum12", true
	case AssetTypeAssetTypePoolShare:
		return "LiquidityPool", true
	}
	return "-", false
}

// NewChangeTrustAsset creates a new  ChangeTrustAsset.
func NewChangeTrustAsset(aType AssetType, value interface{}) (result ChangeTrustAsset, err error) {
	result.Type = aType
	switch AssetType(aType) {
	case AssetTypeAssetTypeNative:
		// void
	case AssetTypeAssetTypeCreditAlphanum4:
		tv, ok := value.(AlphaNum4)
		if !ok {
			err = fmt.Errorf("invalid value, must be AlphaNum4")
			return
		}
		result.AlphaNum4 = &tv
	case AssetTypeAssetTypeCreditAlphanum12:
		tv, ok := value.(AlphaNum12)
		if !ok {
			err = fmt.Errorf("invalid value, must be AlphaNum12")
			return
		}
		result.AlphaNum12 = &tv
	case AssetTypeAssetTypePoolShare:
		tv, ok := value.(LiquidityPoolParameters)
		if !ok {
			err = fmt.Errorf("invalid value, must be LiquidityPoolParameters")
			return
		}
		result.LiquidityPool = &tv
	}
	return
}

// MustAlphaNum4 retrieves the AlphaNum4 value from the union,
// panicing if the value is not set.
func (u ChangeTrustAsset) MustAlphaNum4() AlphaNum4 {
	val, ok := u.GetAlphaNum4()

	if !ok {
		panic("arm AlphaNum4 is not set")
	}

	return val
}

// GetAlphaNum4 retrieves the AlphaNum4 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ChangeTrustAsset) GetAlphaNum4() (result AlphaNum4, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AlphaNum4" {
		result = *u.AlphaNum4
		ok = true
	}

	return
}

// MustAlphaNum12 retrieves the AlphaNum12 value from the union,
// panicing if the value is not set.
func (u ChangeTrustAsset) MustAlphaNum12() AlphaNum12 {
	val, ok := u.GetAlphaNum12()

	if !ok {
		panic("arm AlphaNum12 is not set")
	}

	return val
}

// GetAlphaNum12 retrieves the AlphaNum12 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ChangeTrustAsset) GetAlphaNum12() (result AlphaNum12, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AlphaNum12" {
		result = *u.AlphaNum12
		ok = true
	}

	return
}

// MustLiquidityPool retrieves the LiquidityPool value from the union,
// panicing if the value is not set.
func (u ChangeTrustAsset) MustLiquidityPool() LiquidityPoolParameters {
	val, ok := u.GetLiquidityPool()

	if !ok {
		panic("arm LiquidityPool is not set")
	}

	return val
}

// GetLiquidityPool retrieves the LiquidityPool value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ChangeTrustAsset) GetLiquidityPool() (result LiquidityPoolParameters, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "LiquidityPool" {
		result = *u.LiquidityPool
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ChangeTrustAsset) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ChangeTrustAsset) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ChangeTrustAsset)(nil)
	_ encoding.BinaryUnmarshaler = (*ChangeTrustAsset)(nil)
)

// ChangeTrustOp is an XDR Struct defines as:
//
//   struct ChangeTrustOp
//    {
//        ChangeTrustAsset line;
//
//        // if limit is set to 0, deletes the trust line
//        int64 limit;
//    };
//
type ChangeTrustOp struct {
	Line  ChangeTrustAsset
	Limit Int64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ChangeTrustOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ChangeTrustOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ChangeTrustOp)(nil)
	_ encoding.BinaryUnmarshaler = (*ChangeTrustOp)(nil)
)

// AllowTrustOp is an XDR Struct defines as:
//
//   struct AllowTrustOp
//    {
//        AccountID trustor;
//        AssetCode asset;
//
//        // One of 0, AUTHORIZED_FLAG, or AUTHORIZED_TO_MAINTAIN_LIABILITIES_FLAG
//        uint32 authorize;
//    };
//
type AllowTrustOp struct {
	Trustor   AccountId
	Asset     AssetCode
	Authorize Uint32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AllowTrustOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AllowTrustOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AllowTrustOp)(nil)
	_ encoding.BinaryUnmarshaler = (*AllowTrustOp)(nil)
)

// ManageDataOp is an XDR Struct defines as:
//
//   struct ManageDataOp
//    {
//        string64 dataName;
//        DataValue* dataValue; // set to null to clear
//    };
//
type ManageDataOp struct {
	DataName  String64
	DataValue *DataValue
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ManageDataOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ManageDataOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ManageDataOp)(nil)
	_ encoding.BinaryUnmarshaler = (*ManageDataOp)(nil)
)

// BumpSequenceOp is an XDR Struct defines as:
//
//   struct BumpSequenceOp
//    {
//        SequenceNumber bumpTo;
//    };
//
type BumpSequenceOp struct {
	BumpTo SequenceNumber
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BumpSequenceOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BumpSequenceOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BumpSequenceOp)(nil)
	_ encoding.BinaryUnmarshaler = (*BumpSequenceOp)(nil)
)

// CreateClaimableBalanceOp is an XDR Struct defines as:
//
//   struct CreateClaimableBalanceOp
//    {
//        Asset asset;
//        int64 amount;
//        Claimant claimants<10>;
//    };
//
type CreateClaimableBalanceOp struct {
	Asset     Asset
	Amount    Int64
	Claimants []Claimant `xdrmaxsize:"10"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s CreateClaimableBalanceOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *CreateClaimableBalanceOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*CreateClaimableBalanceOp)(nil)
	_ encoding.BinaryUnmarshaler = (*CreateClaimableBalanceOp)(nil)
)

// ClaimClaimableBalanceOp is an XDR Struct defines as:
//
//   struct ClaimClaimableBalanceOp
//    {
//        ClaimableBalanceID balanceID;
//    };
//
type ClaimClaimableBalanceOp struct {
	BalanceId ClaimableBalanceId
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimClaimableBalanceOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimClaimableBalanceOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimClaimableBalanceOp)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimClaimableBalanceOp)(nil)
)

// BeginSponsoringFutureReservesOp is an XDR Struct defines as:
//
//   struct BeginSponsoringFutureReservesOp
//    {
//        AccountID sponsoredID;
//    };
//
type BeginSponsoringFutureReservesOp struct {
	SponsoredId AccountId
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BeginSponsoringFutureReservesOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BeginSponsoringFutureReservesOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BeginSponsoringFutureReservesOp)(nil)
	_ encoding.BinaryUnmarshaler = (*BeginSponsoringFutureReservesOp)(nil)
)

// RevokeSponsorshipType is an XDR Enum defines as:
//
//   enum RevokeSponsorshipType
//    {
//        REVOKE_SPONSORSHIP_LEDGER_ENTRY = 0,
//        REVOKE_SPONSORSHIP_SIGNER = 1
//    };
//
type RevokeSponsorshipType int32

const (
	RevokeSponsorshipTypeRevokeSponsorshipLedgerEntry RevokeSponsorshipType = 0
	RevokeSponsorshipTypeRevokeSponsorshipSigner      RevokeSponsorshipType = 1
)

var revokeSponsorshipTypeMap = map[int32]string{
	0: "RevokeSponsorshipTypeRevokeSponsorshipLedgerEntry",
	1: "RevokeSponsorshipTypeRevokeSponsorshipSigner",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for RevokeSponsorshipType
func (e RevokeSponsorshipType) ValidEnum(v int32) bool {
	_, ok := revokeSponsorshipTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e RevokeSponsorshipType) String() string {
	name, _ := revokeSponsorshipTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s RevokeSponsorshipType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *RevokeSponsorshipType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*RevokeSponsorshipType)(nil)
	_ encoding.BinaryUnmarshaler = (*RevokeSponsorshipType)(nil)
)

// RevokeSponsorshipOpSigner is an XDR NestedStruct defines as:
//
//   struct
//        {
//            AccountID accountID;
//            SignerKey signerKey;
//        }
//
type RevokeSponsorshipOpSigner struct {
	AccountId AccountId
	SignerKey SignerKey
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s RevokeSponsorshipOpSigner) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *RevokeSponsorshipOpSigner) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*RevokeSponsorshipOpSigner)(nil)
	_ encoding.BinaryUnmarshaler = (*RevokeSponsorshipOpSigner)(nil)
)

// RevokeSponsorshipOp is an XDR Union defines as:
//
//   union RevokeSponsorshipOp switch (RevokeSponsorshipType type)
//    {
//    case REVOKE_SPONSORSHIP_LEDGER_ENTRY:
//        LedgerKey ledgerKey;
//    case REVOKE_SPONSORSHIP_SIGNER:
//        struct
//        {
//            AccountID accountID;
//            SignerKey signerKey;
//        } signer;
//    };
//
type RevokeSponsorshipOp struct {
	Type      RevokeSponsorshipType
	LedgerKey *LedgerKey
	Signer    *RevokeSponsorshipOpSigner
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RevokeSponsorshipOp) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RevokeSponsorshipOp
func (u RevokeSponsorshipOp) ArmForSwitch(sw int32) (string, bool) {
	switch RevokeSponsorshipType(sw) {
	case RevokeSponsorshipTypeRevokeSponsorshipLedgerEntry:
		return "LedgerKey", true
	case RevokeSponsorshipTypeRevokeSponsorshipSigner:
		return "Signer", true
	}
	return "-", false
}

// NewRevokeSponsorshipOp creates a new  RevokeSponsorshipOp.
func NewRevokeSponsorshipOp(aType RevokeSponsorshipType, value interface{}) (result RevokeSponsorshipOp, err error) {
	result.Type = aType
	switch RevokeSponsorshipType(aType) {
	case RevokeSponsorshipTypeRevokeSponsorshipLedgerEntry:
		tv, ok := value.(LedgerKey)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKey")
			return
		}
		result.LedgerKey = &tv
	case RevokeSponsorshipTypeRevokeSponsorshipSigner:
		tv, ok := value.(RevokeSponsorshipOpSigner)
		if !ok {
			err = fmt.Errorf("invalid value, must be RevokeSponsorshipOpSigner")
			return
		}
		result.Signer = &tv
	}
	return
}

// MustLedgerKey retrieves the LedgerKey value from the union,
// panicing if the value is not set.
func (u RevokeSponsorshipOp) MustLedgerKey() LedgerKey {
	val, ok := u.GetLedgerKey()

	if !ok {
		panic("arm LedgerKey is not set")
	}

	return val
}

// GetLedgerKey retrieves the LedgerKey value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RevokeSponsorshipOp) GetLedgerKey() (result LedgerKey, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "LedgerKey" {
		result = *u.LedgerKey
		ok = true
	}

	return
}

// MustSigner retrieves the Signer value from the union,
// panicing if the value is not set.
func (u RevokeSponsorshipOp) MustSigner() RevokeSponsorshipOpSigner {
	val, ok := u.GetSigner()

	if !ok {
		panic("arm Signer is not set")
	}

	return val
}

// GetSigner retrieves the Signer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RevokeSponsorshipOp) GetSigner() (result RevokeSponsorshipOpSigner, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Signer" {
		result = *u.Signer
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s RevokeSponsorshipOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *RevokeSponsorshipOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*RevokeSponsorshipOp)(nil)
	_ encoding.BinaryUnmarshaler = (*RevokeSponsorshipOp)(nil)
)

// ClawbackOp is an XDR Struct defines as:
//
//   struct ClawbackOp
//    {
//        Asset asset;
//        MuxedAccount from;
//        int64 amount;
//    };
//
type ClawbackOp struct {
	Asset  Asset
	From   MuxedAccount
	Amount Int64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClawbackOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClawbackOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClawbackOp)(nil)
	_ encoding.BinaryUnmarshaler = (*ClawbackOp)(nil)
)

// ClawbackClaimableBalanceOp is an XDR Struct defines as:
//
//   struct ClawbackClaimableBalanceOp
//    {
//        ClaimableBalanceID balanceID;
//    };
//
type ClawbackClaimableBalanceOp struct {
	BalanceId ClaimableBalanceId
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClawbackClaimableBalanceOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClawbackClaimableBalanceOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClawbackClaimableBalanceOp)(nil)
	_ encoding.BinaryUnmarshaler = (*ClawbackClaimableBalanceOp)(nil)
)

// SetTrustLineFlagsOp is an XDR Struct defines as:
//
//   struct SetTrustLineFlagsOp
//    {
//        AccountID trustor;
//        Asset asset;
//
//        uint32 clearFlags; // which flags to clear
//        uint32 setFlags;   // which flags to set
//    };
//
type SetTrustLineFlagsOp struct {
	Trustor    AccountId
	Asset      Asset
	ClearFlags Uint32
	SetFlags   Uint32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SetTrustLineFlagsOp) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SetTrustLineFlagsOp) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SetTrustLineFlagsOp)(nil)
	_ encoding.BinaryUnmarshaler = (*SetTrustLineFlagsOp)(nil)
)

// LiquidityPoolFeeV18 is an XDR Const defines as:
//
//   const LIQUIDITY_POOL_FEE_V18 = 30;
//
const LiquidityPoolFeeV18 = 30

// OperationBody is an XDR NestedUnion defines as:
//
//   union switch (OperationType type)
//        {
//        case CREATE_ACCOUNT:
//            CreateAccountOp createAccountOp;
//        case PAYMENT:
//            PaymentOp paymentOp;
//        case PATH_PAYMENT_STRICT_RECEIVE:
//            PathPaymentStrictReceiveOp pathPaymentStrictReceiveOp;
//        case MANAGE_SELL_OFFER:
//            ManageSellOfferOp manageSellOfferOp;
//        case CREATE_PASSIVE_SELL_OFFER:
//            CreatePassiveSellOfferOp createPassiveSellOfferOp;
//        case SET_OPTIONS:
//            SetOptionsOp setOptionsOp;
//        case CHANGE_TRUST:
//            ChangeTrustOp changeTrustOp;
//        case ALLOW_TRUST:
//            AllowTrustOp allowTrustOp;
//        case ACCOUNT_MERGE:
//            MuxedAccount destination;
//        case INFLATION:
//            void;
//        case MANAGE_DATA:
//            ManageDataOp manageDataOp;
//        case BUMP_SEQUENCE:
//            BumpSequenceOp bumpSequenceOp;
//        case MANAGE_BUY_OFFER:
//            ManageBuyOfferOp manageBuyOfferOp;
//        case PATH_PAYMENT_STRICT_SEND:
//            PathPaymentStrictSendOp pathPaymentStrictSendOp;
//        case CREATE_CLAIMABLE_BALANCE:
//            CreateClaimableBalanceOp createClaimableBalanceOp;
//        case CLAIM_CLAIMABLE_BALANCE:
//            ClaimClaimableBalanceOp claimClaimableBalanceOp;
//        case BEGIN_SPONSORING_FUTURE_RESERVES:
//            BeginSponsoringFutureReservesOp beginSponsoringFutureReservesOp;
//        case END_SPONSORING_FUTURE_RESERVES:
//            void;
//        case REVOKE_SPONSORSHIP:
//            RevokeSponsorshipOp revokeSponsorshipOp;
//        case CLAWBACK:
//            ClawbackOp clawbackOp;
//        case CLAWBACK_CLAIMABLE_BALANCE:
//            ClawbackClaimableBalanceOp clawbackClaimableBalanceOp;
//        case SET_TRUST_LINE_FLAGS:
//            SetTrustLineFlagsOp setTrustLineFlagsOp;
//        }
//
type OperationBody struct {
	Type                            OperationType
	CreateAccountOp                 *CreateAccountOp
	PaymentOp                       *PaymentOp
	PathPaymentStrictReceiveOp      *PathPaymentStrictReceiveOp
	ManageSellOfferOp               *ManageSellOfferOp
	CreatePassiveSellOfferOp        *CreatePassiveSellOfferOp
	SetOptionsOp                    *SetOptionsOp
	ChangeTrustOp                   *ChangeTrustOp
	AllowTrustOp                    *AllowTrustOp
	Destination                     *MuxedAccount
	ManageDataOp                    *ManageDataOp
	BumpSequenceOp                  *BumpSequenceOp
	ManageBuyOfferOp                *ManageBuyOfferOp
	PathPaymentStrictSendOp         *PathPaymentStrictSendOp
	CreateClaimableBalanceOp        *CreateClaimableBalanceOp
	ClaimClaimableBalanceOp         *ClaimClaimableBalanceOp
	BeginSponsoringFutureReservesOp *BeginSponsoringFutureReservesOp
	RevokeSponsorshipOp             *RevokeSponsorshipOp
	ClawbackOp                      *ClawbackOp
	ClawbackClaimableBalanceOp      *ClawbackClaimableBalanceOp
	SetTrustLineFlagsOp             *SetTrustLineFlagsOp
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OperationBody) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OperationBody
func (u OperationBody) ArmForSwitch(sw int32) (string, bool) {
	switch OperationType(sw) {
	case OperationTypeCreateAccount:
		return "CreateAccountOp", true
	case OperationTypePayment:
		return "PaymentOp", true
	case OperationTypePathPaymentStrictReceive:
		return "PathPaymentStrictReceiveOp", true
	case OperationTypeManageSellOffer:
		return "ManageSellOfferOp", true
	case OperationTypeCreatePassiveSellOffer:
		return "CreatePassiveSellOfferOp", true
	case OperationTypeSetOptions:
		return "SetOptionsOp", true
	case OperationTypeChangeTrust:
		return "ChangeTrustOp", true
	case OperationTypeAllowTrust:
		return "AllowTrustOp", true
	case OperationTypeAccountMerge:
		return "Destination", true
	case OperationTypeInflation:
		return "", true
	case OperationTypeManageData:
		return "ManageDataOp", true
	case OperationTypeBumpSequence:
		return "BumpSequenceOp", true
	case OperationTypeManageBuyOffer:
		return "ManageBuyOfferOp", true
	case OperationTypePathPaymentStrictSend:
		return "PathPaymentStrictSendOp", true
	case OperationTypeCreateClaimableBalance:
		return "CreateClaimableBalanceOp", true
	case OperationTypeClaimClaimableBalance:
		return "ClaimClaimableBalanceOp", true
	case OperationTypeBeginSponsoringFutureReserves:
		return "BeginSponsoringFutureReservesOp", true
	case OperationTypeEndSponsoringFutureReserves:
		return "", true
	case OperationTypeRevokeSponsorship:
		return "RevokeSponsorshipOp", true
	case OperationTypeClawback:
		return "ClawbackOp", true
	case OperationTypeClawbackClaimableBalance:
		return "ClawbackClaimableBalanceOp", true
	case OperationTypeSetTrustLineFlags:
		return "SetTrustLineFlagsOp", true
	}
	return "-", false
}

// NewOperationBody creates a new  OperationBody.
func NewOperationBody(aType OperationType, value interface{}) (result OperationBody, err error) {
	result.Type = aType
	switch OperationType(aType) {
	case OperationTypeCreateAccount:
		tv, ok := value.(CreateAccountOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAccountOp")
			return
		}
		result.CreateAccountOp = &tv
	case OperationTypePayment:
		tv, ok := value.(PaymentOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentOp")
			return
		}
		result.PaymentOp = &tv
	case OperationTypePathPaymentStrictReceive:
		tv, ok := value.(PathPaymentStrictReceiveOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be PathPaymentStrictReceiveOp")
			return
		}
		result.PathPaymentStrictReceiveOp = &tv
	case OperationTypeManageSellOffer:
		tv, ok := value.(ManageSellOfferOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageSellOfferOp")
			return
		}
		result.ManageSellOfferOp = &tv
	case OperationTypeCreatePassiveSellOffer:
		tv, ok := value.(CreatePassiveSellOfferOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreatePassiveSellOfferOp")
			return
		}
		result.CreatePassiveSellOfferOp = &tv
	case OperationTypeSetOptions:
		tv, ok := value.(SetOptionsOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetOptionsOp")
			return
		}
		result.SetOptionsOp = &tv
	case OperationTypeChangeTrust:
		tv, ok := value.(ChangeTrustOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ChangeTrustOp")
			return
		}
		result.ChangeTrustOp = &tv
	case OperationTypeAllowTrust:
		tv, ok := value.(AllowTrustOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be AllowTrustOp")
			return
		}
		result.AllowTrustOp = &tv
	case OperationTypeAccountMerge:
		tv, ok := value.(MuxedAccount)
		if !ok {
			err = fmt.Errorf("invalid value, must be MuxedAccount")
			return
		}
		result.Destination = &tv
	case OperationTypeInflation:
		// void
	case OperationTypeManageData:
		tv, ok := value.(ManageDataOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageDataOp")
			return
		}
		result.ManageDataOp = &tv
	case OperationTypeBumpSequence:
		tv, ok := value.(BumpSequenceOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be BumpSequenceOp")
			return
		}
		result.BumpSequenceOp = &tv
	case OperationTypeManageBuyOffer:
		tv, ok := value.(ManageBuyOfferOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageBuyOfferOp")
			return
		}
		result.ManageBuyOfferOp = &tv
	case OperationTypePathPaymentStrictSend:
		tv, ok := value.(PathPaymentStrictSendOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be PathPaymentStrictSendOp")
			return
		}
		result.PathPaymentStrictSendOp = &tv
	case OperationTypeCreateClaimableBalance:
		tv, ok := value.(CreateClaimableBalanceOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateClaimableBalanceOp")
			return
		}
		result.CreateClaimableBalanceOp = &tv
	case OperationTypeClaimClaimableBalance:
		tv, ok := value.(ClaimClaimableBalanceOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ClaimClaimableBalanceOp")
			return
		}
		result.ClaimClaimableBalanceOp = &tv
	case OperationTypeBeginSponsoringFutureReserves:
		tv, ok := value.(BeginSponsoringFutureReservesOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be BeginSponsoringFutureReservesOp")
			return
		}
		result.BeginSponsoringFutureReservesOp = &tv
	case OperationTypeEndSponsoringFutureReserves:
		// void
	case OperationTypeRevokeSponsorship:
		tv, ok := value.(RevokeSponsorshipOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be RevokeSponsorshipOp")
			return
		}
		result.RevokeSponsorshipOp = &tv
	case OperationTypeClawback:
		tv, ok := value.(ClawbackOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ClawbackOp")
			return
		}
		result.ClawbackOp = &tv
	case OperationTypeClawbackClaimableBalance:
		tv, ok := value.(ClawbackClaimableBalanceOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ClawbackClaimableBalanceOp")
			return
		}
		result.ClawbackClaimableBalanceOp = &tv
	case OperationTypeSetTrustLineFlags:
		tv, ok := value.(SetTrustLineFlagsOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetTrustLineFlagsOp")
			return
		}
		result.SetTrustLineFlagsOp = &tv
	}
	return
}

// MustCreateAccountOp retrieves the CreateAccountOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateAccountOp() CreateAccountOp {
	val, ok := u.GetCreateAccountOp()

	if !ok {
		panic("arm CreateAccountOp is not set")
	}

	return val
}

// GetCreateAccountOp retrieves the CreateAccountOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateAccountOp() (result CreateAccountOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateAccountOp" {
		result = *u.CreateAccountOp
		ok = true
	}

	return
}

// MustPaymentOp retrieves the PaymentOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustPaymentOp() PaymentOp {
	val, ok := u.GetPaymentOp()

	if !ok {
		panic("arm PaymentOp is not set")
	}

	return val
}

// GetPaymentOp retrieves the PaymentOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetPaymentOp() (result PaymentOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PaymentOp" {
		result = *u.PaymentOp
		ok = true
	}

	return
}

// MustPathPaymentStrictReceiveOp retrieves the PathPaymentStrictReceiveOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustPathPaymentStrictReceiveOp() PathPaymentStrictReceiveOp {
	val, ok := u.GetPathPaymentStrictReceiveOp()

	if !ok {
		panic("arm PathPaymentStrictReceiveOp is not set")
	}

	return val
}

// GetPathPaymentStrictReceiveOp retrieves the PathPaymentStrictReceiveOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetPathPaymentStrictReceiveOp() (result PathPaymentStrictReceiveOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PathPaymentStrictReceiveOp" {
		result = *u.PathPaymentStrictReceiveOp
		ok = true
	}

	return
}

// MustManageSellOfferOp retrieves the ManageSellOfferOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageSellOfferOp() ManageSellOfferOp {
	val, ok := u.GetManageSellOfferOp()

	if !ok {
		panic("arm ManageSellOfferOp is not set")
	}

	return val
}

// GetManageSellOfferOp retrieves the ManageSellOfferOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageSellOfferOp() (result ManageSellOfferOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageSellOfferOp" {
		result = *u.ManageSellOfferOp
		ok = true
	}

	return
}

// MustCreatePassiveSellOfferOp retrieves the CreatePassiveSellOfferOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreatePassiveSellOfferOp() CreatePassiveSellOfferOp {
	val, ok := u.GetCreatePassiveSellOfferOp()

	if !ok {
		panic("arm CreatePassiveSellOfferOp is not set")
	}

	return val
}

// GetCreatePassiveSellOfferOp retrieves the CreatePassiveSellOfferOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreatePassiveSellOfferOp() (result CreatePassiveSellOfferOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreatePassiveSellOfferOp" {
		result = *u.CreatePassiveSellOfferOp
		ok = true
	}

	return
}

// MustSetOptionsOp retrieves the SetOptionsOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustSetOptionsOp() SetOptionsOp {
	val, ok := u.GetSetOptionsOp()

	if !ok {
		panic("arm SetOptionsOp is not set")
	}

	return val
}

// GetSetOptionsOp retrieves the SetOptionsOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetSetOptionsOp() (result SetOptionsOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SetOptionsOp" {
		result = *u.SetOptionsOp
		ok = true
	}

	return
}

// MustChangeTrustOp retrieves the ChangeTrustOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustChangeTrustOp() ChangeTrustOp {
	val, ok := u.GetChangeTrustOp()

	if !ok {
		panic("arm ChangeTrustOp is not set")
	}

	return val
}

// GetChangeTrustOp retrieves the ChangeTrustOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetChangeTrustOp() (result ChangeTrustOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ChangeTrustOp" {
		result = *u.ChangeTrustOp
		ok = true
	}

	return
}

// MustAllowTrustOp retrieves the AllowTrustOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustAllowTrustOp() AllowTrustOp {
	val, ok := u.GetAllowTrustOp()

	if !ok {
		panic("arm AllowTrustOp is not set")
	}

	return val
}

// GetAllowTrustOp retrieves the AllowTrustOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetAllowTrustOp() (result AllowTrustOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AllowTrustOp" {
		result = *u.AllowTrustOp
		ok = true
	}

	return
}

// MustDestination retrieves the Destination value from the union,
// panicing if the value is not set.
func (u OperationBody) MustDestination() MuxedAccount {
	val, ok := u.GetDestination()

	if !ok {
		panic("arm Destination is not set")
	}

	return val
}

// GetDestination retrieves the Destination value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetDestination() (result MuxedAccount, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Destination" {
		result = *u.Destination
		ok = true
	}

	return
}

// MustManageDataOp retrieves the ManageDataOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageDataOp() ManageDataOp {
	val, ok := u.GetManageDataOp()

	if !ok {
		panic("arm ManageDataOp is not set")
	}

	return val
}

// GetManageDataOp retrieves the ManageDataOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageDataOp() (result ManageDataOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageDataOp" {
		result = *u.ManageDataOp
		ok = true
	}

	return
}

// MustBumpSequenceOp retrieves the BumpSequenceOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustBumpSequenceOp() BumpSequenceOp {
	val, ok := u.GetBumpSequenceOp()

	if !ok {
		panic("arm BumpSequenceOp is not set")
	}

	return val
}

// GetBumpSequenceOp retrieves the BumpSequenceOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetBumpSequenceOp() (result BumpSequenceOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "BumpSequenceOp" {
		result = *u.BumpSequenceOp
		ok = true
	}

	return
}

// MustManageBuyOfferOp retrieves the ManageBuyOfferOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageBuyOfferOp() ManageBuyOfferOp {
	val, ok := u.GetManageBuyOfferOp()

	if !ok {
		panic("arm ManageBuyOfferOp is not set")
	}

	return val
}

// GetManageBuyOfferOp retrieves the ManageBuyOfferOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageBuyOfferOp() (result ManageBuyOfferOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageBuyOfferOp" {
		result = *u.ManageBuyOfferOp
		ok = true
	}

	return
}

// MustPathPaymentStrictSendOp retrieves the PathPaymentStrictSendOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustPathPaymentStrictSendOp() PathPaymentStrictSendOp {
	val, ok := u.GetPathPaymentStrictSendOp()

	if !ok {
		panic("arm PathPaymentStrictSendOp is not set")
	}

	return val
}

// GetPathPaymentStrictSendOp retrieves the PathPaymentStrictSendOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetPathPaymentStrictSendOp() (result PathPaymentStrictSendOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PathPaymentStrictSendOp" {
		result = *u.PathPaymentStrictSendOp
		ok = true
	}

	return
}

// MustCreateClaimableBalanceOp retrieves the CreateClaimableBalanceOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateClaimableBalanceOp() CreateClaimableBalanceOp {
	val, ok := u.GetCreateClaimableBalanceOp()

	if !ok {
		panic("arm CreateClaimableBalanceOp is not set")
	}

	return val
}

// GetCreateClaimableBalanceOp retrieves the CreateClaimableBalanceOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateClaimableBalanceOp() (result CreateClaimableBalanceOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateClaimableBalanceOp" {
		result = *u.CreateClaimableBalanceOp
		ok = true
	}

	return
}

// MustClaimClaimableBalanceOp retrieves the ClaimClaimableBalanceOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustClaimClaimableBalanceOp() ClaimClaimableBalanceOp {
	val, ok := u.GetClaimClaimableBalanceOp()

	if !ok {
		panic("arm ClaimClaimableBalanceOp is not set")
	}

	return val
}

// GetClaimClaimableBalanceOp retrieves the ClaimClaimableBalanceOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetClaimClaimableBalanceOp() (result ClaimClaimableBalanceOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ClaimClaimableBalanceOp" {
		result = *u.ClaimClaimableBalanceOp
		ok = true
	}

	return
}

// MustBeginSponsoringFutureReservesOp retrieves the BeginSponsoringFutureReservesOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustBeginSponsoringFutureReservesOp() BeginSponsoringFutureReservesOp {
	val, ok := u.GetBeginSponsoringFutureReservesOp()

	if !ok {
		panic("arm BeginSponsoringFutureReservesOp is not set")
	}

	return val
}

// GetBeginSponsoringFutureReservesOp retrieves the BeginSponsoringFutureReservesOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetBeginSponsoringFutureReservesOp() (result BeginSponsoringFutureReservesOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "BeginSponsoringFutureReservesOp" {
		result = *u.BeginSponsoringFutureReservesOp
		ok = true
	}

	return
}

// MustRevokeSponsorshipOp retrieves the RevokeSponsorshipOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustRevokeSponsorshipOp() RevokeSponsorshipOp {
	val, ok := u.GetRevokeSponsorshipOp()

	if !ok {
		panic("arm RevokeSponsorshipOp is not set")
	}

	return val
}

// GetRevokeSponsorshipOp retrieves the RevokeSponsorshipOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetRevokeSponsorshipOp() (result RevokeSponsorshipOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RevokeSponsorshipOp" {
		result = *u.RevokeSponsorshipOp
		ok = true
	}

	return
}

// MustClawbackOp retrieves the ClawbackOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustClawbackOp() ClawbackOp {
	val, ok := u.GetClawbackOp()

	if !ok {
		panic("arm ClawbackOp is not set")
	}

	return val
}

// GetClawbackOp retrieves the ClawbackOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetClawbackOp() (result ClawbackOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ClawbackOp" {
		result = *u.ClawbackOp
		ok = true
	}

	return
}

// MustClawbackClaimableBalanceOp retrieves the ClawbackClaimableBalanceOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustClawbackClaimableBalanceOp() ClawbackClaimableBalanceOp {
	val, ok := u.GetClawbackClaimableBalanceOp()

	if !ok {
		panic("arm ClawbackClaimableBalanceOp is not set")
	}

	return val
}

// GetClawbackClaimableBalanceOp retrieves the ClawbackClaimableBalanceOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetClawbackClaimableBalanceOp() (result ClawbackClaimableBalanceOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ClawbackClaimableBalanceOp" {
		result = *u.ClawbackClaimableBalanceOp
		ok = true
	}

	return
}

// MustSetTrustLineFlagsOp retrieves the SetTrustLineFlagsOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustSetTrustLineFlagsOp() SetTrustLineFlagsOp {
	val, ok := u.GetSetTrustLineFlagsOp()

	if !ok {
		panic("arm SetTrustLineFlagsOp is not set")
	}

	return val
}

// GetSetTrustLineFlagsOp retrieves the SetTrustLineFlagsOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetSetTrustLineFlagsOp() (result SetTrustLineFlagsOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SetTrustLineFlagsOp" {
		result = *u.SetTrustLineFlagsOp
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s OperationBody) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *OperationBody) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*OperationBody)(nil)
	_ encoding.BinaryUnmarshaler = (*OperationBody)(nil)
)

// Operation is an XDR Struct defines as:
//
//   struct Operation
//    {
//        // sourceAccount is the account used to run the operation
//        // if not set, the runtime defaults to "sourceAccount" specified at
//        // the transaction level
//        MuxedAccount* sourceAccount;
//
//        union switch (OperationType type)
//        {
//        case CREATE_ACCOUNT:
//            CreateAccountOp createAccountOp;
//        case PAYMENT:
//            PaymentOp paymentOp;
//        case PATH_PAYMENT_STRICT_RECEIVE:
//            PathPaymentStrictReceiveOp pathPaymentStrictReceiveOp;
//        case MANAGE_SELL_OFFER:
//            ManageSellOfferOp manageSellOfferOp;
//        case CREATE_PASSIVE_SELL_OFFER:
//            CreatePassiveSellOfferOp createPassiveSellOfferOp;
//        case SET_OPTIONS:
//            SetOptionsOp setOptionsOp;
//        case CHANGE_TRUST:
//            ChangeTrustOp changeTrustOp;
//        case ALLOW_TRUST:
//            AllowTrustOp allowTrustOp;
//        case ACCOUNT_MERGE:
//            MuxedAccount destination;
//        case INFLATION:
//            void;
//        case MANAGE_DATA:
//            ManageDataOp manageDataOp;
//        case BUMP_SEQUENCE:
//            BumpSequenceOp bumpSequenceOp;
//        case MANAGE_BUY_OFFER:
//            ManageBuyOfferOp manageBuyOfferOp;
//        case PATH_PAYMENT_STRICT_SEND:
//            PathPaymentStrictSendOp pathPaymentStrictSendOp;
//        case CREATE_CLAIMABLE_BALANCE:
//            CreateClaimableBalanceOp createClaimableBalanceOp;
//        case CLAIM_CLAIMABLE_BALANCE:
//            ClaimClaimableBalanceOp claimClaimableBalanceOp;
//        case BEGIN_SPONSORING_FUTURE_RESERVES:
//            BeginSponsoringFutureReservesOp beginSponsoringFutureReservesOp;
//        case END_SPONSORING_FUTURE_RESERVES:
//            void;
//        case REVOKE_SPONSORSHIP:
//            RevokeSponsorshipOp revokeSponsorshipOp;
//        case CLAWBACK:
//            ClawbackOp clawbackOp;
//        case CLAWBACK_CLAIMABLE_BALANCE:
//            ClawbackClaimableBalanceOp clawbackClaimableBalanceOp;
//        case SET_TRUST_LINE_FLAGS:
//            SetTrustLineFlagsOp setTrustLineFlagsOp;
//        }
//        body;
//    };
//
type Operation struct {
	SourceAccount *MuxedAccount
	Body          OperationBody
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Operation) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Operation) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Operation)(nil)
	_ encoding.BinaryUnmarshaler = (*Operation)(nil)
)

// OperationIdId is an XDR NestedStruct defines as:
//
//   struct
//        {
//            AccountID sourceAccount;
//            SequenceNumber seqNum;
//            uint32 opNum;
//        }
//
type OperationIdId struct {
	SourceAccount AccountId
	SeqNum        SequenceNumber
	OpNum         Uint32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s OperationIdId) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *OperationIdId) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*OperationIdId)(nil)
	_ encoding.BinaryUnmarshaler = (*OperationIdId)(nil)
)

// OperationId is an XDR Union defines as:
//
//   union OperationID switch (EnvelopeType type)
//    {
//    case ENVELOPE_TYPE_OP_ID:
//        struct
//        {
//            AccountID sourceAccount;
//            SequenceNumber seqNum;
//            uint32 opNum;
//        } id;
//    };
//
type OperationId struct {
	Type EnvelopeType
	Id   *OperationIdId
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OperationId) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OperationId
func (u OperationId) ArmForSwitch(sw int32) (string, bool) {
	switch EnvelopeType(sw) {
	case EnvelopeTypeEnvelopeTypeOpId:
		return "Id", true
	}
	return "-", false
}

// NewOperationId creates a new  OperationId.
func NewOperationId(aType EnvelopeType, value interface{}) (result OperationId, err error) {
	result.Type = aType
	switch EnvelopeType(aType) {
	case EnvelopeTypeEnvelopeTypeOpId:
		tv, ok := value.(OperationIdId)
		if !ok {
			err = fmt.Errorf("invalid value, must be OperationIdId")
			return
		}
		result.Id = &tv
	}
	return
}

// MustId retrieves the Id value from the union,
// panicing if the value is not set.
func (u OperationId) MustId() OperationIdId {
	val, ok := u.GetId()

	if !ok {
		panic("arm Id is not set")
	}

	return val
}

// GetId retrieves the Id value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationId) GetId() (result OperationIdId, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Id" {
		result = *u.Id
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s OperationId) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *OperationId) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*OperationId)(nil)
	_ encoding.BinaryUnmarshaler = (*OperationId)(nil)
)

// MemoType is an XDR Enum defines as:
//
//   enum MemoType
//    {
//        MEMO_NONE = 0,
//        MEMO_TEXT = 1,
//        MEMO_ID = 2,
//        MEMO_HASH = 3,
//        MEMO_RETURN = 4
//    };
//
type MemoType int32

const (
	MemoTypeMemoNone   MemoType = 0
	MemoTypeMemoText   MemoType = 1
	MemoTypeMemoId     MemoType = 2
	MemoTypeMemoHash   MemoType = 3
	MemoTypeMemoReturn MemoType = 4
)

var memoTypeMap = map[int32]string{
	0: "MemoTypeMemoNone",
	1: "MemoTypeMemoText",
	2: "MemoTypeMemoId",
	3: "MemoTypeMemoHash",
	4: "MemoTypeMemoReturn",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for MemoType
func (e MemoType) ValidEnum(v int32) bool {
	_, ok := memoTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e MemoType) String() string {
	name, _ := memoTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s MemoType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *MemoType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*MemoType)(nil)
	_ encoding.BinaryUnmarshaler = (*MemoType)(nil)
)

// Memo is an XDR Union defines as:
//
//   union Memo switch (MemoType type)
//    {
//    case MEMO_NONE:
//        void;
//    case MEMO_TEXT:
//        string text<28>;
//    case MEMO_ID:
//        uint64 id;
//    case MEMO_HASH:
//        Hash hash; // the hash of what to pull from the content server
//    case MEMO_RETURN:
//        Hash retHash; // the hash of the tx you are rejecting
//    };
//
type Memo struct {
	Type    MemoType
	Text    *string `xdrmaxsize:"28"`
	Id      *Uint64
	Hash    *Hash
	RetHash *Hash
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u Memo) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of Memo
func (u Memo) ArmForSwitch(sw int32) (string, bool) {
	switch MemoType(sw) {
	case MemoTypeMemoNone:
		return "", true
	case MemoTypeMemoText:
		return "Text", true
	case MemoTypeMemoId:
		return "Id", true
	case MemoTypeMemoHash:
		return "Hash", true
	case MemoTypeMemoReturn:
		return "RetHash", true
	}
	return "-", false
}

// NewMemo creates a new  Memo.
func NewMemo(aType MemoType, value interface{}) (result Memo, err error) {
	result.Type = aType
	switch MemoType(aType) {
	case MemoTypeMemoNone:
		// void
	case MemoTypeMemoText:
		tv, ok := value.(string)
		if !ok {
			err = fmt.Errorf("invalid value, must be string")
			return
		}
		result.Text = &tv
	case MemoTypeMemoId:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.Id = &tv
	case MemoTypeMemoHash:
		tv, ok := value.(Hash)
		if !ok {
			err = fmt.Errorf("invalid value, must be Hash")
			return
		}
		result.Hash = &tv
	case MemoTypeMemoReturn:
		tv, ok := value.(Hash)
		if !ok {
			err = fmt.Errorf("invalid value, must be Hash")
			return
		}
		result.RetHash = &tv
	}
	return
}

// MustText retrieves the Text value from the union,
// panicing if the value is not set.
func (u Memo) MustText() string {
	val, ok := u.GetText()

	if !ok {
		panic("arm Text is not set")
	}

	return val
}

// GetText retrieves the Text value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetText() (result string, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Text" {
		result = *u.Text
		ok = true
	}

	return
}

// MustId retrieves the Id value from the union,
// panicing if the value is not set.
func (u Memo) MustId() Uint64 {
	val, ok := u.GetId()

	if !ok {
		panic("arm Id is not set")
	}

	return val
}

// GetId retrieves the Id value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Id" {
		result = *u.Id
		ok = true
	}

	return
}

// MustHash retrieves the Hash value from the union,
// panicing if the value is not set.
func (u Memo) MustHash() Hash {
	val, ok := u.GetHash()

	if !ok {
		panic("arm Hash is not set")
	}

	return val
}

// GetHash retrieves the Hash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetHash() (result Hash, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Hash" {
		result = *u.Hash
		ok = true
	}

	return
}

// MustRetHash retrieves the RetHash value from the union,
// panicing if the value is not set.
func (u Memo) MustRetHash() Hash {
	val, ok := u.GetRetHash()

	if !ok {
		panic("arm RetHash is not set")
	}

	return val
}

// GetRetHash retrieves the RetHash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetRetHash() (result Hash, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RetHash" {
		result = *u.RetHash
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Memo) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Memo) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Memo)(nil)
	_ encoding.BinaryUnmarshaler = (*Memo)(nil)
)

// TimeBounds is an XDR Struct defines as:
//
//   struct TimeBounds
//    {
//        TimePoint minTime;
//        TimePoint maxTime; // 0 here means no maxTime
//    };
//
type TimeBounds struct {
	MinTime TimePoint
	MaxTime TimePoint
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TimeBounds) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TimeBounds) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TimeBounds)(nil)
	_ encoding.BinaryUnmarshaler = (*TimeBounds)(nil)
)

// MaxOpsPerTx is an XDR Const defines as:
//
//   const MAX_OPS_PER_TX = 100;
//
const MaxOpsPerTx = 100

// TransactionV0Ext is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        }
//
type TransactionV0Ext struct {
	V int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionV0Ext) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionV0Ext
func (u TransactionV0Ext) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	}
	return "-", false
}

// NewTransactionV0Ext creates a new  TransactionV0Ext.
func NewTransactionV0Ext(v int32, value interface{}) (result TransactionV0Ext, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionV0Ext) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionV0Ext) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionV0Ext)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionV0Ext)(nil)
)

// TransactionV0 is an XDR Struct defines as:
//
//   struct TransactionV0
//    {
//        uint256 sourceAccountEd25519;
//        uint32 fee;
//        SequenceNumber seqNum;
//        TimeBounds* timeBounds;
//        Memo memo;
//        Operation operations<MAX_OPS_PER_TX>;
//        union switch (int v)
//        {
//        case 0:
//            void;
//        }
//        ext;
//    };
//
type TransactionV0 struct {
	SourceAccountEd25519 Uint256
	Fee                  Uint32
	SeqNum               SequenceNumber
	TimeBounds           *TimeBounds
	Memo                 Memo
	Operations           []Operation `xdrmaxsize:"100"`
	Ext                  TransactionV0Ext
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionV0) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionV0) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionV0)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionV0)(nil)
)

// TransactionV0Envelope is an XDR Struct defines as:
//
//   struct TransactionV0Envelope
//    {
//        TransactionV0 tx;
//        /* Each decorated signature is a signature over the SHA256 hash of
//         * a TransactionSignaturePayload */
//        DecoratedSignature signatures<20>;
//    };
//
type TransactionV0Envelope struct {
	Tx         TransactionV0
	Signatures []DecoratedSignature `xdrmaxsize:"20"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionV0Envelope) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionV0Envelope) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionV0Envelope)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionV0Envelope)(nil)
)

// TransactionExt is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        }
//
type TransactionExt struct {
	V int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionExt
func (u TransactionExt) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	}
	return "-", false
}

// NewTransactionExt creates a new  TransactionExt.
func NewTransactionExt(v int32, value interface{}) (result TransactionExt, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionExt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionExt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionExt)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionExt)(nil)
)

// Transaction is an XDR Struct defines as:
//
//   struct Transaction
//    {
//        // account used to run the transaction
//        MuxedAccount sourceAccount;
//
//        // the fee the sourceAccount will pay
//        uint32 fee;
//
//        // sequence number to consume in the account
//        SequenceNumber seqNum;
//
//        // validity range (inclusive) for the last ledger close time
//        TimeBounds* timeBounds;
//
//        Memo memo;
//
//        Operation operations<MAX_OPS_PER_TX>;
//
//        // reserved for future use
//        union switch (int v)
//        {
//        case 0:
//            void;
//        }
//        ext;
//    };
//
type Transaction struct {
	SourceAccount MuxedAccount
	Fee           Uint32
	SeqNum        SequenceNumber
	TimeBounds    *TimeBounds
	Memo          Memo
	Operations    []Operation `xdrmaxsize:"100"`
	Ext           TransactionExt
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Transaction) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Transaction) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Transaction)(nil)
	_ encoding.BinaryUnmarshaler = (*Transaction)(nil)
)

// TransactionV1Envelope is an XDR Struct defines as:
//
//   struct TransactionV1Envelope
//    {
//        Transaction tx;
//        /* Each decorated signature is a signature over the SHA256 hash of
//         * a TransactionSignaturePayload */
//        DecoratedSignature signatures<20>;
//    };
//
type TransactionV1Envelope struct {
	Tx         Transaction
	Signatures []DecoratedSignature `xdrmaxsize:"20"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionV1Envelope) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionV1Envelope) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionV1Envelope)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionV1Envelope)(nil)
)

// FeeBumpTransactionInnerTx is an XDR NestedUnion defines as:
//
//   union switch (EnvelopeType type)
//        {
//        case ENVELOPE_TYPE_TX:
//            TransactionV1Envelope v1;
//        }
//
type FeeBumpTransactionInnerTx struct {
	Type EnvelopeType
	V1   *TransactionV1Envelope
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u FeeBumpTransactionInnerTx) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of FeeBumpTransactionInnerTx
func (u FeeBumpTransactionInnerTx) ArmForSwitch(sw int32) (string, bool) {
	switch EnvelopeType(sw) {
	case EnvelopeTypeEnvelopeTypeTx:
		return "V1", true
	}
	return "-", false
}

// NewFeeBumpTransactionInnerTx creates a new  FeeBumpTransactionInnerTx.
func NewFeeBumpTransactionInnerTx(aType EnvelopeType, value interface{}) (result FeeBumpTransactionInnerTx, err error) {
	result.Type = aType
	switch EnvelopeType(aType) {
	case EnvelopeTypeEnvelopeTypeTx:
		tv, ok := value.(TransactionV1Envelope)
		if !ok {
			err = fmt.Errorf("invalid value, must be TransactionV1Envelope")
			return
		}
		result.V1 = &tv
	}
	return
}

// MustV1 retrieves the V1 value from the union,
// panicing if the value is not set.
func (u FeeBumpTransactionInnerTx) MustV1() TransactionV1Envelope {
	val, ok := u.GetV1()

	if !ok {
		panic("arm V1 is not set")
	}

	return val
}

// GetV1 retrieves the V1 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u FeeBumpTransactionInnerTx) GetV1() (result TransactionV1Envelope, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "V1" {
		result = *u.V1
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s FeeBumpTransactionInnerTx) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *FeeBumpTransactionInnerTx) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*FeeBumpTransactionInnerTx)(nil)
	_ encoding.BinaryUnmarshaler = (*FeeBumpTransactionInnerTx)(nil)
)

// FeeBumpTransactionExt is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        }
//
type FeeBumpTransactionExt struct {
	V int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u FeeBumpTransactionExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of FeeBumpTransactionExt
func (u FeeBumpTransactionExt) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	}
	return "-", false
}

// NewFeeBumpTransactionExt creates a new  FeeBumpTransactionExt.
func NewFeeBumpTransactionExt(v int32, value interface{}) (result FeeBumpTransactionExt, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s FeeBumpTransactionExt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *FeeBumpTransactionExt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*FeeBumpTransactionExt)(nil)
	_ encoding.BinaryUnmarshaler = (*FeeBumpTransactionExt)(nil)
)

// FeeBumpTransaction is an XDR Struct defines as:
//
//   struct FeeBumpTransaction
//    {
//        MuxedAccount feeSource;
//        int64 fee;
//        union switch (EnvelopeType type)
//        {
//        case ENVELOPE_TYPE_TX:
//            TransactionV1Envelope v1;
//        }
//        innerTx;
//        union switch (int v)
//        {
//        case 0:
//            void;
//        }
//        ext;
//    };
//
type FeeBumpTransaction struct {
	FeeSource MuxedAccount
	Fee       Int64
	InnerTx   FeeBumpTransactionInnerTx
	Ext       FeeBumpTransactionExt
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s FeeBumpTransaction) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *FeeBumpTransaction) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*FeeBumpTransaction)(nil)
	_ encoding.BinaryUnmarshaler = (*FeeBumpTransaction)(nil)
)

// FeeBumpTransactionEnvelope is an XDR Struct defines as:
//
//   struct FeeBumpTransactionEnvelope
//    {
//        FeeBumpTransaction tx;
//        /* Each decorated signature is a signature over the SHA256 hash of
//         * a TransactionSignaturePayload */
//        DecoratedSignature signatures<20>;
//    };
//
type FeeBumpTransactionEnvelope struct {
	Tx         FeeBumpTransaction
	Signatures []DecoratedSignature `xdrmaxsize:"20"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s FeeBumpTransactionEnvelope) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *FeeBumpTransactionEnvelope) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*FeeBumpTransactionEnvelope)(nil)
	_ encoding.BinaryUnmarshaler = (*FeeBumpTransactionEnvelope)(nil)
)

// TransactionEnvelope is an XDR Union defines as:
//
//   union TransactionEnvelope switch (EnvelopeType type)
//    {
//    case ENVELOPE_TYPE_TX_V0:
//        TransactionV0Envelope v0;
//    case ENVELOPE_TYPE_TX:
//        TransactionV1Envelope v1;
//    case ENVELOPE_TYPE_TX_FEE_BUMP:
//        FeeBumpTransactionEnvelope feeBump;
//    };
//
type TransactionEnvelope struct {
	Type    EnvelopeType
	V0      *TransactionV0Envelope
	V1      *TransactionV1Envelope
	FeeBump *FeeBumpTransactionEnvelope
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionEnvelope) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionEnvelope
func (u TransactionEnvelope) ArmForSwitch(sw int32) (string, bool) {
	switch EnvelopeType(sw) {
	case EnvelopeTypeEnvelopeTypeTxV0:
		return "V0", true
	case EnvelopeTypeEnvelopeTypeTx:
		return "V1", true
	case EnvelopeTypeEnvelopeTypeTxFeeBump:
		return "FeeBump", true
	}
	return "-", false
}

// NewTransactionEnvelope creates a new  TransactionEnvelope.
func NewTransactionEnvelope(aType EnvelopeType, value interface{}) (result TransactionEnvelope, err error) {
	result.Type = aType
	switch EnvelopeType(aType) {
	case EnvelopeTypeEnvelopeTypeTxV0:
		tv, ok := value.(TransactionV0Envelope)
		if !ok {
			err = fmt.Errorf("invalid value, must be TransactionV0Envelope")
			return
		}
		result.V0 = &tv
	case EnvelopeTypeEnvelopeTypeTx:
		tv, ok := value.(TransactionV1Envelope)
		if !ok {
			err = fmt.Errorf("invalid value, must be TransactionV1Envelope")
			return
		}
		result.V1 = &tv
	case EnvelopeTypeEnvelopeTypeTxFeeBump:
		tv, ok := value.(FeeBumpTransactionEnvelope)
		if !ok {
			err = fmt.Errorf("invalid value, must be FeeBumpTransactionEnvelope")
			return
		}
		result.FeeBump = &tv
	}
	return
}

// MustV0 retrieves the V0 value from the union,
// panicing if the value is not set.
func (u TransactionEnvelope) MustV0() TransactionV0Envelope {
	val, ok := u.GetV0()

	if !ok {
		panic("arm V0 is not set")
	}

	return val
}

// GetV0 retrieves the V0 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionEnvelope) GetV0() (result TransactionV0Envelope, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "V0" {
		result = *u.V0
		ok = true
	}

	return
}

// MustV1 retrieves the V1 value from the union,
// panicing if the value is not set.
func (u TransactionEnvelope) MustV1() TransactionV1Envelope {
	val, ok := u.GetV1()

	if !ok {
		panic("arm V1 is not set")
	}

	return val
}

// GetV1 retrieves the V1 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionEnvelope) GetV1() (result TransactionV1Envelope, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "V1" {
		result = *u.V1
		ok = true
	}

	return
}

// MustFeeBump retrieves the FeeBump value from the union,
// panicing if the value is not set.
func (u TransactionEnvelope) MustFeeBump() FeeBumpTransactionEnvelope {
	val, ok := u.GetFeeBump()

	if !ok {
		panic("arm FeeBump is not set")
	}

	return val
}

// GetFeeBump retrieves the FeeBump value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionEnvelope) GetFeeBump() (result FeeBumpTransactionEnvelope, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "FeeBump" {
		result = *u.FeeBump
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionEnvelope) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionEnvelope) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionEnvelope)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionEnvelope)(nil)
)

// TransactionSignaturePayloadTaggedTransaction is an XDR NestedUnion defines as:
//
//   union switch (EnvelopeType type)
//        {
//        // Backwards Compatibility: Use ENVELOPE_TYPE_TX to sign ENVELOPE_TYPE_TX_V0
//        case ENVELOPE_TYPE_TX:
//            Transaction tx;
//        case ENVELOPE_TYPE_TX_FEE_BUMP:
//            FeeBumpTransaction feeBump;
//        }
//
type TransactionSignaturePayloadTaggedTransaction struct {
	Type    EnvelopeType
	Tx      *Transaction
	FeeBump *FeeBumpTransaction
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionSignaturePayloadTaggedTransaction) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionSignaturePayloadTaggedTransaction
func (u TransactionSignaturePayloadTaggedTransaction) ArmForSwitch(sw int32) (string, bool) {
	switch EnvelopeType(sw) {
	case EnvelopeTypeEnvelopeTypeTx:
		return "Tx", true
	case EnvelopeTypeEnvelopeTypeTxFeeBump:
		return "FeeBump", true
	}
	return "-", false
}

// NewTransactionSignaturePayloadTaggedTransaction creates a new  TransactionSignaturePayloadTaggedTransaction.
func NewTransactionSignaturePayloadTaggedTransaction(aType EnvelopeType, value interface{}) (result TransactionSignaturePayloadTaggedTransaction, err error) {
	result.Type = aType
	switch EnvelopeType(aType) {
	case EnvelopeTypeEnvelopeTypeTx:
		tv, ok := value.(Transaction)
		if !ok {
			err = fmt.Errorf("invalid value, must be Transaction")
			return
		}
		result.Tx = &tv
	case EnvelopeTypeEnvelopeTypeTxFeeBump:
		tv, ok := value.(FeeBumpTransaction)
		if !ok {
			err = fmt.Errorf("invalid value, must be FeeBumpTransaction")
			return
		}
		result.FeeBump = &tv
	}
	return
}

// MustTx retrieves the Tx value from the union,
// panicing if the value is not set.
func (u TransactionSignaturePayloadTaggedTransaction) MustTx() Transaction {
	val, ok := u.GetTx()

	if !ok {
		panic("arm Tx is not set")
	}

	return val
}

// GetTx retrieves the Tx value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionSignaturePayloadTaggedTransaction) GetTx() (result Transaction, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Tx" {
		result = *u.Tx
		ok = true
	}

	return
}

// MustFeeBump retrieves the FeeBump value from the union,
// panicing if the value is not set.
func (u TransactionSignaturePayloadTaggedTransaction) MustFeeBump() FeeBumpTransaction {
	val, ok := u.GetFeeBump()

	if !ok {
		panic("arm FeeBump is not set")
	}

	return val
}

// GetFeeBump retrieves the FeeBump value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionSignaturePayloadTaggedTransaction) GetFeeBump() (result FeeBumpTransaction, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "FeeBump" {
		result = *u.FeeBump
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionSignaturePayloadTaggedTransaction) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionSignaturePayloadTaggedTransaction) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionSignaturePayloadTaggedTransaction)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionSignaturePayloadTaggedTransaction)(nil)
)

// TransactionSignaturePayload is an XDR Struct defines as:
//
//   struct TransactionSignaturePayload
//    {
//        Hash networkId;
//        union switch (EnvelopeType type)
//        {
//        // Backwards Compatibility: Use ENVELOPE_TYPE_TX to sign ENVELOPE_TYPE_TX_V0
//        case ENVELOPE_TYPE_TX:
//            Transaction tx;
//        case ENVELOPE_TYPE_TX_FEE_BUMP:
//            FeeBumpTransaction feeBump;
//        }
//        taggedTransaction;
//    };
//
type TransactionSignaturePayload struct {
	NetworkId         Hash
	TaggedTransaction TransactionSignaturePayloadTaggedTransaction
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionSignaturePayload) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionSignaturePayload) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionSignaturePayload)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionSignaturePayload)(nil)
)

// ClaimAtomType is an XDR Enum defines as:
//
//   enum ClaimAtomType
//    {
//        CLAIM_ATOM_TYPE_V0 = 0,
//        CLAIM_ATOM_TYPE_ORDER_BOOK = 1
//    };
//
type ClaimAtomType int32

const (
	ClaimAtomTypeClaimAtomTypeV0        ClaimAtomType = 0
	ClaimAtomTypeClaimAtomTypeOrderBook ClaimAtomType = 1
)

var claimAtomTypeMap = map[int32]string{
	0: "ClaimAtomTypeClaimAtomTypeV0",
	1: "ClaimAtomTypeClaimAtomTypeOrderBook",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ClaimAtomType
func (e ClaimAtomType) ValidEnum(v int32) bool {
	_, ok := claimAtomTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e ClaimAtomType) String() string {
	name, _ := claimAtomTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimAtomType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimAtomType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimAtomType)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimAtomType)(nil)
)

// ClaimOfferAtomV0 is an XDR Struct defines as:
//
//   struct ClaimOfferAtomV0
//    {
//        // emitted to identify the offer
//        uint256 sellerEd25519; // Account that owns the offer
//        int64 offerID;
//
//        // amount and asset taken from the owner
//        Asset assetSold;
//        int64 amountSold;
//
//        // amount and asset sent to the owner
//        Asset assetBought;
//        int64 amountBought;
//    };
//
type ClaimOfferAtomV0 struct {
	SellerEd25519 Uint256
	OfferId       Int64
	AssetSold     Asset
	AmountSold    Int64
	AssetBought   Asset
	AmountBought  Int64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimOfferAtomV0) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimOfferAtomV0) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimOfferAtomV0)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimOfferAtomV0)(nil)
)

// ClaimOfferAtom is an XDR Struct defines as:
//
//   struct ClaimOfferAtom
//    {
//        // emitted to identify the offer
//        AccountID sellerID; // Account that owns the offer
//        int64 offerID;
//
//        // amount and asset taken from the owner
//        Asset assetSold;
//        int64 amountSold;
//
//        // amount and asset sent to the owner
//        Asset assetBought;
//        int64 amountBought;
//    };
//
type ClaimOfferAtom struct {
	SellerId     AccountId
	OfferId      Int64
	AssetSold    Asset
	AmountSold   Int64
	AssetBought  Asset
	AmountBought Int64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimOfferAtom) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimOfferAtom) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimOfferAtom)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimOfferAtom)(nil)
)

// ClaimAtom is an XDR Union defines as:
//
//   union ClaimAtom switch (ClaimAtomType type)
//    {
//    case CLAIM_ATOM_TYPE_V0:
//        ClaimOfferAtomV0 v0;
//    case CLAIM_ATOM_TYPE_ORDER_BOOK:
//        ClaimOfferAtom orderBook;
//    };
//
type ClaimAtom struct {
	Type      ClaimAtomType
	V0        *ClaimOfferAtomV0
	OrderBook *ClaimOfferAtom
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ClaimAtom) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ClaimAtom
func (u ClaimAtom) ArmForSwitch(sw int32) (string, bool) {
	switch ClaimAtomType(sw) {
	case ClaimAtomTypeClaimAtomTypeV0:
		return "V0", true
	case ClaimAtomTypeClaimAtomTypeOrderBook:
		return "OrderBook", true
	}
	return "-", false
}

// NewClaimAtom creates a new  ClaimAtom.
func NewClaimAtom(aType ClaimAtomType, value interface{}) (result ClaimAtom, err error) {
	result.Type = aType
	switch ClaimAtomType(aType) {
	case ClaimAtomTypeClaimAtomTypeV0:
		tv, ok := value.(ClaimOfferAtomV0)
		if !ok {
			err = fmt.Errorf("invalid value, must be ClaimOfferAtomV0")
			return
		}
		result.V0 = &tv
	case ClaimAtomTypeClaimAtomTypeOrderBook:
		tv, ok := value.(ClaimOfferAtom)
		if !ok {
			err = fmt.Errorf("invalid value, must be ClaimOfferAtom")
			return
		}
		result.OrderBook = &tv
	}
	return
}

// MustV0 retrieves the V0 value from the union,
// panicing if the value is not set.
func (u ClaimAtom) MustV0() ClaimOfferAtomV0 {
	val, ok := u.GetV0()

	if !ok {
		panic("arm V0 is not set")
	}

	return val
}

// GetV0 retrieves the V0 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ClaimAtom) GetV0() (result ClaimOfferAtomV0, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "V0" {
		result = *u.V0
		ok = true
	}

	return
}

// MustOrderBook retrieves the OrderBook value from the union,
// panicing if the value is not set.
func (u ClaimAtom) MustOrderBook() ClaimOfferAtom {
	val, ok := u.GetOrderBook()

	if !ok {
		panic("arm OrderBook is not set")
	}

	return val
}

// GetOrderBook retrieves the OrderBook value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ClaimAtom) GetOrderBook() (result ClaimOfferAtom, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "OrderBook" {
		result = *u.OrderBook
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimAtom) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimAtom) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimAtom)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimAtom)(nil)
)

// CreateAccountResultCode is an XDR Enum defines as:
//
//   enum CreateAccountResultCode
//    {
//        // codes considered as "success" for the operation
//        CREATE_ACCOUNT_SUCCESS = 0, // account was created
//
//        // codes considered as "failure" for the operation
//        CREATE_ACCOUNT_MALFORMED = -1,   // invalid destination
//        CREATE_ACCOUNT_UNDERFUNDED = -2, // not enough funds in source account
//        CREATE_ACCOUNT_LOW_RESERVE =
//            -3, // would create an account below the min reserve
//        CREATE_ACCOUNT_ALREADY_EXIST = -4 // account already exists
//    };
//
type CreateAccountResultCode int32

const (
	CreateAccountResultCodeCreateAccountSuccess      CreateAccountResultCode = 0
	CreateAccountResultCodeCreateAccountMalformed    CreateAccountResultCode = -1
	CreateAccountResultCodeCreateAccountUnderfunded  CreateAccountResultCode = -2
	CreateAccountResultCodeCreateAccountLowReserve   CreateAccountResultCode = -3
	CreateAccountResultCodeCreateAccountAlreadyExist CreateAccountResultCode = -4
)

var createAccountResultCodeMap = map[int32]string{
	0:  "CreateAccountResultCodeCreateAccountSuccess",
	-1: "CreateAccountResultCodeCreateAccountMalformed",
	-2: "CreateAccountResultCodeCreateAccountUnderfunded",
	-3: "CreateAccountResultCodeCreateAccountLowReserve",
	-4: "CreateAccountResultCodeCreateAccountAlreadyExist",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateAccountResultCode
func (e CreateAccountResultCode) ValidEnum(v int32) bool {
	_, ok := createAccountResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e CreateAccountResultCode) String() string {
	name, _ := createAccountResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s CreateAccountResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *CreateAccountResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*CreateAccountResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*CreateAccountResultCode)(nil)
)

// CreateAccountResult is an XDR Union defines as:
//
//   union CreateAccountResult switch (CreateAccountResultCode code)
//    {
//    case CREATE_ACCOUNT_SUCCESS:
//        void;
//    default:
//        void;
//    };
//
type CreateAccountResult struct {
	Code CreateAccountResultCode
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAccountResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAccountResult
func (u CreateAccountResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateAccountResultCode(sw) {
	case CreateAccountResultCodeCreateAccountSuccess:
		return "", true
	default:
		return "", true
	}
}

// NewCreateAccountResult creates a new  CreateAccountResult.
func NewCreateAccountResult(code CreateAccountResultCode, value interface{}) (result CreateAccountResult, err error) {
	result.Code = code
	switch CreateAccountResultCode(code) {
	case CreateAccountResultCodeCreateAccountSuccess:
		// void
	default:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s CreateAccountResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *CreateAccountResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*CreateAccountResult)(nil)
	_ encoding.BinaryUnmarshaler = (*CreateAccountResult)(nil)
)

// PaymentResultCode is an XDR Enum defines as:
//
//   enum PaymentResultCode
//    {
//        // codes considered as "success" for the operation
//        PAYMENT_SUCCESS = 0, // payment successfully completed
//
//        // codes considered as "failure" for the operation
//        PAYMENT_MALFORMED = -1,          // bad input
//        PAYMENT_UNDERFUNDED = -2,        // not enough funds in source account
//        PAYMENT_SRC_NO_TRUST = -3,       // no trust line on source account
//        PAYMENT_SRC_NOT_AUTHORIZED = -4, // source not authorized to transfer
//        PAYMENT_NO_DESTINATION = -5,     // destination account does not exist
//        PAYMENT_NO_TRUST = -6,       // destination missing a trust line for asset
//        PAYMENT_NOT_AUTHORIZED = -7, // destination not authorized to hold asset
//        PAYMENT_LINE_FULL = -8,      // destination would go above their limit
//        PAYMENT_NO_ISSUER = -9       // missing issuer on asset
//    };
//
type PaymentResultCode int32

const (
	PaymentResultCodePaymentSuccess          PaymentResultCode = 0
	PaymentResultCodePaymentMalformed        PaymentResultCode = -1
	PaymentResultCodePaymentUnderfunded      PaymentResultCode = -2
	PaymentResultCodePaymentSrcNoTrust       PaymentResultCode = -3
	PaymentResultCodePaymentSrcNotAuthorized PaymentResultCode = -4
	PaymentResultCodePaymentNoDestination    PaymentResultCode = -5
	PaymentResultCodePaymentNoTrust          PaymentResultCode = -6
	PaymentResultCodePaymentNotAuthorized    PaymentResultCode = -7
	PaymentResultCodePaymentLineFull         PaymentResultCode = -8
	PaymentResultCodePaymentNoIssuer         PaymentResultCode = -9
)

var paymentResultCodeMap = map[int32]string{
	0:  "PaymentResultCodePaymentSuccess",
	-1: "PaymentResultCodePaymentMalformed",
	-2: "PaymentResultCodePaymentUnderfunded",
	-3: "PaymentResultCodePaymentSrcNoTrust",
	-4: "PaymentResultCodePaymentSrcNotAuthorized",
	-5: "PaymentResultCodePaymentNoDestination",
	-6: "PaymentResultCodePaymentNoTrust",
	-7: "PaymentResultCodePaymentNotAuthorized",
	-8: "PaymentResultCodePaymentLineFull",
	-9: "PaymentResultCodePaymentNoIssuer",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PaymentResultCode
func (e PaymentResultCode) ValidEnum(v int32) bool {
	_, ok := paymentResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e PaymentResultCode) String() string {
	name, _ := paymentResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PaymentResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PaymentResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PaymentResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*PaymentResultCode)(nil)
)

// PaymentResult is an XDR Union defines as:
//
//   union PaymentResult switch (PaymentResultCode code)
//    {
//    case PAYMENT_SUCCESS:
//        void;
//    default:
//        void;
//    };
//
type PaymentResult struct {
	Code PaymentResultCode
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentResult
func (u PaymentResult) ArmForSwitch(sw int32) (string, bool) {
	switch PaymentResultCode(sw) {
	case PaymentResultCodePaymentSuccess:
		return "", true
	default:
		return "", true
	}
}

// NewPaymentResult creates a new  PaymentResult.
func NewPaymentResult(code PaymentResultCode, value interface{}) (result PaymentResult, err error) {
	result.Code = code
	switch PaymentResultCode(code) {
	case PaymentResultCodePaymentSuccess:
		// void
	default:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PaymentResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PaymentResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PaymentResult)(nil)
	_ encoding.BinaryUnmarshaler = (*PaymentResult)(nil)
)

// PathPaymentStrictReceiveResultCode is an XDR Enum defines as:
//
//   enum PathPaymentStrictReceiveResultCode
//    {
//        // codes considered as "success" for the operation
//        PATH_PAYMENT_STRICT_RECEIVE_SUCCESS = 0, // success
//
//        // codes considered as "failure" for the operation
//        PATH_PAYMENT_STRICT_RECEIVE_MALFORMED = -1, // bad input
//        PATH_PAYMENT_STRICT_RECEIVE_UNDERFUNDED =
//            -2, // not enough funds in source account
//        PATH_PAYMENT_STRICT_RECEIVE_SRC_NO_TRUST =
//            -3, // no trust line on source account
//        PATH_PAYMENT_STRICT_RECEIVE_SRC_NOT_AUTHORIZED =
//            -4, // source not authorized to transfer
//        PATH_PAYMENT_STRICT_RECEIVE_NO_DESTINATION =
//            -5, // destination account does not exist
//        PATH_PAYMENT_STRICT_RECEIVE_NO_TRUST =
//            -6, // dest missing a trust line for asset
//        PATH_PAYMENT_STRICT_RECEIVE_NOT_AUTHORIZED =
//            -7, // dest not authorized to hold asset
//        PATH_PAYMENT_STRICT_RECEIVE_LINE_FULL =
//            -8, // dest would go above their limit
//        PATH_PAYMENT_STRICT_RECEIVE_NO_ISSUER = -9, // missing issuer on one asset
//        PATH_PAYMENT_STRICT_RECEIVE_TOO_FEW_OFFERS =
//            -10, // not enough offers to satisfy path
//        PATH_PAYMENT_STRICT_RECEIVE_OFFER_CROSS_SELF =
//            -11, // would cross one of its own offers
//        PATH_PAYMENT_STRICT_RECEIVE_OVER_SENDMAX = -12 // could not satisfy sendmax
//    };
//
type PathPaymentStrictReceiveResultCode int32

const (
	PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveSuccess          PathPaymentStrictReceiveResultCode = 0
	PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveMalformed        PathPaymentStrictReceiveResultCode = -1
	PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveUnderfunded      PathPaymentStrictReceiveResultCode = -2
	PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveSrcNoTrust       PathPaymentStrictReceiveResultCode = -3
	PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveSrcNotAuthorized PathPaymentStrictReceiveResultCode = -4
	PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveNoDestination    PathPaymentStrictReceiveResultCode = -5
	PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveNoTrust          PathPaymentStrictReceiveResultCode = -6
	PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveNotAuthorized    PathPaymentStrictReceiveResultCode = -7
	PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveLineFull         PathPaymentStrictReceiveResultCode = -8
	PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveNoIssuer         PathPaymentStrictReceiveResultCode = -9
	PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveTooFewOffers     PathPaymentStrictReceiveResultCode = -10
	PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveOfferCrossSelf   PathPaymentStrictReceiveResultCode = -11
	PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveOverSendmax      PathPaymentStrictReceiveResultCode = -12
)

var pathPaymentStrictReceiveResultCodeMap = map[int32]string{
	0:   "PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveSuccess",
	-1:  "PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveMalformed",
	-2:  "PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveUnderfunded",
	-3:  "PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveSrcNoTrust",
	-4:  "PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveSrcNotAuthorized",
	-5:  "PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveNoDestination",
	-6:  "PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveNoTrust",
	-7:  "PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveNotAuthorized",
	-8:  "PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveLineFull",
	-9:  "PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveNoIssuer",
	-10: "PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveTooFewOffers",
	-11: "PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveOfferCrossSelf",
	-12: "PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveOverSendmax",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PathPaymentStrictReceiveResultCode
func (e PathPaymentStrictReceiveResultCode) ValidEnum(v int32) bool {
	_, ok := pathPaymentStrictReceiveResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e PathPaymentStrictReceiveResultCode) String() string {
	name, _ := pathPaymentStrictReceiveResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PathPaymentStrictReceiveResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PathPaymentStrictReceiveResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PathPaymentStrictReceiveResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*PathPaymentStrictReceiveResultCode)(nil)
)

// SimplePaymentResult is an XDR Struct defines as:
//
//   struct SimplePaymentResult
//    {
//        AccountID destination;
//        Asset asset;
//        int64 amount;
//    };
//
type SimplePaymentResult struct {
	Destination AccountId
	Asset       Asset
	Amount      Int64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SimplePaymentResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SimplePaymentResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SimplePaymentResult)(nil)
	_ encoding.BinaryUnmarshaler = (*SimplePaymentResult)(nil)
)

// PathPaymentStrictReceiveResultSuccess is an XDR NestedStruct defines as:
//
//   struct
//        {
//            ClaimAtom offers<>;
//            SimplePaymentResult last;
//        }
//
type PathPaymentStrictReceiveResultSuccess struct {
	Offers []ClaimAtom
	Last   SimplePaymentResult
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PathPaymentStrictReceiveResultSuccess) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PathPaymentStrictReceiveResultSuccess) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PathPaymentStrictReceiveResultSuccess)(nil)
	_ encoding.BinaryUnmarshaler = (*PathPaymentStrictReceiveResultSuccess)(nil)
)

// PathPaymentStrictReceiveResult is an XDR Union defines as:
//
//   union PathPaymentStrictReceiveResult switch (PathPaymentStrictReceiveResultCode code)
//    {
//    case PATH_PAYMENT_STRICT_RECEIVE_SUCCESS:
//        struct
//        {
//            ClaimAtom offers<>;
//            SimplePaymentResult last;
//        } success;
//    case PATH_PAYMENT_STRICT_RECEIVE_NO_ISSUER:
//        Asset noIssuer; // the asset that caused the error
//    default:
//        void;
//    };
//
type PathPaymentStrictReceiveResult struct {
	Code     PathPaymentStrictReceiveResultCode
	Success  *PathPaymentStrictReceiveResultSuccess
	NoIssuer *Asset
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PathPaymentStrictReceiveResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PathPaymentStrictReceiveResult
func (u PathPaymentStrictReceiveResult) ArmForSwitch(sw int32) (string, bool) {
	switch PathPaymentStrictReceiveResultCode(sw) {
	case PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveSuccess:
		return "Success", true
	case PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveNoIssuer:
		return "NoIssuer", true
	default:
		return "", true
	}
}

// NewPathPaymentStrictReceiveResult creates a new  PathPaymentStrictReceiveResult.
func NewPathPaymentStrictReceiveResult(code PathPaymentStrictReceiveResultCode, value interface{}) (result PathPaymentStrictReceiveResult, err error) {
	result.Code = code
	switch PathPaymentStrictReceiveResultCode(code) {
	case PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveSuccess:
		tv, ok := value.(PathPaymentStrictReceiveResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be PathPaymentStrictReceiveResultSuccess")
			return
		}
		result.Success = &tv
	case PathPaymentStrictReceiveResultCodePathPaymentStrictReceiveNoIssuer:
		tv, ok := value.(Asset)
		if !ok {
			err = fmt.Errorf("invalid value, must be Asset")
			return
		}
		result.NoIssuer = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u PathPaymentStrictReceiveResult) MustSuccess() PathPaymentStrictReceiveResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PathPaymentStrictReceiveResult) GetSuccess() (result PathPaymentStrictReceiveResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// MustNoIssuer retrieves the NoIssuer value from the union,
// panicing if the value is not set.
func (u PathPaymentStrictReceiveResult) MustNoIssuer() Asset {
	val, ok := u.GetNoIssuer()

	if !ok {
		panic("arm NoIssuer is not set")
	}

	return val
}

// GetNoIssuer retrieves the NoIssuer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PathPaymentStrictReceiveResult) GetNoIssuer() (result Asset, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "NoIssuer" {
		result = *u.NoIssuer
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PathPaymentStrictReceiveResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PathPaymentStrictReceiveResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PathPaymentStrictReceiveResult)(nil)
	_ encoding.BinaryUnmarshaler = (*PathPaymentStrictReceiveResult)(nil)
)

// PathPaymentStrictSendResultCode is an XDR Enum defines as:
//
//   enum PathPaymentStrictSendResultCode
//    {
//        // codes considered as "success" for the operation
//        PATH_PAYMENT_STRICT_SEND_SUCCESS = 0, // success
//
//        // codes considered as "failure" for the operation
//        PATH_PAYMENT_STRICT_SEND_MALFORMED = -1, // bad input
//        PATH_PAYMENT_STRICT_SEND_UNDERFUNDED =
//            -2, // not enough funds in source account
//        PATH_PAYMENT_STRICT_SEND_SRC_NO_TRUST =
//            -3, // no trust line on source account
//        PATH_PAYMENT_STRICT_SEND_SRC_NOT_AUTHORIZED =
//            -4, // source not authorized to transfer
//        PATH_PAYMENT_STRICT_SEND_NO_DESTINATION =
//            -5, // destination account does not exist
//        PATH_PAYMENT_STRICT_SEND_NO_TRUST =
//            -6, // dest missing a trust line for asset
//        PATH_PAYMENT_STRICT_SEND_NOT_AUTHORIZED =
//            -7, // dest not authorized to hold asset
//        PATH_PAYMENT_STRICT_SEND_LINE_FULL = -8, // dest would go above their limit
//        PATH_PAYMENT_STRICT_SEND_NO_ISSUER = -9, // missing issuer on one asset
//        PATH_PAYMENT_STRICT_SEND_TOO_FEW_OFFERS =
//            -10, // not enough offers to satisfy path
//        PATH_PAYMENT_STRICT_SEND_OFFER_CROSS_SELF =
//            -11, // would cross one of its own offers
//        PATH_PAYMENT_STRICT_SEND_UNDER_DESTMIN = -12 // could not satisfy destMin
//    };
//
type PathPaymentStrictSendResultCode int32

const (
	PathPaymentStrictSendResultCodePathPaymentStrictSendSuccess          PathPaymentStrictSendResultCode = 0
	PathPaymentStrictSendResultCodePathPaymentStrictSendMalformed        PathPaymentStrictSendResultCode = -1
	PathPaymentStrictSendResultCodePathPaymentStrictSendUnderfunded      PathPaymentStrictSendResultCode = -2
	PathPaymentStrictSendResultCodePathPaymentStrictSendSrcNoTrust       PathPaymentStrictSendResultCode = -3
	PathPaymentStrictSendResultCodePathPaymentStrictSendSrcNotAuthorized PathPaymentStrictSendResultCode = -4
	PathPaymentStrictSendResultCodePathPaymentStrictSendNoDestination    PathPaymentStrictSendResultCode = -5
	PathPaymentStrictSendResultCodePathPaymentStrictSendNoTrust          PathPaymentStrictSendResultCode = -6
	PathPaymentStrictSendResultCodePathPaymentStrictSendNotAuthorized    PathPaymentStrictSendResultCode = -7
	PathPaymentStrictSendResultCodePathPaymentStrictSendLineFull         PathPaymentStrictSendResultCode = -8
	PathPaymentStrictSendResultCodePathPaymentStrictSendNoIssuer         PathPaymentStrictSendResultCode = -9
	PathPaymentStrictSendResultCodePathPaymentStrictSendTooFewOffers     PathPaymentStrictSendResultCode = -10
	PathPaymentStrictSendResultCodePathPaymentStrictSendOfferCrossSelf   PathPaymentStrictSendResultCode = -11
	PathPaymentStrictSendResultCodePathPaymentStrictSendUnderDestmin     PathPaymentStrictSendResultCode = -12
)

var pathPaymentStrictSendResultCodeMap = map[int32]string{
	0:   "PathPaymentStrictSendResultCodePathPaymentStrictSendSuccess",
	-1:  "PathPaymentStrictSendResultCodePathPaymentStrictSendMalformed",
	-2:  "PathPaymentStrictSendResultCodePathPaymentStrictSendUnderfunded",
	-3:  "PathPaymentStrictSendResultCodePathPaymentStrictSendSrcNoTrust",
	-4:  "PathPaymentStrictSendResultCodePathPaymentStrictSendSrcNotAuthorized",
	-5:  "PathPaymentStrictSendResultCodePathPaymentStrictSendNoDestination",
	-6:  "PathPaymentStrictSendResultCodePathPaymentStrictSendNoTrust",
	-7:  "PathPaymentStrictSendResultCodePathPaymentStrictSendNotAuthorized",
	-8:  "PathPaymentStrictSendResultCodePathPaymentStrictSendLineFull",
	-9:  "PathPaymentStrictSendResultCodePathPaymentStrictSendNoIssuer",
	-10: "PathPaymentStrictSendResultCodePathPaymentStrictSendTooFewOffers",
	-11: "PathPaymentStrictSendResultCodePathPaymentStrictSendOfferCrossSelf",
	-12: "PathPaymentStrictSendResultCodePathPaymentStrictSendUnderDestmin",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PathPaymentStrictSendResultCode
func (e PathPaymentStrictSendResultCode) ValidEnum(v int32) bool {
	_, ok := pathPaymentStrictSendResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e PathPaymentStrictSendResultCode) String() string {
	name, _ := pathPaymentStrictSendResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PathPaymentStrictSendResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PathPaymentStrictSendResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PathPaymentStrictSendResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*PathPaymentStrictSendResultCode)(nil)
)

// PathPaymentStrictSendResultSuccess is an XDR NestedStruct defines as:
//
//   struct
//        {
//            ClaimAtom offers<>;
//            SimplePaymentResult last;
//        }
//
type PathPaymentStrictSendResultSuccess struct {
	Offers []ClaimAtom
	Last   SimplePaymentResult
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PathPaymentStrictSendResultSuccess) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PathPaymentStrictSendResultSuccess) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PathPaymentStrictSendResultSuccess)(nil)
	_ encoding.BinaryUnmarshaler = (*PathPaymentStrictSendResultSuccess)(nil)
)

// PathPaymentStrictSendResult is an XDR Union defines as:
//
//   union PathPaymentStrictSendResult switch (PathPaymentStrictSendResultCode code)
//    {
//    case PATH_PAYMENT_STRICT_SEND_SUCCESS:
//        struct
//        {
//            ClaimAtom offers<>;
//            SimplePaymentResult last;
//        } success;
//    case PATH_PAYMENT_STRICT_SEND_NO_ISSUER:
//        Asset noIssuer; // the asset that caused the error
//    default:
//        void;
//    };
//
type PathPaymentStrictSendResult struct {
	Code     PathPaymentStrictSendResultCode
	Success  *PathPaymentStrictSendResultSuccess
	NoIssuer *Asset
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PathPaymentStrictSendResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PathPaymentStrictSendResult
func (u PathPaymentStrictSendResult) ArmForSwitch(sw int32) (string, bool) {
	switch PathPaymentStrictSendResultCode(sw) {
	case PathPaymentStrictSendResultCodePathPaymentStrictSendSuccess:
		return "Success", true
	case PathPaymentStrictSendResultCodePathPaymentStrictSendNoIssuer:
		return "NoIssuer", true
	default:
		return "", true
	}
}

// NewPathPaymentStrictSendResult creates a new  PathPaymentStrictSendResult.
func NewPathPaymentStrictSendResult(code PathPaymentStrictSendResultCode, value interface{}) (result PathPaymentStrictSendResult, err error) {
	result.Code = code
	switch PathPaymentStrictSendResultCode(code) {
	case PathPaymentStrictSendResultCodePathPaymentStrictSendSuccess:
		tv, ok := value.(PathPaymentStrictSendResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be PathPaymentStrictSendResultSuccess")
			return
		}
		result.Success = &tv
	case PathPaymentStrictSendResultCodePathPaymentStrictSendNoIssuer:
		tv, ok := value.(Asset)
		if !ok {
			err = fmt.Errorf("invalid value, must be Asset")
			return
		}
		result.NoIssuer = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u PathPaymentStrictSendResult) MustSuccess() PathPaymentStrictSendResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PathPaymentStrictSendResult) GetSuccess() (result PathPaymentStrictSendResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// MustNoIssuer retrieves the NoIssuer value from the union,
// panicing if the value is not set.
func (u PathPaymentStrictSendResult) MustNoIssuer() Asset {
	val, ok := u.GetNoIssuer()

	if !ok {
		panic("arm NoIssuer is not set")
	}

	return val
}

// GetNoIssuer retrieves the NoIssuer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PathPaymentStrictSendResult) GetNoIssuer() (result Asset, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "NoIssuer" {
		result = *u.NoIssuer
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PathPaymentStrictSendResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PathPaymentStrictSendResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PathPaymentStrictSendResult)(nil)
	_ encoding.BinaryUnmarshaler = (*PathPaymentStrictSendResult)(nil)
)

// ManageSellOfferResultCode is an XDR Enum defines as:
//
//   enum ManageSellOfferResultCode
//    {
//        // codes considered as "success" for the operation
//        MANAGE_SELL_OFFER_SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        MANAGE_SELL_OFFER_MALFORMED = -1, // generated offer would be invalid
//        MANAGE_SELL_OFFER_SELL_NO_TRUST =
//            -2,                              // no trust line for what we're selling
//        MANAGE_SELL_OFFER_BUY_NO_TRUST = -3, // no trust line for what we're buying
//        MANAGE_SELL_OFFER_SELL_NOT_AUTHORIZED = -4, // not authorized to sell
//        MANAGE_SELL_OFFER_BUY_NOT_AUTHORIZED = -5,  // not authorized to buy
//        MANAGE_SELL_OFFER_LINE_FULL = -6, // can't receive more of what it's buying
//        MANAGE_SELL_OFFER_UNDERFUNDED = -7, // doesn't hold what it's trying to sell
//        MANAGE_SELL_OFFER_CROSS_SELF =
//            -8, // would cross an offer from the same user
//        MANAGE_SELL_OFFER_SELL_NO_ISSUER = -9, // no issuer for what we're selling
//        MANAGE_SELL_OFFER_BUY_NO_ISSUER = -10, // no issuer for what we're buying
//
//        // update errors
//        MANAGE_SELL_OFFER_NOT_FOUND =
//            -11, // offerID does not match an existing offer
//
//        MANAGE_SELL_OFFER_LOW_RESERVE =
//            -12 // not enough funds to create a new Offer
//    };
//
type ManageSellOfferResultCode int32

const (
	ManageSellOfferResultCodeManageSellOfferSuccess           ManageSellOfferResultCode = 0
	ManageSellOfferResultCodeManageSellOfferMalformed         ManageSellOfferResultCode = -1
	ManageSellOfferResultCodeManageSellOfferSellNoTrust       ManageSellOfferResultCode = -2
	ManageSellOfferResultCodeManageSellOfferBuyNoTrust        ManageSellOfferResultCode = -3
	ManageSellOfferResultCodeManageSellOfferSellNotAuthorized ManageSellOfferResultCode = -4
	ManageSellOfferResultCodeManageSellOfferBuyNotAuthorized  ManageSellOfferResultCode = -5
	ManageSellOfferResultCodeManageSellOfferLineFull          ManageSellOfferResultCode = -6
	ManageSellOfferResultCodeManageSellOfferUnderfunded       ManageSellOfferResultCode = -7
	ManageSellOfferResultCodeManageSellOfferCrossSelf         ManageSellOfferResultCode = -8
	ManageSellOfferResultCodeManageSellOfferSellNoIssuer      ManageSellOfferResultCode = -9
	ManageSellOfferResultCodeManageSellOfferBuyNoIssuer       ManageSellOfferResultCode = -10
	ManageSellOfferResultCodeManageSellOfferNotFound          ManageSellOfferResultCode = -11
	ManageSellOfferResultCodeManageSellOfferLowReserve        ManageSellOfferResultCode = -12
)

var manageSellOfferResultCodeMap = map[int32]string{
	0:   "ManageSellOfferResultCodeManageSellOfferSuccess",
	-1:  "ManageSellOfferResultCodeManageSellOfferMalformed",
	-2:  "ManageSellOfferResultCodeManageSellOfferSellNoTrust",
	-3:  "ManageSellOfferResultCodeManageSellOfferBuyNoTrust",
	-4:  "ManageSellOfferResultCodeManageSellOfferSellNotAuthorized",
	-5:  "ManageSellOfferResultCodeManageSellOfferBuyNotAuthorized",
	-6:  "ManageSellOfferResultCodeManageSellOfferLineFull",
	-7:  "ManageSellOfferResultCodeManageSellOfferUnderfunded",
	-8:  "ManageSellOfferResultCodeManageSellOfferCrossSelf",
	-9:  "ManageSellOfferResultCodeManageSellOfferSellNoIssuer",
	-10: "ManageSellOfferResultCodeManageSellOfferBuyNoIssuer",
	-11: "ManageSellOfferResultCodeManageSellOfferNotFound",
	-12: "ManageSellOfferResultCodeManageSellOfferLowReserve",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageSellOfferResultCode
func (e ManageSellOfferResultCode) ValidEnum(v int32) bool {
	_, ok := manageSellOfferResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e ManageSellOfferResultCode) String() string {
	name, _ := manageSellOfferResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ManageSellOfferResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ManageSellOfferResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ManageSellOfferResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*ManageSellOfferResultCode)(nil)
)

// ManageOfferEffect is an XDR Enum defines as:
//
//   enum ManageOfferEffect
//    {
//        MANAGE_OFFER_CREATED = 0,
//        MANAGE_OFFER_UPDATED = 1,
//        MANAGE_OFFER_DELETED = 2
//    };
//
type ManageOfferEffect int32

const (
	ManageOfferEffectManageOfferCreated ManageOfferEffect = 0
	ManageOfferEffectManageOfferUpdated ManageOfferEffect = 1
	ManageOfferEffectManageOfferDeleted ManageOfferEffect = 2
)

var manageOfferEffectMap = map[int32]string{
	0: "ManageOfferEffectManageOfferCreated",
	1: "ManageOfferEffectManageOfferUpdated",
	2: "ManageOfferEffectManageOfferDeleted",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageOfferEffect
func (e ManageOfferEffect) ValidEnum(v int32) bool {
	_, ok := manageOfferEffectMap[v]
	return ok
}

// String returns the name of `e`
func (e ManageOfferEffect) String() string {
	name, _ := manageOfferEffectMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ManageOfferEffect) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ManageOfferEffect) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ManageOfferEffect)(nil)
	_ encoding.BinaryUnmarshaler = (*ManageOfferEffect)(nil)
)

// ManageOfferSuccessResultOffer is an XDR NestedUnion defines as:
//
//   union switch (ManageOfferEffect effect)
//        {
//        case MANAGE_OFFER_CREATED:
//        case MANAGE_OFFER_UPDATED:
//            OfferEntry offer;
//        default:
//            void;
//        }
//
type ManageOfferSuccessResultOffer struct {
	Effect ManageOfferEffect
	Offer  *OfferEntry
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageOfferSuccessResultOffer) SwitchFieldName() string {
	return "Effect"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageOfferSuccessResultOffer
func (u ManageOfferSuccessResultOffer) ArmForSwitch(sw int32) (string, bool) {
	switch ManageOfferEffect(sw) {
	case ManageOfferEffectManageOfferCreated:
		return "Offer", true
	case ManageOfferEffectManageOfferUpdated:
		return "Offer", true
	default:
		return "", true
	}
}

// NewManageOfferSuccessResultOffer creates a new  ManageOfferSuccessResultOffer.
func NewManageOfferSuccessResultOffer(effect ManageOfferEffect, value interface{}) (result ManageOfferSuccessResultOffer, err error) {
	result.Effect = effect
	switch ManageOfferEffect(effect) {
	case ManageOfferEffectManageOfferCreated:
		tv, ok := value.(OfferEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be OfferEntry")
			return
		}
		result.Offer = &tv
	case ManageOfferEffectManageOfferUpdated:
		tv, ok := value.(OfferEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be OfferEntry")
			return
		}
		result.Offer = &tv
	default:
		// void
	}
	return
}

// MustOffer retrieves the Offer value from the union,
// panicing if the value is not set.
func (u ManageOfferSuccessResultOffer) MustOffer() OfferEntry {
	val, ok := u.GetOffer()

	if !ok {
		panic("arm Offer is not set")
	}

	return val
}

// GetOffer retrieves the Offer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageOfferSuccessResultOffer) GetOffer() (result OfferEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Effect))

	if armName == "Offer" {
		result = *u.Offer
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ManageOfferSuccessResultOffer) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ManageOfferSuccessResultOffer) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ManageOfferSuccessResultOffer)(nil)
	_ encoding.BinaryUnmarshaler = (*ManageOfferSuccessResultOffer)(nil)
)

// ManageOfferSuccessResult is an XDR Struct defines as:
//
//   struct ManageOfferSuccessResult
//    {
//        // offers that got claimed while creating this offer
//        ClaimAtom offersClaimed<>;
//
//        union switch (ManageOfferEffect effect)
//        {
//        case MANAGE_OFFER_CREATED:
//        case MANAGE_OFFER_UPDATED:
//            OfferEntry offer;
//        default:
//            void;
//        }
//        offer;
//    };
//
type ManageOfferSuccessResult struct {
	OffersClaimed []ClaimAtom
	Offer         ManageOfferSuccessResultOffer
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ManageOfferSuccessResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ManageOfferSuccessResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ManageOfferSuccessResult)(nil)
	_ encoding.BinaryUnmarshaler = (*ManageOfferSuccessResult)(nil)
)

// ManageSellOfferResult is an XDR Union defines as:
//
//   union ManageSellOfferResult switch (ManageSellOfferResultCode code)
//    {
//    case MANAGE_SELL_OFFER_SUCCESS:
//        ManageOfferSuccessResult success;
//    default:
//        void;
//    };
//
type ManageSellOfferResult struct {
	Code    ManageSellOfferResultCode
	Success *ManageOfferSuccessResult
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageSellOfferResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageSellOfferResult
func (u ManageSellOfferResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageSellOfferResultCode(sw) {
	case ManageSellOfferResultCodeManageSellOfferSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageSellOfferResult creates a new  ManageSellOfferResult.
func NewManageSellOfferResult(code ManageSellOfferResultCode, value interface{}) (result ManageSellOfferResult, err error) {
	result.Code = code
	switch ManageSellOfferResultCode(code) {
	case ManageSellOfferResultCodeManageSellOfferSuccess:
		tv, ok := value.(ManageOfferSuccessResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageOfferSuccessResult")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageSellOfferResult) MustSuccess() ManageOfferSuccessResult {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSellOfferResult) GetSuccess() (result ManageOfferSuccessResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ManageSellOfferResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ManageSellOfferResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ManageSellOfferResult)(nil)
	_ encoding.BinaryUnmarshaler = (*ManageSellOfferResult)(nil)
)

// ManageBuyOfferResultCode is an XDR Enum defines as:
//
//   enum ManageBuyOfferResultCode
//    {
//        // codes considered as "success" for the operation
//        MANAGE_BUY_OFFER_SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        MANAGE_BUY_OFFER_MALFORMED = -1,     // generated offer would be invalid
//        MANAGE_BUY_OFFER_SELL_NO_TRUST = -2, // no trust line for what we're selling
//        MANAGE_BUY_OFFER_BUY_NO_TRUST = -3,  // no trust line for what we're buying
//        MANAGE_BUY_OFFER_SELL_NOT_AUTHORIZED = -4, // not authorized to sell
//        MANAGE_BUY_OFFER_BUY_NOT_AUTHORIZED = -5,  // not authorized to buy
//        MANAGE_BUY_OFFER_LINE_FULL = -6,   // can't receive more of what it's buying
//        MANAGE_BUY_OFFER_UNDERFUNDED = -7, // doesn't hold what it's trying to sell
//        MANAGE_BUY_OFFER_CROSS_SELF = -8, // would cross an offer from the same user
//        MANAGE_BUY_OFFER_SELL_NO_ISSUER = -9, // no issuer for what we're selling
//        MANAGE_BUY_OFFER_BUY_NO_ISSUER = -10, // no issuer for what we're buying
//
//        // update errors
//        MANAGE_BUY_OFFER_NOT_FOUND =
//            -11, // offerID does not match an existing offer
//
//        MANAGE_BUY_OFFER_LOW_RESERVE = -12 // not enough funds to create a new Offer
//    };
//
type ManageBuyOfferResultCode int32

const (
	ManageBuyOfferResultCodeManageBuyOfferSuccess           ManageBuyOfferResultCode = 0
	ManageBuyOfferResultCodeManageBuyOfferMalformed         ManageBuyOfferResultCode = -1
	ManageBuyOfferResultCodeManageBuyOfferSellNoTrust       ManageBuyOfferResultCode = -2
	ManageBuyOfferResultCodeManageBuyOfferBuyNoTrust        ManageBuyOfferResultCode = -3
	ManageBuyOfferResultCodeManageBuyOfferSellNotAuthorized ManageBuyOfferResultCode = -4
	ManageBuyOfferResultCodeManageBuyOfferBuyNotAuthorized  ManageBuyOfferResultCode = -5
	ManageBuyOfferResultCodeManageBuyOfferLineFull          ManageBuyOfferResultCode = -6
	ManageBuyOfferResultCodeManageBuyOfferUnderfunded       ManageBuyOfferResultCode = -7
	ManageBuyOfferResultCodeManageBuyOfferCrossSelf         ManageBuyOfferResultCode = -8
	ManageBuyOfferResultCodeManageBuyOfferSellNoIssuer      ManageBuyOfferResultCode = -9
	ManageBuyOfferResultCodeManageBuyOfferBuyNoIssuer       ManageBuyOfferResultCode = -10
	ManageBuyOfferResultCodeManageBuyOfferNotFound          ManageBuyOfferResultCode = -11
	ManageBuyOfferResultCodeManageBuyOfferLowReserve        ManageBuyOfferResultCode = -12
)

var manageBuyOfferResultCodeMap = map[int32]string{
	0:   "ManageBuyOfferResultCodeManageBuyOfferSuccess",
	-1:  "ManageBuyOfferResultCodeManageBuyOfferMalformed",
	-2:  "ManageBuyOfferResultCodeManageBuyOfferSellNoTrust",
	-3:  "ManageBuyOfferResultCodeManageBuyOfferBuyNoTrust",
	-4:  "ManageBuyOfferResultCodeManageBuyOfferSellNotAuthorized",
	-5:  "ManageBuyOfferResultCodeManageBuyOfferBuyNotAuthorized",
	-6:  "ManageBuyOfferResultCodeManageBuyOfferLineFull",
	-7:  "ManageBuyOfferResultCodeManageBuyOfferUnderfunded",
	-8:  "ManageBuyOfferResultCodeManageBuyOfferCrossSelf",
	-9:  "ManageBuyOfferResultCodeManageBuyOfferSellNoIssuer",
	-10: "ManageBuyOfferResultCodeManageBuyOfferBuyNoIssuer",
	-11: "ManageBuyOfferResultCodeManageBuyOfferNotFound",
	-12: "ManageBuyOfferResultCodeManageBuyOfferLowReserve",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageBuyOfferResultCode
func (e ManageBuyOfferResultCode) ValidEnum(v int32) bool {
	_, ok := manageBuyOfferResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e ManageBuyOfferResultCode) String() string {
	name, _ := manageBuyOfferResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ManageBuyOfferResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ManageBuyOfferResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ManageBuyOfferResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*ManageBuyOfferResultCode)(nil)
)

// ManageBuyOfferResult is an XDR Union defines as:
//
//   union ManageBuyOfferResult switch (ManageBuyOfferResultCode code)
//    {
//    case MANAGE_BUY_OFFER_SUCCESS:
//        ManageOfferSuccessResult success;
//    default:
//        void;
//    };
//
type ManageBuyOfferResult struct {
	Code    ManageBuyOfferResultCode
	Success *ManageOfferSuccessResult
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageBuyOfferResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageBuyOfferResult
func (u ManageBuyOfferResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageBuyOfferResultCode(sw) {
	case ManageBuyOfferResultCodeManageBuyOfferSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageBuyOfferResult creates a new  ManageBuyOfferResult.
func NewManageBuyOfferResult(code ManageBuyOfferResultCode, value interface{}) (result ManageBuyOfferResult, err error) {
	result.Code = code
	switch ManageBuyOfferResultCode(code) {
	case ManageBuyOfferResultCodeManageBuyOfferSuccess:
		tv, ok := value.(ManageOfferSuccessResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageOfferSuccessResult")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageBuyOfferResult) MustSuccess() ManageOfferSuccessResult {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageBuyOfferResult) GetSuccess() (result ManageOfferSuccessResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ManageBuyOfferResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ManageBuyOfferResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ManageBuyOfferResult)(nil)
	_ encoding.BinaryUnmarshaler = (*ManageBuyOfferResult)(nil)
)

// SetOptionsResultCode is an XDR Enum defines as:
//
//   enum SetOptionsResultCode
//    {
//        // codes considered as "success" for the operation
//        SET_OPTIONS_SUCCESS = 0,
//        // codes considered as "failure" for the operation
//        SET_OPTIONS_LOW_RESERVE = -1,      // not enough funds to add a signer
//        SET_OPTIONS_TOO_MANY_SIGNERS = -2, // max number of signers already reached
//        SET_OPTIONS_BAD_FLAGS = -3,        // invalid combination of clear/set flags
//        SET_OPTIONS_INVALID_INFLATION = -4,      // inflation account does not exist
//        SET_OPTIONS_CANT_CHANGE = -5,            // can no longer change this option
//        SET_OPTIONS_UNKNOWN_FLAG = -6,           // can't set an unknown flag
//        SET_OPTIONS_THRESHOLD_OUT_OF_RANGE = -7, // bad value for weight/threshold
//        SET_OPTIONS_BAD_SIGNER = -8,             // signer cannot be masterkey
//        SET_OPTIONS_INVALID_HOME_DOMAIN = -9,    // malformed home domain
//        SET_OPTIONS_AUTH_REVOCABLE_REQUIRED =
//            -10 // auth revocable is required for clawback
//    };
//
type SetOptionsResultCode int32

const (
	SetOptionsResultCodeSetOptionsSuccess               SetOptionsResultCode = 0
	SetOptionsResultCodeSetOptionsLowReserve            SetOptionsResultCode = -1
	SetOptionsResultCodeSetOptionsTooManySigners        SetOptionsResultCode = -2
	SetOptionsResultCodeSetOptionsBadFlags              SetOptionsResultCode = -3
	SetOptionsResultCodeSetOptionsInvalidInflation      SetOptionsResultCode = -4
	SetOptionsResultCodeSetOptionsCantChange            SetOptionsResultCode = -5
	SetOptionsResultCodeSetOptionsUnknownFlag           SetOptionsResultCode = -6
	SetOptionsResultCodeSetOptionsThresholdOutOfRange   SetOptionsResultCode = -7
	SetOptionsResultCodeSetOptionsBadSigner             SetOptionsResultCode = -8
	SetOptionsResultCodeSetOptionsInvalidHomeDomain     SetOptionsResultCode = -9
	SetOptionsResultCodeSetOptionsAuthRevocableRequired SetOptionsResultCode = -10
)

var setOptionsResultCodeMap = map[int32]string{
	0:   "SetOptionsResultCodeSetOptionsSuccess",
	-1:  "SetOptionsResultCodeSetOptionsLowReserve",
	-2:  "SetOptionsResultCodeSetOptionsTooManySigners",
	-3:  "SetOptionsResultCodeSetOptionsBadFlags",
	-4:  "SetOptionsResultCodeSetOptionsInvalidInflation",
	-5:  "SetOptionsResultCodeSetOptionsCantChange",
	-6:  "SetOptionsResultCodeSetOptionsUnknownFlag",
	-7:  "SetOptionsResultCodeSetOptionsThresholdOutOfRange",
	-8:  "SetOptionsResultCodeSetOptionsBadSigner",
	-9:  "SetOptionsResultCodeSetOptionsInvalidHomeDomain",
	-10: "SetOptionsResultCodeSetOptionsAuthRevocableRequired",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for SetOptionsResultCode
func (e SetOptionsResultCode) ValidEnum(v int32) bool {
	_, ok := setOptionsResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e SetOptionsResultCode) String() string {
	name, _ := setOptionsResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SetOptionsResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SetOptionsResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SetOptionsResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*SetOptionsResultCode)(nil)
)

// SetOptionsResult is an XDR Union defines as:
//
//   union SetOptionsResult switch (SetOptionsResultCode code)
//    {
//    case SET_OPTIONS_SUCCESS:
//        void;
//    default:
//        void;
//    };
//
type SetOptionsResult struct {
	Code SetOptionsResultCode
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetOptionsResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetOptionsResult
func (u SetOptionsResult) ArmForSwitch(sw int32) (string, bool) {
	switch SetOptionsResultCode(sw) {
	case SetOptionsResultCodeSetOptionsSuccess:
		return "", true
	default:
		return "", true
	}
}

// NewSetOptionsResult creates a new  SetOptionsResult.
func NewSetOptionsResult(code SetOptionsResultCode, value interface{}) (result SetOptionsResult, err error) {
	result.Code = code
	switch SetOptionsResultCode(code) {
	case SetOptionsResultCodeSetOptionsSuccess:
		// void
	default:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SetOptionsResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SetOptionsResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SetOptionsResult)(nil)
	_ encoding.BinaryUnmarshaler = (*SetOptionsResult)(nil)
)

// ChangeTrustResultCode is an XDR Enum defines as:
//
//   enum ChangeTrustResultCode
//    {
//        // codes considered as "success" for the operation
//        CHANGE_TRUST_SUCCESS = 0,
//        // codes considered as "failure" for the operation
//        CHANGE_TRUST_MALFORMED = -1,     // bad input
//        CHANGE_TRUST_NO_ISSUER = -2,     // could not find issuer
//        CHANGE_TRUST_INVALID_LIMIT = -3, // cannot drop limit below balance
//                                         // cannot create with a limit of 0
//        CHANGE_TRUST_LOW_RESERVE =
//            -4, // not enough funds to create a new trust line,
//        CHANGE_TRUST_SELF_NOT_ALLOWED = -5, // trusting self is not allowed
//        CHANGE_TRUST_TRUST_LINE_MISSING = -6, // Asset trustline is missing for pool
//        CHANGE_TRUST_CANNOT_DELETE = -7, // Asset trustline is still referenced in a pool
//        CHANGE_TRUST_NOT_AUTH_MAINTAIN_LIABILITIES = -8 // Asset trustline is deauthorized
//    };
//
type ChangeTrustResultCode int32

const (
	ChangeTrustResultCodeChangeTrustSuccess                    ChangeTrustResultCode = 0
	ChangeTrustResultCodeChangeTrustMalformed                  ChangeTrustResultCode = -1
	ChangeTrustResultCodeChangeTrustNoIssuer                   ChangeTrustResultCode = -2
	ChangeTrustResultCodeChangeTrustInvalidLimit               ChangeTrustResultCode = -3
	ChangeTrustResultCodeChangeTrustLowReserve                 ChangeTrustResultCode = -4
	ChangeTrustResultCodeChangeTrustSelfNotAllowed             ChangeTrustResultCode = -5
	ChangeTrustResultCodeChangeTrustTrustLineMissing           ChangeTrustResultCode = -6
	ChangeTrustResultCodeChangeTrustCannotDelete               ChangeTrustResultCode = -7
	ChangeTrustResultCodeChangeTrustNotAuthMaintainLiabilities ChangeTrustResultCode = -8
)

var changeTrustResultCodeMap = map[int32]string{
	0:  "ChangeTrustResultCodeChangeTrustSuccess",
	-1: "ChangeTrustResultCodeChangeTrustMalformed",
	-2: "ChangeTrustResultCodeChangeTrustNoIssuer",
	-3: "ChangeTrustResultCodeChangeTrustInvalidLimit",
	-4: "ChangeTrustResultCodeChangeTrustLowReserve",
	-5: "ChangeTrustResultCodeChangeTrustSelfNotAllowed",
	-6: "ChangeTrustResultCodeChangeTrustTrustLineMissing",
	-7: "ChangeTrustResultCodeChangeTrustCannotDelete",
	-8: "ChangeTrustResultCodeChangeTrustNotAuthMaintainLiabilities",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ChangeTrustResultCode
func (e ChangeTrustResultCode) ValidEnum(v int32) bool {
	_, ok := changeTrustResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e ChangeTrustResultCode) String() string {
	name, _ := changeTrustResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ChangeTrustResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ChangeTrustResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ChangeTrustResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*ChangeTrustResultCode)(nil)
)

// ChangeTrustResult is an XDR Union defines as:
//
//   union ChangeTrustResult switch (ChangeTrustResultCode code)
//    {
//    case CHANGE_TRUST_SUCCESS:
//        void;
//    default:
//        void;
//    };
//
type ChangeTrustResult struct {
	Code ChangeTrustResultCode
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ChangeTrustResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ChangeTrustResult
func (u ChangeTrustResult) ArmForSwitch(sw int32) (string, bool) {
	switch ChangeTrustResultCode(sw) {
	case ChangeTrustResultCodeChangeTrustSuccess:
		return "", true
	default:
		return "", true
	}
}

// NewChangeTrustResult creates a new  ChangeTrustResult.
func NewChangeTrustResult(code ChangeTrustResultCode, value interface{}) (result ChangeTrustResult, err error) {
	result.Code = code
	switch ChangeTrustResultCode(code) {
	case ChangeTrustResultCodeChangeTrustSuccess:
		// void
	default:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ChangeTrustResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ChangeTrustResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ChangeTrustResult)(nil)
	_ encoding.BinaryUnmarshaler = (*ChangeTrustResult)(nil)
)

// AllowTrustResultCode is an XDR Enum defines as:
//
//   enum AllowTrustResultCode
//    {
//        // codes considered as "success" for the operation
//        ALLOW_TRUST_SUCCESS = 0,
//        // codes considered as "failure" for the operation
//        ALLOW_TRUST_MALFORMED = -1,     // asset is not ASSET_TYPE_ALPHANUM
//        ALLOW_TRUST_NO_TRUST_LINE = -2, // trustor does not have a trustline
//                                        // source account does not require trust
//        ALLOW_TRUST_TRUST_NOT_REQUIRED = -3,
//        ALLOW_TRUST_CANT_REVOKE = -4,     // source account can't revoke trust,
//        ALLOW_TRUST_SELF_NOT_ALLOWED = -5 // trusting self is not allowed
//    };
//
type AllowTrustResultCode int32

const (
	AllowTrustResultCodeAllowTrustSuccess          AllowTrustResultCode = 0
	AllowTrustResultCodeAllowTrustMalformed        AllowTrustResultCode = -1
	AllowTrustResultCodeAllowTrustNoTrustLine      AllowTrustResultCode = -2
	AllowTrustResultCodeAllowTrustTrustNotRequired AllowTrustResultCode = -3
	AllowTrustResultCodeAllowTrustCantRevoke       AllowTrustResultCode = -4
	AllowTrustResultCodeAllowTrustSelfNotAllowed   AllowTrustResultCode = -5
)

var allowTrustResultCodeMap = map[int32]string{
	0:  "AllowTrustResultCodeAllowTrustSuccess",
	-1: "AllowTrustResultCodeAllowTrustMalformed",
	-2: "AllowTrustResultCodeAllowTrustNoTrustLine",
	-3: "AllowTrustResultCodeAllowTrustTrustNotRequired",
	-4: "AllowTrustResultCodeAllowTrustCantRevoke",
	-5: "AllowTrustResultCodeAllowTrustSelfNotAllowed",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for AllowTrustResultCode
func (e AllowTrustResultCode) ValidEnum(v int32) bool {
	_, ok := allowTrustResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e AllowTrustResultCode) String() string {
	name, _ := allowTrustResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AllowTrustResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AllowTrustResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AllowTrustResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*AllowTrustResultCode)(nil)
)

// AllowTrustResult is an XDR Union defines as:
//
//   union AllowTrustResult switch (AllowTrustResultCode code)
//    {
//    case ALLOW_TRUST_SUCCESS:
//        void;
//    default:
//        void;
//    };
//
type AllowTrustResult struct {
	Code AllowTrustResultCode
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AllowTrustResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AllowTrustResult
func (u AllowTrustResult) ArmForSwitch(sw int32) (string, bool) {
	switch AllowTrustResultCode(sw) {
	case AllowTrustResultCodeAllowTrustSuccess:
		return "", true
	default:
		return "", true
	}
}

// NewAllowTrustResult creates a new  AllowTrustResult.
func NewAllowTrustResult(code AllowTrustResultCode, value interface{}) (result AllowTrustResult, err error) {
	result.Code = code
	switch AllowTrustResultCode(code) {
	case AllowTrustResultCodeAllowTrustSuccess:
		// void
	default:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AllowTrustResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AllowTrustResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AllowTrustResult)(nil)
	_ encoding.BinaryUnmarshaler = (*AllowTrustResult)(nil)
)

// AccountMergeResultCode is an XDR Enum defines as:
//
//   enum AccountMergeResultCode
//    {
//        // codes considered as "success" for the operation
//        ACCOUNT_MERGE_SUCCESS = 0,
//        // codes considered as "failure" for the operation
//        ACCOUNT_MERGE_MALFORMED = -1,       // can't merge onto itself
//        ACCOUNT_MERGE_NO_ACCOUNT = -2,      // destination does not exist
//        ACCOUNT_MERGE_IMMUTABLE_SET = -3,   // source account has AUTH_IMMUTABLE set
//        ACCOUNT_MERGE_HAS_SUB_ENTRIES = -4, // account has trust lines/offers
//        ACCOUNT_MERGE_SEQNUM_TOO_FAR = -5,  // sequence number is over max allowed
//        ACCOUNT_MERGE_DEST_FULL = -6,       // can't add source balance to
//                                            // destination balance
//        ACCOUNT_MERGE_IS_SPONSOR = -7       // can't merge account that is a sponsor
//    };
//
type AccountMergeResultCode int32

const (
	AccountMergeResultCodeAccountMergeSuccess       AccountMergeResultCode = 0
	AccountMergeResultCodeAccountMergeMalformed     AccountMergeResultCode = -1
	AccountMergeResultCodeAccountMergeNoAccount     AccountMergeResultCode = -2
	AccountMergeResultCodeAccountMergeImmutableSet  AccountMergeResultCode = -3
	AccountMergeResultCodeAccountMergeHasSubEntries AccountMergeResultCode = -4
	AccountMergeResultCodeAccountMergeSeqnumTooFar  AccountMergeResultCode = -5
	AccountMergeResultCodeAccountMergeDestFull      AccountMergeResultCode = -6
	AccountMergeResultCodeAccountMergeIsSponsor     AccountMergeResultCode = -7
)

var accountMergeResultCodeMap = map[int32]string{
	0:  "AccountMergeResultCodeAccountMergeSuccess",
	-1: "AccountMergeResultCodeAccountMergeMalformed",
	-2: "AccountMergeResultCodeAccountMergeNoAccount",
	-3: "AccountMergeResultCodeAccountMergeImmutableSet",
	-4: "AccountMergeResultCodeAccountMergeHasSubEntries",
	-5: "AccountMergeResultCodeAccountMergeSeqnumTooFar",
	-6: "AccountMergeResultCodeAccountMergeDestFull",
	-7: "AccountMergeResultCodeAccountMergeIsSponsor",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for AccountMergeResultCode
func (e AccountMergeResultCode) ValidEnum(v int32) bool {
	_, ok := accountMergeResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e AccountMergeResultCode) String() string {
	name, _ := accountMergeResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountMergeResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountMergeResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountMergeResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountMergeResultCode)(nil)
)

// AccountMergeResult is an XDR Union defines as:
//
//   union AccountMergeResult switch (AccountMergeResultCode code)
//    {
//    case ACCOUNT_MERGE_SUCCESS:
//        int64 sourceAccountBalance; // how much got transferred from source account
//    default:
//        void;
//    };
//
type AccountMergeResult struct {
	Code                 AccountMergeResultCode
	SourceAccountBalance *Int64
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountMergeResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AccountMergeResult
func (u AccountMergeResult) ArmForSwitch(sw int32) (string, bool) {
	switch AccountMergeResultCode(sw) {
	case AccountMergeResultCodeAccountMergeSuccess:
		return "SourceAccountBalance", true
	default:
		return "", true
	}
}

// NewAccountMergeResult creates a new  AccountMergeResult.
func NewAccountMergeResult(code AccountMergeResultCode, value interface{}) (result AccountMergeResult, err error) {
	result.Code = code
	switch AccountMergeResultCode(code) {
	case AccountMergeResultCodeAccountMergeSuccess:
		tv, ok := value.(Int64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Int64")
			return
		}
		result.SourceAccountBalance = &tv
	default:
		// void
	}
	return
}

// MustSourceAccountBalance retrieves the SourceAccountBalance value from the union,
// panicing if the value is not set.
func (u AccountMergeResult) MustSourceAccountBalance() Int64 {
	val, ok := u.GetSourceAccountBalance()

	if !ok {
		panic("arm SourceAccountBalance is not set")
	}

	return val
}

// GetSourceAccountBalance retrieves the SourceAccountBalance value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountMergeResult) GetSourceAccountBalance() (result Int64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "SourceAccountBalance" {
		result = *u.SourceAccountBalance
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountMergeResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountMergeResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountMergeResult)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountMergeResult)(nil)
)

// InflationResultCode is an XDR Enum defines as:
//
//   enum InflationResultCode
//    {
//        // codes considered as "success" for the operation
//        INFLATION_SUCCESS = 0,
//        // codes considered as "failure" for the operation
//        INFLATION_NOT_TIME = -1
//    };
//
type InflationResultCode int32

const (
	InflationResultCodeInflationSuccess InflationResultCode = 0
	InflationResultCodeInflationNotTime InflationResultCode = -1
)

var inflationResultCodeMap = map[int32]string{
	0:  "InflationResultCodeInflationSuccess",
	-1: "InflationResultCodeInflationNotTime",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for InflationResultCode
func (e InflationResultCode) ValidEnum(v int32) bool {
	_, ok := inflationResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e InflationResultCode) String() string {
	name, _ := inflationResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s InflationResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *InflationResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*InflationResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*InflationResultCode)(nil)
)

// InflationPayout is an XDR Struct defines as:
//
//   struct InflationPayout // or use PaymentResultAtom to limit types?
//    {
//        AccountID destination;
//        int64 amount;
//    };
//
type InflationPayout struct {
	Destination AccountId
	Amount      Int64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s InflationPayout) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *InflationPayout) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*InflationPayout)(nil)
	_ encoding.BinaryUnmarshaler = (*InflationPayout)(nil)
)

// InflationResult is an XDR Union defines as:
//
//   union InflationResult switch (InflationResultCode code)
//    {
//    case INFLATION_SUCCESS:
//        InflationPayout payouts<>;
//    default:
//        void;
//    };
//
type InflationResult struct {
	Code    InflationResultCode
	Payouts *[]InflationPayout
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u InflationResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of InflationResult
func (u InflationResult) ArmForSwitch(sw int32) (string, bool) {
	switch InflationResultCode(sw) {
	case InflationResultCodeInflationSuccess:
		return "Payouts", true
	default:
		return "", true
	}
}

// NewInflationResult creates a new  InflationResult.
func NewInflationResult(code InflationResultCode, value interface{}) (result InflationResult, err error) {
	result.Code = code
	switch InflationResultCode(code) {
	case InflationResultCodeInflationSuccess:
		tv, ok := value.([]InflationPayout)
		if !ok {
			err = fmt.Errorf("invalid value, must be []InflationPayout")
			return
		}
		result.Payouts = &tv
	default:
		// void
	}
	return
}

// MustPayouts retrieves the Payouts value from the union,
// panicing if the value is not set.
func (u InflationResult) MustPayouts() []InflationPayout {
	val, ok := u.GetPayouts()

	if !ok {
		panic("arm Payouts is not set")
	}

	return val
}

// GetPayouts retrieves the Payouts value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u InflationResult) GetPayouts() (result []InflationPayout, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Payouts" {
		result = *u.Payouts
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s InflationResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *InflationResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*InflationResult)(nil)
	_ encoding.BinaryUnmarshaler = (*InflationResult)(nil)
)

// ManageDataResultCode is an XDR Enum defines as:
//
//   enum ManageDataResultCode
//    {
//        // codes considered as "success" for the operation
//        MANAGE_DATA_SUCCESS = 0,
//        // codes considered as "failure" for the operation
//        MANAGE_DATA_NOT_SUPPORTED_YET =
//            -1, // The network hasn't moved to this protocol change yet
//        MANAGE_DATA_NAME_NOT_FOUND =
//            -2, // Trying to remove a Data Entry that isn't there
//        MANAGE_DATA_LOW_RESERVE = -3, // not enough funds to create a new Data Entry
//        MANAGE_DATA_INVALID_NAME = -4 // Name not a valid string
//    };
//
type ManageDataResultCode int32

const (
	ManageDataResultCodeManageDataSuccess         ManageDataResultCode = 0
	ManageDataResultCodeManageDataNotSupportedYet ManageDataResultCode = -1
	ManageDataResultCodeManageDataNameNotFound    ManageDataResultCode = -2
	ManageDataResultCodeManageDataLowReserve      ManageDataResultCode = -3
	ManageDataResultCodeManageDataInvalidName     ManageDataResultCode = -4
)

var manageDataResultCodeMap = map[int32]string{
	0:  "ManageDataResultCodeManageDataSuccess",
	-1: "ManageDataResultCodeManageDataNotSupportedYet",
	-2: "ManageDataResultCodeManageDataNameNotFound",
	-3: "ManageDataResultCodeManageDataLowReserve",
	-4: "ManageDataResultCodeManageDataInvalidName",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageDataResultCode
func (e ManageDataResultCode) ValidEnum(v int32) bool {
	_, ok := manageDataResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e ManageDataResultCode) String() string {
	name, _ := manageDataResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ManageDataResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ManageDataResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ManageDataResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*ManageDataResultCode)(nil)
)

// ManageDataResult is an XDR Union defines as:
//
//   union ManageDataResult switch (ManageDataResultCode code)
//    {
//    case MANAGE_DATA_SUCCESS:
//        void;
//    default:
//        void;
//    };
//
type ManageDataResult struct {
	Code ManageDataResultCode
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageDataResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageDataResult
func (u ManageDataResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageDataResultCode(sw) {
	case ManageDataResultCodeManageDataSuccess:
		return "", true
	default:
		return "", true
	}
}

// NewManageDataResult creates a new  ManageDataResult.
func NewManageDataResult(code ManageDataResultCode, value interface{}) (result ManageDataResult, err error) {
	result.Code = code
	switch ManageDataResultCode(code) {
	case ManageDataResultCodeManageDataSuccess:
		// void
	default:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ManageDataResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ManageDataResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ManageDataResult)(nil)
	_ encoding.BinaryUnmarshaler = (*ManageDataResult)(nil)
)

// BumpSequenceResultCode is an XDR Enum defines as:
//
//   enum BumpSequenceResultCode
//    {
//        // codes considered as "success" for the operation
//        BUMP_SEQUENCE_SUCCESS = 0,
//        // codes considered as "failure" for the operation
//        BUMP_SEQUENCE_BAD_SEQ = -1 // `bumpTo` is not within bounds
//    };
//
type BumpSequenceResultCode int32

const (
	BumpSequenceResultCodeBumpSequenceSuccess BumpSequenceResultCode = 0
	BumpSequenceResultCodeBumpSequenceBadSeq  BumpSequenceResultCode = -1
)

var bumpSequenceResultCodeMap = map[int32]string{
	0:  "BumpSequenceResultCodeBumpSequenceSuccess",
	-1: "BumpSequenceResultCodeBumpSequenceBadSeq",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for BumpSequenceResultCode
func (e BumpSequenceResultCode) ValidEnum(v int32) bool {
	_, ok := bumpSequenceResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e BumpSequenceResultCode) String() string {
	name, _ := bumpSequenceResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BumpSequenceResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BumpSequenceResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BumpSequenceResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*BumpSequenceResultCode)(nil)
)

// BumpSequenceResult is an XDR Union defines as:
//
//   union BumpSequenceResult switch (BumpSequenceResultCode code)
//    {
//    case BUMP_SEQUENCE_SUCCESS:
//        void;
//    default:
//        void;
//    };
//
type BumpSequenceResult struct {
	Code BumpSequenceResultCode
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BumpSequenceResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of BumpSequenceResult
func (u BumpSequenceResult) ArmForSwitch(sw int32) (string, bool) {
	switch BumpSequenceResultCode(sw) {
	case BumpSequenceResultCodeBumpSequenceSuccess:
		return "", true
	default:
		return "", true
	}
}

// NewBumpSequenceResult creates a new  BumpSequenceResult.
func NewBumpSequenceResult(code BumpSequenceResultCode, value interface{}) (result BumpSequenceResult, err error) {
	result.Code = code
	switch BumpSequenceResultCode(code) {
	case BumpSequenceResultCodeBumpSequenceSuccess:
		// void
	default:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BumpSequenceResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BumpSequenceResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BumpSequenceResult)(nil)
	_ encoding.BinaryUnmarshaler = (*BumpSequenceResult)(nil)
)

// CreateClaimableBalanceResultCode is an XDR Enum defines as:
//
//   enum CreateClaimableBalanceResultCode
//    {
//        CREATE_CLAIMABLE_BALANCE_SUCCESS = 0,
//        CREATE_CLAIMABLE_BALANCE_MALFORMED = -1,
//        CREATE_CLAIMABLE_BALANCE_LOW_RESERVE = -2,
//        CREATE_CLAIMABLE_BALANCE_NO_TRUST = -3,
//        CREATE_CLAIMABLE_BALANCE_NOT_AUTHORIZED = -4,
//        CREATE_CLAIMABLE_BALANCE_UNDERFUNDED = -5
//    };
//
type CreateClaimableBalanceResultCode int32

const (
	CreateClaimableBalanceResultCodeCreateClaimableBalanceSuccess       CreateClaimableBalanceResultCode = 0
	CreateClaimableBalanceResultCodeCreateClaimableBalanceMalformed     CreateClaimableBalanceResultCode = -1
	CreateClaimableBalanceResultCodeCreateClaimableBalanceLowReserve    CreateClaimableBalanceResultCode = -2
	CreateClaimableBalanceResultCodeCreateClaimableBalanceNoTrust       CreateClaimableBalanceResultCode = -3
	CreateClaimableBalanceResultCodeCreateClaimableBalanceNotAuthorized CreateClaimableBalanceResultCode = -4
	CreateClaimableBalanceResultCodeCreateClaimableBalanceUnderfunded   CreateClaimableBalanceResultCode = -5
)

var createClaimableBalanceResultCodeMap = map[int32]string{
	0:  "CreateClaimableBalanceResultCodeCreateClaimableBalanceSuccess",
	-1: "CreateClaimableBalanceResultCodeCreateClaimableBalanceMalformed",
	-2: "CreateClaimableBalanceResultCodeCreateClaimableBalanceLowReserve",
	-3: "CreateClaimableBalanceResultCodeCreateClaimableBalanceNoTrust",
	-4: "CreateClaimableBalanceResultCodeCreateClaimableBalanceNotAuthorized",
	-5: "CreateClaimableBalanceResultCodeCreateClaimableBalanceUnderfunded",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateClaimableBalanceResultCode
func (e CreateClaimableBalanceResultCode) ValidEnum(v int32) bool {
	_, ok := createClaimableBalanceResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e CreateClaimableBalanceResultCode) String() string {
	name, _ := createClaimableBalanceResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s CreateClaimableBalanceResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *CreateClaimableBalanceResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*CreateClaimableBalanceResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*CreateClaimableBalanceResultCode)(nil)
)

// CreateClaimableBalanceResult is an XDR Union defines as:
//
//   union CreateClaimableBalanceResult switch (CreateClaimableBalanceResultCode code)
//    {
//    case CREATE_CLAIMABLE_BALANCE_SUCCESS:
//        ClaimableBalanceID balanceID;
//    default:
//        void;
//    };
//
type CreateClaimableBalanceResult struct {
	Code      CreateClaimableBalanceResultCode
	BalanceId *ClaimableBalanceId
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateClaimableBalanceResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateClaimableBalanceResult
func (u CreateClaimableBalanceResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateClaimableBalanceResultCode(sw) {
	case CreateClaimableBalanceResultCodeCreateClaimableBalanceSuccess:
		return "BalanceId", true
	default:
		return "", true
	}
}

// NewCreateClaimableBalanceResult creates a new  CreateClaimableBalanceResult.
func NewCreateClaimableBalanceResult(code CreateClaimableBalanceResultCode, value interface{}) (result CreateClaimableBalanceResult, err error) {
	result.Code = code
	switch CreateClaimableBalanceResultCode(code) {
	case CreateClaimableBalanceResultCodeCreateClaimableBalanceSuccess:
		tv, ok := value.(ClaimableBalanceId)
		if !ok {
			err = fmt.Errorf("invalid value, must be ClaimableBalanceId")
			return
		}
		result.BalanceId = &tv
	default:
		// void
	}
	return
}

// MustBalanceId retrieves the BalanceId value from the union,
// panicing if the value is not set.
func (u CreateClaimableBalanceResult) MustBalanceId() ClaimableBalanceId {
	val, ok := u.GetBalanceId()

	if !ok {
		panic("arm BalanceId is not set")
	}

	return val
}

// GetBalanceId retrieves the BalanceId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateClaimableBalanceResult) GetBalanceId() (result ClaimableBalanceId, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "BalanceId" {
		result = *u.BalanceId
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s CreateClaimableBalanceResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *CreateClaimableBalanceResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*CreateClaimableBalanceResult)(nil)
	_ encoding.BinaryUnmarshaler = (*CreateClaimableBalanceResult)(nil)
)

// ClaimClaimableBalanceResultCode is an XDR Enum defines as:
//
//   enum ClaimClaimableBalanceResultCode
//    {
//        CLAIM_CLAIMABLE_BALANCE_SUCCESS = 0,
//        CLAIM_CLAIMABLE_BALANCE_DOES_NOT_EXIST = -1,
//        CLAIM_CLAIMABLE_BALANCE_CANNOT_CLAIM = -2,
//        CLAIM_CLAIMABLE_BALANCE_LINE_FULL = -3,
//        CLAIM_CLAIMABLE_BALANCE_NO_TRUST = -4,
//        CLAIM_CLAIMABLE_BALANCE_NOT_AUTHORIZED = -5
//
//    };
//
type ClaimClaimableBalanceResultCode int32

const (
	ClaimClaimableBalanceResultCodeClaimClaimableBalanceSuccess       ClaimClaimableBalanceResultCode = 0
	ClaimClaimableBalanceResultCodeClaimClaimableBalanceDoesNotExist  ClaimClaimableBalanceResultCode = -1
	ClaimClaimableBalanceResultCodeClaimClaimableBalanceCannotClaim   ClaimClaimableBalanceResultCode = -2
	ClaimClaimableBalanceResultCodeClaimClaimableBalanceLineFull      ClaimClaimableBalanceResultCode = -3
	ClaimClaimableBalanceResultCodeClaimClaimableBalanceNoTrust       ClaimClaimableBalanceResultCode = -4
	ClaimClaimableBalanceResultCodeClaimClaimableBalanceNotAuthorized ClaimClaimableBalanceResultCode = -5
)

var claimClaimableBalanceResultCodeMap = map[int32]string{
	0:  "ClaimClaimableBalanceResultCodeClaimClaimableBalanceSuccess",
	-1: "ClaimClaimableBalanceResultCodeClaimClaimableBalanceDoesNotExist",
	-2: "ClaimClaimableBalanceResultCodeClaimClaimableBalanceCannotClaim",
	-3: "ClaimClaimableBalanceResultCodeClaimClaimableBalanceLineFull",
	-4: "ClaimClaimableBalanceResultCodeClaimClaimableBalanceNoTrust",
	-5: "ClaimClaimableBalanceResultCodeClaimClaimableBalanceNotAuthorized",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ClaimClaimableBalanceResultCode
func (e ClaimClaimableBalanceResultCode) ValidEnum(v int32) bool {
	_, ok := claimClaimableBalanceResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e ClaimClaimableBalanceResultCode) String() string {
	name, _ := claimClaimableBalanceResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimClaimableBalanceResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimClaimableBalanceResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimClaimableBalanceResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimClaimableBalanceResultCode)(nil)
)

// ClaimClaimableBalanceResult is an XDR Union defines as:
//
//   union ClaimClaimableBalanceResult switch (ClaimClaimableBalanceResultCode code)
//    {
//    case CLAIM_CLAIMABLE_BALANCE_SUCCESS:
//        void;
//    default:
//        void;
//    };
//
type ClaimClaimableBalanceResult struct {
	Code ClaimClaimableBalanceResultCode
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ClaimClaimableBalanceResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ClaimClaimableBalanceResult
func (u ClaimClaimableBalanceResult) ArmForSwitch(sw int32) (string, bool) {
	switch ClaimClaimableBalanceResultCode(sw) {
	case ClaimClaimableBalanceResultCodeClaimClaimableBalanceSuccess:
		return "", true
	default:
		return "", true
	}
}

// NewClaimClaimableBalanceResult creates a new  ClaimClaimableBalanceResult.
func NewClaimClaimableBalanceResult(code ClaimClaimableBalanceResultCode, value interface{}) (result ClaimClaimableBalanceResult, err error) {
	result.Code = code
	switch ClaimClaimableBalanceResultCode(code) {
	case ClaimClaimableBalanceResultCodeClaimClaimableBalanceSuccess:
		// void
	default:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClaimClaimableBalanceResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClaimClaimableBalanceResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClaimClaimableBalanceResult)(nil)
	_ encoding.BinaryUnmarshaler = (*ClaimClaimableBalanceResult)(nil)
)

// BeginSponsoringFutureReservesResultCode is an XDR Enum defines as:
//
//   enum BeginSponsoringFutureReservesResultCode
//    {
//        // codes considered as "success" for the operation
//        BEGIN_SPONSORING_FUTURE_RESERVES_SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        BEGIN_SPONSORING_FUTURE_RESERVES_MALFORMED = -1,
//        BEGIN_SPONSORING_FUTURE_RESERVES_ALREADY_SPONSORED = -2,
//        BEGIN_SPONSORING_FUTURE_RESERVES_RECURSIVE = -3
//    };
//
type BeginSponsoringFutureReservesResultCode int32

const (
	BeginSponsoringFutureReservesResultCodeBeginSponsoringFutureReservesSuccess          BeginSponsoringFutureReservesResultCode = 0
	BeginSponsoringFutureReservesResultCodeBeginSponsoringFutureReservesMalformed        BeginSponsoringFutureReservesResultCode = -1
	BeginSponsoringFutureReservesResultCodeBeginSponsoringFutureReservesAlreadySponsored BeginSponsoringFutureReservesResultCode = -2
	BeginSponsoringFutureReservesResultCodeBeginSponsoringFutureReservesRecursive        BeginSponsoringFutureReservesResultCode = -3
)

var beginSponsoringFutureReservesResultCodeMap = map[int32]string{
	0:  "BeginSponsoringFutureReservesResultCodeBeginSponsoringFutureReservesSuccess",
	-1: "BeginSponsoringFutureReservesResultCodeBeginSponsoringFutureReservesMalformed",
	-2: "BeginSponsoringFutureReservesResultCodeBeginSponsoringFutureReservesAlreadySponsored",
	-3: "BeginSponsoringFutureReservesResultCodeBeginSponsoringFutureReservesRecursive",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for BeginSponsoringFutureReservesResultCode
func (e BeginSponsoringFutureReservesResultCode) ValidEnum(v int32) bool {
	_, ok := beginSponsoringFutureReservesResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e BeginSponsoringFutureReservesResultCode) String() string {
	name, _ := beginSponsoringFutureReservesResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BeginSponsoringFutureReservesResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BeginSponsoringFutureReservesResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BeginSponsoringFutureReservesResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*BeginSponsoringFutureReservesResultCode)(nil)
)

// BeginSponsoringFutureReservesResult is an XDR Union defines as:
//
//   union BeginSponsoringFutureReservesResult switch (BeginSponsoringFutureReservesResultCode code)
//    {
//    case BEGIN_SPONSORING_FUTURE_RESERVES_SUCCESS:
//        void;
//    default:
//        void;
//    };
//
type BeginSponsoringFutureReservesResult struct {
	Code BeginSponsoringFutureReservesResultCode
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BeginSponsoringFutureReservesResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of BeginSponsoringFutureReservesResult
func (u BeginSponsoringFutureReservesResult) ArmForSwitch(sw int32) (string, bool) {
	switch BeginSponsoringFutureReservesResultCode(sw) {
	case BeginSponsoringFutureReservesResultCodeBeginSponsoringFutureReservesSuccess:
		return "", true
	default:
		return "", true
	}
}

// NewBeginSponsoringFutureReservesResult creates a new  BeginSponsoringFutureReservesResult.
func NewBeginSponsoringFutureReservesResult(code BeginSponsoringFutureReservesResultCode, value interface{}) (result BeginSponsoringFutureReservesResult, err error) {
	result.Code = code
	switch BeginSponsoringFutureReservesResultCode(code) {
	case BeginSponsoringFutureReservesResultCodeBeginSponsoringFutureReservesSuccess:
		// void
	default:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BeginSponsoringFutureReservesResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BeginSponsoringFutureReservesResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BeginSponsoringFutureReservesResult)(nil)
	_ encoding.BinaryUnmarshaler = (*BeginSponsoringFutureReservesResult)(nil)
)

// EndSponsoringFutureReservesResultCode is an XDR Enum defines as:
//
//   enum EndSponsoringFutureReservesResultCode
//    {
//        // codes considered as "success" for the operation
//        END_SPONSORING_FUTURE_RESERVES_SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        END_SPONSORING_FUTURE_RESERVES_NOT_SPONSORED = -1
//    };
//
type EndSponsoringFutureReservesResultCode int32

const (
	EndSponsoringFutureReservesResultCodeEndSponsoringFutureReservesSuccess      EndSponsoringFutureReservesResultCode = 0
	EndSponsoringFutureReservesResultCodeEndSponsoringFutureReservesNotSponsored EndSponsoringFutureReservesResultCode = -1
)

var endSponsoringFutureReservesResultCodeMap = map[int32]string{
	0:  "EndSponsoringFutureReservesResultCodeEndSponsoringFutureReservesSuccess",
	-1: "EndSponsoringFutureReservesResultCodeEndSponsoringFutureReservesNotSponsored",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for EndSponsoringFutureReservesResultCode
func (e EndSponsoringFutureReservesResultCode) ValidEnum(v int32) bool {
	_, ok := endSponsoringFutureReservesResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e EndSponsoringFutureReservesResultCode) String() string {
	name, _ := endSponsoringFutureReservesResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s EndSponsoringFutureReservesResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *EndSponsoringFutureReservesResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*EndSponsoringFutureReservesResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*EndSponsoringFutureReservesResultCode)(nil)
)

// EndSponsoringFutureReservesResult is an XDR Union defines as:
//
//   union EndSponsoringFutureReservesResult switch (EndSponsoringFutureReservesResultCode code)
//    {
//    case END_SPONSORING_FUTURE_RESERVES_SUCCESS:
//        void;
//    default:
//        void;
//    };
//
type EndSponsoringFutureReservesResult struct {
	Code EndSponsoringFutureReservesResultCode
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u EndSponsoringFutureReservesResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of EndSponsoringFutureReservesResult
func (u EndSponsoringFutureReservesResult) ArmForSwitch(sw int32) (string, bool) {
	switch EndSponsoringFutureReservesResultCode(sw) {
	case EndSponsoringFutureReservesResultCodeEndSponsoringFutureReservesSuccess:
		return "", true
	default:
		return "", true
	}
}

// NewEndSponsoringFutureReservesResult creates a new  EndSponsoringFutureReservesResult.
func NewEndSponsoringFutureReservesResult(code EndSponsoringFutureReservesResultCode, value interface{}) (result EndSponsoringFutureReservesResult, err error) {
	result.Code = code
	switch EndSponsoringFutureReservesResultCode(code) {
	case EndSponsoringFutureReservesResultCodeEndSponsoringFutureReservesSuccess:
		// void
	default:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s EndSponsoringFutureReservesResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *EndSponsoringFutureReservesResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*EndSponsoringFutureReservesResult)(nil)
	_ encoding.BinaryUnmarshaler = (*EndSponsoringFutureReservesResult)(nil)
)

// RevokeSponsorshipResultCode is an XDR Enum defines as:
//
//   enum RevokeSponsorshipResultCode
//    {
//        // codes considered as "success" for the operation
//        REVOKE_SPONSORSHIP_SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        REVOKE_SPONSORSHIP_DOES_NOT_EXIST = -1,
//        REVOKE_SPONSORSHIP_NOT_SPONSOR = -2,
//        REVOKE_SPONSORSHIP_LOW_RESERVE = -3,
//        REVOKE_SPONSORSHIP_ONLY_TRANSFERABLE = -4,
//        REVOKE_SPONSORSHIP_MALFORMED = -5
//    };
//
type RevokeSponsorshipResultCode int32

const (
	RevokeSponsorshipResultCodeRevokeSponsorshipSuccess          RevokeSponsorshipResultCode = 0
	RevokeSponsorshipResultCodeRevokeSponsorshipDoesNotExist     RevokeSponsorshipResultCode = -1
	RevokeSponsorshipResultCodeRevokeSponsorshipNotSponsor       RevokeSponsorshipResultCode = -2
	RevokeSponsorshipResultCodeRevokeSponsorshipLowReserve       RevokeSponsorshipResultCode = -3
	RevokeSponsorshipResultCodeRevokeSponsorshipOnlyTransferable RevokeSponsorshipResultCode = -4
	RevokeSponsorshipResultCodeRevokeSponsorshipMalformed        RevokeSponsorshipResultCode = -5
)

var revokeSponsorshipResultCodeMap = map[int32]string{
	0:  "RevokeSponsorshipResultCodeRevokeSponsorshipSuccess",
	-1: "RevokeSponsorshipResultCodeRevokeSponsorshipDoesNotExist",
	-2: "RevokeSponsorshipResultCodeRevokeSponsorshipNotSponsor",
	-3: "RevokeSponsorshipResultCodeRevokeSponsorshipLowReserve",
	-4: "RevokeSponsorshipResultCodeRevokeSponsorshipOnlyTransferable",
	-5: "RevokeSponsorshipResultCodeRevokeSponsorshipMalformed",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for RevokeSponsorshipResultCode
func (e RevokeSponsorshipResultCode) ValidEnum(v int32) bool {
	_, ok := revokeSponsorshipResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e RevokeSponsorshipResultCode) String() string {
	name, _ := revokeSponsorshipResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s RevokeSponsorshipResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *RevokeSponsorshipResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*RevokeSponsorshipResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*RevokeSponsorshipResultCode)(nil)
)

// RevokeSponsorshipResult is an XDR Union defines as:
//
//   union RevokeSponsorshipResult switch (RevokeSponsorshipResultCode code)
//    {
//    case REVOKE_SPONSORSHIP_SUCCESS:
//        void;
//    default:
//        void;
//    };
//
type RevokeSponsorshipResult struct {
	Code RevokeSponsorshipResultCode
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RevokeSponsorshipResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RevokeSponsorshipResult
func (u RevokeSponsorshipResult) ArmForSwitch(sw int32) (string, bool) {
	switch RevokeSponsorshipResultCode(sw) {
	case RevokeSponsorshipResultCodeRevokeSponsorshipSuccess:
		return "", true
	default:
		return "", true
	}
}

// NewRevokeSponsorshipResult creates a new  RevokeSponsorshipResult.
func NewRevokeSponsorshipResult(code RevokeSponsorshipResultCode, value interface{}) (result RevokeSponsorshipResult, err error) {
	result.Code = code
	switch RevokeSponsorshipResultCode(code) {
	case RevokeSponsorshipResultCodeRevokeSponsorshipSuccess:
		// void
	default:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s RevokeSponsorshipResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *RevokeSponsorshipResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*RevokeSponsorshipResult)(nil)
	_ encoding.BinaryUnmarshaler = (*RevokeSponsorshipResult)(nil)
)

// ClawbackResultCode is an XDR Enum defines as:
//
//   enum ClawbackResultCode
//    {
//        // codes considered as "success" for the operation
//        CLAWBACK_SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        CLAWBACK_MALFORMED = -1,
//        CLAWBACK_NOT_CLAWBACK_ENABLED = -2,
//        CLAWBACK_NO_TRUST = -3,
//        CLAWBACK_UNDERFUNDED = -4
//    };
//
type ClawbackResultCode int32

const (
	ClawbackResultCodeClawbackSuccess            ClawbackResultCode = 0
	ClawbackResultCodeClawbackMalformed          ClawbackResultCode = -1
	ClawbackResultCodeClawbackNotClawbackEnabled ClawbackResultCode = -2
	ClawbackResultCodeClawbackNoTrust            ClawbackResultCode = -3
	ClawbackResultCodeClawbackUnderfunded        ClawbackResultCode = -4
)

var clawbackResultCodeMap = map[int32]string{
	0:  "ClawbackResultCodeClawbackSuccess",
	-1: "ClawbackResultCodeClawbackMalformed",
	-2: "ClawbackResultCodeClawbackNotClawbackEnabled",
	-3: "ClawbackResultCodeClawbackNoTrust",
	-4: "ClawbackResultCodeClawbackUnderfunded",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ClawbackResultCode
func (e ClawbackResultCode) ValidEnum(v int32) bool {
	_, ok := clawbackResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e ClawbackResultCode) String() string {
	name, _ := clawbackResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClawbackResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClawbackResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClawbackResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*ClawbackResultCode)(nil)
)

// ClawbackResult is an XDR Union defines as:
//
//   union ClawbackResult switch (ClawbackResultCode code)
//    {
//    case CLAWBACK_SUCCESS:
//        void;
//    default:
//        void;
//    };
//
type ClawbackResult struct {
	Code ClawbackResultCode
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ClawbackResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ClawbackResult
func (u ClawbackResult) ArmForSwitch(sw int32) (string, bool) {
	switch ClawbackResultCode(sw) {
	case ClawbackResultCodeClawbackSuccess:
		return "", true
	default:
		return "", true
	}
}

// NewClawbackResult creates a new  ClawbackResult.
func NewClawbackResult(code ClawbackResultCode, value interface{}) (result ClawbackResult, err error) {
	result.Code = code
	switch ClawbackResultCode(code) {
	case ClawbackResultCodeClawbackSuccess:
		// void
	default:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClawbackResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClawbackResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClawbackResult)(nil)
	_ encoding.BinaryUnmarshaler = (*ClawbackResult)(nil)
)

// ClawbackClaimableBalanceResultCode is an XDR Enum defines as:
//
//   enum ClawbackClaimableBalanceResultCode
//    {
//        // codes considered as "success" for the operation
//        CLAWBACK_CLAIMABLE_BALANCE_SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        CLAWBACK_CLAIMABLE_BALANCE_DOES_NOT_EXIST = -1,
//        CLAWBACK_CLAIMABLE_BALANCE_NOT_ISSUER = -2,
//        CLAWBACK_CLAIMABLE_BALANCE_NOT_CLAWBACK_ENABLED = -3
//    };
//
type ClawbackClaimableBalanceResultCode int32

const (
	ClawbackClaimableBalanceResultCodeClawbackClaimableBalanceSuccess            ClawbackClaimableBalanceResultCode = 0
	ClawbackClaimableBalanceResultCodeClawbackClaimableBalanceDoesNotExist       ClawbackClaimableBalanceResultCode = -1
	ClawbackClaimableBalanceResultCodeClawbackClaimableBalanceNotIssuer          ClawbackClaimableBalanceResultCode = -2
	ClawbackClaimableBalanceResultCodeClawbackClaimableBalanceNotClawbackEnabled ClawbackClaimableBalanceResultCode = -3
)

var clawbackClaimableBalanceResultCodeMap = map[int32]string{
	0:  "ClawbackClaimableBalanceResultCodeClawbackClaimableBalanceSuccess",
	-1: "ClawbackClaimableBalanceResultCodeClawbackClaimableBalanceDoesNotExist",
	-2: "ClawbackClaimableBalanceResultCodeClawbackClaimableBalanceNotIssuer",
	-3: "ClawbackClaimableBalanceResultCodeClawbackClaimableBalanceNotClawbackEnabled",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ClawbackClaimableBalanceResultCode
func (e ClawbackClaimableBalanceResultCode) ValidEnum(v int32) bool {
	_, ok := clawbackClaimableBalanceResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e ClawbackClaimableBalanceResultCode) String() string {
	name, _ := clawbackClaimableBalanceResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClawbackClaimableBalanceResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClawbackClaimableBalanceResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClawbackClaimableBalanceResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*ClawbackClaimableBalanceResultCode)(nil)
)

// ClawbackClaimableBalanceResult is an XDR Union defines as:
//
//   union ClawbackClaimableBalanceResult switch (ClawbackClaimableBalanceResultCode code)
//    {
//    case CLAWBACK_CLAIMABLE_BALANCE_SUCCESS:
//        void;
//    default:
//        void;
//    };
//
type ClawbackClaimableBalanceResult struct {
	Code ClawbackClaimableBalanceResultCode
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ClawbackClaimableBalanceResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ClawbackClaimableBalanceResult
func (u ClawbackClaimableBalanceResult) ArmForSwitch(sw int32) (string, bool) {
	switch ClawbackClaimableBalanceResultCode(sw) {
	case ClawbackClaimableBalanceResultCodeClawbackClaimableBalanceSuccess:
		return "", true
	default:
		return "", true
	}
}

// NewClawbackClaimableBalanceResult creates a new  ClawbackClaimableBalanceResult.
func NewClawbackClaimableBalanceResult(code ClawbackClaimableBalanceResultCode, value interface{}) (result ClawbackClaimableBalanceResult, err error) {
	result.Code = code
	switch ClawbackClaimableBalanceResultCode(code) {
	case ClawbackClaimableBalanceResultCodeClawbackClaimableBalanceSuccess:
		// void
	default:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ClawbackClaimableBalanceResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ClawbackClaimableBalanceResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ClawbackClaimableBalanceResult)(nil)
	_ encoding.BinaryUnmarshaler = (*ClawbackClaimableBalanceResult)(nil)
)

// SetTrustLineFlagsResultCode is an XDR Enum defines as:
//
//   enum SetTrustLineFlagsResultCode
//    {
//        // codes considered as "success" for the operation
//        SET_TRUST_LINE_FLAGS_SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        SET_TRUST_LINE_FLAGS_MALFORMED = -1,
//        SET_TRUST_LINE_FLAGS_NO_TRUST_LINE = -2,
//        SET_TRUST_LINE_FLAGS_CANT_REVOKE = -3,
//        SET_TRUST_LINE_FLAGS_INVALID_STATE = -4
//    };
//
type SetTrustLineFlagsResultCode int32

const (
	SetTrustLineFlagsResultCodeSetTrustLineFlagsSuccess      SetTrustLineFlagsResultCode = 0
	SetTrustLineFlagsResultCodeSetTrustLineFlagsMalformed    SetTrustLineFlagsResultCode = -1
	SetTrustLineFlagsResultCodeSetTrustLineFlagsNoTrustLine  SetTrustLineFlagsResultCode = -2
	SetTrustLineFlagsResultCodeSetTrustLineFlagsCantRevoke   SetTrustLineFlagsResultCode = -3
	SetTrustLineFlagsResultCodeSetTrustLineFlagsInvalidState SetTrustLineFlagsResultCode = -4
)

var setTrustLineFlagsResultCodeMap = map[int32]string{
	0:  "SetTrustLineFlagsResultCodeSetTrustLineFlagsSuccess",
	-1: "SetTrustLineFlagsResultCodeSetTrustLineFlagsMalformed",
	-2: "SetTrustLineFlagsResultCodeSetTrustLineFlagsNoTrustLine",
	-3: "SetTrustLineFlagsResultCodeSetTrustLineFlagsCantRevoke",
	-4: "SetTrustLineFlagsResultCodeSetTrustLineFlagsInvalidState",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for SetTrustLineFlagsResultCode
func (e SetTrustLineFlagsResultCode) ValidEnum(v int32) bool {
	_, ok := setTrustLineFlagsResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e SetTrustLineFlagsResultCode) String() string {
	name, _ := setTrustLineFlagsResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SetTrustLineFlagsResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SetTrustLineFlagsResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SetTrustLineFlagsResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*SetTrustLineFlagsResultCode)(nil)
)

// SetTrustLineFlagsResult is an XDR Union defines as:
//
//   union SetTrustLineFlagsResult switch (SetTrustLineFlagsResultCode code)
//    {
//    case SET_TRUST_LINE_FLAGS_SUCCESS:
//        void;
//    default:
//        void;
//    };
//
type SetTrustLineFlagsResult struct {
	Code SetTrustLineFlagsResultCode
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetTrustLineFlagsResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetTrustLineFlagsResult
func (u SetTrustLineFlagsResult) ArmForSwitch(sw int32) (string, bool) {
	switch SetTrustLineFlagsResultCode(sw) {
	case SetTrustLineFlagsResultCodeSetTrustLineFlagsSuccess:
		return "", true
	default:
		return "", true
	}
}

// NewSetTrustLineFlagsResult creates a new  SetTrustLineFlagsResult.
func NewSetTrustLineFlagsResult(code SetTrustLineFlagsResultCode, value interface{}) (result SetTrustLineFlagsResult, err error) {
	result.Code = code
	switch SetTrustLineFlagsResultCode(code) {
	case SetTrustLineFlagsResultCodeSetTrustLineFlagsSuccess:
		// void
	default:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SetTrustLineFlagsResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SetTrustLineFlagsResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SetTrustLineFlagsResult)(nil)
	_ encoding.BinaryUnmarshaler = (*SetTrustLineFlagsResult)(nil)
)

// OperationResultCode is an XDR Enum defines as:
//
//   enum OperationResultCode
//    {
//        opINNER = 0, // inner object result is valid
//
//        opBAD_AUTH = -1,            // too few valid signatures / wrong network
//        opNO_ACCOUNT = -2,          // source account was not found
//        opNOT_SUPPORTED = -3,       // operation not supported at this time
//        opTOO_MANY_SUBENTRIES = -4, // max number of subentries already reached
//        opEXCEEDED_WORK_LIMIT = -5, // operation did too much work
//        opTOO_MANY_SPONSORING = -6  // account is sponsoring too many entries
//    };
//
type OperationResultCode int32

const (
	OperationResultCodeOpInner             OperationResultCode = 0
	OperationResultCodeOpBadAuth           OperationResultCode = -1
	OperationResultCodeOpNoAccount         OperationResultCode = -2
	OperationResultCodeOpNotSupported      OperationResultCode = -3
	OperationResultCodeOpTooManySubentries OperationResultCode = -4
	OperationResultCodeOpExceededWorkLimit OperationResultCode = -5
	OperationResultCodeOpTooManySponsoring OperationResultCode = -6
)

var operationResultCodeMap = map[int32]string{
	0:  "OperationResultCodeOpInner",
	-1: "OperationResultCodeOpBadAuth",
	-2: "OperationResultCodeOpNoAccount",
	-3: "OperationResultCodeOpNotSupported",
	-4: "OperationResultCodeOpTooManySubentries",
	-5: "OperationResultCodeOpExceededWorkLimit",
	-6: "OperationResultCodeOpTooManySponsoring",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for OperationResultCode
func (e OperationResultCode) ValidEnum(v int32) bool {
	_, ok := operationResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e OperationResultCode) String() string {
	name, _ := operationResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s OperationResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *OperationResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*OperationResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*OperationResultCode)(nil)
)

// OperationResultTr is an XDR NestedUnion defines as:
//
//   union switch (OperationType type)
//        {
//        case CREATE_ACCOUNT:
//            CreateAccountResult createAccountResult;
//        case PAYMENT:
//            PaymentResult paymentResult;
//        case PATH_PAYMENT_STRICT_RECEIVE:
//            PathPaymentStrictReceiveResult pathPaymentStrictReceiveResult;
//        case MANAGE_SELL_OFFER:
//            ManageSellOfferResult manageSellOfferResult;
//        case CREATE_PASSIVE_SELL_OFFER:
//            ManageSellOfferResult createPassiveSellOfferResult;
//        case SET_OPTIONS:
//            SetOptionsResult setOptionsResult;
//        case CHANGE_TRUST:
//            ChangeTrustResult changeTrustResult;
//        case ALLOW_TRUST:
//            AllowTrustResult allowTrustResult;
//        case ACCOUNT_MERGE:
//            AccountMergeResult accountMergeResult;
//        case INFLATION:
//            InflationResult inflationResult;
//        case MANAGE_DATA:
//            ManageDataResult manageDataResult;
//        case BUMP_SEQUENCE:
//            BumpSequenceResult bumpSeqResult;
//        case MANAGE_BUY_OFFER:
//            ManageBuyOfferResult manageBuyOfferResult;
//        case PATH_PAYMENT_STRICT_SEND:
//            PathPaymentStrictSendResult pathPaymentStrictSendResult;
//        case CREATE_CLAIMABLE_BALANCE:
//            CreateClaimableBalanceResult createClaimableBalanceResult;
//        case CLAIM_CLAIMABLE_BALANCE:
//            ClaimClaimableBalanceResult claimClaimableBalanceResult;
//        case BEGIN_SPONSORING_FUTURE_RESERVES:
//            BeginSponsoringFutureReservesResult beginSponsoringFutureReservesResult;
//        case END_SPONSORING_FUTURE_RESERVES:
//            EndSponsoringFutureReservesResult endSponsoringFutureReservesResult;
//        case REVOKE_SPONSORSHIP:
//            RevokeSponsorshipResult revokeSponsorshipResult;
//        case CLAWBACK:
//            ClawbackResult clawbackResult;
//        case CLAWBACK_CLAIMABLE_BALANCE:
//            ClawbackClaimableBalanceResult clawbackClaimableBalanceResult;
//        case SET_TRUST_LINE_FLAGS:
//            SetTrustLineFlagsResult setTrustLineFlagsResult;
//        }
//
type OperationResultTr struct {
	Type                                OperationType
	CreateAccountResult                 *CreateAccountResult
	PaymentResult                       *PaymentResult
	PathPaymentStrictReceiveResult      *PathPaymentStrictReceiveResult
	ManageSellOfferResult               *ManageSellOfferResult
	CreatePassiveSellOfferResult        *ManageSellOfferResult
	SetOptionsResult                    *SetOptionsResult
	ChangeTrustResult                   *ChangeTrustResult
	AllowTrustResult                    *AllowTrustResult
	AccountMergeResult                  *AccountMergeResult
	InflationResult                     *InflationResult
	ManageDataResult                    *ManageDataResult
	BumpSeqResult                       *BumpSequenceResult
	ManageBuyOfferResult                *ManageBuyOfferResult
	PathPaymentStrictSendResult         *PathPaymentStrictSendResult
	CreateClaimableBalanceResult        *CreateClaimableBalanceResult
	ClaimClaimableBalanceResult         *ClaimClaimableBalanceResult
	BeginSponsoringFutureReservesResult *BeginSponsoringFutureReservesResult
	EndSponsoringFutureReservesResult   *EndSponsoringFutureReservesResult
	RevokeSponsorshipResult             *RevokeSponsorshipResult
	ClawbackResult                      *ClawbackResult
	ClawbackClaimableBalanceResult      *ClawbackClaimableBalanceResult
	SetTrustLineFlagsResult             *SetTrustLineFlagsResult
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OperationResultTr) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OperationResultTr
func (u OperationResultTr) ArmForSwitch(sw int32) (string, bool) {
	switch OperationType(sw) {
	case OperationTypeCreateAccount:
		return "CreateAccountResult", true
	case OperationTypePayment:
		return "PaymentResult", true
	case OperationTypePathPaymentStrictReceive:
		return "PathPaymentStrictReceiveResult", true
	case OperationTypeManageSellOffer:
		return "ManageSellOfferResult", true
	case OperationTypeCreatePassiveSellOffer:
		return "CreatePassiveSellOfferResult", true
	case OperationTypeSetOptions:
		return "SetOptionsResult", true
	case OperationTypeChangeTrust:
		return "ChangeTrustResult", true
	case OperationTypeAllowTrust:
		return "AllowTrustResult", true
	case OperationTypeAccountMerge:
		return "AccountMergeResult", true
	case OperationTypeInflation:
		return "InflationResult", true
	case OperationTypeManageData:
		return "ManageDataResult", true
	case OperationTypeBumpSequence:
		return "BumpSeqResult", true
	case OperationTypeManageBuyOffer:
		return "ManageBuyOfferResult", true
	case OperationTypePathPaymentStrictSend:
		return "PathPaymentStrictSendResult", true
	case OperationTypeCreateClaimableBalance:
		return "CreateClaimableBalanceResult", true
	case OperationTypeClaimClaimableBalance:
		return "ClaimClaimableBalanceResult", true
	case OperationTypeBeginSponsoringFutureReserves:
		return "BeginSponsoringFutureReservesResult", true
	case OperationTypeEndSponsoringFutureReserves:
		return "EndSponsoringFutureReservesResult", true
	case OperationTypeRevokeSponsorship:
		return "RevokeSponsorshipResult", true
	case OperationTypeClawback:
		return "ClawbackResult", true
	case OperationTypeClawbackClaimableBalance:
		return "ClawbackClaimableBalanceResult", true
	case OperationTypeSetTrustLineFlags:
		return "SetTrustLineFlagsResult", true
	}
	return "-", false
}

// NewOperationResultTr creates a new  OperationResultTr.
func NewOperationResultTr(aType OperationType, value interface{}) (result OperationResultTr, err error) {
	result.Type = aType
	switch OperationType(aType) {
	case OperationTypeCreateAccount:
		tv, ok := value.(CreateAccountResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAccountResult")
			return
		}
		result.CreateAccountResult = &tv
	case OperationTypePayment:
		tv, ok := value.(PaymentResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentResult")
			return
		}
		result.PaymentResult = &tv
	case OperationTypePathPaymentStrictReceive:
		tv, ok := value.(PathPaymentStrictReceiveResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be PathPaymentStrictReceiveResult")
			return
		}
		result.PathPaymentStrictReceiveResult = &tv
	case OperationTypeManageSellOffer:
		tv, ok := value.(ManageSellOfferResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageSellOfferResult")
			return
		}
		result.ManageSellOfferResult = &tv
	case OperationTypeCreatePassiveSellOffer:
		tv, ok := value.(ManageSellOfferResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageSellOfferResult")
			return
		}
		result.CreatePassiveSellOfferResult = &tv
	case OperationTypeSetOptions:
		tv, ok := value.(SetOptionsResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetOptionsResult")
			return
		}
		result.SetOptionsResult = &tv
	case OperationTypeChangeTrust:
		tv, ok := value.(ChangeTrustResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ChangeTrustResult")
			return
		}
		result.ChangeTrustResult = &tv
	case OperationTypeAllowTrust:
		tv, ok := value.(AllowTrustResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be AllowTrustResult")
			return
		}
		result.AllowTrustResult = &tv
	case OperationTypeAccountMerge:
		tv, ok := value.(AccountMergeResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountMergeResult")
			return
		}
		result.AccountMergeResult = &tv
	case OperationTypeInflation:
		tv, ok := value.(InflationResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be InflationResult")
			return
		}
		result.InflationResult = &tv
	case OperationTypeManageData:
		tv, ok := value.(ManageDataResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageDataResult")
			return
		}
		result.ManageDataResult = &tv
	case OperationTypeBumpSequence:
		tv, ok := value.(BumpSequenceResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be BumpSequenceResult")
			return
		}
		result.BumpSeqResult = &tv
	case OperationTypeManageBuyOffer:
		tv, ok := value.(ManageBuyOfferResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageBuyOfferResult")
			return
		}
		result.ManageBuyOfferResult = &tv
	case OperationTypePathPaymentStrictSend:
		tv, ok := value.(PathPaymentStrictSendResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be PathPaymentStrictSendResult")
			return
		}
		result.PathPaymentStrictSendResult = &tv
	case OperationTypeCreateClaimableBalance:
		tv, ok := value.(CreateClaimableBalanceResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateClaimableBalanceResult")
			return
		}
		result.CreateClaimableBalanceResult = &tv
	case OperationTypeClaimClaimableBalance:
		tv, ok := value.(ClaimClaimableBalanceResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ClaimClaimableBalanceResult")
			return
		}
		result.ClaimClaimableBalanceResult = &tv
	case OperationTypeBeginSponsoringFutureReserves:
		tv, ok := value.(BeginSponsoringFutureReservesResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be BeginSponsoringFutureReservesResult")
			return
		}
		result.BeginSponsoringFutureReservesResult = &tv
	case OperationTypeEndSponsoringFutureReserves:
		tv, ok := value.(EndSponsoringFutureReservesResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be EndSponsoringFutureReservesResult")
			return
		}
		result.EndSponsoringFutureReservesResult = &tv
	case OperationTypeRevokeSponsorship:
		tv, ok := value.(RevokeSponsorshipResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be RevokeSponsorshipResult")
			return
		}
		result.RevokeSponsorshipResult = &tv
	case OperationTypeClawback:
		tv, ok := value.(ClawbackResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ClawbackResult")
			return
		}
		result.ClawbackResult = &tv
	case OperationTypeClawbackClaimableBalance:
		tv, ok := value.(ClawbackClaimableBalanceResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ClawbackClaimableBalanceResult")
			return
		}
		result.ClawbackClaimableBalanceResult = &tv
	case OperationTypeSetTrustLineFlags:
		tv, ok := value.(SetTrustLineFlagsResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetTrustLineFlagsResult")
			return
		}
		result.SetTrustLineFlagsResult = &tv
	}
	return
}

// MustCreateAccountResult retrieves the CreateAccountResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateAccountResult() CreateAccountResult {
	val, ok := u.GetCreateAccountResult()

	if !ok {
		panic("arm CreateAccountResult is not set")
	}

	return val
}

// GetCreateAccountResult retrieves the CreateAccountResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateAccountResult() (result CreateAccountResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateAccountResult" {
		result = *u.CreateAccountResult
		ok = true
	}

	return
}

// MustPaymentResult retrieves the PaymentResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustPaymentResult() PaymentResult {
	val, ok := u.GetPaymentResult()

	if !ok {
		panic("arm PaymentResult is not set")
	}

	return val
}

// GetPaymentResult retrieves the PaymentResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetPaymentResult() (result PaymentResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PaymentResult" {
		result = *u.PaymentResult
		ok = true
	}

	return
}

// MustPathPaymentStrictReceiveResult retrieves the PathPaymentStrictReceiveResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustPathPaymentStrictReceiveResult() PathPaymentStrictReceiveResult {
	val, ok := u.GetPathPaymentStrictReceiveResult()

	if !ok {
		panic("arm PathPaymentStrictReceiveResult is not set")
	}

	return val
}

// GetPathPaymentStrictReceiveResult retrieves the PathPaymentStrictReceiveResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetPathPaymentStrictReceiveResult() (result PathPaymentStrictReceiveResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PathPaymentStrictReceiveResult" {
		result = *u.PathPaymentStrictReceiveResult
		ok = true
	}

	return
}

// MustManageSellOfferResult retrieves the ManageSellOfferResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageSellOfferResult() ManageSellOfferResult {
	val, ok := u.GetManageSellOfferResult()

	if !ok {
		panic("arm ManageSellOfferResult is not set")
	}

	return val
}

// GetManageSellOfferResult retrieves the ManageSellOfferResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageSellOfferResult() (result ManageSellOfferResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageSellOfferResult" {
		result = *u.ManageSellOfferResult
		ok = true
	}

	return
}

// MustCreatePassiveSellOfferResult retrieves the CreatePassiveSellOfferResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreatePassiveSellOfferResult() ManageSellOfferResult {
	val, ok := u.GetCreatePassiveSellOfferResult()

	if !ok {
		panic("arm CreatePassiveSellOfferResult is not set")
	}

	return val
}

// GetCreatePassiveSellOfferResult retrieves the CreatePassiveSellOfferResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreatePassiveSellOfferResult() (result ManageSellOfferResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreatePassiveSellOfferResult" {
		result = *u.CreatePassiveSellOfferResult
		ok = true
	}

	return
}

// MustSetOptionsResult retrieves the SetOptionsResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustSetOptionsResult() SetOptionsResult {
	val, ok := u.GetSetOptionsResult()

	if !ok {
		panic("arm SetOptionsResult is not set")
	}

	return val
}

// GetSetOptionsResult retrieves the SetOptionsResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetSetOptionsResult() (result SetOptionsResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SetOptionsResult" {
		result = *u.SetOptionsResult
		ok = true
	}

	return
}

// MustChangeTrustResult retrieves the ChangeTrustResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustChangeTrustResult() ChangeTrustResult {
	val, ok := u.GetChangeTrustResult()

	if !ok {
		panic("arm ChangeTrustResult is not set")
	}

	return val
}

// GetChangeTrustResult retrieves the ChangeTrustResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetChangeTrustResult() (result ChangeTrustResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ChangeTrustResult" {
		result = *u.ChangeTrustResult
		ok = true
	}

	return
}

// MustAllowTrustResult retrieves the AllowTrustResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustAllowTrustResult() AllowTrustResult {
	val, ok := u.GetAllowTrustResult()

	if !ok {
		panic("arm AllowTrustResult is not set")
	}

	return val
}

// GetAllowTrustResult retrieves the AllowTrustResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetAllowTrustResult() (result AllowTrustResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AllowTrustResult" {
		result = *u.AllowTrustResult
		ok = true
	}

	return
}

// MustAccountMergeResult retrieves the AccountMergeResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustAccountMergeResult() AccountMergeResult {
	val, ok := u.GetAccountMergeResult()

	if !ok {
		panic("arm AccountMergeResult is not set")
	}

	return val
}

// GetAccountMergeResult retrieves the AccountMergeResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetAccountMergeResult() (result AccountMergeResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AccountMergeResult" {
		result = *u.AccountMergeResult
		ok = true
	}

	return
}

// MustInflationResult retrieves the InflationResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustInflationResult() InflationResult {
	val, ok := u.GetInflationResult()

	if !ok {
		panic("arm InflationResult is not set")
	}

	return val
}

// GetInflationResult retrieves the InflationResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetInflationResult() (result InflationResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "InflationResult" {
		result = *u.InflationResult
		ok = true
	}

	return
}

// MustManageDataResult retrieves the ManageDataResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageDataResult() ManageDataResult {
	val, ok := u.GetManageDataResult()

	if !ok {
		panic("arm ManageDataResult is not set")
	}

	return val
}

// GetManageDataResult retrieves the ManageDataResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageDataResult() (result ManageDataResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageDataResult" {
		result = *u.ManageDataResult
		ok = true
	}

	return
}

// MustBumpSeqResult retrieves the BumpSeqResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustBumpSeqResult() BumpSequenceResult {
	val, ok := u.GetBumpSeqResult()

	if !ok {
		panic("arm BumpSeqResult is not set")
	}

	return val
}

// GetBumpSeqResult retrieves the BumpSeqResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetBumpSeqResult() (result BumpSequenceResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "BumpSeqResult" {
		result = *u.BumpSeqResult
		ok = true
	}

	return
}

// MustManageBuyOfferResult retrieves the ManageBuyOfferResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageBuyOfferResult() ManageBuyOfferResult {
	val, ok := u.GetManageBuyOfferResult()

	if !ok {
		panic("arm ManageBuyOfferResult is not set")
	}

	return val
}

// GetManageBuyOfferResult retrieves the ManageBuyOfferResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageBuyOfferResult() (result ManageBuyOfferResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageBuyOfferResult" {
		result = *u.ManageBuyOfferResult
		ok = true
	}

	return
}

// MustPathPaymentStrictSendResult retrieves the PathPaymentStrictSendResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustPathPaymentStrictSendResult() PathPaymentStrictSendResult {
	val, ok := u.GetPathPaymentStrictSendResult()

	if !ok {
		panic("arm PathPaymentStrictSendResult is not set")
	}

	return val
}

// GetPathPaymentStrictSendResult retrieves the PathPaymentStrictSendResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetPathPaymentStrictSendResult() (result PathPaymentStrictSendResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PathPaymentStrictSendResult" {
		result = *u.PathPaymentStrictSendResult
		ok = true
	}

	return
}

// MustCreateClaimableBalanceResult retrieves the CreateClaimableBalanceResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateClaimableBalanceResult() CreateClaimableBalanceResult {
	val, ok := u.GetCreateClaimableBalanceResult()

	if !ok {
		panic("arm CreateClaimableBalanceResult is not set")
	}

	return val
}

// GetCreateClaimableBalanceResult retrieves the CreateClaimableBalanceResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateClaimableBalanceResult() (result CreateClaimableBalanceResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateClaimableBalanceResult" {
		result = *u.CreateClaimableBalanceResult
		ok = true
	}

	return
}

// MustClaimClaimableBalanceResult retrieves the ClaimClaimableBalanceResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustClaimClaimableBalanceResult() ClaimClaimableBalanceResult {
	val, ok := u.GetClaimClaimableBalanceResult()

	if !ok {
		panic("arm ClaimClaimableBalanceResult is not set")
	}

	return val
}

// GetClaimClaimableBalanceResult retrieves the ClaimClaimableBalanceResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetClaimClaimableBalanceResult() (result ClaimClaimableBalanceResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ClaimClaimableBalanceResult" {
		result = *u.ClaimClaimableBalanceResult
		ok = true
	}

	return
}

// MustBeginSponsoringFutureReservesResult retrieves the BeginSponsoringFutureReservesResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustBeginSponsoringFutureReservesResult() BeginSponsoringFutureReservesResult {
	val, ok := u.GetBeginSponsoringFutureReservesResult()

	if !ok {
		panic("arm BeginSponsoringFutureReservesResult is not set")
	}

	return val
}

// GetBeginSponsoringFutureReservesResult retrieves the BeginSponsoringFutureReservesResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetBeginSponsoringFutureReservesResult() (result BeginSponsoringFutureReservesResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "BeginSponsoringFutureReservesResult" {
		result = *u.BeginSponsoringFutureReservesResult
		ok = true
	}

	return
}

// MustEndSponsoringFutureReservesResult retrieves the EndSponsoringFutureReservesResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustEndSponsoringFutureReservesResult() EndSponsoringFutureReservesResult {
	val, ok := u.GetEndSponsoringFutureReservesResult()

	if !ok {
		panic("arm EndSponsoringFutureReservesResult is not set")
	}

	return val
}

// GetEndSponsoringFutureReservesResult retrieves the EndSponsoringFutureReservesResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetEndSponsoringFutureReservesResult() (result EndSponsoringFutureReservesResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "EndSponsoringFutureReservesResult" {
		result = *u.EndSponsoringFutureReservesResult
		ok = true
	}

	return
}

// MustRevokeSponsorshipResult retrieves the RevokeSponsorshipResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustRevokeSponsorshipResult() RevokeSponsorshipResult {
	val, ok := u.GetRevokeSponsorshipResult()

	if !ok {
		panic("arm RevokeSponsorshipResult is not set")
	}

	return val
}

// GetRevokeSponsorshipResult retrieves the RevokeSponsorshipResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetRevokeSponsorshipResult() (result RevokeSponsorshipResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RevokeSponsorshipResult" {
		result = *u.RevokeSponsorshipResult
		ok = true
	}

	return
}

// MustClawbackResult retrieves the ClawbackResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustClawbackResult() ClawbackResult {
	val, ok := u.GetClawbackResult()

	if !ok {
		panic("arm ClawbackResult is not set")
	}

	return val
}

// GetClawbackResult retrieves the ClawbackResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetClawbackResult() (result ClawbackResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ClawbackResult" {
		result = *u.ClawbackResult
		ok = true
	}

	return
}

// MustClawbackClaimableBalanceResult retrieves the ClawbackClaimableBalanceResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustClawbackClaimableBalanceResult() ClawbackClaimableBalanceResult {
	val, ok := u.GetClawbackClaimableBalanceResult()

	if !ok {
		panic("arm ClawbackClaimableBalanceResult is not set")
	}

	return val
}

// GetClawbackClaimableBalanceResult retrieves the ClawbackClaimableBalanceResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetClawbackClaimableBalanceResult() (result ClawbackClaimableBalanceResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ClawbackClaimableBalanceResult" {
		result = *u.ClawbackClaimableBalanceResult
		ok = true
	}

	return
}

// MustSetTrustLineFlagsResult retrieves the SetTrustLineFlagsResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustSetTrustLineFlagsResult() SetTrustLineFlagsResult {
	val, ok := u.GetSetTrustLineFlagsResult()

	if !ok {
		panic("arm SetTrustLineFlagsResult is not set")
	}

	return val
}

// GetSetTrustLineFlagsResult retrieves the SetTrustLineFlagsResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetSetTrustLineFlagsResult() (result SetTrustLineFlagsResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SetTrustLineFlagsResult" {
		result = *u.SetTrustLineFlagsResult
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s OperationResultTr) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *OperationResultTr) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*OperationResultTr)(nil)
	_ encoding.BinaryUnmarshaler = (*OperationResultTr)(nil)
)

// OperationResult is an XDR Union defines as:
//
//   union OperationResult switch (OperationResultCode code)
//    {
//    case opINNER:
//        union switch (OperationType type)
//        {
//        case CREATE_ACCOUNT:
//            CreateAccountResult createAccountResult;
//        case PAYMENT:
//            PaymentResult paymentResult;
//        case PATH_PAYMENT_STRICT_RECEIVE:
//            PathPaymentStrictReceiveResult pathPaymentStrictReceiveResult;
//        case MANAGE_SELL_OFFER:
//            ManageSellOfferResult manageSellOfferResult;
//        case CREATE_PASSIVE_SELL_OFFER:
//            ManageSellOfferResult createPassiveSellOfferResult;
//        case SET_OPTIONS:
//            SetOptionsResult setOptionsResult;
//        case CHANGE_TRUST:
//            ChangeTrustResult changeTrustResult;
//        case ALLOW_TRUST:
//            AllowTrustResult allowTrustResult;
//        case ACCOUNT_MERGE:
//            AccountMergeResult accountMergeResult;
//        case INFLATION:
//            InflationResult inflationResult;
//        case MANAGE_DATA:
//            ManageDataResult manageDataResult;
//        case BUMP_SEQUENCE:
//            BumpSequenceResult bumpSeqResult;
//        case MANAGE_BUY_OFFER:
//            ManageBuyOfferResult manageBuyOfferResult;
//        case PATH_PAYMENT_STRICT_SEND:
//            PathPaymentStrictSendResult pathPaymentStrictSendResult;
//        case CREATE_CLAIMABLE_BALANCE:
//            CreateClaimableBalanceResult createClaimableBalanceResult;
//        case CLAIM_CLAIMABLE_BALANCE:
//            ClaimClaimableBalanceResult claimClaimableBalanceResult;
//        case BEGIN_SPONSORING_FUTURE_RESERVES:
//            BeginSponsoringFutureReservesResult beginSponsoringFutureReservesResult;
//        case END_SPONSORING_FUTURE_RESERVES:
//            EndSponsoringFutureReservesResult endSponsoringFutureReservesResult;
//        case REVOKE_SPONSORSHIP:
//            RevokeSponsorshipResult revokeSponsorshipResult;
//        case CLAWBACK:
//            ClawbackResult clawbackResult;
//        case CLAWBACK_CLAIMABLE_BALANCE:
//            ClawbackClaimableBalanceResult clawbackClaimableBalanceResult;
//        case SET_TRUST_LINE_FLAGS:
//            SetTrustLineFlagsResult setTrustLineFlagsResult;
//        }
//        tr;
//    default:
//        void;
//    };
//
type OperationResult struct {
	Code OperationResultCode
	Tr   *OperationResultTr
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OperationResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OperationResult
func (u OperationResult) ArmForSwitch(sw int32) (string, bool) {
	switch OperationResultCode(sw) {
	case OperationResultCodeOpInner:
		return "Tr", true
	default:
		return "", true
	}
}

// NewOperationResult creates a new  OperationResult.
func NewOperationResult(code OperationResultCode, value interface{}) (result OperationResult, err error) {
	result.Code = code
	switch OperationResultCode(code) {
	case OperationResultCodeOpInner:
		tv, ok := value.(OperationResultTr)
		if !ok {
			err = fmt.Errorf("invalid value, must be OperationResultTr")
			return
		}
		result.Tr = &tv
	default:
		// void
	}
	return
}

// MustTr retrieves the Tr value from the union,
// panicing if the value is not set.
func (u OperationResult) MustTr() OperationResultTr {
	val, ok := u.GetTr()

	if !ok {
		panic("arm Tr is not set")
	}

	return val
}

// GetTr retrieves the Tr value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResult) GetTr() (result OperationResultTr, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Tr" {
		result = *u.Tr
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s OperationResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *OperationResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*OperationResult)(nil)
	_ encoding.BinaryUnmarshaler = (*OperationResult)(nil)
)

// TransactionResultCode is an XDR Enum defines as:
//
//   enum TransactionResultCode
//    {
//        txFEE_BUMP_INNER_SUCCESS = 1, // fee bump inner transaction succeeded
//        txSUCCESS = 0,                // all operations succeeded
//
//        txFAILED = -1, // one of the operations failed (none were applied)
//
//        txTOO_EARLY = -2,         // ledger closeTime before minTime
//        txTOO_LATE = -3,          // ledger closeTime after maxTime
//        txMISSING_OPERATION = -4, // no operation was specified
//        txBAD_SEQ = -5,           // sequence number does not match source account
//
//        txBAD_AUTH = -6,             // too few valid signatures / wrong network
//        txINSUFFICIENT_BALANCE = -7, // fee would bring account below reserve
//        txNO_ACCOUNT = -8,           // source account not found
//        txINSUFFICIENT_FEE = -9,     // fee is too small
//        txBAD_AUTH_EXTRA = -10,      // unused signatures attached to transaction
//        txINTERNAL_ERROR = -11,      // an unknown error occurred
//
//        txNOT_SUPPORTED = -12,         // transaction type not supported
//        txFEE_BUMP_INNER_FAILED = -13, // fee bump inner transaction failed
//        txBAD_SPONSORSHIP = -14        // sponsorship not confirmed
//    };
//
type TransactionResultCode int32

const (
	TransactionResultCodeTxFeeBumpInnerSuccess TransactionResultCode = 1
	TransactionResultCodeTxSuccess             TransactionResultCode = 0
	TransactionResultCodeTxFailed              TransactionResultCode = -1
	TransactionResultCodeTxTooEarly            TransactionResultCode = -2
	TransactionResultCodeTxTooLate             TransactionResultCode = -3
	TransactionResultCodeTxMissingOperation    TransactionResultCode = -4
	TransactionResultCodeTxBadSeq              TransactionResultCode = -5
	TransactionResultCodeTxBadAuth             TransactionResultCode = -6
	TransactionResultCodeTxInsufficientBalance TransactionResultCode = -7
	TransactionResultCodeTxNoAccount           TransactionResultCode = -8
	TransactionResultCodeTxInsufficientFee     TransactionResultCode = -9
	TransactionResultCodeTxBadAuthExtra        TransactionResultCode = -10
	TransactionResultCodeTxInternalError       TransactionResultCode = -11
	TransactionResultCodeTxNotSupported        TransactionResultCode = -12
	TransactionResultCodeTxFeeBumpInnerFailed  TransactionResultCode = -13
	TransactionResultCodeTxBadSponsorship      TransactionResultCode = -14
)

var transactionResultCodeMap = map[int32]string{
	1:   "TransactionResultCodeTxFeeBumpInnerSuccess",
	0:   "TransactionResultCodeTxSuccess",
	-1:  "TransactionResultCodeTxFailed",
	-2:  "TransactionResultCodeTxTooEarly",
	-3:  "TransactionResultCodeTxTooLate",
	-4:  "TransactionResultCodeTxMissingOperation",
	-5:  "TransactionResultCodeTxBadSeq",
	-6:  "TransactionResultCodeTxBadAuth",
	-7:  "TransactionResultCodeTxInsufficientBalance",
	-8:  "TransactionResultCodeTxNoAccount",
	-9:  "TransactionResultCodeTxInsufficientFee",
	-10: "TransactionResultCodeTxBadAuthExtra",
	-11: "TransactionResultCodeTxInternalError",
	-12: "TransactionResultCodeTxNotSupported",
	-13: "TransactionResultCodeTxFeeBumpInnerFailed",
	-14: "TransactionResultCodeTxBadSponsorship",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for TransactionResultCode
func (e TransactionResultCode) ValidEnum(v int32) bool {
	_, ok := transactionResultCodeMap[v]
	return ok
}

// String returns the name of `e`
func (e TransactionResultCode) String() string {
	name, _ := transactionResultCodeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionResultCode) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionResultCode) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionResultCode)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionResultCode)(nil)
)

// InnerTransactionResultResult is an XDR NestedUnion defines as:
//
//   union switch (TransactionResultCode code)
//        {
//        // txFEE_BUMP_INNER_SUCCESS is not included
//        case txSUCCESS:
//        case txFAILED:
//            OperationResult results<>;
//        case txTOO_EARLY:
//        case txTOO_LATE:
//        case txMISSING_OPERATION:
//        case txBAD_SEQ:
//        case txBAD_AUTH:
//        case txINSUFFICIENT_BALANCE:
//        case txNO_ACCOUNT:
//        case txINSUFFICIENT_FEE:
//        case txBAD_AUTH_EXTRA:
//        case txINTERNAL_ERROR:
//        case txNOT_SUPPORTED:
//        // txFEE_BUMP_INNER_FAILED is not included
//        case txBAD_SPONSORSHIP:
//            void;
//        }
//
type InnerTransactionResultResult struct {
	Code    TransactionResultCode
	Results *[]OperationResult
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u InnerTransactionResultResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of InnerTransactionResultResult
func (u InnerTransactionResultResult) ArmForSwitch(sw int32) (string, bool) {
	switch TransactionResultCode(sw) {
	case TransactionResultCodeTxSuccess:
		return "Results", true
	case TransactionResultCodeTxFailed:
		return "Results", true
	case TransactionResultCodeTxTooEarly:
		return "", true
	case TransactionResultCodeTxTooLate:
		return "", true
	case TransactionResultCodeTxMissingOperation:
		return "", true
	case TransactionResultCodeTxBadSeq:
		return "", true
	case TransactionResultCodeTxBadAuth:
		return "", true
	case TransactionResultCodeTxInsufficientBalance:
		return "", true
	case TransactionResultCodeTxNoAccount:
		return "", true
	case TransactionResultCodeTxInsufficientFee:
		return "", true
	case TransactionResultCodeTxBadAuthExtra:
		return "", true
	case TransactionResultCodeTxInternalError:
		return "", true
	case TransactionResultCodeTxNotSupported:
		return "", true
	case TransactionResultCodeTxBadSponsorship:
		return "", true
	}
	return "-", false
}

// NewInnerTransactionResultResult creates a new  InnerTransactionResultResult.
func NewInnerTransactionResultResult(code TransactionResultCode, value interface{}) (result InnerTransactionResultResult, err error) {
	result.Code = code
	switch TransactionResultCode(code) {
	case TransactionResultCodeTxSuccess:
		tv, ok := value.([]OperationResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be []OperationResult")
			return
		}
		result.Results = &tv
	case TransactionResultCodeTxFailed:
		tv, ok := value.([]OperationResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be []OperationResult")
			return
		}
		result.Results = &tv
	case TransactionResultCodeTxTooEarly:
		// void
	case TransactionResultCodeTxTooLate:
		// void
	case TransactionResultCodeTxMissingOperation:
		// void
	case TransactionResultCodeTxBadSeq:
		// void
	case TransactionResultCodeTxBadAuth:
		// void
	case TransactionResultCodeTxInsufficientBalance:
		// void
	case TransactionResultCodeTxNoAccount:
		// void
	case TransactionResultCodeTxInsufficientFee:
		// void
	case TransactionResultCodeTxBadAuthExtra:
		// void
	case TransactionResultCodeTxInternalError:
		// void
	case TransactionResultCodeTxNotSupported:
		// void
	case TransactionResultCodeTxBadSponsorship:
		// void
	}
	return
}

// MustResults retrieves the Results value from the union,
// panicing if the value is not set.
func (u InnerTransactionResultResult) MustResults() []OperationResult {
	val, ok := u.GetResults()

	if !ok {
		panic("arm Results is not set")
	}

	return val
}

// GetResults retrieves the Results value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u InnerTransactionResultResult) GetResults() (result []OperationResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Results" {
		result = *u.Results
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s InnerTransactionResultResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *InnerTransactionResultResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*InnerTransactionResultResult)(nil)
	_ encoding.BinaryUnmarshaler = (*InnerTransactionResultResult)(nil)
)

// InnerTransactionResultExt is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        }
//
type InnerTransactionResultExt struct {
	V int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u InnerTransactionResultExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of InnerTransactionResultExt
func (u InnerTransactionResultExt) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	}
	return "-", false
}

// NewInnerTransactionResultExt creates a new  InnerTransactionResultExt.
func NewInnerTransactionResultExt(v int32, value interface{}) (result InnerTransactionResultExt, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s InnerTransactionResultExt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *InnerTransactionResultExt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*InnerTransactionResultExt)(nil)
	_ encoding.BinaryUnmarshaler = (*InnerTransactionResultExt)(nil)
)

// InnerTransactionResult is an XDR Struct defines as:
//
//   struct InnerTransactionResult
//    {
//        // Always 0. Here for binary compatibility.
//        int64 feeCharged;
//
//        union switch (TransactionResultCode code)
//        {
//        // txFEE_BUMP_INNER_SUCCESS is not included
//        case txSUCCESS:
//        case txFAILED:
//            OperationResult results<>;
//        case txTOO_EARLY:
//        case txTOO_LATE:
//        case txMISSING_OPERATION:
//        case txBAD_SEQ:
//        case txBAD_AUTH:
//        case txINSUFFICIENT_BALANCE:
//        case txNO_ACCOUNT:
//        case txINSUFFICIENT_FEE:
//        case txBAD_AUTH_EXTRA:
//        case txINTERNAL_ERROR:
//        case txNOT_SUPPORTED:
//        // txFEE_BUMP_INNER_FAILED is not included
//        case txBAD_SPONSORSHIP:
//            void;
//        }
//        result;
//
//        // reserved for future use
//        union switch (int v)
//        {
//        case 0:
//            void;
//        }
//        ext;
//    };
//
type InnerTransactionResult struct {
	FeeCharged Int64
	Result     InnerTransactionResultResult
	Ext        InnerTransactionResultExt
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s InnerTransactionResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *InnerTransactionResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*InnerTransactionResult)(nil)
	_ encoding.BinaryUnmarshaler = (*InnerTransactionResult)(nil)
)

// InnerTransactionResultPair is an XDR Struct defines as:
//
//   struct InnerTransactionResultPair
//    {
//        Hash transactionHash;          // hash of the inner transaction
//        InnerTransactionResult result; // result for the inner transaction
//    };
//
type InnerTransactionResultPair struct {
	TransactionHash Hash
	Result          InnerTransactionResult
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s InnerTransactionResultPair) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *InnerTransactionResultPair) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*InnerTransactionResultPair)(nil)
	_ encoding.BinaryUnmarshaler = (*InnerTransactionResultPair)(nil)
)

// TransactionResultResult is an XDR NestedUnion defines as:
//
//   union switch (TransactionResultCode code)
//        {
//        case txFEE_BUMP_INNER_SUCCESS:
//        case txFEE_BUMP_INNER_FAILED:
//            InnerTransactionResultPair innerResultPair;
//        case txSUCCESS:
//        case txFAILED:
//            OperationResult results<>;
//        default:
//            void;
//        }
//
type TransactionResultResult struct {
	Code            TransactionResultCode
	InnerResultPair *InnerTransactionResultPair
	Results         *[]OperationResult
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionResultResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionResultResult
func (u TransactionResultResult) ArmForSwitch(sw int32) (string, bool) {
	switch TransactionResultCode(sw) {
	case TransactionResultCodeTxFeeBumpInnerSuccess:
		return "InnerResultPair", true
	case TransactionResultCodeTxFeeBumpInnerFailed:
		return "InnerResultPair", true
	case TransactionResultCodeTxSuccess:
		return "Results", true
	case TransactionResultCodeTxFailed:
		return "Results", true
	default:
		return "", true
	}
}

// NewTransactionResultResult creates a new  TransactionResultResult.
func NewTransactionResultResult(code TransactionResultCode, value interface{}) (result TransactionResultResult, err error) {
	result.Code = code
	switch TransactionResultCode(code) {
	case TransactionResultCodeTxFeeBumpInnerSuccess:
		tv, ok := value.(InnerTransactionResultPair)
		if !ok {
			err = fmt.Errorf("invalid value, must be InnerTransactionResultPair")
			return
		}
		result.InnerResultPair = &tv
	case TransactionResultCodeTxFeeBumpInnerFailed:
		tv, ok := value.(InnerTransactionResultPair)
		if !ok {
			err = fmt.Errorf("invalid value, must be InnerTransactionResultPair")
			return
		}
		result.InnerResultPair = &tv
	case TransactionResultCodeTxSuccess:
		tv, ok := value.([]OperationResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be []OperationResult")
			return
		}
		result.Results = &tv
	case TransactionResultCodeTxFailed:
		tv, ok := value.([]OperationResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be []OperationResult")
			return
		}
		result.Results = &tv
	default:
		// void
	}
	return
}

// MustInnerResultPair retrieves the InnerResultPair value from the union,
// panicing if the value is not set.
func (u TransactionResultResult) MustInnerResultPair() InnerTransactionResultPair {
	val, ok := u.GetInnerResultPair()

	if !ok {
		panic("arm InnerResultPair is not set")
	}

	return val
}

// GetInnerResultPair retrieves the InnerResultPair value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionResultResult) GetInnerResultPair() (result InnerTransactionResultPair, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "InnerResultPair" {
		result = *u.InnerResultPair
		ok = true
	}

	return
}

// MustResults retrieves the Results value from the union,
// panicing if the value is not set.
func (u TransactionResultResult) MustResults() []OperationResult {
	val, ok := u.GetResults()

	if !ok {
		panic("arm Results is not set")
	}

	return val
}

// GetResults retrieves the Results value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionResultResult) GetResults() (result []OperationResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Results" {
		result = *u.Results
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionResultResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionResultResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionResultResult)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionResultResult)(nil)
)

// TransactionResultExt is an XDR NestedUnion defines as:
//
//   union switch (int v)
//        {
//        case 0:
//            void;
//        }
//
type TransactionResultExt struct {
	V int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionResultExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionResultExt
func (u TransactionResultExt) ArmForSwitch(sw int32) (string, bool) {
	switch int32(sw) {
	case 0:
		return "", true
	}
	return "-", false
}

// NewTransactionResultExt creates a new  TransactionResultExt.
func NewTransactionResultExt(v int32, value interface{}) (result TransactionResultExt, err error) {
	result.V = v
	switch int32(v) {
	case 0:
		// void
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionResultExt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionResultExt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionResultExt)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionResultExt)(nil)
)

// TransactionResult is an XDR Struct defines as:
//
//   struct TransactionResult
//    {
//        int64 feeCharged; // actual fee charged for the transaction
//
//        union switch (TransactionResultCode code)
//        {
//        case txFEE_BUMP_INNER_SUCCESS:
//        case txFEE_BUMP_INNER_FAILED:
//            InnerTransactionResultPair innerResultPair;
//        case txSUCCESS:
//        case txFAILED:
//            OperationResult results<>;
//        default:
//            void;
//        }
//        result;
//
//        // reserved for future use
//        union switch (int v)
//        {
//        case 0:
//            void;
//        }
//        ext;
//    };
//
type TransactionResult struct {
	FeeCharged Int64
	Result     TransactionResultResult
	Ext        TransactionResultExt
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionResult)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionResult)(nil)
)

// Hash is an XDR Typedef defines as:
//
//   typedef opaque Hash[32];
//
type Hash [32]byte

// XDRMaxSize implements the Sized interface for Hash
func (e Hash) XDRMaxSize() int {
	return 32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Hash) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Hash) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Hash)(nil)
	_ encoding.BinaryUnmarshaler = (*Hash)(nil)
)

// Uint256 is an XDR Typedef defines as:
//
//   typedef opaque uint256[32];
//
type Uint256 [32]byte

// XDRMaxSize implements the Sized interface for Uint256
func (e Uint256) XDRMaxSize() int {
	return 32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Uint256) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Uint256) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Uint256)(nil)
	_ encoding.BinaryUnmarshaler = (*Uint256)(nil)
)

// Uint32 is an XDR Typedef defines as:
//
//   typedef unsigned int uint32;
//
type Uint32 uint32

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Uint32) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Uint32) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Uint32)(nil)
	_ encoding.BinaryUnmarshaler = (*Uint32)(nil)
)

// Int32 is an XDR Typedef defines as:
//
//   typedef int int32;
//
type Int32 int32

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Int32) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Int32) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Int32)(nil)
	_ encoding.BinaryUnmarshaler = (*Int32)(nil)
)

// Uint64 is an XDR Typedef defines as:
//
//   typedef unsigned hyper uint64;
//
type Uint64 uint64

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Uint64) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Uint64) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Uint64)(nil)
	_ encoding.BinaryUnmarshaler = (*Uint64)(nil)
)

// Int64 is an XDR Typedef defines as:
//
//   typedef hyper int64;
//
type Int64 int64

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Int64) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Int64) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Int64)(nil)
	_ encoding.BinaryUnmarshaler = (*Int64)(nil)
)

// CryptoKeyType is an XDR Enum defines as:
//
//   enum CryptoKeyType
//    {
//        KEY_TYPE_ED25519 = 0,
//        KEY_TYPE_PRE_AUTH_TX = 1,
//        KEY_TYPE_HASH_X = 2,
//        // MUXED enum values for supported type are derived from the enum values
//        // above by ORing them with 0x100
//        KEY_TYPE_MUXED_ED25519 = 0x100
//    };
//
type CryptoKeyType int32

const (
	CryptoKeyTypeKeyTypeEd25519      CryptoKeyType = 0
	CryptoKeyTypeKeyTypePreAuthTx    CryptoKeyType = 1
	CryptoKeyTypeKeyTypeHashX        CryptoKeyType = 2
	CryptoKeyTypeKeyTypeMuxedEd25519 CryptoKeyType = 256
)

var cryptoKeyTypeMap = map[int32]string{
	0:   "CryptoKeyTypeKeyTypeEd25519",
	1:   "CryptoKeyTypeKeyTypePreAuthTx",
	2:   "CryptoKeyTypeKeyTypeHashX",
	256: "CryptoKeyTypeKeyTypeMuxedEd25519",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CryptoKeyType
func (e CryptoKeyType) ValidEnum(v int32) bool {
	_, ok := cryptoKeyTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e CryptoKeyType) String() string {
	name, _ := cryptoKeyTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s CryptoKeyType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *CryptoKeyType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*CryptoKeyType)(nil)
	_ encoding.BinaryUnmarshaler = (*CryptoKeyType)(nil)
)

// PublicKeyType is an XDR Enum defines as:
//
//   enum PublicKeyType
//    {
//        PUBLIC_KEY_TYPE_ED25519 = KEY_TYPE_ED25519
//    };
//
type PublicKeyType int32

const (
	PublicKeyTypePublicKeyTypeEd25519 PublicKeyType = 0
)

var publicKeyTypeMap = map[int32]string{
	0: "PublicKeyTypePublicKeyTypeEd25519",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PublicKeyType
func (e PublicKeyType) ValidEnum(v int32) bool {
	_, ok := publicKeyTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e PublicKeyType) String() string {
	name, _ := publicKeyTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PublicKeyType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PublicKeyType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PublicKeyType)(nil)
	_ encoding.BinaryUnmarshaler = (*PublicKeyType)(nil)
)

// SignerKeyType is an XDR Enum defines as:
//
//   enum SignerKeyType
//    {
//        SIGNER_KEY_TYPE_ED25519 = KEY_TYPE_ED25519,
//        SIGNER_KEY_TYPE_PRE_AUTH_TX = KEY_TYPE_PRE_AUTH_TX,
//        SIGNER_KEY_TYPE_HASH_X = KEY_TYPE_HASH_X
//    };
//
type SignerKeyType int32

const (
	SignerKeyTypeSignerKeyTypeEd25519   SignerKeyType = 0
	SignerKeyTypeSignerKeyTypePreAuthTx SignerKeyType = 1
	SignerKeyTypeSignerKeyTypeHashX     SignerKeyType = 2
)

var signerKeyTypeMap = map[int32]string{
	0: "SignerKeyTypeSignerKeyTypeEd25519",
	1: "SignerKeyTypeSignerKeyTypePreAuthTx",
	2: "SignerKeyTypeSignerKeyTypeHashX",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for SignerKeyType
func (e SignerKeyType) ValidEnum(v int32) bool {
	_, ok := signerKeyTypeMap[v]
	return ok
}

// String returns the name of `e`
func (e SignerKeyType) String() string {
	name, _ := signerKeyTypeMap[int32(e)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SignerKeyType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SignerKeyType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SignerKeyType)(nil)
	_ encoding.BinaryUnmarshaler = (*SignerKeyType)(nil)
)

// PublicKey is an XDR Union defines as:
//
//   union PublicKey switch (PublicKeyType type)
//    {
//    case PUBLIC_KEY_TYPE_ED25519:
//        uint256 ed25519;
//    };
//
type PublicKey struct {
	Type    PublicKeyType
	Ed25519 *Uint256
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PublicKey) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PublicKey
func (u PublicKey) ArmForSwitch(sw int32) (string, bool) {
	switch PublicKeyType(sw) {
	case PublicKeyTypePublicKeyTypeEd25519:
		return "Ed25519", true
	}
	return "-", false
}

// NewPublicKey creates a new  PublicKey.
func NewPublicKey(aType PublicKeyType, value interface{}) (result PublicKey, err error) {
	result.Type = aType
	switch PublicKeyType(aType) {
	case PublicKeyTypePublicKeyTypeEd25519:
		tv, ok := value.(Uint256)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint256")
			return
		}
		result.Ed25519 = &tv
	}
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u PublicKey) MustEd25519() Uint256 {
	val, ok := u.GetEd25519()

	if !ok {
		panic("arm Ed25519 is not set")
	}

	return val
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PublicKey) GetEd25519() (result Uint256, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ed25519" {
		result = *u.Ed25519
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PublicKey) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PublicKey) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PublicKey)(nil)
	_ encoding.BinaryUnmarshaler = (*PublicKey)(nil)
)

// SignerKey is an XDR Union defines as:
//
//   union SignerKey switch (SignerKeyType type)
//    {
//    case SIGNER_KEY_TYPE_ED25519:
//        uint256 ed25519;
//    case SIGNER_KEY_TYPE_PRE_AUTH_TX:
//        /* SHA-256 Hash of TransactionSignaturePayload structure */
//        uint256 preAuthTx;
//    case SIGNER_KEY_TYPE_HASH_X:
//        /* Hash of random 256 bit preimage X */
//        uint256 hashX;
//    };
//
type SignerKey struct {
	Type      SignerKeyType
	Ed25519   *Uint256
	PreAuthTx *Uint256
	HashX     *Uint256
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SignerKey) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SignerKey
func (u SignerKey) ArmForSwitch(sw int32) (string, bool) {
	switch SignerKeyType(sw) {
	case SignerKeyTypeSignerKeyTypeEd25519:
		return "Ed25519", true
	case SignerKeyTypeSignerKeyTypePreAuthTx:
		return "PreAuthTx", true
	case SignerKeyTypeSignerKeyTypeHashX:
		return "HashX", true
	}
	return "-", false
}

// NewSignerKey creates a new  SignerKey.
func NewSignerKey(aType SignerKeyType, value interface{}) (result SignerKey, err error) {
	result.Type = aType
	switch SignerKeyType(aType) {
	case SignerKeyTypeSignerKeyTypeEd25519:
		tv, ok := value.(Uint256)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint256")
			return
		}
		result.Ed25519 = &tv
	case SignerKeyTypeSignerKeyTypePreAuthTx:
		tv, ok := value.(Uint256)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint256")
			return
		}
		result.PreAuthTx = &tv
	case SignerKeyTypeSignerKeyTypeHashX:
		tv, ok := value.(Uint256)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint256")
			return
		}
		result.HashX = &tv
	}
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u SignerKey) MustEd25519() Uint256 {
	val, ok := u.GetEd25519()

	if !ok {
		panic("arm Ed25519 is not set")
	}

	return val
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SignerKey) GetEd25519() (result Uint256, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ed25519" {
		result = *u.Ed25519
		ok = true
	}

	return
}

// MustPreAuthTx retrieves the PreAuthTx value from the union,
// panicing if the value is not set.
func (u SignerKey) MustPreAuthTx() Uint256 {
	val, ok := u.GetPreAuthTx()

	if !ok {
		panic("arm PreAuthTx is not set")
	}

	return val
}

// GetPreAuthTx retrieves the PreAuthTx value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SignerKey) GetPreAuthTx() (result Uint256, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PreAuthTx" {
		result = *u.PreAuthTx
		ok = true
	}

	return
}

// MustHashX retrieves the HashX value from the union,
// panicing if the value is not set.
func (u SignerKey) MustHashX() Uint256 {
	val, ok := u.GetHashX()

	if !ok {
		panic("arm HashX is not set")
	}

	return val
}

// GetHashX retrieves the HashX value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SignerKey) GetHashX() (result Uint256, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "HashX" {
		result = *u.HashX
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SignerKey) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SignerKey) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SignerKey)(nil)
	_ encoding.BinaryUnmarshaler = (*SignerKey)(nil)
)

// Signature is an XDR Typedef defines as:
//
//   typedef opaque Signature<64>;
//
type Signature []byte

// XDRMaxSize implements the Sized interface for Signature
func (e Signature) XDRMaxSize() int {
	return 64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Signature) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Signature) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Signature)(nil)
	_ encoding.BinaryUnmarshaler = (*Signature)(nil)
)

// SignatureHint is an XDR Typedef defines as:
//
//   typedef opaque SignatureHint[4];
//
type SignatureHint [4]byte

// XDRMaxSize implements the Sized interface for SignatureHint
func (e SignatureHint) XDRMaxSize() int {
	return 4
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s SignatureHint) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *SignatureHint) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*SignatureHint)(nil)
	_ encoding.BinaryUnmarshaler = (*SignatureHint)(nil)
)

// NodeId is an XDR Typedef defines as:
//
//   typedef PublicKey NodeID;
//
type NodeId PublicKey

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u NodeId) SwitchFieldName() string {
	return PublicKey(u).SwitchFieldName()
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PublicKey
func (u NodeId) ArmForSwitch(sw int32) (string, bool) {
	return PublicKey(u).ArmForSwitch(sw)
}

// NewNodeId creates a new  NodeId.
func NewNodeId(aType PublicKeyType, value interface{}) (result NodeId, err error) {
	u, err := NewPublicKey(aType, value)
	result = NodeId(u)
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u NodeId) MustEd25519() Uint256 {
	return PublicKey(u).MustEd25519()
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u NodeId) GetEd25519() (result Uint256, ok bool) {
	return PublicKey(u).GetEd25519()
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s NodeId) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *NodeId) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*NodeId)(nil)
	_ encoding.BinaryUnmarshaler = (*NodeId)(nil)
)

// Curve25519Secret is an XDR Struct defines as:
//
//   struct Curve25519Secret
//    {
//        opaque key[32];
//    };
//
type Curve25519Secret struct {
	Key [32]byte `xdrmaxsize:"32"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Curve25519Secret) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Curve25519Secret) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Curve25519Secret)(nil)
	_ encoding.BinaryUnmarshaler = (*Curve25519Secret)(nil)
)

// Curve25519Public is an XDR Struct defines as:
//
//   struct Curve25519Public
//    {
//        opaque key[32];
//    };
//
type Curve25519Public struct {
	Key [32]byte `xdrmaxsize:"32"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Curve25519Public) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Curve25519Public) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Curve25519Public)(nil)
	_ encoding.BinaryUnmarshaler = (*Curve25519Public)(nil)
)

// HmacSha256Key is an XDR Struct defines as:
//
//   struct HmacSha256Key
//    {
//        opaque key[32];
//    };
//
type HmacSha256Key struct {
	Key [32]byte `xdrmaxsize:"32"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s HmacSha256Key) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *HmacSha256Key) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*HmacSha256Key)(nil)
	_ encoding.BinaryUnmarshaler = (*HmacSha256Key)(nil)
)

// HmacSha256Mac is an XDR Struct defines as:
//
//   struct HmacSha256Mac
//    {
//        opaque mac[32];
//    };
//
type HmacSha256Mac struct {
	Mac [32]byte `xdrmaxsize:"32"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s HmacSha256Mac) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *HmacSha256Mac) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*HmacSha256Mac)(nil)
	_ encoding.BinaryUnmarshaler = (*HmacSha256Mac)(nil)
)

var fmtTest = fmt.Sprint("this is a dummy usage of fmt")
