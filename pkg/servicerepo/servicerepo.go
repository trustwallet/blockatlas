package servicerepo

/**
 * ServiceRepository: a simple repository of string-keyed interfaces,
 * intended to be used as repo for internal services/components, to be used by each other.
 * Intended usage:
 * - Create one instance of the repo, early on.
 * - Pass this repo to service initialization, in a tree of calls.  This may seem cumbersome, but it ensures determined initialization order (avoid circular dependencies)
 * - Services should add themself to the repo, by interface.  Service type should be private, only interface methods should be public.
 * - Services can take out other services they need, they should store them in local variables.
 * - If interfaces are added, there is no dependency to implementations, mocking is easy to achieve.
 */

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
