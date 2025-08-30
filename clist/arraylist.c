#include <stdio.h>
#include <stdlib.h>
#include "arraylist.h"

void ArrayList_Init(ArrayList *list, size_t size) {
    list->v = calloc(size, sizeof(int));
    list->inserted = 0;
    list->capacity = size;
}

size_t ArrayList_Size(ArrayList *list) {
    return list->inserted;
}

int ArrayList_Get(ArrayList *list, size_t index) {
    if (index>=list->inserted) {
        fprintf(stderr,"Indice invalido\n");
        return -1;
    }
    return list->v[index];
}

void ArrayList_Add(ArrayList *list, int e) {
    if (list->inserted == list->capacity) {
        ArrayList_doubleV(list);
    }
    list->v[list->inserted] = e;
    list->inserted++;
}

void ArrayList_doubleV(ArrayList *list) {
    size_t newSize = (list->capacity == 0) ? 1 : list->capacity * 2;
    
    int *newV = realloc(list->v, newSize * sizeof(int));
    free(list->v);
    list->v = newV;
    list->capacity = newSize;
}

void ArrayList_RemoveOnIndex(ArrayList *list, size_t index) {
    if (index >= list->inserted) {
        fprintf(stderr,"Indice invalido\n"); 
        return;
    }
    for (size_t i = index; i < list->inserted-1; ++i) {
        list->v[i] = list->v[i+1];
    }
    --list->inserted;
    return;

}

void ArrayList_Set(ArrayList *list, int e, size_t index) {
    if (index >= list->inserted){
        fprintf(stderr,"Indice invalido\n"); 
        return;
    }
    list->v[index] = e;
    return;
}

int ArrayList_Pop(ArrayList *list) {
    if (list->inserted == 0){
        fprintf(stderr,"ArrayList vazio\n");
        return;
    }

    --list->inserted;
    return list->v[list->inserted];

}

void ArrayList_Destroy(ArrayList *list) {
    if (list->v != NULL) {
        free(list->v);
        list->v = NULL;
    }
    list->inserted = 0;
    list->capacity = 0;
}
