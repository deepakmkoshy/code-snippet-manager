package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

func saveSnippet(keyword string, code string) error {
	db, err := bolt.Open("snippets.db", 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("snippets"))
		if err != nil {
			return err
		}

		err = b.Put([]byte(keyword), []byte(code))
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func getSnippet(keyword string) (string, error) {
	db, err := bolt.Open("snippets.db", 0400, nil)
	if err != nil {
		return "", err
	}
	defer db.Close()

	var value []byte
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
	db, err := bolt.Open("snippets.db", 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("snippets"))
		if b == nil {
			return errors.New("bucket not found")
		}

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

	db, err := bolt.Open("snippets.db", 0600, &bolt.Options{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("snippets"))
		if b == nil {
			return errors.New("bucket not found")
		}

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
	scanner := bufio.NewScanner(os.Stdin)

	var rootCmd = &cobra.Command{
		Use:   "my-snippets",
		Short: "A CLI code snippet manager",
		Long:  "A CLI snippet manager that allows you to manage your code snippets from the command line",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("It allows users to quickly retrieve and insert commonly used code snippets into their current projects, saving time and effort.")
			fmt.Println("Usage:\n\tmy-snippets [flags]\n\tmy-snippets [command]")
			fmt.Println("Run 'my-snippets --help' for usage.")

		},
	}

	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a new code snippet",
		Long:  `Add a new code snippet to the manager.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			key := args[0]

			fmt.Println("Enter the code snippet (press Ctrl+D when finished):")
			var lines []string
			for scanner.Scan() {
				line := scanner.Text()
				if line == "EOF" {
					break
				}
				lines = append(lines, line)
			}
			snippet := strings.Join(lines, "\n")

			saveSnippet(key, snippet)
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

			if err != nil {
				fmt.Println("Error:", err)
				return
			} else {
				fmt.Println("\n" + code + "\n")
			}

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

			// w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight|tabwriter.Debug)

			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\nSaved Code Snippets\n    ---------    \n")

			if err != nil {
				fmt.Println(err)
				return
			}
			if len(snippets) == 0 {
				fmt.Println("No snippets found.")
				return
			}

			// Find the length of the longest key
			maxKeyLength := 3 // Minimum key length is 3 ("add" command)
			for key := range snippets {
				if len(key) > maxKeyLength {
					maxKeyLength = len(key)
				}
			}

			fmt.Printf("%-*s %s\n", maxKeyLength+10, "Key", "Snippet")
			fmt.Printf("%s %s\n", strings.Repeat("-", maxKeyLength+10), strings.Repeat("-", 50))
			for key, snippet := range snippets {
				fmt.Printf("\n")
				snippetLines := strings.Split(snippet, "\n")
				if len(snippetLines) == 1 {
					fmt.Printf("%-*s %s\n", maxKeyLength+10, key, snippetLines[0])
				} else {
					fmt.Printf("%-*s %s\n", maxKeyLength+10, key, snippetLines[0])
					for _, line := range snippetLines[1:] {
						fmt.Printf("%-*s %s\n", maxKeyLength+10, "", line)
					}
				}
			}

			// fmt.Printf("%-*s  %s\n", 30, "Key", "Snippet")
			// fmt.Printf("%-*s  %s\n", 30, "---", "-------")

			// for key, snippet := range snippets {
			// 	fmt.Printf("%-*s  %-s\n", 30, key, snippet)
			// }

			// for key, value := range snippets {
			// 	// Split multi-line snippets into lines and join with newline character
			// 	snippetLines := strings.Split(value, "|")
			// 	snippet := strings.Join(snippetLines, "\n")

			// 	fmt.Printf("%-20s %s\n", key, snippet)
			// }

			// for k, v := range snippets {
			// 	fmt.Fprintln(w, k+"\t"+v)
			// }

			// w.Flush()

			// for keyword, code := range snippets {
			// 	fmt.Printf("\n%s\t\t:\t%s\n", keyword, code)
			// }
		},
	}

	rootCmd.AddCommand(listCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
