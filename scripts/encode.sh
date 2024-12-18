#!/bin/bash

BASE64=$(base64 -w 0 $1)
echo '<img alt="fill" src="data:image/jpeg;base64,'"${BASE64}"'">'
