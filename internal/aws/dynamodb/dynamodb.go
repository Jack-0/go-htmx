package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDBService struct {
	client *dynamodb.DynamoDB
}


func NewDynamoDBService(region, endpoint string) (*DynamoDBService, error) {
	config := aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(endpoint),
	}
	sess, err := session.NewSession(&config)
	if err != nil {
		return nil, err
	}

	svc := dynamodb.New(sess)
	return &DynamoDBService{client: svc}, nil
}

func (d *DynamoDBService) AddItem(table string, item any) error {

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Marshalling map err:", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(table),
	}

	_, err = d.client.PutItem(input)

	if err != nil {
		fmt.Println("Put item err:", err)
		return err
	}

	return err
}

func (d *DynamoDBService) GetItem(table string, pk string, sk string) (*dynamodb.GetItemOutput, error) {
	key := map[string]*dynamodb.AttributeValue{
		pk: {
			S: aws.String(sk),
		},
	}
	input := &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key:       key,
	}

	result, err := d.client.GetItem(input)

	if err != nil {
		fmt.Println("Get item err:", err)
		return nil, err
	}

	return result, err
}

func (d *DynamoDBService) ListTables() ([]string, error) {
	result, err := d.client.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		return nil, err
	}

	tables := make([]string, len(result.TableNames))
	for i, table := range result.TableNames {
		tables[i] = *table
	}
	return tables, nil
}
