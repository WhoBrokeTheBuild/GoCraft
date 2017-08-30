package main

import (
    "os"
    "io"
    "path/filepath"
    "compress/zlib"
    "encoding/binary"
    "bufio"
    "fmt"
    "math"
    "errors"
    "github.com/go-gl/gl/v4.1-core/gl"
)

var (
	ErrListUnknown = errors.New("Lists of unknown type aren't supported")
)

type Chunk struct {
    X    int
    Z    int
    Data [16][16][256]BlockID
    VAO  uint32
	VBOs [2]uint32
}

func (c *Chunk) GetLength() (int) {
    return len(c.Data) * len(c.Data[0]) * len(c.Data[0][0])
}

func (ch *Chunk) Load(dir string, x int, z int, program uint32) (error) {
    mcaName := fmt.Sprintf("r.%v.%v.mca", x>>5, z>>5)
    mcaPath := filepath.Join(dir, "region", mcaName)

    file, err := os.Open(mcaPath)
    if err != nil {
        return err
    }
    defer file.Close()

    ch.X = x
    ch.Z = z

    file.Seek(int64(4*((x&31)+(z&31)*32)), io.SeekStart)

	var (
        location uint32
		length          uint32
		compressionType byte
	)

    err = binary.Read(file, binary.BigEndian, &location)
    if err != nil {
        return err
    }

    file.Seek(int64(4096 * (int(location) >> 8)), io.SeekStart)

    err = binary.Read(file, binary.BigEndian, &length)
    if err != nil {
        return err
    }

    err = binary.Read(file, binary.BigEndian, &compressionType)
    if err != nil {
        return err
    }

    zfile, err := zlib.NewReader(file)
    if err != nil {
        return err
    }
    defer zfile.Close()

    chunkData := new(chunkData)
	chunkData.sections = make([]*sectionData, 0)
	if err := chunkData.parse(NewChunkReader(zfile), false); err != nil {
		return err
	}

    if len(chunkData.sections) != 0 {
		for _, section := range chunkData.sections {
			for i, blockId := range section.blocks {
				var metadata byte
				if i&1 == 1 {
					metadata = section.data[i/2] >> 4
				} else {
					metadata = section.data[i/2] & 0xf
				}
				// Note that the old format is XZY and the new format is YZX
				x, z, y := indexToCoords(i, 16, 16)
                y += 16 * section.y
                // coordsToIndex(x, z, y+16*section.y, 16, 256)
				ch.Data[x][z][y] = BlockID{ int(blockId), int(metadata) }
			}
		}
	} else {
		//if chunkData.blocks != nil && chunkData.data != nil {
		//	for i, blockId := range chunkData.blocks {
		//		var metadata byte
		//		if i&1 == 1 {
		//			metadata = chunkData.data[i/2] >> 4
		//		} else {
		//			metadata = chunkData.data[i/2] & 0xf
		//		}
        //        ch.Data[i] = BlockID{ int(blockId + (metadata << 8)), 0 }
		//	}
		//}
    }

    allBlocks := GetAllBlocks()

	gl.GenVertexArrays(1, &ch.VAO)
	gl.BindVertexArray(ch.VAO)

	gl.GenBuffers(2, &ch.VBOs[0])

	gl.BindBuffer(gl.ARRAY_BUFFER, ch.VBOs[0])
	gl.BufferData(gl.ARRAY_BUFFER, ch.GetLength() * 36 * 3 * 4, nil, gl.DYNAMIC_DRAW)

	gl.BindBuffer(gl.ARRAY_BUFFER, ch.VBOs[1])
    gl.BufferData(gl.ARRAY_BUFFER, ch.GetLength() * 36 * 2 * 4, nil, gl.DYNAMIC_DRAW)

    vertOffset := 0
    texcoordOffset := 0
    for row := range ch.Data {
        for col := range ch.Data[row] {
            for slice := range ch.Data[row][col] {
                id := ch.Data[row][col][slice]
                block, ok := allBlocks[id]
                if !ok {
                    block = allBlocks[BlockID{ 0, 0 }]
                }
                verts, texcoords := block.GetData()

                position := []float32{ float32(row), float32(slice), float32(col) }

                for i := range verts {
                    verts[i] += position[i % 3]
                }

                gl.BindBuffer(gl.ARRAY_BUFFER, ch.VBOs[0])
                gl.BufferSubData(gl.ARRAY_BUFFER, vertOffset, len(verts) * 4, gl.Ptr(verts))
                vertOffset += len(verts) * 4

                gl.BindBuffer(gl.ARRAY_BUFFER, ch.VBOs[1])
                gl.BufferSubData(gl.ARRAY_BUFFER, texcoordOffset, len(texcoords) * 4, gl.Ptr(texcoords))
                texcoordOffset += len(texcoords) * 4
            }
        }
    }

    gl.BindBuffer(gl.ARRAY_BUFFER, ch.VBOs[0])
    vertAttr := uint32(gl.GetAttribLocation(program, gl.Str("in_vert\x00")))
    gl.EnableVertexAttribArray(vertAttr)
    gl.VertexAttribPointer(vertAttr, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

    gl.BindBuffer(gl.ARRAY_BUFFER, ch.VBOs[1])
	texcoordAttr := uint32(gl.GetAttribLocation(program, gl.Str("in_texcoord\x00")))
	gl.EnableVertexAttribArray(texcoordAttr)
	gl.VertexAttribPointer(texcoordAttr, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))

    return nil
}

