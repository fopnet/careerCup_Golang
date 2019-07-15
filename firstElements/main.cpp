#include <iostream>
using namespace std;

// Swap function to interchange
// the value of variables x and y
void swap(int &x, int &y)
{
    int temp = x;
    x = y;
    y = temp;
}

// Min Heap Class
// arr holds reference to an integer
// array size indicate the number of
// elements in Min Heap
class MinHeap
{

    int size;
    int *arr;
    int length;

public:
    // Constructor to initialize the size and arr
    MinHeap(int size, int length, int input[]);

    // Min Heapify function, that assumes that
    // 2*i+1 and 2*i+2 are min heap and fix the
    // heap property for i.
    void heapify(int i);

    // Build the min heap, by calling heapify
    // for all non-leaf nodes.
    void buildHeap();

    void printArray();
};

// Constructor to initialize data
// members and creating mean heap
MinHeap::MinHeap(int k, int length, int input[])
{
    // Initializing arr and size
    // int length = sizeof(input) ;
    // printf("input length %lu \n", length);

    this->size = k;
    this->arr = input;
    this->length = length;

    // Building the Min Heap
    buildHeap();
}

// Min Heapify function, that assumes
// 2*i+1 and 2*i+2 are min heap and
// fix min heap property for i

void MinHeap::printArray()
{
    // int length = sizeof(&arr) / sizeof(&arr[0]);
    // printf("length %lu %s\n", length, length == 2 ? "true" : "false");

    std::cout << "\n";
    for (int i = 0; i < length; i++)
    {
        std::cout << arr[i] << " ";
    }
}

/// heapify only not leafs
void MinHeap::heapify(int i)
{
    // If Leaf Node, Simply return
    if (i >= size / 2)
        return;

    // variable to store the smallest element
    // index out of i, 2*i+1 and 2*i+2
    int smallest;

    // Index of left node
    int left = 2 * i + 1;

    // Index of right node
    int right = 2 * i + 2;

    // Select minimum from left node and
    // current node i, and store the minimum
    // index in smallest variable
    smallest = arr[left] < arr[i] ? left : i;

    // If right child exist, compare and
    // update the smallest variable
    if (right < size)
        smallest = arr[right] < arr[smallest]
                       ? right
                       : smallest;

    // If Node i violates the min heap
    // property, swap  current node i with
    // smallest to fix the min-heap property
    // and recursively call heapify for node smallest.
    if (smallest != i)
    {
        swap(arr[i], arr[smallest]);
        heapify(smallest);
    }
}

// Build Min Heap
void MinHeap::buildHeap()
{
    // cout << "\nbefore heapyfy ";
    // printArray();

    // Calling Heapify for all non leaf nodes
    for (int i = size / 2 - 1; i >= 0; i--)
    {
        heapify(i);
    }
    cout << "after buildHeap ";
    printArray();
}

void FirstKelements(int arr[], int size, int k)
{
    // Creating Min Heap for given
    // array with only k elements
    MinHeap *m = new MinHeap(k, size, arr);

    // Loop For each element in array
    // after the kth element
    for (int i = k; i < size; i++)
    {

        // if current element is smaller
        // than minimum element, do nothing
        // and continue to next element
        if (arr[0] > arr[i])
            continue;

        // Otherwise Change minimum element to
        // current element, and call heapify to
        // restore the heap property
        else
        {
            arr[0] = arr[i];
            m->heapify(0);
            //  m->printArray();
        }

        // cout << "\nParcial => ";
        // m->printArray();
    }
    // Now min heap contains k maximum
    // elements, Iterate and print
    cout << "\nResult => ";
    for (int i = 0; i < k; i++)
    {
        cout << arr[i] << " ";
    }

    // cout << "\n Array final state";
    // m->printArray();
}
// Driver Program
int main()
{

    // int arr[] = {11, 3, 2, 1, 15, 5, 4, 45, 88, 96, 50, 45};
    int arr[] = {1, 10, 2, 15, 11, 7, 3, 22, 16, 14, 12, 18, 8, 6, 4, 31, 25, 28, 17, 29, 23, 24, 13, 30, 21, 26, 9, 27, 19, 20, 5};
    // const int LENGTH = 31;
    // int arr[LENGTH] = {};
    // for (int i = 0; i < LENGTH; i++)
    //     arr[i] = LENGTH - i;

    //     std::cout << "\n";
    // for (int i = 0; i < LENGTH; i++)
    // {
    //     std::cout << arr[i] << " ";
    // }

    int size = sizeof(arr) / sizeof(arr[0]); // 12

    // printf("size %lu\n", sizeof(arr));
    // printf("size [0] %lu\n", sizeof(arr[0]));
    // printf("size %lu\n", size);

    // Size of Min Heap
    int k = 10;

    FirstKelements(arr, size, k);

    return 0;
}
// This code is contributed by Ankur Goel