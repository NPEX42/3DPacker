package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

  


	bytes, err := os.ReadFile("res/teapot.obj")
	if err != nil {
		fmt.Printf("Failed To Open TestCube.obj: %s\n", err)
	}

	source := string(bytes)


	lines := strings.Split(source, "\n")


  var faces [][]int
  var positions [][3]float64
  var normals [][3]float64
  var uvs [][2]float64


	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		parts := strings.Split(line, " ")
    switch parts[0] {
		case "v":
			v := ParseVertex(parts[1:])
			//fmt.Printf("Vertex: %v\n", v)
      positions = append(positions, v)
    case "vn":
      normal := ParseVertex(parts[1:])
      //fmt.Printf("Normal: %v\n", normal)
      normals = append(normals, normal)  
    case "vt":
      uv := ParseVertex2D(parts[1:])
      //fmt.Printf("TexCoord: %v\n", uv)
      uvs = append(uvs, uv) 

		case "f":
      face := ParseTriangleVerts(parts[1:])
			fmt.Printf("Face: %v\n", face)
      faces = append(faces, face)
		}
	}

  var s strings.Builder

  for _, face := range faces {
    for _, index := range face {
      s.WriteString(fmt.Sprintf("%d ", index))
    }
  }
  s.WriteString("\n");

  for _, position := range positions {
    for _, axis_value := range position {
      s.WriteString(fmt.Sprintf("%02.5f ", axis_value))
    }
    s.WriteRune(';')
  }
  s.WriteString("\n");

  for _, normal := range normals {
    for _, axis_value := range normal {
      s.WriteString(fmt.Sprintf("%02.5f ", axis_value))
    }
    s.WriteRune(';')
  }
  s.WriteString("\n");

  for _, normal := range uvs {
    for _, axis_value := range normal {
      s.WriteString(fmt.Sprintf("%02.5f ", axis_value))
    }
    s.WriteRune(';')
  }
  s.WriteString("\n");



  os.WriteFile("TestCube.mdl", []byte(s.String()), os.ModePerm);

}

func ParseVertex(parts []string) [3]float64 {
	x, _ := strconv.ParseFloat(parts[0], 32)
	y, _ := strconv.ParseFloat(parts[1], 32)
	z, _ := strconv.ParseFloat(parts[2], 32)

	return [3]float64{
		x, y, z,
	}
}

func ParseVertex2D(parts []string) [2]float64 {
	x, _ := strconv.ParseFloat(parts[0], 32)
	y, _ := strconv.ParseFloat(parts[1], 32)

	return [2]float64{
		x, y,
	}
}

func ParseTriangleVerts(parts []string) []int {
	var verts []int
	for i := 0; i < 3; i++ {
items := strings.Split(parts[i], "/")
		index, err := strconv.Atoi(items[0])
		if err != nil {
			continue
		}
		verts = append(verts, index - 1)
	}

	return verts
}
