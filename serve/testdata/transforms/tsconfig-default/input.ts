async function test() {
	const obj = { a: 1, b: 2 };
	const { a, ...rest } = obj;
	return rest;
}
