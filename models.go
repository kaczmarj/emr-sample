package main

import (
	"fmt"
	"log"
	"reflect"
)

type Patient struct {
	Id   int    `json:"id"`
	Name string `json:"name,omitempty"`
	DOB  string `json:"dob,omitempty"`
}

func (p *Patient) ModifyName(name string) {
	if name == "" {
		log.Printf("not modifying name of patient %d because name is empty", p.Id)
	} else {
		p.Name = name
	}
}

func (p *Patient) ModifyDOB(dob string) {
	if dob == "" {
		log.Printf("not modifying dob of patient %d because dob is empty", p.Id)
	} else {
		p.DOB = dob
	}
}

type PatientList []*Patient

func AddPatientToList(p *Patient) {
	Patients = append(Patients, p)
}

func (p *PatientList) Get(id int) (*Patient, error) {
	if id < 0 || id >= len(Patients) {
		return nil, fmt.Errorf("patient %d not found", id)
	}
	return Patients[id], nil
}

func (p *PatientList) Equals(q *PatientList) bool {
	if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	}
	if len(*p) != len(*q) {
		return false
	}
	for i := range *p {
		if !reflect.DeepEqual((*p)[i], (*q)[i]) {
			return false
		}
	}
	return true
}

func NewPatient(name, dob string) *Patient {
	id := len(Patients)
	if name == "" {
		log.Printf("name cannot be empty\n")
	} else if dob == "" {
		log.Printf("DOB cannot be empty\n")
	}
	return &Patient{Id: id, Name: name, DOB: dob}
}

func EditPatient(id int, name, dob string) error {
	pt, err := Patients.Get(id)
	if err != nil {
		return fmt.Errorf("patient %d not found", id)
	}
	pt.ModifyName(name)
	pt.ModifyDOB(dob)
	Patients[id] = pt
	return nil
}
