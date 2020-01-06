package main

import (
	"fmt"
	"reflect"
)

var Patients = make(PatientList, 0)

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
