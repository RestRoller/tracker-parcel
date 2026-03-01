package main

import (
    "errors"
    "sync"
)

type Parcel struct {
    ID     string
    Weight float64
    Status string
}

var (
    parcels = make(map[string]*Parcel)
    mutex   = &sync.RWMutex{}
)

func NewParcel(id string, weight float64) *Parcel {
    return &Parcel{
        ID:     id,
        Weight: weight,
        Status: "created",
    }
}

func SaveParcel(parcel *Parcel) error {
    mutex.Lock()
    defer mutex.Unlock()
    
    if parcel.ID == "" {
        return errors.New("parcel ID cannot be empty")
    }
    
    parcels[parcel.ID] = parcel
    return nil
}

func GetParcel(id string) (*Parcel, error) {
    mutex.RLock()
    defer mutex.RUnlock()
    
    parcel, exists := parcels[id]
    if !exists {
        return nil, errors.New("parcel not found")
    }
    return parcel, nil
}

func UpdateParcel(parcel *Parcel) error {
    mutex.Lock()
    defer mutex.Unlock()
    
    if _, exists := parcels[parcel.ID]; !exists {
        return errors.New("parcel not found")
    }
    
    parcels[parcel.ID] = parcel
    return nil
}

func DeleteParcel(id string) error {
    mutex.Lock()
    defer mutex.Unlock()
    
    if _, exists := parcels[id]; !exists {
        return errors.New("parcel not found")
    }
    
    delete(parcels, id)
    return nil
}
