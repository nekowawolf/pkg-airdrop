package airdrop

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertAirdropFree(t *testing.T) {
	name := "initia2"
	task := "TESTNET"
	link := "https://initiatest.com"
	level := "hard"
	hasil, err := InsertAirdropFree(name, task, link, level)
	if err != nil {
		t.Errorf("Failed to insert AirdropFree: %v", err)
	} else {
		fmt.Printf("Inserted AirdropFree ID: %v\n", hasil)
	}
}

func TestInsertAirdropPaid(t *testing.T) {
	name := "test"
	task := "HOLD"
	link := "https://test.com"
	level := "medium"
	hasil, err := InsertAirdropPaid(name, task, link, level)
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

func TestGetAllAirdrop(t *testing.T) {
	allAirdrops, err := GetAllAirdrop()
	if err != nil {
		t.Errorf("Failed to retrieve all airdrops: %v", err)
		return
	}

	if len(allAirdrops) == 0 {
		t.Errorf("No airdrops found.")
		return
	}

	fmt.Printf("Retrieved %d airdrops:\n", len(allAirdrops))
	for _, airdrop := range allAirdrops {
		fmt.Printf("%v\n", airdrop)
	}

	freeAirdrops, err := GetAllAirdropFree()
	if err != nil {
		t.Errorf("Failed to retrieve AirdropFree data: %v", err)
		return
	}

	paidAirdrops, err := GetAllAirdropPaid()
	if err != nil {
		t.Errorf("Failed to retrieve AirdropPaid data: %v", err)
		return
	}

	expectedCount := len(freeAirdrops) + len(paidAirdrops)
	if len(allAirdrops) != expectedCount {
		t.Errorf("Expected %d airdrops, but got %d", expectedCount, len(allAirdrops))
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

func TestGetAirdropFreeByName(t *testing.T) {
	name := "Initia"
	airdrop, err := GetAirdropFreeByName(name)
	if err != nil {
		t.Fatalf("Error calling GetAirdropFreeByName: %v", err)
	}

	fmt.Println("Free Airdrops found:", airdrop)
}

func TestGetAirdropPaidByName(t *testing.T) {
	name := "Dflow"
	airdrop, err := GetAirdropPaidByName(name)
	if err != nil {
		t.Fatalf("Error calling GetAirdropPaidByName: %v", err)
	}

	fmt.Println("Paid Airdrops found:", airdrop)
}

func TestUpdateAirdropFreeByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("66ea30e2017690fa6e447744")
	if err != nil {
		t.Errorf("Invalid ID format: %v", err)
		return
	}

	newName := "Initia 3"
	newTask := "TESTNET"
	newLink := "https://initiatest.com/updated"
	newLevel := "Easy"

	err = UpdateAirdropFreeByID(id, newName, newTask, newLink, newLevel)
	if err != nil {
		t.Errorf("Failed to update AirdropFree by ID: %v", err)
		return
	}

	airdrop, err := GetAirdropFreeByID(id)
	if err != nil {
		t.Errorf("Failed to retrieve AirdropFree by ID after update: %v", err)
		return
	}

	if airdrop.Name != newName || airdrop.Task != newTask || airdrop.Link != newLink || airdrop.Level != newLevel {
		t.Errorf("AirdropFree not updated correctly. Got: %+v", airdrop)
	} else {
		fmt.Printf("AirdropFree updated successfully: %+v\n", airdrop)
	}
}

func TestUpdateAirdropPaidByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("66ea2fbcf75fbbced4145728")
	if err != nil {
		t.Errorf("Invalid ID format: %v", err)
		return
	}

	newName := "Dflow 3"
	newTask := "Retro"
	newLink := "https://app.test/updated"
	newLevel := "Medium"

	err = UpdateAirdropPaidByID(id, newName, newTask, newLink, newLevel)
	if err != nil {
		t.Errorf("Failed to update AirdropPaid by ID: %v", err)
		return
	}

	airdrop, err := GetAirdropPaidByID(id)
	if err != nil {
		t.Errorf("Failed to retrieve AirdropPaid by ID after update: %v", err)
		return
	}

	if airdrop.Name != newName || airdrop.Task != newTask || airdrop.Link != newLink || airdrop.Level != newLevel {
		t.Errorf("AirdropPaid not updated correctly. Got: %+v", airdrop)
	} else {
		fmt.Printf("AirdropPaid updated successfully: %+v\n", airdrop)
	}
}

func TestDeleteAirdropFreeByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("66ea30e2017690fa6e447744")
	if err != nil {
		t.Errorf("Invalid ID format: %v", err)
		return
	}

	err = DeleteAirdropFreeByID(id)
	if err != nil {
		t.Errorf("Failed to delete AirdropFree by ID: %v", err)
		return
	}

	airdrop, err := GetAirdropFreeByID(id)
	if err == nil {
		t.Errorf("Expected no document, but found AirdropFree: %+v", airdrop)
	} else {
		fmt.Printf("AirdropFree deleted successfully, no document found with ID: %s\n", id.Hex())
	}
}

func TestDeleteAirdropPaidByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("66ea2fbcf75fbbced4145728")
	if err != nil {
		t.Errorf("Invalid ID format: %v", err)
		return
	}

	err = DeleteAirdropPaidByID(id)
	if err != nil {
		t.Errorf("Failed to delete AirdropPaid by ID: %v", err)
		return
	}

	airdrop, err := GetAirdropPaidByID(id)
	if err == nil {
		t.Errorf("Expected no document, but found AirdropPaid: %+v", airdrop)
	} else {
		fmt.Printf("AirdropPaid deleted successfully, no document found with ID: %s\n", id.Hex())
	}
}