// person_cgo.cpp

#include "person.h"

extern "C"
{
#include "person_cgo.h"
}

struct Person_T : Person
{
  Person_T(const char *name, int age) : Person(name, age) {}
  ~Person_T() {}
};

Person_T *NewPerson(const char *name, int age)
{
  return new Person_T(name, age);
}

void PersonSay(Person_T *person)
{
  person->Say();
}

void DeletePerson(Person_T *person)
{
  delete person;
}