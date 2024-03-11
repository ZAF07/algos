class Node: 
    def __init__(self, val, left= None, right=None) -> None:
       self.val = val
       self.left = left  
       self.right = right 


    # add a new node
    def add_node(self, value: int) -> bool:
        # TODO: Implement recursive insert for binary tree
        def add_new_leaf(node, val):
            # Larger than current node
            if val > node.val and node.right is None:
                node.right = Node(val, None, None)
                return True

            # Less than current node
            if val < node.val and node.left is None:
                node.left = Node(val, None, None)
                return True

            if val > node.val and node.right is not None:
                return add_new_leaf(node.right, val)
            
            if val < node.val and node.left is not None:
                return add_new_leaf(node.left, val)
            return False

        return add_new_leaf(self, value)

    # Get sum of values within given min and max range
    def sum(self, min: int, max:int) -> int:
        def _sum(node, l, r):
            sum = 0
            if node.val >= l and node.val <= r:
                sum += node.val
            
            if node.left is not None:
                sum += _sum(node.left, l, r)
            if node.right is not None:
                sum += _sum(node.right, l, r)
            return sum
        
        return _sum(self, min, max)

    # Search for value using DFS
    def search_dfs(self, target: int) -> int:

        def _dfs_recursive(node, target) -> int:
            if node.val == target:
                return node.val
        
            if node.val > target and node.left is not None:
                return _dfs_recursive(node.left, target)

            if node.val < target and node.right is not None:
                return _dfs_recursive(node.right, target)
            
            return -1
        
        return _dfs_recursive(self, target)
    
    # Counts the number of nodes in the BST
    def count_nodes(self):
        def count(node):
            sum = 0

            if node.left is not None:
                sum += 1
                sum += count(node.left)
            
            if node.right is not None:
                sum += 1
                sum += count(node.right)
            return sum
        return count(self)




root = Node(3)
print('Add -> ', root.add_node(2))
print('Add -> ', root.add_node(1))
print('Add -> ', root.add_node(4))
print('Add -> ', root.add_node(5))
print('Add -> ', root.add_node(5))
print('Sum -> ', root.sum(1,5))
print('Count node --> ', root.count_nodes())
