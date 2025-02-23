interface IMyInterface {
  listKey(name: string): void;
}

class MyClass implements IMyInterface {
  listKey(name: string) {
    console.log(name)
  }
}

export default new MyClass();
