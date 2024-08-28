package airdrop

import (
	"fmt"
	"testing"
)

func TestInsertAirdropFree(t *testing.T) {
	name := "Beacon"
	task := "GAME"
	link := "https://nfq.thebeacon.gg/?referral=0xNekowawolf"
	hasil, err := InsertAirdropFree(name, task, link)
	if err != nil {
		t.Errorf("Failed to insert AirdropFree: %v", err)
	} else {
		fmt.Printf("Inserted AirdropFree ID: %v\n", hasil)
	}
}

func TestInsertAirdropPaid(t *testing.T) {
	name := "Solv Protocol"
	task := "HOLD"
	link := "https://app.solv.finance/points"
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
