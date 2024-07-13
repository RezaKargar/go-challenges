package main

import "testing"

func TestArmstrongNumber_IsArmstrong_WhenANonArmstrongNumberGiven_ShouldReturnFalse(t *testing.T) {
	//Arrange
	var sut MyInt = 372

	//Act
	result := sut.IsArmstrong()

	//Assert
	if result {
		t.Errorf("Expected to be true, but got false")
	}
}

func TestArmstrongNumber_IsArmstrong_WhenANonArmstrongNumberGiven_ShouldReturnTrue(t *testing.T) {
	//Arrange
	var sut MyInt = 371

	//Act
	result := sut.IsArmstrong()

	//Assert
	if !result {
		t.Errorf("Expected to be true, but got false")
	}
}
