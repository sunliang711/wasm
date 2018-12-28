package types

import (
	"fmt"
)

type IndexMap struct {
	m         map[uint64]interface{}
	minIndex  uint64
	maxIndex  uint64
	lastIndex uint64
}

func NewIndexMap(minIndex, maxIndex uint64) *IndexMap {
	return &IndexMap{
		m:         make(map[uint64]interface{}),
		minIndex:  minIndex,
		maxIndex:  maxIndex,
		lastIndex: minIndex - 1,
	}
}

func (im *IndexMap) IsFull() bool {
	if uint64(len(im.m)) == im.maxIndex-im.minIndex+1 {
		return true
	}
	return false
}

func (im *IndexMap) checkRange() error {
	if im.lastIndex < im.minIndex || im.lastIndex > im.maxIndex {
		return fmt.Errorf(ErrIndexMapLastIndexOutOfRange)
	}
	return nil
}

func (im *IndexMap) Add(failIndex uint64, val interface{}) (uint64, error) {
	if im.IsFull() {
		return failIndex, fmt.Errorf(ErrIndexMapFull)
	}
	for {
		im.lastIndex += 1
		if im.lastIndex > im.maxIndex {
			im.lastIndex = im.minIndex
		}
		if _, ok := im.m[im.lastIndex]; !ok {
			break
		}
	}
	err := im.checkRange()
	if err != nil {
		return 0, err
	}
	im.m[im.lastIndex] = val
	return im.lastIndex, nil
}

func (im *IndexMap) Insert(index uint64, val interface{}) error {
	err := im.checkRange()
	if err != nil {
		return err
	}
	if _, ok := im.m[index]; !ok {
		im.m[index ] = val
	} else {
		return fmt.Errorf(ErrIndexExist)
	}
	return nil
}

func (im *IndexMap) Remove(index uint64) error {
	err := im.checkRange()
	if err != nil {
		return err
	}
	if _, ok := im.m[index]; !ok {
		return fmt.Errorf(ErrIndexNotExist)
	} else {
		delete(im.m, index)
		return nil
	}
}

func (im *IndexMap) Contains(index uint64) bool {
	err := im.checkRange()
	if err != nil {
		return false
	}
	if _, ok := im.m[index]; !ok {
		return false
	} else {
		return true
	}
}

func (im *IndexMap) Get(index uint64) (interface{}, error) {
	if im.Contains(index) {
		return im.m[index], nil
	} else {
		return nil, fmt.Errorf(ErrIndexNotExist)
	}
}

func (im *IndexMap) Size() uint64 {
	return uint64(len(im.m))
}

func (im *IndexMap) MinIndex() uint64 {
	return im.minIndex
}

func (im *IndexMap) MaxIndex() uint64 {
	return im.maxIndex
}
