class MinHeap:
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
        while i != 0 and self.heap[self._parent(i)] > self.heap[i]:
            self.heap[self._parent(i)], self.heap[i] = self.heap[i], self.heap[self._parent(i)]
            i = self._parent(i) 
    
    
    def _heapify(self, i):
        smallest = i
        
        left = self._left(i)
        right = self._right(i)
        
        #check left
        if i < len(self.heap) and left < len(self.heap) and self.heap[left] < self.heap[smallest]:
            smallest = left
        if i < len(self.heap) and right < len(self.heap) and self.heap[right] < self.heap[smallest]:
            smallest = right
            
        if smallest != i:
            self.heap[i], self.heap[smallest] = self.heap[smallest], self.heap[i]
            self._heapify(smallest)
        
    def extract_min(self):
        if not self.heap:
            return None
        
        if len(self.heap) == 1:
            return self.heap.pop()
        
        root = self.heap[0]
        
        self.heap[0] = self.heap.pop()
        self._heapify(0)
        return root
        
    def get_min(self):
        return self.heap[0] if self.heap else None
    
    def __str__(self):
        return str(self.heap)
        
        