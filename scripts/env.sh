#!/bin/bash

echo "Git hooks by default are in .git/hooks,"
echo "but in this case I make scripts/.githooks/ directory, for doint this, put in the shell the following:"
echo "$ git config core.hooksPath scripts/.githooks"

export URL="192.168.1.65"
export PORT=":5000"

./server