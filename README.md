
# Go library for the ATLAS SMS API.

atlas-sms-gateway-go is a Go client library for accessing the ATLAS SMS API.

Where possible, the services available on the client groups the API into logical chunks and correspond to the structure of the ATLAS API documentation.

## Usage

``` go

import (
	atlas "atlas-sms-gateway-go"
	"fmt"
)

func main() {
	client := atlas.NewBasicAuthClient(API_KEY, API_SECRET)
	sms := atlas.Sms{
		To:      "+2507XXXXXX",
		Sender:  "XXXXXXXXX",
		Content: "Hello From the other side",
	}
	resp, err := client.SendSMS(&sms)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(resp))

}

```
