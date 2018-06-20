# go-servicebus

### Usage
SbConfig struct:

Key = key value in service bus

KeyType = type of key, i.e. RootManageSharedAccessKey

Endpoint = https://<service_bus_name>.servicebus.windows.net

```
type SbConfig struct {
	Key string
	KeyType string
	Endpoint string
}
```
SbMessage struct:

Body = message body

Endpoint = queue or topic name

Props = key/value pairs to be sent as message properties

```
type SbMessage struct {
	Body string
	Endpoint string
	Props map[string]string
}
```
Sample message:
```
{
	"Body": "sample message",
	"Props": {
		"property1": "value1",
		"property2": "value2"
	},
	"Endpoint": "<queue_or_topic>"
}

```

1. Create/serialize json to SbMessage
2. Create a SbConfig item similar to below
2. Send Message
```
jsonString := "json string matching the sample message shown above"
var msg servicebus.SbMessage
json.Unmarshal([]byte(jsonString), &msg)
config := &sb.SbConfig{SB_KEY, SB_KEYTYPE, SB_URL}
success, err := config.Send(&msg) //returns (bool,error)
```
