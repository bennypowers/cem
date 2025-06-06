export class Simple {}

export class Subclass extends Simple {}

@decorator
@decorator()
@decorator('decorator')
export class Decorated extends Simple {}

class Foo {}
