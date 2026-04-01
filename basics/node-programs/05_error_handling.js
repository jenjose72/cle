/*
	NODE.JS PROGRAM 5: ERROR HANDLING (TRY-CATCH)
	Concepts:
	- 'try...catch' for catching runtime errors.
	- Creating and 'throwing' custom Error objects.
	- The 'finally' block (runs cleanup regardless).
*/

function divide(a, b) {
	// Custom error check
	if (b === 0) {
		throw new Error("Cannot divide by zero!");
	}
	return a / b;
}

// Success Scenario
try {
	console.log("Success Result: 10 / 2 =", divide(10, 2));
} catch (err) {
	console.log("Caught Error:", err.message);
}

// Error Scenario
console.log("\nError Scenario:");
try {
	const result = divide(10, 0);
	console.log("Result is:", result);
} catch (err) {
	console.log("Caught Error (as expected):", err.message);
} finally {
	console.log("Operation Complete (This runs anyway).");
}
