declare class Player {
  // ambient declarations carry only body-less signatures; they are emitted
  // as members, and there is no implementation signature to drop
  onReady(cb: () => void): void;
  pauseVideo(): void;
}
