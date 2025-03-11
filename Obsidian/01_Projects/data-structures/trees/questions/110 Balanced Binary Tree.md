[Question](https://leetcode.com/problems/balanced-binary-tree/description/)

It is quite similar to [[543 Diameter of Binary Tree]] check out more and more notes need to be added.

```python
class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        self.balanced = True
        self.depth(root)
        return self.balanced

    def depth(self, node):
        
        if node is None:
            return 0

        right_depth = self.depth(node.right)
        left_depth = self.depth(node.left)

        if abs(right_depth - left_depth) > 1:
            self.balanced = False

        return max(right_depth, left_depth) + 1
```

[[Drawing 2024-04-22 00.08.18.excalidraw]]