#!/bin/bash

echo "Git hooks by default are in .git/hooks,"
echo "but in this case I make scripts/.githooks/ directory, for doint this, put in the shell the following:"
echo "$ git config core.hooksPath scripts/.githooks"

# Server statements
export URL="192.168.1.65"
export PORT=":5000"

# Mongo uri
export URI_MONGO="mongodb://localhost:27017/?readPreference=primary&directConnection=true&ssl=false"

./server