package utils

//Create response message var
func Message(success bool, message string) map[string]interface{} {
	return map[string]interface{}{"success": success, "message": message}
}
