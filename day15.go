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

type chunk struct {
	fromX, toX int64
	checked    bool
}

func day15_2(input string) string {
	analyzeMin := int64(0)
	analyzeMax := int64(4000000)
	//analyzeMax := int64(20)
	re := regexp.MustCompile("Sensor at x=(-?\\d+), y=(-?\\d+): closest beacon is at x=(-?\\d+), y=(-?\\d+)")

	type beaconSensor struct {
		beaconX, beaconY, sensorX, sensorY int64
	}

	rows := strings.Split(input, "\n")
	beaconSensors := make([]beaconSensor, len(rows))
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
	}

	chunks := make([]chunk, len(beaconSensors))
	for analyzeRow := analyzeMin; analyzeRow <= analyzeMax; analyzeRow++ {
		chunks = make([]chunk, len(beaconSensors)) //reset
		for cnt, bs := range beaconSensors {
			distance := manhattan(bs.sensorX, bs.sensorY, bs.beaconX, bs.beaconY)
			diff := int64(math.Abs(float64(analyzeRow - bs.sensorY)))
			if diff <= distance {
				coverageFrom := bs.sensorX - (distance - diff)
				coverageTo := bs.sensorX + (distance - diff)
				chunks[cnt] = chunk{
					fromX: coverageFrom,
					toX:   coverageTo,
				}
			}
		}

		res := findNonOverlapping(analyzeMax, 0, chunks)
		if res >= 0 {
			return strconv.Itoa(int(4000000*res + analyzeRow))
		}
	}

	return ""
}

func findNonOverlapping(to, check int64, chunks []chunk) int64 {
	newTo := int64(0)
	foundAny := false
	for cnt, c := range chunks {
		if c.checked {
			continue
		}
		if c.fromX <= check {
			if c.toX >= newTo {
				newTo = c.toX
			}

			if newTo >= to {
				return -1
			}
			foundAny = true
			chunks[cnt].checked = true
		}
	}

	if !foundAny {
		return check + 1
	}
	return findNonOverlapping(to, newTo, chunks)
}

func manhattan(x1, y1, x2, y2 int64) int64 {
	out := math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2))
	return int64(out)
}
