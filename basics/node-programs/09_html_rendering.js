const http = require('http');

/*
	NODE.JS PROGRAM 9: HTML RENDERING & TEMPLATES
	Concepts:
	- Serving HTML strings from a node server.
	- Using template literals (ES6) for dynamic content in HTML string.
	- Loop rendering in a string-based template.
*/

// PageData (equivalent to the Go struct)
const pageData = {
	title: "Node.js Programming Language",
	author: "Ryan Dahl",
	language: "JavaScript (V8)",
	features: ["Non-blocking I/O", "V8 Engine", "NPM Ecosystem", "Unified Language (Front/Back)"]
};

const server = http.createServer((req, res) => {
	// Set correct MIME type to render text as HTML
	res.setHeader('Content-Type', 'text/html');
	
	// Create the HTML template string
	const html = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>${pageData.title}</title>
		<style>
			body { font-family: Arial, sans-serif; padding: 2em; background-color: #f4f4f4; }
			h1 { color: #007d9c; }
			ul { color: #333; }
		</style>
	</head>
	<body>
		<h1>${pageData.title}</h1>
		<p><strong>Developed by:</strong> ${pageData.author}</p>
		<p><strong>Language:</strong> ${pageData.language}</p>
		<h3>Key Features:</h3>
		<ul>
			${pageData.features.map(f => `<li>${f}</li>`).join('')}
		</ul>
	</body>
	</html>
	`;
	
	res.end(html);
});

const PORT = 3001;
server.listen(PORT, () => {
	console.log(`Server starting on http://localhost:${PORT}...`);
});
