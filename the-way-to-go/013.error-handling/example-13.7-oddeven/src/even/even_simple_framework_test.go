package even
import (
    "testing"
    "reflect"
    "strings"
    "runtime"
    "strconv"
    "fmt"
)

type Case struct {
    input []interface{}
    output []interface{}
    fn interface{}
    t *testing.T
}

func (c *Case) Run() {
    fv := reflect.ValueOf(c.fn)
    fp := runtime.FuncForPC(fv.Pointer())
    if fv.Kind() == reflect.Func {
        // process input parameter
        params := inputToParams(c.input)

        // call function
        retSlice := fv.Call(params)

        // check return value count is valid?
        if len(retSlice) != len(c.output) {
            c.t.Logf("%s(%s) return value count: (%v). wanted: (%v).", fp.Name(), paramsToStr(params), len(retSlice), len(c.output))
            c.t.Fail()
        }

        // check return value is expected?
        satisfyExpected := true
        for i, v := range retSlice {
            if v.Interface() != c.output[i] {
                c.t.Logf("%s(%s) [return value index: %d] => (%v). wanted: (%v)", fp.Name(), paramsToStr(params), i, valueToStr(v), valueToStr(reflect.ValueOf(c.output[i])))
                satisfyExpected = false
            }
        }
        if !satisfyExpected {
            c.t.Fail()
        }
    }
}

func paramsToStr(params []reflect.Value) string {
    strArr := make([]string, len(params))
    for i, v := range params {
        strArr[i] = valueToStr(v)
    }
    return strings.Join(strArr, ",")
}
func inputToParams(input []interface{}) []reflect.Value {
    params := make([]reflect.Value, len(input))
    for i, v := range input {
        params[i] = reflect.ValueOf(v)
    }
    return params
}
func valueToStr(v reflect.Value) string {
    switch v.Kind() {
        case reflect.Bool:
            return strconv.FormatBool(v.Bool())
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            return strconv.FormatInt(v.Int(), 10)
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            return strconv.FormatUint(v.Uint(), 10)
        case reflect.Ptr, reflect.Uintptr, reflect.UnsafePointer:
            return fmt.Sprintf("%p", v.Pointer())
        case reflect.Float32:
            return strconv.FormatFloat(v.Float(), 'E', 2, 32)
        case reflect.Float64:
            return strconv.FormatFloat(v.Float(), 'E', 2, 64)
        case reflect.Complex64, reflect.Complex128:
            return fmt.Sprintf("%E", v.Complex())
        case reflect.Array, reflect.Slice:
            strArr := make([]string, v.Len())
            for i := 0; i < v.Len(); i ++ {
                strArr[i] = valueToStr(v.Index(i))
            }
            return "[" + strings.Join(strArr, ", ") + "]"
        case reflect.Map:
            strArr := make([]string, v.Len())
            keys := v.MapKeys()
            i := 0
            for _, k := range keys {
                strArr[i] = valueToStr(k) + ": " + valueToStr(v.MapIndex(k))
                i ++
            }
            return "{" + strings.Join(strArr, ", ") + "}"
        case reflect.Chan:
            return fmt.Sprintf("<chan %p>", v.Pointer())
        case reflect.Func:
            fp := runtime.FuncForPC(v.Pointer())
            return fmt.Sprintf("%s%s", fp.Name(), v.String())
        case reflect.String:
            return "\"" + v.String() + "\""
        case reflect.Struct:
            funcString := v.MethodByName("String")
            if funcString.IsValid() {
                retArr := funcString.Call([]reflect.Value{})
                if len(retArr) == 1 {
                    return retArr[0].String()
                }
            }

            strArr := make([]string, v.NumField())
            for i := 0; i < v.NumField(); i ++ {
                strArr[i] = valueToStr(v.Field(i))
            }
            return fmt.Sprintf("struct %v {%s}", v.Type(), strings.Join(strArr, ", "))
        case reflect.Interface:
            return fmt.Sprintf("<interface %p>", v)
    }
    return "unknown"
}
func anyToStr(v interface{}) string {
    return valueToStr(reflect.ValueOf(v))
}

func NewCase(t *testing.T, fn interface{}, nInput, nOutput int, params ...interface{}) *Case {
    input  := make([]interface{}, nInput)
    output := make([]interface{}, nOutput)
    n := 0
    for i := 0; i < nInput; i ++ {
        input[i] = params[n]
        n ++
    }
    for i := 0; i < nOutput; i ++ {
        output[i] = params[n]
        n ++
    }
    return &Case{input, output, fn, t}
}

func TestCases(t *testing.T) {
    NewCase(t, UdLog, 1, 1, 10, true).Run()

    NewCase(t, Even, 1, 1, 1, false).Run()
    NewCase(t, Even, 1, 1, 2, true).Run()
    NewCase(t, Even, 1, 1, 3, false).Run()
    NewCase(t, Even, 1, 1, 4, true).Run()

    NewCase(t, Odd, 1, 1, 1, true).Run()
    NewCase(t, Odd, 1, 1, 2, false).Run()
    NewCase(t, Odd, 1, 1, 3, true).Run()
    NewCase(t, Odd, 1, 1, 4, false).Run()

/* Test for Value to String
    iArr := new([5]int)
    iArr[0] = 1
    iArr[1] = 3
    iArr[2] = 5
    iArr[3] = 7
    iArr[4] = 9
    byteSlice := []byte{'h', 'e', 'l', 'l', 'o'}
    fmt.Printf("%s\n", anyToStr(true))
    fmt.Printf("%s\n", anyToStr(0))
    fmt.Printf("%s\n", anyToStr(0.32843))
    fmt.Printf("%s\n", anyToStr(3.2837 + 12.233i))
    fmt.Printf("%s\n", anyToStr(*iArr))
    fmt.Printf("%s\n", anyToStr([]int{2, 4, 6, 8, 10}))
    fmt.Printf("%s\n", anyToStr(map[string]int{"one": 1, "two": 2, "three": 3}))
    fmt.Printf("%s\n", anyToStr(UdLog))
    fmt.Printf("%s\n", anyToStr("Hello World"))
    fmt.Printf("%s\n", anyToStr(byteSlice))
    fmt.Printf("%s\n", anyToStr(Case{inputR, outputR, UdLog}))
*/
}
