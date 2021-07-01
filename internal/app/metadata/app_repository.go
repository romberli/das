package metadata

import (
	"errors"
	"fmt"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"

	"github.com/romberli/log"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/metadata"
)

var _ metadata.AppRepo = (*AppRepo)(nil)

type AppRepo struct {
	Database middleware.Pool
}

// NewAppRepo returns *AppRepo with given middleware.Pool
func NewAppRepo(db middleware.Pool) *AppRepo {
	return &AppRepo{db}
}

// NewAppRepoWithGlobal returns *AppRepo with global mysql pool
func NewAppRepoWithGlobal() *AppRepo {
	return NewAppRepo(global.DASMySQLPool)
}

// Execute executes command with arguments on the middleware
func (ar *AppRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := ar.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata AppRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
func (ar *AppRepo) Transaction() (middleware.Transaction, error) {
	return ar.Database.Transaction()
}

// GetAll gets all apps from the middleware
func (ar *AppRepo) GetAll() ([]metadata.App, error) {
	sql := `
		select id, app_name, level, owner_id, del_flag, create_time, last_update_time
		from t_meta_app_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata AppRepo.GetAll() sql: \n%s", sql)

	result, err := ar.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*AppInfo
	appInfoList := make([]*AppInfo, result.RowNumber())
	for i := range appInfoList {
		appInfoList[i] = NewEmptyAppInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(appInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []dependency.Entity
	appList := make([]metadata.App, result.RowNumber())
	for i := range appList {
		appList[i] = appInfoList[i]
	}

	return appList, nil
}

// GetByID gets an app by the identity from the middleware
func (ar *AppRepo) GetByID(id int) (metadata.App, error) {
	sql := `
		select id, app_name, level,owner_id, del_flag, create_time, last_update_time
		from t_meta_app_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata AppRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := ar.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata AppInfo.GetByID(): data does not exists, id: %d", id))
	case 1:
		appInfo := NewEmptyAppInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(appInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return appInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata AppInfo.GetByID(): duplicate key exists, id: %d", id))
	}
}

// GetID gets the identity with given app name from the middleware
func (ar *AppRepo) GetID(systemName string) (int, error) {
	sql := `select id from t_meta_app_info where del_flag = 0 and app_name = ?;`
	log.Debugf("metadata AppRepo.GetID() select sql: %s", sql)
	result, err := ar.Execute(sql, systemName)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// GetAppSystemByName gets the app by name from the middleware
func (ar *AppRepo) GetAppByName(appName string) (metadata.App, error) {
	sql := `select id from t_meta_app_info where del_flag = 0 and app_name = ?;`
	log.Debugf("metadata AppRepo.GetAppByName() select sql: %s", sql)
	result, err := ar.Execute(sql, appName)
	if err != nil {
		return nil, err
	}

	id, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	if err != nil {
		return nil, err
	}

	return ar.GetByID(id)
}

// GetDBIDList gets a database identity list that app uses
func (ar *AppRepo) GetDBIDList(id int) ([]int, error) {
	sql := `
		select db_id
		from t_meta_app_info ai
				 inner join t_meta_app_db_map adm on ai.id = adm.app_id
		where ai.del_flag = 0
		  and adm.del_flag = 0
		  and ai.id = ?;
	`
	log.Debugf("metadata AppRepo.GetDBIDList() select sql: %s", sql)
	result, err := ar.Execute(sql, id)
	if err != nil {
		return nil, err
	}

	resultNum := result.RowNumber()
	dbIDList := make([]int, resultNum)

	for row := 0; row < resultNum; row++ {
		dbID, err := result.GetInt(row, constant.ZeroInt)
		if err != nil {
			return nil, err
		}

		dbIDList[row] = dbID
	}

	return dbIDList, nil

}

// Create creates an app in the middleware
func (ar *AppRepo) Create(app metadata.App) (metadata.App, error) {
	sql := `insert into t_meta_app_info(app_name, level, owner_id) values(?, ?, ?);`
	log.Debugf("metadata AppRepo.Create() insert sql: %s", sql)
	// execute
	_, err := ar.Execute(sql, app.GetAppName(), app.GetLevel(), app.GetOwnerID())
	if err != nil {
		return nil, err
	}
	// get id
	id, err := ar.GetID(app.GetAppName())
	if err != nil {
		return nil, err
	}
	// get entity
	return ar.GetByID(id)
}

// Update updates the app in the middleware
func (ar *AppRepo) Update(app metadata.App) error {
	sql := `update t_meta_app_info set app_name = ?, level = ?, owner_id = ?, del_flag = ? where id = ?;`
	log.Debugf("metadata AppRepo.Update() update sql: %s", sql)
	_, err := ar.Execute(sql, app.GetAppName(), app.GetLevel(), app.GetOwnerID(), app.GetDelFlag(), app.Identity())

	return err
}

// Delete deletes the app in the middleware
func (ar *AppRepo) Delete(id int) error {
	tx, err := ar.Transaction()
	if err != nil {
		return err
	}
	defer func() {
		err = tx.Close()
		if err != nil {
			log.Errorf("metadata AppRepo.Delete(): close database connection failed.\n%s", err.Error())
		}
	}()

	err = tx.Begin()
	if err != nil {
		return err
	}
	sql := `delete from t_meta_app_info where id = ?;`
	log.Debugf("metadata AppRepo.Delete() delete sql(t_meta_app_info): %s", sql)
	_, err = tx.Execute(sql, id)
	if err != nil {
		return err
	}
	sql = `delete from t_meta_app_db_map where app_id = ?;`
	log.Debugf("metadata AppRepo.Delete() delete sql(t_meta_app_db_map): %s", sql)
	_, err = tx.Execute(sql, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// AddDB adds a new map of app and database in the middleware
func (ar *AppRepo) AddDB(appID, dbID int) error {
	sql := `insert into t_meta_app_db_map(app_id, db_id) values(?, ?);`
	log.Debugf("metadata AppRepo.AddDB() insert sql: %s", sql)
	_, err := ar.Execute(sql, appID, dbID)

	return err
}

// DeleteDB delete the map of app and database in the middleware
func (ar *AppRepo) DeleteDB(appID, dbID int) error {
	sql := `delete from t_meta_app_db_map where app_id = ? and db_id = ?;`
	log.Debugf("metadata AppRepo.DeleteDB() delete sql: %s", sql)
	_, err := ar.Execute(sql, appID, dbID)

	return err
}
