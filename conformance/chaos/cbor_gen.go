// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package chaos

import (
	"fmt"
	"io"
	"math"
	"sort"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"

	address "github.com/filecoin-project/go-address"
	abi "github.com/filecoin-project/go-state-types/abi"
	exitcode "github.com/filecoin-project/go-state-types/exitcode"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufState = []byte{130}

func (t *State) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufState); err != nil {
		return err
	}

	// t.Value (string) (string)
	if len(t.Value) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Value was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Value))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Value)); err != nil {
		return err
	}

	// t.Unmarshallable ([]*chaos.UnmarshallableCBOR) (slice)
	if len(t.Unmarshallable) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Unmarshallable was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Unmarshallable))); err != nil {
		return err
	}
	for _, v := range t.Unmarshallable {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}

	}
	return nil
}

func (t *State) UnmarshalCBOR(r io.Reader) (err error) {
	*t = State{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Value (string) (string)

	{
		sval, err := cbg.ReadString(cr)
		if err != nil {
			return err
		}

		t.Value = string(sval)
	}
	// t.Unmarshallable ([]*chaos.UnmarshallableCBOR) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Unmarshallable: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Unmarshallable = make([]*UnmarshallableCBOR, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error
			_ = maj
			_ = extra
			_ = err

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.Unmarshallable[i] = new(UnmarshallableCBOR)
					if err := t.Unmarshallable[i].UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.Unmarshallable[i] pointer: %w", err)
					}
				}

			}

		}
	}
	return nil
}

var lengthBufCallerValidationArgs = []byte{131}

func (t *CallerValidationArgs) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufCallerValidationArgs); err != nil {
		return err
	}

	// t.Branch (chaos.CallerValidationBranch) (int64)
	if t.Branch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Branch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.Branch-1)); err != nil {
			return err
		}
	}

	// t.Addrs ([]address.Address) (slice)
	if len(t.Addrs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Addrs was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Addrs))); err != nil {
		return err
	}
	for _, v := range t.Addrs {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}

	}

	// t.Types ([]cid.Cid) (slice)
	if len(t.Types) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Types was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Types))); err != nil {
		return err
	}
	for _, v := range t.Types {

		if err := cbg.WriteCid(cw, v); err != nil {
			return xerrors.Errorf("failed to write cid field v: %w", err)
		}

	}
	return nil
}

func (t *CallerValidationArgs) UnmarshalCBOR(r io.Reader) (err error) {
	*t = CallerValidationArgs{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Branch (chaos.CallerValidationBranch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Branch = CallerValidationBranch(extraI)
	}
	// t.Addrs ([]address.Address) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Addrs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Addrs = make([]address.Address, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error
			_ = maj
			_ = extra
			_ = err

			{

				if err := t.Addrs[i].UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.Addrs[i]: %w", err)
				}

			}

		}
	}
	// t.Types ([]cid.Cid) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Types: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Types = make([]cid.Cid, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error
			_ = maj
			_ = extra
			_ = err

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.Types[i]: %w", err)
				}

				t.Types[i] = c

			}

		}
	}
	return nil
}

var lengthBufCreateActorArgs = []byte{132}

func (t *CreateActorArgs) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufCreateActorArgs); err != nil {
		return err
	}

	// t.UndefActorCID (bool) (bool)
	if err := cbg.WriteBool(w, t.UndefActorCID); err != nil {
		return err
	}

	// t.ActorCID (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.ActorCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.ActorCID: %w", err)
	}

	// t.UndefAddress (bool) (bool)
	if err := cbg.WriteBool(w, t.UndefAddress); err != nil {
		return err
	}

	// t.Address (address.Address) (struct)
	if err := t.Address.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *CreateActorArgs) UnmarshalCBOR(r io.Reader) (err error) {
	*t = CreateActorArgs{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.UndefActorCID (bool) (bool)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.UndefActorCID = false
	case 21:
		t.UndefActorCID = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	// t.ActorCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.ActorCID: %w", err)
		}

		t.ActorCID = c

	}
	// t.UndefAddress (bool) (bool)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.UndefAddress = false
	case 21:
		t.UndefAddress = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	// t.Address (address.Address) (struct)

	{

		if err := t.Address.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Address: %w", err)
		}

	}
	return nil
}

var lengthBufResolveAddressResponse = []byte{130}

func (t *ResolveAddressResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufResolveAddressResponse); err != nil {
		return err
	}

	// t.Address (address.Address) (struct)
	if err := t.Address.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Success (bool) (bool)
	if err := cbg.WriteBool(w, t.Success); err != nil {
		return err
	}
	return nil
}

func (t *ResolveAddressResponse) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ResolveAddressResponse{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Address (address.Address) (struct)

	{

		if err := t.Address.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Address: %w", err)
		}

	}
	// t.Success (bool) (bool)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.Success = false
	case 21:
		t.Success = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	return nil
}

var lengthBufSendArgs = []byte{132}

