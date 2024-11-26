#!/bin/bash
echo $(date +%F)
printf -v date '%(%Y-%m-%d)T\n' -1
