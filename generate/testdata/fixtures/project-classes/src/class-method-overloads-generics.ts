export class Lookup<Key extends string> {
  find(key: Key): string | undefined;
  find(keys: Key[]): (string | undefined)[];
  find(keyOrKeys: Key | Key[]): string | undefined | (string | undefined)[] {
    return undefined;
  }

  map<Value>(callback: (key: Key) => Value): Value[];
  map<Value>(callback: (key: Key, index: number) => Value): Value[];
  map<Value>(callback: (key: Key, index?: number) => Value): Value[] {
    return [];
  }
}
