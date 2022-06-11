package dynamo

import (
	"fmt"
	"myAppSyncApp-API/serverless/appsync"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

/* ID                string      `json:"id"`
Email             string      `json:"email"`
OrderStatus       OrderStatus `json:"orderStatus"`
FirstName         string     `json:"firstName"`
LastName          string     `json:"lastName"`
PrescriberAddress string     `json:"prescriberAddress"`
DeliveryAddress   string     `json:"deliveryAddress"`
ExtraDetails      string     `json:"extraDetails"`
FutureOrderDate   string     `json:"futureOrderDate"`
PaymentAmount     string     `json:"paymentAmount"`
BrainTreeID       string      `json:"brainTreeId"`  How come this is presnt at the time of crating order ??*/

// CreateAppointment create a course and the first appointment.

/*

{
    "id": "order-user_b944efb2-466f-4f17-8fd9-380dbacf12831652100469",
    "email": "ronghester@gmail.com",
    "orderStatus": "preparing",
    "orderItems": [
        {
            "brandDesc": "Complete Blink-N-Clean lens drops",
            "prodDesc": "Complete Blink-N-Clean lens drops (AMO UK Ltd)",
            "drugName": "Complete Blink-N-Clean lens drops (AMO UK Ltd)",
            "variantName": "Complete Blink-N-Clean lens drops",
            "vtmId": "",
            "vmpId": "10748911000001102",
            "packages": null,
            "isRestricted": false,
            "form": "Other",
            "ingredients": [],
            "iq": "",
            "isDiscontinued": false,
            "ampId": "10734211000001104",
            "barcodes": null,
            "quantity": 1
        }
    ],
    "prescriberAddress": ", Shap Health Centre, Peggy Nut Croft, Penrith, CA10 3LW",
    "deliveryAddress": "20 Concord Way, undefined, Dukinfield, SK16 4DB",
    "extraDetails": "",
    "futureOrderDate": "2022-05-17T12:48:00.000Z",
    "freeDisposable": true,
    "paymentAmount": null,
    "brainTreeId": null
}
*/

func (s Store) CreateTodoDB(in appsync.NewTodo) (out appsync.Todo, err error) {

	id := s.generateID()

	out = appsync.Todo{
		ID:   id,
		Text: in.Text,
		User: &appsync.User{
			ID: in.UserID,
		},
		Done: false,
	}

	appRec := newTodoRecord(out)

	appItem, err := dynamodbattribute.MarshalMap(appRec)

	if err != nil {
		return out, fmt.Errorf("dynamo.Store.Create todo error marshalling  record: %v", err)
	}

	_, err = s.db.PutItem(&dynamodb.PutItemInput{
		TableName: &s.tableName,
		Item:      appItem,
	})

	if err != nil {
		return out, fmt.Errorf("dynamo.Store.Create Todo: error in saving : %v", err)
	}

	//out = appItem
	return out, nil
}

// Performing the update via db data resolver.
/*func (s Store) UpdateTodo(pk string, sk string, done bool) (out appsync.Todo, err error) {

	_, err = s.db.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String(s.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: aws.String(fmt.Sprintf("user_%s", pk))},
			"sk": {S: aws.String(sk)},
		},
		UpdateExpression: aws.String("SET done = :done"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":done": {BOOL: aws.Bool(done)},
		},
		ConditionExpression: aws.String("attribute_exists(pk)"),
	})

	if err != nil {
		fmt.Printf("In error -- %+v", err)
		return out, fmt.Errorf("dynamo.Store.Update : error updating record: %v", err)
	}

	res_todo, _, err := s.GetTodo(pk, sk)

	if err != nil {
		return out, fmt.Errorf("dynamo.Store.Update : error updating record: %v", err)
	}
	return res_todo, nil
}

func (s Store) GetTodo(pk string, sk string) (out appsync.Todo, found bool, err error) {

	res, err := s.db.Query(&dynamodb.QueryInput{
		TableName:              aws.String(s.tableName),
		KeyConditionExpression: aws.String("pk=:pk AND  sk = :sk"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pk": {S: aws.String(fmt.Sprintf("user_%s", pk))},
			":sk": {S: aws.String(sk)},
		},
	})

	if err != nil {
		return out, false, fmt.Errorf("dynamo.Store: error getting record: %v", err)
	}
	if len(res.Items) == 0 {
		return out, false, nil
	}

	var rec appsync.Todo
	err = dynamodbattribute.UnmarshalMap(res.Items[0], &rec)
	if err != nil {
		return out, false, fmt.Errorf("dynamo.Store.Get ToDO: error: %v", err)
	}

	return rec, true, nil

}*/
