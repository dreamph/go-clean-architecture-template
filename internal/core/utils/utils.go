package utils

import (
	"bytes"
	"html/template"
	"math"
	"math/big"
	"math/rand"
	"net/url"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"fmt"

	"github.com/jinzhu/copier"
	"github.com/samber/lo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func RemoveFirstZero(value string) string {
	if IsEmpty(value) {
		return value
	}
	if strings.HasPrefix(value, "0") {
		return value[1:]
	}
	return value
}

func SplitBy(value string, sep string) []string {
	if IsEmpty(value) {
		return nil
	}
	return strings.Split(value, sep)
}

func Length[T any](data *[]T) int64 {
	if data == nil {
		return 0
	}
	return int64(len(*data))
}

func FormatFullName(firstName string, lastName string) string {
	return cases.Title(language.Und).String(firstName + " " + lastName)
}

func RemoveDataFromSlice(list []string, data string) []string {
	var result []string
	if list == nil {
		return nil
	}
	for _, row := range list {
		if row != data {
			result = append(result, row)
		}
	}
	return result
}

func GetValueByMaxLength(val string, maxLength int) string {
	if len(val) > maxLength {
		return val[:maxLength]
	}
	return val
}

func InterfaceToString(data interface{}) string {
	if data == nil {
		return ""
	}
	return fmt.Sprintf("%v", data)
}

/*func DifferenceInFirstList(firstList []string, secondList []string) []string {
	var result []string
	for _, row := range secondList {
		if !IsContainsString(firstList, row) {
			result = append(result, row)
		}
	}
	return result
}*/

func GetArrayValueByIndex(list []string, index int) string {
	if len(list) < index {
		return ""
	}
	return list[index]
}

func PatchValue(toValue interface{}, fromValue interface{}) error {
	return copier.Copy(toValue, fromValue)
}

func ToChar(i int) string {
	return string(rune('A' + i))
}

func Retry(attempts int, sleep time.Duration, retryFunc func() error) error {
	for {
		err := retryFunc()
		if err == nil {
			return nil
		}
		attempts--
		if attempts > 0 {
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep += jitter / 2

			time.Sleep(sleep)
			continue
		}
		return err
	}
}

func RetryWithResult[T any](attempts int, sleep time.Duration, retryFunc func() (T, error)) (T, error) {
	for {
		result, err := retryFunc()
		if err == nil {
			return result, nil
		}
		attempts--
		if attempts > 0 {
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep += jitter / 2

			time.Sleep(sleep)
			continue
		}
		return result, err
	}
}

func MillisecondToTime(millisecond int64) time.Time {
	return time.Unix(0, millisecond*int64(1000000))
}

func ToFloat64(val string) float64 {
	result, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0
	}
	return result
}

func ToFloat64WithDefault(val string, defaultVal float64) float64 {
	result, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return defaultVal
	}
	return result
}

func ToInt64(val string) int64 {
	result, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return result
}

func NumberToString(val interface{}) string {
	return fmt.Sprintf("%v", val)
}

func Int64ToStringWithDefault(val int64, valid bool, defaultVal string) string {
	if !valid {
		return defaultVal
	}
	return strconv.FormatInt(val, 10)
}

func RemoveAllParamFromURL(fullURL string) string {
	//fullURL := "https://sodaudevusers.s3.dualstack.us-east-1.amazonaws.com/profile/022a04d9-db88-409f-9ddb-4c9896a0644c_profile.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIATQARX4ZZMBHQ2YUE%2F20210823%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210823T032931Z&X-Amz-Expires=604800&X-Amz-SignedHeaders=host&X-Amz-Signature=74312d48565801d698afca1608debabe1150bc65dfb76d788b2031bd7e56f20e"
	//baseUrl, _ := url.Parse(fullURL)
	i := strings.Index(fullURL, "?")
	if i == -1 {
		return fullURL
	}
	return fullURL[:i]
}

func IsAlphanumeric(str string) bool {
	var isStringAlphabetic = regexp.MustCompile(`^[a-zA-Z0-9_]*$`).MatchString
	return isStringAlphabetic(str)
}

/*
func FormatCurrency(value float64) string {
	result := fmt.Sprintf("%v", value)
	if math.Mod(value, 1) == 0 || math.Mod(value*10, 1) == 0 {
		result = fmt.Sprintf("%.2f", value)
	}
	return result
}
*/

func SplitByLengthToArrays(s string, n int) []string {
	var ss []string
	for i := 1; i < len(s); i++ {
		if i%n == 0 {
			ss = append(ss, s[:i])
			s = s[i:]
			i = 1
		}
	}
	ss = append(ss, s)
	return ss
}

func Contains[T comparable](collection []T, element T) bool {
	return lo.Contains(collection, element)
}

func Find[T any](collection []T, predicate func(T) bool) (T, bool) {
	return lo.Find(collection, predicate)
}

func FindIndexOf[T any](collection []T, predicate func(T) bool) (T, int, bool) {
	return lo.FindIndexOf(collection, predicate)
}

func Filter[V any](collection []V, predicate func(V, int) bool) []V {
	return lo.Filter(collection, predicate)
}

func Uniq[T comparable](collection []T) []T {
	return lo.Uniq(collection)
}

func UniqBy[T any, U comparable](collection []T, iteratee func(T) U) []T {
	return lo.UniqBy(collection, iteratee)
}

func Map[T any, R any](collection []T, iteratee func(T, int) R) []R {
	return lo.Map(collection, iteratee)
}

func Difference[T comparable](list1 []T, list2 []T) ([]T, []T) {
	return lo.Difference(list1, list2)
}

