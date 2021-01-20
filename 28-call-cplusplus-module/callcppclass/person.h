//person.h

#define _GLIBCXX_USE_CXX11_ABI 0

#include <iostream>
#include <string>

using namespace std;

struct Person
{
private:
  string name;
  int age;

public:
  Person(string name, int age);
  void Say();
};

Person::Person(string name, int age)
{
  this->name = name;
  this->age = age;
}

void Person::Say()
{
  cout << "name: " << this->name << ", age: " << this->age;
}
