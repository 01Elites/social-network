interface File {
  toBase64(): Promise<string>;
}

File.prototype.toBase64 = function (): Promise<string> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => {
      const b64 = (reader.result as string).split(',')[1];
      resolve(b64);
    };
    reader.onerror = () => reject(reader.error);
    reader.readAsDataURL(this);
  });
};