func indexToCoords(i, aMax, bMax int) (a, b, c int) {
	a = i % aMax
	b = (i / aMax) % bMax
	c = i / (aMax * bMax)
	return
}

func coordsToIndex(a, b, c, bMax, cMax int) int {
	return c + cMax*(b+bMax*a)
}

func yzxToXzy(yzx, xMax, zMax, yMax int) int {
	x := yzx % xMax
	z := (yzx / xMax) % zMax
	y := (yzx / (xMax * zMax)) % yMax

	// yzx := x + xMax*(z + zMax*y)
	xzy := y + yMax*(z+zMax*x)
	return xzy
}

type ChunkTypeId byte

const (
	TagStructEnd ChunkTypeId = 0  // No name. Single zero byte.
	TagInt8      ChunkTypeId = 1  // A single signed byte (8 bits)
	TagInt16     ChunkTypeId = 2  // A signed short (16 bits, big endian)
	TagInt32     ChunkTypeId = 3  // A signed int (32 bits, big endian)
	TagInt64     ChunkTypeId = 4  // A signed long (64 bits, big endian)
	TagFloat32   ChunkTypeId = 5  // A floating point value (32 bits, big endian, IEEE 754-2008, binary32)
	TagFloat64   ChunkTypeId = 6  // A floating point value (64 bits, big endian, IEEE 754-2008, binary64)
	TagByteArray ChunkTypeId = 7  // { TAG_Int length; An array of bytes of unspecified format. The length of this array is <length> bytes }
	TagString    ChunkTypeId = 8  // { TAG_Short length; An array of bytes defining a string in UTF-8 format. The length of this array is <length> bytes }
	TagList      ChunkTypeId = 9  // { TAG_Byte tagId; TAG_Int length; A sequential list of Tags (not Named Tags), of type <typeId>. The length of this array is <length> Tags. } Notes: All tags share the same type.
	TagStruct    ChunkTypeId = 10 // { A sequential list of Named Tags. This array keeps going until a TAG_End is found.; TAG_End end } Notes: If there's a nested TAG_Compound within this tag, that one will also have a TAG_End, so simply reading until the next TAG_End will not work. The names of the named tags have to be unique within each TAG_Compound The order of the tags is not guaranteed.
	TagIntArray  ChunkTypeId = 11 // { TAG_Int length; An array of ints. The length of this array is <length> ints }
)

type ChunkReader struct {
    r *bufio.Reader
}

func NewChunkReader(r io.Reader) *ChunkReader {
    return &ChunkReader{bufio.NewReader(r)}
}

func (r *ChunkReader) ReadTag() (typeId ChunkTypeId, name string, err error) {
	typeId, err = r.readTypeId()
	if err != nil || typeId == 0 {
		return typeId, "", err
	}

	name, err = r.ReadString()
	if err != nil {
		return typeId, name, err
	}

	return typeId, name, nil
}

