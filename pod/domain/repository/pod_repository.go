package repository

import (
	"paas-learning/pod/domain/model"

	"github.com/jinzhu/gorm"
)

type PodRepositoryInterface interface {
	InitTable() error

	CreatePod(*model.Pod) (int64, error)

	DeletePodByID(int64) error
	// DeletePodByName(string) error

	UpdatePod(*model.Pod) error

	FindPodByID(int64) (*model.Pod, error)
	FindPodByName(string) (*model.Pod, error)
	FindAll() ([]model.Pod, error)
}

func NewPodRepository(db *gorm.DB) PodRepositoryInterface {
	return &PodRepository{mysqlDb: db}
}

type PodRepository struct {
	mysqlDb *gorm.DB
}

func (p *PodRepository) InitTable() error {
	return p.mysqlDb.CreateTable(&model.Pod{}, &model.PodPort{}, &model.PodEnv{}).Error
}

func (p *PodRepository) CreatePod(pod *model.Pod) (int64, error) {
	return pod.ID, p.mysqlDb.Create(pod).Error
}

func (p *PodRepository) DeletePodByID(podID int64) error {
	tx := p.mysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	if err := p.mysqlDb.Where("pod_id = ?", podID).Delete(&model.Pod{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := p.mysqlDb.Where("pod_id = ?", podID).Delete(&model.PodPort{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := p.mysqlDb.Where("pod_id = ?", podID).Delete(&model.PodEnv{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (p *PodRepository) DeletePodByName(s string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PodRepository) UpdatePod(pod *model.Pod) error {
	return p.mysqlDb.Model(pod).Update(pod).Error
}

// TODO 更新 PodPort PodEnv

func (p *PodRepository) FindPodByID(podID int64) (pod *model.Pod, err error) {
	pod = &model.Pod{}
	return pod, p.mysqlDb.Preload("PodPort").Preload("PodEnv").First(pod, "id = ?", podID).Error
}

func (p *PodRepository) FindPodByName(podName string) (pod *model.Pod, err error) {
	pod = &model.Pod{}
	return pod, p.mysqlDb.Preload("PodPort").Preload("PodEnv").First(pod, "name = ?", podName).Error
}

func (p *PodRepository) FindAll() (pods []model.Pod, err error) {
	return pods, p.mysqlDb.Preload("PodPort").Preload("PodEnv").Find(&pods).Error
}
