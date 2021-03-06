package checkers

import (
	"github.com/google/uuid"
)

type IdChecker interface {
	IsEmpty() bool
	IsValid() bool
}

type UuidString string

type UuidGeneric uuid.UUID

func (inst UuidString) IsEmpty() bool {
	return len(inst) == 0
}

func (inst UuidString) IsValid() bool {
	if _, err := uuid.Parse(string(inst)); err != nil {
		return false
	}
	return true
}

func (inst UuidGeneric) IsEmpty() bool {
	id := uuid.UUID(inst)
	return id[0] == 0
}

func (inst UuidGeneric) IsValid() bool {
	id := uuid.UUID(inst)
	idStr := id.String()
	if parsed, e := uuid.Parse(idStr); parsed[0] == 0 || e != nil {
		return false
	}
	return true
}
