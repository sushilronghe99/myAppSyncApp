package dynamo

import (
	"fmt"
	"myAppSyncApp-API/serverless/appsync"
)

type recordKey struct {
	PK string `json:"pk"`
	SK string `json:"sk"`
}

type todoRecord struct {
	recordKey
	appsync.Todo
}

func newTodoRecord(c appsync.Todo) todoRecord {
	//data := fmt.Sprintf("email_%s", c.Email)

	return todoRecord{
		recordKey: recordKey{
			PK: fmt.Sprintf("user_%s", c.User.ID),
			SK: c.ID,
		},
		Todo: c,
	}
}
