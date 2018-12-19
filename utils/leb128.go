package utils

import (
	"fmt"
	"io"
	"reflect"
	"wasm/types"
)

func DecodeUInt32(rd io.Reader) (uint32, error) {
	buf := make([]byte, 1)
	maxBits := 32
	maxBytes := (maxBits + 6) / 7
	numBytes := 0
	bytes := make([]byte, maxBytes)
	for numBytes < maxBytes {
		_, err := rd.Read(buf)
		if err != nil {
			return 0, err
		}

		byt := buf[0]
		bytes[numBytes] = byt
		numBytes += 1
		if byt&0x80 == 0 {
			break
		}
	}

	var result uint32
	for i := 0; i < maxBytes; i++ {
		result |= uint32(bytes[i]&0x7f) << uint32(i*7)
	}

	return result, nil
}

func DecodeInt32(rd io.Reader) (int32, error) {
	buf := make([]byte, 1)
	maxBits := 32
	maxBytes := (maxBits + 6) / 7
	numBytes := 0
	bytes := make([]byte, maxBytes)

	signExtendShift := 32
	for numBytes < maxBytes {
		_, err := rd.Read(buf)
		if err != nil {
			return 0, err
		}

		byt := buf[0]
		bytes[numBytes] = byt
		numBytes += 1
		signExtendShift -= 7
		if byt&0x80 == 0 {
			break
		}
	}

	var result int32
	for i := 0; i < maxBytes; i++ {
		result |= int32(bytes[i]&0x7f) << uint32(i*7)
	}
	result = (result << uint(signExtendShift)) >> uint(signExtendShift)
	return result, nil
}

func DecodeUInt64(rd io.Reader) (uint64, error) {
	buf := make([]byte, 1)
	maxBits := 64
	maxBytes := (maxBits + 6) / 7
	numBytes := 0
	bytes := make([]byte, maxBytes)
	for numBytes < maxBytes {
		_, err := rd.Read(buf)
		if err != nil {
			return 0, err
		}

		byt := buf[0]
		bytes[numBytes] = byt
		numBytes += 1
		if byt&0x80 == 0 {
			break
		}
	}

	var result uint64
	for i := 0; i < maxBytes; i++ {
		result |= uint64(bytes[i]&0x7f) << uint64(i*7)
	}

	return result, nil
}

func DecodeU1(rd io.Reader) (byte, error) {
	buf := make([]byte, 1)
	maxBits := 1
	maxBytes := (maxBits + 6) / 7
	numBytes := 0
	bytes := make([]byte, maxBytes)
	for numBytes < maxBytes {
		_, err := rd.Read(buf)
		if err != nil {
			return 0, err
		}

		byt := buf[0]
		bytes[numBytes] = byt
		numBytes += 1
		if byt&0x80 == 0 {
			break
		}
	}

	var result byte
	for i := 0; i < maxBytes; i++ {
		result |= byte(bytes[i]&0x7f) << byte(i*7)
	}

	return result, nil
}

