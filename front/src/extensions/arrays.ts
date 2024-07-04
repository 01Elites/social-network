interface Array<T> {
  random(): T;
}

Array.prototype.random = function <T>(): T {
  const randomIndex = Math.floor(Math.random() * this.length);
  return this[randomIndex];
};
