/*
	NODE.JS PROGRAM 2: VARIABLES, TYPES & THE FOR LOOP
	Concepts:
	- 'const' for constants (cannot reassign).
	- 'let' for variables (scoped to the block).
	- 'for' and 'while' loops.
	- Dynamic typing: Variables can hold any type without a declared type name.
*/

// Variables
const age = 30; // Constant
let name = "Alice"; // Let (changeable)

console.log(`User: ${name}, Age: ${age}`);

// Traditional For Loop
console.log("\nTraditional For Loop:");
for (let i = 1; i <= 5; i++) {
	console.log(`Iteration ${i}`);
}

// While Loop
console.log("\nWhile Loop:");
let counter = 0;
while (counter < 3) {
	console.log(`Counter is ${counter}`);
	counter++;
}

// Array Iteration (a common JS pattern)
const items = ["Apple", "Banana", "Cherry"];
console.log("\nArray Iteration:");
items.forEach((item, index) => {
	console.log(`${index}: ${item}`);
});
