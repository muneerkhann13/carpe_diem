package log4

/*

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reached here")

	body, _ := ioutil.ReadAll(r.Body)
	var a bytes.Buffer
	json.Compact(&a, body)
	fmt.Println(a)
	fmt.Println(string(body))
	//WriteLog("request_log", "info", string(a))
	WriteLog("response_log", "info", "sfdjsdlfjsldfjslkdfsldfjsldkfjsdfsd")

}


func main() {

	LoadConfiguration("example.json")
	http.HandleFunc("/kamal", handler)
	http.ListenAndServe(":8080", nil)
}
*/
