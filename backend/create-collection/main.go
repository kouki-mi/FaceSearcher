package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"encoding/json"
)

//リクエストボディ
type MyRequestBody struct {
	CollectionId string `json:"collectionId"`
}

// awsRekognitionのコレクションを作成する
func createCollection(svc *rekognition.Rekognition, collectionId string) (*rekognition.CreateCollectionOutput, error){
	// コレクションを作成する
	input := &rekognition.CreateCollectionInput{
		CollectionId: aws.String(collectionId),
	}

	result, err := svc.CreateCollection(input)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	fmt.Println(result)
	return result, nil
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// リクエストボディからコレクションIDを取得
	requestBody := request.Body
	var myRequestBody MyRequestBody
	err := json.Unmarshal([]byte(requestBody), &myRequestBody) 
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, nil
	}

	// セッションを作成
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Rekognitionサービスのクライアントを作成
	svc := rekognition.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))
	result, err := createCollection(svc, myRequestBody.CollectionId)

	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		Body:       result.String(),
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(handler)
}