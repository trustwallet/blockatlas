package servicerepo

import (
	"testing"
)

// A custom service type
type serviceOneI interface {
	Name() string
}

type serviceOneType struct {
	name string
}

func (s *serviceOneType) Name() string {
	return s.name
}

// constructor
func newOne(name string) serviceOneI {
	n := new(serviceOneType)
	n.name = name
	return n
}

// specialized getter on ServiceRepo (with specialized return type, type name constant and casting)
func (s *ServiceRepo) getServiceOne() serviceOneI {
	return s.Get("servicerepo.serviceOneType").(serviceOneI)
}

func TestAddName(t *testing.T) {
	sr := New()
	service1 := newOne("S1")
	const serviceName = "ServiceOne"
	sr.AddName(serviceName, service1)
	if len(sr.s) != 1 {
		t.Errorf("Wrong size")
	}

	retrieved := sr.Get(serviceName)
	if retrieved == nil {
		t.Errorf("Could not retrieve")
	}
}

func TestAdd(t *testing.T) {
	sr := New()
	service1 := newOne("S1")
	sr.Add(service1)
	if len(sr.s) != 1 {
		t.Errorf("Wrong size")
	}

	retrieved := sr.Get("servicerepo.serviceOneType")
	if retrieved == nil {
		t.Errorf("Could not retrieve")
	}

	// check casting and content
	retrievedService := retrieved.(serviceOneI)
	if retrievedService.Name() != "S1" {
		t.Errorf("Wrong content")
	}
}

func TestSingleton(t *testing.T) {
	sr := Inst()
	if sr == nil {
		t.Errorf("Nil service repo")
	}
	// contents is undefined
}

func TestCustomGet(t *testing.T) {
	sr := New()
	service1 := newOne("S1")
	sr.Add(service1)
	if len(sr.s) != 1 {
		t.Errorf("Wrong size")
	}

	retrieved := sr.getServiceOne()
	if retrieved.Name() != "S1" {
		t.Errorf("Wrong content")
	}
}

func TestGetNonexistentPanics(t *testing.T) {
	sr := New()
	service1 := newOne("S1")
	sr.Add(service1)

	defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
	}()
	
	_ = sr.Get("NO_SUCH_SERVICE")
}
