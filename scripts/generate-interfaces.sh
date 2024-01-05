#!/usr/bin/env bash

FILE_NAME=${GOFILE//.go/}
STRUCT_NAME=$(echo "${FILE_NAME:0:1}" | tr '[:lower:]' '[:upper:]')"${FILE_NAME:1}"
INTERFACE_NAME="Interface${STRUCT_NAME}"
INTERFACE_FILENAME="interface${STRUCT_NAME}.go"

echo "ifacemaker: ${GOFILE} -> ${INTERFACE_FILENAME}"

ifacemaker  \
  -f "$GOFILE" \
  -s "$STRUCT_NAME" \
  -i "$INTERFACE_NAME"\
  -p "$GOPACKAGE" \
  -o "$INTERFACE_FILENAME" \
  -y "${INTERFACE_NAME} is the interface implemented by ${STRUCT_NAME}" \
  -c "code generated by ifacemaker; DO NOT EDIT."