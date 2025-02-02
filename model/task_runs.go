package model

import (
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/utils"
)

// GetRunsForTask is just a join of all runs that belong to a task.
func (m *Model) GetRunsForTask(id, page, perPage int64) ([]*Run, *errors.Error) {
	runs := []*Run{}

	page, perPage, err := utils.ScopePaginationInt(page, perPage)
	if err != nil {
		return nil, err
	}

	if err := m.Order("id DESC").Limit(perPage).Offset(page*perPage).Where("task_id = ?", id).Find(&runs).Error; err != nil {
		return nil, errors.New(err)
	}

	return runs, nil
}

// CountRunsForTask retrieves the total count of runs for the given task.
func (m *Model) CountRunsForTask(id int64) (int64, *errors.Error) {
	var count int64

	if err := m.Model(&Run{}).Where("task_id = ?", id).Count(&count).Error; err != nil {
		return 0, errors.New(err)
	}

	return count, nil
}
