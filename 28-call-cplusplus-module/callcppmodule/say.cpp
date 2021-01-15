// say.cpp

#include <iostream>

extern "C" {
    #include "say.h"
}

void Say(const char* s) {
    std::cout << s;
}