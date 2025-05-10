package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type PhoneNumber struct {
	ID        string `json:"id"`
	Number    string `json:"number"`
	ContactID string `json:"contact_id"`
}

type Contact struct {
	ID           string        `json:"id"`
	FirstName    string        `json:"first_name"`
	LastName     string        `json:"last_name"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
}

type SearchResponse struct {
	Message string    `json:"message"`
	Data    []Contact `json:"data"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <baseurl>")
		return
	}
	baseURL := os.Args[1]
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Contact Client ---")
		fmt.Println("1. Create contact")
		fmt.Println("2. Search contacts")
		fmt.Println("3. Update contact")
		fmt.Println("4. Exit")
		fmt.Print("Select option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			createContact(baseURL, reader)
		case "2":
			searchContacts(baseURL, reader)
		case "3":
			updateContact(baseURL, reader)
		case "4":
			fmt.Println("Exiting client.")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func createContact(baseURL string, reader *bufio.Reader) {
	fmt.Print("First name: ")
	firstName, _ := reader.ReadString('\n')
	firstName = strings.TrimSpace(firstName)

	fmt.Print("Last name: ")
	lastName, _ := reader.ReadString('\n')
	lastName = strings.TrimSpace(lastName)

	var phoneNumbers []PhoneNumber
	for {
		fmt.Print("Phone number (empty to finish): ")
		phone, _ := reader.ReadString('\n')
		phone = strings.TrimSpace(phone)
		if phone == "" {
			break
		}
		phoneNumbers = append(phoneNumbers, PhoneNumber{Number: phone})
	}
	contact := Contact{
		FirstName:    firstName,
		LastName:     lastName,
		PhoneNumbers: phoneNumbers,
	}

	sendPost(baseURL+"/contact", contact)
}

func searchContacts(baseURL string, reader *bufio.Reader) {
	fmt.Print("Enter search param (name, last name or phone): ")
	param, _ := reader.ReadString('\n')
	param = strings.TrimSpace(param)

	resp, err := http.Get("http://" + baseURL + "/contact/search?param=" + param)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	var result SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	fmt.Printf("HTTP %d - %s\n", resp.StatusCode, result.Message)
	for _, c := range result.Data {
		fmt.Printf("ID: %s, Name: %s %s\n", c.ID, c.FirstName, c.LastName)
		for _, p := range c.PhoneNumbers {
			fmt.Printf("   Phone ID: %s, Number: %s\n", p.ID, p.Number)
		}
	}
}

func updateContact(baseURL string, reader *bufio.Reader) {
	fmt.Print("Enter contact ID: ")
	contactID, _ := reader.ReadString('\n')
	contactID = strings.TrimSpace(contactID)

	fmt.Print("Enter first name: ")
	firstName, _ := reader.ReadString('\n')
	firstName = strings.TrimSpace(firstName)

	fmt.Print("Enter last name: ")
	lastName, _ := reader.ReadString('\n')
	lastName = strings.TrimSpace(lastName)

	var phoneNumbers []map[string]string
	for {
		fmt.Print("Enter phone ID (or leave empty to stop): ")
		phoneID, _ := reader.ReadString('\n')
		phoneID = strings.TrimSpace(phoneID)
		if phoneID == "" {
			break
		}

		fmt.Print("Enter phone number: ")
		number, _ := reader.ReadString('\n')
		number = strings.TrimSpace(number)

		phoneNumbers = append(phoneNumbers, map[string]string{
			"phone_id": phoneID,
			"number":   number,
		})
	}

	payload := map[string]interface{}{
		"contact_id":    contactID,
		"first_name":    firstName,
		"last_name":     lastName,
		"phone_numbers": phoneNumbers,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	url := "http://" + baseURL + "/contact/update"
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Printf("HTTP %d - %v\n", resp.StatusCode, result["message"])
}

func sendPost(url string, payload any) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	resp, err := http.Post("http://"+url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	var response SearchResponse
	json.NewDecoder(resp.Body).Decode(&response)

	fmt.Printf("Response: %s (HTTP %d)\n", response.Message, resp.StatusCode)
}
