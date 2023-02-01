package DomainGroupTypes

import (
	DomainGroupTypesEntity "gontact/Domain/GroupTypes/Entity"
	DomainGroupTypesVO "gontact/Domain/GroupTypes/VO"
	InterfaceGroupTypes "gontact/Interface/GroupTypes"
)

func GetGroupTypes(vo DomainGroupTypesVO.GroupTypesVO) []DomainGroupTypesEntity.GroupTypesEntity {
	return getGroupTypesCollection(InterfaceGroupTypes.GetDbGroupTypes(vo))
}
