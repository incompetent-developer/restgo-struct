package restgostruct

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Cloner is function for deep cloning map[string]interface{}
func Cloner(data map[string]interface{}) map[string]interface{} {
	clone := make(map[string]interface{})
	for key, value := range data {
		switch value.(type) {
		case map[string]interface{}:
			clone[key] = Cloner(value.(map[string]interface{}))

		default:
			clone[key] = value

		}
	}
	return clone
}

// RandStringRunes is random string in runes
func RandStringRunes(runes []rune, n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

// RandIntInRange (min, max int)
//
// Function is random int between min -> (max-1). Ex: 0 -> 3 possible number is (0, 1, 2)
//
func RandIntInRange(min, max int) int {
	return rand.Intn(max-min) + min
}

// EqualSliceString is compare 2 slice of string is equal ?
func EqualSliceString(s1, s2 []string) bool {
	sort.Strings(s1)
	sort.Strings(s2)

	return reflect.DeepEqual(s1, s2)
}

// ItemExistsInSlice is check item exists in slice,arry
func ItemExistsInSlice(slice interface{}, item interface{}) bool {
	arr := reflect.ValueOf(slice)

	if arr.Kind() != reflect.Array && arr.Kind() != reflect.Slice {
		return false
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

// ItemStringContainExistsInStrArray is check item exists in array
func ItemStringContainExistsInStrArray(slice []string, item string, isPrefix bool) (string, bool) {

	for _, data := range slice {
		if isPrefix {
			if strings.HasPrefix(item, data) {
				return data, true
			}
		} else {
			if strings.Contains(item, data) {
				return data, true
			}
		}
	}

	return "", false
}

// InterString function
// Describtion: this function use for convert value in interface into string.
func InterString(value interface{}) (string, error) {
	switch value.(type) {
	case string:
		return value.(string), nil
	case int:
		return fmt.Sprintf("%d", value.(int)), nil
	case uint:
		return fmt.Sprintf("%d", value.(uint)), nil
	case int32:
		return fmt.Sprintf("%d", value.(int32)), nil
	case uint32:
		return fmt.Sprintf("%d", value.(uint32)), nil
	case int64:
		return fmt.Sprintf("%d", value.(int64)), nil
	case uint64:
		return fmt.Sprintf("%d", value.(uint64)), nil
	case float32:
		return fmt.Sprintf("%0.f", value.(float32)), nil
	case float64:
		return fmt.Sprintf("%0.f", value.(float64)), nil
	case json.Number:
		return string(value.(json.Number)), nil
	default:
		return "", fmt.Errorf("Cannot parse value (%v) type (%T) to string", value, value)
	}
}

// InterFloat function
// Describtion: this function use for convert value in interface into float64.
func InterFloat(value interface{}) (float64, error) {
	switch value.(type) {
	case string:
		return strconv.ParseFloat(value.(string), 64)
	case int:
		return float64(value.(int)), nil
	case uint:
		return float64(value.(uint)), nil
	case int32:
		return float64(value.(int32)), nil
	case uint32:
		return float64(value.(uint32)), nil
	case int64:
		return float64(value.(int64)), nil
	case uint64:
		return float64(value.(uint64)), nil
	case float32:
		return float64(value.(float32)), nil
	case float64:
		return value.(float64), nil
	case json.Number:
		return strconv.ParseFloat(string(value.(json.Number)), 64)
	default:
		return 0, fmt.Errorf("Cannot parse value (%v) type (%T) to float64", value, value)
	}
}

// InterInt function
// Describtion: this function use for convert value in interface into integer.
func InterInt(value interface{}) (int, error) {
	switch value.(type) {
	case string:
		return strconv.Atoi(value.(string))
	case int:
		return value.(int), nil
	case uint:
		return int(value.(uint)), nil
	case int32:
		return int(value.(int32)), nil
	case uint32:
		return int(value.(uint32)), nil
	case int64:
		return int(value.(int64)), nil
	case uint64:
		return int(value.(uint64)), nil
	case float32:
		return int(value.(float32)), nil
	case float64:
		return int(value.(float64)), nil
	case json.Number:
		return strconv.Atoi(string(value.(json.Number)))
	default:
		return 0, fmt.Errorf("Cannot parse value (%v) type (%T) to float64", value, value)
	}
}

// Round function
// Description: this function use for round value float64 into integer.
func Round(num float64) int64 {
	return int64(num + math.Copysign(0.5, num))
}

// ToFixed function
// Description: this function use for round value with precition float64 into integer.
func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}

// Display function
func Display(data interface{}) {
	disp, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(disp))
}

// ProjectName is return a project name from parent folder
func ProjectName() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	spath := strings.Split(wd, "/")

	return spath[(len(spath) - 1)], nil
}

