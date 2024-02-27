package pathpermissions

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"

	// _ "github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) upsertManyPathPermissions(items models.PathPermissions) (err error) {
	_, err = r.db().NewInsert().On("CONFLICT (resource) DO UPDATE").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("guest_allow = EXCLUDED.guest_allow").
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteManyPathPermissions(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.PathPermission)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertManyPathPermissionsRoles(items models.PathPermissionsRoles) (err error) {
	_, err = r.db().NewInsert().On("CONFLICT (path_permission_id, role_id) DO UPDATE").
		Set("path_permission_id = EXCLUDED.path_permission_id").
		Set("role_id = EXCLUDED.role_id").
		Model(&items).
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteManyPathPermissionsRoles(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.PathPermissionRole)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) getAllPathPermissions() (models.PathPermissions, error) {
	items := make(models.PathPermissions, 0)
	err := r.db().NewSelect().
		Model(&items).
		Relation("PathPermissionsRoles").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) getAllPathPermissionsAdmin() (item models.PathPermissionsWithCount, err error) {
	item.PathPermissions = make(models.PathPermissions, 0)
	query := r.db().NewSelect().Model(&item.PathPermissions).
		Relation("PathPermissionsRoles")

	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) checkPathPermissions(path string, roleID string) error {
	if roleID == "" {
		return r.db().NewSelect().
			Model(&models.PathPermission{}).
			Where("path_permissions.resource = ? and path_permissions.guest_allow = true", path).
			Scan(r.ctx)
	}

	return r.db().NewSelect().
		Model(&models.PathPermission{}).
		Join("JOIN path_permissions_roles ppr on ppr.path_permission_id = path_permissions.id and ppr.role_id = ?", roleID).
		Where("path_permissions.resource = ?", path).
		Scan(r.ctx)
}

func (r *Repository) getPathPermissionsByRoleID(id string) (models.PathPermissions, error) {
	items := make(models.PathPermissions, 0)
	err := r.db().NewSelect().
		Model(&items).
		Relation("PathPermissionsRoles").
		Join("JOIN path_permissions_roles ppr on ppr.path_permission_id = path_permissions.id and ppr.role_id = ?", id).
		// Where("path_permissions.path_permissions_roles.role_id = ?", id).
		Scan(r.ctx)
	return items, err
}
