
var node4 = {
  'name': 'Node 4',
  'val': 4,
  'left': null,
  'right': null,
}
var node3 = {
  'name': 'Node 3',
  'val': 3,
  'left': null,
  'right': null,
};
var node2 = {
  'name': 'Node 2',
  'val': 2,
  'left': node4,
  'right': null,
};

var node1 = {
  'name': 'Node 1',
  'val': 1,
  'left': node2,
  'right': node3,
}


const result = [];
const dfs = (node, target) => {
  console.log('NODE: ', node.name);
  if (node.val == target) {
    return node.val
    // return node.val
  }
  const left = node.left;
  const right = node.right; 

  if (left != null) {
   return dfs(node.left, target)
  }
  
  if (right != null) {
     return dfs(node.right, target)
  }

}

console.log(dfs(node1, 2));
