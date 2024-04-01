package dynamodb

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
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

