const fs = require('fs/promises');

/*
	NODE.JS PROGRAM 7: FILE I/O (READING & WRITING)
	Concepts:
	- 'fs/promises' for easy async file operations.
	- 'await fs.writeFile()' and 'await fs.readFile()'.
	- Converting data between strings and buffers.
*/

async function main() {
	const filename = "testfile_node.txt";
	const content = "Line 1: Node.js is powerful!\nLine 2: File handling is fully asynchronous.";
	
	// WRITE FILE
	console.log(`Writing to: ${filename}`);
	try {
		await fs.writeFile(filename, content);
		console.log("File written successfully!");
	} catch (err) {
		console.error("Error writing file:", err);
		return;
	}
	
	// READ FILE
	console.log("\nReading from file:");
	try {
		// utf8 encoding is important for reading text
		const data = await fs.readFile(filename, 'utf8');
		console.log("File Content Found:");
		console.log(data);
	} catch (err) {
		console.error("Error reading file:", err);
	}
}

main();
