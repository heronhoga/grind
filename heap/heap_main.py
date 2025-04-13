from max_heap import MaxHeap
print("--Max Heap--")
heap = MaxHeap()

heap.insert(5)
heap.insert(9)
heap.insert(2)
heap.insert(3)
heap.insert(7)
heap.insert(4)

#print heap
print(heap)

print("Extract the max", heap.extract_max())

#print heap after extraction
print("After extraction, heap: ", heap)
print("--Max Heap--\n")

print("--Min Heap--")
print("--Min Heap--")
