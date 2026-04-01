/*
	NODE.JS PROGRAM 6: CONCURRENCY (ASYNC/AWAIT & PROMISES)
	Concepts:
	- 'async' and 'await' keywords for handling non-blocking tasks.
	- 'setTimeout' for delayed execution (simulating work).
	- Managing multiple tasks concurrently using 'Promise.all'.
*/

// Function returning a Promise (to simulate async work)
const work = (id) => {
	return new Promise((resolve) => {
		console.log(`Worker ${id}: Starting work...`);
		const delay = id * 500; // ms
		setTimeout(() => {
			console.log(`Worker ${id}: Finished work.`);
			resolve(id);
		}, delay);
	});
};

async function main() {
	console.log("Concurrency (Running 3 workers concurrently):");
	
	// Create an array of tasks (Promises)
	const tasks = [work(1), work(2), work(3)];
	
	// Wait for ALL to finish!
	const results = await Promise.all(tasks);
	console.log("\nAll workers finished. Results:", results);
	
	// Simple delayed greeting
	const greet = async () => "Greetings from an Async Function!";
	console.log("\nReceived via Async:", await greet());
}

main();
