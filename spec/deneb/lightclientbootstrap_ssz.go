// Code generated by fastssz. DO NOT EDIT.
// Hash: 6e2caefc28a9b853b346650cc414ff8babe4e5136fa8ace0b5e5c15ad44a07e7
// Version: 0.1.2
package deneb

import (
	"github.com/attestantio/go-eth2-client/spec/altair"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the LightClientBootstrap object
func (l *LightClientBootstrap) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(l)
}

// MarshalSSZTo ssz marshals the LightClientBootstrap object to a target array
func (l *LightClientBootstrap) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(168)

	// Offset (0) 'Header'
	dst = ssz.WriteOffset(dst, offset)
	if l.Header == nil {
		l.Header = new(LightClientHeader)
	}
	offset += l.Header.SizeSSZ()

	// Offset (1) 'CurrentSyncCommittee'
	dst = ssz.WriteOffset(dst, offset)
	if l.CurrentSyncCommittee == nil {
		l.CurrentSyncCommittee = new(altair.SyncCommittee)
	}
	offset += l.CurrentSyncCommittee.SizeSSZ()

	// Field (2) 'CurrentSyncCommitteeBranch'
	if size := len(l.CurrentSyncCommitteeBranch); size != 5 {
		err = ssz.ErrVectorLengthFn("LightClientBootstrap.CurrentSyncCommitteeBranch", size, 5)
		return
	}
	for ii := 0; ii < 5; ii++ {
		if size := len(l.CurrentSyncCommitteeBranch[ii]); size != 32 {
			err = ssz.ErrBytesLengthFn("LightClientBootstrap.CurrentSyncCommitteeBranch[ii]", size, 32)
			return
		}
		dst = append(dst, l.CurrentSyncCommitteeBranch[ii]...)
	}

	// Field (0) 'Header'
	if dst, err = l.Header.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'CurrentSyncCommittee'
	if dst, err = l.CurrentSyncCommittee.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the LightClientBootstrap object
func (l *LightClientBootstrap) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 168 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1 uint64

	// Offset (0) 'Header'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 168 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'CurrentSyncCommittee'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (2) 'CurrentSyncCommitteeBranch'
	l.CurrentSyncCommitteeBranch = make([][]byte, 5)
	for ii := 0; ii < 5; ii++ {
		if cap(l.CurrentSyncCommitteeBranch[ii]) == 0 {
			l.CurrentSyncCommitteeBranch[ii] = make([]byte, 0, len(buf[8:168][ii*32:(ii+1)*32]))
		}
		l.CurrentSyncCommitteeBranch[ii] = append(l.CurrentSyncCommitteeBranch[ii], buf[8:168][ii*32:(ii+1)*32]...)
	}

	// Field (0) 'Header'
	{
		buf = tail[o0:o1]
		if l.Header == nil {
			l.Header = new(LightClientHeader)
		}
		if err = l.Header.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'CurrentSyncCommittee'
	{
		buf = tail[o1:]
		if l.CurrentSyncCommittee == nil {
			l.CurrentSyncCommittee = new(altair.SyncCommittee)
		}
		if err = l.CurrentSyncCommittee.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the LightClientBootstrap object
func (l *LightClientBootstrap) SizeSSZ() (size int) {
	size = 168

	// Field (0) 'Header'
	if l.Header == nil {
		l.Header = new(LightClientHeader)
	}
	size += l.Header.SizeSSZ()

	// Field (1) 'CurrentSyncCommittee'
	if l.CurrentSyncCommittee == nil {
		l.CurrentSyncCommittee = new(altair.SyncCommittee)
	}
	size += l.CurrentSyncCommittee.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the LightClientBootstrap object
func (l *LightClientBootstrap) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(l)
}

// HashTreeRootWith ssz hashes the LightClientBootstrap object with a hasher
func (l *LightClientBootstrap) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Header'
	if err = l.Header.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'CurrentSyncCommittee'
	if err = l.CurrentSyncCommittee.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'CurrentSyncCommitteeBranch'
	{
		if size := len(l.CurrentSyncCommitteeBranch); size != 5 {
			err = ssz.ErrVectorLengthFn("LightClientBootstrap.CurrentSyncCommitteeBranch", size, 5)
			return
		}
		subIndx := hh.Index()
		for _, i := range l.CurrentSyncCommitteeBranch {
			if len(i) != 32 {
				err = ssz.ErrBytesLength
				return
			}
			hh.Append(i)
		}
		hh.Merkleize(subIndx)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the LightClientBootstrap object
func (l *LightClientBootstrap) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(l)
}
