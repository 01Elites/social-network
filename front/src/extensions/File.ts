  interface File {
    toBase64(): Promise<string>;
  }


File.prototype.toBase64 = function (): Promise<string> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => resolve(reader.result as string);
    reader.onerror = () => reject(reader.error);
    reader.readAsDataURL(this);
  });
}