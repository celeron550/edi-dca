#include "linkedList.h"
#include <stdlib.h>
#include <stdio.h>

typedef struct LinkedList {
    size_t Inserted;
    Node *Head;
};

typedef struct Node {
    int val;
    int *next;
};

size_t LinkedList_Size(LinkedList *list) {
    return list->Inserted;
}

int LinkedList_Get(LinkedList *list, int index) {
    if (index < 0 || index >=list->Inserted) {
        fprintf(stderr,"Indice invalido\n");
        return -1;
    }
    Node *aux = list->Head;
    for (int i = 0; i < index; ++i){
        aux = aux->next;
    }
    return aux->val;
}

void LinkedList_Add(LinkedList *list, int val) {
    Node *Novo = {
        Novo->val = val,
        Novo->next = NULL,
    };

    Node *aux = list->Head;
    while (aux->next != NULL) {
        aux = aux->next;
    }
    aux->next = Novo;
    ++list->Inserted;
    return;
}

void LinkedList_RemoveOnIndex(LinkedList *list, int index) {
    if (index < 0 || index >=list->Inserted) {
        fprintf(stderr,"Indice invalido\n");
        return -1;
    }
    if(list->Inserted == 1) { // so tem um elemento na lista
        list->Head = NULL;
        --list->Inserted;
        return;
    }
    if (index == 0) {
        list->Head = list->Head->next;
        --list->Inserted;
        return;     
    }
    Node *aux = list->Head;

    for (int i = 0; i < index-1; ++i){
        aux = aux->next;
    }
    if (index == list->Inserted -1 ){
        aux->next = NULL;
        --list->Inserted;
        return;
    }
    Node *prox = aux->next;
    prox = aux->next;

    aux->next = prox;
    --list->Inserted;
    return;
}
int LinkedList_Pop(LinkedList *list)
{
    return 0;
}

void LinkedList_Set(LinkedList *list, int e, size_t index)
{
}
