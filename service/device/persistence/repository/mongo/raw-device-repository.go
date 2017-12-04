package mongo

import (
	"github.com/alexx2012/go-home/service/device/persistence/model"
	"github.com/alexx2012/go-home/service/device/persistence/repository"
	"gopkg.in/mgo.v2"
)

const RawDevicesCollectionName = "raw_devices"

type RawDeviceRepository struct {
	db *mgo.Database
}

func NewRawDeviceRepository(DB *mgo.Database) *RawDeviceRepository {
	return &RawDeviceRepository{db: DB}
}

func (r *RawDeviceRepository) Save(device *model.RawDevice) (err error) {
	_, err = r.db.C(RawDevicesCollectionName).UpsertId(device.Id, device)
	return
}

func (r *RawDeviceRepository) FindAll() (out []*model.RawDevice, err error) {
	err = r.db.C(RawDevicesCollectionName).Find(nil).All(&out)
	return
}

func (r *RawDeviceRepository) FindById(id uint16) (*model.RawDevice, error) {
	return r.prepareFindResult(r.db.C(RawDevicesCollectionName).FindId(id))
}

func (r *RawDeviceRepository) RemoveById(id uint16) error {
	return r.db.C(RawDevicesCollectionName).RemoveId(id)
}

func (r *RawDeviceRepository) prepareFindResult(query *mgo.Query) (*model.RawDevice, error) {
	var (
		device     = &model.RawDevice{}
		count, err = query.Count()
	)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, repository.ErrRawDeviceNotFound
	}

	if err = query.One(device); err != nil {
		return nil, err
	}

	return device, nil
}
