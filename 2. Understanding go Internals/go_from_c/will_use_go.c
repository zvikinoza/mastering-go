#include <stdio.h>
#include "go_from_c.h"

int main() {
    GoInt x = 12;
    GoInt y = 2;

    printf("About to call go function\n");
    PrintMessage();

    GoInt p = Multiply(x, y);
    printf("Product is %d", (int)p);
    printf("Wow It worked!");
    return 0;
}