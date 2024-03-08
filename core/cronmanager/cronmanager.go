package cronmanager

import (
	"sync"

	"github.com/robfig/cron/v3"
	"iceline-hosting.com/core/logger"
)

type CronID cron.EntryID

type CronManager struct {
	logger     logger.Logger
	cron       *cron.Cron
	jobFuncMap map[string]cron.EntryID
	enabledMap map[string]bool
	mutex      sync.Mutex
}

func NewCronManager(logger logger.Logger) (*CronManager, error) {
	return &CronManager{
		cron:       cron.New(cron.WithSeconds()),
		logger:     logger,
		jobFuncMap: make(map[string]cron.EntryID),
		enabledMap: make(map[string]bool),
	}, nil
}

func (cm *CronManager) Start() {
	cm.cron.Start()
}

func (cm *CronManager) Stop() {
	cm.cron.Stop()
}

func (cm *CronManager) IsSchedulerIDExists(schedulerID string) bool {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	_, exists := cm.jobFuncMap[schedulerID]

	return exists

}

func (cm *CronManager) DeleteSchedulerID(schedulerID string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	if entryID, exists := cm.jobFuncMap[schedulerID]; exists {
		cm.cron.Remove(entryID)
		delete(cm.jobFuncMap, schedulerID)
		delete(cm.enabledMap, schedulerID)
	}

}

func (cm *CronManager) UpdateEnableSchedulerID(schedulerID string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	cm.enabledMap[schedulerID] = true
}

func (cm *CronManager) UpdateDisableSchedulerID(schedulerID string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	cm.enabledMap[schedulerID] = false
}

func (cm *CronManager) ReturnCronIdOnError() CronID {
	return CronID(0)
}

func (cm *CronManager) Schedule(expression string, jobFunc func()) (CronID, error) {
	entryID, err := cm.cron.AddFunc(expression, jobFunc)

	return CronID(entryID), err
}

func (cm *CronManager) AddSchedulerEntryID(schedule_id string, entryID cron.EntryID) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	cm.jobFuncMap[schedule_id] = entryID

}