func DecodeVarInt(rd io.Reader, maxBits int, value interface{}) (int, error) {
	maxBytes := (maxBits + 6) / 7
	bytes := make([]byte, maxBytes)
	numBytes := 0

	signExtendShift := 0
	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return numBytes, fmt.Errorf(types.ErrNotPtr)
	}
	isSign := false
	switch rv.Elem().Kind() {
	case reflect.Int8:
		signExtendShift = 8
		isSign = true
	case reflect.Int16:
		signExtendShift = 16
		isSign = true
	case reflect.Int32:
		signExtendShift = 32
		isSign = true
	case reflect.Int64:
		signExtendShift = 64
		isSign = true
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	default:
		return numBytes, fmt.Errorf(types.ErrIntPtr)
	}
	buf := make([]byte, 1)
	for numBytes < maxBytes {
		_, err := rd.Read(buf)
		if err != nil {
			return numBytes, err
		}
		byt := buf[0]
		bytes[numBytes] = byt
		numBytes += 1
		signExtendShift -= 7
		if byt&0x80 == 0 {
			break
		}
	}

	if isSign {
		switch rv.Elem().Kind() {
		case reflect.Int8:
			var result int8
			for i := 0; i < maxBytes; i++ {
				result |= int8((bytes[i] & 0x7f) << uint(i*7))
			}
			result = (result << uint(signExtendShift)) >> uint(signExtendShift)
			rv.Elem().SetInt(int64(result))
		case reflect.Int16:
			var result int16
			for i := 0; i < maxBytes; i++ {
				result |= int16((bytes[i] & 0x7f) << uint(i*7))
			}
			result = (result << uint(signExtendShift)) >> uint(signExtendShift)
			rv.Elem().SetInt(int64(result))
		case reflect.Int32:
			var result int32
			for i := 0; i < maxBytes; i++ {
				result |= int32((bytes[i] & 0x7f) << uint(i*7))
			}
			result = (result << uint(signExtendShift)) >> uint(signExtendShift)
			rv.Elem().SetInt(int64(result))
		case reflect.Int64:
			var result int64
			for i := 0; i < maxBytes; i++ {
				result |= int64((bytes[i] & 0x7f) << uint(i*7))
			}
			result = (result << uint(signExtendShift)) >> uint(signExtendShift)
			rv.Elem().SetInt(int64(result))
		}
	} else {
		var result uint64
		for i := 0; i < maxBytes; i++ {
			result |= uint64(bytes[i]&0x7f) << uint64(i*7)
		}
		rv.Elem().SetUint(result)
	}
	return numBytes, nil
	//I8 signExtendShift = (I8)sizeof(Value) * 8;
	//while(numBytes < maxBytes)
	//{
	//	U8 byte = *stream.advance(1);
	//	bytes[numBytes] = byte;
	//	++numBytes;
	//	signExtendShift -= 7;
	//	if(!(byte & 0x80)) { break; }
	//};
	//
	//// Ensure that the input does not encode more than maxBits of data.
	//enum
	//{
	//	numUsedBitsInLastByte = maxBits - (maxBytes - 1) * 7,
	//		numUnusedBitsInLast = 8 - numUsedBitsInLastByte,
	//	lastBitUsedMask = U8(1 << (numUsedBitsInLastByte - 1)),
	//	lastByteUsedMask = U8(1 << numUsedBitsInLastByte) - U8(1),
	//	lastByteSignedMask = U8(~U8(lastByteUsedMask) & ~U8(0x80))
	//};
	//const U8 lastByte = bytes[maxBytes - 1];
	//if(!std::is_signed<Value>::value)
	//{
	//if((lastByte & ~lastByteUsedMask) != 0)
	//{
	//throw FatalSerializationException(
	//"Invalid unsigned LEB encoding: unused bits in final byte must be 0");
	//}
	//}
	//else
	//{
	//const I8 signBit = I8((lastByte & lastBitUsedMask) << numUnusedBitsInLast);
	//const I8 signExtendedLastBit = signBit >> numUnusedBitsInLast;
	//if((lastByte & ~lastByteUsedMask) != (signExtendedLastBit & lastByteSignedMask))
	//{
	//throw FatalSerializationException(
	//"Invalid signed LEB encoding: unused bits in final byte must match the "
	//"most-significant used bit");
	//}
	//}
	//
	//// Decode the buffer's bytes into the output integer.
	//value = 0;
	//for(Uptr byteIndex = 0; byteIndex < maxBytes; ++byteIndex)
	//{ value |= Value(U64(bytes[byteIndex] & ~0x80) << U64(byteIndex * 7)); }
	//
	//// Sign extend the output integer to the full size of Value.
	//if(std::is_signed<Value>::value && signExtendShift > 0)
	//{ value = Value(value << signExtendShift) >> signExtendShift; }

}
