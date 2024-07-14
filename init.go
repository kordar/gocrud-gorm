package gocrud_gorm

import (
	"github.com/kordar/gocrud"
)

func InitExec() {

	gocrud.AddExecute("=", EQ, "gorm")
	gocrud.AddExecute("EQ", EQ, "gorm")
	gocrud.AddExecute("!=", NEQ, "gorm")
	gocrud.AddExecute("<>", NEQ, "gorm")
	gocrud.AddExecute("LT", LT, "gorm")
	gocrud.AddExecute("<", LT, "gorm")
	gocrud.AddExecute("LE", LE, "gorm")
	gocrud.AddExecute("<=", LE, "gorm")
	gocrud.AddExecute("GT", GT, "gorm")
	gocrud.AddExecute(">", GT, "gorm")
	gocrud.AddExecute("GE", GE, "gorm")
	gocrud.AddExecute(">=", GE, "gorm")
	gocrud.AddExecute("NEQ", NEQ, "gorm")
	gocrud.AddExecute("IN", IN, "gorm")
	gocrud.AddExecute("NOTIN", NOTIN, "gorm")
	gocrud.AddExecute("LIKE", LIKE, "gorm")
	gocrud.AddExecute("NOTLIKE", NOTLIKE, "gorm")
	gocrud.AddExecute("LIKELEFT", LIKELEFT, "gorm")
	gocrud.AddExecute("LIKERIGHT", LIKERIGHT, "gorm")
	gocrud.AddExecute("BETWEEN", BETWEEN, "gorm")
	gocrud.AddExecute("NOTBETWEEN", NOTBETWEEN, "gorm")
	gocrud.AddExecute("ISNULL", ISNULL, "gorm")
	gocrud.AddExecute("ISNOTNULL", ISNOTNULL, "gorm")

	gocrud.AddExecute("ASC", ASC, "gorm")
	gocrud.AddExecute("DESC", DESC, "gorm")

	gocrud.AddExecute("SAVE", SAVE, "gorm")
	gocrud.AddExecute("SETVAL", SETVAL, "gorm")
	gocrud.AddExecute("UPDATE", UPDATE, "gorm")
	gocrud.AddExecute("UPDATES", UPDATES, "gorm")
	gocrud.AddExecute("CREATE", CREATE, "gorm")
	gocrud.AddExecute("PAGE", PAGE, "gorm")
	gocrud.AddExecute("DELETE", DELETE, "gorm")

}