func (t *SendArgs) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufSendArgs); err != nil {
		return err
	}

	// t.To (address.Address) (struct)
	if err := t.To.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Value (big.Int) (struct)
	if err := t.Value.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Method (abi.MethodNum) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Method)); err != nil {
		return err
	}

	// t.Params ([]uint8) (slice)
	if len(t.Params) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Params was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Params))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Params); err != nil {
		return err
	}

	return nil
}

func (t *SendArgs) UnmarshalCBOR(r io.Reader) (err error) {
	*t = SendArgs{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.To (address.Address) (struct)

	{

		if err := t.To.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.To: %w", err)
		}

	}
	// t.Value (big.Int) (struct)

	{

		if err := t.Value.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Value: %w", err)
		}

	}
	// t.Method (abi.MethodNum) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Method = abi.MethodNum(extra)

	}
	// t.Params ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Params: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Params = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Params); err != nil {
		return err
	}

	return nil
}

var lengthBufSendReturn = []byte{130}

func (t *SendReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufSendReturn); err != nil {
		return err
	}

	// t.Return (builtin.CBORBytes) (slice)
	if len(t.Return) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Return was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Return))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Return); err != nil {
		return err
	}

	// t.Code (exitcode.ExitCode) (int64)
	if t.Code >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Code)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.Code-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *SendReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*t = SendReturn{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Return (builtin.CBORBytes) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Return: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Return = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Return); err != nil {
		return err
	}

	// t.Code (exitcode.ExitCode) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Code = exitcode.ExitCode(extraI)
	}
	return nil
}

var lengthBufMutateStateArgs = []byte{130}

func (t *MutateStateArgs) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufMutateStateArgs); err != nil {
		return err
	}

	// t.Value (string) (string)
	if len(t.Value) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Value was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Value))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Value)); err != nil {
		return err
	}

	// t.Branch (chaos.MutateStateBranch) (int64)
	if t.Branch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Branch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.Branch-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *MutateStateArgs) UnmarshalCBOR(r io.Reader) (err error) {
	*t = MutateStateArgs{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Value (string) (string)

	{
		sval, err := cbg.ReadString(cr)
		if err != nil {
			return err
		}

		t.Value = string(sval)
	}
	// t.Branch (chaos.MutateStateBranch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Branch = MutateStateBranch(extraI)
	}
	return nil
}

var lengthBufAbortWithArgs = []byte{131}

func (t *AbortWithArgs) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufAbortWithArgs); err != nil {
		return err
	}

	// t.Code (exitcode.ExitCode) (int64)
	if t.Code >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Code)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.Code-1)); err != nil {
			return err
		}
	}

	// t.Message (string) (string)
	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Message))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Message)); err != nil {
		return err
	}

	// t.Uncontrolled (bool) (bool)
	if err := cbg.WriteBool(w, t.Uncontrolled); err != nil {
		return err
	}
	return nil
}

func (t *AbortWithArgs) UnmarshalCBOR(r io.Reader) (err error) {
	*t = AbortWithArgs{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Code (exitcode.ExitCode) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Code = exitcode.ExitCode(extraI)
	}
	// t.Message (string) (string)

	{
		sval, err := cbg.ReadString(cr)
		if err != nil {
			return err
		}

		t.Message = string(sval)
	}
	// t.Uncontrolled (bool) (bool)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.Uncontrolled = false
	case 21:
		t.Uncontrolled = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	return nil
}

var lengthBufInspectRuntimeReturn = []byte{134}

func (t *InspectRuntimeReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufInspectRuntimeReturn); err != nil {
		return err
	}

	// t.Caller (address.Address) (struct)
	if err := t.Caller.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Receiver (address.Address) (struct)
	if err := t.Receiver.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ValueReceived (big.Int) (struct)
	if err := t.ValueReceived.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.CurrEpoch (abi.ChainEpoch) (int64)
	if t.CurrEpoch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.CurrEpoch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.CurrEpoch-1)); err != nil {
			return err
		}
	}

	// t.CurrentBalance (big.Int) (struct)
	if err := t.CurrentBalance.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.State (chaos.State) (struct)
	if err := t.State.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *InspectRuntimeReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*t = InspectRuntimeReturn{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 6 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Caller (address.Address) (struct)

	{

		if err := t.Caller.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Caller: %w", err)
		}

	}
	// t.Receiver (address.Address) (struct)

	{

		if err := t.Receiver.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Receiver: %w", err)
		}

	}
	// t.ValueReceived (big.Int) (struct)

	{

		if err := t.ValueReceived.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.ValueReceived: %w", err)
		}

	}
	// t.CurrEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.CurrEpoch = abi.ChainEpoch(extraI)
	}
	// t.CurrentBalance (big.Int) (struct)

	{

		if err := t.CurrentBalance.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.CurrentBalance: %w", err)
		}

	}
	// t.State (chaos.State) (struct)

	{

		if err := t.State.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.State: %w", err)
		}

	}
	return nil
}
