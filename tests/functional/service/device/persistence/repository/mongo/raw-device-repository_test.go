package mongo_tests

import (
	"testing"

	"fmt"
	"os"

	"github.com/alexx2012/go-home/service/device/persistence/model"
	"github.com/alexx2012/go-home/service/device/persistence/repository/mongo"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

func TestFindById(t *testing.T) {
	var (
		repository  = getRepository(t)
		actual, err = repository.FindById(2)
		expected    = &model.RawDevice{
			Id:             2,
			ParentDeviceId: 3,
			SubDeviceId:    4,
			SubDeviceType:  5,
		}
	)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestFindAll(t *testing.T) {
	var (
		repository  = getRepository(t)
		actual, err = repository.FindAll()
		expected    = []*model.RawDevice{
			{
				Id:             1,
				ParentDeviceId: 2,
				SubDeviceId:    3,
				SubDeviceType:  4,
			},
			{
				Id:             2,
				ParentDeviceId: 3,
				SubDeviceId:    4,
				SubDeviceType:  5,
			},
		}
	)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func getRepository(t *testing.T) *mongo.RawDeviceRepository {
	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	assert.Nil(t, err)

	mongoDB := session.DB(fmt.Sprintf("%s_test", mongo.RawDevicesCollectionName))

	repo := mongo.NewRawDeviceRepository(mongoDB)

	repo.Save(&model.RawDevice{
		Id:             1,
		ParentDeviceId: 2,
		SubDeviceId:    3,
		SubDeviceType:  4,
	})

	repo.Save(&model.RawDevice{
		Id:             2,
		ParentDeviceId: 3,
		SubDeviceId:    4,
		SubDeviceType:  5,
	})

	return repo
}
