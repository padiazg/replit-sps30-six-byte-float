package main

import (
	"encoding/binary"
	"fmt"
	"math"
)

// AirQualityReading data readed from sensor
type AirQualityReading struct {
	// all big-endian, IEEE754 float value
	massPM1             float32 // Mass Concentration PM1.0 [μg/m3]
	massPM25            float32 // Mass Concentration PM2.5 [μg/m3]
	massPM4             float32 // Mass Concentration PM4.0 [μg/m3]
	massPM10            float32 // Mass Concentration PM10 [μg/m3]
	numberPM05          float32 // Number Concentration PM0.5 [#/cm3]
	numberPM1           float32 // Number Concentration PM1.0 [#/cm3]
	numberPM25          float32 // Number Concentration PM2.5 [#/cm3]
	numberPM4           float32 // Number Concentration PM4.0 [#/cm3]
	numberPM10          float32 // Number Concentration PM10 [#/cm3]
	typicalParticleSize float32 // Typical Particle Size [μm]
}

func byteArrayToFloat32(data []byte) float32 {
	a0 := binary.BigEndian.Uint32(data)
	a1 := math.Float32frombits(a0)
	return a1
} 

func main() {
	// var b0 = []byte{0x42, 0x4f, 0xc2, 0x05, 0x03, 0xa5}
  // var i0 = []byte{b0[4], b0[3], b0[1], b0[0]}
  // var i1 = []byte{b0[0], b0[1], b0[3], b0[4]}

	// a0 := binary.LittleEndian.Uint32(i0)
	// a2 := math.Float32frombits(a0)
	// fmt.Println(a2)

  // a3 := binary.BigEndian.Uint32(i1)
  // a4 := math.Float32frombits(a3)
  // fmt.Println(a4)

  // var f0 float32 = 51.754894
  // b1 := Float32bytes(f0)
  // fmt.Printf("%X\n", b1)

// pm0.5 count: 17.551914
// pm1   count: 36.319 ug: 6.074
// pm2.5 count: 47.804 ug: 16.789
// pm4   count: 50.137 ug: 25.178
// pm10  count: 50.473 ug: 26.856
// pm_typ: 1.281431

  b0 := []byte{0x40,0xc2,0x2d,0x5f,0xce,0xa7,0x41,0x86,0x20,0x4f,0xaf,0x43,0x41,0xc9,0x33,0x6c,0xb0,0xdf,0x41,0xd6,0x5e,0xd8,0xe1,0x82,0x41,0x8c,0xfb,0x6a,0x52,0x26,0x42,0x11,0xa3,0x47,0x09,0x2e,0x42,0x3f,0x3a,0x37,0xa1,0x50,0x42,0x48,0x55,0x8c,0xb8,0x10,0x42,0x49,0x64,0xe4,0x68,0x76,0x3f,0xa4,0x92,0x05,0xf2}

	var measurement AirQualityReading
	measurement.massPM1 = byteArrayToFloat32([]byte{b0[0], b0[1], b0[3], b0[4]})
	measurement.massPM25 = byteArrayToFloat32([]byte{b0[6], b0[7], b0[9], b0[10]})
	measurement.massPM4 = byteArrayToFloat32([]byte{b0[12], b0[13], b0[15], b0[16]})
	measurement.massPM10 = byteArrayToFloat32([]byte{b0[18], b0[19], b0[21], b0[22]})
	measurement.numberPM05 = byteArrayToFloat32([]byte{b0[24], b0[25], b0[27], b0[27]})
	measurement.numberPM1 = byteArrayToFloat32([]byte{b0[30], b0[31], b0[33], b0[34]})
	measurement.numberPM25 = byteArrayToFloat32([]byte{b0[36], b0[37], b0[39], b0[40]})
	measurement.numberPM4 = byteArrayToFloat32([]byte{b0[42], b0[43], b0[45], b0[46]})
	measurement.numberPM10 = byteArrayToFloat32([]byte{b0[48], b0[49], b0[51], b0[52]})
	measurement.typicalParticleSize = byteArrayToFloat32([]byte{b0[54], b0[55], b0[57], b0[58]})
  fmt.Printf("%v", measurement)
}
