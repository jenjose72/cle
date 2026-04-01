import {
  DynamoDBClient,
  CreateTableCommand,
  DeleteTableCommand,
  PutItemCommand,
  UpdateItemCommand
} from "@aws-sdk/client-dynamodb";

const client = new DynamoDBClient({
  region: "ap-south-1", // Chennai nearby 😄
});

// 1️⃣ Create Table
async function createTable() {
  const params = {
    TableName: "Students",
    AttributeDefinitions: [
      { AttributeName: "id", AttributeType: "S" }
    ],
    KeySchema: [
      { AttributeName: "id", KeyType: "HASH" }
    ],
    BillingMode: "PAY_PER_REQUEST"
  };

  try {
    const data = await client.send(new CreateTableCommand(params));
    console.log("Table Created:", data);
  } catch (err) {
    console.error("Error creating table:", err);
  }
}

// 2️⃣ Insert Item (needed before update)
async function insertItem() {
  const params = {
    TableName: "Students",
    Item: {
      id: { S: "1" },
      name: { S: "Jen" },
      age: { N: "21" }
    }
  };

  await client.send(new PutItemCommand(params));
  console.log("Item inserted");
}

// 3️⃣ Modify Field (Update)
async function updateItem() {
  const params = {
    TableName: "Students",
    Key: {
      id: { S: "1" }
    },
    UpdateExpression: "SET age = :newAge",
    ExpressionAttributeValues: {
      ":newAge": { N: "22" }
    }
  };

  try {
    const data = await client.send(new UpdateItemCommand(params));
    console.log("Item Updated:", data);
  } catch (err) {
    console.error("Error updating:", err);
  }
}

// 4️⃣ Delete Table
async function deleteTable() {
  const params = {
    TableName: "Students"
  };

  try {
    const data = await client.send(new DeleteTableCommand(params));
    console.log("Table Deleted:", data);
  } catch (err) {
    console.error("Error deleting table:", err);
  }
}

// Run sequence
async function main() {
  await createTable();
  await new Promise(r => setTimeout(r, 5000)); // wait for table creation
  await insertItem();
  await updateItem();
  await deleteTable();
}

main();