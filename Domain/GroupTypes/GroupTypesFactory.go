package DomainGroupTypes

import (
	"database/sql"

	DomainGroupTypesEntity "gontact/Domain/GroupTypes/Entity"
)

func getGroupTypesCollection(rows *sql.Rows) []DomainGroupTypesEntity.GroupTypesEntity {
	var collection []DomainGroupTypesEntity.GroupTypesEntity
	defer rows.Close()

	for rows.Next() {
		var entity DomainGroupTypesEntity.GroupTypesEntity
		rows.Scan(&entity.Id, &entity.Name, &entity.CreatedBy)
		collection = append(collection, entity)
	}

	return collection
}
