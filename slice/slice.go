package slice

import (
	"sort"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StringIsMember(s string, ss []string) bool {
	for _, member := range ss {
		if member == s {
			return true
		}
	}

	return false
}

func IntIsMember(i int, ii []int) bool {
	for _, member := range ii {
		if member == i {
			return true
		}
	}

	return false
}

func ObjectIDIsMember(s primitive.ObjectID, ss []primitive.ObjectID) bool {
	for _, member := range ss {
		if member == s {
			return true
		}
	}

	return false
}

func StringIsMemberCaseInsensitive(s string, ss []string) (bool, string) {
	for _, member := range ss {
		if strings.ToLower(member) == strings.ToLower(s) {
			return true, member
		}
	}

	return false, ""
}

func StringIsPrefixMemberCaseInsensitive(s string, ss []string) bool {
	lowerS := strings.ToLower(s)

	for _, member := range ss {
		lowerMember := strings.ToLower(member)

		// Prefix matchins is specified by a trailing *, E.g. KW*
		if strings.HasSuffix(lowerMember, "*") {
			prefixRemoved := strings.TrimSuffix(lowerMember, "*")
			if strings.HasPrefix(lowerS, prefixRemoved) && lowerS != prefixRemoved {
				return true
			}
		} else {
			if lowerMember == lowerS {
				return true
			}
		}
	}

	return false
}

func FilterWithBlackList(myStrings, allStrings []string, isBlacklist bool) (result []string) {
	if isBlacklist {
		for _, s := range allStrings {
			if !StringIsMember(s, myStrings) {
				result = append(result, s)
			}
		}
	} else {
		result = myStrings
	}

	return
}

func IsValidWithBlackList(s string, myStrings []string, isBlacklist bool) bool {
	targetIsOnList := StringIsMember(s, myStrings)
	if isBlacklist {
		return !targetIsOnList
	}

	return targetIsOnList

}

func IsValidCaseInsensitiveByPrefixWithBlackList(s string, myStrings []string, isBlacklist bool) bool {
	targetIsOnList := StringIsPrefixMemberCaseInsensitive(s, myStrings)
	if isBlacklist {
		return !targetIsOnList
	}

	return targetIsOnList
}

func StringSliceToCommaSeparatedList(ss []string) string {
	return strings.Join(ss, ",")
}

func StringSliceRemoveBlanks(ss []string) (res []string) {
	for i := range ss {
		if ss[i] != "" {
			res = append(res, ss[i])
		}
	}

	return
}

func StringSliceJoinIf(ss []string, sep string) string {
	return strings.Join(StringSliceRemoveBlanks(ss), sep)
}

func StringSliceHasNonBlanks(ss []string) bool {
	return len(StringSliceRemoveBlanks(ss)) > 0
}

func StringSliceRemoveDuplicates(ss []string) (res []string) {
	m := map[string]bool{}

	for i := range ss {
		m[ss[i]] = true
	}

	for s := range m {
		res = append(res, s)
	}

	return
}

func StringSliceAddToSet(s string, ss []string) []string {
	if StringIsMember(s, ss) {
		return ss
	}

	return append(ss, s)
}

func StringSliceSort(ss []string, desc bool) {
	if desc {
		sort.Slice(ss, func(i, j int) bool {
			return ss[i] > ss[j]
		})
	} else {
		sort.Slice(ss, func(i, j int) bool {
			return ss[i] < ss[j]
		})
	}
}

func LastNonEmptyString(ss ...string) (res string) {
	if len(ss) == 0 {
		return
	}

	for i := range ss {
		if ss[i] != "" {
			res = ss[i]
		}
	}

	return
}

func CompareStringSlicesOrderIrrelevant(ss1, ss2 []string) bool {
	// Low hanging fruit first
	if len(ss1) != len(ss2) {
		return false
	}

	// Next low hanging fruit:
	// single member slice - just compare the single member
	if len(ss1) == 1 {
		return ss1[0] == ss2[0]
	}

	// At this point, the only thing left are
	// two slices of the same length and longer than 1
	// Loop through one of the slices, checking that each member
	// is also a member of the other
	for i := range ss1 {
		if !StringIsMember(ss1[i], ss2) {
			return false
		}
	}

	return true
}

func NeededToMakeSliceLongEnoughForIndex(listLength, index int) int {
	if index < listLength {
		return 0
	}

	return (index - listLength) + 1
}
