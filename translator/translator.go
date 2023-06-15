package translator

type Translator struct {
}

type ITranslator interface {
	Translate(messageData []byte, messageAttributes map[string]string) TranslatedData
}

func New() ITranslator {
	return &Translator{}
}
