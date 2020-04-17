package servicerepo

import (
	"reflect"
)

type ServiceRepo struct {
	s map[string]interface{}
}

var singleInstance *ServiceRepo

func Inst() *ServiceRepo {
	if singleInstance != nil {
		return singleInstance
	}
	singleInstance = New()
	return singleInstance
}

func New() *ServiceRepo {
	newInstance := new(ServiceRepo)
	newInstance.s = make(map[string]interface{})
	return newInstance
}

func typeName(obj interface{}) string {
	typeName := reflect.TypeOf(obj).String()
	if len(typeName) >= 1 && typeName[0] == '*' {
		typeName = typeName[1:]
	}
	return typeName
}

func (s *ServiceRepo) Add(service interface{}) {
	typeName := typeName(service)
	s.AddName(typeName, service)
}

func (s *ServiceRepo) AddName(typeName string, service interface{}) {
	s.s[typeName] = service
}

func (s *ServiceRepo) Get(typeName string) interface{} {
	if _, ok := s.s[typeName]; !ok {
		// not found
		panic("ServiceRepo: service not found!  name: " + typeName)
	}
	return s.s[typeName]
}
