class MaxHeap:
    def __init__(self):
        self.heap = []
        
    def _parent(self, i):
        return (i-1) // 2
    
    def _left(self, i):
        return (2 * i) + 1
    
    def _right(self, i):
        return (2 * i) + 2
    
    def insert(self, key):
        self.heap.append(key)
        i = len(self.heap) - 1
        #bubble up the new key node
        while i != 0 and self.heap[self._parent(i)] < self.heap[i]:
            self.heap[i], self.heap[self._parent(i)] = self.heap[self._parent(i)], self.heap[i]
            i = self._parent(i)
            
    def extract_max(self):
        if not self.heap:
            return None
        
        if len(self.heap) == 1:
            return self.heap.pop()
        
        root = self.heap[0]
        
        self.heap[0] = self.heap.pop()
        
        self._heapify(0)
        return root
            
    def _heapify(self, i):
        largest = i
        left = self._left(i)
        right = self._right(i)
        
        if left < len(self.heap) and self.heap[i] > self.heap[largest]:
            largest = left
        if right < len(self.heap) and self.heap[i] > self.heap[largest]:
            largest = right
        
        if largest != i:
            self.heap[i], self.heap[largest] = self.heap[largest], self.heap[i]
            self._heapify(largest)
            
    def get_max(self):
        return self.heap[0] if self.heap else None
    
    def __str__(self):
        return str(self.heap)