from max_heap import MaxHeap
from min_heap import MinHeap
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

print("Extract the max", heap.extract_max())

#print heap after extraction
print("After extraction, heap: ", heap)

print("--Max Heap--\n")

print("--Min Heap--")

minHeap = MinHeap()

minHeap.insert(5)
minHeap.insert(9)
minHeap.insert(2)
minHeap.insert(3)
minHeap.insert(7)
minHeap.insert(4)

#print heap
print(minHeap)

print("Extract the min", minHeap.extract_min())

#print heap after extraction
print("After extraction, heap: ", minHeap)

print("Extract the min", minHeap.extract_min())

#print heap after extraction
print("After extraction, heap: ", minHeap)

print("--Min Heap--")
