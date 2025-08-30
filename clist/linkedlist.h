#pragma once

typedef struct LinkedList LinkedList;

typedef struct Node Node;

size_t LinkedList_Size(LinkedList *list);

int LinkedList_Get(LinkedList *list, int index);

void LinkedList_Add(LinkedList *list, int val);

void LinkedList_RemoveOnIndex(LinkedList *list, int index);

int LinkedList_Pop(LinkedList *list);

void LinkedList_Set(LinkedList *list, int e, size_t index);