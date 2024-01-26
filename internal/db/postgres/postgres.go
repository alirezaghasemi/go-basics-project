package postgres

import (
	"strings"
	"text/template"
	"database/sql"

	"github.com/alirezaghasemi/go-basics-project/pkg/config/config.go"
)

const connString = "postgres://{{.User}}:{{ .Pass }}@{{ .Host }}:{{ .Port }}/{{ .Database }}@sslmode=disable"

func buildConnectionStringOrPanic(cnf config.Postgres) string {
	// temp, err := template.New("ConnString").Parse(connString)

	// if err != nil {
	// 	panic(err)
	// }

	// bf := bytes.NewBuffer(make([]byte, 0))
	// os.Stdout
	sb := strings.Builder{}
	temp := template.Must(template.New("ConnString").Parse(connString))
	if err := temp.Execute(&sb, cnf); err != nil {
		panic(err)
	}

	return sb.String()
}
