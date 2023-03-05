# my-snippets CLI


It allows users to quickly retrieve and insert commonly used code snippets into their current projects, saving time and effort.

Usage:

```
my-snippets [command]
```

Run   ` `for usage.

### Installation

Download the [my-snippets](CLI/my-snippets) executable file and export the path in your PATH variable


1. From the directory the my-snippets is saved run the following cmd

   ```bash
   pwd
   ```
2. Run the command `export PATH=$PATH:/full/path/to/your/directory` to add the directory to the PATH variable.
```bash
 export PATH=$PATH:/full/path/to/your/directory
```

3. Verify that the directory has been added to the PATH variable by running

   ```bash
   echo PATH
   ```

### **Usage**

---



A CLI snippet manager that allows you to manage your code snippets from the command line

**Usage:**

```
my-snippets [flags]
```

```bash
my-snippets [command]
```

**Available Commands:**

`add`        Add a new code snippet

  `completion`  Generate the autocompletion script for the specified shell

`get`        get existing code snippet

 `help`        Help about any command

`list`        list all existing code snippet

  `rm`          Remove code snippet

**Flags:**

`-h, --help`   help for my-snippets

Use   `my-snippets [command] --help`  **for more information about a command.**

---
