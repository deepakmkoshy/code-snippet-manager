# code-snippet-manager README

It allows users to quickly retrieve and insert commonly used code snippets into their current projects, saving time and effort.

## Problem we are solving

Sometimes situations arise when we have to reuse code in a different project, and would have been better if they were saved somewhere and can be accessed easily.
This CLI tool aims to do exactly that.

## Description and Usage

Usage:

```
my-snippets [command]
```


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

```bash
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

## Timeline

Initially we thought to build a VS code extension tool which if enabled can create code snippets very quickly. We used Javascript yo template to generate a VS code extension. Although we were able to obtain the text which is selected in VS code and set it into clipboard, we were unable to add it as a VS code snippet.

So we tried a Golang CLI approach by using BoltDB to store the snippets locally.

In future, we plan to make the extension work and also host the golang as a backend server and facilitate code sharing between developers.
