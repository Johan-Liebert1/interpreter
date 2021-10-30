package constants

var INT_FLOAT_OPERATIONS = map[string]map[string]bool{
	INTEGER: {
		FLOAT:   true,
		INTEGER: true,
	},
	FLOAT: {
		FLOAT:   true,
		INTEGER: true,
	},
}

var INT_FLOAT_STRING_OPERATIONS = map[string]map[string]bool{
	INTEGER: {
		FLOAT:   true,
		INTEGER: true,
	},
	FLOAT: {
		FLOAT:   true,
		INTEGER: true,
	},
	STRING: {
		STRING: true,
	},
}

var VAR_TYPE_TO_TOKEN_TYPE = map[string]string{
	INTEGER_TYPE: INTEGER,
	FLOAT_TYPE:   FLOAT,
	STRING_TYPE:  STRING,
	BOOLEAN_TYPE: BOOLEAN,
}

/*
allowedOperation = {
	PLUS (the operation): {
		INTEGER (left supported operand): {
			FLOAT: true, (allowed right operand type)
			INTEGER: true (allowed right operand type)
		}
	}
}
*/
var ALLOWED_OPERATIONS_ON_TYPES = map[string]map[string]map[string]bool{
	PLUS: {
		INTEGER: {
			FLOAT:   true,
			INTEGER: true,
		},
		FLOAT: {
			FLOAT:   true,
			INTEGER: true,
		},
		STRING: {
			STRING: true,
		},
	},
	MUL: {
		INTEGER: {
			FLOAT:   true,
			INTEGER: true,
		},
		FLOAT: {
			FLOAT:   true,
			INTEGER: true,
		},
		STRING: {
			INTEGER: true,
		},
	},
	MODULO: {
		INTEGER: {
			INTEGER: true,
		},
	},
	MINUS:       INT_FLOAT_OPERATIONS,
	FLOAT_DIV:   INT_FLOAT_OPERATIONS,
	INTEGER_DIV: INT_FLOAT_OPERATIONS,
	EXPONENT:    INT_FLOAT_OPERATIONS,

	GREATER_THAN:          INT_FLOAT_STRING_OPERATIONS,
	GREATER_THAN_EQUAL_TO: INT_FLOAT_STRING_OPERATIONS,
	LESS_THAN:             INT_FLOAT_STRING_OPERATIONS,
	LESS_THAN_EQUAL_TO:    INT_FLOAT_STRING_OPERATIONS,
	EQUALITY:              INT_FLOAT_STRING_OPERATIONS,
	NOT_EQUAL_TO:          INT_FLOAT_STRING_OPERATIONS,
}
