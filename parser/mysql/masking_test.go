package mysql

import (
	"mysql-parser/masker"
	"mysql-parser/parser/base"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMySQLExtractSensitiveField(t *testing.T) {
	const (
		defaultDatabase = "db"
	)
	var (
		defaultDatabaseSchema = &base.SensitiveSchemaInfo{
			DatabaseList: []base.DatabaseSchema{
				{
					Name: defaultDatabase,
					SchemaList: []base.SchemaSchema{
						{
							Name: "",
							TableList: []base.TableSchema{
								{
									Name: "t",
									ColumnList: []base.ColumnInfo{
										{
											Name:              "a",
											MaskingAttributes: base.NewMaskingAttributes(masker.NewDefaultFullMasker()),
										},
										{
											Name:              "b",
											MaskingAttributes: base.NewMaskingAttributes(masker.NewNoneMasker()),
										},
										{
											Name:              "c",
											MaskingAttributes: base.NewMaskingAttributes(masker.NewNoneMasker()),
										},
										{
											Name:              "d",
											MaskingAttributes: base.NewMaskingAttributes(masker.NewDefaultRangeMasker()),
										},
									},
								},
							},
						},
					},
				},
				{
					Name: defaultDatabase,
					SchemaList: []base.SchemaSchema{
						{
							Name: "",
							TableList: []base.TableSchema{
								{
									Name: "e_common_config",
									ColumnList: []base.ColumnInfo{
										{
											Name:              "id",
											MaskingAttributes: base.NewMaskingAttributes(masker.NewNoneMasker()),
										},
										{
											Name:              "config_type",
											MaskingAttributes: base.NewMaskingAttributes(masker.NewDefaultFullMasker()),
										},
										{
											Name:              "config_value",
											MaskingAttributes: base.NewMaskingAttributes(masker.NewNoneMasker()),
										},
										{
											Name:              "create_by",
											MaskingAttributes: base.NewMaskingAttributes(masker.NewNoneMasker()),
										},
									},
								},
							},
						},
					},
				},
			},
		}
	)
	tests := []struct {
		statement  string
		schemaInfo *base.SensitiveSchemaInfo
		fieldList  []base.SensitiveField
	}{
		{
			// Test for case-insensitive column names.
			statement:  `SELECT * FROM (select * from (select a from t) t1 join t as t2 using(A)) result LIMIT 10000;`,
			schemaInfo: defaultDatabaseSchema,
			fieldList: []base.SensitiveField{
				{
					Name:              "a",
					MaskingAttributes: base.NewMaskingAttributes(masker.NewDefaultFullMasker()),
				},
				{
					Name:              "b",
					MaskingAttributes: base.NewMaskingAttributes(masker.NewNoneMasker()),
				},
				{
					Name:              "c",
					MaskingAttributes: base.NewMaskingAttributes(masker.NewNoneMasker()),
				},
				{
					Name:              "d",
					MaskingAttributes: base.NewMaskingAttributes(masker.NewDefaultRangeMasker()),
				},
			},
		},
		{
			// Test for case-insensitive column names.
			statement: `SELECT
				  un,
				  CONCAT(un, 'abc') AS c,
				  CONCAT(CONCAT(un, 'abc'), 'def') AS final_result
				FROM
				  (
					SELECT
					  config_type AS un,
					  CONCAT(config_type, 'abc') AS c
					FROM
					  e_common_config
				  ) AS subquery;`,
			schemaInfo: defaultDatabaseSchema,
			fieldList: []base.SensitiveField{
				{
					Name:              "un",
					MaskingAttributes: base.NewMaskingAttributes(masker.NewDefaultFullMasker()),
				},
				{
					Name:              "c",
					MaskingAttributes: base.NewMaskingAttributes(masker.NewDefaultFullMasker()),
				},
				{
					Name:              "final_result",
					MaskingAttributes: base.NewMaskingAttributes(masker.NewDefaultFullMasker()),
				},
			},
		},
	}

	for _, test := range tests {
		res, err := GetMaskedFields(test.statement, defaultDatabase, test.schemaInfo)
		require.NoError(t, err, test.statement)
		require.Equal(t, test.fieldList, res, test.statement)
	}
}
