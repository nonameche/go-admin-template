package tables

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUsersTable(ctx *context.Context) table.Table {

	users := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := users.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Gender", "gender", db.Tinyint)
	info.AddField("City", "city", db.Varchar)
	info.AddField("Ip", "ip", db.Varchar)
	info.AddField("Phone", "phone", db.Varchar)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("users").SetTitle("Users").SetDescription("Users").SetAction(template.HTML(`<a href="http://google.com"><i class="fa fa-google"></i></a>`))

	formList := users.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Gender", "gender", db.Tinyint, form.Number)
	formList.AddField("City", "city", db.Varchar, form.Text)
	formList.AddField("Ip", "ip", db.Varchar, form.Ip)
	formList.AddField("Phone", "phone", db.Varchar, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("users").SetTitle("Users").SetDescription("Users")

	return users
}