// IsFloatInt is function for checking type float/int and return boolean
// eg. floatValue = 1 return true
// eg. floatValue = 1.2 return false
func IsFloatInt(floatValue float64) bool {
	return math.Mod(floatValue, 1.0) == 0
}

// ConvertBoolToInt64 is convert boolean to int64
func ConvertBoolToInt64(b bool) int64 {
	converter := map[bool]int64{
		true:  1,
		false: 0,
	}

	return converter[b]
}

// GetBoolean is get all boolean
func GetBoolean() map[string]string {
	return map[string]string{
		"true":  "Yes",
		"false": "No",
	}
}

// StringContainsEmpty is check string is empty ?
func StringContainsEmpty(ss ...string) bool {
	if len(ss) == 0 {
		return true
	}
	for _, s := range ss {
		if s == "" {
			return true
		}
	}
	return false
}

// NumberFormatWithComma is format number with comma
func NumberFormatWithComma(number interface{}) string {
	p := message.NewPrinter(language.English)

	switch number.(type) {
	case int:
		return p.Sprintf("%d", number)
	case float64:
		return p.Sprintf("%.2f", number)
	case string:
		return number.(string)
	default:
		return ""
	}
}

//ConvertTimestampToNano is time string converter
func ConvertTimestampToNano(timestamp string) string {
	if len(timestamp) != 19 {
		suffix := strings.Repeat("0", (19 - len(timestamp)))
		timestamp = timestamp + suffix
	}
	return timestamp
}

//ConvertNanoTimestampToTime is covert nano timstamp to time
func ConvertNanoTimestampToTime(nanoTimestamp int64) time.Time {
	return time.Unix(nanoTimestamp/1000000000, 0)
}

// GenerateIDProc is return check digit
func GenerateIDProc(prefix string, size int) (string, error) {
	ramdomLen := size - len(prefix)

	temp := fmt.Sprintf("%s%s", prefix, RandStringRunes(GetNumberRunes(), ramdomLen))

	var (
		sum int
	)

	for index := len(temp) - 2; index >= 0; index-- {

		parsed, err := strconv.Atoi(string(temp[index]))
		if err != nil {
			return "", err
		}

		if index%2 == 0 {
			multiply := parsed * 2

			if multiply >= 10 {
				multiply = multiply - 9
			}

			sum += multiply
		} else {
			sum += parsed
		}

	}

	checkDigit := sum % 10
	if checkDigit != 0 {
		checkDigit = 10 - checkDigit
	}

	return fmt.Sprintf("%s%d", temp[:len(temp)-1], checkDigit), nil
}

// Valid a card number
func Valid(number string) bool {
	var sum int
	var alternate bool

	numberLen := len(number)

	if numberLen < 13 || numberLen > 19 {
		return false
	}

	for i := numberLen - 1; i > -1; i-- {
		mod, _ := strconv.Atoi(string(number[i]))
		if alternate {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}

		alternate = !alternate

		sum += mod
	}

	return sum%10 == 0
}

// GenerateID is return check digit
func GenerateID(prefix string) (string, error) {
	ts := fmt.Sprintf("%d", time.Now().UnixNano())

	tid, err := GenerateIDProc(ts, 30)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s", prefix, tid), nil
}
