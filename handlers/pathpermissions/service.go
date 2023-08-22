package pathpermissions

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) UpsertManyPathPermissions(paths models.PathPermissions) error {
	if len(paths) == 0 {
		return nil
	}
	err := s.repository.upsertManyPathPermissions(paths)
	if err != nil {
		return err
	}

	if len(paths.GetPathPermissionsRolesForDelete()) > 0 {
		err = s.repository.deleteManyPathPermissionsRoles(paths.GetPathPermissionsRolesForDelete())
		if err != nil {
			return err
		}
	}
	if len(paths.GetPathPermissionsRoles()) > 0 {
		err = s.repository.upsertManyPathPermissionsRoles(paths.GetPathPermissionsRoles())
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) GetAllPathPermissions() (models.PathPermissions, error) {
	return s.repository.getAllPathPermissions()
}

func (s *Service) GetAllPathPermissionsAdmin() (models.PathPermissionsWithCount, error) {
	return s.repository.getAllPathPermissionsAdmin()
}

func (s *Service) CheckPathPermissions(path string, roleID string) error {
	return s.repository.checkPathPermissions(path, roleID)
}

func (s *Service) GetPathPermissionsByRoleID(id string) (models.PathPermissions, error) {
	return s.repository.getPathPermissionsByRoleID(id)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