func (r *ChunkReader) ReadListHeader() (itemTypeId ChunkTypeId, length int, err error) {
	length = 0

	itemTypeId, err = r.readTypeId()
	if err == nil {
		length, err = r.ReadInt32()
	}

	return
}

func (r *ChunkReader) ReadString() (string, error) {
	var length, err1 = r.ReadInt16()
	if err1 != nil {
		return "", err1
	}

	var bytes = make([]byte, length)
	var _, err = io.ReadFull(r.r, bytes)
	return string(bytes), err
}

func (r *ChunkReader) ReadByteList() ([]byte, error) {
	var length, err1 = r.ReadInt32()
	if err1 != nil {
		return nil, err1
	}

	var bytes = make([]byte, length)
	var _, err = io.ReadFull(r.r, bytes)
	return bytes, err
}

func (r *ChunkReader) ReadIntList() ([]int, error) {
	length, err := r.ReadInt32()
	if err != nil {
		return nil, err
	}

	ints := make([]int, length)
	for i := 0; i < length; i++ {
		ints[i], err = r.ReadInt32()
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}

func (r *ChunkReader) ReadInt8() (int, error) {
	return r.readIntN(1)
}

func (r *ChunkReader) ReadInt16() (int, error) {
	return r.readIntN(2)
}

func (r *ChunkReader) ReadInt32() (int, error) {
	return r.readIntN(4)
}

func (r *ChunkReader) ReadInt64() (int, error) {
	return r.readIntN(8)
}

func (r *ChunkReader) ReadFloat32() (float32, error) {
	x, err := r.readUintN(4)
	return math.Float32frombits(uint32(x)), err
}

func (r *ChunkReader) ReadFloat64() (float64, error) {
	x, err := r.readUintN(8)
	return math.Float64frombits(x), err
}

func (r *ChunkReader) readTypeId() (ChunkTypeId, error) {
    id, err := r.r.ReadByte()
    return ChunkTypeId(id), err
}

func (r *ChunkReader) readIntN(n int) (int, error) {
	var a int = 0

	for i := 0; i < n; i++ {
		var b, err = r.r.ReadByte()
		if err != nil {
			return a, err
		}
		a = a<<8 + int(b)
	}

	return a, nil
}

func (r *ChunkReader) readUintN(n int) (uint64, error) {
	var a uint64 = 0

	for i := 0; i < n; i++ {
		var b, err = r.r.ReadByte()
		if err != nil {
			return a, err
		}
		a = a<<8 + uint64(b)
	}

	return a, nil
}

func (r *ChunkReader) ReadStruct() (map[string]interface{}, error) {
	s := make(map[string]interface{})
	for {
		typeId, name, err := r.ReadTag()
		if err != nil {
			return s, err
		}
		if typeId == TagStructEnd {
			break
		}
		x, err := r.ReadValue(typeId)
		s[name] = x
		if err != nil {
			return s, err
		}
	}
	return s, nil
}

func (r *ChunkReader) ReadValue(typeId ChunkTypeId) (interface{}, error) {
	switch typeId {
	case TagStruct:
		return r.ReadStruct()
	case TagStructEnd:
		return nil, nil
	case TagByteArray:
		return r.ReadByteList()
	case TagInt8:
		return r.ReadInt8()
	case TagInt16:
		return r.ReadInt16()
	case TagInt32:
		return r.ReadInt32()
	case TagInt64:
		return r.ReadInt64()
	case TagFloat32:
		return r.ReadFloat32()
	case TagFloat64:
		return r.ReadFloat64()
	case TagString:
		return r.ReadString()
	case TagList:
		itemTypeId, length, err := r.ReadListHeader()
		if err != nil {
			return nil, err
		}
		switch ChunkTypeId(itemTypeId) {
		case TagInt8:
			list := make([]int, length)
			for i := 0; i < length; i++ {
				x, err := r.ReadInt8()
				list[i] = x
				if err != nil {
					return list, err
				}
			}
			return list, nil
		case TagFloat32:
			list := make([]float32, length)
			for i := 0; i < length; i++ {
				x, err := r.ReadFloat32()
				list[i] = x
				if err != nil {
					return list, err
				}
			}
			return list, nil
		case TagFloat64:
			list := make([]float64, length)
			for i := 0; i < length; i++ {
				x, err := r.ReadFloat64()
				list[i] = x
				if err != nil {
					return list, err
				}
			}
			return list, nil
		case TagStruct:
			list := make([]interface{}, length)
			for i := 0; i < length; i++ {
				s := make(map[string]interface{})
				s, err := r.ReadStruct()
				list[i] = s
				if err != nil {
					return list, err
				}
			}
			return list, nil
		default:
			return nil, errors.New(fmt.Sprintf("reading lists of typeId %d not supported. length:%d", itemTypeId, length))
		}
	}

	return nil, errors.New(fmt.Sprintf("reading typeId %d not supported", typeId))
}

type chunkData struct {
	xPos, zPos int
	blocks     []byte
	data       []byte
	section    *sectionData
	sections   []*sectionData
}


type sectionData struct {
	y      int
	blocks []byte
	data   []byte
}

func (chunk *chunkData) parse(r *ChunkReader, listStruct bool) error {
	structDepth := 0
	if listStruct {
		structDepth++
	}

	for {
		typeId, name, err := r.ReadTag()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		switch typeId {
		case TagStruct:
			structDepth++
		case TagStructEnd:
			structDepth--
			if structDepth == 0 {
				return nil
			}
		case TagByteArray:
			bytes, err := r.ReadByteList()
			if err != nil {
				return err
			}
			if name == "Blocks" {
				if chunk.section != nil {
					chunk.section.blocks = bytes
				} else {
					chunk.blocks = bytes
				}
			} else if name == "Data" {
				if chunk.section != nil {
					chunk.section.data = bytes
				} else {
					chunk.data = bytes
				}
			}
		case TagIntArray:
			_, err := r.ReadIntList()
			if err != nil {
				return err
			}
		case TagInt8:
			number, err := r.ReadInt8()
			if err != nil {
				return err
			}
			if name == "Y" {
				chunk.section.y = int(number)
			}
		case TagInt16:
			_, err := r.ReadInt16()
			if err != nil {
				return err
			}
		case TagInt32:
			number, err := r.ReadInt32()
			if err != nil {
				return err
			}

			if name == "xPos" {
				chunk.xPos = number
			}
			if name == "zPos" {
				chunk.zPos = number
			}
		case TagInt64:
			_, err := r.ReadInt64()
			if err != nil {
				return err
			}
		case TagFloat32:
			_, err := r.ReadFloat32()
			if err != nil {
				return err
			}
		case TagFloat64:
			_, err := r.ReadFloat64()
			if err != nil {
				return err
			}
		case TagString:
			_, err := r.ReadString()
			if err != nil {
				return err
			}
		case TagList:
			itemTypeId, length, err := r.ReadListHeader()
			if err != nil {
				return err
			}
			switch itemTypeId {
			case TagInt8:
				for i := 0; i < length; i++ {
					_, err := r.ReadInt8()
					if err != nil {
						return err
					}
				}
			case TagFloat32:
				for i := 0; i < length; i++ {
					_, err := r.ReadFloat32()
					if err != nil {
						return err
					}
				}
			case TagFloat64:
				for i := 0; i < length; i++ {
					_, err := r.ReadFloat64()
					if err != nil {
						return err
					}
				}
			case TagStruct:
				for i := 0; i < length; i++ {
					if name == "Sections" {
						chunk.section = new(sectionData)
						chunk.sections = append(chunk.sections, chunk.section)
					}
					err := chunk.parse(r, true)
					if err != nil {
						return err
					}
				}
            case TagStructEnd:

			default:
				fmt.Printf("# %s list todo(%v) %v\n", name, itemTypeId, length)
				return ErrListUnknown
			}
		default:
			fmt.Printf("# %s todo(%d)\n", name, typeId)
		}
	}

	return nil
}
