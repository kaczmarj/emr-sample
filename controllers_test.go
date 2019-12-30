package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGETPatients(t *testing.T) {
	t.Run("get empty patient collection", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/patients", nil)
		response := httptest.NewRecorder()
		Patients = make(PatientList, 0)
		GetPatients(response, request, nil)
		y := "[]\n"
		y_ := response.Body.String()
		if y != y_ {
			t.Errorf("got %q but expected %q", y_, y)
		}
		if response.Code != http.StatusOK {
			t.Errorf("got code %d but expected code %d", response.Code, http.StatusOK)
		}
	})

	t.Run("get nonempty patient collection", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/patients", nil)
		response := httptest.NewRecorder()
		Patients = make(PatientList, 0)
		AddPatientToList(&Patient{Name: "john", DOB: "20190512"})
		GetPatients(response, request, nil)
		if response.Code != http.StatusOK {
			t.Errorf("got code %d but expected code %d", response.Code, http.StatusOK)
		}
		var pts PatientList
		err := json.Unmarshal(response.Body.Bytes(), &pts)
		if err != nil {
			t.Errorf("error unmarshalling json")
		}
		if !Patients.Equals(&pts) {
			t.Errorf("patient slices not equal")
		}
	})
}

func TestPatient(t *testing.T) {
	Patients = make(PatientList, 0)

	t.Run("get unknown patient", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/patients/10", nil)
		response := httptest.NewRecorder()
		GetPatient(response, request, httprouter.Params{httprouter.Param{Key: "id", Value: "10"}})
		y := "patient not found\n"
		y_ := response.Body.String()
		if y != y_ {
			t.Errorf("got %q but expected %q", y_, y)
		}
		if response.Code != http.StatusNotFound {
			t.Errorf("got code %d but expected code %d", response.Code, http.StatusNotFound)
		}
	})

	p := Patient{Name: "john", DOB: "19990812"}
	AddPatientToList(&p)
	AddPatientToList(&Patient{Name: "joe", DOB: "20000101"})
	AddPatientToList(&Patient{Name: "kim", DOB: "20000101"})
	AddPatientToList(&Patient{Name: "ellen", DOB: "20000101"})

	t.Run("get known patient", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/patients/10", nil)
		response := httptest.NewRecorder()
		GetPatient(response, request, httprouter.Params{httprouter.Param{Key: "id", Value: "0"}})
		if response.Code != http.StatusOK {
			t.Errorf("got code %d but expected code %d", response.Code, http.StatusOK)
		}
		var p_ Patient
		err := json.Unmarshal(response.Body.Bytes(), &p_)
		if err != nil {
			t.Errorf("error unmarshalling json")
		}
		if p != p_ {
			t.Errorf("patients not equal")
		}
	})

	t.Run("edit known patient", func(t *testing.T) {
		form := url.Values{
			"name": {"mark"},
			"dob":  {"2000"}}
		request, _ := http.NewRequest(http.MethodPatch, "/patients/0", strings.NewReader(form.Encode()))
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		response := httptest.NewRecorder()
		PatchPatient(response, request, httprouter.Params{httprouter.Param{Key: "id", Value: "0"}})
		if response.Code != http.StatusOK {
			t.Errorf("got code %d but expected code %d", response.Code, http.StatusOK)
		}
		if Patients[0].Name != "mark" || Patients[0].DOB != "2000" {
			t.Errorf("expected name to be mark and dob to be 2000")
		}
	})

	t.Run("edit unknown patient", func(t *testing.T) {
		form := url.Values{
			"name": {"mark"},
			"dob":  {"2000"}}
		request, _ := http.NewRequest(http.MethodPatch, "/patients/10", strings.NewReader(form.Encode()))
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		response := httptest.NewRecorder()
		PatchPatient(response, request, httprouter.Params{httprouter.Param{Key: "id", Value: "10"}})
		if response.Code != http.StatusNotFound {
			t.Errorf("got code %d but expected code %d", response.Code, http.StatusNotFound)
		}
	})
}
