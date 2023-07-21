package framework

const UnknownString = "unknown"
const UndefinedString = "undefined"

const EmptyString = ""
const EmptyStringAsInt = 0

func AdjustString(message string, additions []string) string {
	msg := message

	for i := 0; i < len(additions); i++ {
		msg = msg + additions[i]
	}

	return msg
}

func StringIsEmpty(value string) bool {
	if value == EmptyString 
	
}

func HandleString(value string, handling string) string {
	var res string

	if value.(type) == nil {

	}
}