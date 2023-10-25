DIRECTORY_COUNT=$(find . -type d | wc -l)
FILE_COUNT=$(find . -type f | wc -l)
echo $(($DIRECTORY_COUNT + $FILE_COUNT))