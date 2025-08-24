typedef struct ArrayList ArrayList;

void ArrayList_Init(ArrayList *list, size_t size);
size_t ArrayList_Size(ArrayList *list);
int ArrayList_Get(ArrayList *list, size_t index);
void ArrayList_Add(ArrayList *list, int e);
void ArrayList_doubleV(ArrayList *list);
void ArrayList_RemoveOnIndex(ArrayList *list, size_t index);
void ArrayList_Pop(ArrayList *list);
void ArrayList_Set(ArrayList *list, int e, size_t index);
void ArrayList_Destroy(ArrayList *list);

typedef struct ArrayList
{
    int* v;
    size_t inserted;
    size_t capacity;

}ArrayList;