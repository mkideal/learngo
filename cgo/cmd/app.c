#include "app.h"

typedef int(*ExecFunc) (void*);
static ExecFunc h_exec = NULL;

void load() {
    void *handle;
    char *error;

    handle = dlopen("../lib.so", RTLD_LAZY);
    if (!handle) {
        fprintf(stderr, "%s\n", dlerror());
        exit(EXIT_FAILURE);
    }
    dlerror();

    *(void **) (&h_exec) = dlsym(handle, "Exec");
    if ((error = dlerror()) != NULL) {
        fprintf(stderr, "%s\n", error);
        exit(EXIT_FAILURE);
    }
}

void exec(void* ptr) {
	h_exec(ptr);
};
