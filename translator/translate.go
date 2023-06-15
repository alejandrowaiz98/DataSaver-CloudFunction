package translator

func (t *Translator) Translate(messageData []byte, messageAttributes map[string]string) TranslatedData {

	td := TranslatedData{}

	td.Origen = messageAttributes["Origen"]
	td.Date = messageAttributes["Date"]
	td.DestinyCollection = messageAttributes["DestinyCollection"]
	td.MessageID = messageAttributes["MessageID"]

	//TODO: find if data message is sended in json format or byte format. If byte format, then this should have a decoder!
	td.Message = messageData

	return td
}
