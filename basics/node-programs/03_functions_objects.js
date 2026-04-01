/*
	NODE.JS PROGRAM 3: FUNCTIONS, OBJECTS & ARROW FUNCTIONS
	Concepts:
	- Traditional function declarations vs. Arrow functions (ES6).
	- Passing objects to functions (equivalent to memory reference in JS).
*/

// Standard function
function add(a, b) {
	return a + b;
}

// Arrow function (more concise)
const multiply = (a, b) => a * b;

console.log("Add: 5 + 5 =", add(5, 5));
console.log("Mult: 10 * 2 =", multiply(10, 2));

// JS Objects (similar to Go's structs, but dynamic)
let user = {
	name: "Jack",
	score: 10
};

console.log("\nOriginal User Score:", user.score);

// Objects are passed by reference in JS
function updateScore(obj) {
	obj.score += 50; // Modifies the actual object
}

updateScore(user);
console.log("After updateScore User Score:", user.score);
