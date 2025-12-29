export enum LogLevel {
  Debug,
  Info,
  Warn,
  Error
}

export class Logger {
  constructor(private level: LogLevel = LogLevel.Info) {}

  log(message: string, level: LogLevel = LogLevel.Info): void {
    if (level >= this.level) {
      console.log(`[${LogLevel[level]}] ${message}`);
    }
  }
}