func DifferenceInFirstList[T comparable](firstList []T, secondList []T) []T {
	_, r2 := Difference(firstList, secondList)
	return r2
}

func Chunk[T any](collection []T, size int) [][]T {
	return lo.Chunk(collection, size)
}

func DropWhile[V any](collection []V, predicate func(V) bool) []V {
	return lo.DropWhile(collection, predicate)
}

func MergeMaps[K comparable, V any](maps ...map[K]V) map[K]V {
	if maps == nil {
		return nil
	}

	merged := make(map[K]V)
	for _, mapValue := range maps {
		if mapValue == nil {
			continue
		}
		for key, val := range mapValue {
			merged[key] = val
		}
	}

	return merged
}

func Reverse[T any](list []T) []T {
	return lo.Reverse(list)
}

func Masking(data string, maskingDigit int) string {
	return MaskingWithSign(data, "X", maskingDigit)
}

func MaskingWithSign(data string, sign string, maskingDigit int) string {
	if maskingDigit > len(data) {
		return data
	}
	var result string
	for i := 0; i < len(data)-maskingDigit; i++ {
		result += sign
	}
	runes := []rune(data)
	result += string(runes[len(data)-maskingDigit:])
	return result
}

func GetStackTrace(e interface{}) string {
	buf := make([]byte, 1<<20)
	buf = buf[:runtime.Stack(buf, false)]
	msg := fmt.Sprintf("panic: %v\n%s\n", e, buf)
	return msg
}

func GetStackTraceByError(err error) string {
	return fmt.Sprintf("%+v", err)
}

func RemoveSensitiveParams(jsn string, sensitiveParams []string, filteredLabel string) string {
	if jsn == "" {
		return ""
	}
	for _, param := range sensitiveParams {
		jsn = removeSensitiveParams(jsn, param, filteredLabel)
	}
	return jsn
}

func removeSensitiveParams(jsn string, param string, filteredLabel string) string {
	rWithSlash := *regexp.MustCompile(`\\"` + param + `\\":.*?"(.*?)\\"`)
	jsn = rWithSlash.ReplaceAllString(jsn, fmt.Sprintf(`\"`+param+`\": \"`+filteredLabel+`\"`))

	rWithoutSlash := *regexp.MustCompile(`"` + param + `":.*?"(.*?)"`)
	return rWithoutSlash.ReplaceAllString(jsn, fmt.Sprintf(`"`+param+`": "`+filteredLabel+`"`))
}

func PtrToValue[T any](val *T) T {
	if val == nil {
		var obj T
		return obj
	}
	return *val
}

func ValueToPtr[T any](val T) *T {
	return &val
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func GetSymbolPair(source, destination string) string {
	return destination + "_" + source
}

func BigIntAsFloat64(i *big.Int, decimals int) float64 {
	f := new(big.Float)
	f.SetPrec(100)
	f.SetInt(i)
	f.Quo(f, big.NewFloat(math.Pow10(decimals)))

	format := "%." + strconv.FormatInt(int64(decimals), 10) + "f"
	fl64, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)

	return fl64
}

func IsSlice(value interface{}) bool {
	typeOf := reflect.TypeOf(value)
	return typeOf.Kind() == reflect.Slice
}

func TrimString(s string) string {
	return strings.TrimSpace(s)
}

func IsURLExcludes(originalURL string, urlExcludes []string) bool {
	if len(urlExcludes) == 0 {
		return false
	}
	for _, urlExclude := range urlExcludes {
		if originalURL == urlExclude {
			return true
		}
	}
	return false
}

/*
func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}*/

func CapitalizedFirstStringOfSlice(list []string) []string {
	var result []string
	c := cases.Title(language.Und)
	for _, row := range list {
		result = append(result, c.String(row))
	}
	return result
}

func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func RemoveSpace(val string) string {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		return ""
	}
	return re.ReplaceAllString(val, "")
}

func RemoveBackslashNByBytes(val []byte) []byte {
	val = bytes.ReplaceAll(val, []byte("\n"), []byte(""))
	return bytes.ReplaceAll(val, []byte(`\n`), []byte(""))
}

func TemplateExecute(templateStr string, data interface{}) (string, error) {
	tem := template.New("template")
	_, err := tem.Parse(templateStr)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tem.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func GetPathFromUrl(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return u.Path
}

func JoinByOptionData(sep string, requiredData string, optionalData ...string) string {
	vals := []string{requiredData}
	if len(optionalData) != 0 {
		vals = append(vals, optionalData...)
	}
	return strings.Join(vals, sep)
}

func ServiceName(serviceName string, data ...string) string {
	return JoinByOptionData(":", serviceName, data...)
}

func GetStringValueIfEmpty(value string, defaultValue string) string {
	if IsEmpty(value) {
		return defaultValue
	}
	return value
}

func GetValueByCondition[T any](condition bool, whenTrueValue T, whenFalseValue T) T {
	if condition {
		return whenTrueValue
	}
	return whenFalseValue
}

func CompareBigInt(val1 *big.Int, val2 *big.Int) bool {
	if val1.Cmp(val2) == 0 {
		return true
	}
	return false
}

func IsValidUTF8(data []byte) bool {
	return utf8.Valid(data)
}

func IsValidUTF8String(data string) bool {
	return utf8.ValidString(data)
}

func ToValidUTF8(data string, checkBeforeConvert bool) string {
	if IsEmpty(data) {
		return data
	}
	if checkBeforeConvert {
		if IsValidUTF8String(data) {
			return data
		}
	}
	return strings.ToValidUTF8(data, "")
}

func SafeDefer(fn func() error) {
	_ = fn()
}

func ToBigInt(val string) *big.Int {
	result := new(big.Int)
	result.SetString(val, 10)
	return result
}
