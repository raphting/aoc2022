package main

import (
	"sort"
	"strconv"
	"strings"
)

type messageItem struct {
	isList  bool
	value   int
	subList []*messageItem
}

type pairMessage struct {
	a, b []*messageItem
}

func day13_1(input string) string {
	pairs := strings.Split(input, "\n\n")

	pairMessages := make([]pairMessage, len(pairs))
	for cnt, p := range pairs {
		rows := strings.Split(p, "\n")
		a := strings.Split(rows[0], "")
		b := strings.Split(rows[1], "")

		_, pairMessages[cnt].a = parseMessage(a)
		_, pairMessages[cnt].b = parseMessage(b)
	}

	result := 0
	for cnt, messages := range pairMessages {
		if isInOrder(messages.a, messages.b) {
			result += cnt + 1
		}
	}

	return strconv.Itoa(result)
}

func day13_2(input string) string {
	rows := strings.Split(input, "\n")
	allMessages := make([]string, 0)
	for _, r := range rows {
		if r != "" {
			allMessages = append(allMessages, r)
		}
	}
	allMessages = append(allMessages, "[[2]]")
	allMessages = append(allMessages, "[[6]]")

	sort.Slice(allMessages, func(i, j int) bool {
		_, a := parseMessage(strings.Split(allMessages[i], ""))
		_, b := parseMessage(strings.Split(allMessages[j], ""))
		return isInOrder(a, b)
	})

	result := 1
	for cnt, a := range allMessages {
		if a == "[[2]]" || a == "[[6]]" {
			result *= cnt + 1
		}
	}

	return strconv.Itoa(result)
}

// was used to check correct parsing of input
func reversePrint(m []*messageItem) string {
	out := ""
	for cnt, e := range m {
		if e.isList {
			out += "["
			out += reversePrint(e.subList)
			out += "]"
			if cnt != len(m)-1 {
				out += ","
			}
		} else {
			out += strconv.Itoa(e.value)
			if cnt != len(m)-1 {
				out += ","
			}
		}
	}
	return out
}

func isInOrder(a, b []*messageItem) bool {
	for {
		if len(a) == 0 {
			return true
		}
		itemA := a[0]
		if len(b) == 0 {
			return false
		}
		itemB := b[0]

		a = a[1:]
		b = b[1:]

		if !itemA.isList && !itemB.isList {
			if itemA.value == itemB.value {
				continue
			}

			return itemA.value < itemB.value
		}

		if itemA.isList && itemB.isList {
			if len(itemA.subList) == 0 && len(itemB.subList) == 0 {
				continue
			}
			if len(itemA.subList) == 0 {
				return true
			}
			if len(itemB.subList) == 0 {
				return false
			}

			firstA := itemA.subList[0]
			itemA.subList = itemA.subList[1:]
			newA := append([]*messageItem{firstA}, itemA)
			newA = append(newA, a...)

			firstB := itemB.subList[0]
			itemB.subList = itemB.subList[1:]
			newB := append([]*messageItem{firstB}, itemB)
			newB = append(newB, b...)
			return isInOrder(newA, newB)
		}

		if itemA.isList && !itemB.isList {
			// convert itemB to list
			bList := &messageItem{
				isList:  true,
				subList: []*messageItem{{isList: false, value: itemB.value}},
			}
			return isInOrder(append([]*messageItem{itemA}, a...), append([]*messageItem{bList}, b...))
		}

		if !itemA.isList && itemB.isList {
			// convert itemA to list
			aList := &messageItem{
				isList:  true,
				subList: []*messageItem{{isList: false, value: itemA.value}},
			}
			return isInOrder(append([]*messageItem{aList}, a...), append([]*messageItem{itemB}, b...))
		}
	}
}

func parseMessage(m []string) (int, []*messageItem) {
	work := make([]*messageItem, 0)
	for i := 0; i < len(m); i++ {
		if m[i] == "[" {
			processed, sub := parseMessage(m[i+1:])
			work = append(work, &messageItem{isList: true, subList: sub})
			i += processed
			continue
		}
		if m[i] == "," {
			continue
		}
		if m[i] == "]" {
			return i + 1, work
		}
		n, _ := strconv.ParseInt(m[i], 10, 64)

		// check if i+1 is a number as well (I know this is a little hacky)
		if i < len(m)-2 {
			if m[i+1] == "0" {
				n = 10
				i++
			}
		}

		work = append(work, &messageItem{isList: false, value: int(n)})
	}
	return len(m), work
}
