package model

import (
	"bytes"
	"encoding/binary"
	"math"
)

type Model struct {
	Header    ModelHeader
	ModelData ModelData
	Checksum  uint64
}

type ModelHeader struct {
	Magic       string
	VertexCount uint32
	IndexCount  uint32
}

type ModelData struct {
	Positions [][3]float64
	Normals   [][3]float64
	TexCoords [][2]float64
	Indices   [][]int
}

func  NewModel() *Model {
	return &Model{
		Header:    ModelHeader{},
		ModelData: ModelData{},
		Checksum:  0,
	}
}

func (mdl *Model) AddPosition(pos [3]float64) {
  mdl.ModelData.Positions = append(mdl.ModelData.Positions, pos)
	for _, value := range pos {
		mdl.Checksum += math.Float64bits(value)
	}
}

func (mdl *Model) AddNormal(pos [3]float64) {
	mdl.ModelData.Normals = append(mdl.ModelData.Normals, pos)
	for _, value := range pos {
		mdl.Checksum += math.Float64bits(value)
	}
}

func (mdl *Model) AddTexCoord(pos [2]float64) {
	mdl.ModelData.TexCoords = append(mdl.ModelData.TexCoords, pos)
	for _, value := range pos {
		mdl.Checksum += math.Float64bits(value)
	}
}

func (mdl *Model) AddIndices(pos []int) {
	mdl.ModelData.Indices = append(mdl.ModelData.Indices, pos)
	for _, value := range pos {
		mdl.Checksum += uint64(value)
	}
}

func (mdl *Model) Bytes() []byte {
  buf := new(bytes.Buffer)
  mdl.Header.VertexCount = uint32(len(mdl.ModelData.Positions))
  mdl.Header.IndexCount  = uint32(len(mdl.ModelData.Indices))
 
  binary.Write(buf, binary.BigEndian, []byte("MDL01"));

  binary.Write(buf, binary.BigEndian, mdl.Header.VertexCount)
  binary.Write(buf, binary.BigEndian, mdl.Header.IndexCount)
  binary.Write(buf, binary.BigEndian, mdl.Checksum)


  for _, v := range mdl.ModelData.Positions {
    for _, p := range v {
      binary.Write(buf, binary.BigEndian, p)
    }
  }

  return buf.Bytes()
}



