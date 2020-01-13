package main

import (
    "encoding/json"
    "fmt"
	"net/http"
	"io/ioutil"
	"os"
	
)

// type []CityData struct{

// 	CountryName string
// 	CapitalLatitude float64
// 	CapitalLongitude float64
// }


//City Latitude/Longitude 
//Structure is nested to match the structure of the API
type City struct{
	Name string
	Latitude float64 //`json: "latitutde"`
	Longitude float64 // `json:"longitude"`
	Currently Weather `json: "currently"`	 
	}

type Weather struct{
	Summary string
}
//Weather summary in these cities


var berlin City

var url string = "https://api.darksky.net/forecast/58ffcf3ca186b5068bc9918ad2c16d8e/9.55,44.050000/"

// func GetUrlList(){

//}

//FetchResponse gets response 
func FetchResponse(url string)(res *http.Response){
	res, err := http.Get(url)
	if err != nil{
		fmt.Println(err)
			os.Exit(1)
	}
	return res
}

//ResponseBodyToByte will convert data to Byte format
func ResponseBodyToByte(res *http.Response) ([]byte) {
	item, err := ioutil.ReadAll(res.Body)
	if err != nil {
	fmt.Println(err)
		os.Exit(1)
	}
	return item
}

//UnmarshalBodyToPointer using json, storing in city instance
func UnmarshalBodyToPointer(Body []byte, city *City){
	err2 := json.Unmarshal(Body, &city)
	if err2 != nil {
		fmt.Println(err2)
			os.Exit(1)
	}
}

//GetWeatherData will return weather as struct
func GetWeatherData(){
	res := FetchResponse(url)
	body := ResponseBodyToByte(res)
	UnmarshalBodyToPointer(body, &berlin)
	
}


func main() {

	GetWeatherData()
	fmt.Println(berlin)

	// //GetWeatherData()
	// fmt.Println(berlin.Currently, berlin.Latitude, berlin.Longitude)
	// test := "http://techslides.com/demos/country-capitals.json"
	// r, err3 := http.Get(test)
	// if err3 != nil {
	// 	fmt.Println(err3)
	// 		os.Exit(1)
 
	//  }

	//  Item, err := ioutil.ReadAll(r.Body)
	//  if err != nil {
	// 	fmt.Println(err)
	// 		os.Exit(1)
 
	//  }

	 
	//  listOfCities := []CityData{}

	//  err5 := json.Unmarshal(Item, &listOfCities)
	//  if err5 != nil {
	// 	 fmt.Println(err5)
	// 		 os.Exit(1)
 
	//   }

	// fmt.Println(Item)

	}
