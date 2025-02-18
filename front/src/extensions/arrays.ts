interface Array<T> {
  /**
   * Returns a random element from the array.
   */
  random(): T;
  /**
   * Removes the first occurrence of a specific object from the array.
   * @param item The item to remove from the array.
   */
  remove(item: T): T[];
}

Array.prototype.random = function <T>() {
  const randomIndex = Math.floor(Math.random() * this.length);
  return this[randomIndex];
};

Array.prototype.remove = function <T>(item: T) {
  const index = this.indexOf(item);
  if (index === -1) return this;
  this.splice(index, 1);
  return this;
};
