export class Repository<Item extends { id: string }> {
  items: Item[] = [];

  constructor(public initial: Item) {
    this.items = [initial];
  }

  get first(): Item | undefined {
    return this.items[0];
  }

  add(item: Item): number {
    this.items.push(item);
    return this.items.length;
  }

  map<Result>(project: (item: Item) => Result): Result[] {
    return this.items.map(project);
  }

  createScoped() {
    class Scoped<Item> {
      item: Item | undefined;
    }
    return Scoped;
  }
}
