export /** In-statement doc. */ class InsideExport extends Event {}

export default /** In-statement default doc. */ class InsideDefaultExport extends Event {}

function decorator() {
  return function (target: any) { return target; };
}

@decorator()
/** In-statement doc after decorator. */
export class DecoratedInsideExport extends Event {}
