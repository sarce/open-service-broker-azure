package postgresqldb

import "github.com/Azure/open-service-broker-azure/pkg/service"

func (m *module) GetCatalog() (service.Catalog, error) {
	return service.NewCatalog([]service.Service{
		//all-in-one
		service.NewService(
			&service.ServiceProperties{
				ID:          "b43b4bba-5741-4d98-a10b-17dc5cee0175",
				Name:        "azure-postgresqldb",
				Description: "Azure Database for PostgreSQL (Experimental)",
				Metadata: &service.ServiceMetadata{
					DisplayName: "Azure Database for PostgreSQL",
					ImageUrl: "https://azure.microsoft.com/svghandler/postgresql/" +
						"?width=200",
					LongDescription: "Enterprise-ready fully managed community PostgreSQL" +
						" (Experimental)",
					DocumentationUrl: "https://docs.microsoft.com/en-us/azure/postgresql/",
					SupportUrl:       "https://azure.microsoft.com/en-us/support/",
				},
				Bindable: true,
				Tags:     []string{"Azure", "PostgreSQL", "Database"},
			},
			m.allInOneManager,
			service.NewPlan(&service.PlanProperties{
				ID:          "b2ed210f-6a10-4593-a6c4-964e6b6fad62",
				Name:        "basic50",
				Description: "Basic Tier, 50 DTUs",
				Free:        false,
				Extended: map[string]interface{}{
					"skuName":        "PGSQLB50",
					"skuTier":        "Basic",
					"skuCapacityDTU": 50,
				},
				Metadata: &service.ServicePlanMetadata{
					DisplayName: "Basic Tier",
					Bullets:     []string{"50 DTUs"},
				},
			}),
			service.NewPlan(&service.PlanProperties{
				ID:          "843d7d03-9306-447e-8c19-25ccc4ac30d7",
				Name:        "basic100",
				Description: "Basic Tier, 100 DTUs",
				Free:        false,
				Extended: map[string]interface{}{
					"skuName":        "PGSQLB100",
					"skuTier":        "Basic",
					"skuCapacityDTU": 100,
				},
				Metadata: &service.ServicePlanMetadata{
					DisplayName: "Basic Tier",
					Bullets:     []string{"100 DTUs"},
				},
			}),
		),
		//dbms only
		service.NewService(
			&service.ServiceProperties{
				ID:          "d3f74b44-79bc-4d1e-bf7d-c247c2b851f9",
				Name:        "azure-postgresqldb-dbms-only",
				Description: "Azure Database for PostgreSQL DBMS Only (Experimental)",
				Metadata: &service.ServiceMetadata{
					DisplayName: "Azure Database for PostgreSQL (DBMS Only)",
					ImageUrl: "https://azure.microsoft.com/svghandler/postgresql/" +
						"?width=200",
					LongDescription: "Enterprise-ready fully managed community PostgreSQL" +
						" (Experimental)",
					DocumentationUrl: "https://docs.microsoft.com/en-us/azure/postgresql/",
					SupportUrl:       "https://azure.microsoft.com/en-us/support/",
				},
				Bindable:       false,
				Tags:           []string{"Azure", "PostgreSQL", "Database"},
				ChildServiceID: "25434f16-d762-41c7-bbdd-8045d7f74ca",
			},
			m.dbmsOnlyManager,
			service.NewPlan(&service.PlanProperties{
				ID:          "bf389028-8dcc-433a-ab6f-0ee9b8db142f",
				Name:        "basic50",
				Description: "Basic Tier, 50 DTUs",
				Free:        false,
				Extended: map[string]interface{}{
					"skuName":        "PGSQLB50",
					"skuTier":        "Basic",
					"skuCapacityDTU": 50,
				},
				Metadata: &service.ServicePlanMetadata{
					DisplayName: "Basic Tier",
					Bullets:     []string{"50 DTUs"},
				},
			}),
			service.NewPlan(&service.PlanProperties{
				ID:          "58633c61-942c-42cb-b22c-346a4c594b8e",
				Name:        "basic100",
				Description: "Basic Tier, 100 DTUs",
				Free:        false,
				Extended: map[string]interface{}{
					"skuName":        "PGSQLB100",
					"skuTier":        "Basic",
					"skuCapacityDTU": 100,
				},
				Metadata: &service.ServicePlanMetadata{
					DisplayName: "Basic Tier",
					Bullets:     []string{"100 DTUs"},
				},
			}),
		),
		//db only
		service.NewService(
			&service.ServiceProperties{
				ID:          "25434f16-d762-41c7-bbdd-8045d7f74ca6",
				Name:        "azure-postgresql-db-only",
				Description: "Azure Database for PostgreSQL Database Only (Experimental)",
				Metadata: &service.ServiceMetadata{
					DisplayName: "Azure Database for PostgreSQL (Database Only)",
					ImageUrl: "https://azure.microsoft.com/svghandler/postgresql/" +
						"?width=200",
					LongDescription: "Enterprise-ready fully managed community PostgreSQL" +
						" (Experimental)",
					DocumentationUrl: "https://docs.microsoft.com/en-us/azure/postgresql/",
					SupportUrl:       "https://azure.microsoft.com/en-us/support/",
				},
				Bindable:        true,
				Tags:            []string{"Azure", "PostgreSQL", "Database"},
				ParentServiceID: "d3f74b44-79bc-4d1e-bf7d-c247c2b851f9",
			},
			m.dbOnlyManager,
			service.NewPlan(&service.PlanProperties{
				ID:          "df6f5ef1-e602-406b-ba73-09c107d1e31b",
				Name:        "postgressql-db-only",
				Description: "Azure Databse for PostgreSQL (db Only)",
				Free:        false,
				Metadata: &service.ServicePlanMetadata{
					DisplayName: "Azure Databse for PostgreSQL (DB Only)",
				},
			}),
		),
	}), nil
}
