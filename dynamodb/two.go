package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var client *dynamodb.Client

func initClient() {
	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1"))
	client = dynamodb.NewFromConfig(cfg)
}

// 1️⃣ Create Table
func createTable() {
	_, err := client.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		TableName: "Students",
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: "id",
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: "id",
				KeyType:       types.KeyTypeHash,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	})

	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}
	fmt.Println("Table Created")
}

// 2️⃣ Insert Item
func insertItem() {
	_, err := client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: "Students",
		Item: map[string]types.AttributeValue{
			"id":   &types.AttributeValueMemberS{Value: "1"},
			"name": &types.AttributeValueMemberS{Value: "Jen"},
			"age":  &types.AttributeValueMemberN{Value: "21"},
		},
	})

	if err != nil {
		fmt.Println("Error inserting:", err)
		return
	}
	fmt.Println("Item Inserted")
}

// 3️⃣ Update Item
func updateItem() {
	_, err := client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: "Students",
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: "1"},
		},
		UpdateExpression: "SET age = :newAge",
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":newAge": &types.AttributeValueMemberN{Value: "22"},
		},
	})

	if err != nil {
		fmt.Println("Error updating:", err)
		return
	}
	fmt.Println("Item Updated")
}

// 4️⃣ Delete Table
func deleteTable() {
	_, err := client.DeleteTable(context.TODO(), &dynamodb.DeleteTableInput{
		TableName: "Students",
	})

	if err != nil {
		fmt.Println("Error deleting table:", err)
		return
	}
	fmt.Println("Table Deleted")
}

func main() {
	initClient()

	createTable()
	time.Sleep(5 * time.Second) // wait for table creation

	insertItem()
	updateItem()
	deleteTable()
}