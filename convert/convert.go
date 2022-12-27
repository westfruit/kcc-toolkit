package convert

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/encoding/simplifiedchinese"
)

//float64 - float64
func Float64DivFloat64(a float64, b float64) float64 {
	da := decimal.NewFromFloat(a)
	db := decimal.NewFromFloat(b)
	res, _ := da.Div(db).Float64()
	return res
}

func IntDivInt(a int, b int) int {
	da := decimal.New(int64(a), 0)
	db := decimal.New(int64(b), 0)
	return int(da.Sub(db).IntPart())
}

func Float64AddFloat64(a float64, b float64) float64 {
	da := decimal.NewFromFloat(a)
	db := decimal.NewFromFloat(b)
	aa, _ := da.Add(db).Float64()
	return aa
}

func Float64SubFloat64(a float64, b float64) string {
	da := decimal.NewFromFloat(a)
	db := decimal.NewFromFloat(b)
	res, _ := da.Sub(db).Float64()

	return fmt.Sprintf("%.2f", res)
}

func ObjToJson(obj interface{}) string {

	if obj == nil {
		return ""
	}

	b, err := json.Marshal(obj)
	if err != nil {
		logrus.Info("ObjToJson, error, ", err)
		return ""
	}

	return string(b)
}

func JsonToObj(jsonString string, obj interface{}) error {

	if len(jsonString) == 0 {
		return errors.New("JSON字符串为空")
	}

	err := json.Unmarshal([]byte(jsonString), &obj)
	if err != nil {
		logrus.Info("JsonToObj, error, ", err)
		return err
	}

	return nil
}

func ConvertStr2GBK(str string) string {
	data, err := simplifiedchinese.GBK.NewEncoder().String(str)
	if err != nil {
		logrus.Error("ConvertStr2GBK err, ", err)
		return ""
	}

	return data
}

func ConvertGBK2Str(gbkStr string) string {
	data, err := simplifiedchinese.GBK.NewDecoder().String(gbkStr)
	if err != nil {
		logrus.Error("ConvertGBK2Str err, ", err)
		return ""
	}

	return data
}

//string转int
func StringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return n
}

//string转int32
func StringToInt32(s string) int32 {
	n, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	}

	return int32(n)
}

//string转int64
func StringToInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}

	return n
}

//convert number to big.Int type
func StringToBigInt(s string) *big.Int {
	// convert number to big.Int type
	ip := new(big.Int)
	ip.SetString(s, 10) //base 10

	return ip
}

//string转float32
func StringToFloat32(s string) float32 {
	if f, err := strconv.ParseFloat(s, 32); err == nil {
		return float32(f)
	}

	return 0
}

//string转float64
func StringToFloat64(s string) float64 {
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}

	return 0
}

/**
Int64ToString
*/
func Int64ToString(value int64) string {
	str := strconv.FormatInt(value, 10)
	return str
}

func Float32ToString(n float32) string {
	return strconv.FormatFloat(float64(n), 'f', -1, 32)
}

func Float64ToString(n float64) string {
	return strconv.FormatFloat(n, 'f', -1, 64)
}

/*
浮点数转为百分数字符串
*/
func Float64ToPercentStr(x float64) string {

	s := fmt.Sprintf("%.2f%%", x*100)

	return s
}

func InterfaceToString(x interface{}) string {

	s := fmt.Sprintf("%v", x)

	return s
}

func InterfaceToSqlString(x interface{}) string {

	switch i := x.(type) {
	case int64:
		return Int64ToString(i)
	case int32:
		return Int32ToString(i)
	case int:
		return IntToString(i)
	case string:
		return "'" + i + "'"
	}

	return ""
}

func InterfaceToFloat64(x interface{}) float64 {

	s := fmt.Sprintf("%v", x)

	f := StringToFloat64(s)

	return f
}

func IntToString(x int) string {

	s := fmt.Sprintf("%d", x)

	return s
}

//int32转string
func Int32ToString(n int32) string {

	n64 := int64(n)

	s := strconv.FormatInt(n64, 10)

	return s

}

var floatType = reflect.TypeOf(float64(0))
var stringType = reflect.TypeOf("")

func GetFloat(unk interface{}) (float64, error) {
	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	case string:
		return strconv.ParseFloat(i, 64)
	default:
		v := reflect.ValueOf(unk)
		v = reflect.Indirect(v)
		if v.Type().ConvertibleTo(floatType) {
			fv := v.Convert(floatType)
			return fv.Float(), nil
		} else if v.Type().ConvertibleTo(stringType) {
			sv := v.Convert(stringType)
			s := sv.String()
			return strconv.ParseFloat(s, 64)
		} else {
			return math.NaN(), fmt.Errorf("Can't convert %v to float64", v.Type())
		}
	}
}

// func Struct2Int64Slice(slice interface{}, field string) []int64 {
// 	var ids []int64
// 	s := reflect.ValueOf(slice)
// 	if s.Kind() != reflect.Slice {
// 		return ids
// 	}
// 	for i := 0; i < s.Len(); i++ {
// 		ids = append(ids, s.Index(i).FieldByName(field).Interface().(int64))
// 	}
// 	return ids
// }

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}

	return data
}

func HexToInt(hexString string) (int64, error) {
	if !strings.HasPrefix(hexString, "0x") {
		hexString = "0x" + hexString
	}

	val, ok := big.NewInt(0).SetString(hexString, 0)
	if !ok {
		return 0, fmt.Errorf("hex string to int error")
	}
	return val.Int64(), nil
}
