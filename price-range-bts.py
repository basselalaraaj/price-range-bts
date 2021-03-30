from typing import List


class BinarySearchTree:
    def __init__(self):
        self.root = None

    def insert(self, val: int) -> None:
        if not self.root:
            self.root = Node(val)
        else:
            self.root.insert(self.root, val)


class Node:
    def __init__(self, val):
        self.left = None
        self.val = val
        self.right = None

    def insert(self, tree: 'Node', val: int) -> 'Node':
        if not tree:
            return Node(val)
        if val < tree.val:
            tree.left = self.insert(tree.left, val)
        if val > tree.val:
            tree.right = self.insert(tree.right, val)

        return tree


def productsInRange(root: 'Node', low: int, high: int) -> List[int]:
    orderedPrices: List[int] = []

    def preorder(node: 'Node'):
        if not node:
            return
        if low <= node.val <= high:
            orderedPrices.append(node.val)
        if low <= node.val:
            preorder(node.left)
        if node.val <= high:
            preorder(node.right)

    preorder(root)
    return orderedPrices


bst = BinarySearchTree()
bst.insert(9)
bst.insert(6)
bst.insert(14)
bst.insert(20)
bst.insert(1)
bst.insert(30)
bst.insert(8)
bst.insert(17)
bst.insert(5)

print(productsInRange(bst.root, 7, 20))
