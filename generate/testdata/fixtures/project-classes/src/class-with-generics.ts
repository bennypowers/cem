export class Repository<Item extends { id: string }> {
  items: Item[] = [];

  add(item: Item): number {
    this.items.push(item);
    return this.items.length;
  }

  map<Result>(project: (item: Item) => Result): Result[] {
    return this.items.map(project);
  }
}
