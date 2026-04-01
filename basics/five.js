function asyncTask() {
  return new Promise(resolve => {
    setTimeout(() => resolve("Task done"), 2000);
  });
}

async function run() {
  console.log("Starting task...");
  const result = await asyncTask();
  console.log(result);
}

run();