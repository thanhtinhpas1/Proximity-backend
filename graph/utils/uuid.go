package utils

import "github.com/google/uuid"

func ToUuids(uuidArr []string) ([]uuid.UUID, error) {
	uuids := make([]uuid.UUID, len(uuidArr))
	for i, v := range uuidArr {
		uuid, err := uuid.Parse(v)
		if err != nil {
			return nil, err
		}
		uuids[i] = uuid
	}
	return uuids, nil
}
