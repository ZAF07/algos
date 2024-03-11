class BSTNode {
  constructor(value) {
    this.val = value;
    this.left = undefined;
    this.right = undefined;
  }

  /**
   *
   * @param {number} val
   * @returns {bool}
   * @description Takes a new value as a parameter and recursively tries to find a placement in the BST for the new value. Returns True if successful and false otherwise
   */
  addNode(val) {
    const newNode = new BSTNode(val);

    const addNewNode = (node, val) => {
      if (val < node.val && node.left == undefined) {
        node.left = newNode;
        return true;
      }

      if (val > node.val && node.right == undefined) {
        node.right = newNode;
        return true;
      }

      if (val < node.val && node.left != undefined) {
        return addNewNode(node.left, val);
      }

      if (val > node.val && node.right != undefined) {
        return addNewNode(node.right, val);
      }
      return false;
    };

    return addNewNode(this, val);
  }

  /**
   *
   * @param {number} min
   * @param {number} max
   * @returns {number}
   * @description Takes a min and max range and returns the sum of node values in the BST whose values fall between the range
   */
  sum(min, max) {
    const sumValues = (node, min, max) => {
      let sum = 0;
      if (node.val >= min && node.val <= max) {
        sum += node.val;
      }

      if (node.left != undefined) {
        sum += sumValues(node.left, min, max);
      }

      if (node.right != undefined) {
        sum += sumValues(node.right, min, max);
      }
      return sum;
    };
    return sumValues(this, min, max);
  }
}

const node = new BSTNode(3);
console.log(node.addNode(2));
console.log(node.addNode(4));
console.log("Sum --> ", node.sum(1, 6));
