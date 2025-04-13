class MinHeap:
    def __init__(self):
        self.heap = []
        
    def _parent(self, i):
        return (i-1) // 2
    
    def _left(self, i):
        return (2 * i) + 1
    
    def _right(self, i):
        return (2 * i) + 2