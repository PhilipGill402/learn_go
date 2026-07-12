package main

import (
	"fmt"
	"os"
	"strings"
	"errors"
	"bufio"
)

type Command int;
const (
	Add Command = iota
	Delete
	List
	Unknown
);

func deleteRow(pathname string, row int) error {
	file, err := os.Open(pathname);
	if (err != nil) {
		return err;
	}
	
	var lines []string;
	scanner := bufio.NewScanner(file);
	count := 0;

	for scanner.Scan() {
		line := scanner.Text();

		if (count != row) {
			lines = append(lines, line);
		}

		count++;
	}
	
	file.Close();
	err = scanner.Err();

	if (err != nil) {
		return err
	}


	file, err = os.OpenFile(pathname, os.O_WRONLY|os.O_TRUNC, 0644)
	if (err != nil) {
		return err
	}
	defer file.Close();

	for _, line := range lines {
		_, err := fmt.Fprintln(file, line);
		if (err != nil) {
			return err;
		}
	}	

	return nil;
}

func argToCommand(cmd string) (Command, error) {
	cmd = strings.ToLower(cmd);	
	switch (cmd) {
		case "add":
			return Add, nil;
		case "delete":
			return Delete, nil;
		case "list":
			return List, nil;
		default:
			return Unknown, errors.New("Unrecognized command"); 

	}
}

func handleAdd(pathname, item string) error {
	file, err := os.OpenFile(pathname, os.O_WRONLY | os.O_APPEND, 0644);
	if (err != nil) {
		return err;
	}
	defer file.Close();

	item = item + "\n";
	
	_, err = file.Write([]byte(item));
	
	return err;
}

func handleDelete(pathname, item string) error {
	file, err := os.Open(pathname);
	if (err != nil) {
		return err;
	}	

	idx := -1;
	scanner := bufio.NewScanner(file);
	count := 0;

	for scanner.Scan()  {
		line := scanner.Text();
		if (line == item) {
			idx = count;
			break;
		}

		count++;
	}

	if (idx == -1) {
		return errors.New("Item not found");
	}

	err = deleteRow(pathname, idx);
	if (err != nil) {
		return err;
	}

	return nil;
}

func handleList(pathname string) error {
	file, err := os.Open(pathname);
	if (err != nil) {
		return err;
	}
	defer file.Close();

	scanner := bufio.NewScanner(file);

	for scanner.Scan() {
		line := scanner.Text();
		fmt.Println(line);
	}

	err = scanner.Err();

	return err;
}

func main() {
	if (len(os.Args) < 2) {
		fmt.Println("Usage: todo <command> [args]");
		return;
	}

	pathname := "todo.txt"
	
	file, err := os.OpenFile(pathname, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644);
	if (err != nil) {
		fmt.Println("Error:", err);
		return;
	}
	defer file.Close();

	usr_cmd := os.Args[1];
	cmd, err := argToCommand(usr_cmd);
	if (err != nil) {
		fmt.Println("Error:", err);	
		return;
	}

	if (cmd == Add) {
		err := handleAdd(pathname, os.Args[2]);
		if (err != nil) {
			fmt.Println("Error:", err);
			return;
		}
	} else if (cmd == Delete) {
		err := handleDelete(pathname, os.Args[2]);
		if (err != nil) {
			fmt.Println("Error:", err);
			return;
		}
	} else if (cmd == List) {
		err := handleList(pathname);
		if (err != nil) {
			fmt.Println("Error:", err);
			return;
		}
	}
}
