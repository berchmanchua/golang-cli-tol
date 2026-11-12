#!/bin/bash
set -e

echo "Building tol CLI..."
go build -o tol cmd/tol/main.go

echo "Installing to /usr/local/bin (requires sudo)..."
sudo mv tol /usr/local/bin/tol

echo "Adding aliases to ~/.bashrc..."
cat << 'EOF' >> ~/.bashrc

# Tol CLI Aliases
tols() {
   local LAST_CMD
    
    if [ -n "$ZSH_VERSION" ]; then
        # Zsh: Read directly from the internal history associative array
        LAST_CMD="${history[$((HISTCMD-1))]}"
    else
        # Bash: Use 'history -p' to securely expand '!!' (the previous command)
        # This completely avoids the subshell history-wipe issue.
        LAST_CMD=$(history -p '!!')
    fi

    # Pass it to your tool
    tol save --cmd "$LAST_CMD" --desc "$1"
    echo "Saved: $LAST_CMD"
}
tolf() {
    tol find "$@"
}
EOF

echo "Installation complete. Please run: source ~/.bashrc"


