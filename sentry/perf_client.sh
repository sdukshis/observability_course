#!/bin/bash

while true
do
    curl  "localhost:3000/slow"
    curl  "localhost:3000/fast"
done
