package airdrop

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertAirdropFree(t *testing.T) {
	name := "Beacon2"
	task := "GAME2"
	link := "https://nfq.thebeacon.gg/?referral=0xNekowawolf2"
	hasil, err := InsertAirdropFree(name, task, link)
	if err != nil {
		t.Errorf("Failed to insert AirdropFree: %v", err)
	} else {
		fmt.Printf("Inserted AirdropFree ID: %v\n", hasil)
	}
}

func TestInsertAirdropPaid(t *testing.T) {
	name := "Solv Protocol2"
	task := "HOLD"
	link := "https://app.solv.finance/points2"
	hasil, err := InsertAirdropPaid(name, task, link)
	if err != nil {
		t.Errorf("Failed to insert AirdropPaid: %v", err)
	} else {
		fmt.Printf("Inserted AirdropPaid ID: %v\n", hasil)
	}
}

func TestGetAllAirdropFree(t *testing.T) {
	data, err := GetAllAirdropFree()
	if err != nil {
		t.Errorf("Failed to retrieve AirdropFree data: %v", err)
	} else if len(data) == 0 {
		t.Errorf("No AirdropFree data found")
	} else {
		fmt.Printf("Retrieved AirdropFree data: %v\n", data)
	}
}

func TestGetAllAirdropPaid(t *testing.T) {
	data, err := GetAllAirdropPaid()
	if err != nil {
		t.Errorf("Failed to retrieve AirdropPaid data: %v", err)
	} else if len(data) == 0 {
		t.Errorf("No AirdropPaid data found")
	} else {
		fmt.Printf("Retrieved AirdropPaid data: %v\n", data)
	}
}

func TestGetAirdropFreeByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("66cfa14dd39e7e3b0c85b295") 
	if err != nil {
		t.Errorf("Invalid ID format: %v", err)
		return
	}

	airdrop, err := GetAirdropFreeByID(id)
	if err != nil {
		t.Errorf("Failed to retrieve AirdropFree by ID: %v", err)
	} else {
		fmt.Printf("Retrieved AirdropFree by ID: %v\n", airdrop)
	}
}

func TestGetAirdropPaidByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("66cfa14dd39e7e3b0c85b297")
	if err != nil {
		t.Errorf("Invalid ID format: %v", err)
		return
	}

	airdrop, err := GetAirdropPaidByID(id)
	if err != nil {
		t.Errorf("Failed to retrieve AirdropPaid by ID: %v", err)
	} else {
		fmt.Printf("Retrieved AirdropPaid by ID: %v\n", airdrop)
	}
}

func TestUpdateAirdropByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("66dc8cf3b1bc5acfc329268b")
	if err != nil {
		t.Fatalf("Invalid ID format: %v", err)
	}

	err = UpdateAirdropByID("airdrop_paid", id, "git3", "hub3", "https://example.com3")
	if err != nil {
		t.Fatalf("Failed to update Airdrop by ID: %v", err)
	}

	updatedAirdrop, err := GetAirdropFreeByID(id)
	if err != nil {
		t.Fatalf("Failed to retrieve updated AirdropFree by ID: %v", err)
	}

	if updatedAirdrop.Name != "git3" || updatedAirdrop.Task != "hub3" || updatedAirdrop.Link != "https://example.com3" {
		t.Errorf("AirdropFree data not updated correctly: got %v", updatedAirdrop)
	}
}

func TestDeleteAirdropByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("66dc8784c047f376f4c45294")
	if err != nil {
		t.Errorf("Invalid ID format: %v", err)
		return
	}

	err = DeleteAirdropByID("airdrop_paid", id)
	if err != nil {
		t.Errorf("Failed to delete AirdropFree by ID: %v", err)
	} else {
		fmt.Printf("Deleted AirdropFree by ID: %v\n", id.Hex())
	}
}