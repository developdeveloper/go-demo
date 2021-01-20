// person_cgo.h

typedef struct Person_T Person_T;

Person_T* NewPerson(const char* name, int age);
void PersonSay(Person_T *person);
void DeletePerson(Person_T* person);