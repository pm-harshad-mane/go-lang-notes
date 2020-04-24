package main 

import (
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func wakandaHandlerOld(httpRespWriter http.ResponseWriter, httpRequest *http.Request){	
	me2 := new(Person)
	decoder2 := json.NewDecoder(httpRequest.Body) // using r.Body to parse
	err2 := decoder2.Decode(me2)
	if err2 != nil {
		fmt.Println("errors: ", err2)
	}
	fmt.Println("Me2 ", me2) // success

	buf2, _ := ioutil.ReadAll(httpRequest.Body)
	fmt.Println(buf2) // fails

	httpRespWriter.WriteHeader(http.StatusOK)
	httpRespWriter.Write([]byte( fmt.Sprintf("{success: \"%s\", statusMsg: \"%s\"}")))
}

func wakandaHandlerNew(httpRespWriter http.ResponseWriter, httpRequest *http.Request){
	buf, _ := ioutil.ReadAll(httpRequest.Body)
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	httpRequest.Body = rdr2 // assigning a copy here back to body

	rebuestBodyString := fmt.Sprintf("%s", rdr1)
	fmt.Println("rebuestBodyString copied value as string: ", rebuestBodyString) // printing the copid value // success

	me2 := new(Person)
	decoder2 := json.NewDecoder(httpRequest.Body) // using r.Body to parse
	err2 := decoder2.Decode(me2)
	if err2 != nil {
		fmt.Println("errors: ", err2)
	}
	fmt.Println("Me2 ", me2) // success

	httpRespWriter.WriteHeader(http.StatusOK)
	httpRespWriter.Write([]byte( fmt.Sprintf("{success: \"%s\", statusMsg: \"%s\"}")))
}

func main(){
	fmt.Println("Wakanda Forever!!! X ")	
	http.HandleFunc("/wakandaOld", wakandaHandlerOld)
	http.HandleFunc("/wakandaNew", wakandaHandlerNew)
	if err := http.ListenAndServe(":99", nil); err != nil {
		fmt.Println("Failed to start server", err)
	}
}

/*

	curl --location --request POST 'http://localhost:99/wakandaOld' \
	--header 'Content-Type: text/plain' \
	--data-raw '{ "name": "Harshad", "age": 34 }'

	curl --location --request POST 'http://localhost:99/wakandaNew' \
	--header 'Content-Type: text/plain' \
	--data-raw '{ "name": "Harshad", "age": 34 }'

*/
