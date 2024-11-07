#!/usr/bin/env bash

mockgen -source=types/interfaces/binary_parser.go -destination=types/testutil/binary_parser_mock.go -package=testutil