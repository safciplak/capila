#!/usr/bin/env sh


# Remove all outdated mocks since they'll be used to validate the interfaces
for file in $(find -name "*_mock.go") ; do
  # Remove the file
  rm "$file"
done
# Generate mocks for all the interfaces, this will output the mocks in the same package in the following format:
# mocks_$filename.go
mockery -all -inpkg

# Loops over each file matching the mockery format
for file in $(find -name "mock_*.go") ; do
  # Strip the mock_ from the front and lowercase the i from Interface
  newfile=${file//mock_I/i}
  # Add _mock to the end
  newfile=${newfile//.go/_mock.go}
  # Move/rename the file
  mv -v "$file" "${newfile}"
done