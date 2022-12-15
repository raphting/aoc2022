package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func day15_1(input string) string {
	//analyzeRow := int64(10)
	analyzeRow := int64(2000000)
	re := regexp.MustCompile("Sensor at x=(-?\\d+), y=(-?\\d+): closest beacon is at x=(-?\\d+), y=(-?\\d+)")

	type beaconSensor struct {
		beaconX, beaconY, sensorX, sensorY int64
	}

	rows := strings.Split(input, "\n")
	beaconSensors := make([]beaconSensor, len(rows))
	beaconAffectedRowX := make([]int64, 0)
	for cnt, r := range rows {
		found := re.FindStringSubmatch(r)
		sX, _ := strconv.ParseInt(found[1], 10, 64)
		sY, _ := strconv.ParseInt(found[2], 10, 64)
		bX, _ := strconv.ParseInt(found[3], 10, 64)
		bY, _ := strconv.ParseInt(found[4], 10, 64)
		beaconSensors[cnt] = beaconSensor{
			beaconX: bX,
			beaconY: bY,
			sensorX: sX,
			sensorY: sY,
		}
		if bY == analyzeRow {
			beaconAffectedRowX = append(beaconAffectedRowX, bX)
		}
	}

	affectedRow := make(map[int64]bool)
	for _, bs := range beaconSensors {
		distance := manhattan(bs.sensorX, bs.sensorY, bs.beaconX, bs.beaconY)
		diff := int64(math.Abs(float64(analyzeRow - bs.sensorY)))
		if diff <= distance {
			for i := int64(0); i <= distance-diff; i++ {
				coverageA := bs.sensorX + i
				coverageB := bs.sensorX - i
				notA := true
				notB := true
				for _, beaconPlaced := range beaconAffectedRowX {
					if beaconPlaced == coverageA {
						notA = false
					}
					if beaconPlaced == coverageB {
						notB = false
					}
				}
				if notA {
					affectedRow[coverageA] = true
				}
				if notB {
					affectedRow[coverageB] = true
				}
			}
		}
	}

	return strconv.Itoa(len(affectedRow))
}

func manhattan(x1, y1, x2, y2 int64) int64 {
	out := math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2))
	return int64(out)
}
