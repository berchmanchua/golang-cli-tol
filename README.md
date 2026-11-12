# golang-cli-tol 

> **⚠️ Early Stage Notice:** This project is currently in active, early-stage development. It is highly experimental, and the CLI commands, database schemas, and API contracts may change rapidly as the architecture evolves into a full Model Context Protocol (MCP) server.

`golang-cli-tol` is a lightning-fast, AI-powered terminal cheatsheet and command memory tool. Built in Go, it replaces your standard bash history with a local SQLite vector database, allowing you to save and retrieve complex terminal commands using natural language semantic search.

💡 **Fun Fact:** In Filipino slang, ***"tol"*** (short for *utol*) means a sibling, brother, or a very close friend—which is exactly what this tool is designed to be: a reliable companion for your terminal that always remembers what you forgot.

While currently functioning as a highly optimized standalone CLI tool, `golang-cli-tol` is architected as the foundational storage layer for a future Model Context Protocol (MCP) server.

## Features

* **True Semantic Search:** Uses K-Nearest Neighbors (KNN) vector math via the `sqlite-vec` C-extension to find commands based on meaning, not just exact keyword matches.
* **Multi-Provider AI Factory:** Seamlessly swap between local and cloud embedding models. Standardizes all vector output to 768 dimensions for cross-compatibility.
  * **Ollama** (Default, Local, Zero-cost)
  * **OpenAI**
  * **Google Gemini**
  * **Claude / Voyage AI**
* **Shadow Table Architecture:** Safely separates standard relational text payloads from the high-performance memory-mapped `vec0` virtual indices.
* **Frictionless Workflow:** Includes shell aliases (`tols`, `tolf`) to instantly capture your last executed command.

## Prerequisites

Because `golang-cli-tol` leverages the native `sqlite-vec` C-extension for maximum performance, you must have Go and standard C development headers installed on your system.

**For WSL / Ubuntu:**
```bash
sudo apt update
sudo apt install -y libsqlite3-dev build-essential
```

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/berchmanchua/golang-cli-tol.git
   cd golang-cli-tol
   ```
2. Download dependencies (including the CGO bindings):
   ```bash
   go mod tidy
   ```
3. Build and install the binary (requires sudo to move to `/usr/local/bin`):
   ```bash
   bash scripts/install.sh
   ```
4. Reload your shell environment to activate the aliases:
   ```bash
   source ~/.bashrc
   ```

## Configuration (AI Providers)

By default, the tool assumes you are running **Ollama** locally with the `nomic-embed-text` model. You can change this by exporting environment variables in your `~/.bashrc`.

**Note:** You *must* stick to one provider for your database. If you change providers, their geometric coordinate systems will clash. You will need to delete `tol.db` and start fresh if you swap LLMs.

### 1. Ollama (Default)
```bash
export TOL_AI_PROVIDER="ollama"
# Optional overrides:
# export OLLAMA_HOST="http://localhost:11434"
# export OLLAMA_MODEL="nomic-embed-text"
```

### 2. OpenAI
```bash
export TOL_AI_PROVIDER="openai"
export OPENAI_API_KEY="sk-..."
```

### 3. Google Gemini
```bash
export TOL_AI_PROVIDER="gemini"
export GEMINI_API_KEY="AIza..."
```

### 4. Claude (Voyage AI)
*Anthropic officially partners with Voyage AI for embeddings.*
```bash
export TOL_AI_PROVIDER="claude"
export VOYAGE_API_KEY="pa-..."
```

## Usage

### The Fast Workflow (Aliases)

The installation script injects `tols` and `tolf` into your environment. 

**Saving a command:**
Run a complex command, then immediately type `tols` followed by what it does. It will automatically pull the command from your bash history and index it.
```bash
$ kubectl delete pod stuck-pod --grace-period=0 --force
$ tols "force nuke a stuck kubernetes pod"
Saved: kubectl delete pod stuck-pod --grace-period=0 --force
```

**Finding a command:**
Ask for the command using natural language.
```bash
$ tolf "how to forcefully remove a k8s pod"

Match found:
Description: force nuke a stuck kubernetes pod
Command: kubectl delete pod stuck-pod --grace-period=0 --force
```

### The Manual CLI Workflow

If you want to manually insert a command without running it first, use the `tol` binary directly:
```bash
tol save --cmd "docker rm -vf \$(docker ps -aq)" --desc "force remove all docker containers"
tol find "clean up all my docker containers"
```
