package main

import (
    "testing"
)

func TestNewParcel(t *testing.T) {
    parcel := NewParcel("test-1", 10.5)
    
    if parcel.ID != "test-1" {
        t.Errorf("Expected ID 'test-1', got '%s'", parcel.ID)
    }
    
    if parcel.Weight != 10.5 {
        t.Errorf("Expected weight 10.5, got %f", parcel.Weight)
    }
    
    if parcel.Status != "created" {
        t.Errorf("Expected status 'created', got '%s'", parcel.Status)
    }
}

func TestSaveAndGetParcel(t *testing.T) {
    parcel := NewParcel("test-2", 15.0)
    
    // Save parcel
    err := SaveParcel(parcel)
    if err != nil {
        t.Fatalf("Failed to save parcel: %v", err)
    }
    
    // Get parcel
    loaded, err := GetParcel("test-2")
    if err != nil {
        t.Fatalf("Failed to get parcel: %v", err)
    }
    
    if loaded.Weight != 15.0 {
        t.Errorf("Expected weight 15.0, got %f", loaded.Weight)
    }
}

func TestUpdateParcel(t *testing.T) {
    parcel := NewParcel("test-3", 20.0)
    
    err := SaveParcel(parcel)
    if err != nil {
        t.Fatalf("Failed to save parcel: %v", err)
    }
    
    parcel.Weight = 25.0
    parcel.Status = "updated"
    
    err = UpdateParcel(parcel)
    if err != nil {
        t.Fatalf("Failed to update parcel: %v", err)
    }
    
    loaded, err := GetParcel("test-3")
    if err != nil {
        t.Fatalf("Failed to get parcel: %v", err)
    }
    
    if loaded.Weight != 25.0 {
        t.Errorf("Expected weight 25.0, got %f", loaded.Weight)
    }
    
    if loaded.Status != "updated" {
        t.Errorf("Expected status 'updated', got '%s'", loaded.Status)
    }
}

func TestDeleteParcel(t *testing.T) {
    parcel := NewParcel("test-4", 30.0)
    
    err := SaveParcel(parcel)
    if err != nil {
        t.Fatalf("Failed to save parcel: %v", err)
    }
    
    err = DeleteParcel("test-4")
    if err != nil {
        t.Fatalf("Failed to delete parcel: %v", err)
    }
    
    _, err = GetParcel("test-4")
    if err == nil {
        t.Error("Expected error when getting deleted parcel")
    }
}
