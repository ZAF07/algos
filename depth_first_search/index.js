var node3 = {
  'val': 3,
  'left': null,
  'right': null,
};
var node2 = {
  'val': 2,
  'left': node3,
  'right': null,
};

var node1 = {
  'val': 1,
  'left': null,
  'right': node2,
}


const result = [];
const dfs = (node, res) => {
  if (node.val != null || node.val != undefined) {
    res.push(node.val)
  }
  
  const left = node.left;
  const right = node.right; 

  if (left != null) {
  dfs(node.left, res)
  }
  
  if (right != null) {
    dfs(node.right, res)
  }

}

dfs(node1, result)
console.log(result);
