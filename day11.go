package main

import (
	"sort"
	"strconv"
	"strings"
)

func day11_1(input string) string {
	rows := strings.Split(input, "\n")

	type monkeyLogic struct {
		items                 []int
		itemsInspected        int
		operationSymbol       string
		operationValue        int
		operationReferenceOld bool
		testDivider           int
		ifTrue, ifFalse       int
	}

	monkeys := make([]monkeyLogic, 0)
	currentMonkey := monkeyLogic{}
	for _, r := range rows {
		if strings.HasPrefix(r, "Monkey ") {
			currentMonkey = monkeyLogic{}
		}

		if strings.HasPrefix(r, "  Starting items:") {
			fields := strings.Split(r, ": ")
			items := strings.Split(fields[1], ", ")
			for _, i := range items {
				itemVal, err := strconv.ParseInt(i, 10, 64)
				if err != nil {
					panic(err)
				}
				currentMonkey.items = append(currentMonkey.items, int(itemVal))
			}
		}

		if strings.HasPrefix(r, "  Operation: new = old") {
			fields := strings.Split(r, " old ")
			operation := strings.Split(fields[1], " ")
			currentMonkey.operationSymbol = operation[0]

			if operation[1] == "old" {
				currentMonkey.operationReferenceOld = true
			} else {
				currentMonkey.operationReferenceOld = false
				opVal, err := strconv.ParseInt(operation[1], 10, 64)
				if err != nil {
					panic(err)
				}
				currentMonkey.operationValue = int(opVal)
			}

		}

		if strings.HasPrefix(r, "  Test: divisible by") {
			fields := strings.Split(r, " by ")
			divBy, err := strconv.ParseInt(fields[1], 10, 64)
			if err != nil {
				panic(err)
			}
			currentMonkey.testDivider = int(divBy)
		}

		if strings.HasPrefix(r, "    If true: throw to monkey") {
			fields := strings.Split(r, " monkey ")
			toMonkey, err := strconv.ParseInt(fields[1], 10, 64)
			if err != nil {
				panic(err)
			}
			currentMonkey.ifTrue = int(toMonkey)
		}

		if strings.HasPrefix(r, "    If false: throw to monkey") {
			fields := strings.Split(r, " monkey ")
			toMonkey, err := strconv.ParseInt(fields[1], 10, 64)
			if err != nil {
				panic(err)
			}
			currentMonkey.ifFalse = int(toMonkey)
			monkeys = append(monkeys, currentMonkey)
		}
	}

	for i := 0; i < 20; i++ {
		for monkey := 0; monkey < len(monkeys); monkey++ {
			for _, item := range monkeys[monkey].items {
				monkeys[monkey].itemsInspected++
				worryLevel := 0
				switch monkeys[monkey].operationSymbol {
				case "*":
					worryLevel = item * monkeys[monkey].operationValue
					if monkeys[monkey].operationReferenceOld {
						worryLevel = item * item
					}
				case "+":
					worryLevel = item + monkeys[monkey].operationValue
					if monkeys[monkey].operationReferenceOld {
						worryLevel = item + item
					}
				}

				worryLevel = worryLevel / 3

				if worryLevel%monkeys[monkey].testDivider == 0 {
					monkeys[monkeys[monkey].ifTrue].items = append(monkeys[monkeys[monkey].ifTrue].items, worryLevel)
				} else {
					monkeys[monkeys[monkey].ifFalse].items = append(monkeys[monkeys[monkey].ifFalse].items, worryLevel)
				}
			}
			monkeys[monkey].items = make([]int, 0)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].itemsInspected > monkeys[j].itemsInspected
	})

	return strconv.Itoa(monkeys[0].itemsInspected * monkeys[1].itemsInspected)
}

func day11_2(input string) string {
	rows := strings.Split(input, "\n")

	type monkeyLogic struct {
		items                 []int
		itemsInspected        int
		operationSymbol       string
		operationValue        int
		operationReferenceOld bool
		testDivider           int
		ifTrue, ifFalse       int
	}

	monkeys := make([]monkeyLogic, 0)
	currentMonkey := monkeyLogic{}
	for _, r := range rows {
		if strings.HasPrefix(r, "Monkey ") {
			currentMonkey = monkeyLogic{}
		}

		if strings.HasPrefix(r, "  Starting items:") {
			fields := strings.Split(r, ": ")
			items := strings.Split(fields[1], ", ")
			for _, i := range items {
				itemVal, err := strconv.ParseInt(i, 10, 64)
				if err != nil {
					panic(err)
				}
				currentMonkey.items = append(currentMonkey.items, int(itemVal))
			}
		}

		if strings.HasPrefix(r, "  Operation: new = old") {
			fields := strings.Split(r, " old ")
			operation := strings.Split(fields[1], " ")
			currentMonkey.operationSymbol = operation[0]

			if operation[1] == "old" {
				currentMonkey.operationReferenceOld = true
			} else {
				currentMonkey.operationReferenceOld = false
				opVal, err := strconv.ParseInt(operation[1], 10, 64)
				if err != nil {
					panic(err)
				}
				currentMonkey.operationValue = int(opVal)
			}

		}

		if strings.HasPrefix(r, "  Test: divisible by") {
			fields := strings.Split(r, " by ")
			divBy, err := strconv.ParseInt(fields[1], 10, 64)
			if err != nil {
				panic(err)
			}
			currentMonkey.testDivider = int(divBy)
		}

		if strings.HasPrefix(r, "    If true: throw to monkey") {
			fields := strings.Split(r, " monkey ")
			toMonkey, err := strconv.ParseInt(fields[1], 10, 64)
			if err != nil {
				panic(err)
			}
			currentMonkey.ifTrue = int(toMonkey)
		}

		if strings.HasPrefix(r, "    If false: throw to monkey") {
			fields := strings.Split(r, " monkey ")
			toMonkey, err := strconv.ParseInt(fields[1], 10, 64)
			if err != nil {
				panic(err)
			}
			currentMonkey.ifFalse = int(toMonkey)
			monkeys = append(monkeys, currentMonkey)
		}
	}

	// calculate prime field
	pField := 1
	for _, m := range monkeys {
		pField *= m.testDivider
	}

	for i := 0; i < 10_000; i++ {
		for monkey := 0; monkey < len(monkeys); monkey++ {
			for _, item := range monkeys[monkey].items {
				monkeys[monkey].itemsInspected++
				worryLevel := 0
				switch monkeys[monkey].operationSymbol {
				case "*":
					worryLevel = item * monkeys[monkey].operationValue
					if monkeys[monkey].operationReferenceOld {
						worryLevel = item * item
					}
				case "+":
					worryLevel = item + monkeys[monkey].operationValue
					if monkeys[monkey].operationReferenceOld {
						worryLevel = item + item
					}
				}

				worryLevel = worryLevel % pField

				if worryLevel%monkeys[monkey].testDivider == 0 {
					monkeys[monkeys[monkey].ifTrue].items = append(monkeys[monkeys[monkey].ifTrue].items, worryLevel)
				} else {
					monkeys[monkeys[monkey].ifFalse].items = append(monkeys[monkeys[monkey].ifFalse].items, worryLevel)
				}
			}
			monkeys[monkey].items = make([]int, 0)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].itemsInspected > monkeys[j].itemsInspected
	})

	return strconv.Itoa(monkeys[0].itemsInspected * monkeys[1].itemsInspected)
}
