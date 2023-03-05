package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

func saveSnippet(keyword string, code string) error {
	// Open the BoltDB database file for read/write access, creating it if it doesn't exist
	db, err := bolt.Open("snippets.db", 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	// Begin a read/write transaction
	err = db.Update(func(tx *bolt.Tx) error {
		// Get the snippets bucket (or create it if it doesn't exist)
		b, err := tx.CreateBucketIfNotExists([]byte("snippets"))
		if err != nil {
			return err
		}

		// Store the code snippet in the bucket with the given keyword as the key
		err = b.Put([]byte(keyword), []byte(code))
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func getSnippet(keyword string) (string, error) {
	// Open the BoltDB database file for read-only access
	db, err := bolt.Open("snippets.db", 0400, nil)
	if err != nil {
		return "", err
	}
	defer db.Close()

	var value []byte
	// Get the value of the given key from the "snippets" bucket
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("snippets"))
		if b == nil {
			return errors.New("bucket not found")
		}
		value = b.Get([]byte(keyword))
		if value == nil {
			return errors.New("snippet not found")
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	return string(value), nil
}

func deleteSnippet(keyword string) error {
	// Open the BoltDB database file for read/write access
	db, err := bolt.Open("snippets.db", 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	// Begin a read/write transaction
	err = db.Update(func(tx *bolt.Tx) error {
		// Get the snippets bucket
		b := tx.Bucket([]byte("snippets"))
		if b == nil {
			return errors.New("bucket not found")
		}

		// Delete the key-value pair with the given keyword
		err = b.Delete([]byte(keyword))
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func listSnippets() (map[string]string, error) {
	snippets := make(map[string]string)

	// Open the BoltDB database file for read-only access
	db, err := bolt.Open("snippets.db", 0600, &bolt.Options{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// View the database content
	err = db.View(func(tx *bolt.Tx) error {
		// Get the snippets bucket
		b := tx.Bucket([]byte("snippets"))
		if b == nil {
			return errors.New("bucket not found")
		}

		// Iterate over all the key-value pairs in the bucket
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			snippets[string(k)] = string(v)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return snippets, nil
}

func main() {

	var rootCmd = &cobra.Command{
		Use:   "my-snippets",
		Short: "A CLI code snippet manager",
		Long:  "A CLI snippet manager that allows you to manage your code snippets from the command line",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, World!")
		},
	}

	// Add a new sub-command
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a new code snippet",
		Long:  `Add a new code snippet to the manager.`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			key := args[0]
			code := args[1]

			saveSnippet(key, code)
			// TODO: Add code to add the snippet to the manager
			// For example, you could write the snippet to a file
			// in a specified directory or store it in a database.

			fmt.Printf("Snippet '%s' added successfully.\n", key)
		},
	}

	rootCmd.AddCommand(addCmd)

	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "get existing code snippet",
		Long:  `Get a code snippet from the manager.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			key := args[0]

			code, err := getSnippet(key)
			// TODO: Add code to add the snippet to the manager
			// For example, you could write the snippet to a file
			// in a specified directory or store it in a database.
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Printf("%s", code)
		},
	}

	rootCmd.AddCommand(getCmd)

	var rmCmd = &cobra.Command{
		Use:   "rm",
		Short: "Remove code snippet",
		Long:  `Remove code snippet from the manager.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			key := args[0]

			deleteSnippet(key)
			// TODO: Add code to add the snippet to the manager
			// For example, you could write the snippet to a file
			// in a specified directory or store it in a database.

			fmt.Printf("Snippet '%s' deleted successfully.\n", key)
		},
	}

	rootCmd.AddCommand(rmCmd)

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "list all existing code snippet",
		Long:  `List all code snippet from the manager.`,
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			snippets, err := listSnippets()
			if err != nil {
				log.Fatal(err)
			}
			for keyword, code := range snippets {
				fmt.Printf("%s : %s\n", keyword, code)
			}
		},
	}

	rootCmd.AddCommand(listCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
