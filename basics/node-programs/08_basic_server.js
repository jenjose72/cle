const http = require('http');
const url = require('url');

/*
	NODE.JS PROGRAM 8: BASIC HTTP SERVER
	Concepts:
	- Using the built-in 'http' module.
	- Handling 'request' and 'response'.
	- Parsing URLs and query parameters.
*/

const server = http.createServer((req, res) => {
	// Parse the URL to get the path and query
	const parsedUrl = url.parse(req.url, true);
	
	// Set the response header
	res.setHeader('Content-Type', 'text/plain');
	
	if (parsedUrl.pathname === '/') {
		// Home logic
		res.end("Welcome to the Node.js Home Page!");
	} else if (parsedUrl.pathname === '/hello') {
		// Extract query parameter (e.g., /hello?name=Coder)
		const name = parsedUrl.query.name || "Guest";
		res.end(`Hello, ${name}!`);
	} else {
		res.statusCode = 404;
		res.end("Page not found");
	}
});

const PORT = 3000;
server.listen(PORT, () => {
	console.log(`Server starting on http://localhost:${PORT}...`);
	console.log(`Try visiting http://localhost:${PORT} or http://localhost:${PORT}/hello?name=YourName`);
});
