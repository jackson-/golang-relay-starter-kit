package data

// Data model structs

type Operation struct {
	Id   string `json:"id"`
	Title string `json:"title"`
}

type OperationList struct {
	Id   string `json:"id"`
	Operations []Operation `json:"operations"`
}

// Mock data

var currentOperation = &Operation{
	Id: "3",
	Title: "Max",
}

var allOperations = &OperationList{
		Id:"1",
		Operations:[]Operation {
		{
			Id: "1",
			Title: "Mean",
		},
		{
			Id: "2",
			Title: "Median",
		},
		{
			Id: "3",
			Title: "Max",
		},
		{
			Id: "4",
			Title: "Min",
		},
		{
			Id: "5",
			Title: "STD",
		},
		{
			Id: "6",
			Title: "Sum",
		},
		{
			Id: "7",
			Title: "Product",
		},
	},
}

// Data getters/setters
func GetOperation() *Operation {
	return currentOperation
}

func GetAllOperations() *OperationList{
	return allOperations
}
