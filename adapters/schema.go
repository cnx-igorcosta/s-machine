package adapters

type Schema struct {
	InputName    string
	Type         string
	OutputName   string
	Required     bool
	DefaultValue interface{}
}

func Transform(input map[string]interface{}, schemas []Schema) (output map[string]interface{}) {
	for _, s := range schemas {
		inputName := input[s.InputName]
	}
}
