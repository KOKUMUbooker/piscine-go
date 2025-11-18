#! /bin/bash

find . -name "*.sh" | cut -c 3- | sed "s/.sh//g"