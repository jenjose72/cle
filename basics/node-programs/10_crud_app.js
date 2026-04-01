const http = require('http');

/*
	NODE.JS PROGRAM 10: SIMPLE CRUD API (In-Memory)
	Concepts:
	- RESTful API structure with Node.js built-ins.
	- JSON parsing and stringifying.
	- Using an Array as an in-memory database.
	- Handling body data in an async/streaming fashion.
*/

// In-Memory Database (equiv for Go's Map)
let tasks = [
	{ id: 1, title: "Learn JavaScript", completed: false }
];
let nextID = 2;

// Utility to parse JSON body from the incoming request stream
async function getBody(req) {
	return new Promise((resolve, reject) => {
		let body = '';
		req.on('data', chunk => body += chunk.toString());
		req.on('end', () => resolve(JSON.parse(body || '{}')));
		req.on('error', err => reject(err));
	});
}

const server = http.createServer(async (req, res) => {
	res.setHeader('Content-Type', 'application/json');

	const urlParts = req.url.split('/');
	const id = parseInt(urlParts[2]); // For /tasks/1

	// ROUTE: /tasks (List/Create)
	if (req.url === '/tasks') {
		if (req.method === 'GET') {
			// READ ALL
			res.end(JSON.stringify(tasks));
		} else if (req.method === 'POST') {
			// CREATE
			try {
				const body = await getBody(req);
				const newTask = { id: nextID++, ...body };
				tasks.push(newTask);
				res.statusCode = 201;
				res.end(JSON.stringify(newTask));
			} catch (err) {
				res.statusCode = 400;
				res.end(JSON.stringify({ error: "Invalid JSON" }));
			}
		}
	} 
	// ROUTE: /tasks/:id (GET/PUT/DELETE)
	else if (urlParts[1] === 'tasks' && id) {
		const taskIndex = tasks.findIndex(t => t.id === id);
		if (taskIndex === -1) {
			res.statusCode = 404;
			res.end(JSON.stringify({ error: "Task not found" }));
			return;
		}

		if (req.method === 'GET') {
			// READ SINGLE
			res.end(JSON.stringify(tasks[taskIndex]));
		} else if (req.method === 'PUT') {
			// UPDATE
			const body = await getBody(req);
			tasks[taskIndex] = { ...tasks[taskIndex], ...body };
			res.end(JSON.stringify(tasks[taskIndex]));
		} else if (req.method === 'DELETE') {
			// DELETE
			tasks.splice(taskIndex, 1);
			res.statusCode = 204;
			res.end();
		}
	} else {
		res.statusCode = 404;
		res.end(JSON.stringify({ error: "Endpoint not found" }));
	}
});

const PORT = 3002;
server.listen(PORT, () => {
	console.log(`CRUD Server starting on http://localhost:${PORT}...`);
	console.log("Try these endpoints:");
	console.log(` - GET/POST http://localhost:${PORT}/tasks`);
	console.log(` - GET/PUT/DELETE http://localhost:${PORT}/tasks/1`);
});
