// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// AlertsColumns holds the columns for the "alerts" table.
	AlertsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "scenario", Type: field.TypeString},
		{Name: "bucket_id", Type: field.TypeString},
		{Name: "message", Type: field.TypeString},
		{Name: "events_count", Type: field.TypeInt32},
		{Name: "started_at", Type: field.TypeTime},
		{Name: "stopped_at", Type: field.TypeTime},
		{Name: "source_ip", Type: field.TypeString, Nullable: true},
		{Name: "source_range", Type: field.TypeString, Nullable: true},
		{Name: "source_as_number", Type: field.TypeString, Nullable: true},
		{Name: "source_as_name", Type: field.TypeString, Nullable: true},
		{Name: "source_country", Type: field.TypeString, Nullable: true},
		{Name: "source_latitude", Type: field.TypeFloat32, Nullable: true},
		{Name: "source_longitude", Type: field.TypeFloat32, Nullable: true},
		{Name: "source_scope", Type: field.TypeString},
		{Name: "source_value", Type: field.TypeString},
		{Name: "capacity", Type: field.TypeInt32},
		{Name: "leak_speed", Type: field.TypeString},
		{Name: "machine_alerts", Type: field.TypeInt, Nullable: true},
	}
	// AlertsTable holds the schema information for the "alerts" table.
	AlertsTable = &schema.Table{
		Name:       "alerts",
		Columns:    AlertsColumns,
		PrimaryKey: []*schema.Column{AlertsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "alerts_machines_alerts",
				Columns: []*schema.Column{AlertsColumns[20]},

				RefColumns: []*schema.Column{MachinesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BlockersColumns holds the columns for the "blockers" table.
	BlockersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "api_key", Type: field.TypeString},
		{Name: "revoked", Type: field.TypeBool},
		{Name: "ip_address", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "type", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "until", Type: field.TypeTime, Nullable: true},
		{Name: "last_pull", Type: field.TypeTime},
	}
	// BlockersTable holds the schema information for the "blockers" table.
	BlockersTable = &schema.Table{
		Name:        "blockers",
		Columns:     BlockersColumns,
		PrimaryKey:  []*schema.Column{BlockersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DecisionsColumns holds the columns for the "decisions" table.
	DecisionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "until", Type: field.TypeTime},
		{Name: "scenario", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "start_ip", Type: field.TypeInt64, Nullable: true},
		{Name: "end_ip", Type: field.TypeInt64, Nullable: true},
		{Name: "scope", Type: field.TypeString},
		{Name: "target", Type: field.TypeString},
		{Name: "alert_decisions", Type: field.TypeInt, Nullable: true},
	}
	// DecisionsTable holds the schema information for the "decisions" table.
	DecisionsTable = &schema.Table{
		Name:       "decisions",
		Columns:    DecisionsColumns,
		PrimaryKey: []*schema.Column{DecisionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "decisions_alerts_decisions",
				Columns: []*schema.Column{DecisionsColumns[10]},

				RefColumns: []*schema.Column{AlertsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "time", Type: field.TypeTime},
		{Name: "serialized", Type: field.TypeString},
		{Name: "alert_events", Type: field.TypeInt, Nullable: true},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "events_alerts_events",
				Columns: []*schema.Column{EventsColumns[5]},

				RefColumns: []*schema.Column{AlertsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MachinesColumns holds the columns for the "machines" table.
	MachinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "machine_id", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "ip_address", Type: field.TypeString},
		{Name: "is_validated", Type: field.TypeBool},
		{Name: "status", Type: field.TypeString, Nullable: true},
	}
	// MachinesTable holds the schema information for the "machines" table.
	MachinesTable = &schema.Table{
		Name:        "machines",
		Columns:     MachinesColumns,
		PrimaryKey:  []*schema.Column{MachinesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// MetaColumns holds the columns for the "meta" table.
	MetaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "alert_metas", Type: field.TypeInt, Nullable: true},
	}
	// MetaTable holds the schema information for the "meta" table.
	MetaTable = &schema.Table{
		Name:       "meta",
		Columns:    MetaColumns,
		PrimaryKey: []*schema.Column{MetaColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "meta_alerts_metas",
				Columns: []*schema.Column{MetaColumns[5]},

				RefColumns: []*schema.Column{AlertsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AlertsTable,
		BlockersTable,
		DecisionsTable,
		EventsTable,
		MachinesTable,
		MetaTable,
	}
)

func init() {
	AlertsTable.ForeignKeys[0].RefTable = MachinesTable
	DecisionsTable.ForeignKeys[0].RefTable = AlertsTable
	EventsTable.ForeignKeys[0].RefTable = AlertsTable
	MetaTable.ForeignKeys[0].RefTable = AlertsTable
}
