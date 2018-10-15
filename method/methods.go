package method

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*type SksoilHandler struct {

}
type Sksoil_schemaHandler struct {

}*/

var Url="https://sksoils.usask.ca/api/v2/sksoil/"
var apiKey="0c6c4a2baf6ac608f2ef74354b1bbc2651a05675000983a397b07a5df96b1eac"

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func generateGetRequest(url string)*http.Request  {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	query := req.URL.Query()
	query.Add("api_key",apiKey)
	req.URL.RawQuery=query.Encode()
	fmt.Println(req.URL.String())
	return req
}
func getBody(url string)[]byte{
	client := &http.Client{}
	req:=generateGetRequest(url)
	res,err:=client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body,_:=ioutil.ReadAll(res.Body)
	res.Body.Close()
	return body
}
func SksoilHandler(w http.ResponseWriter,r *http.Request){
	enableCors(&w)
	if r.Method=="GET"{
		body:=getBody(Url)
		w.Write(body)
	}

}
func SkSchemaHandler(w http.ResponseWriter,r *http.Request){
	enableCors(&w)
	following := r.URL.Path[len("/sksoil/_schema/"):]
	thisUrl:=Url+"_schema/"
	if following== ""{
		if r.Method=="GET"{
			body:=getBody(thisUrl)
			w.Write(body)
		}
	}else {
		thisUrl:=thisUrl+following
		if r.Method=="GET"{
			body:=getBody(thisUrl)
			w.Write(body)
		}
	}

}
func SkTableHandler(w http.ResponseWriter,r *http.Request)  {
	enableCors(&w)
	following := r.URL.Path[len("/sksoil/_table/"):]
	thisUrl:=Url+"_table/"
	if following== ""{
		if r.Method=="GET"{
			body:=getBody(thisUrl)
			w.Write(body)
		}
	}else {
		thisUrl:=thisUrl+following
		if r.Method=="GET"{
			body:=getBody(thisUrl)
			w.Write(body)
		}
	}
}
func SkProcHandler(w http.ResponseWriter,r *http.Request)   {
	enableCors(&w)
	following := r.URL.Path[len("/sksoil/_proc/"):]
	thisUrl:=Url+"_proc/"
	if following== ""{
		if r.Method=="GET"{
			body:=getBody(thisUrl)
			w.Write(body)
		}
	}else {
		thisUrl:=thisUrl+following
		if r.Method=="GET"{
			body:=getBody(thisUrl)
			w.Write(body)
		}
	}
}
func SkFuncHandler(w http.ResponseWriter,r *http.Request)   {
	enableCors(&w)
	following := r.URL.Path[len("/sksoil/_func/"):]
	thisUrl:=Url+"_func/"
	if following== ""{
		if r.Method=="GET"{
			body:=getBody(thisUrl)
			w.Write(body)
		}
	}else {
		thisUrl:=thisUrl+following
		if r.Method=="GET"{
			body:=getBody(thisUrl)
			w.Write(body)
		}
	}
}