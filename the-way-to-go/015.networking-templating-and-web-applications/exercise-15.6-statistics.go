package main
/*
    * get input numbers
    * print them
    * print their mean
    * print their median
 */

import (
    "net/http"
    "io"
    "fmt"
    "regexp"
    "sort"
    "strings"
    "strconv"
)

const (
    html_form = "<html><head><title>Statistics</title></head><body><form method=\"POST\" action=\"\"><input type=\"text\" name=\"inp\" value=\"%s\"/><br /><input type=\"submit\"/></form>%s</body></html>"
    html_table_fmt = "<table border=\"1\" cellspacing=\"0\" cellpading=\"0\"><tr><th colspan=\"2\">Results</th></tr><tr><th>Numbers</th><td>%s</td></tr><tr><th>Count</th><td>%d</td></tr><tr><th>Mean</th><td>%0.6f</td></tr><tr><th>Median</th><td>%0.6f</td></tr></table>"
)

var reSplit *regexp.Regexp

func init() {
    re, err := regexp.Compile("\\D+")
    if err != nil {
        panic(err.Error())
    }
    reSplit = re
}

func main() {
    http.HandleFunc("/statistics", StatisticsServ)
    err := http.ListenAndServe("0.0.0.0:8088", nil)
    if err != nil {
        fmt.Printf("web-server startup failed: %s\n", err.Error())
    }
}

func StatisticsServ(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    inputString  := req.FormValue("inp")
    contentTable := ""
    if req.Method == http.MethodPost {
        inputEles := reSplit.Split(inputString, -1)

        var inputNums sort.IntSlice = make([]int, 0, len(inputEles))
        sum := 0
        for i, v := range inputEles {
            n, err := strconv.ParseInt(v, 10, 64)
            if err == nil {
                inputNums = append(inputNums, int(n))
                sum += int(n)
            } else {
                inputEles[i] = "NaN"
            }
        }
        numStr := "[" + strings.Join(inputEles, ", ") + "]"

        var median float64
        count := len(inputNums)
        sort.Sort(inputNums)
        if count % 2 == 0 {
            median = float64(inputNums[count / 2] + inputNums[count / 2 + 1]) / 2.0
        } else {
            median = float64(inputNums[count / 2])
        }

        contentTable = fmt.Sprintf(html_table_fmt, numStr, count, float64(sum) / float64(count), median)
    }

    io.WriteString(w, fmt.Sprintf(html_form, inputString, contentTable))
}
