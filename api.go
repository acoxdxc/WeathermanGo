package main
import (
    "encoding/json"
    "fmt"
	"net/http"
	"io/ioutil"
	"os"	
)

//CityData todo
 type CityData struct{
	CapitalName string
 	CapitalLatitude string
 	CapitalLongitude string
 }

 //CityDataSlice todo
 type CityDataSlice struct{
	 Collection []CityData
 }

//City Latitude/Longitude 
//Structure is nested to match the structure of the API
type City struct{
	Name string
	Latitude float64 //`json: "latitutde"`
	Longitude float64 // `json:"longitude"`
	Currently Weather `json: "currently"`	 
}
//Weather nested into City
type Weather struct{
	Summary string
}
//Weather summary in these cities


var berlin City

var list []CityData

var url string = "https://api.darksky.net/forecast/58ffcf3ca186b5068bc9918ad2c16d8e/9.55,44.050000/"

var protocol string = "https"
var domain string = "api.darksky.net"
var route string = "forecast"
var cryptokey string = "58ffcf3ca186b5068bc9918ad2c16d8e"
var baseURL string = protocol + "://" + domain + "/"
var routeURL string = baseURL + route + "/"
var routeWithCryptoURL string = routeURL + cryptokey + "/"

// function with iterator inside
//      url_for_this_Request = routeWithCryptoURL + lat + ',' + long + '/'
//      request url_for_this_Request
//      


//locationURL todo
var locationURL string = "http://techslides.com/demos/country-capitals.json"
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


//UnmarshalBodyToPointer2 todo
func UnmarshalBodyToPointer2(Body []byte, city *CityData){
	err2 := json.Unmarshal(Body, &city)
	if err2 != nil {
		fmt.Println(err2)
			os.Exit(1)
	}
}

//GetLocationData will return Lat/Long/Name as a struct
func GetLocationData() ([]CityData){
	res := FetchResponse(locationURL)
	body := ResponseBodyToByte(res)
	keys := make([]CityData,0)

	json.Unmarshal(body, &keys)
//	fmt.Printf("%#v", keys)

	
	return keys
}

//GetWeatherData will return weather as struct
func GetWeatherData(url string)(*City){
	res := FetchResponse(url)
	body := ResponseBodyToByte(res)
	UnmarshalBodyToPointer(body, &berlin)
	//fmt.Println(berlin)
	return &berlin
	
}
//GetManyWeathers todo
func GetManyWeathers(routeWithCryptoURL string) ([]string){
	cities := GetLocationData()

//	json.Marshal(coord)
	var sunnyCities []string
	for _,city := range cities{
		name := city.CapitalName
		lat := city.CapitalLatitude
		long := city.CapitalLongitude
		var url string = routeWithCryptoURL + lat + "," + long + "/"
		var thisweather *City = GetWeatherData(url)
		
		var currently Weather = thisweather.Currently
		if currently.Summary == "Clear"{
			sunnyCities = append(sunnyCities, name)	
		}
	}
	return sunnyCities

	//var lat string = "15.35"
	//var long string = "44.200000"
	//var url string = routeWithCryptoURL + lat + "," + long + "/"
	//GetWeatherData(url)

}
//GetCoordinates todo
func GetCoordinates(){
	coord := GetLocationData()
	json.Marshal(coord)
	fmt.Println(coord)

	


	//Get key (a slice of structs) which has Lat/long
	//Create function to create an array of latitude and longitude values
	//substitue this values into url and unmarshal via iteration
	//Use range to iterate over Keys
	//Append to a new array


}

func main() {

	var manyWeathers []string = GetManyWeathers(routeWithCryptoURL)
	fmt.Println(manyWeathers)

//	ret := GetLocationData()
//	fmt.Printf("%#v", ret)
	//GetCoordinates()
	
	

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
