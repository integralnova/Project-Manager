package permit_tracker

import (
	"fmt"

	"github.com/integralnova/Project-Manager/models"
)

// working slice of permits. To be used with search and filter
type OpenPermits map[string]models.PermitsModel


// OpenPermit adds a permit to the OpenPermits. When you arealdy have models.permitsModel
// use OpenPermitfromDB to get it from the database
// Idk if I need these two TODO
func (op *OpenPermits) OpenPermitFromMemory(permit *models.PermitsModel) error {
	if permit == nil {
		return fmt.Errorf("permit cannot be nil")
	}
	(*op)[permit.PermitID] = *permit
	return nil
}


// OpenPermitFromDB gets a permit from the database
func (op *OpenPermits) OpenPermitFromDB(m *models.Datatings, permitID string) error {
	permit, err := m.GetPermitByName(permitID)
	if err != nil {
		return fmt.Errorf("permit not found")
	}
	(*op)[permitID] = permit
	return nil

}

func (op *OpenPermits) OpenPermitByID(permitID string) (models.PermitsModel, error) {
	if permit, ok := (*op)[permitID]; ok {
		return permit, nil
	}
	return models.PermitsModel{}, fmt.Errorf("permit not found")
}

func (op *OpenPermits) RemoveOpenPermit(permitID string) error {

	if _, ok := (*op)[permitID]; !ok {
		delete(*op, permitID)
		return nil
	}

	return fmt.Errorf("permit not in open permits")

}
