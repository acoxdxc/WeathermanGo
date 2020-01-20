package main
import ("fmt")
import "testing"

func TestGetWeather(t *testing.T) {
	var berlin City
	got := (berlin)
	want := "{ 9.55 44.05 {Partly Cloudy}}"

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}